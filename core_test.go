// Copyright 2019 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package mir_test

import (
	. "github.com/alimy/mir/v2"
)

// entry mux style URN entry
type muxEntry struct {
	Chain      Chain   `mir:"-"`
	Group      Group   `mir:"v1"`
	DotHandler Any     `mir:"/dot/handler/#."`
	get        Get     `mir:"/get/"`
	put        Put     `mir:"/put/"`
	post       Post    `mir:"/post/"`
	delete     Delete  `mir:"/delete/"`
	head       Head    `mir:"/head/"`
	patch      Patch   `mir:"/patch/"`
	trace      Trace   `mir:"/trace/"`
	connect    Connect `mir:"/connect/"`
	options    Options `mir:"/options/"`
	any        Any     `mir:"/any/"`
	alias      Get     `mir:"/alias/#GetAlias"`
	chainFunc1 Get     `mir:"/chainfunc1/#-ChainFunc"`
	chainFunc2 Get     `mir:"/chainfunc2/#GetChainFunc2&ChainFunc"`
	query      Get     `mir:"/query/?filter={filter}"`
	full       Get     `mir:"//{subdomain}.domain.com:8013/full/{other}/{id:[0-9]+}?filter={filter}&foo=bar&index={index:[0-9]+}#GetFull"`
}

// ginEntry gin,echo,httrouter style URN entry
type ginEntry struct {
	Chain      Chain   `mir:"-"`
	DotHandler Any     `mir:"/dot/handler/#."`
	group      Group   `mir:"v1"`
	get        Get     `mir:"/get/"`
	put        Put     `mir:"/put/"`
	post       Post    `mir:"/post/"`
	delete     Delete  `mir:"/delete/"`
	head       Head    `mir:"/head/"`
	patch      Patch   `mir:"/patch/"`
	trace      Trace   `mir:"/trace/"`
	connect    Connect `mir:"/connect/"`
	options    Options `mir:"/options/"`
	any        Any     `mir:"/any/"`
	alias      Get     `mir:"/alias/#GetAlias"`
	chainFunc1 Get     `mir:"/chainfunc1/#-ChainFunc"`
	chainFunc2 Get     `mir:"/chainfunc2/#GetChainFunc2&ChainFunc"`
	full       Get     `mir:"/full/:other/:name#GetFull"`
}

// irisEntry iris style URN entry
type irisEntry struct {
	Chain      Chain   `mir:"-"`
	DotHandler Any     `mir:"/dot/handler/#."`
	group      Group   `mir:"v1"`
	get        Get     `mir:"/get/"`
	put        Put     `mir:"/put/"`
	post       Post    `mir:"/post/"`
	delete     Delete  `mir:"/delete/"`
	head       Head    `mir:"/head/"`
	patch      Patch   `mir:"/patch/"`
	trace      Trace   `mir:"/trace/"`
	connect    Connect `mir:"/connect/"`
	options    Options `mir:"/options/"`
	any        Any     `mir:"/any/"`
	alias      Get     `mir:"/alias/#GetAlias"`
	chainFunc1 Get     `mir:"/chainfunc1/#-ChainFunc"`
	chainFunc2 Get     `mir:"/chainfunc2/#GetChainFunc2&ChainFunc"`
	full       Get     `mir:"/full/{other:string}/{name:string range(1,200) else 400}#GetFull"`
}

// urbanEntry used to test custom mir tag name entry
type urbanEntry struct {
	group Group `urban:"v1"`
	get   Get   `urban:"/get/"`
}

// errGroupEntry used to test direct assign group info to no exported group filed occurs error
type errGroupEntry struct {
	group Group `mir:"-"`
	get   Get   `mir:"/get/"`
}

// errChainEntry used to test assign exported chain field occurs error
type errChainEntry struct {
	chain Chain `mir:"-"`
	get   Get   `mir:"/get/"`
}

// errNoMethodEntry used to test no method define occurs error
type errNoMethodEntry struct {
	get Get `mir:"/get/"`
}

