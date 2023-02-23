package main

import (
	cli "bgg01578/menu/go-client-generated"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/antihax/optional"
)

func main() {
	client := cli.NewAPIClient(cli.NewConfiguration())
	var choice int
	for {
		clearScreen()
		fmt.Println("=== Menu ===")
		fmt.Println("1. FindPetByStatus")
		fmt.Println("2. AddPet")
		fmt.Println("3. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("FindPetByStatusCalled")		
			var status,apikey string

			fmt.Printf(">> Status : ")
			fmt.Scanln(&status) 
			fmt.Printf(">> ApiKey : ")
			fmt.Scanln(&apikey)
			FindPetByStatus(client,status,apikey)
		case 2:
			fmt.Println("AddPet")
		case 3:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}

		fmt.Println("Press Enter to continue...")
		fmt.Scanln()
	}
}

func clearScreen() {
	cmd := exec.Command("clear") // for Unix-like systems
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // for Windows
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func FindPetByStatus(client *cli.APIClient, status string, apikey string ){
	pets, res, err  := client.PetApi.FindPetsByStatus(context.Background(),&cli.PetApiFindPetsByStatusOpts{Status: optional.NewString(status), ApiKey: optional.NewString(apikey)})

	// log.Print(pets,res,err)
	// data, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()


	if err==nil {
		for _,pet := range pets{
			log.Print("Id :",pet.Id)
			log.Print("Name :", pet.Name)		
			log.Print("\n\n")
		}
		log.Println(res.Status)
		log.Println(res.Proto)
		log.Print("\n\n")
	} else {
		log.Print(err.Error())
	}



	
}