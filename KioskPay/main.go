package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Lee el contenido del archivo HTML
	htmlContent, err := readFile("templates/index.html")
	if err != nil {
		http.Error(w, "Error al leer el archivo HTML", http.StatusInternalServerError)
		return
	}

	// Escribe el contenido en la respuesta HTTP
	tmpl, err := template.New("index").Parse(htmlContent)
	if err != nil {
		http.Error(w, "Error al analizar el archivo HTML", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func readFile(filename string) (string, error) {
	// Obtiene la ruta absoluta del archivo
	filePath, err := filepath.Abs(filename)
	if err != nil {
		return "", err
	}

	// Lee el contenido del archivo
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func main() {
	http.HandleFunc("/", handler)
	// Manejador para los archivos estáticos (CSS y JavaScript)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/templates/img/", http.StripPrefix("/templates/img/", http.FileServer(http.Dir("templates/img"))))

	//Abrir puerto local 8080
	fmt.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}