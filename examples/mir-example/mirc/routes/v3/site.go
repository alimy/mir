// Copyright 2025 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package v3

import (
	"github.com/alimy/mir/mir-example/v5/model"
	. "github.com/alimy/mir/v5"
)

// Site site v3 interface info
type Site struct {
	Schema           `mir:"v3"`
	Index            func(Get)                                               `mir:"index"`
	Articles         func(Get)                                               `mir:"articles/:category"`
	NextTweets       func(Any, model.TweetsReq) model.TweetsResp             `mir:"tweets/next" render:"jsonp"`
	PrevTweets       func(Post, Get, Head, model.TweetsReq) model.TweetsResp `mir:"tweets/prev"`
	Login            func(Post, model.LoginReq) model.LoginResp              `mir:"user/login"`
	Logout           func(Post)                                              `mir:"user/logout"`
	Assets           func(Get, Context)                                      `mir:"assets/:name"`
	Statics          func(Get, Context, model.LoginReq)                      `mir:"statics/:name"`
	AnyStaticks      func(Any, Context)                                      `mir:"anystaticks/:name"`
	ManyResources    func(Get, Head, Options, Context)                       `mir:"resources/:name"`
	MultiAttachments func(Get, Head, Options, Chain, Context)                `mir:"attachments/:name"`
}
