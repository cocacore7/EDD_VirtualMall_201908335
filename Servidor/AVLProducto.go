package Servidor

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type producto struct {
	Nombre string
	Codigo int
	Descripcion string
	Precio int
	Cantidad int
	Imagen string
}

type nodoproducto struct {
	producto producto
	Factor int
	id int
	izq *nodoproducto
	der *nodoproducto
}

type ArbolProducto struct {
	raiz *nodoproducto
}

func NewArbolProducto() *ArbolProducto{
	return &ArbolProducto{nil}
}

func NewNodoProducto(producto producto,id int) *nodoproducto{
	return &nodoproducto{producto,0,id,nil,nil}
}

func NewProducto(nombre string, codigo int, descripcion string, precio int, cantidad int, imagen string ) *producto{
	return &producto{nombre,codigo,descripcion,precio,cantidad,imagen}
}

func rotIIProducto(n *nodoproducto,n1 *nodoproducto) *nodoproducto{
	n.izq = n1.der
	n1.der = n
	if n1.Factor == -1{
		n.Factor = 0
		n1.Factor = 0
	}else{
		n.Factor = -1
		n1.Factor = 1
	}
	return n1
}

func rotDDProducto(n *nodoproducto, n1 *nodoproducto) *nodoproducto{
	n.der = n1.izq
	n1.izq = n
	if n1.Factor == 1{
		n.Factor = 0
		n1.Factor = 0
	}else{
		n.Factor = 1
		n1.Factor = -1
	}
	return n1
}

func rotDIProducto(n *nodoproducto, n1 *nodoproducto) *nodoproducto{
	n2:=n1.izq
	n.der = n2.izq
	n2.izq = n
	n1.izq = n2.der
	n2.der = n1
	if n2.Factor == 1{
		n.Factor = -1
	}else{
		n.Factor = 0
	}
	if n2.Factor == -1{
		n1.Factor = 1
	}else{
		n1.Factor = 0
	}
	n2.Factor = 0
	return n2
}

func rotIDProducto(n*nodoproducto,n1*nodoproducto)*nodoproducto{
	n2:= n1.der
	n.izq = n2.der
	n2.der = n
	n1.der = n2.izq
	n2.izq = n1
	if n2.Factor == 1{
		n1.Factor = -1
	}else{
		n1.Factor = 0
	}
	if n2.Factor == -1{
		n.Factor = 1
	}else{
		n.Factor = 0
	}
	n2.Factor = 0
	return n2
}

func insertarAVLProducto(raiz *nodoproducto, producto producto,id int, bandera *bool) *nodoproducto{
	var n1 *nodoproducto
	if raiz == nil{
		raiz = NewNodoProducto(producto, id)
		*bandera = true
	}else if producto.Codigo < raiz.producto.Codigo{
		izq:=insertarAVLProducto(raiz.izq,producto,id,bandera)
		raiz.izq = izq
		if *bandera{
			switch raiz.Factor {
			case 1:
				raiz.Factor = 0
				*bandera = false
				break
			case 0:
				raiz.Factor = -1
				break
			case -1:
				n1 = raiz.izq
				if n1.Factor == -1{
					raiz=rotIIProducto(raiz,n1)
				}else{
					raiz=rotIDProducto(raiz,n1)
				}
				*bandera=false
			}
		}
	}else if producto.Codigo > raiz.producto.Codigo{
		der:=insertarAVLProducto(raiz.der,producto,id,bandera)
		raiz.der =der
		if *bandera{
			switch raiz.Factor {
			case 1:
				n1 = raiz.der
				if n1.Factor == 1{
					raiz = rotDDProducto(raiz,n1)
				}else {
					raiz = rotDIProducto(raiz,n1)
				}
				*bandera = false
				break
			case 0:
				raiz.Factor = 1
				break
			case -1:
				raiz.Factor = 0
				*bandera = false
			}
		}
	}else if producto.Codigo == raiz.producto.Codigo{
		raiz.producto.Cantidad = raiz.producto.Cantidad + producto.Cantidad
		fmt.Println("Codigo: "+strconv.Itoa(producto.Codigo)+" Ya Existe En Arbol De Productos")
	}
	return raiz
}

func (this *ArbolProducto) InsertarAVLProducto(producto producto,id int) {
	b:=false
	a:=&b
	this.raiz = insertarAVLProducto(this.raiz,producto,id,a)
}

