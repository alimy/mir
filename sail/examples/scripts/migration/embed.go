// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

//go:build migration
// +build migration

package migration

import (
	"embed"
)

//go:embed **/*
var Files embed.FS
