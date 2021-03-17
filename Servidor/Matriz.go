package Servidor

import(
	"reflect"
)

type nodoPedido struct {
	Der, Arr, Aba, Izq interface{}
	dia int
	NoPedido int
	Tienda string
	Categoria string
	Calificacion int
	codigos []int
}

type NodoCabeceraVertical struct {
	Der, Arr, Aba, Izq interface{}
	Categoria string
}

type NodoCabeceraHorizontal struct {
	Der, Arr, Aba, Izq interface{}
	dia int
}

type Matriz struct {
	CabH *NodoCabeceraHorizontal
	CabV *NodoCabeceraVertical
}

func newPedido(dia int, noPedido int, Tienda string, Categoria string, Calificacion int, codigos []int) *nodoPedido{
	return &nodoPedido{nil,nil,nil,nil,dia,noPedido,Tienda,Categoria,Calificacion,codigos}
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

func (this *Matriz) crearHorizontal(dia int) *NodoCabeceraHorizontal{
	if this.CabH == nil{
		nueva:=&NodoCabeceraHorizontal{Der:nil, Izq: nil, Aba: nil, Arr: nil,dia: dia}
		this.CabH=nueva
		return nueva
	}
	var aux interface{} = this.CabH
	if dia < aux.(*NodoCabeceraHorizontal).dia{
		nueva:=&NodoCabeceraHorizontal{Der:nil, Izq: nil, Aba: nil, Arr: nil,dia: dia}
		nueva.Der = this.CabH
		this.CabH.Izq = nueva
		this.CabH = nueva
		return nueva
	}
	for aux.(*NodoCabeceraHorizontal).Der!=nil{
		if dia > aux.(*NodoCabeceraHorizontal).dia && dia < aux.(*NodoCabeceraHorizontal).Der.(NodoCabeceraHorizontal).dia{
			nueva:=&NodoCabeceraHorizontal{Der:nil, Izq: nil, Aba: nil, Arr: nil,dia: dia}
			tmp:= aux.(*NodoCabeceraHorizontal).Der
			tmp.(*NodoCabeceraHorizontal).Izq=nueva
			nueva.Der = tmp
			aux.(*NodoCabeceraHorizontal).Der=nueva
			nueva.Izq = aux
			return nueva
		}
		aux = aux.(*NodoCabeceraHorizontal).Der
	}
	nueva:=&NodoCabeceraHorizontal{Der:nil, Izq: nil, Aba: nil, Arr: nil,dia: dia}
	aux.(*NodoCabeceraHorizontal).Der=nueva
	nueva.Izq=aux
	return nueva
}

func (this *Matriz) crearVertical(Categoria string) *NodoCabeceraVertical{
	if this.CabV == nil{
		nueva:=&NodoCabeceraVertical{Der:nil, Izq: nil, Aba: nil, Arr: nil,Categoria: Categoria}
		this.CabV = nueva
		return nueva
	}
	var aux interface{} = this.CabV
	if Categoria <= aux.(*NodoCabeceraVertical).Categoria{
		nueva:=&NodoCabeceraVertical{Der:nil, Izq: nil, Aba: nil, Arr: nil,Categoria: Categoria}
		nueva.Aba = this.CabV
		this.CabV.Arr = nueva
		this.CabV = nueva
		return nueva
	}
	for aux.(*NodoCabeceraVertical).Aba != nil{
		if Categoria > aux.(*NodoCabeceraVertical).Categoria && Categoria <= aux.(*NodoCabeceraVertical).Aba.(NodoCabeceraVertical).Categoria{
			nueva:=&NodoCabeceraVertical{Der:nil, Izq: nil, Aba: nil, Arr: nil,Categoria: Categoria}
			tmp:= aux.(*NodoCabeceraVertical).Aba
			tmp.(*NodoCabeceraVertical).Arr = nueva
			nueva.Aba = tmp
			aux.(*NodoCabeceraVertical).Aba = nueva
			nueva.Arr = aux
			return nueva
		}
		aux = aux.(*NodoCabeceraVertical).Aba
	}
	nueva:=&NodoCabeceraVertical{Der:nil, Izq: nil, Aba: nil, Arr: nil,Categoria: Categoria}
	aux.(*NodoCabeceraVertical).Aba=nueva
	nueva.Arr = aux
	return nueva
}

func (this *Matriz) obtenerUltimoV(cabecera *NodoCabeceraHorizontal, Categoria string) interface{}{
	if cabecera.Aba==nil{
		return cabecera
	}
	aux:=cabecera.Aba
	if Categoria <= aux.(*nodoPedido).Categoria{
		return cabecera
	}
	for aux.(*nodoPedido).Aba != nil{
		if Categoria > aux.(*nodoPedido).Categoria && Categoria < aux.(*nodoPedido).Aba.(*nodoPedido).Categoria{
			return aux
		}
		aux = aux.(*nodoPedido).Aba
	}
	if Categoria <= aux.(*nodoPedido).Categoria{
		return aux.(*nodoPedido).Arr
	}
	return aux
}

func (this *Matriz) obtenerUltimoH(cabecera *NodoCabeceraVertical, dia int) interface{}{
	if cabecera.Der==nil{
		return cabecera
	}
	aux:=cabecera.Der
	if dia <= aux.(*nodoPedido).dia{
		return cabecera
	}
	for aux.(*nodoPedido).Der != nil{
		if dia > aux.(*nodoPedido).dia && dia < aux.(*nodoPedido).Der.(*nodoPedido).dia{
			return aux
		}
		aux = aux.(*nodoPedido).Der
	}
	if dia <= aux.(*nodoPedido).dia{
		return aux.(*nodoPedido).Izq
	}
	return aux
}

func (this *Matriz) Agregar(nueva *nodoPedido){
	vertical:= this.getVertical(nueva.Categoria)
	horizontal:= this.getHorizontal(nueva.dia)
	if vertical == nil{
		vertical = this.crearVertical(nueva.Categoria)
	}
	if horizontal == nil{
		horizontal = this.crearHorizontal(nueva.dia)
	}
	izquierda:=this.obtenerUltimoH(vertical.(*NodoCabeceraVertical),nueva.dia)
	superior:= this.obtenerUltimoV(horizontal.(*NodoCabeceraHorizontal),nueva.Categoria)
	if reflect.TypeOf(izquierda).String()=="*Servidor.nodoPedido"{
		if izquierda.(*nodoPedido).Der == nil{
			izquierda.(*nodoPedido).Der=nueva
			nueva.Izq = izquierda
		}else{
			tmp:= izquierda.(*nodoPedido).Der
			izquierda.(*nodoPedido).Der = nueva
			nueva.Izq = izquierda
			tmp.(*nodoPedido).Izq = nueva
			nueva.Der = tmp
		}
	}else{
		if izquierda.(*NodoCabeceraVertical).Der == nil{
			izquierda.(*NodoCabeceraVertical).Der=nueva
			nueva.Izq = izquierda
		}else{
			tmp:=izquierda.(*NodoCabeceraVertical).Der
			izquierda.(*NodoCabeceraVertical).Der = nueva
			nueva.Izq = izquierda
			tmp.(*nodoPedido).Izq = nueva
			nueva.Der = tmp
		}
	}
	if reflect.TypeOf(superior).String()=="*Servidor.nodoPedido"{
		if superior.(*nodoPedido).Aba == nil{
			superior.(*nodoPedido).Aba=nueva
			nueva.Arr = superior
		}else{
			tmp:= superior.(*nodoPedido).Aba
			superior.(*nodoPedido).Aba = nueva
			nueva.Arr = superior
			tmp.(*nodoPedido).Arr = nueva
			nueva.Aba = tmp
		}
	}else{
		if superior.(*NodoCabeceraHorizontal).Aba == nil{
			superior.(*NodoCabeceraHorizontal).Aba=nueva
			nueva.Arr = superior
		}else{
			tmp:=superior.(*NodoCabeceraHorizontal).Aba
			superior.(*NodoCabeceraHorizontal).Aba = nueva
			nueva.Arr = superior
			tmp.(*nodoPedido).Arr = nueva
			nueva.Aba = tmp
		}
	}
}
