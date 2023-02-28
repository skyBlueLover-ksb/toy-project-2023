package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func resp2User(r *http.Response) *User {
	myUser := new(User)
	json.NewDecoder(r.Body).Decode(myUser)
	return myUser
}

func readUsersList() string {
	resp, err := http.Get(requestURL + "/users")
	if err != nil || resp.StatusCode != http.StatusOK {
		return ""
	}
	data, _ := io.ReadAll(resp.Body)
	return string(data)
}

func readUser(id int) *User {
	resp, err := http.Get(requestURL + "/users/" + strconv.Itoa(id))
	if resp.StatusCode != http.StatusOK || err != nil {
		return nil
	}

	return resp2User(resp)
}

func createUser(first_name, last_name, email string) (*http.Response, error) {

	resp, err := http.Post(requestURL+"/users", "application/json",
		strings.NewReader(`{
			"first_name":"`+first_name+`", `+
			`"last_name":"`+last_name+`", `+
			`"email":"`+email+`" }`))
	return resp, err
}

func updateUser(id, first_name, last_name, email string) {
	req, _ := http.NewRequest("PUT", requestURL+"/users",
		strings.NewReader(`{
			"id":`+id+`", `+
			`"first_name":"`+first_name+`", `+
			`"last_name":"`+last_name+`", `+
			`"email":"`+email+`" }`))
	resp, err := http.DefaultClient.Do(req)
	if resp.StatusCode != http.StatusOK || err != nil {
		fmt.Println("오류가 발생하였습니다.")
		os.Exit(0)
	}
	data, _ := io.ReadAll(resp.Body)
	if strings.Contains(string(data), "No User ID:") {
		fmt.Println(id + "번 회원이 현재 없습니다.")
	} else {
		fmt.Println(id + "번 회원이 업데이트되었습니다.")
	}

}

func deleteUser(id string) {
	req, _ := http.NewRequest("DELETE", requestURL+"/users/"+id, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("오류가 발생하였습니다.")
		return
	}
	fmt.Println(id + "번 회원이 삭제되었습니다.")
}

const requestURL = "http://127.0.0.1:3000"

func main() {
	for true {

		userInput := ""
		fmt.Println("C/R/U/D 중 하나를 입력하세요. 아니면 종료됩니다.")
		fmt.Scanln(&userInput)

		switch userInput {
		case "C":
			first_name, last_name, email := "", "", ""
			fmt.Println("성, 이름, 이메일을 순서대로 한 줄씩 입력해주세요.")
			fmt.Scanln(&last_name)
			fmt.Scanln(&first_name)
			fmt.Scanln(&email)
			resp, err := createUser(first_name, last_name, email)
			if err != nil || resp.StatusCode != http.StatusCreated {
				fmt.Println("오류가 발생했습니다. 프로그램을 종료합니다.")
				os.Exit(0)
			}
			id := resp2User(resp).ID
			fmt.Printf("%d번 회원이 생성되었습니다.\n", id)

		case "R":
			option := ""
			fmt.Println("모든 회원 정보를 읽으려면 A를, 아니면 조회하고자 하는 유저의 id를 입력하세요.")
			fmt.Scanln(&option)
			if option == "A" {
				fmt.Println("전체 회원 조회.")
				fmt.Println(readUsersList())
			} else {
				id, _ := strconv.Atoi(option)
				fmt.Printf("%d번 회원 조회.", id)
				myUser := readUser(id)
				if myUser != nil {
					fmt.Println(myUser)
				} else {
					fmt.Println("오류가 발생하였습니다.")
				}
			}
		case "U":
			id, first_name, last_name, email := "", "", "", ""
			fmt.Println("업데이트하고자 하는 회원의 id를 입력하십시오")
			fmt.Scanln(&id)
			fmt.Println("성, 이름, 이메일을 순서대로 한 줄씩 입력해주세요.")
			fmt.Scanln(&last_name)
			fmt.Scanln(&first_name)
			fmt.Scanln(&email)
			updateUser(id, first_name, last_name, email)
		case "D":
			id := ""
			fmt.Println("지우고자 하는 회원의 id를 입력하십시오.")
			fmt.Scanln(&id)
			deleteUser(id)

		default:
			os.Exit(0)
		}
		fmt.Println("작업을 완료했습니다.")
	}
}
