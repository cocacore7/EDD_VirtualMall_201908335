package Servidor

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type ArbolB struct{
	k 		int
	Raiz 	*NodoArbolB
}

func NewArbolB(nivel int) *ArbolB{
	a:= ArbolB{nivel,nil}
	nodoraiz:= NewNodoAB(nivel)
	a.Raiz = nodoraiz
	return &a
}

func(this ArbolB) InsertarAB(newkey *key, bandera bool) bool {
	if this.Raiz.keys[0] == nil{
		this.Raiz.Colocar(0,newkey)
	}else if this.Raiz.keys[0].Izquierdo == nil{
		lugarinsertado := -1
		node:=this.Raiz
		lugarinsertado = this.colocarNodoAB(node, newkey)
		if lugarinsertado == -5{
			fmt.Println("Usuario Con DPI: " + strconv.Itoa(newkey.Value.Dpi)+" Ya Existe")
			return false
		}else if lugarinsertado!=-1{
			if lugarinsertado == node.Max-1{
				medio:= node.Max/2
				llavecentral:=node.keys[medio]
				derecho:=NewNodoAB(this.k)
				izquierdo:=NewNodoAB(this.k)
				indiceizquierdo:=0
				indicederecho:=0
				for j:=0;j<node.Max;j++{
					//AQUI REVISAR
					if node.keys[j].Value.Dpi<llavecentral.Value.Dpi{
						izquierdo.Colocar(indiceizquierdo,node.keys[j])
						indiceizquierdo++
						node.Colocar(j,nil)
					}else if node.keys[j].Value.Dpi>llavecentral.Value.Dpi{
						derecho.Colocar(indicederecho,node.keys[j])
						indicederecho++
						node.Colocar(j,nil)
					}
				}
				node.Colocar(medio,nil)
				this.Raiz = node
				this.Raiz.Colocar(0,llavecentral)
				izquierdo.Padre = this.Raiz
				derecho.Padre = this.Raiz
				llavecentral.Izquierdo = izquierdo
				llavecentral.Derecho = derecho
			}
		}
	}else if this.Raiz.keys[0].Izquierdo!=nil{
		node:=this.Raiz
		for node.keys[0].Izquierdo!=nil{
			loop:=0
			for i:=0;i<node.Max;i,loop=i+1,loop+1{
				if node.keys[i]!=nil{
					if node.keys[i].Value.Dpi==newkey.Value.Dpi{
						fmt.Println("DPI: " + strconv.Itoa(newkey.Value.Dpi)+" Ya Existe")
						return false
					}else if node.keys[i].Value.Dpi>newkey.Value.Dpi{
						node = node.keys[i].Izquierdo
						break
					}
				}else{
					node=node.keys[i-1].Derecho
					break
				}
			}
			if loop == node.Max{
				node=node.keys[loop-1].Derecho
			}
		}
		indiceColocado:=this.colocarNodoAB(node,newkey)
		if indiceColocado == -5{
			fmt.Println("Usuario Con DPI: " + strconv.Itoa(newkey.Value.Dpi)+" Ya Existe")
			return false
		}else if indiceColocado==node.Max-1{
			for node.Padre!=nil{
				indicemedio:=node.Max/2
				llavecentral:=node.keys[indicemedio]
				izquierdo:=NewNodoAB(this.k)
				derecho:=NewNodoAB(this.k)
				indiceizquierdo,indicederecho:=0,0
				for i:=0;i<node.Max;i++{
					//AQUI REVISAR
					if node.keys[i].Value.Dpi<llavecentral.Value.Dpi{
						izquierdo.Colocar(indiceizquierdo,node.keys[i])
						indiceizquierdo++
						node.Colocar(i,nil)
					}else if node.keys[i].Value.Dpi>llavecentral.Value.Dpi{
						derecho.Colocar(indicederecho,node.keys[i])
						indicederecho++
						node.Colocar(i,nil)
					}
				}
				node.Colocar(indicemedio,nil)
				llavecentral.Izquierdo=izquierdo
				llavecentral.Derecho=derecho
				node= node.Padre
				izquierdo.Padre=node
				derecho.Padre=node
				for i:=0;i<izquierdo.Max;i++{
					if izquierdo.keys[i]!=nil{
						if izquierdo.keys[i].Izquierdo!=nil{
							izquierdo.keys[i].Izquierdo.Padre=izquierdo
						}
						if izquierdo.keys[i].Derecho!=nil{
							izquierdo.keys[i].Derecho.Padre=izquierdo
						}
					}
				}
				for i:=0;i<derecho.Max;i++{
					if derecho.keys[i]!=nil{
						if derecho.keys[i].Izquierdo!=nil{
							derecho.keys[i].Izquierdo.Padre=derecho
						}
						if derecho.keys[i].Derecho!=nil{
							derecho.keys[i].Derecho.Padre=derecho
						}
					}
				}
				lugarcolocado:=this.colocarNodoAB(node,llavecentral)
				if lugarcolocado == -5 {
					fmt.Println("Usuario Con DPI: " + strconv.Itoa(newkey.Value.Dpi)+" Ya Existe")
					return false
				}else if lugarcolocado==node.Max-1{
					if node.Padre==nil{
						indicecentralraiz:=node.Max/2
						llavecentralraiz:=node.keys[indicecentralraiz]
						izquierdoraiz:=NewNodoAB(this.k)
						derechoraiz:=NewNodoAB(this.k)
						indicederechoraiz, indiceizquierdoraiz:=0,0
						for i:=0;i<node.Max;i++{
							//AQUI REVISAR
							if node.keys[i].Value.Dpi<llavecentralraiz.Value.Dpi{
								izquierdoraiz.Colocar(indiceizquierdoraiz,node.keys[i])
								indiceizquierdoraiz++
								node.Colocar(i,nil)
							}else if node.keys[i].Value.Dpi>llavecentralraiz.Value.Dpi{
								derechoraiz.Colocar(indicederechoraiz,node.keys[i])
								indicederechoraiz++
								node.Colocar(i,nil)
							}
						}
						node.Colocar(indicecentralraiz,nil)
						node.Colocar(0,llavecentralraiz)
						for i:=0;i<this.k;i++{
							if izquierdoraiz.keys[i]!=nil{
								izquierdoraiz.keys[i].Izquierdo.Padre=izquierdoraiz
								izquierdoraiz.keys[i].Derecho.Padre=izquierdoraiz
							}
						}
						for i:=0;i<this.k;i++{
							if derechoraiz.keys[i]!=nil{
								derechoraiz.keys[i].Izquierdo.Padre=derechoraiz
								derechoraiz.keys[i].Derecho.Padre=derechoraiz
							}
						}
						llavecentralraiz.Izquierdo=izquierdoraiz
						llavecentralraiz.Derecho=derechoraiz
						izquierdoraiz.Padre=node
						derechoraiz.Padre=node
						this.Raiz=node
					}
					continue
				}else{
					break
				}
			}
		}
	}
	return bandera
}

