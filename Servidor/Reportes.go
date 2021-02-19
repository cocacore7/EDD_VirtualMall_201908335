package Servidor

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

var cant int
var canta int

//Recursividad Grafo
func reportes(i int, f int, n int, letra int, dep int){
	//Metodo Recursivo
	if i+5 <= len(vec){
		f = f + 5
		escribir(i,f,n,letra,dep)
	}else if i+4 <= len(vec){
		f = f + 4
		escribir(i,f,n,letra,dep)
	}else if i+3 <= len(vec){
		f = f + 3
		escribir(i,f,n,letra,dep)
	}else if i+2 <= len(vec){
		f = f + 2
		escribir(i,f,n,letra,dep)
	}else if i+1 <= len(vec){
		f = f + 1
		escribir(i,f,n,letra,dep)
	} else{
		return
	}
}

//Crear Grafo
func escribir(i int, f int, n int, letra int, dep int){
	arch, _ := os.Create("archivo" + strconv.Itoa(n) + ".dot")

	if dep == len(data.Datos[0].Departamentos){
		dep = 0
		letra++
	}

	//Encabezado
	_, _ = arch.WriteString("digraph G{" + "\n")
	_, _ = arch.WriteString(`compound=true;` + "\n")

	//Primer Subgrafo
	pos := i
	_, _ = arch.WriteString(`subgraph cluster0{` + "\n")
	_, _ = arch.WriteString(`edge[minlen=0.1, dir=fordware]` + "\n")
	_, _ = arch.WriteString(`struct1 [shape=record,label="` + data.Datos[letra].Indice + `|` + data.Datos[letra].Departamentos[dep].Departamentos + `|{ 1 | pos: ` + strconv.Itoa(pos) + `}"];` + "\n")
	_, _ = arch.WriteString(`struct2 [shape=record,label="` + data.Datos[letra].Indice + `|` + data.Datos[letra].Departamentos[dep].Departamentos + `|{ 2 | pos: ` + strconv.Itoa(pos+1) + `}"];` + "\n")
	_, _ = arch.WriteString(`struct3 [shape=record,label="` + data.Datos[letra].Indice + `|` + data.Datos[letra].Departamentos[dep].Departamentos + `|{ 3 | pos: ` + strconv.Itoa(pos+2) + `}"];` + "\n")
	_, _ = arch.WriteString(`struct4 [shape=record,label="` + data.Datos[letra].Indice + `|` + data.Datos[letra].Departamentos[dep].Departamentos + `|{ 4 | pos: ` + strconv.Itoa(pos+3) + `}"];` + "\n")
	_, _ = arch.WriteString(`struct5 [shape=record,label="` + data.Datos[letra].Indice + `|` + data.Datos[letra].Departamentos[dep].Departamentos + `|{ 5 | pos: ` + strconv.Itoa(pos+4) + `}"];` + "\n")
	_, _ = arch.WriteString(`struct1 -> struct2;` + "\n")
	_, _ = arch.WriteString(`struct2 -> struct3;` + "\n")
	_, _ = arch.WriteString(`struct3 -> struct4;` + "\n")
	_, _ = arch.WriteString(`struct4 -> struct5;` + "\n")
	_, _ = arch.WriteString("}" + "\n")

	if dep < len(data.Datos[0].Departamentos){ dep++ }

	var con []int
	in := i
	cant = 1
	if !vec[in].Vacio(){
		_, _ = arch.WriteString(`subgraph cluster1{` + "\n")
		_, _ = arch.WriteString(`edge[dir=both]` + "\n")

		//Creacion Tiendas
		aux := vec[in].primero
		con = append(con, cant)
		for aux != nil{
			_, _ = arch.WriteString("nodo" + strconv.Itoa(cant) + `[shape=record,label="{` + aux.tienda.nombre + " | " + aux.tienda.contacto + `}",fillcolor=red];` + "\n")
			cant++
			aux = aux.sig
		}

		//Cantidad de Tiendas
		contador := 0
		aux = vec[in].primero
		for aux != nil{
			contador++
			aux = aux.sig
		}

		//Conexiones Tiendas
		caux := cant-contador
		canta = caux
		aux = vec[in].primero
		conaux := 1
		for aux != nil{
			canta++
			if contador != 1 && conaux < contador{
				_, _ = arch.WriteString("nodo" + strconv.Itoa(caux) + " -> nodo" + strconv.Itoa(canta)  + "; \n")
				_, _ = arch.WriteString("nodo" + strconv.Itoa(canta) + " -> nodo" + strconv.Itoa(caux)  + "; \n")
			}
			conaux++
			caux++
			aux = aux.sig
		}
		_, _ = arch.WriteString("}" + "\n")
	}
	in++

	if !vec[in].Vacio(){
		_, _ = arch.WriteString(`subgraph cluster2{` + "\n")
		_, _ = arch.WriteString(`edge[dir=both]` + "\n")

		//Creacion Tiendas
		aux := vec[in].primero
		con = append(con, cant)
		for aux != nil{
			_, _ = arch.WriteString("nodo" + strconv.Itoa(cant) + `[shape=record,label="{` + aux.tienda.nombre + " | " + aux.tienda.contacto + `}",fillcolor=red];` + "\n")
			cant++
			aux = aux.sig
		}

		//Cantidad de Tiendas
		contador := 0
		aux = vec[in].primero
		for aux != nil{
			contador++
			aux = aux.sig
		}

		//Conexiones Tiendas
		caux := cant-contador
		canta = caux
		aux = vec[in].primero
		conaux := 1
		for aux != nil{
			canta++
			if contador != 1 && conaux < contador{
				_, _ = arch.WriteString("nodo" + strconv.Itoa(caux) + " -> nodo" + strconv.Itoa(canta)  + "; \n")
				_, _ = arch.WriteString("nodo" + strconv.Itoa(canta) + " -> nodo" + strconv.Itoa(caux)  + "; \n")
			}
			conaux++
			caux++
			aux = aux.sig
		}
		_, _ = arch.WriteString("}" + "\n")
	}
	in++

	if !vec[in].Vacio(){
		_, _ = arch.WriteString(`subgraph cluster3{` + "\n")
		_, _ = arch.WriteString(`edge[dir=both]` + "\n")

		//Creacion Tiendas
		aux := vec[in].primero
		con = append(con, cant)
		for aux != nil{
			_, _ = arch.WriteString("nodo" + strconv.Itoa(cant) + `[shape=record,label="{` + aux.tienda.nombre + " | " + aux.tienda.contacto + `}",fillcolor=red];` + "\n")
			cant++
			aux = aux.sig
		}

		//Cantidad de Tiendas
		contador := 0
		aux = vec[in].primero
		for aux != nil{
			contador++
			aux = aux.sig
		}

		//Conexiones Tiendas
		caux := cant-contador
		canta = caux
		aux = vec[in].primero
		conaux := 1
		for aux != nil{
			canta++
			if contador != 1 && conaux < contador{
				_, _ = arch.WriteString("nodo" + strconv.Itoa(caux) + " -> nodo" + strconv.Itoa(canta)  + "; \n")
				_, _ = arch.WriteString("nodo" + strconv.Itoa(canta) + " -> nodo" + strconv.Itoa(caux)  + "; \n")
			}
			conaux++
			caux++
			aux = aux.sig
		}
		_, _ = arch.WriteString("}" + "\n")
	}
	in++

	if !vec[in].Vacio(){
		_, _ = arch.WriteString(`subgraph cluster4{` + "\n")
		_, _ = arch.WriteString(`edge[dir=both]` + "\n")

		//Creacion Tiendas
		aux := vec[in].primero
		con = append(con, cant)
		for aux != nil{
			_, _ = arch.WriteString("nodo" + strconv.Itoa(cant) + `[shape=record,label="{` + aux.tienda.nombre + " | " + aux.tienda.contacto + `}",fillcolor=red];` + "\n")
			cant++
			aux = aux.sig
		}

		//Cantidad de Tiendas
		contador := 0
		aux = vec[in].primero
		for aux != nil{
			contador++
			aux = aux.sig
		}

		//Conexiones Tiendas
		caux := cant-contador
		canta = caux
		aux = vec[in].primero
		conaux := 1
		for aux != nil{
			canta++
			if contador != 1 && conaux < contador{
				_, _ = arch.WriteString("nodo" + strconv.Itoa(caux) + " -> nodo" + strconv.Itoa(canta)  + "; \n")
				_, _ = arch.WriteString("nodo" + strconv.Itoa(canta) + " -> nodo" + strconv.Itoa(caux)  + "; \n")
			}
			conaux++
			caux++
			aux = aux.sig
		}
		_, _ = arch.WriteString("}" + "\n")
	}
	in++

	if !vec[in].Vacio(){
		_, _ = arch.WriteString(`subgraph cluster5{` + "\n")
		_, _ = arch.WriteString(`edge[dir=both]` + "\n")

		//Creacion Tiendas
		aux := vec[in].primero
		con = append(con, cant)
		for aux != nil{
			_, _ = arch.WriteString("nodo" + strconv.Itoa(cant) + `[shape=record,label="{` + aux.tienda.nombre + " | " + aux.tienda.contacto + `}",fillcolor=red];` + "\n")
			cant++
			aux = aux.sig
		}

		//Cantidad de Tiendas
		contador := 0
		aux = vec[in].primero
		for aux != nil{
			contador++
			aux = aux.sig
		}

		//Conexiones Tiendas
		caux := cant-contador
		canta = caux
		aux = vec[in].primero
		conaux := 1
		for aux != nil{
			canta++
			if contador != 1 && conaux < contador{
				_, _ = arch.WriteString("nodo" + strconv.Itoa(caux) + " -> nodo" + strconv.Itoa(canta)  + "; \n")
				_, _ = arch.WriteString("nodo" + strconv.Itoa(canta) + " -> nodo" + strconv.Itoa(caux)  + "; \n")
			}
			conaux++
			caux++
			aux = aux.sig
		}
		_, _ = arch.WriteString("}" + "\n")
	}

	ultimo := 0
	in = i
	if !vec[in].Vacio(){
		_, _ = arch.WriteString("struct1 -> nodo"+ strconv.Itoa(con[ultimo]) + " [lhead=cluster1];" + "\n")
		ultimo++
	}
	in++
	if !vec[in].Vacio(){
		_, _ = arch.WriteString("struct2 -> nodo"+ strconv.Itoa(con[ultimo]) + " [lhead=cluster2];" + "\n")
		ultimo++
	}
	in++
	if !vec[in].Vacio(){
		_, _ = arch.WriteString("struct3 -> nodo"+ strconv.Itoa(con[ultimo]) + " [lhead=cluster3];" + "\n")
		ultimo++
	}
	in++
	if !vec[in].Vacio(){
		_, _ = arch.WriteString("struct4 -> nodo"+ strconv.Itoa(con[ultimo]) + " [lhead=cluster4];" + "\n")
		ultimo++
	}
	in++
	if !vec[in].Vacio(){
		_, _ = arch.WriteString("struct5 -> nodo"+ strconv.Itoa(con[ultimo]) + " [lhead=cluster5];" + "\n")
	}

	_, _ = arch.WriteString("}" + "\n")
	arch.Close()

	//Crear Archivo
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./archivo" + strconv.Itoa(n) + ".dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("outfile" + strconv.Itoa(n) + ".png", cmd, os.FileMode(mode))
	i = f
	n++
	//Recursividad
	reportes(i,f,n,letra,dep)
}

