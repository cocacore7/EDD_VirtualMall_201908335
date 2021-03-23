package Servidor

import (
	"io/ioutil"
	"os"
	"os/exec"
	"reflect"
	"strconv"
)

type NodoCabeceraVertical struct {
	Der, Arr, Aba, Izq  interface{}
	Categoria 			string
	nodog 				string
}

type NodoCabeceraHorizontal struct {
	Der, Arr, Aba, Izq  interface{}
	dia 				int
	nodog 				string
}

type Matriz struct {
	CabH *NodoCabeceraHorizontal
	CabV *NodoCabeceraVertical
}

func newMatriz() *Matriz{
	return &Matriz{nil,nil}
}

func (this *Matriz) getHorizontal(dia int) interface{}{
	if this.CabH == nil{
		return nil
	}
	var aux interface{} = this.CabH
	for aux != nil{
		if aux.(*NodoCabeceraHorizontal).dia == dia{
			return aux
		}
		aux = aux.(*NodoCabeceraHorizontal).Der
	}
	return nil
}

func (this *Matriz) getVertical(Categoria string) interface{}{
	if this.CabV == nil{
		return nil
	}
	var aux interface{} = this.CabV
	for aux != nil{
		if aux.(*NodoCabeceraVertical).Categoria == Categoria{
			return aux
		}
		aux = aux.(*NodoCabeceraVertical).Aba
	}
	return nil
}

func (this *Matriz) crearHorizontal(dia int, nodog string) *NodoCabeceraHorizontal{
	if this.CabH == nil{
		nueva:=&NodoCabeceraHorizontal{Der:nil, Izq: nil, Aba: nil, Arr: nil,dia: dia, nodog: nodog}
		this.CabH=nueva
		return nueva
	}
	var aux interface{} = this.CabH
	if dia < aux.(*NodoCabeceraHorizontal).dia{
		nueva:=&NodoCabeceraHorizontal{Der:nil, Izq: nil, Aba: nil, Arr: nil,dia: dia, nodog: nodog}
		nueva.Der = this.CabH
		this.CabH.Izq = nueva
		this.CabH = nueva
		return nueva
	}
	for aux.(*NodoCabeceraHorizontal).Der!=nil{
		if dia > aux.(*NodoCabeceraHorizontal).dia && dia < aux.(*NodoCabeceraHorizontal).Der.(NodoCabeceraHorizontal).dia{
			nueva:=&NodoCabeceraHorizontal{Der:nil, Izq: nil, Aba: nil, Arr: nil,dia: dia, nodog: nodog}
			tmp:= aux.(*NodoCabeceraHorizontal).Der
			tmp.(*NodoCabeceraHorizontal).Izq=nueva
			nueva.Der = tmp
			aux.(*NodoCabeceraHorizontal).Der=nueva
			nueva.Izq = aux
			return nueva
		}
		aux = aux.(*NodoCabeceraHorizontal).Der
	}
	nueva:=&NodoCabeceraHorizontal{Der:nil, Izq: nil, Aba: nil, Arr: nil,dia: dia, nodog: nodog}
	aux.(*NodoCabeceraHorizontal).Der=nueva
	nueva.Izq=aux
	return nueva
}

func (this *Matriz) crearVertical(Categoria string, nodog string) *NodoCabeceraVertical{
	if this.CabV == nil{
		nueva:=&NodoCabeceraVertical{Der:nil, Izq: nil, Aba: nil, Arr: nil,Categoria: Categoria, nodog: nodog}
		this.CabV = nueva
		return nueva
	}
	var aux interface{} = this.CabV
	if Categoria <= aux.(*NodoCabeceraVertical).Categoria{
		nueva:=&NodoCabeceraVertical{Der:nil, Izq: nil, Aba: nil, Arr: nil,Categoria: Categoria, nodog: nodog}
		nueva.Aba = this.CabV
		this.CabV.Arr = nueva
		this.CabV = nueva
		return nueva
	}
	for aux.(*NodoCabeceraVertical).Aba != nil{
		if Categoria > aux.(*NodoCabeceraVertical).Categoria && Categoria <= aux.(*NodoCabeceraVertical).Aba.(NodoCabeceraVertical).Categoria{
			nueva:=&NodoCabeceraVertical{Der:nil, Izq: nil, Aba: nil, Arr: nil,Categoria: Categoria, nodog: nodog}
			tmp:= aux.(*NodoCabeceraVertical).Aba
			tmp.(*NodoCabeceraVertical).Arr = nueva
			nueva.Aba = tmp
			aux.(*NodoCabeceraVertical).Aba = nueva
			nueva.Arr = aux
			return nueva
		}
		aux = aux.(*NodoCabeceraVertical).Aba
	}
	nueva:=&NodoCabeceraVertical{Der:nil, Izq: nil, Aba: nil, Arr: nil,Categoria: Categoria, nodog: nodog}
	aux.(*NodoCabeceraVertical).Aba=nueva
	nueva.Arr = aux
	return nueva
}

