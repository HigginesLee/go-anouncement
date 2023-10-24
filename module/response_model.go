package module

type Article struct {
	ID         int    `json:"id" json:"itemId"`
	Title      string `json:"title" json:"docPageName"`
	SubTitle   string `json:"subtitle"`
	CreateTime int64  `json:"createTime" json:"onlineTime"`
	UpdateTime int64  `json:"UpdateTime"`
	Status     int32  `json:"status"`
}

type JingDongArticle struct {
	ID         int    `json:"id"`
	Title      string `json:"articleTitle"`
	CreateTime string `json:"created"`
	UpdateTime string `json:"modified"`
}

type DouDianJsonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ArticleData []Article `json:"articles"`
		Total       int       `json:"total"`
	} `json:"data"`
}

type XhongshuJsonResponse struct {
	Code      int32 `json:"error_code"`
	IsSuccess bool  `json:"success"`
	Data      struct {
		ArticleData []Article `json:"announcementDetails"`
		Total       int       `json:"total"`
	} `json:"data"`
}

type KuaishouJsonResponse struct {
	Code    int32  `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Total       int       `json:"totalCount"`
		ArticleData []Article `json:"list"`
	} `json:"data"`
}

type JingDongJsonResponse struct {
	Code      int32  `json:"code"`
	Message   string `json:"message"`
	IsSuccess bool   `json:"success"`
	Data      struct {
		Total       int               `json:"totalArticles"`
		ArticleData []JingDongArticle `json:"josCmsArticle"`
	} `json:"data"`
}

type WeimobJsonResponse struct {
	Code    int8   `json:"errcode"`
	Message string `json:"errmsg"`
	Data    struct {
		Total       int       `json:"total"`
		ArticleData []Article `json:"data"`
	} `json:"data"`
}
