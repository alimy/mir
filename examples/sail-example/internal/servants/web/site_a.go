// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package web

import (
	api "github.com/alimy/mir/sail-example/v5/auto/api/v1"
	"github.com/alimy/mir/sail-example/v5/internal/servants/base"
)

type siteSrvA struct {
	base.BaseSrv
	api.UnimplementedSiteServant
}

type siteChainA struct {
	api.UnimplementedSiteChain
}

func newSiteSrvA() api.Site {
	return &siteSrvA{}
}

func newSiteChainA() api.SiteChain {
	return &siteChainA{}
}