func (this *Matriz) obtenerUltimoV(cabecera *NodoCabeceraHorizontal, Categoria string) interface{}{
	if cabecera.Aba==nil{
		return cabecera
	}
	aux:=cabecera.Aba
	if Categoria <= aux.(*ArregloNodoPedido).Cola.primero.Pedido.Categoria{
		return cabecera
	}
	for aux.(*ArregloNodoPedido).Aba != nil{
		if Categoria > aux.(*ArregloNodoPedido).Cola.primero.Pedido.Categoria && Categoria < aux.(*ArregloNodoPedido).Aba.(*ArregloNodoPedido).Cola.primero.Pedido.Categoria{
			return aux
		}
		aux = aux.(*ArregloNodoPedido).Aba
	}
	if Categoria <= aux.(*ArregloNodoPedido).Cola.primero.Pedido.Categoria{
		return aux.(*ArregloNodoPedido).Arr
	}
	return aux
}

func (this *Matriz) obtenerUltimoH(cabecera *NodoCabeceraVertical, dia int) interface{}{
	if cabecera.Der==nil{
		return cabecera
	}
	aux:=cabecera.Der
	if dia <= aux.(*ArregloNodoPedido).Cola.primero.Pedido.dia{
		return cabecera
	}
	for aux.(*ArregloNodoPedido).Der != nil{
		if dia > aux.(*ArregloNodoPedido).Cola.primero.Pedido.dia && dia < aux.(*ArregloNodoPedido).Der.(*ArregloNodoPedido).Cola.primero.Pedido.dia{
			return aux
		}
		aux = aux.(*ArregloNodoPedido).Der
	}
	if dia <= aux.(*ArregloNodoPedido).Cola.primero.Pedido.dia{
		return aux.(*ArregloNodoPedido).Izq
	}
	return aux
}

