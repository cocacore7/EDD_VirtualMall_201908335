package Json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Datos struct{
	datos []indice `json:"Datos"`
}

type indice struct{
	indice string `json:"Indice"`
	departamento []depar `json:"Departamentos"`
}

type depar struct {
	nombre string `json:"Nombre"`
	tienda []tienda `json:"Tiendas"`
}

type tienda struct {
	nombre string `json:"Nombre"`
	descripcion string `json:"Descripcion"`
	contacto string `json:"Contacto"`
	calificacion int `json:"Calificacion"`
}

func Leer(){
	/*archivo, err := os.Open("C:\\Users\\usuario\\OneDrive\\Escritorio\\primero.json")
	if err != nil{
		fmt.Println(err)
	}
	defer archivo.Close()

	valorB, _ := ioutil.ReadAll(archivo)
	var datos Datos//map[string]interface{}


	err = json.Unmarshal(valorB, &datos)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(datos)*/

	archivo, err := ioutil.ReadFile("C:\\Users\\usuario\\OneDrive\\Escritorio\\primero.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var slice map[string]interface{}

	err = json.Unmarshal(archivo, &slice)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Datos: ",slice)
}
