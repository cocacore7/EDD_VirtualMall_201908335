package Servidor

type NodoArbolB struct {
	Max 	int
	Padre 	*NodoArbolB
	keys 	[]*key
}

func NewNodoAB(Max int) *NodoArbolB{
	keys:= make([]*key,Max)
	n:=NodoArbolB{Max,nil,keys}
	return &n
}

func (this *NodoArbolB)Colocar(i int, llave *key){
	this.keys[i] = llave
}