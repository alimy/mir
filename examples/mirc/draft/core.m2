//mir:syntax v0.1-alpha.1

package core

message ArticleReq {
     Content string  `json:"content"`
}

message ArticleRes {
     Code   int      `json:"code"`
     Msg    string   `json:"msg"`
     PostId string   `json:"postid"`
}

message IndexHead {
     LastTime string `param:"lastTime"`
}
