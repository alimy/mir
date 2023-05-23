// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	api "github.com/alimy/mir-example/v3/mirc/auto/api/v1"
)

type siteSrvA struct {
	api.UnimplementedSiteServant
}

type siteChainA struct {
	api.UnimplementedSiteChain
}

type siteRenderA struct {
	*api.UnimplementedSiteRender
}

func newSiteSrvA() api.Site {
	return &siteSrvA{}
}

func newSiteRenderA() api.SiteRender {
	return &siteRenderA{
		UnimplementedSiteRender: &api.UnimplementedSiteRender{
			RenderAny: renderAny,
		},
	}
}

func newSiteChainA() api.SiteChain {
	return &siteChainA{}
}
