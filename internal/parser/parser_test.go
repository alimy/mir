// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package parser

import (
	"testing"

	"github.com/alimy/mir/v3/internal/core/tag"
)

type site struct {
	Chain    tag.Chain `mir:"-"`
	Index    tag.Get   `mir:"/index/"`
	Articles tag.Get   `mir:"/articles/:category/"`
}

type siteV1 struct {
	Chain    tag.Chain `mir:"-"`
	Group    tag.Group `mir:"v1"`
	Index    tag.Get   `mir:"/index/"`
	Articles tag.Get   `mir:"/articles/:category/"`
}

type siteV2 struct {
	Group    tag.Group `mir:"v2"`
	Index    tag.Get   `mir:"/index/"`
	Articles tag.Get   `mir:"/articles/:category/"`
	Category tag.Get   `mir:"/category/"`
}

func TestMirParser_Parse(t *testing.T) {
	p := &mirParser{tagName: defaultTag}

	entries := []interface{}{
		new(site),
		new(siteV1),
		new(siteV2),
	}
	ds, err := p.Parse(entries)
	if err != nil {
		t.Error("want nil error but not")
	}
	if len(ds) != 3 {
		t.Fatal("want 3 item but not")
	}

	iface, exist := ds.Get("")
	if !exist || len(iface) != 1 {
		t.Error("want a correct iface but not")
	}
	site := iface["site"]
	if site == nil || len(site.Fields) != 2 {
		t.Error("want a correct iface but not")
	}

	iface, exist = ds.Get("v1")
	if !exist || len(iface) != 1 {
		t.Error("want a correct iface but not")
	}
	site = iface["siteV1"]
	if site == nil || len(site.Fields) != 2 {
		t.Error("want a correct iface but not")
	}

	iface, exist = ds.Get("v2")
	if !exist || len(iface) != 1 {
		t.Error("want a correct iface but not")
	}
	site = iface["siteV2"]
	if site == nil || len(site.Fields) != 3 {
		t.Error("want a correct iface but not")
	}
}
