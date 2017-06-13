package server

import (
	"dev.model.360baige.com/models"
	"github.com/smallnest/rpcx/plugin"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx"
	"github.com/astaxie/beego"
	"time"
)

func init() {
	// 服务
	services := map[string]interface{}{
		"Account": &models.Account{},
		"AccountItem": &models.AccountItem{},
		"Transaction": &models.Transaction{},
		"App": &models.Application{},
		"AppTpl": &models.ApplicationTpl{},
		"AttendanceGroup":&models.ApplicationTpl{},
		"AttendanceRecord":&models.AttendanceRecord{},
		"AttendanceSetup":&models.AttendanceSetup{},
		"AttendanceShift":&models.AttendanceShift{},
		"AttendanceShiftItem":&models.AttendanceShiftItem{},
		"AttendanceShiftRecord":&models.AttendanceShiftRecord{},
		"Card":&models.Card{},
		"Company":&models.Company{},
		"Logger":&models.Logger{},
		"Machine":&models.Machine{},
		"Order":&models.Order{},
		"User":&models.User{},
		"UserPosition":&models.UserPosition{},
		"Person":&models.Person{},
		"PersonRelation":&models.PersonRelation{},
		"PersonStructure":&models.PersonStructure{},
		"Structure":&models.Structure{},
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