func graficar(n * ArbolProducto, bandera bool) bool{
	if n!=nil{
		if n.raiz!=nil{
			arch, _ := os.Create("ArbolProducto.dot")
			_, _ = arch.WriteString("digraph G{" + "\n")
			_, _ = arch.WriteString(`rankdir=UD;` + "\n")
			_, _ = arch.WriteString(`concentrate=true;` + "\n")
			_, _ = arch.WriteString(n.raiz.Interno())
			_, _ = arch.WriteString("}" + "\n")
			_ = arch.Close()
			path, _ := exec.LookPath("dot")
			cmd, _ := exec.Command(path, "-Tpng", "./ArbolProducto.dot").Output()
			mode := 0777
			_ = ioutil.WriteFile("ArbolProducto.png", cmd, os.FileMode(mode))
			bandera=true
		}
	}
	return bandera
}

func (this *nodoproducto) Interno() string{
	var etiqueta string
	etiqueta = "nodo"+strconv.Itoa(this.id)+"[shape=record,label=\"Factor: "+strconv.Itoa(this.Factor)+"|{Codigo: "+strconv.Itoa(this.producto.Codigo)+"|"+this.producto.Nombre+"|Precio: "+strconv.Itoa(this.producto.Precio)+"}|Cantidad: "+strconv.Itoa(this.producto.Cantidad)+"\"];\n"
	if this.izq != nil{
		etiqueta+=this.izq.Interno() + "nodo"+strconv.Itoa(this.id)+"->nodo"+strconv.Itoa(this.izq.id)+";\n"
	}
	if this.der != nil{
		etiqueta+=this.der.Interno() + "nodo"+strconv.Itoa(this.id)+"->nodo"+strconv.Itoa(this.der.id)+";\n"
	}
	return etiqueta
}

func ProuctosModificado(this *nodoproducto, productos []Productos) []Productos{
	if this.izq !=nil{
		productos = ProuctosModificado(this.izq,productos)
		productos = append(productos, Productos{Nombre: this.producto.Nombre, Codigo: this.producto.Codigo, Descripcion: this.producto.Descripcion, Precio: this.producto.Precio, Cantidad: this.producto.Cantidad, Imagen: this.producto.Imagen})
	}
	if this.der!=nil{
		productos = ProuctosModificado(this.der,productos)
		if validar(productos, true, this.producto.Codigo){
			p:=Productos{Nombre: this.producto.Nombre, Codigo: this.producto.Codigo, Descripcion: this.producto.Descripcion, Precio: this.producto.Precio, Cantidad: this.producto.Cantidad, Imagen: this.producto.Imagen}
			productos = append(productos, p)
		}
	}
	if this.izq ==nil && this.der == nil{
		productos = append(productos, Productos{Nombre: this.producto.Nombre, Codigo: this.producto.Codigo, Descripcion: this.producto.Descripcion, Precio: this.producto.Precio, Cantidad: this.producto.Cantidad, Imagen: this.producto.Imagen})
	}
	return productos
}

func validar(productos []Productos, bandera bool, codigo int)bool{
	for x:=0;x<len(productos);x++ {
		if productos[x].Codigo == codigo{
			bandera = false
			break
		}
	}
	return bandera
}

func buscarCodigoPedido(raiz *nodoproducto, codigo int, bandera bool) bool{
	if codigo < raiz.producto.Codigo{
		bandera = buscarCodigoPedido(raiz.izq,codigo,bandera)
	}else if codigo > raiz.producto.Codigo{
		bandera = buscarCodigoPedido(raiz.der,codigo,bandera)
	}else if codigo == raiz.producto.Codigo{
		bandera = true
	}
	return bandera
}

func ValidarExistencias(raiz *nodoproducto, codigo int, bandera bool) bool{
	if codigo < raiz.producto.Codigo{
		bandera = ValidarExistencias(raiz.izq,codigo,bandera)
	}else if codigo > raiz.producto.Codigo{
		bandera = ValidarExistencias(raiz.der,codigo,bandera)
	}else if codigo == raiz.producto.Codigo{
		if raiz.producto.Cantidad>0{return true}else{return false}
	}
	return bandera
}

func RestarStock(raiz *nodoproducto, codigo int) *nodoproducto{
	if codigo < raiz.producto.Codigo{
		izq := RestarStock(raiz.izq,codigo)
		raiz.izq = izq
	}else if codigo > raiz.producto.Codigo{
		der := RestarStock(raiz.der,codigo)
		raiz.der = der
	}else if codigo == raiz.producto.Codigo{
		raiz.producto.Cantidad--
	}
	return raiz
}

func SumarStock(raiz *nodoproducto, codigo int) *nodoproducto{
	if codigo < raiz.producto.Codigo{
		izq := SumarStock(raiz.izq,codigo)
		raiz.izq = izq
	}else if codigo > raiz.producto.Codigo{
		der := SumarStock(raiz.der,codigo)
		raiz.der = der
	}else if codigo == raiz.producto.Codigo{
		raiz.producto.Cantidad++
	}
	return raiz
}
