package server

import (
	"dev.model.360baige.com/rpcx/plugin"
	"dev.model.360baige.com/models"
	"github.com/astaxie/beego/logs"
)

func init() {
	log := logs.NewLogger()
	log.SetLogger(logs.AdapterConsole)
	log.Debug("rpcx run start")
	go run()
	log.Debug("rpcx run end")
}

func run() {
	//服务
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
	plugin.Register(services)
}