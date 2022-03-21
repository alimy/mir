package servants

import (
	v1 "github.com/alimy/mir-example/mirc/auto/api/v1"
)

type siteV1Srv struct {
	v1.UnimplementedSiteServant
}

func newSiteV1Srv() v1.Site {
	return &siteV1Srv{}
}