func (this *Matriz) Agregar(nueva *Pedido){
	vertical:= this.getVertical(nueva.Categoria)
	horizontal:= this.getHorizontal(nueva.dia)
	if vertical == nil{
		vertical = this.crearVertical(nueva.Categoria,"n"+strconv.Itoa(NodoG))
		NodoG++
	}
	if horizontal == nil{
		horizontal = this.crearHorizontal(nueva.dia,"n"+strconv.Itoa(NodoG))
		NodoG++
	}
	izquierda:=this.obtenerUltimoH(vertical.(*NodoCabeceraVertical),nueva.dia)
	superior:= this.obtenerUltimoV(horizontal.(*NodoCabeceraHorizontal),nueva.Categoria)
	if reflect.TypeOf(izquierda).String()=="*Servidor.ArregloNodoPedido"{
		if izquierda.(*ArregloNodoPedido).Der == nil{
			if reflect.TypeOf(superior).String()=="*Servidor.ArregloNodoPedido"{
				if superior.(*ArregloNodoPedido).Cola.primero.Pedido.dia == izquierda.(*ArregloNodoPedido).Cola.primero.Pedido.dia{
					insertarPedido(nueva,izquierda.(*ArregloNodoPedido).Cola)
				}else{
					arreglonuevo:=newArregloPedidos("n"+strconv.Itoa(NodoG))
					insertarPedido(nueva,arreglonuevo.Cola)
					izquierda.(*ArregloNodoPedido).Der=arreglonuevo
					arreglonuevo.Izq = izquierda
				}
			}else{
				if superior.(*NodoCabeceraHorizontal).dia == izquierda.(*ArregloNodoPedido).Cola.primero.Pedido.dia{
					insertarPedido(nueva,izquierda.(*ArregloNodoPedido).Cola)
				}else{
					arreglonuevo:=newArregloPedidos("n"+strconv.Itoa(NodoG))
					insertarPedido(nueva,arreglonuevo.Cola)
					izquierda.(*ArregloNodoPedido).Der=arreglonuevo
					arreglonuevo.Izq = izquierda
				}
			}
		}else{
			tmp:= izquierda.(*ArregloNodoPedido).Der
			if reflect.TypeOf(superior).String()=="*Servidor.ArregloNodoPedido"{
				if superior.(*ArregloNodoPedido).Cola.primero.Pedido.dia == tmp.(*ArregloNodoPedido).Cola.primero.Pedido.dia{
					insertarPedido(nueva,tmp.(*ArregloNodoPedido).Cola)
				}else{
					arreglonuevo:=newArregloPedidos("n"+strconv.Itoa(NodoG))
					insertarPedido(nueva,arreglonuevo.Cola)
					izquierda.(*NodoCabeceraVertical).Der = arreglonuevo
					arreglonuevo.Izq = izquierda
					tmp.(*ArregloNodoPedido).Izq = arreglonuevo
					arreglonuevo.Der = tmp
				}
			}else{
				if superior.(*NodoCabeceraHorizontal).dia == tmp.(*ArregloNodoPedido).Cola.primero.Pedido.dia{
					insertarPedido(nueva,tmp.(*ArregloNodoPedido).Cola)
				}else{
					arreglonuevo:=newArregloPedidos("n"+strconv.Itoa(NodoG))
					insertarPedido(nueva,arreglonuevo.Cola)
					izquierda.(*NodoCabeceraVertical).Der = arreglonuevo
					arreglonuevo.Izq = izquierda
					tmp.(*ArregloNodoPedido).Izq = arreglonuevo
					arreglonuevo.Der = tmp
				}
			}
		}
	}else{
		if izquierda.(*NodoCabeceraVertical).Der == nil{
			arreglonuevo:=newArregloPedidos("n"+strconv.Itoa(NodoG))
			insertarPedido(nueva,arreglonuevo.Cola)
			izquierda.(*NodoCabeceraVertical).Der=arreglonuevo
			arreglonuevo.Izq = izquierda
		}else{
			tmp:=izquierda.(*NodoCabeceraVertical).Der
			if reflect.TypeOf(superior).String()=="*Servidor.ArregloNodoPedido"{
				if superior.(*ArregloNodoPedido).Cola.primero.Pedido.dia == tmp.(*ArregloNodoPedido).Cola.primero.Pedido.dia{
					insertarPedido(nueva,tmp.(*ArregloNodoPedido).Cola)
				}else{
					arreglonuevo:=newArregloPedidos("n"+strconv.Itoa(NodoG))
					insertarPedido(nueva,arreglonuevo.Cola)
					izquierda.(*NodoCabeceraVertical).Der = arreglonuevo
					arreglonuevo.Izq = izquierda
					tmp.(*ArregloNodoPedido).Izq = arreglonuevo
					arreglonuevo.Der = tmp
				}
			}else{
				if superior.(*NodoCabeceraHorizontal).dia == tmp.(*ArregloNodoPedido).Cola.primero.Pedido.dia{
					insertarPedido(nueva,tmp.(*ArregloNodoPedido).Cola)
				}else{
					arreglonuevo:=newArregloPedidos("n"+strconv.Itoa(NodoG))
					insertarPedido(nueva,arreglonuevo.Cola)
					izquierda.(*NodoCabeceraVertical).Der = arreglonuevo
					arreglonuevo.Izq = izquierda
					tmp.(*ArregloNodoPedido).Izq = arreglonuevo
					arreglonuevo.Der = tmp
				}
			}
		}
	}
	if reflect.TypeOf(superior).String()=="*Servidor.ArregloNodoPedido"{
		if superior.(*ArregloNodoPedido).Aba == nil{
			if reflect.TypeOf(izquierda).String()=="*Servidor.ArregloNodoPedido"{
				if izquierda.(*ArregloNodoPedido).Cola.primero.Pedido.Categoria == superior.(*ArregloNodoPedido).Cola.primero.Pedido.Categoria{
					insertarPedido(nueva,superior.(*ArregloNodoPedido).Cola)
				}else{
					arreglonuevo:=newArregloPedidos("n"+strconv.Itoa(NodoG))
					NodoG++
					insertarPedido(nueva,arreglonuevo.Cola)
					superior.(*ArregloNodoPedido).Aba=arreglonuevo
					arreglonuevo.Arr = superior
				}
			}else{
				if izquierda.(*NodoCabeceraVertical).Categoria == superior.(*ArregloNodoPedido).Cola.primero.Pedido.Categoria{
					insertarPedido(nueva,superior.(*ArregloNodoPedido).Cola)
				}else{
					arreglonuevo:=newArregloPedidos("n"+strconv.Itoa(NodoG))
					NodoG++
					insertarPedido(nueva,arreglonuevo.Cola)
					superior.(*ArregloNodoPedido).Aba=arreglonuevo
					arreglonuevo.Arr = superior
				}
			}
		}else{
			tmp:= superior.(*ArregloNodoPedido).Aba
			if reflect.TypeOf(izquierda).String()=="*Servidor.ArregloNodoPedido"{
				if izquierda.(*ArregloNodoPedido).Cola.primero.Pedido.Categoria == tmp.(*ArregloNodoPedido).Cola.primero.Pedido.Categoria{
					insertarPedido(nueva,tmp.(*ArregloNodoPedido).Cola)
				}else{
					arreglonuevo:=newArregloPedidos("n"+strconv.Itoa(NodoG))
					NodoG++
					insertarPedido(nueva,arreglonuevo.Cola)
					superior.(*NodoCabeceraVertical).Aba = arreglonuevo
					arreglonuevo.Arr = superior
					tmp.(*ArregloNodoPedido).Arr = arreglonuevo
					arreglonuevo.Aba = tmp
				}
			}else{
				if izquierda.(*NodoCabeceraVertical).Categoria == tmp.(*ArregloNodoPedido).Cola.primero.Pedido.Categoria{
					insertarPedido(nueva,tmp.(*ArregloNodoPedido).Cola)
				}else{
					arreglonuevo:=newArregloPedidos("n"+strconv.Itoa(NodoG))
					NodoG++
					insertarPedido(nueva,arreglonuevo.Cola)
					superior.(*NodoCabeceraHorizontal).Aba = arreglonuevo
					arreglonuevo.Arr = superior
					tmp.(*ArregloNodoPedido).Arr = arreglonuevo
					arreglonuevo.Aba = tmp
				}
			}
		}
	}else{
		if superior.(*NodoCabeceraHorizontal).Aba == nil{
			arreglonuevo:=newArregloPedidos("n"+strconv.Itoa(NodoG))
			NodoG++
			insertarPedido(nueva,arreglonuevo.Cola)
			superior.(*NodoCabeceraHorizontal).Aba=arreglonuevo
			arreglonuevo.Arr = superior
		}else{
			tmp:=superior.(*NodoCabeceraHorizontal).Aba
			if reflect.TypeOf(izquierda).String()=="*Servidor.ArregloNodoPedido"{
				if izquierda.(*ArregloNodoPedido).Cola.primero.Pedido.Categoria == tmp.(*ArregloNodoPedido).Cola.primero.Pedido.Categoria{
					insertarPedido(nueva,tmp.(*ArregloNodoPedido).Cola)
				}else{
					arreglonuevo:=newArregloPedidos("n"+strconv.Itoa(NodoG))
					NodoG++
					insertarPedido(nueva,arreglonuevo.Cola)
					superior.(*NodoCabeceraHorizontal).Aba = arreglonuevo
					arreglonuevo.Arr = superior
					tmp.(*ArregloNodoPedido).Arr = arreglonuevo
					arreglonuevo.Aba = tmp
				}
			}else{
				if izquierda.(*NodoCabeceraVertical).Categoria == tmp.(*ArregloNodoPedido).Cola.primero.Pedido.Categoria{
					insertarPedido(nueva,tmp.(*ArregloNodoPedido).Cola)
				}else{
					arreglonuevo:=newArregloPedidos("n"+strconv.Itoa(NodoG))
					NodoG++
					insertarPedido(nueva,arreglonuevo.Cola)
					superior.(*NodoCabeceraHorizontal).Aba = arreglonuevo
					arreglonuevo.Arr = superior
					tmp.(*ArregloNodoPedido).Arr = arreglonuevo
					arreglonuevo.Aba = tmp
				}
			}
		}
	}
}

