package main

import (
	"fmt"
	"html/template"
)

// chargement de la page
func loadTemplate(filename string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		fmt.Println("Error loading template:", err)
	}
	return tmpl, err
}

// et gestion de tout les donnee a renvoye
func createData(Statue bool, Ask string, Answer []string, Page string, Alllastopt []bool, Ms string, ResulteSerie int, Serie []string) interface{} {
	return struct {
		Statue       bool
		Ask          string
		Answer       []string
		Page         string
		Alllastopt   []bool
		Ms           string
		ResulteSerie int
		Values       []string
	}{
		Statue:       Statue,
		Ask:          Ask,
		Answer:       Answer,
		Page:         Page,
		Alllastopt:   Alllastopt,
		Ms:           Ms,
		ResulteSerie: ResulteSerie,
		Values:       Serie,
	}
}