func(this *ArbolB)colocarNodoAB(node *NodoArbolB,newkey *key) int{
	index := -1
	for i:=0;i<node.Max;i++{
		if node.keys[i] == nil{
			placed := false
			for j:=i-1;j>=0;j--{
				if node.keys[j].Value.Dpi==newkey.Value.Dpi{
					index = -5
					return index
				}else if node.keys[j].Value.Dpi>newkey.Value.Dpi{
					node.Colocar(j+1,node.keys[j])
				}else{
					node.Colocar(j+1,newkey)
					node.keys[j].Derecho=newkey.Izquierdo
					if (j+2)<this.k && node.keys[j+2]!=nil{
						node.keys[j+2].Izquierdo=newkey.Derecho
					}
					placed=true
					break
				}
			}
			if placed==false{
				node.Colocar(0,newkey)
				node.keys[1].Izquierdo=newkey.Derecho
			}
			index=i
			break
		}
	}
	return index
}

func(this ArbolB) BuscarAB(newkey *key) Usuario {
	if this.Raiz.keys[0].Izquierdo == nil{
		lugarinsertado := -1
		node:=this.Raiz
		lugarinsertado = this.buscarNodoAB(node, newkey)
		if lugarinsertado == -5{
			return this.buscarNodoAB2(node,newkey)
		}else if lugarinsertado == -6{
			return Usuario{Dpi: 0,Nombre: "NoContra",Correo: "",Password: "",Cuenta: ""}
		}
	}else if this.Raiz.keys[0].Izquierdo!=nil{
		node:=this.Raiz
		for node.keys[0].Izquierdo!=nil{
			loop:=0
			for i:=0;i<node.Max;i,loop=i+1,loop+1{
				if node.keys[i]!=nil{
					if node.keys[i].Value.Dpi==newkey.Value.Dpi{
						if node.keys[i].Value.Password==newkey.Value.Password{
							return this.buscarNodoAB2(node,newkey)
						}else{
							return Usuario{Dpi: 0,Nombre: "NoContra",Correo: "",Password: "",Cuenta: ""}
						}
					}else if node.keys[i].Value.Dpi>newkey.Value.Dpi{
						node = node.keys[i].Izquierdo
						break
					}
				}else{
					node=node.keys[i-1].Derecho
					break
				}
			}
			if loop == node.Max{
				node=node.keys[loop-1].Derecho
			}
		}
		indiceColocado:=this.buscarNodoAB(node,newkey)
		if indiceColocado == -5{
			return this.buscarNodoAB2(node,newkey)
		}else if indiceColocado == -6{
			return Usuario{Dpi: 0,Nombre: "NoContra",Correo: "",Password: "",Cuenta: ""}
		}
	}
	return Usuario{Dpi: 0,Nombre: "",Correo: "",Password: "",Cuenta: ""}
}

