package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler( w http.ResponseWriter , r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w , "Error %v" , err)
		return
	}

	fmt.Fprintf(w , "Success")
	fname := r.FormValue("fname")
	sname := r.FormValue("sname")
	fmt.Fprintf(w, "Name : %v " , fname)
	
	fmt.Fprintf(w , "Last Name %v: " , sname)

}

func helloHandler(w http.ResponseWriter , r *http.Request){

	if r.URL.Path != "/Hello" {
		http.Error( w , "404 not found " , http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w , " Method is not supported" , http.StatusNotFound)
		return
	}

	fmt.Fprintf(w , "Hello")


}

func main(){


	fileServer := http.FileServer(http.Dir("./Static"))

	http.Handle("/" , fileServer)
	http.HandleFunc("/Hello" , helloHandler)
	http.HandleFunc("/form" , formHandler)

	fmt.Println("Starting Server at port 8080:")
	if err := http.ListenAndServe(":8080",nil); err != nil {

		log.Fatal(err)
	}

}