// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	api "github.com/alimy/mir-example/v3/mirc/auto/api/v3"
)

type siteSrvC struct {
	api.UnimplementedSiteServant
}

type siteBindingC struct {
	*api.UnimplementedSiteBinding
}

type siteRenderC struct {
	*api.UnimplementedSiteRender
}

func newSiteSrvC() api.Site {
	return &siteSrvC{}
}

func newSiteBindingC() api.SiteBinding {
	return &siteBindingC{
		UnimplementedSiteBinding: &api.UnimplementedSiteBinding{
			BindAny: bindAny,
		},
	}
}

func newSiteRenderC() api.SiteRender {
	return &siteRenderC{
		UnimplementedSiteRender: &api.UnimplementedSiteRender{
			RenderAny: renderAny,
		},
	}
}