func(this *ArbolB)buscarNodoAB(node *NodoArbolB,newkey *key) int{
	index := -1
	for i:=0;i<node.Max;i++{
		if node.keys[i] == nil{
			for j:=i-1;j>=0;j--{
				if node.keys[j].Value.Dpi==newkey.Value.Dpi{
					if node.keys[j].Value.Password==newkey.Value.Password{
						index = -5
					}else{
						index = -6
					}
					return index
				}
			}
			index=i
			break
		}
	}
	return index
}

func(this *ArbolB)buscarNodoAB2(node *NodoArbolB,newkey *key) Usuario{
	for i:=0;i<node.Max;i++{
		if node.keys[i] == nil{
			for j:=i-1;j>=0;j--{
				if node.keys[j].Value.Dpi==newkey.Value.Dpi{
					if node.keys[j].Value.Password==newkey.Value.Password{
						return node.keys[j].Value
					}
					return Usuario{Dpi: 0,Nombre: "NoContra",Correo: "",Password: "",Cuenta: ""}
				}
			}
			break
		}
	}
	return Usuario{Dpi: 0,Nombre: "",Correo: "",Password: "",Cuenta: ""}
}

func (this *ArbolB) GraficarABSC() {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "digraph G{\nnode[shape=record]\nedge[color=\"green\"]\n")
	m := make(map[string]*NodoArbolB)
	graficarABSC(this.Raiz, &builder, m, nil, 0)
	fmt.Fprintf(&builder, "}")
	f, err := os.Create("ArbolSC.dot")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(builder.String())
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written succesfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./ArbolSC.dot").Output()
	mode := int(0772)
	ioutil.WriteFile("ArbolSC.png", cmd, os.FileMode(mode))
}

func graficarABSC(actual *NodoArbolB, cad *strings.Builder, arr map[string]*NodoArbolB, padre *NodoArbolB, pos int) {
	if actual == nil {
		return
	}
	j := 0
	contiene2 := arr[fmt.Sprint(&(*actual))]
	if contiene2 != nil {
		arr[fmt.Sprint(&(*actual))] = nil
		return
	} else {
		arr[fmt.Sprint(&(*actual))] = actual
	}
	fmt.Fprintf(cad, "node%p[color=\".7 .3 1.0\",label=\"", &(*actual))
	enlace := true
	for i := 0; i < actual.Max; i++ {
		if actual.keys[i] == nil {
			return
		} else {
			if enlace {
				if i != actual.Max-1 {
					fmt.Fprintf(cad, "<f%d>|", j)
				} else {
					fmt.Fprintf(cad, "<f%d>", j)
					break
				}
				enlace = false
				i--
				j++

			} else {
				fmt.Fprintf(cad, "{<f%d>DPI: %d|", j, actual.keys[i].Value.Dpi)
				fmt.Fprintf(cad, "Nombre: %s|", actual.keys[i].Value.Nombre)
				fmt.Fprintf(cad, "Correo: %s|", actual.keys[i].Value.Correo)
				fmt.Fprintf(cad, "Password: %s|", actual.keys[i].Value.Password)
				fmt.Fprintf(cad, "Cuenta: %s}|", actual.keys[i].Value.Cuenta)
				j++

				enlace = true
				if i < actual.Max-1 {
					if actual.keys[i+1] == nil {
						fmt.Fprintf(cad, "<f%d>", j)
						j++
						break
					}
				}
			}
		}
	}
	fmt.Fprintf(cad, "\"]\n")
	ji := 0
	for i := 0; i < actual.Max; i++ {
		if actual.keys[i] == nil {
			break
		}
		graficarABSC(actual.keys[i].Izquierdo, cad, arr, actual, ji)
		ji++
		ji++
		graficarABSC(actual.keys[i].Derecho, cad, arr, actual, ji)
		ji++
		ji--
	}
	if padre != nil {
		fmt.Fprintf(cad, "node%p:f%d->node%p\n", &(*padre), pos, &(*actual))
	}
}

