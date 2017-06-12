package plugin

import (
	"github.com/smallnest/rpcx"
	"github.com/smallnest/rpcx/plugin"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/clientselector"
	"github.com/astaxie/beego"
	"time"
)

/**
 * 注册服务
 */
func Register(services map[string]interface{}) {
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

/**
 * 获取服务
 */
func Call(etcdURL, serviceName, methodName string, args, reply interface{}) error {
	// RandomSelect RoundRobin WeightedRoundRobin ConsistentHash
	s := clientselector.NewEtcdClientSelector([]string{etcdURL}, "/rpcx/" + serviceName, time.Minute, rpcx.RandomSelect, time.Minute)
	client := rpcx.NewClient(s)
	// Failfast Failover Failtry Broadcast Forking
	client.FailMode = rpcx.Failover
	err := client.Call(serviceName + "." + methodName, args, &reply)
	client.Close()
	return err
}
