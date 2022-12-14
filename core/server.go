package core

import (
	"fmt"
	"time"

	"github.com/haoleiqin/gin-flip-api/global"
	"github.com/haoleiqin/gin-flip-api/initialize"
	"github.com/haoleiqin/gin-flip-api/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	fmt.Printf("启动成功 当前版本:V2.5.0 启动成功!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
