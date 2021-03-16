package Servidor


import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type Año struct {
	Año int
	mes *listaMes
}

type nodoAño struct {
	Año Año
	Factor int
	id int
	izq *nodoAño
	der *nodoAño
}

type ArbolAño struct {
	raiz *nodoAño
}

func NewArbolAño() *ArbolAño{
	return &ArbolAño{nil}
}

func NewNodoAño(año Año,id int) *nodoAño{
	return &nodoAño{año,0,id,nil,nil}
}

func NewAño(año int) *Año{
	return &Año{año, newListaMes()}
}

func rotIIAño(n *nodoAño,n1 *nodoAño) *nodoAño{
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

func rotDDAño(n *nodoAño, n1 *nodoAño) *nodoAño{
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

func rotDIAño(n *nodoAño, n1 *nodoAño) *nodoAño{
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

func rotIDAño(n*nodoAño,n1*nodoAño)*nodoAño{
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

func insertarAVLAño(raiz *nodoAño, año Año,id int, bandera *bool) *nodoAño{
	var n1 *nodoAño
	if raiz == nil{
		raiz = NewNodoAño(año, id)
		*bandera = true
	}else if año.Año < raiz.Año.Año{
		izq:=insertarAVLAño(raiz.izq,año,id,bandera)
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
					raiz=rotIIAño(raiz,n1)
				}else{
					raiz=rotIDAño(raiz,n1)
				}
				*bandera=false
			}
		}
	}else if año.Año > raiz.Año.Año{
		der:=insertarAVLAño(raiz.der,año,id,bandera)
		raiz.der =der
		if *bandera{
			switch raiz.Factor {
			case 1:
				n1 = raiz.der
				if n1.Factor == 1{
					raiz = rotDDAño(raiz,n1)
				}else {
					raiz = rotDIAño(raiz,n1)
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
	}else if año.Año == raiz.Año.Año{
		fmt.Println("Dato Repetido, Ya Existe En Arbol")
	}
	return raiz
}

func (this *ArbolAño) InsertarAVLAño(año Año,id int) {
	b:=false
	a:=&b
	this.raiz = insertarAVLAño(this.raiz,año,id,a)
}

func graficarAño(n * ArbolAño, num string){
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

func (this *nodoAño) Interno() string{
	var etiqueta string
	etiqueta = "nodo"+strconv.Itoa(this.id)+"[label=\"Año: "+strconv.Itoa(this.Año.Año)+"\"];\n"
	if this.izq != nil{
		etiqueta+=this.izq.Interno() + "nodo"+strconv.Itoa(this.id)+"->nodo"+strconv.Itoa(this.izq.id)+";\n"
	}
	if this.der != nil{
		etiqueta+=this.der.Interno() + "nodo"+strconv.Itoa(this.id)+"->nodo"+strconv.Itoa(this.der.id)+";\n"
	}
	return etiqueta
}

func insertarMesYPedido(raiz *nodoAño, año int,mes string, dia int, posv int,t tiendaPedido,codigos []int){
	if año < raiz.Año.Año{
		insertarMesYPedido(raiz.izq,año,mes,dia,posv,t,codigos)
	}else if año > raiz.Año.Año{
		insertarMesYPedido(raiz.izq,año,mes,dia,posv,t,codigos)
	}else if año == raiz.Año.Año{
		if raiz.Año.mes.primero != nil{
			aux := raiz.Año.mes.primero
			ban:=true
			for aux != nil{
				if aux.Mes.mes == mes{
					ban = false
					break
				}
				aux = aux.sig
			}
			if ban{
				insertarMes(newMes(mes),raiz.Año.mes)
			}
		}else{
			insertarMes(newMes(mes),raiz.Año.mes)
		}

		raiz.Año.mes = raiz.Año.mes.IngresarPedido(mes,dia,posv,t,codigos)
	}
}

