package servants

import "github.com/alimy/mir-example/mirc/auto/api"

type siteSrv struct {
	api.UnimplementedSiteServant
}

func newSiteSrv() api.Site {
	return &siteSrv{}
}