func (this *ArbolB) GraficarABCS() {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "digraph G{\nnode[shape=record]\nedge[color=\"green\"]\n")
	m := make(map[string]*NodoArbolB)
	graficarABCS(this.Raiz, &builder, m, nil, 0)
	fmt.Fprintf(&builder, "}")
	f, err := os.Create("ArbolCS.dot")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(builder.String())
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written succesfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./ArbolCS.dot").Output()
	mode := int(0772)
	ioutil.WriteFile("ArbolCS.png", cmd, os.FileMode(mode))
}

func graficarABCS(actual *NodoArbolB, cad *strings.Builder, arr map[string]*NodoArbolB, padre *NodoArbolB, pos int) {
	if actual == nil {
		return
	}
	j := 0
	contiene2 := arr[fmt.Sprint(&(*actual))]
	if contiene2 != nil {
		arr[fmt.Sprint(&(*actual))] = nil
		return
	} else {
		arr[fmt.Sprint(&(*actual))] = actual
	}
	fmt.Fprintf(cad, "node%p[color=\"red\",label=\"", &(*actual))
	enlace := true
	for i := 0; i < actual.Max; i++ {
		if actual.keys[i] == nil {
			return
		} else {
			if enlace {
				if i != actual.Max-1 {
					fmt.Fprintf(cad, "<f%d>|", j)
				} else {
					fmt.Fprintf(cad, "<f%d>", j)
					break
				}
				enlace = false
				i--
				j++

			} else {
				cryptoPass := sha256.New()
				cryptoPass.Write([]byte(actual.keys[i].Value.Password))
				password := strings.ReplaceAll(base64.URLEncoding.EncodeToString(cryptoPass.Sum(nil)),"\"","\\\"")[0:10]
				password = strings.ReplaceAll(password,"}","\\")
				fmt.Fprintf(cad, "{<f%d>DPI: %d|", j, actual.keys[i].Value.Dpi)
				fmt.Fprintf(cad, "Nombre: %s|", actual.keys[i].Value.Nombre)
				fmt.Fprintf(cad, "Correo: %s|", actual.keys[i].Value.Correo)
				fmt.Fprintf(cad, "Password: %s|", password)
				fmt.Fprintf(cad, "Cuenta: %s}|", actual.keys[i].Value.Cuenta)
				j++

				enlace = true
				if i < actual.Max-1 {
					if actual.keys[i+1] == nil {
						fmt.Fprintf(cad, "<f%d>", j)
						j++
						break
					}
				}
			}
		}
	}
	fmt.Fprintf(cad, "\"]\n")
	ji := 0
	for i := 0; i < actual.Max; i++ {
		if actual.keys[i] == nil {
			break
		}
		graficarABCS(actual.keys[i].Izquierdo, cad, arr, actual, ji)
		ji++
		ji++
		graficarABCS(actual.keys[i].Derecho, cad, arr, actual, ji)
		ji++
		ji--
	}
	if padre != nil {
		fmt.Fprintf(cad, "node%p:f%d->node%p\n", &(*padre), pos, &(*actual))
	}
}

