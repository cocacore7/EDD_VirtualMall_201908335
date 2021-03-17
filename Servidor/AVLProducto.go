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
		fmt.Println("Codigo: "+strconv.Itoa(producto.Codigo)+" Ya Existe En Arbol De Productos")
	}
	return raiz
}

func (this *ArbolProducto) InsertarAVLProducto(producto producto,id int) {
	b:=false
	a:=&b
	this.raiz = insertarAVLProducto(this.raiz,producto,id,a)
}

func graficar(n * ArbolProducto, num string){
	arch, _ := os.Create("archivo"+num+".dot")
	_, _ = arch.WriteString("digraph G{" + "\n")
	_, _ = arch.WriteString(`rankdir=UD;` + "\n")
	_, _ = arch.WriteString(`concentrate=true;` + "\n")
	_, _ = arch.WriteString(n.raiz.Interno())
	_, _ = arch.WriteString("}" + "\n")
	_ = arch.Close()
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./archivo"+num+".dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("outfile"+num+".png", cmd, os.FileMode(mode))
}

func (this *nodoproducto) Interno() string{
	var etiqueta string
	etiqueta = "nodo"+strconv.Itoa(this.id)+"[label=\"{Codigo: "+this.producto.Nombre+"|"+strconv.Itoa(this.producto.Codigo)+"|"+strconv.Itoa(this.producto.Precio)+"}|Factor: "+strconv.Itoa(this.Factor)+"\"];\n"
	if this.izq != nil{
		etiqueta+=this.izq.Interno() + "nodo"+strconv.Itoa(this.id)+"->nodo"+strconv.Itoa(this.izq.id)+";\n"
	}
	if this.der != nil{
		etiqueta+=this.der.Interno() + "nodo"+strconv.Itoa(this.id)+"->nodo"+strconv.Itoa(this.der.id)+";\n"
	}
	return etiqueta
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
