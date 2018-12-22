package main

import (
	"gitee.com/cooge/impush/clog"
	"gitee.com/cooge/impush/cluster"
	"gitee.com/cooge/impush/config"
	"gitee.com/cooge/impush/rest"
	"gitee.com/cooge/impush/ws"
	"gitee.com/cooge/impush/xhttp"
	"gitee.com/cooge/impush/xtcp"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	/**
	初始化配置
	*/
	config.Init()

	/**
	  初始化log
	*/

	clog.Init(config.LOG_OPEN)

	/**
	  集群启动
	*/
	if config.CLUSTER_START {
		cluster.Start()
	}
	/**
	启动tcp协议
	*/
	if config.TCP_START {
		xtcp.Start()
	}

	/**
	启动websocket协议
	*/

	if config.WS_START {
		ws.Start()
	}

	/**
	http推送协议
	*/

	if config.EX_START {
		xhttp.Start()
	}

	/**
	接口启动
	*/
	rest.Start()
}
