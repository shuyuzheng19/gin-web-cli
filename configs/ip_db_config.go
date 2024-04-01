package configs

import (
	"gin-web/helper"

	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
)

var IpDB *xdb.Searcher

func InitIpDBConfig(dbPath string) {
	var err error

	IpDB, err = xdb.NewWithFileOnly(dbPath)

	if err != nil {
		helper.PanicErrorAndMessage(err, "加载IP数据库失败")
	}
}
