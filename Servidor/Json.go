package Servidor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Data struct{
	Data []Datos `json:"Datos"`
}

type Datos struct{
	indice string `json:"Indice"`
	Departamentos []Departamentos `json:"Departamentos"`
}

type Departamentos struct {
	nombre string `json:"Nombre"`
	Tiendas []Tiendas `json:"Tiendas"`
}

type Tiendas struct {
	nombre string `json:"Nombre"`
	descripcion string `json:"Descripcion"`
	contacto string `json:"Contacto"`
	calificacion int `json:"Calificacion"`
}

func Leer(){
	archivo, err := os.Open("C:\\Users\\usuario\\OneDrive\\Escritorio\\primero.json")
	if err != nil{
		fmt.Println(err)
	}
	defer archivo.Close()

	valorB, _ := ioutil.ReadAll(archivo)
	var data Data


	err = json.Unmarshal(valorB, &data)
	if err != nil{
		log.Fatal(err)
	}

	for i := 0; i < len(data.Data); i++ {
		fmt.Println("Indice: " + data.Data[i].indice)
		for j:=0; j<len(data.Data[i].Departamentos);j++{
			fmt.Println("Departamento: " + data.Data[i].Departamentos[j].nombre)
			for x:=0;x<len(data.Data[i].Departamentos[j].Tiendas);x++{
				fmt.Println("Tienda: " + data.Data[i].Departamentos[j].Tiendas[x].nombre)
				fmt.Println("Contacto: " + data.Data[i].Departamentos[j].Tiendas[x].contacto)
				fmt.Println("Descripcion: " + data.Data[i].Departamentos[j].Tiendas[x].descripcion)
				fmt.Println("Calificacion: " + strconv.Itoa(data.Data[i].Departamentos[j].Tiendas[x].calificacion))
			}
		}
	}
}
