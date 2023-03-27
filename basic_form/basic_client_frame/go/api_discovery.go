/*
 * Simple MEC Discovery API
 *
 * # Find your nearest MEC platform --- Network operators will typically have multiple MEC sites in a given territory. Connecting your application to a server on the closest MEC platform means the lowest latency - however, the physical location of a user is not an accurate match to the closest MEC site, due to the way operator networks are configured. This API returns the MEC platforms with the _shortest network path_ to the client making the request, and hence the lowest propagation delay. * If you have a server instance deployed there, connect to it to gain the lowest latency * Or if not, you may wish to deploy an instance there using the APIs of the cloud provider supporting that zone.    This API is intended to be called by a client application hosted on a UE attached to the operator network. _Note that the API parameters have been listed in this 'simple API' to align with the full API, but are optional and may not be supported by the API server_ ---
 *
 * API version: 0.8.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
 
 func Test1(w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json; charset=UTF-8")
 
	 fmt.Fprint(w,"Hello Client Test1")
	 log.Println("Hello Client Test1 log")
 
	 w.WriteHeader(http.StatusOK)
 }
 
 func Test2(w http.ResponseWriter, r *http.Request) {
	 w.Header().Set("Content-Type", "application/json; charset=UTF-8")
 
	 fmt.Fprint(w,"Hello Client Test2")
	 log.Println("Hello Client Test2 log")
 
	 w.WriteHeader(http.StatusOK)
 }
 
 
 
 //	http://localhost:8081/{value}


 func TestRequest(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	value := vars["value"]

	value ="(value:"+value+")-[client send]" 
	log.Println(value)

	// 요청 객체 생성
	req, _ := http.NewRequest("GET", "http://localhost:8080/testrequest/"+value, nil)


	// HTTP 클라이언트를 사용하여 서버(http://localhost:8080) 로 요청 보내기
	res, _ := http.DefaultClient.Do(req)

	// 응답 본문 읽기
	body, _ := ioutil.ReadAll(res.Body)

	// 응답 본문 출력 - response (from request to server)
	fmt.Fprint(w, string(body)+"-[client received]")
	log.Println(string(body)+"-[client received]")

	 // 응답 객체 닫기
	 res.Body.Close()
 }