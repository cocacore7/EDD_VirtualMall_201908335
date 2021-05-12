package Servidor

import (
	"math"
	"time"
)

type NodoHash struct {
	hash 	int
	valor 	string
	Año 	int
	Mes 	int
	Dia 	int
	Hora 	int
	Minuto 	int
	Seg 	int
	subarreglo *HashTable
}

type HashTable struct{
	size int
	carga int
	porcentaje int
	arreglo []*NodoHash
}

func NewTable()*HashTable{
	arreglo:=make([]*NodoHash,7)
	return &HashTable{7,0,60,arreglo}
}

func(this *HashTable)insertar (nuevo int, valor string){
	var ahora time.Time
	ahora = time.Now()
	mes := obtenernumMes()
	nuevo_nodo:=NodoHash{nuevo,valor,ahora.Year(),mes,ahora.Day(),ahora.Hour(),ahora.Minute(),ahora.Second(),NewTable()}
	pos := this.posicion(nuevo, valor)
	this.arreglo[pos]=&nuevo_nodo
	this.carga++
	if((this.carga*100)/this.size)>this.porcentaje{
		sizenuevo:=this.size
		for{
			sizenuevo++
			correcto := esprimo(sizenuevo)
			if correcto{
				break
			}
		}
		nuevo_array:=make([]*NodoHash,sizenuevo)
		antiguo:=this.arreglo
		this.arreglo = nuevo_array
		this.size = sizenuevo
		aux:=0
		for i:=0;i<len(antiguo);i++ {
			if antiguo[i]!= nil{
				aux = this.posicion(antiguo[i].hash, antiguo[i].valor)
				nuevo_array[aux] = antiguo[i]
			}
		}
	}
}

func obtenernumMes() int {
	mes := 0
	switch time.Now().Month() {
	case time.January:
		mes = 1
		break
	case time.February:
		mes = 2
		break
	case time.March:
		mes = 3
		break
	case time.April:
		mes = 4
		break
	case time.May:
		mes = 5
		break
	case time.June:
		mes += 6
		break
	case time.July:
		mes += 7
		break
	case time.August:
		mes += 8
		break
	case time.September:
		mes += 9
		break
	case time.October:
		mes += 10
		break
	case time.November:
		mes += 11
		break
	case time.December:
		mes += 12
	}
	return mes
}

func esprimo(num int) bool {
	contDivisores := 0
	for i := 1; i <= int(num/2); i++ {
		if num%i == 0 {
			contDivisores++
		}
		if contDivisores > 1 {
			return false
		}
	}
	return true
}

func(this *HashTable)posicion(clave int, valor string) int{
	i:=0
	p:=int(math.Trunc(float64(this.size) * ((0.2520 * float64(clave)) - math.Trunc(0.2520*float64(clave)))))
	for this.arreglo[p]!=nil{
		i++
		p= p + int(math.Trunc(float64(this.size) * ((0.2520 * float64(i*i)) - math.Trunc(0.2520*float64(i*i)))))
		if p>=this.size{
			p = tamañoposicioncorrecto(p,this.size)
		}
	}
	return p
}

func tamañoposicioncorrecto(posactual int, tamañotabla int) int{
	if posactual>=tamañotabla{
		posactual = posactual - tamañotabla
		if posactual>=tamañotabla{
			posactual = tamañoposicioncorrecto(posactual,tamañotabla)
		}else{
			return posactual
		}
	}
	return posactual
}

func(this *HashTable)Buscar () []NodoHash{
	regreso := make([]NodoHash,0)
	for i := 0; i< this.size; i++ {
		if this.arreglo[i] != nil{
			regreso = append(regreso, *this.arreglo[i])
		}
	}
	if len(regreso) != 0{
		quicksort(regreso,0,len(regreso)-1)
	}
	return regreso
}

func quicksort(arreglo []NodoHash, start int, end int) {
	part := partition(arreglo, start, end)
	if (part - 1) > start {
		quicksort(arreglo, start, part-1)
	}
	if (part + 1) < end {
		quicksort(arreglo, part+1, end)
	}
}

func partition(arreglo []NodoHash, start int, end int) int {
	pivote := (arreglo)[end]
	panio := pivote.Año
	pmes := pivote.Mes
	pdia := pivote.Dia
	phora := pivote.Hora
	pmin := pivote.Minuto
	psec := pivote.Seg
	for i := start; i < end; i++ {
		anio := (arreglo)[i].Año
		mes := (arreglo)[i].Mes
		dia := (arreglo)[i].Dia
		hora := (arreglo)[i].Hora
		min := (arreglo)[i].Minuto
		sec := (arreglo)[i].Seg
		if anio < panio {
			tmp := (arreglo)[start]
			(arreglo)[start] = (arreglo)[i]
			(arreglo)[i] = tmp
			start++
		} else if anio == panio {
			if mes < pmes {
				tmp := (arreglo)[start]
				(arreglo)[start] = (arreglo)[i]
				(arreglo)[i] = tmp
				start++
			} else if mes == pmes {
				if dia < pdia {
					tmp := (arreglo)[start]
					(arreglo)[start] = (arreglo)[i]
					(arreglo)[i] = tmp
					start++
				} else if dia == pdia {
					if hora < phora {
						tmp := (arreglo)[start]
						(arreglo)[start] = (arreglo)[i]
						(arreglo)[i] = tmp
						start++
					} else if hora == phora {
						if min < pmin {
							tmp := (arreglo)[start]
							(arreglo)[start] = (arreglo)[i]
							(arreglo)[i] = tmp
							start++
						} else if min == pmin {
							if sec < psec {
								tmp := (arreglo)[start]
								(arreglo)[start] = (arreglo)[i]
								(arreglo)[i] = tmp
								start++
							}
						}
					}
				}
			}
		}
	}
	tmp := (arreglo)[start]
	(arreglo)[start] = pivote
	(arreglo)[end] = tmp
	return start
}
