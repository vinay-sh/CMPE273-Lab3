package main

import(
"fmt"
"net/http"
"github.com/julienschmidt/httprouter"
"encoding/json"
)

type jsonobject struct{
	Key string
	Value string
}


var data [10]jsonobject
var i int

func PostKeyValue(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Println("Inside PostKeyValue")
	fmt.Println("Value of i",i)


 key := p.ByName("key_id")
 value := p.ByName("value")

 data[i].Key = key
 data[i].Value = value

 fmt.Println("Key_id is", data[i].Key)
 fmt.Println("Value is", data[i].Value)

  i++

fmt.Println("Value of i",i)
 
}


func GetKey(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

	key := p.ByName("key_id")

	fmt.Println("key_id is", key)

	fmt.Println("First value is", data[1].Key, data[1].Value)

	for j:=0;j<10;j++ {
		if data[j].Key == key{
			response:= jsonobject{data[j].Key, data[j].Value}
			resp, err := json.Marshal(response)
    		if err != nil {
    		panic(err)
    		}
    	rw.Write(resp)
    	fmt.Println("Marshaled response is", string(resp))
		}	
	}

} 

func GetAllKeys(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

	for j:=0;j<i;j++ {
			response:= jsonobject{data[j].Key, data[j].Value}
			resp, err := json.Marshal(response)
    		if err != nil {
    		panic(err)
    		}
    	rw.Write(resp)
    	fmt.Println("Marshaled response is", string(resp))
		}	
	}

func main(){
	 fmt.Println("Server 3001 running")
     mux := httprouter.New()
     mux.PUT("/keys/:key_id/:value",PostKeyValue) 
     mux.GET("/keys/:key_id",GetKey)
     mux.GET("/keys",GetAllKeys)
     server := http.Server{
             Addr:        "127.0.0.1:3001",
             Handler: mux,
     }
     server.ListenAndServe()
 }