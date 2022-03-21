package servants

import (
	v2 "github.com/alimy/mir-example/mirc/auto/api/v2"
)

type siteV2Srv struct {
	v2.UnimplementedSiteServant
}

func newSiteV2Srv() v2.Site {
	return &siteV2Srv{}
}
