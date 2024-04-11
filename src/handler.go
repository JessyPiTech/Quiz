package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// gestion page index
func inde(w http.ResponseWriter, r *http.Request) {
	Page = "Acceuill"
	Data := createData(Statue, Ask, Answer, Page, Alllastopt, Ms, ResulteSerie, Serie)
	err := tmpl.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// guestion de la serie choisi genre noel été paque ....et des mode hard medium....
func ThemePage(w http.ResponseWriter, r *http.Request) {
	PlayZero()
	if r.Method == http.MethodPost {
		r.ParseForm()
		selectedModeAndTheme := r.FormValue("Mode")

		//d'abord on recupere la valeur renvoyer par la page
		//puis on recupere le fichier lié
		Mode = "asset/ask/" + selectedModeAndTheme + ".txt"
		//on decoupe la valeur pour definir  le theme et le mode dans le jeu diretement
		Intermediaire := strings.SplitN(selectedModeAndTheme, "-", 2)
		switch Intermediaire[0] {
		case "Easy":
			switch Intermediaire[1] {
			case "Geographie":
				SerieNum = 0
			case "Histoire":
				SerieNum = 3
			case "Noel":
				SerieNum = 6
			case "Paque":
				SerieNum = 9
			case "Eté":
				SerieNum = 12
			}
		case "Medium":
			switch Intermediaire[1] {
			case "Geographie":
				SerieNum = 1
			case "Histoire":
				SerieNum = 4
			case "Noel":
				SerieNum = 7
			case "Paque":
				SerieNum = 10
			case "Eté":
				SerieNum = 13
			}
		case "Hard":
			switch Intermediaire[1] {
			case "Geographie":
				SerieNum = 2
			case "Histoire":
				SerieNum = 5
			case "Noel":
				SerieNum = 8
			case "Paque":
				SerieNum = 11
			case "Eté":
				SerieNum = 14
			}
		}
		http.Redirect(w, r, "/NewQuest", http.StatusSeeOther)
	}
	Page = "Mode"
	Ask = ""
	Answer = make([]string, 4)
	Data := createData(Statue, Ask, Answer, Page, Alllastopt, Ms, ResulteSerie, Serie)
	err := tmpl.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewQuestPage(w http.ResponseWriter, r *http.Request) {
	//on verifi le nombre de question deja faite
	//si il a fini la serie on appele notre page score
	if q >= 8 {
		http.Redirect(w, r, "/Score", http.StatusSeeOther)
	}
	//sinon on prend une nouvelle question et les nouvelle reponce
	Page = "Jeu"
	Ask, AnswerAll = ChooseAsks(Mode, q)
	Answer = make([]string, 4)
	Alllastopt = make([]bool, 4)
	for i := 0; i < 4; i++ {
		Answer[i] = AnswerAll[i][1].(string)
		fmt.Println(Answer[i])
	}
	//puis on redirige vers notre page
	http.Redirect(w, r, "/Quest", http.StatusSeeOther)
}

// page des questions
func QuestPage(w http.ResponseWriter, r *http.Request) {
	Page = "Jeu"
	Data := createData(Statue, Ask, Answer, Page, Alllastopt, Ms, ResulteSerie, Serie)
	err := tmpl.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// verification de la reponse pour que on la met en rouge si c'est faux
func VerifPage(w http.ResponseWriter, r *http.Request) {
	Page = "Jeu"
	var selectedOption string
	if r.Method == http.MethodPost {
		r.ParseForm()
		selectedOption = r.FormValue("option")
		EssaySerie++
	}
	for i := 0; i < 4; i++ {
		if selectedOption == AnswerAll[i][1].(string) {
			if AnswerAll[i][0].(bool) {
				fmt.Println("You have won!")
				Page = "Win"
				http.Redirect(w, r, "/Win", http.StatusFound)
				q++
			} else {
				Alllastopt[i] = true
				fmt.Println("You have lost.")
				Page = "Deafeat"
				http.Redirect(w, r, "/Quest", http.StatusFound)
			}
			break
		}
	}
	Data := createData(Statue, Ask, Answer, Page, Alllastopt, Ms, ResulteSerie, Serie)
	err := tmpl.Execute(w, Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// page de bonne reponce
func WinPage(w http.ResponseWriter, r *http.Request) {
	Page = "Win"
	Data := createData(Statue, Ask, Answer, Page, Alllastopt, Ms, ResulteSerie, Serie)
	err := tmpl.Execute(w, Data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

// gestion du resultat de la serie
func Score(w http.ResponseWriter, r *http.Request) {
	var resulte float64
	resulte = 10 / float64(EssaySerie)
	ResulteSerie := int(resulte * 100.0)
	//on le met sur 100% pour l'affichage
	Serie[SerieNum] = strconv.Itoa(ResulteSerie)
	fmt.Println("le resultat de la serie " + Serie[SerieNum] + "%")
	Page = "Resulte"
	//puis on met le profil joueur a jour
	mise()
	Data := createData(Statue, Ask, Answer, Page, Alllastopt, Ms, ResulteSerie, Serie)
	err := tmpl.Execute(w, Data)
	//puis on remet tout a 0
	PlayZero()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CompteHandler(w http.ResponseWriter, r *http.Request) {
	Page = "Connection"
	//deja on remet tout a 0 dès la connection pour pouvoir changer de compte au besoin
	Ms = ""
	pseudo = ""
	password = ""
	Statue = false
	//si la page envoie un post
	if r.Method == http.MethodPost {
		pseudo = strings.ToLower(r.FormValue("pseudo"))
		password = strings.ToLower(r.FormValue("password"))
		//on recup le pseudo et mots de pass
		contenu := password + "\n"
		filename := "players/" + pseudo + ".txt"
		//si le fichier si dessu n'existe pas on le cree avec comme nom le pseudo et avec comme comptenue notre mots de passe
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			contenu = password + "\n" + "0,0,0,0,0,0,0,0,0,0,0,0,0,0,0" + "\n"
			creerFichierJoueur(pseudo, contenu)
		}
		//la on recupe tout nos donner du fichier text
		infos, err := SplitPlayerFile(filename) //la on recupe la tout le text du fichier wo.txt
		if err != nil {
			fmt.Println(err)
			return
		}
		// Vérifiez si le mot de passe correspond
		if strings.TrimSpace(infos[0]) == password {
			DonneeDeSerie(infos[1])
			//et on met statue de connection a true pour ne plus a avoir a se connecter jusqu'a la deconexion
			Statue = true
			//puis on redirige a l'acceuil
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} else {
			//sinon on garde le statue a false
			//et on affiche l'erreur sur la page avec ms
			Statue = false
			Ms = "Mot de passe incorrect."
		}
		//on cree notre data
		Data := createData(Statue, Ask, Answer, Page, Alllastopt, Ms, ResulteSerie, Serie)
		//puis on execute notre page avec le contenue assoscier
		tmpl.Execute(w, Data)
	}
}
