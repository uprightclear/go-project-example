package website

import (
	"go-project-example/internal/app/type/storage"
	"go-project-example/internal/app/type/website"
	"go-project-example/internal/pkg/db"
)

func Test(req website.TestReq) (resp website.TestResp, err error) {
	var relayoplog storage.RelayOpLog
	if err := db.GetMySQL().Table(storage.RelayOpLog{}.TableName()).
		Where("id = ?", req.ID).First(&relayoplog).Error; err != nil {

	}
	return
}
