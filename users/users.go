package users

import (
	"bufio"
	"encoding/csv"
	"github.com/qwertypomy/printers/dao/factory"
	"github.com/qwertypomy/printers/models"
	"github.com/qwertypomy/rand_string"
	"os"
	"path/filepath"
	"strings"
)

func Populate(config models.Config) (err error) {
	factoryDao := factory.FactoryDao{Engine: config.Engine}
	userDao := factoryDao.GetUserDaoInterface()

	path, _ := filepath.Abs("./users/users.csv")
	f, err := os.Open(path)
	if err != nil {
		return
	}
	r := csv.NewReader(bufio.NewReader(f))
	arr, err := r.ReadAll()
	if err != nil {
		return
	}

	admin := models.User{
		Name:     "Podvalnyi Mikhail",
		Email:    "podvalnyi.mikhail@gmail.com",
		Password: rand_string.String(16),
		IsAdmin:  true,
	}
	err = userDao.CreateUser(&admin)
	if err != nil {
		return
	}
	for _, v := range arr[1:] {
		if len([]rune(v[1])) > 45 {
			continue
		}
		user := models.User{
			Name:     v[1],
			Email:    strings.ToLower(strings.Trim(v[0],"_ -")) + "@gmail.com",
			Password: rand_string.String(16),
		}
		err = userDao.CreateUser(&user)
		if err != nil {
			return
		}
	}

	return
}
