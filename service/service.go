package service

import (
	"encoding/json"
	"fmt"
	"go_announce/module"
	"io"
	"log"
	"net/http"
)

// 抖店平台公告

type Article struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	SubTitle   string `json:"subtitle"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"UpdateTime"`
	Status     int    `json:"status"`
}

type JsonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		ArticleData []Article `json:"articles"`
		Total       int       `json:"total"`
	} `json:"data"`
}

func DouDianAnounceMent() {
	url := "https://op.jinritemai.com/doc/external/open/queryDocArticleList?pageIndex=0&pageSize=10&status=1&dirId=5&orderType=3"
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("抖店平台公告获取失败, err: %s", err)
		return
	}
	defer response.Body.Close()
	var respJson JsonResponse
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &respJson)
	if err != nil {
		log.Fatalf("抖店平台公告解析失败, err: %s", err)
		return
	}
	total := respJson.Data.Total
	first_articles := respJson.Data.ArticleData[0]
	// 比较本地数据库存储的total和first_article时间, 如果total对应不上或更新时间对应不上则进行更新并发出提醒, 如果total相等
	anouncement, _ := module.QueryAnouncement("Doudian")
	dbTotal, dbLastTime := anouncement.Total, anouncement.LastTime
	if total != dbTotal || first_articles.UpdateTime != dbLastTime {
		var a module.AnounceModule
		a.Platform = "Doudian"
		a.Total = total
		a.LastTime = first_articles.UpdateTime
		fmt.Printf("%v", a)
		err := a.SaveAnouncement()
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}
}

//小红书平台公告

func XiaohongshuAnouncement() {

}

// 淘宝平台公告

func TaobaoAnouncement() {
	//url := ""
}

// 1688平台公告

func AlibabaAnouncement() {

}

// 国际站平台公告

func AliAnouncement() {

}

// 微盟平台公告

func WeiMobAnouncement() {

}

// 有赞平台公告

func YouzanAnouncement() {

}

// 拼多多平台公告

func PinduoduoAnouncement() {

}

// 京东平台公告

func JingDongAnouncement() {

}

// 快手平台公告

func KuaishouAnouncement() {

}

// 爱采购平台公告

func AiCaiGouAnouncement() {

}

// 千川平台公告

func QianchuanAnouncement() {

}