func (this *ArbolB) GraficarABC() {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "digraph G{\nnode[shape=record]\nedge[color=\"green\"]\n")
	m := make(map[string]*NodoArbolB)
	graficarABC(this.Raiz, &builder, m, nil, 0)
	fmt.Fprintf(&builder, "}")
	f, err := os.Create("ArbolC.dot")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(builder.String())
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written succesfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./ArbolC.dot").Output()
	mode := int(0772)
	ioutil.WriteFile("ArbolC.png", cmd, os.FileMode(mode))
}

func graficarABC(actual *NodoArbolB, cad *strings.Builder, arr map[string]*NodoArbolB, padre *NodoArbolB, pos int) {
	if actual == nil {
		return
	}
	j := 0
	contiene2 := arr[fmt.Sprint(&(*actual))]
	if contiene2 != nil {
		arr[fmt.Sprint(&(*actual))] = nil
		return
	} else {
		arr[fmt.Sprint(&(*actual))] = actual
	}
	fmt.Fprintf(cad, "node%p[color=\"blue\",label=\"", &(*actual))
	enlace := true
	for i := 0; i < actual.Max; i++ {
		if actual.keys[i] == nil {
			return
		} else {
			if enlace {
				if i != actual.Max-1 {
					fmt.Fprintf(cad, "<f%d>|", j)
				} else {
					fmt.Fprintf(cad, "<f%d>", j)
					break
				}
				enlace = false
				i--
				j++

			} else {
				cryptoPass := sha256.New()
				cryptoPass.Write([]byte(strconv.Itoa(actual.keys[i].Value.Dpi)))
				dpi := strings.ReplaceAll(base64.URLEncoding.EncodeToString(cryptoPass.Sum(nil)),"\"","\\\"")[0:10]
				dpi = strings.ReplaceAll(dpi,"}","\\")
				fmt.Fprintf(cad, "{<f%d>DPI: %d|", j, dpi)

				cryptoPass2 := sha256.New()
				cryptoPass2.Write([]byte(actual.keys[i].Value.Nombre))
				nombre := strings.ReplaceAll(base64.URLEncoding.EncodeToString(cryptoPass2.Sum(nil)),"\"","\\\"")[0:10]
				nombre = strings.ReplaceAll(nombre,"}","\\")
				fmt.Fprintf(cad, "Nombre: %s|", nombre)

				cryptoPass3 := sha256.New()
				cryptoPass3.Write([]byte(actual.keys[i].Value.Correo))
				correo := strings.ReplaceAll(base64.URLEncoding.EncodeToString(cryptoPass3.Sum(nil)),"\"","\\\"")[0:10]
				correo = strings.ReplaceAll(correo,"}","\\")
				fmt.Fprintf(cad, "Correo: %s|", correo)

				cryptoPass4 := sha256.New()
				cryptoPass4.Write([]byte(actual.keys[i].Value.Password))
				password := strings.ReplaceAll(base64.URLEncoding.EncodeToString(cryptoPass4.Sum(nil)),"\"","\\\"")[0:10]
				password = strings.ReplaceAll(password,"}","\\")
				fmt.Fprintf(cad, "Password: %s|", password)

				cryptoPass5 := sha256.New()
				cryptoPass5.Write([]byte(actual.keys[i].Value.Cuenta))
				cuenta := strings.ReplaceAll(base64.URLEncoding.EncodeToString(cryptoPass5.Sum(nil)),"\"","\\\"")[0:10]
				cuenta = strings.ReplaceAll(cuenta,"}","\\")
				fmt.Fprintf(cad, "Cuenta: %s}|",cuenta)

				j++

				enlace = true
				if i < actual.Max-1 {
					if actual.keys[i+1] == nil {
						fmt.Fprintf(cad, "<f%d>", j)
						j++
						break
					}
				}
			}
		}
	}
	fmt.Fprintf(cad, "\"]\n")
	ji := 0
	for i := 0; i < actual.Max; i++ {
		if actual.keys[i] == nil {
			break
		}
		graficarABC(actual.keys[i].Izquierdo, cad, arr, actual, ji)
		ji++
		ji++
		graficarABC(actual.keys[i].Derecho, cad, arr, actual, ji)
		ji++
		ji--
	}
	if padre != nil {
		fmt.Fprintf(cad, "node%p:f%d->node%p\n", &(*padre), pos, &(*actual))
	}
}
