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
