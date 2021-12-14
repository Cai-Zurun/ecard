package main

import (
	_ "ecard_gateway/db"
	"ecard_gateway/ecard/server"
	"ecard_gateway/ecard/ws_server"
	"ecard_gateway/log"
	"ecard_gateway/utils"
	"net/http"
	_ "net/http/pprof"
)

func init() {
	log.InitLogger("./log/ecard.log", "Debug", 100, 1)
}

func main() {
	//程序崩溃时恢复，并记录崩溃原因
	utils.RecoverAndLog()

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Error("pprof failed: %v", err)
		}
	}()

	go ws_server.StartWebsocketServer()

	//TcpServer: 接收设备数据存入hbase
	server.GetTcpServer().Start()
}