func (this * Matriz)GraficarMatriz(){
	arch, _ := os.Create("GraficoMatriz.dot")
	_, _ = arch.WriteString("graph grid{" + "\n")
	_, _ = arch.WriteString("layout=dot" + "\n")
	_, _ = arch.WriteString("labelloc = \"t\"" + "\n")
	_, _ = arch.WriteString("node [shape=box]" + "\n")
	_, _ = arch.WriteString("edge [style=none]" + "\n")
	_, _ = arch.WriteString("p [shape=record]" + "\n")

	_, _ = arch.WriteString(this.getNodos())
	_, _ = arch.WriteString(this.getConexiones())
	_, _ = arch.WriteString(this.getRanked())

	_, _ = arch.WriteString("}" + "\n")
	arch.Close()

	//Crear Archivo
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./GraficoMatriz.dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("GraficoMatriz.png", cmd, os.FileMode(mode))
}

func (this *Matriz) getNodos() string{
	etiqueta := ""
	C:=1
	var aux interface{} = this.CabH
	if aux != nil{
		for aux != nil {
			etiqueta += aux.(*NodoCabeceraHorizontal).nodog+` [shape=record,color=red,label="` + strconv.Itoa(aux.(*NodoCabeceraHorizontal).dia) + `"];` + "\n"
			tmp := aux.(*NodoCabeceraHorizontal).Aba
			for tmp != nil {
				etiqueta += tmp.(*ArregloNodoPedido).nodog+` [shape=record,color=blue,label="C` + strconv.Itoa(C) + `"];` + "\n"
				C++
				tmp = tmp.(*ArregloNodoPedido).Aba
			}
			aux = aux.(*NodoCabeceraHorizontal).Der
		}
	}
	var aux2 interface{} = this.CabV
	if aux2 != nil{
		for aux2 != nil {
			etiqueta += aux2.(*NodoCabeceraVertical).nodog+` [shape=record,color=green,label="` + aux2.(*NodoCabeceraVertical).Categoria + `"];` + "\n"
			aux2 = aux2.(*NodoCabeceraVertical).Aba
		}
	}
	return etiqueta
}

