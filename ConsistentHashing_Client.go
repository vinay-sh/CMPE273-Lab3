package main

import(
"fmt"
"net/http"
"github.com/julienschmidt/httprouter"
"hash/fnv"
"bytes"
)

type jsonobject struct{
	Key string
	Value string
}


var data [10]jsonobject
var i int
var k [10]string
var v [10]string
var server [3]string
var CHashServer map[string]uint32
var CHashClient map[string]uint32
var cServerAssigned map[string]string
var s [3] uint32


func result(u uint32) string {
	var ss [3] uint32
	ss = s
	sort(&ss)
	if u < ss[0] || u > ss[2]{
		return mapserver(ss[0])
	}
	if u>ss[0] && u<ss[1]{
		return mapserver(ss[1])
	}
	return mapserver(ss[2])
}


func sort(serv *[3]uint32){
	if serv[0] > serv[1] {
		swap(&serv[0], &serv[1])		
	}
	if serv[0] > serv[2] {
		swap(&serv[0], &serv[2])	
	}
	if serv[1] > serv[2] {
		swap(&serv[1], &serv[2])
	}
}

func swap(a *uint32, b *uint32){
	var temp uint32
	temp=*a
	*a=*b
	*b=temp
}

func mapserver(u uint32) string{
	if u == CHashServer[server[0]]{
		return server[0]		
	}	else if u == CHashServer[server[1]]{
		return server[1]
	}	else if u == CHashServer[server[2]]{
		return server[2]
	}
	return "nil"
			
}

func SendRequest(a string, b string, c string){
	url:=a+"/"+b+"/"+c
	fmt.Println("Value of Request", url)
	jsonStr := []byte(`{}`)
	r, err := http.NewRequest("PUT",url, bytes.NewBuffer(jsonStr)) 
	client := &http.Client{}
	response, err:=client.Do(r)	
	if err != nil {
	panic(err)
	}
	fmt.Println(response.StatusCode)
	defer response.Body.Close()
}


func GetRequest(a string){
	xyz := "\""+a+"\""
	fmt.Println("Value of string to int conversion of key", xyz)
	url:=cServerAssigned["a"]+"/"+a
	fmt.Println("Value of Request", url)
	jsonStr := []byte(`{}`)
	r, err := http.NewRequest("GET",url, bytes.NewBuffer(jsonStr)) 
	client := &http.Client{}
	response, err:=client.Do(r)	
	if err != nil {
	panic(err)
	}
	fmt.Println(response.StatusCode)
	defer response.Body.Close()
}

func hash(s string) uint32 {
        h := fnv.New32a()
        h.Write([]byte(s))
        return h.Sum32()
}


func PostKeyValue(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {

//var c1, c2, c3, s1, s2, s3, c4, c5, c6, c7, c8, c9, c10 uint32
var str string


server[0] = "http://localhost:3000/keys"
server[1] = "http://localhost:3001/keys"
server[2] = "http://localhost:3002/keys"

v[0] = "a"
v[1] = "b"
v[2] = "c"
v[3] = "d"
v[4] = "e"
v[5] = "f"
v[6] = "g"
v[7] = "h"
v[8] = "i"
v[9] = "j"

k[0] = "1"
k[1] = "2"
k[2] = "3"
k[3] = "4"
k[4] = "5"
k[5] = "6"
k[6] = "7"
k[7] = "8"
k[8] = "9"
k[9] = "10"

CHashServer[server[0]] = hash("http://localhost:3000/keys")
fmt.Println("Hashed Value of Server A is",CHashServer[server[0]])
s[0]=CHashServer[server[0]]

CHashServer[server[1]] = hash("http://localhost:3001/keys")
fmt.Println("Hashed Value of Server B is", CHashServer[server[1]])
s[1]=CHashServer[server[1]]

CHashServer[server[2]] = hash("http://localhost:3002/keys")
fmt.Println("Hashed Value of Server C is", CHashServer[server[2]])	 
s[2]=CHashServer[server[2]]

CHashClient[v[0]] = hash("1=>a")
fmt.Println("Hashed Value of K-V 1 is", CHashClient[v[0]])
str = result(CHashClient[v[0]])
cServerAssigned[v[0]] = str
SendRequest(str, k[0], v[0]) 

CHashClient[v[1]] = hash("2=>b")
fmt.Println("Hashed Value of K-V 2 is", CHashClient[v[1]])	
str = result(CHashClient[v[1]])
cServerAssigned[v[1]] = str
SendRequest(str, k[1], v[1])  

CHashClient[v[2]] = hash("3=>c")
fmt.Println("Hashed Value of K-V 3 is", CHashClient[v[2]])	 
str = result(CHashClient[v[2]])
cServerAssigned[v[2]] = str
SendRequest(str, k[2], v[2]) 

CHashClient[v[3]] = hash("4=>d")
fmt.Println("Hashed Value of K-V 4 is", CHashClient[v[3]])	
str = result(CHashClient[v[3]])
cServerAssigned[v[3]] = str
SendRequest(str, k[3], v[3])  

CHashClient[v[4]] = hash("5=>e")
fmt.Println("Hashed Value of K-V 5 is", CHashClient[v[4]])	 
str = result(CHashClient[v[4]])
cServerAssigned[v[4]] = str
SendRequest(str, k[4], v[4]) 

CHashClient[v[5]] = hash("6=>f")
fmt.Println("Hashed Value of K-V 6 is", CHashClient[v[5]])	
str = result(CHashClient[v[5]])
cServerAssigned[v[5]] = str
SendRequest(str, k[5], v[5]) 

CHashClient[v[6]] = hash("7=>g")
fmt.Println("Hashed Value of K-V 7 is", CHashClient[v[6]])	
str = result(CHashClient[v[6]])
cServerAssigned[v[6]] = str
SendRequest(str, k[6], v[6]) 

CHashClient[v[7]] = hash("8=>h")
fmt.Println("Hashed Value of K-V 8 is", CHashClient[v[7]])	 
str = result(CHashClient[v[7]])
cServerAssigned[v[7]] = str
SendRequest(str, k[7], v[7]) 

CHashClient[v[8]] = hash("9=>i")
fmt.Println("Hashed Value of K-V 9 is", CHashClient[v[8]])	 
str = result(CHashClient[v[8]])
cServerAssigned[v[8]] = str
SendRequest(str, k[8], v[8]) 

CHashClient[v[9]] = hash("10=>j")
fmt.Println("Hashed Value of K-V 10 is", CHashClient[v[9]])	 
str = result(CHashClient[v[9]])
cServerAssigned[v[9]] = str
SendRequest(str, k[9], v[9]) 

}


func GetKey(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var key string
	key = p.ByName("key_id")
	GetRequest(key)

}

func main() {
	 fmt.Println("Inside Main")
	 CHashServer = make(map[string]uint32)
	 CHashClient = make(map[string]uint32)
	 cServerAssigned = make(map[string]string)
     mux := httprouter.New()
     mux.PUT("/keys",PostKeyValue)
     mux.GET("/keys/:key_id",GetKey)
     server := http.Server{
             Addr:        "127.0.0.1:8080",
             Handler: mux,
     }
     server.ListenAndServe()
}