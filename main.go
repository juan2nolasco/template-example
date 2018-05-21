package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type PageVariables struct {
	Date string
	Time string
}

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	return ":8080"
}

func main() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(getPort(), nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	now := time.Now() // find the time right now
	HomePageVars := PageVariables{
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	t, err := template.ParseFiles("homepage.html") //parse the html file homepage.html
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
