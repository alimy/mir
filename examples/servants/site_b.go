// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	api "github.com/alimy/mir/examples/v4/mirc/auto/api/v2"
)

type siteSrvB struct {
	baseSrv

	api.UnimplementedSiteServant
}

func newSiteSrvB() api.Site {
	return &siteSrvB{}
}
