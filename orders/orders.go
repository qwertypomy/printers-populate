package orders

import (
	"github.com/qwertypomy/printers/models"
	"github.com/qwertypomy/printers/dao/factory"
	"math/rand"
)
func Populate(config models.Config) (err error) {
	factoryDao := factory.FactoryDao{Engine: config.Engine}
	orderDao := factoryDao.GetOrderDaoInterface()
	userDao := factoryDao.GetUserDaoInterface()
	printerDao := factoryDao.GetPrinterDaoInterface()

	users, err := userDao.UserList()
	if err != nil {
		return err
	}
	printers, err := printerDao.PrinterList()
	if err != nil {
		return err
	}
	usersN := len(users)
	printersN := len(printers)

	for i:=0; i<1000; i++ {
		user := users[rand.Int()%(usersN - 1) + 1]
		order := models.Order{UserID:user.ID, Status:uint(rand.Int()) % 6}
		err := orderDao.CreateOrder(&order)
		if err != nil {
			return err
		}
		onlyOnePrinter := (rand.Int() % 10) != 0
		var amount uint
		if onlyOnePrinter {
			amount = 1
		} else {
			amount = uint(rand.Int()) % 29 + 1
		}
		for amount > 0 {
			am := uint(rand.Int()) % (amount) + 1
			amount -= am;
			printer := printers[rand.Int()%printersN]
			orderHasPrinter := models.OrderHasPrinter{PrinterID:printer.ID, OrderID:order.ID, Amount: am}
			orderDao.CreateOrderHasPrinter(&orderHasPrinter)
			if err != nil {
				return err
			}
		}
	}
	return
}
