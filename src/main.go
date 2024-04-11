// coucou curtis, bonne année va en bas stp
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	//tout nos varible qui  sont accessibles partout dans notre programme.
	tmpl           *template.Template
	httpServer     *http.Server
	AnswerAll      [][]interface{}
	Ask            string
	Answer         []string
	Page           string
	Mode           string
	SelectedOption string
	Alllastopt     []bool
	Statue         bool
	Ms             = ""
	pseudo         = ""
	password       = ""
	b              = false
	q              int
	EssaySerie     int
	ResulteSerie   int
	Serie          []string
	SerieNum       int
)

// Fonction principale
func main() {
	//on charge notre page html
	var err error
	tmpl, err = loadTemplate("asset/html/index.html")
	if err != nil {
		// Handle template loading error
		fmt.Println("Error loading template:", err)
		return
	}
	//la on gere se qui il a dans l'url pour dire genre si t'es dans cet url fais se code...
	http.HandleFunc("/", inde)
	http.HandleFunc("/Compte", CompteHandler)
	http.HandleFunc("/Theme", ThemePage)
	http.HandleFunc("/NewQuest", NewQuestPage)
	http.HandleFunc("/Win", WinPage)
	http.HandleFunc("/Score", Score)
	http.HandleFunc("/Quest", QuestPage)
	http.HandleFunc("/HandlerVerif", VerifPage)

	//start du serveur
	fs := http.FileServer(http.Dir("asset"))
	http.Handle("/asset/", http.StripPrefix("/asset/", fs))
	fmt.Println("Server started on :5656")
	http.ListenAndServe(":5656", nil)
}
