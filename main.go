package main

import (
	"fmt"
	"net/http"
)

func httpFlood(w http.ResponseWriter,r *http.Request){
	_,_=fmt.Fprintf(w,"<h1>hello world</h1>")
	fmt.Printf("r is %+v",*r)
	fmt.Println("访问成功")
}

func main(){
	http.HandleFunc("/httpflood",httpFlood)
	err:=http.ListenAndServe(":9090",nil)
	if err != nil{
		fmt.Printf("some thing wrong %v",err)
	}
}