func retorno(in int, fi int) []byte{
	var reg Datos
	reg = data
	for i:=0;i<len(data.Datos);i++{
		for j:=0;j<len(data.Datos[i].Departamentos);j++{
			if in+5 <= len(vec){
				fi = fi + 5
			}else if in+4 <= len(vec){
				fi = fi + 4
			}else if in+3 <= len(vec){
				fi = fi + 3
			}else if in+2 <= len(vec){
				fi = fi + 2
			}else if in+1 <= len(vec){
				fi = fi + 1
			}
			reg.Datos[i].Departamentos[j].Tiendas = obtenerT(in, fi)
			in = fi
		}
	}
	arch, _ := os.Create("Salida.json")
	crearJson, _ := json.MarshalIndent(reg,"","    ")
	_, _ = arch.WriteString(string(crearJson))
	arch.Close()
	return crearJson
}

func obtenerT(i int, f int) []Tiendas{
	var t []Tiendas
	var aux Tiendas
	for i<f{
		a := vec[i].primero
		for a != nil{
			aux.Tiendas = a.tienda.nombre
			aux.Descripcion = a.tienda.descripcion
			aux.Contacto = a.tienda.contacto
			aux.Calificacion = a.tienda.calif
			t = append(t, aux)
			a = a.sig
		}
		i++
	}
	return t
}
