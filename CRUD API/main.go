package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{

	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`


}

type Director struct{

	Firstname string `json:"fname"`
	Lastname string `json:"lname"`

}

var movies []Movie
func getmovies(w http.ResponseWriter , r *http.Request){

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)



}

func main(){

	r := mux.NewRouter() 
 
	movies = append(movies, Movie{ID: "1" , Isbn: "123454" , Title: "movie1" , Director: &Director{Firstname: "Mike" , Lastname: "Oxmall"}})
	movies = append(movies, Movie{ID: "2" , Isbn: "543223" , Title: "movie2" , Director: &Director{Firstname: "Tim" , Lastname: "Cook"}})



	r.HandleFunc("/movies" , getmovies).Methods("GET")
	r.HandleFunc("/movies/{id}" , getmovie).Methods("GET")
	r.HandleFunc("/movies" , createmovie).Methods("POST")
	r.HandleFunc("/movies/{id}" , updatemovie).Methods("PUT")
	r.HandleFunc("/movies/{id}" , deletemovie).Methods("DELETE")

	fmt.Printf("Starting Server at Port 8080")


	if err := http.ListenAndServe(":8080",nil); err != nil{
		log.Fatal(err)
	}

}