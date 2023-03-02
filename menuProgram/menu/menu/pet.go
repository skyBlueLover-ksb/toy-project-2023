package menu

import (
	cli "bgg01578/menu/go-client-generated"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/antihax/optional"
)

const (
	FINDPETSBYSTATUS = 1
	GETPETBYID=2
	ADDPET=3
	DELETEPET=4
	GETPETLIST=5
	CLEARPET=6
	EXIT=7
)

func PetMenu(client *cli.APIClient){
	var choice int	//for choice
	PrintPetMenu()	//print pet menu
	fmt.Scanln(&choice)

	ClearScreen()
	switch choice {
		case FINDPETSBYSTATUS: //FindPetByStatus
			var status,apikey string
			fmt.Println("[FindPetByStatus]")		
			fmt.Printf(">> Status : ")
			fmt.Scanln(&status) 
			fmt.Printf(">> ApiKey : ")
			fmt.Scanln(&apikey)
			FindPetByStatus(client,status,apikey)

		case GETPETBYID:
			var petId int64
			fmt.Println("[GetPetById]")		
			fmt.Printf(">> PetId : ")
			fmt.Scanln(&petId)
			GetPetById(client,petId)

		case ADDPET:
			var petId,categoryId int64
			var petName,categoryName,Status string
			fmt.Println("[AddPet]")		
			fmt.Printf(">> PetId : ") 
			fmt.Scanln(&petId)
			fmt.Printf(">> PetName : ")
			fmt.Scanln(&petName)
			fmt.Printf(">> Category Id: ")
			fmt.Scanln(&categoryId)
			fmt.Printf(">> Category Name: ")
			fmt.Scanln(&categoryName)
			fmt.Printf(">> Status: ")
			fmt.Scanln(&Status)
			AddPet(client,cli.Pet{Id:petId,Name:petName,Category: &cli.Category{Id:categoryId,Name:categoryName},Status: Status})

		case DELETEPET:
			var petId int64
			var apikey string
			fmt.Println("[DeletePet]")		
			fmt.Printf(">> petId : ")
			fmt.Scanln(&petId) 
			fmt.Printf(">> ApiKey : ")
			fmt.Scanln(&apikey)
			DeletePet(client,petId,apikey)

		case GETPETLIST:
			ClearScreen()
			fmt.Println("[GetList]")
			GetPetList(client)

		case CLEARPET:
			fmt.Println("[ClearPet]")
			ClearPet(client)

		case EXIT:
			os.Exit(0)

		default:
			fmt.Println("Invalid choice")
	}
}

func PrintPetMenu(){
	fmt.Println("[Pet Menu]")
	fmt.Println("1. FindPetByStatus")
	fmt.Println("2. GetPetById")
	fmt.Println("3. AddPet")
	fmt.Println("4. DeletePet")
	fmt.Println("5. GetList")
	fmt.Println("6. ClearPet")
	fmt.Println("7. Exit")
	fmt.Print("Enter your choice: ")
}

func FindPetByStatus(client *cli.APIClient, status string, apikey string ){
	pets, res, err  := client.PetApi.FindPetsByStatus(context.Background(),
	&cli.PetApiFindPetsByStatusOpts{Status: optional.NewString(status), 
		ApiKey: optional.NewString(apikey)})


	idx := 1
	if err==nil {
		for _,pet := range pets{
			log.Print("<Pet(",idx, ")>")
			idx++
			log.Print("Id :",pet.Id)
			log.Print("Name :", pet.Name)		
			log.Print("Status :", pet.Status,"\n\n")
		}
		log.Println(res.Status)
		log.Println(res.Proto)
	} else {
		log.Print(err.Error())
	}	
}

func GetPetById(client *cli.APIClient, petId int64 ){
	pet, res, err  := client.PetApi.GetPetById(context.Background(),petId)
	if err==nil {	
		log.Print("<Response>\n")
		log.Print("response Status :" ,res.Status)
		log.Print("response Proto :",res.Proto,"\n\n")

		log.Print("<Pet>\n")
		log.Print("Id :",pet.Id)
		log.Print("Name :", pet.Name)	
		log.Print("Status :", pet.Status)	
	} else {
		log.Print(err.Error())
	}	
}

func AddPet(client *cli.APIClient, body cli.Pet){
	pet, res, err  := client.PetApi.AddPet(context.Background(),body)

	if(err==nil){ //추가 성공
		log.Print("<Pet Added>\n")
		log.Println("Response Status:",res.Status)
		log.Println("CategoryId : " ,pet.Category.Id )
		log.Println("CategoryName : " ,pet.Category.Name )	
		log.Println("Id : " ,pet.Id )
		log.Println("Name : " ,pet.Name )
		log.Println("Status : ", pet.Status)
	}else{ //추가 실패
		log.Println(err.Error())
	}
}

func DeletePet(client *cli.APIClient, petId int64, apikey string){
	res, err  := client.PetApi.DeletePet(context.Background(),petId,&cli.PetApiDeletePetOpts{ApiKey:optional.NewString(apikey)})

	if err==nil { //삭제 성공
		log.Println("Response Status:",res.Status)
	}else{ //삭제 실패
		log.Println(err.Error())
	}
}

func GetPetList(client *cli.APIClient){
	pets, res, err  := client.PetApi.GetPetList(context.Background())	
	if err!=nil {
		log.Print(err.Error())
		return
	}

	idx := 0
	if err==nil {
		for _,pet := range pets{
			idx++
			log.Print("<Pet(",idx, ")>")
			log.Print("Id :",pet.Id)
			log.Print("Name :", pet.Name)		
			log.Print("Status :", pet.Status,"\n\n")
		}
		log.Println(res.Status)
		log.Println(res.Proto)
	}else{
		log.Print(err.Error())
	}
}

func ClearPet(client *cli.APIClient){

	res,err  := client.PetApi.ClearPet(context.Background())	
	if err==nil {
		log.Print("<Success>")
		log.Print(res.Status)
		log.Print(res.Proto,"\n\n")
	} else {
		log.Print("<Fail>")
		log.Print(err.Error())
	}	
}