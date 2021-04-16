package Servidor

type key struct{
	Value 		Usuario
	Izquierdo 	*NodoArbolB
	Derecho 	*NodoArbolB
}

func NewKeyAB(valor Usuario) *key{
	k:=key{valor,nil,nil}
	return &k
}
