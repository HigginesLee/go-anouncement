package service

import (
	"bytes"
	"encoding/json"
	"go_announce/module"
	"go_announce/utils"
	"io"
	"log"
	"net/http"
	"time"
)

// 抖店平台公告

func DouDianAnounceMent() {
	url := "https://op.jinritemai.com/doc/external/open/queryDocArticleList?pageIndex=0&pageSize=10&status=1&dirId=5&orderType=3"
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("抖店平台公告获取失败, err: %s", err)
		return
	}
	defer response.Body.Close()
	var respJson module.DouDianJsonResponse
	body, err := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &respJson)
	if err != nil {
		log.Fatalf("抖店平台公告解析失败, err: %s", err)
		return
	}
	total := respJson.Data.Total
	firstArticles := respJson.Data.ArticleData[0]
	// 比较本地数据库存储的total和first_article时间, 如果total对应不上或更新时间对应不上则进行更新并发出提醒, 如果total相等
	anouncement, _ := module.QueryAnouncement("Doudian")
	dbTotal, dbLastTime := anouncement.Total, anouncement.LastTime
	if total != dbTotal || firstArticles.UpdateTime != dbLastTime {
		// 更新数据库
		var a module.AnounceModule
		a.Platform = "Doudian"
		a.Total = total
		a.LastTime = firstArticles.UpdateTime
		err := a.SaveAnouncement()
		utils.CheckErr(err)
		// 钉钉通知平台公告有更新
		err = utils.DingCli.SendLinkMessage("抖店平台公告更新", "抖店平台公告有更新, 请点击查看", "",
			"https://op.jinritemai.com/docs/notice-docs/5")
		utils.CheckErr(err)
	}
}

//小红书平台公告

func XiaohongshuAnouncement() {
	url := "https://open.xiaohongshu.com/api/announcement/getAllAnnouncementDetailNew"
	paramsMap := map[string]int8{"pageNo": 1, "pageSize": 20}
	params, _ := json.Marshal(paramsMap)
	resp, err := http.Post(url, "application/json", bytes.NewReader(params))
	utils.CheckErr(err)
	defer resp.Body.Close()
	var respJson module.XhongshuJsonResponse
	body, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &respJson)
	utils.CheckErr(err)
	total := respJson.Data.Total
	firstArticles := respJson.Data.ArticleData[0]
	announcement, _ := module.QueryAnouncement("Xhongshu")
	dbTotal, dbLastTime := announcement.Total, announcement.LastTime
	if total != dbTotal || dbLastTime != firstArticles.UpdateTime {
		var a module.AnounceModule
		a.Platform = "Xhongshu"
		a.Total = total
		a.LastTime = firstArticles.UpdateTime
		err := a.SaveAnouncement()
		utils.CheckErr(err)
		utils.DingCli.SendLinkMessage("小红书平台公告更新", "小红书平台公告有更新, 请点击查看", "",
			"https://open.xiaohongshu.com/platformSupport/notice/-1")
	}
}

// 淘宝平台公告 TODO

func TaobaoAnouncement() {
	//url := ""
}

// 1688平台公告 TODO

func AlibabaAnouncement() {

}

// 国际站平台公告 TODO

func AliAnouncement() {

}

// 微盟平台公告

func WeiMobAnouncement() {
	url := "https://www.weimobcloud.com/openapi/bulletin/findBulletin"
	requestParams := map[string]interface{}{
		"page":     1,
		"pageSize": 8,
		"param":    map[string]string{"categoryId": ""},
	}
	params, err := json.Marshal(requestParams)
	utils.CheckErr(err)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(params))
	utils.CheckErr(err)
	defer resp.Body.Close()

	var respJson module.WeimobJsonResponse
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &respJson)
	utils.CheckErr(err)
	total := respJson.Data.Total
	firstArticles := respJson.Data.ArticleData[0]
	announcement, _ := module.QueryAnouncement("Weimob")
	dbTotal, dbLastTime := announcement.Total, announcement.LastTime
	if total != dbTotal || dbLastTime != firstArticles.UpdateTime {
		var a module.AnounceModule
		a.Platform = "Weimob"
		a.Total = total
		a.LastTime = firstArticles.UpdateTime
		err := a.SaveAnouncement()
		utils.CheckErr(err)
		utils.DingCli.SendLinkMessage("微盟平台公告更新", "微盟平台公告有更新, 请点击查看", "",
			"https://www.weimobcloud.com/bulletin?pid=")
	}

}

// 有赞平台公告 TODO

func YouzanAnouncement() {

}

// 拼多多平台公告 TODO

func PinduoduoAnouncement() {

}

// 京东平台公告

func JingDongAnouncement() {
	url := "https://joshome.jd.com/doc/getNewJosChannelInfo?channelId=&pageIndex=1&pageSize=20"
	resp, err := http.Get(url)
	utils.CheckErr(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var respJson module.JingDongJsonResponse
	err = json.Unmarshal(body, &respJson)
	utils.CheckErr(err)
	total := respJson.Data.Total
	firstArticles := respJson.Data.ArticleData[0]
	announcements, _ := module.QueryAnouncement("JingDong")
	dbTotal, dbLastTime := announcements.Total, announcements.LastTime
	LastTime, err := time.Parse("2006-01-02T15:04:05.000-07:00", firstArticles.UpdateTime)
	LastUpdateTime := LastTime.Unix()
	if total != dbTotal || dbLastTime != LastUpdateTime {
		var a module.AnounceModule
		a.Platform = "JingDong"
		a.Total = total
		a.LastTime = LastUpdateTime
		err := a.SaveAnouncement()
		utils.CheckErr(err)
		utils.DingCli.SendLinkMessage("京东平台公告更新", "京东平台公告有更新, 请点击查看", "",
			"https://jos.jd.com/platformlist")
	}

}

// 快手平台公告

func KuaishouAnouncement() {
	url := "https://open.kwaixiaodian.com/rest/open/platform/doc/page/list?pageNum=1&pageSize=10&docType=1&docCatalogId=&location=0"
	resp, err := http.Get(url)
	utils.CheckErr(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	utils.CheckErr(err)
	var respJson module.KuaishouJsonResponse
	err = json.Unmarshal(body, &respJson)
	utils.CheckErr(err)
	total := respJson.Data.Total
	firstArticles := respJson.Data.ArticleData[0]
	announcement, _ := module.QueryAnouncement("Kuaishou")
	dbTotal, dbLastTime := announcement.Total, announcement.LastTime
	if total != dbTotal || dbLastTime != firstArticles.UpdateTime {
		var a module.AnounceModule
		a.Platform = "Kuaishou"
		a.Total = total
		a.LastTime = firstArticles.UpdateTime
		err := a.SaveAnouncement()
		utils.CheckErr(err)
		utils.DingCli.SendLinkMessage("快手平台公告更新", "快手平台公告有更新, 请点击查看", "",
			"https://open.kwaixiaodian.com/announcement/list?cateId=all")
	}
}

// 爱采购平台公告 TODO

func AiCaiGouAnouncement() {
}

// 千川平台公告 TODO

func QianchuanAnouncement() {

}
