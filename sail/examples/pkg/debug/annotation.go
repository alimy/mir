// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package debug

import (
	"github.com/sirupsen/logrus"
)

func TODO() {
	logrus.Fatalln("in todo progress")
}

func NotImplemented() {
	logrus.Fatalln("not implemented")
}