// handlerFunc fake handler function
type handlerFunc func() string

// chains fake chain of middleware
type chains []func() string

func (*muxEntry) Get() string {
	return "/get/"
}

func (*muxEntry) Put() string {
	return "/put/"
}

func (*muxEntry) Post() string {
	return "/post/"
}

func (*muxEntry) Delete() string {
	return "/delete/"
}

func (*muxEntry) Head() string {
	return "/head/"
}

func (*muxEntry) Patch() string {
	return "/patch/"
}

func (*muxEntry) Trace() string {
	return "/trace/"
}

func (*muxEntry) Connect() string {
	return "/connect/"
}

func (*muxEntry) Options() string {
	return "/options/"
}

func (*muxEntry) Any() string {
	return "/any/"
}

func (*muxEntry) GetAlias() string {
	return "/alias/"
}

func (*muxEntry) ChainFunc1() string {
	return "/chainfunc1/"
}

func (*muxEntry) GetChainFunc2() string {
	return "/chainfunc2/"
}

func (*muxEntry) ChainFunc() string {
	return "chainFunc"
}

func (*muxEntry) Query() string {
	return "/query/"
}

func (*muxEntry) GetFull() string {
	return "/full/{other}/{id:[0-9]+}"
}

func (*ginEntry) Get() string {
	return "/get/"
}

func (*ginEntry) Put() string {
	return "/put/"
}

func (*ginEntry) Post() string {
	return "/post/"
}

func (*ginEntry) Delete() string {
	return "/delete/"
}

func (*ginEntry) Head() string {
	return "/head/"
}

func (*ginEntry) Patch() string {
	return "/patch/"
}

func (*ginEntry) Trace() string {
	return "/trace/"
}

func (*ginEntry) Connect() string {
	return "/connect/"
}

func (*ginEntry) Options() string {
	return "/options/"
}

func (*ginEntry) Any() string {
	return "/any/"
}

func (*ginEntry) GetAlias() string {
	return "/alias/"
}

func (*ginEntry) ChainFunc1() string {
	return "/chainfunc1/"
}

func (*ginEntry) GetChainFunc2() string {
	return "/chainfunc2/"
}

func (*ginEntry) ChainFunc() string {
	return "chainFunc"
}

func (*ginEntry) Query() string {
	return "/query"
}

func (*ginEntry) GetFull() string {
	return "/full/:other/:name"
}

func (*irisEntry) Get() string {
	return "/get/"
}

func (*irisEntry) Put() string {
	return "/put/"
}

func (*irisEntry) Post() string {
	return "/post/"
}

func (*irisEntry) Delete() string {
	return "/delete/"
}

func (*irisEntry) Head() string {
	return "/head/"
}

func (*irisEntry) Patch() string {
	return "/patch/"
}

func (*irisEntry) Trace() string {
	return "/trace/"
}

func (*irisEntry) Connect() string {
	return "/connect/"
}

func (*irisEntry) Options() string {
	return "/options/"
}

func (*irisEntry) Any() string {
	return "/any/"
}

func (*irisEntry) GetAlias() string {
	return "/alias/"
}

func (*irisEntry) ChainFunc1() string {
	return "/chainfunc1/"
}

func (*irisEntry) GetChainFunc2() string {
	return "/chainfunc2/"
}

func (*irisEntry) ChainFunc() string {
	return "chainFunc"
}

func (*irisEntry) Query() string {
	return "/query"
}

func (*irisEntry) GetFull() string {
	return "/full/{other:string}/{name:string range(1,200) else 400}"
}

func (*urbanEntry) Get() string {
	return "/get/"
}

func (*errChainEntry) Get() string {
	return "/get/"
}

func (*errGroupEntry) Get() string {
	return "/get/"
}

func DotHandler() string {
	return "/dot/handler/"
}

func pingChain() string {
	return "pingChain"
}

func pongChain() string {
	return "pongChain"
}

func mirChains() chains {
	return chains{
		pingChain,
		pongChain,
	}
}
