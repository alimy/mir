// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	"github.com/alimy/mir/mir-example/v4/mirc/auto/api"
)

type siteSrv struct {
	baseSrv
	api.UnimplementedSiteServant
}

func newSiteSrv() api.Site {
	return &siteSrv{}
}
