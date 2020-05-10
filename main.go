package main

import(
"encoding/json"
"log"
"net/http"
"math/rand"
"strconv"
"github.com/gorilla/mux"




)

//Book structure
type Book struct{
	ID  string `json:"id"`
	Isbn  string `json:"isbn"`
	Title  string `json:"title"`
	Author  *Author `json:"author"`
}
//Author structure
type Author struct{
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//init books
var books []Book

//func get all books
func getBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	params := mux.Vars(r)
	for _, item := range books{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}
func createBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}
func updateBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.ID == params["id"]{
			books=append(books[:index],books[index+1:]...)
			var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
			
		}
	}
	json.NewEncoder(w).Encode(books)
}
func deleteBooks(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type","application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.ID == params["id"]{
			books=append(books[:index],books[index+1:]...)
			break
			
		}
	}
	json.NewEncoder(w).Encode(books)
}
func main()  {
	//Router initiall
	r := mux.NewRouter()

	books = append(books, Book{ID:"1", Isbn:"21312", Title:"First Book", Author:&Author{Firstname:"Klevi", Lastname:"Merkuri"}})
	//Route hanglers
	r.HandleFunc("/api/books",getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	r.HandleFunc("/api/books",createBooks).Methods("POST")
	r.HandleFunc("/api/books/{id}",updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}",deleteBooks).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000",r))
}
