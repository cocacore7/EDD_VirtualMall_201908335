package Servidor

import (
	"bufio"
	"container/list"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type NodoMerkle struct {
	valor     string
	contenido string
	derecha   *NodoMerkle
	izquierda *NodoMerkle
}

type ArbolMerkle struct {
	raiz *NodoMerkle
}

func newNodoMerkle(valor string,contenido string, derecha *NodoMerkle, izquierda *NodoMerkle) *NodoMerkle {
	return &NodoMerkle{valor,contenido, derecha, izquierda}
}

func newArbolMerkle() *ArbolMerkle {
	return &ArbolMerkle{}
}

func (this *ArbolMerkle) Insertar(valor string,contenido string) {
	n := newNodoMerkle(valor, contenido, nil, nil)

	if this.raiz == nil {
		lis := list.New()

		lis.PushBack(n)
		h := sha256.New()
		h.Write([]byte(strconv.Itoa(-1)))
		sha256Sum := h.Sum(nil)
		cifrado := fmt.Sprintf("%x", sha256Sum)
		lis.PushBack(newNodoMerkle(strconv.Itoa(-1) , cifrado, nil, nil))
		this.construirArbol(lis)

	} else {
		lis := this.ObtenerLista()

		lis.PushBack(n)
		this.construirArbol(lis)
	}
}

func (this *ArbolMerkle) ObtenerLista() *list.List {
	lis := list.New()
	obtenerLista(lis, this.raiz.izquierda)
	obtenerLista(lis, this.raiz.derecha)
	return lis

}

func obtenerLista(lista *list.List, actual *NodoMerkle) {
	if actual != nil {
		obtenerLista(lista, actual.izquierda)
		if actual.derecha == nil && actual.valor != strconv.Itoa(-1) {
			lista.PushBack(actual)
		}
		obtenerLista(lista, actual.derecha)

	}
}

func (this *ArbolMerkle) construirArbol(lista *list.List) {
	size := float64(lista.Len())

	cantmerkle := 1

	for (size / 2) > 1 {
		cantmerkle++
		size = size / 2
	}

	nodostot := math.Pow(2, float64(cantmerkle))

	for lista.Len() < int(nodostot) {
		h := sha256.New()
		h.Write([]byte(strconv.Itoa(-1)))
		sha256Sum := h.Sum(nil)
		cifrado := fmt.Sprintf("%x", sha256Sum)
		lista.PushBack(newNodoMerkle(strconv.Itoa(-1) ,cifrado, nil, nil))

	}

	for lista.Len() > 1 {
		primero := lista.Front()
		segundo := primero.Next()

		lista.Remove(primero)
		lista.Remove(segundo)

		nodo1 := primero.Value.(*NodoMerkle)

		nodo2 := segundo.Value.(*NodoMerkle)

		h := sha256.New()
		h.Write([]byte(nodo1.contenido+nodo2.contenido))
		sha256Sum := h.Sum(nil)
		cifrado := fmt.Sprintf("%x", sha256Sum)
		nuevo := newNodoMerkle(nodo1.contenido+nodo2.contenido,cifrado, nodo2, nodo1)

		lista.PushBack(nuevo)

	}

	this.raiz = lista.Front().Value.(*NodoMerkle)

}

func Graficararbol(arbol *ArbolMerkle,nombre string) string {
	archivo, _ := os.Create("./" + nombre + ".dot")
	_, _ = archivo.WriteString(arbol.Codigo())
	archivo.Close()
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./"+nombre+".dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("./"+nombre+".png", cmd, os.FileMode(mode))

	f, _ := os.Open("./" + nombre + ".png")

	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	encoded := base64.StdEncoding.EncodeToString(content)

	return encoded
}

func (this *ArbolMerkle) Codigo() string{
	var cadena strings.Builder

	fmt.Fprintf(&cadena, "digraph G{\n")
	fmt.Fprintf(&cadena, "node[shape=\"record\"];\n")
	if this.raiz != nil {
		if (len(this.raiz.contenido)> 10) && (len(this.raiz.valor)> 10){
			fmt.Fprintf(&cadena, "node%p[label=\"<f0>|{<f1>%v | <f3>%v} | <f2>\"];\n", &(*this.raiz), this.raiz.valor[0:10],this.raiz.contenido[0:10])
		}else{
			fmt.Fprintf(&cadena, "node%p[label=\"<f0>|{<f1>%v | <f3>%v} | <f2>\"];\n", &(*this.raiz), this.raiz.valor,this.raiz.contenido[0:10])
		}

		this.generar(&cadena, (this.raiz), this.raiz.izquierda, true)
		this.generar(&cadena, this.raiz, this.raiz.derecha, false)
	}
	fmt.Fprintf(&cadena, "}\n")
	fmt.Println(cadena.String())
	return cadena.String()
}

func (this *ArbolMerkle) generar(cadena *strings.Builder, padre *NodoMerkle, actual *NodoMerkle, izquierda bool) {
	if actual != nil {
		if (len(actual.contenido)> 10) && (len(actual.valor)> 10){
			fmt.Fprintf(cadena, "node%p[label=\"<f0>|{<f1>%v | <f3>%v} | <f2>\"];\n", &(*actual),actual.contenido[0:10] ,actual.valor[0:10])
		}else{
			fmt.Fprintf(cadena, "node%p[label=\"<f0>|{<f1>%v | <f3>%v} | <f2>\"];\n", &(*actual),actual.contenido[0:10] ,actual.valor)
		}

		if izquierda {
			fmt.Fprintf(cadena, "node%p:f0->node%p:f1\n", &(*padre), &(*actual))
		} else {
			fmt.Fprintf(cadena, "node%p:f2->node%p:f1\n", &(*padre), &(*actual))
		}

		this.generar(cadena, actual, actual.izquierda, true)
		this.generar(cadena, actual, actual.derecha, false)
	}
}
