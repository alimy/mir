// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"github.com/alimy/mir-example/v3/mirc/auto/api"
)

type siteSrv struct {
	api.UnimplementedSiteServant
}

type siteBinding struct {
	*api.UnimplementedSiteBinding
}

type siteRender struct {
	*api.UnimplementedSiteRender
}

func newSiteSrv() api.Site {
	return &siteSrv{}
}

func newSiteBinding() api.SiteBinding {
	return &siteBinding{
		UnimplementedSiteBinding: &api.UnimplementedSiteBinding{
			BindAny: bindAny,
		},
	}
}

func newSiteRender() api.SiteRender {
	return &siteRender{
		UnimplementedSiteRender: &api.UnimplementedSiteRender{
			RenderAny: renderAny,
		},
	}
}
