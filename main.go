package main

import (
	"html/template"
	"log"
	"net/http"
)

// Import template.
var tpl = template.Must(template.ParseFiles("template/index.gohtml"))

func defaultHandler(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(res, "index.gohtml", req.Form)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Import CSS
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	// Run a server.
	log.Print("Server is running...")
	http.HandleFunc("/", defaultHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
