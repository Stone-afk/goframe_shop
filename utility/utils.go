package utility

import (
	"math/rand"
	"time"

	"github.com/gogf/gf/v2/os/gtime"

	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/util/gconv"
)

// GetBefore7Date 获取一周前的日期
func GetBefore7Date() (date string) {
	gt := gtime.New(time.Now())
	date = gt.Add(-gtime.D * 6).Format("Y-m-d")
	return
}

// GetRecent7Date 生成最近一周的日期
func GetRecent7Date() (dates []string) {
	gt := gtime.New(time.Now())
	dates = []string{
		gt.Format("Y-m-d"),
		gt.Add(-gtime.D * 1).Format("Y-m-d"),
		gt.Add(-gtime.D * 2).Format("Y-m-d"),
		gt.Add(-gtime.D * 3).Format("Y-m-d"),
		gt.Add(-gtime.D * 4).Format("Y-m-d"),
		gt.Add(-gtime.D * 5).Format("Y-m-d"),
		gt.Add(-gtime.D * 6).Format("Y-m-d"),
	}
	return
}

// RandInt 获取随机整数
func RandInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func GetOrderNum() (number string) {
	rand.Seed(time.Now().UnixNano())
	number = gconv.String(time.Now().UnixNano()) + gconv.String(rand.Intn(1000))
	return
}

// EncryptPassword 密码加密
func EncryptPassword(password, salt string) string {
	return gmd5.MustEncryptString(gmd5.MustEncryptString(password) + gmd5.MustEncryptString(salt))
}
