package Servidor

import(
	"reflect"
)

type NodoCabeceraVertical struct {
	Der, Arr, Aba, Izq  interface{}
	Categoria 			string
}

type NodoCabeceraHorizontal struct {
	Der, Arr, Aba, Izq  interface{}
	dia 				int
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
		vertical = this.crearVertical(nueva.Categoria)
	}
	if horizontal == nil{
		horizontal = this.crearHorizontal(nueva.dia)
	}
	izquierda:=this.obtenerUltimoH(vertical.(*NodoCabeceraVertical),nueva.dia)
	superior:= this.obtenerUltimoV(horizontal.(*NodoCabeceraHorizontal),nueva.Categoria)
	if reflect.TypeOf(izquierda).String()=="*Servidor.ArregloNodoPedido"{
		if izquierda.(*ArregloNodoPedido).Der == nil{
			if reflect.TypeOf(superior).String()=="*Servidor.ArregloNodoPedido"{
				if superior.(*ArregloNodoPedido).Cola.primero.Pedido.dia == izquierda.(*ArregloNodoPedido).Cola.primero.Pedido.dia{
					insertarPedido(nueva,izquierda.(*ArregloNodoPedido).Cola)
				}else{
					arreglonuevo:=newArregloPedidos()
					insertarPedido(nueva,arreglonuevo.Cola)
					izquierda.(*ArregloNodoPedido).Der=arreglonuevo
					arreglonuevo.Izq = izquierda
				}
			}else{
				if superior.(*NodoCabeceraHorizontal).dia == izquierda.(*ArregloNodoPedido).Cola.primero.Pedido.dia{
					insertarPedido(nueva,izquierda.(*ArregloNodoPedido).Cola)
				}else{
					arreglonuevo:=newArregloPedidos()
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
					arreglonuevo:=newArregloPedidos()
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
					arreglonuevo:=newArregloPedidos()
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
			arreglonuevo:=newArregloPedidos()
			insertarPedido(nueva,arreglonuevo.Cola)
			izquierda.(*NodoCabeceraVertical).Der=arreglonuevo
			arreglonuevo.Izq = izquierda
		}else{
			tmp:=izquierda.(*NodoCabeceraVertical).Der
			if reflect.TypeOf(superior).String()=="*Servidor.ArregloNodoPedido"{
				if superior.(*ArregloNodoPedido).Cola.primero.Pedido.dia == tmp.(*ArregloNodoPedido).Cola.primero.Pedido.dia{
					insertarPedido(nueva,tmp.(*ArregloNodoPedido).Cola)
				}else{
					arreglonuevo:=newArregloPedidos()
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
					arreglonuevo:=newArregloPedidos()
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
					arreglonuevo:=newArregloPedidos()
					insertarPedido(nueva,arreglonuevo.Cola)
					superior.(*ArregloNodoPedido).Aba=arreglonuevo
					arreglonuevo.Arr = superior
				}
			}else{
				if izquierda.(*NodoCabeceraVertical).Categoria == superior.(*ArregloNodoPedido).Cola.primero.Pedido.Categoria{
					insertarPedido(nueva,superior.(*ArregloNodoPedido).Cola)
				}else{
					arreglonuevo:=newArregloPedidos()
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
					arreglonuevo:=newArregloPedidos()
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
					arreglonuevo:=newArregloPedidos()
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
			arreglonuevo:=newArregloPedidos()
			insertarPedido(nueva,arreglonuevo.Cola)
			superior.(*NodoCabeceraHorizontal).Aba=arreglonuevo
			arreglonuevo.Arr = superior
		}else{
			tmp:=superior.(*NodoCabeceraHorizontal).Aba
			if reflect.TypeOf(izquierda).String()=="*Servidor.ArregloNodoPedido"{
				if izquierda.(*ArregloNodoPedido).Cola.primero.Pedido.Categoria == tmp.(*ArregloNodoPedido).Cola.primero.Pedido.Categoria{
					insertarPedido(nueva,tmp.(*ArregloNodoPedido).Cola)
				}else{
					arreglonuevo:=newArregloPedidos()
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
					arreglonuevo:=newArregloPedidos()
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
