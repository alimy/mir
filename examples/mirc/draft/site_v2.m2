//mir:syntax v0.1-alpha.1

package v2

import core

// Site v2 service
service Site(group: v2, chain: _) {
    Index(param core.indexParam)                         `get:"/index"`
    Articles()                                           `get, post:"/articles/:category/"`
    Category()                                           `get:"/category/"`
    PostArticle(article core.ArticleReq) core.ArticleRes `post:"/articles/:category/"`
}
