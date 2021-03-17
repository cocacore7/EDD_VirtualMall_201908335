package Servidor

import(
	"reflect"
)

type nodoPedido struct {
	dia 				int
	NoPedido 			int
	Tienda 				string
	Categoria 			string
	Calificacion 		int
	codigos 			[]int
}

type ArregloNodoPedido struct {
	Der, Arr, Aba, Izq  interface{}
	Pedidos []nodoPedido
}

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

func newPedido(dia int, noPedido int, Tienda string, Categoria string, Calificacion int, codigos []int) *nodoPedido{
	return &nodoPedido{dia,noPedido,Tienda,Categoria,Calificacion,codigos}
}

func newArregloPedidos() *ArregloNodoPedido{
	return &ArregloNodoPedido{nil,nil,nil,nil,nil}
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
	if Categoria <= aux.(*ArregloNodoPedido).Pedidos[0].Categoria{
		return cabecera
	}
	for aux.(*ArregloNodoPedido).Aba != nil{
		if Categoria > aux.(*ArregloNodoPedido).Pedidos[0].Categoria && Categoria < aux.(*ArregloNodoPedido).Aba.(*ArregloNodoPedido).Pedidos[0].Categoria{
			return aux
		}
		aux = aux.(*ArregloNodoPedido).Aba
	}
	if Categoria <= aux.(*nodoPedido).Categoria{
		return aux.(*ArregloNodoPedido).Arr
	}
	return aux
}

