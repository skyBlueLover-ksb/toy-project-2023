package main

import (
	cli "bgg01578/menu/go-client-generated"
	menu "bgg01578/menu/menu"
	"fmt"
)

func main() {
	client := cli.NewAPIClient(cli.NewConfiguration())
	for {
		var choice int
		menu.ClearScreen()
		fmt.Println("=== Menu ===")
		fmt.Println("1. Pet")
		fmt.Println("2. Store")
		fmt.Println("3. User")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		
	switch choice {
		case 1:
			menu.ClearScreen()
			menu.PetMenu(client)
		case 2:
			menu.ClearScreen()
			fmt.Println("Store Menu")
		case 3:
			menu.ClearScreen()
			fmt.Println("User Menu")
		default:
			menu.ClearScreen()
			fmt.Println("Invalid choice")
	}

		fmt.Println("Press Enter to continue...")
		fmt.Scanln()
	}
}