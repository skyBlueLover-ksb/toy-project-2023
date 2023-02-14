// You can edit this code!
// Click here and start typing.
package main

type service struct {
	client *APIClient
}
type PetApiService service
type UserApiService service

type APIClient struct {
	common service
	PetApi *PetApiService
	UserApi *UserApiService
}

func (a PetApiService) AddPet(num int, num2 int) {
	println(num+num2)
}

func (a UserApiService) SubPet(num int, num2 int) {
	println(num-num2)
}

func NewAPIClient() *APIClient {
	c := &APIClient{}
	//c.common.client = c

	// API Services
	c.PetApi = (*PetApiService)(&c.common)
	c.UserApi = (*UserApiService)(&c.common)

	return c
}


func main() {
	c := NewAPIClient()
	print(c.PetApi.client)
	println(c.PetApi.client)
	c.PetApi.AddPet(1,2)
	c.UserApi.SubPet(1,2)
}