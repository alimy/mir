package servants

import (
	api "github.com/alimy/mir-example/mirc/auto/api/v2"
)

type siteV2Srv struct {
	api.UnimplementedSiteServant
}

type siteV2Binding struct {
	api.UnimplementedSiteBinding
}

type siteV2Render struct {
	api.UnimplementedSiteRender
}

func newSiteV2Srv() api.Site {
	return &siteV2Srv{}
}

func newSiteV2Binding() api.SiteBinding {
	return &siteV2Binding{}
}

func newSiteV2Render() api.SiteRender {
	return &siteV2Render{}
}
