// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	api "github.com/alimy/mir-example/v3/mirc/auto/api/v1"
)

type siteV1Srv struct {
	api.UnimplementedSiteServant
}

type siteV1Render struct {
	*api.UnimplementedSiteRender
}

func newSiteV1Srv() api.Site {
	return &siteV1Srv{}
}

func newSiteV1Render() api.SiteRender {
	return &siteV1Render{
		UnimplementedSiteRender: &api.UnimplementedSiteRender{
			RenderAny: renderAny,
		},
	}
}
