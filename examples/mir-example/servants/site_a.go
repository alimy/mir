// Copyright 2025 Michael Li <alimy@niubiu.com>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	api "github.com/alimy/mir/mir-example/v5/mirc/auto/api/v1"
)

type siteSrvA struct {
	baseSrv

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
