package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Node struct {
	ID     int64    `json:"id"`
	Lat    float64  `json:"lat"`
	Lon    float64  `json:"lon"`
	Tags   map[string]string `json:"tags"`
}

type Response struct {
	Version     float64 `json:"version"`
	Generator   string  `json:"generator"`
	Nodes       []Node  `json:"elements"`
}

func main() {
	//selecionar los tipos de edificios
    amenities:=[]string{"restaurant","bar","school"}

    //generar la url
    url:="https://overpass-api.de/api/interpreter?data=[out:json];area[name=\"Zaragoza\"]->.a;"
    for _, v := range amenities {
        url+=fmt.Sprintf("node[\"amenity\"=\"%s\"][\"name\"](area.a);",v)
        
    }
    url+="out;"
    // Realizar la petición GET a la API de Overpass Turbo
    resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error al realizar la petición:", err)
		return
	}
	defer resp.Body.Close()

	// Leer la respuesta de la API y decodificarla en un struct Response
	var data Response
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
        fmt.Println(resp.Body)
		fmt.Println("Error al decodificar la respuesta:", err)
		return
	}

	// Imprimir los resultados
	for _, node := range data.Nodes {
		fmt.Println("ID:", node.ID)
		fmt.Println("Latitud:", node.Lat)
		fmt.Println("Longitud:", node.Lon)
		fmt.Println("Nombre:", node.Tags["name"])
		fmt.Println("Tipo:", node.Tags["amenity"])
		fmt.Println()
	}
}
