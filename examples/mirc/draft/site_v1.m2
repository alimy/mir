//mir:syntax v0.1-alpha.1

package v1

message ArticleReq {
     Content string  `json:"content"`
}

message ArticleRes {
    Code   int       `json:"code"`
    Msg    string    `json:"msg"`
    PostId string    `json:"postid"`
}

// Site v1 service
service Site(group: v1) {
    Index()                                    `get:"/index/"`
    Articles()                                 `get:"/articles/:category/"`
    Category()                                 `get:"/category/"`
    PostArticle(article ArticleReq) ArticleRes `post:"/articles/:category/"`
}
