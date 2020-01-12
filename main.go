package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpFlood(w http.ResponseWriter,r *http.Request){
	//读取文件中的数据，为html文件
	if r.Method=="GET"{
		words,_:=ioutil.ReadFile("./http_flood.txt")
		_,_=fmt.Fprintf(w,string(words))
		fmt.Println(r.Method)
		fmt.Println("访问成功")
	}
}

func main(){
	//定义一个方法，url为/httpflood
	http.HandleFunc("/httpFlood",httpFlood)
	err:=http.ListenAndServe(":9090",nil)
	if err != nil{
		fmt.Printf("some thing wrong %v",err)
	}
}