func (this *Matriz) getConexiones() string{
	etiqueta := "p"
	var aux2 interface{} = this.CabV
	for aux2 != nil {
		etiqueta += "--"+aux2.(*NodoCabeceraVertical).nodog
		aux2 = aux2.(*NodoCabeceraVertical).Aba
	}
	etiqueta+="\n"
	var aux interface{} = this.CabH
	for aux != nil {
		etiqueta += aux.(*NodoCabeceraHorizontal).nodog
		tmp := aux.(*NodoCabeceraHorizontal).Aba
		for tmp != nil {
			etiqueta += "--"+tmp.(*ArregloNodoPedido).nodog
			tmp = tmp.(*ArregloNodoPedido).Aba
		}
		aux = aux.(*NodoCabeceraHorizontal).Der
		etiqueta+="\n"
	}
	return etiqueta
}

func (this *Matriz) getRanked() string{
	etiqueta := "rank=same {p"
	var aux2 interface{} = this.CabH
	for aux2 != nil {
		etiqueta += "--"+aux2.(*NodoCabeceraHorizontal).nodog
		aux2 = aux2.(*NodoCabeceraHorizontal).Der
	}
	etiqueta+="}\n"
	var aux interface{} = this.CabV
	for aux != nil {
		etiqueta += "rank=same {"+aux.(*NodoCabeceraVertical).nodog
		tmp := aux.(*NodoCabeceraVertical).Der
		for tmp != nil {
			etiqueta += "--"+tmp.(*ArregloNodoPedido).nodog
			tmp = tmp.(*ArregloNodoPedido).Der
		}
		aux = aux.(*NodoCabeceraVertical).Aba
		etiqueta+="}\n"
	}
	return etiqueta
}

func (this *Matriz) GraColaM(dia int, categoria string,bandera bool) bool{
	var aux interface{} = this.CabH
	for aux != nil {
		if dia == aux.(*NodoCabeceraHorizontal).dia{
			tmp := aux.(*NodoCabeceraHorizontal).Aba
			for tmp != nil {
				if tmp.(*ArregloNodoPedido).Cola.primero!=nil{
					if categoria == tmp.(*ArregloNodoPedido).Cola.primero.Pedido.Categoria{
						bandera = tmp.(*ArregloNodoPedido).Cola.GraficarPedidos()
						break
					}
				}
				tmp = tmp.(*ArregloNodoPedido).Aba
			}
			break
		}
		aux = aux.(*NodoCabeceraHorizontal).Der
	}
	return bandera
}
