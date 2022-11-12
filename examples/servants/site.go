package servants

import (
	"github.com/alimy/mir-example/mirc/auto/api"
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
