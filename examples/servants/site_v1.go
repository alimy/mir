package servants

import (
	api "github.com/alimy/mir-example/mirc/auto/api/v1"
)

type siteV1Srv struct {
	api.UnimplementedSiteServant
}

type siteV1Binding struct {
	api.UnimplementedSiteBinding
}

type siteV1Render struct {
	api.UnimplementedSiteRender
}

func newSiteV1Srv() api.Site {
	return &siteV1Srv{}
}

func newSiteV1Binding() api.SiteBinding {
	return &siteV1Binding{}
}

func newSiteV1Render() api.SiteRender {
	return &siteV1Render{}
}
