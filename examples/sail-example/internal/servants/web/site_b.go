// Copyright 2024 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package web

import (
	api "github.com/alimy/mir/sail-example/v4/auto/api/v2"
	"github.com/alimy/mir/sail-example/v4/internal/servants/base"
)

type siteSrvB struct {
	base.BaseSrv
	api.UnimplementedSiteServant
}

func newSiteSrvB() api.Site {
	return &siteSrvB{}
}