func (this *Matriz) obtenerUltimoH(cabecera *NodoCabeceraVertical, dia int) interface{}{
	if cabecera.Der==nil{
		return cabecera
	}
	aux:=cabecera.Der
	if dia <= aux.(*ArregloNodoPedido).Pedidos[0].dia{
		return cabecera
	}
	for aux.(*ArregloNodoPedido).Der != nil{
		if dia > aux.(*ArregloNodoPedido).Pedidos[0].dia && dia < aux.(*ArregloNodoPedido).Der.(*ArregloNodoPedido).Pedidos[0].dia{
			return aux
		}
		aux = aux.(*ArregloNodoPedido).Der
	}
	if dia <= aux.(*ArregloNodoPedido).Pedidos[0].dia{
		return aux.(*ArregloNodoPedido).Izq
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
	if reflect.TypeOf(izquierda).String()=="*Servidor.ArregloNodoPedido"{
		if izquierda.(*ArregloNodoPedido).Der == nil{
			if reflect.TypeOf(superior).String()=="*Servidor.ArregloNodoPedido"{
				if superior.(*ArregloNodoPedido).Pedidos[0].dia == izquierda.(*ArregloNodoPedido).Pedidos[0].dia{
					izquierda.(*ArregloNodoPedido).Pedidos = append(izquierda.(*ArregloNodoPedido).Pedidos, *nueva)
				}else{
					arreglonuevo:=newArregloPedidos()
					arreglonuevo.Pedidos = append(arreglonuevo.Pedidos, *nueva)
					izquierda.(*ArregloNodoPedido).Der=arreglonuevo
					arreglonuevo.Izq = izquierda
				}
			}else{
				if superior.(*NodoCabeceraHorizontal).dia == izquierda.(*ArregloNodoPedido).Pedidos[0].dia{
					izquierda.(*ArregloNodoPedido).Pedidos = append(izquierda.(*ArregloNodoPedido).Pedidos, *nueva)
				}else{
					arreglonuevo:=newArregloPedidos()
					arreglonuevo.Pedidos = append(arreglonuevo.Pedidos, *nueva)
					izquierda.(*ArregloNodoPedido).Der=arreglonuevo
					arreglonuevo.Izq = izquierda
				}
			}
		}else{
			tmp:= izquierda.(*ArregloNodoPedido).Der
			if reflect.TypeOf(superior).String()=="*Servidor.ArregloNodoPedido"{
				if superior.(*ArregloNodoPedido).Pedidos[0].dia == tmp.(*ArregloNodoPedido).Pedidos[0].dia{
					tmp.(*ArregloNodoPedido).Pedidos = append(tmp.(*ArregloNodoPedido).Pedidos, *nueva)
				}else{
					arreglonuevo:=newArregloPedidos()
					arreglonuevo.Pedidos = append(arreglonuevo.Pedidos, *nueva)
					izquierda.(*NodoCabeceraVertical).Der = arreglonuevo
					arreglonuevo.Izq = izquierda
					tmp.(*ArregloNodoPedido).Izq = arreglonuevo
					arreglonuevo.Der = tmp
				}
			}else{
				if superior.(*NodoCabeceraHorizontal).dia == tmp.(*ArregloNodoPedido).Pedidos[0].dia{
					tmp.(*ArregloNodoPedido).Pedidos = append(tmp.(*ArregloNodoPedido).Pedidos, *nueva)
				}else{
					arreglonuevo:=newArregloPedidos()
					arreglonuevo.Pedidos = append(arreglonuevo.Pedidos, *nueva)
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
			arreglonuevo.Pedidos = append(arreglonuevo.Pedidos, *nueva)
			izquierda.(*NodoCabeceraVertical).Der=arreglonuevo
			arreglonuevo.Izq = izquierda
		}else{
			tmp:=izquierda.(*NodoCabeceraVertical).Der
			if reflect.TypeOf(superior).String()=="*Servidor.ArregloNodoPedido"{
				if superior.(*ArregloNodoPedido).Pedidos[0].dia == tmp.(*ArregloNodoPedido).Pedidos[0].dia{
					tmp.(*ArregloNodoPedido).Pedidos = append(tmp.(*ArregloNodoPedido).Pedidos, *nueva)
				}else{
					arreglonuevo:=newArregloPedidos()
					arreglonuevo.Pedidos = append(arreglonuevo.Pedidos, *nueva)
					izquierda.(*NodoCabeceraVertical).Der = arreglonuevo
					arreglonuevo.Izq = izquierda
					tmp.(*ArregloNodoPedido).Izq = arreglonuevo
					arreglonuevo.Der = tmp
				}
			}else{
				if superior.(*NodoCabeceraHorizontal).dia == tmp.(*ArregloNodoPedido).Pedidos[0].dia{
					tmp.(*ArregloNodoPedido).Pedidos = append(tmp.(*ArregloNodoPedido).Pedidos, *nueva)
				}else{
					arreglonuevo:=newArregloPedidos()
					arreglonuevo.Pedidos = append(arreglonuevo.Pedidos, *nueva)
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
				if izquierda.(*ArregloNodoPedido).Pedidos[0].Categoria == superior.(*ArregloNodoPedido).Pedidos[0].Categoria{
					superior.(*ArregloNodoPedido).Pedidos = append(superior.(*ArregloNodoPedido).Pedidos, *nueva)
				}else{
					arreglonuevo:=newArregloPedidos()
					arreglonuevo.Pedidos = append(arreglonuevo.Pedidos, *nueva)
					superior.(*ArregloNodoPedido).Aba=arreglonuevo
					arreglonuevo.Arr = superior
				}
			}else{
				if izquierda.(*NodoCabeceraVertical).Categoria == superior.(*ArregloNodoPedido).Pedidos[0].Categoria{
					superior.(*ArregloNodoPedido).Pedidos = append(superior.(*ArregloNodoPedido).Pedidos, *nueva)
				}else{
					arreglonuevo:=newArregloPedidos()
					arreglonuevo.Pedidos = append(arreglonuevo.Pedidos, *nueva)
					superior.(*ArregloNodoPedido).Aba=arreglonuevo
					arreglonuevo.Arr = superior
				}
			}
		}else{
			tmp:= superior.(*ArregloNodoPedido).Aba
			if reflect.TypeOf(izquierda).String()=="*Servidor.ArregloNodoPedido"{
				if izquierda.(*ArregloNodoPedido).Pedidos[0].Categoria == tmp.(*ArregloNodoPedido).Pedidos[0].Categoria{
					tmp.(*ArregloNodoPedido).Pedidos = append(tmp.(*ArregloNodoPedido).Pedidos, *nueva)
				}else{
					arreglonuevo:=newArregloPedidos()
					arreglonuevo.Pedidos = append(arreglonuevo.Pedidos, *nueva)
					superior.(*NodoCabeceraVertical).Aba = arreglonuevo
					arreglonuevo.Arr = superior
					tmp.(*ArregloNodoPedido).Arr = arreglonuevo
					arreglonuevo.Aba = tmp
				}
			}else{
				if izquierda.(*NodoCabeceraVertical).Categoria == tmp.(*ArregloNodoPedido).Pedidos[0].Categoria{
					tmp.(*ArregloNodoPedido).Pedidos = append(tmp.(*ArregloNodoPedido).Pedidos, *nueva)
				}else{
					arreglonuevo:=newArregloPedidos()
					arreglonuevo.Pedidos = append(arreglonuevo.Pedidos, *nueva)
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
			arreglonuevo.Pedidos = append(arreglonuevo.Pedidos, *nueva)
			superior.(*NodoCabeceraHorizontal).Aba=arreglonuevo
			arreglonuevo.Arr = superior
		}else{
			tmp:=superior.(*NodoCabeceraHorizontal).Aba
			if reflect.TypeOf(izquierda).String()=="*Servidor.ArregloNodoPedido"{
				if izquierda.(*ArregloNodoPedido).Pedidos[0].Categoria == tmp.(*ArregloNodoPedido).Pedidos[0].Categoria{
					tmp.(*ArregloNodoPedido).Pedidos = append(tmp.(*ArregloNodoPedido).Pedidos, *nueva)
				}else{
					arreglonuevo:=newArregloPedidos()
					arreglonuevo.Pedidos = append(arreglonuevo.Pedidos, *nueva)
					superior.(*NodoCabeceraHorizontal).Aba = arreglonuevo
					arreglonuevo.Arr = superior
					tmp.(*ArregloNodoPedido).Arr = arreglonuevo
					arreglonuevo.Aba = tmp
				}
			}else{
				if izquierda.(*NodoCabeceraVertical).Categoria == tmp.(*ArregloNodoPedido).Pedidos[0].Categoria{
					tmp.(*ArregloNodoPedido).Pedidos = append(tmp.(*ArregloNodoPedido).Pedidos, *nueva)
				}else{
					arreglonuevo:=newArregloPedidos()
					arreglonuevo.Pedidos = append(arreglonuevo.Pedidos, *nueva)
					superior.(*NodoCabeceraHorizontal).Aba = arreglonuevo
					arreglonuevo.Arr = superior
					tmp.(*ArregloNodoPedido).Arr = arreglonuevo
					arreglonuevo.Aba = tmp
				}
			}
		}
	}
}
