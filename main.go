package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {

	route := mux.NewRouter()

	// route path folder public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	// routing
	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/form-project", formAddProject).Methods("GET")
	route.HandleFunc("/detail-project/{projectName}", detailProject).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/add-project", addProject).Methods("POST")

	fmt.Println("server running on port 5000")
	// menjalankan server
	http.ListenAndServe("localhost:5000", route)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("message :" + err.Error()))
		return
	}

	// w.Write([]byte("add project"))
	// w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func formAddProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/addproject.html")

	if err != nil {
		w.Write([]byte("message :" + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}

func addProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Title: " + r.PostForm.Get("projectName"))
	fmt.Println("Content: " + r.PostForm.Get("description"))
	fmt.Println("StartDate : " + r.PostForm.Get("startDate"))
	fmt.Println("EndDate : " + r.PostForm.Get("endDate"))
	fmt.Println("NodeJS : " + r.PostForm.Get("nodeJS"))
	fmt.Println("NextJS : " + r.PostForm.Get("nextJS"))
	fmt.Println("ReactJS : " + r.PostForm.Get("reactJS"))
	fmt.Println("VueJS : " + r.PostForm.Get("vueJS"))

	http.Redirect(w, r, "/", http.StatusMovedPermanently)

}

func detailProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/detailproject.html")

	if err != nil {
		w.Write([]byte("message :" + err.Error()))
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["projectName"])

	// menampung data object
	data := map[string]interface{}{
		"Title":   "Hallo Title",
		"Content": "Lorem ipsum dolor, sit amet consectetur adipisicing elit. Quibusdam, earum quaerat. Sint istinctio nostrum corrupti animi laboriosam eum rerum doloribus nobis odit, veniam modi eligendi ipsam.I llo dolorem excepturi natus! Lorem ipsum dolor sit amet consectetur adipisicing elit. Consectetur, unt. Consequatur ad doloribus autem voluptatum quod ducimus nisi quis, sit, nemo laborum temporibus",
		"Id":      id,
	}

	tmpl.Execute(w, data)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var tmpl, err = template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("message :" + err.Error()))
		return
	}

	tmpl.Execute(w, nil)
}
