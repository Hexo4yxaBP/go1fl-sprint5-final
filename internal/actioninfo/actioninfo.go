package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, dataString := range dataset {
		if err := dp.Parse(dataString); err != nil {
			log.Println(err)
			continue
		}

		str, err := dp.ActionInfo()

		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(str)
	}
}
