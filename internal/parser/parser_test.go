package parser

import (
	"testing"

	"github.com/alimy/mir/v2"
)

type site struct {
	Chain    mir.Chain `mir:"-"`
	Index    mir.Get   `mir:"/index/"`
	Articles mir.Get   `mir:"/articles/:category/"`
}

type siteV1 struct {
	Chain    mir.Chain `mir:"-"`
	Group    mir.Group `mir:"v1"`
	Index    mir.Get   `mir:"/index/"`
	Articles mir.Get   `mir:"/articles/:category/"`
}

type siteV2 struct {
	Group    mir.Group `mir:"v2"`
	Index    mir.Get   `mir:"/index/"`
	Articles mir.Get   `mir:"/articles/:category/"`
	Category mir.Get   `mir:"/category/"`
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
