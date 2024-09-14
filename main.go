package main

import (
	"fmt"
	"log"
	"net/http"
)


func formHandler(w http.ResponseWriter, r *http.Request){
	var err = r.ParseForm()
	if err!=nil{
		fmt.Fprintf(w,"parse form error: %v",err)
		return
	}
	fmt.Fprintf(w,"POST request success full")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w,"Name = %s\n",name)
	fmt.Fprintf(w,"Address = %s\n",address)
}

func helloHandler(w http.ResponseWriter, r *http.Request){

	if r.URL.Path!="/hello" {
		http.Error(w,"404 path not found",http.StatusNotFound)
		return
	}
	if r.Method!="GET"{
		http.Error(w,"method is not supported",http.StatusMethodNotAllowed)
	}
	fmt.Fprintf(w,"Hello people")
}

func main() {
	
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

 fmt.Printf("starting server at port 8080\n")

 var err = http.ListenAndServe(":8080",nil); 
 if err!=nil {
	log.Fatal(err)
 }
}
