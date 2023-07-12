package main

import (
	"bufio"
	"demo/menu"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func main() {

out:
	for {
		fmt.Println("Please select an option:")
		fmt.Println("0) Exit")
		fmt.Println("1) Print Menu")
		fmt.Println("2) Add Item")

		choice, _ := in.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "0":
			fmt.Println("Goodbye!")
			break out
		case "1":
			menu.Print()
		case "2":
			err := menu.AddItem()
			if err != nil {
				fmt.Println(fmt.Errorf("invalid input: %w", err))
			}
		default:
			fmt.Println(strings.ToUpper("Option not recognized, try again") + "!")
		}
	}

}
