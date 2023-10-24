package module

import (
	_ "github.com/mattn/go-sqlite3"
	"go_announce/utils"
)

type AnounceModule struct {
	Platform string
	Total    int
	LastTime int64
}

func (a *AnounceModule) SaveAnouncement() error {
	stmt, err := utils.DB.Prepare("UPDATE anouncement_info SET total = ?, lastTime = ? WHERE platform = ?")
	if utils.CheckErr(err) {
		return err
	}
	result, err := stmt.Exec(a.Total, a.LastTime, a.Platform)
	if utils.CheckErr(err) {
		return err
	}
	affected, err := result.RowsAffected()
	if affected == 0 {
		stmt, err = utils.DB.Prepare("INSERT INTO anouncement_info(platform, total, lastTime) VALUES(?, ?, ?)")
		_, err = stmt.Exec(a.Platform, a.Total, a.LastTime)
	}
	return err
}

func QueryAnouncement(platform string) (AnounceModule, error) {
	var resp AnounceModule
	stmt, err := utils.DB.Prepare("SELECT platform, total, lastTime FROM anouncement_info WHERE platform=?")
	if utils.CheckErr(err) {
		return resp, err
	}
	err = stmt.QueryRow(platform).Scan(&resp.Platform, &resp.Total, &resp.LastTime)
	if utils.CheckErr(err) {
		return resp, err
	}
	return resp, nil

}
