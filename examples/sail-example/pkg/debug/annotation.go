// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package debug

import (
	"log"
	"log/slog"
)

func TODO() {
	slog.Warn("in todo progress")
}

func NotImplemented() {
	log.Fatalln("not implemented")
}
