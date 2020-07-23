package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter().StrictSlash(true)

	route.HandleFunc("/", index)
	route.HandleFunc("/save", saveFile).Methods("POST")

	log.Println("Server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", route))

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("."+"/public/"))))
}

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./public/index.html")

	if err != nil {
		fmt.Println(err)
	}

	t.Execute(w, nil)
}

func saveFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		fmt.Println(err)
	} else if r.Method != "POST" {
		http.Redirect(w, r, "/", 200)
	}

	file := r.PostFormValue("archivo")

	exampleEncodeCoords(file)
}

func exampleEncodeCoords(file string) {
	coords := make([][]float64, 0)
	archivo, err := os.Open(file)

	if err != nil {
		fmt.Println(err)
	}

	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)

	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Location arrived") {
			result := strings.Fields(line)
			coords = append(coords, make([]float64, 0))

			rst := result[9]
			latitud, _ := strconv.ParseFloat(rst[2:13], 64)
			longitud, _ := strconv.ParseFloat(rst[14:26], 64)

			coords[count] = append(coords[count], latitud)
			coords[count] = append(coords[count], longitud)
			count++
		}

	}

	fmt.Println(coords)
}
