package menu

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

type menuItem struct {
	name  string
	price map[string]float32
}

type menu []menuItem

func (m menu) print() {
	for _, value1 := range m {
		fmt.Println(strings.Repeat("-", 30))
		fmt.Println(value1.name)
		for key2, value2 := range value1.price {
			fmt.Printf("\t%10s%10.2f\n", key2, value2)
		}
	}
	fmt.Println(strings.Repeat("-", 30))
}

func Print() {
	data.print()
}

func (m *menu) addItem() error {
	fmt.Println("Enter the name of the new item:")
	name, _ := in.ReadString('\n')
	name = strings.TrimSpace(name)

	for _, item := range *m {
		if item.name == name {
			return errors.New("item already resent")
		}
	}

	*m = append(*m, menuItem{name: name, price: make(map[string]float32)})

	return nil
}

func AddItem() error {
	return data.addItem()
}
