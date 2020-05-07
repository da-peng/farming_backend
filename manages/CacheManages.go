package manages

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
)

var (
	tokenCache cache.Cache
	err        error
)

func init() {
	//配置信息如下所示，配置的信息表示 GC 的时间，表示每隔 60s 会进行一次过期清理：
	//重新授权登录
	tokenCache, err = cache.NewCache("memory", `{"interval":60}`)
	if err != nil {
		beego.Info("缓存初始化异常 ——>", err)
	}
}

// StorageToken  存储token  {token:uid}
func StorageToken(token string, uid int, lastTime time.Time) {
	userInfo := make(map[int]time.Time)
	userInfo[uid] = lastTime
	tokenCache.Put(token, userInfo, 0)
}

// FindUIDByToken 查找Uid
func FindUIDByToken(token string) interface{} {
	value := tokenCache.Get(token)
	return value
}
