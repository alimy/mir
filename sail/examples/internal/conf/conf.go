// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

// package conf contains configure logic

package conf

import (
	"log"

	"github.com/alimy/tryst/cfg"
)

var (
	DatabaseSetting   *databaseConf
	MysqlSetting      *mysqlConf
	PostgresSetting   *postgresConf
	Sqlite3Setting    *sqlite3Conf
	WebServerSetting  *httpServerConf
	BotServerSetting  *httpServerConf
	DocsServerSetting *httpServerConf
	AppSetting        *appConf
)

func setupSetting(suite []string, noDefault bool) error {
	vp, err := newViper()
	if err != nil {
		return err
	}

	// initialize features configure
	ss, kv := featuresInfoFrom(vp, "Features")
	cfg.Initial(ss, kv)
	if len(suite) > 0 {
		cfg.Use(suite, noDefault)
	}

	objects := map[string]any{
		"App":        &AppSetting,
		"WebServer":  &WebServerSetting,
		"BotServer":  &BotServerSetting,
		"DocsServer": &DocsServerSetting,
		"Database":   &DatabaseSetting,
		"MySQL":      &MysqlSetting,
		"Postgres":   &PostgresSetting,
		"Sqlite3":    &Sqlite3Setting,
	}
	for k, v := range objects {
		err := vp.UnmarshalKey(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func Initial(suite []string, noDefault bool) {
	err := setupSetting(suite, noDefault)
	if err != nil {
		log.Fatalf("init.setupSetting err: %s", err)
	}
}

func RunMode() string {
	return AppSetting.RunMode
}
