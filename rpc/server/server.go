package server

import (
	"dev.model.360baige.com/models"
	"github.com/smallnest/rpcx/plugin"
	"github.com/rcrowley/go-metrics"
	"github.com/astaxie/beego/logs"
	"github.com/smallnest/rpcx"
	"github.com/astaxie/beego"
	"time"
)

func init() {
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)
	log.Debug("rpc run start")
	go run()
	log.Debug("rpc run end")
}

func run() {
	//服务
	services := map[string] interface{} {
		"Company.Cloud.Company": &models.Company{},
	}
	register(services)
}

/**
 * 注册服务
 */
func register(services map[string]interface{}) {
	etcdServerRegisterAddr := beego.AppConfig.String("RpcEtcdURL")
	serverRegisterAddr := beego.AppConfig.String("RpcServer")

	server := rpcx.NewServer()
	rplugin := &plugin.EtcdRegisterPlugin{
		ServiceAddress: "tcp@" + serverRegisterAddr,
		EtcdServers:    []string{etcdServerRegisterAddr },
		BasePath:       "/rpcx",
		Metrics:        metrics.NewRegistry(),
		Services:       make([]string, len(services)),
		UpdateInterval: time.Minute,
	}
	rplugin.Start()
	server.PluginContainer.Add(rplugin)
	server.PluginContainer.Add(plugin.NewMetricsPlugin())

	//注册 s
	for name, serv := range services {
		server.RegisterName(name, serv, "weight=1&m=devops")
	}
	//注册 e
	server.Serve("tcp", serverRegisterAddr)
}