package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var NUM *int64

func httpFlood(w http.ResponseWriter,r *http.Request){
	//读取文件中的数据，为html文件
	*NUM=*NUM+1
	if r.Method=="GET"{
		//可以自定义返回码
		w.WriteHeader(200)
		//c1 := http.Cookie{
		//	Name:"first_cookie",
		//	Value:"Go Web Programming",
		//	HttpOnly: true,
		//}
		//http.SetCookie(w, &c1)
		words,_:=ioutil.ReadFile("./http_flood.txt")
		_,_=fmt.Fprintf(w,string(words))
		//fmt.Println(r.Method)
		fmt.Printf("访问成功，总访问数%d",*NUM)
	}else if r.Method=="POST"{
		//接收application/x-www-form-urlencoded编码的数据
		fmt.Println(r.FormValue("id"))
		//接受endcodeing/json编码的数据
		con, _ := ioutil.ReadAll(r.Body)
		fmt.Println(string(con))
		fmt.Printf("访问成功，总访问数:%d",*NUM)
		//todo 写入数据库，模拟dbflood
		//todo 从txt中读取db数据
	}
}

func test(w http.ResponseWriter,r *http.Request){
	fmt.Printf("in func\n")
}


func main(){
	NUM=new(int64)
	//定义一个方法，url为/httpflood
	http.HandleFunc("/httpFlood",httpFlood)
	http.HandleFunc("/",test)
	err:=http.ListenAndServe(":9090",nil)
	//自定义地址和端口
	//err:=http.ListenAndServe("10.28.186.172:9090",nil)
	if err != nil{
		fmt.Printf("some thing wrong %v",err)
	}
}
