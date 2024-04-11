package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// ici on recupe tout les questions dans le fichier text
func ChooseAsks(Mode string, q int) (string, [][]interface{}) {
	//on decoupe chaque ensemble de questions reponces
	//gerne
	//Combien de doigts a une main ?
	//:
	//F-a) Deux
	//F-b) Quatre
	//V-c) Cinq
	//F-d) Six
	Asks, err := SplitTxT(Mode)
	//ensuit on decoupe chaque ensemble en une question et reponces
	//genre
	//V-a)aujourd'hui
	//F-b)demain
	//F-c)hier
	//F-d)apres demain
	Part, err := SplitQuest(Asks[q])
	//puis on decoupe tout l'ensemble de reponce en 4 reponce distincete genre V-a)aujourd'hui
	Answers, err := SplitAnswers(Part[1])
	//puis on decoupe chaque reponce en metant le V  et le a)aujourd'hui separement
	boolValue1, stringValue1 := SplitTF(Answers[0])
	boolValue2, stringValue2 := SplitTF(Answers[1])
	boolValue3, stringValue3 := SplitTF(Answers[2])
	boolValue4, stringValue4 := SplitTF(Answers[3])
	//on enregistre tous sa
	AnswerAll := [][]interface{}{
		{boolValue1, stringValue1},
		{boolValue2, stringValue2},
		{boolValue3, stringValue3},
		{boolValue4, stringValue4},
	}

	// puis on enregistre notre question dans une variable
	Ask := Part[0]
	if err != nil {
		fmt.Println("erreur de decoupe du fichier", err)
	}
	return Ask, AnswerAll

}

// decoupe du fichier text //
// ------------------------------------------------------------------------//
// ici on a tout nos split lié a la lecture des fichiers text
func SplitTxT(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename) //la on recupe la tout le text du fichier wo.txt
	if err != nil {
		fmt.Println("erreur de lecture du fichier", err)
	}
	Asks := strings.SplitN(string(data), "121", 10) //ici on fais en sorte que notre data devienne un string avec string(data) puis on le split a chaque . en 8 separation
	//puis on renvoie notre tableau

	return Asks, nil
}

func SplitQuest(Quest string) ([]string, error) {
	Part := strings.SplitN(string(Quest), ":", 2) //ici on fais en sorte que notre data devienne un string avec string(data) puis on le split a chaque . en 8 separation
	//puis on renvoie notre tableau
	return Part, nil
}

func SplitAnswers(Part string) ([]string, error) {

	if strings.HasSuffix(Part, "\n") {
		Part = strings.TrimLeft(Part, "\n")
	}
	Part = strings.TrimSpace(Part)

	Answers := strings.SplitN(string(Part), "\n", 4) //ici on fais en sorte que notre data devienne un string avec string(data) puis on le split a chaque . en 8 separation
	//puis on renvoie notre tableau
	return Answers, nil
}

func SplitTF(Answers string) (bool, string) {
	convert := strings.SplitN(Answers, "-", 2)
	stringValue := strings.TrimSpace(convert[1])
	bool1, err := stringToBool(convert[0])
	if err != nil {
		fmt.Println(err)
		return false, "" // Gestion de l'erreur en retournant une valeur par défaut
	}
	return bool1, stringValue
}

// ici c'est pour definir un true ou un false en fonction du V ou du  F qui est dans le texte
func stringToBool(s string) (bool, error) {
	switch strings.ToUpper(s) {
	case "V":
		return true, nil
	case "F":
		return false, nil
	default:
		return false, fmt.Errorf("La chaîne doit être 'V' ou 'F'")
	}
}

// ------------------------------------------------------------------------//

// ici on remet tout a zero
func PlayZero() {
	SerieNum = 0
	ResulteSerie = 0
	EssaySerie = 0
	q = 0
}

//-------------------------------------------------------------------------------------------//

// ----------------------------------------a metre dans  un fichier aparte------------------//

// ici c'est la meme chose que dans le hangman-web
func creerFichierJoueur(pseudo string, contenu string) {
	filename := "players/" + pseudo + ".txt"

	//la on le cree le fichier
	fichier, err := os.Create(filename)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier :", err)
		return
	}

	//on cree un writer (tampon)
	writer := bufio.NewWriter(fichier)
	//et on lui donne un comptenu a ecrire
	_, err = writer.WriteString(contenu)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier :", err)
		return
	}
	// on s'assure que toutes les données tamponnées sont écrites dans le fichier
	err = writer.Flush()
	if err != nil {
		fmt.Println("Erreur lors du vidage du tampon dans le fichier :", err)
		return
	}
	//ici le b c'est pour faire une difference de quand on cree un ficheier de quand on fais simplemnt un mise a jour
	if !b {
		fmt.Println("Compte créé avec succès :", pseudo, ",", password)
	}
	// on s'assure ensuite de ferme le fichier a la fin du main
	defer fichier.Close()
}
func supprimerFichier(pseudo string) {
	filename := "players/" + pseudo + ".txt"

	// Vérifiez si le fichier existe
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("Le fichier n'existe pas :", filename)
		return
	}
	//puis supprime
	err := os.Remove(filename)
	if err != nil {
		fmt.Println("Erreur lors de la suppression du fichier :", err)
		return
	}
	if !b {
		fmt.Println("Fichier supprimé avec succès :", filename)
	}

}

func mise() {
	b = true
	//donc d'abord on supprime l'anciene sauvegarde
	supprimerFichier(pseudo)
	//on convertie tout en string pour le metre en contenu a metre dans le .txt
	contenu := password + "\n" + Serie[0] + "," + Serie[1] + "," + Serie[2] + "," + Serie[3] + "," + Serie[3] + "," + Serie[4] + "," + Serie[5] + "," + Serie[6] + "," + Serie[7] + "," + Serie[8] + "," + Serie[9] + "," + Serie[10] + "," + Serie[11] + "," + Serie[12] + "," + Serie[14] + "\n"
	creerFichierJoueur(pseudo, contenu)
	fmt.Println("Sauvegarde créé avec succès :", pseudo, ",", password)
}

//-----------------------------------------------------------------------------//

// -----------------------------------------------------------------------------//
// ici c'est la recuperation des fichier joueur
func SplitPlayerFile(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename) //la on recupe la tout le text du fichier joueur.txt
	if err != nil {
		return nil, err
	}
	infos := strings.SplitN(string(data), "\n", 2) //ici on fais en sorte que notre data devienne un string avec string(data) puis on split tout les données
	//puis on renvoie notre tableau
	return infos, nil
}

// recupereation et separation des serie
func DonneeDeSerie(Donnee string) {
	Serie = SplitSerie(Donnee)
}

func SplitSerie(Donnee string) []string {
	SerieInfo := strings.SplitN(Donnee, ",", 15) //ici on fais en sorte que notre data devienne un string avec string(data) puis on le split a chaque . en 8 separation

	//puis on renvoie notre tableau
	return SerieInfo
}

// -----------------------------------------------------------------------------//
//----------------------------------------a metre dans  un fichier aparte------------------//
