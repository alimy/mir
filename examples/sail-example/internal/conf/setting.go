// Copyright 2025 Michael Li <alimy@niubiu.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package conf

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

//go:embed config.yaml
var configBytes []byte

type appConf struct {
	RunMode string
}

type httpServerConf struct {
	Host         string
	Port         int16
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type databaseConf struct {
	TablePrefix string
	LogLevel    string
}

type mysqlConf struct {
	UserName     string
	Password     string
	Host         string
	DBName       string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type postgresConf map[string]string

type sqlite3Conf struct {
	Path string
}

func (s *httpServerConf) MyAddr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

func (s *httpServerConf) MyReadTimeout() time.Duration {
	return s.ReadTimeout * time.Second
}

func (s *httpServerConf) MyWriteTimeout() time.Duration {
	return s.WriteTimeout * time.Second
}

func (s *mysqlConf) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		s.UserName,
		s.Password,
		s.Host,
		s.DBName,
		s.Charset,
		s.ParseTime,
	)
}

func (s postgresConf) Dsn() string {
	var params []string
	for k, v := range s {
		if len(v) == 0 {
			continue
		}
		lk := strings.ToLower(k)
		tv := strings.Trim(v, " ")
		switch lk {
		case "schema":
			params = append(params, "search_path="+tv)
		case "applicationname":
			params = append(params, "application_name="+tv)
		default:
			params = append(params, lk+"="+tv)
		}
	}
	return strings.Join(params, " ")
}

func (s *sqlite3Conf) Dsn(driverName string) string {
	pragmas := "_foreign_keys=1&_journal_mode=WAL&_synchronous=NORMAL&_busy_timeout=8000"
	if driverName == "sqlite" {
		pragmas = "_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)&_pragma=synchronous(NORMAL)&_pragma=busy_timeout(8000)&_pragma=journal_size_limit(100000000)"
	}
	return fmt.Sprintf("file:%s?%s", s.Path, pragmas)
}

func newViper() (*viper.Viper, error) {
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath(".")
	vp.AddConfigPath("custom/")
	vp.SetConfigType("yaml")
	err := vp.ReadConfig(bytes.NewReader(configBytes))
	if err != nil {
		return nil, err
	}
	if err = vp.MergeInConfig(); err != nil {
		return nil, err
	}
	return vp, nil
}

func featuresInfoFrom(vp *viper.Viper, k string) (map[string][]string, map[string]string) {
	sub := vp.Sub(k)
	keys := sub.AllKeys()

	suites := make(map[string][]string)
	kv := make(map[string]string, len(keys))
	for _, key := range sub.AllKeys() {
		val := sub.Get(key)
		switch v := val.(type) {
		case string:
			kv[key] = v
		case []any:
			suites[key] = sub.GetStringSlice(key)
		}
	}
	return suites, kv
}
