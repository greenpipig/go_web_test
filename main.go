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
		fmt.Printf("访问成功，总访问数%d",*NUM)
	}
}

func main(){
	//定义一个方法，url为/httpflood
	NUM=new(int64)
	http.HandleFunc("/httpFlood",httpFlood)
	err:=http.ListenAndServe(":9090",nil)
	if err != nil{
		fmt.Printf("some thing wrong %v",err)
	}

}
