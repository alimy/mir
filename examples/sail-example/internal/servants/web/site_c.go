// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package web

import (
	api "github.com/alimy/mir/sail-example/v5/auto/api/v3"
	"github.com/alimy/mir/sail-example/v5/internal/servants/base"
)

type siteSrvC struct {
	base.BaseSrv
	api.UnimplementedSiteServant
}

func newSiteSrvC() api.Site {
	return &siteSrvC{}
}
