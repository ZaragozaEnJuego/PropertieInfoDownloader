package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Node struct {
	ID   int64             `json:"id"`
	Lat  float64           `json:"lat"`
	Lon  float64           `json:"lon"`
	Tags map[string]string `json:"tags"`
}

type Response struct {
	Version   float64 `json:"version"`
	Generator string  `json:"generator"`
	Nodes     []Node  `json:"elements"`
}

type Edificio struct {
	Name       string  `json:"name"`
	Kind       string  `json:"kind"`
	Address    string  `json:"addres"`
	Lng        float64 `json:"lng"`
	Lat        float64 `json:"lat"`
	Price      int     `json:"price"`
	BaseIncome int     `json:"baseIncome"`
}

func main() {
	//selecionar los tipos de edificios
	amenities := []string{"restaurant", "bar", "school", "bus_station", "taxi", "library", "school", "university", "college", "clinic", "hospital", "pharmacy", "cafe", "fast_food"}

	var edificios []Edificio

	for _, am := range amenities {
		//generar la url

		url := fmt.Sprintf("https://overpass-api.de/api/interpreter?data=[out:json];area[name=\"Zaragoza\"]->.a;node[\"amenity\"=\"%s\"][\"name\"](area.a);out;", am)

		// Realizar la petición GET a la API de Overpass Turbo
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("Error al realizar la petición:", err)
			return
		}

		// Leer la respuesta de la API y decodificarla en un struct Response
		var data Response
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			fmt.Println(resp.Body)
			fmt.Println("Error al decodificar la respuesta:", err)
			return
		}

		for _, node := range data.Nodes {
			if node.Tags["addr:street"] != "" && node.Tags["addr:housenumber"] != "" {
				edificio := Edificio{
					Name:    node.Tags["name"],
					Kind:    node.Tags["amenity"],
					Address: node.Tags["addr:street"] + ", " + node.Tags["addr:housenumber"],
					Lng:     node.Lon,
					Lat:     node.Lat,
				}
				//TODO: add price and income
				//TODO: change kind to our kinds (transport, health, academic, groceries)
				edificios = append(edificios, edificio)

			}

		}
		resp.Body.Close()

	}
	// Convertir el slice en JSON
	jsonData, err := json.Marshal(edificios)
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return
	}
	// Escribir el JSON en un archivo
	file, err := os.Create("properties.json")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

}
