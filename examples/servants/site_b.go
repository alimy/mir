// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	api "github.com/alimy/mir-example/v3/mirc/auto/api/v2"
)

type siteSrvB struct {
	api.UnimplementedSiteServant
}

type siteChainB struct {
	api.UnimplementedSiteChain
}

type siteBindingB struct {
	*api.UnimplementedSiteBinding
}

type siteRenderB struct {
	*api.UnimplementedSiteRender
}

func newSiteSrvB() api.Site {
	return &siteSrvB{}
}

func newSiteChainB() api.SiteChain {
	return &siteChainB{}
}

func newSiteBindingB() api.SiteBinding {
	return &siteBindingB{
		UnimplementedSiteBinding: &api.UnimplementedSiteBinding{
			BindAny: bindAny,
		},
	}
}

func newSiteRenderB() api.SiteRender {
	return &siteRenderB{
		UnimplementedSiteRender: &api.UnimplementedSiteRender{
			RenderAny: renderAny,
		},
	}
}
