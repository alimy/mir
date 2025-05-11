// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package servants

import (
	api "github.com/alimy/mir/mir-example/v5/mirc/auto/api/v3"
)

type siteSrvC struct {
	baseSrv

	api.UnimplementedSiteServant
}

func newSiteSrvC() api.Site {
	return &siteSrvC{}
}
