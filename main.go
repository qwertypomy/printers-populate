package main

import (
	"log"
	"github.com/qwertypomy/printers/config"
	"github.com/qwertypomy/printers/dao/factory"
	"github.com/qwertypomy/printers_populate/printers"
	"github.com/qwertypomy/printers_populate/users"
	"github.com/qwertypomy/printers_populate/orders"
)

func DeleteAll() {
	Config, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	factoryDao := factory.FactoryDao{Engine: Config.Engine}
	userDao := factoryDao.GetUserDaoInterface()
	printerDao := factoryDao.GetPrinterDaoInterface()
	orderDao := factoryDao.GetOrderDaoInterface()

	err = userDao.DeleteAllUsers()
	if err != nil {
		log.Fatal(err)
	}
	err = printerDao.DeleteAllData()
	if err != nil {
		log.Fatal(err)
	}
	err = orderDao.DeleteAllData()
	if err != nil {
		log.Fatal(err)
	}
}

func PopulateAll() {
	Config, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	err = printers.Populate(Config)
	if err != nil {
		log.Fatal(err)
	}
	err = users.Populate(Config)
	if err != nil {
		log.Fatal(err)
	}
	err = orders.Populate(Config)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	DeleteAll()
	PopulateAll()
}

