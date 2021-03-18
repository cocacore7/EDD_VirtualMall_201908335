package Servidor

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var Cod int

//Cargar tiendas En Json
func cargar(w http.ResponseWriter, r *http.Request){
	var data Datos
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.Unmarshal(body, &data)
	_ = json.NewEncoder(w).Encode(data)
	Crear(data)
}

//Generar Grafos
func getArreglo(w http.ResponseWriter, _ *http.Request){
	reportes(0,0,1,0,0)
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusCreated)
}

//Buscar Tienda Especifica
func tiendaE(w http.ResponseWriter, r *http.Request){
	var t unico
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusFound)
	if vec!=nil{
		_ = json.Unmarshal(body, &t)
		a := posiciont(t)
		if a.Calificacion != 0{
			_ = json.NewEncoder(w).Encode(a)
		}else{
			_ = json.NewEncoder(w).Encode("No Se Encontro Ninguna Tienda Que Cumpla Con Los Parametros")
		}
	}else {
		_ = json.NewEncoder(w).Encode("No Hay Tiendas Cargadas")
	}
}

//Lista De tiendas
func tiendaN(w http.ResponseWriter, r *http.Request){
	var v varios
	vars:=mux.Vars(r)
	b,_:=strconv.Atoi(vars["numero"])
	a := posicionl(b)
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusFound)
	_ = json.Unmarshal(a, &v)
	_ = json.NewEncoder(w).Encode(v)
}

//Eliminar
func elim(w http.ResponseWriter, r *http.Request){
	var t unico2
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusOK)
	_ = json.Unmarshal(body, &t)
	if vec != nil{
		a:= Eliminar(t)
		var v varios
		_ = json.Unmarshal(a, &v)
		_ = json.NewEncoder(w).Encode(v)
	}else{
		_ = json.NewEncoder(w).Encode("No Hay Tiendas Cargadas")
	}
}

//Guardar
func guardar(w http.ResponseWriter, _ *http.Request){
	if vec != nil{
		var regreso Datos
		a := retorno(0,0)
		w.Header().Set("Content-Type","applicattion/json")
		w.WriteHeader(http.StatusAccepted)
		_ = json.Unmarshal(a, &regreso)
		_ = json.NewEncoder(w).Encode(regreso)
	}else{
		_ = json.NewEncoder(w).Encode("No Hay Tiendas Cargadas")
	}
}

//Inventario
func inven(w http.ResponseWriter, r *http.Request){
	var t Inventarios
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusOK)
	_ = json.Unmarshal(body, &t)
	if vec != nil{
		a:= Inventario(t)
		var v Inventarios
		_ = json.Unmarshal(a, &v)
		_ = json.NewEncoder(w).Encode(t)
	}else{
		_ = json.NewEncoder(w).Encode("No Hay Tiendas Cargadas")
	}
}

//Pedidos
func pedido(w http.ResponseWriter, r *http.Request){
	var t Pedidos
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusOK)
	_ = json.Unmarshal(body, &t)
	if vec != nil{
		a:= PedidosJson(t)
		var v Pedidos
		_ = json.Unmarshal(a, &v)
		_ = json.NewEncoder(w).Encode(v)
	}else{
		_ = json.NewEncoder(w).Encode("No Hay Tiendas Cargadas")
	}
}

//Pedidos Desde Carrito
func pedidoCarrito(w http.ResponseWriter, r *http.Request){
	var t Pedidos
	Cod = 0
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusOK)
	_ = json.Unmarshal(body, &t)
	if vec != nil{
		if ValidarPedidosJsonCarrito(t){
			a:= PedidosJsonCarrito(t)
			NoPedido++
			var v Pedidos
			_ = json.Unmarshal(a, &v)
			_ = json.NewEncoder(w).Encode(v)
		}else{
			if Cod == 0{
				_ = json.NewEncoder(w).Encode("No hay Tiendas Cargadas")
			}else{
				_ = json.NewEncoder(w).Encode("Pedido Rechazado, Producto con Codigo: "+strconv.Itoa(Cod)+" No cuenta con Existencia Solicitada")
			}

		}
	}else{
		_ = json.NewEncoder(w).Encode("No Hay Tiendas Cargadas")
	}
}

//Grafico Arbol De Años
func GraficoArbolAños(w http.ResponseWriter, r *http.Request){
	if añosArbol != nil{
		if añosArbol.raiz != nil{
			graficarAño(añosArbol)
			_, _ = fmt.Fprintf(w, "Arbol Generado")
		}else{
			_, _ = fmt.Fprintf(w, "No Existen Años Registrados")
		}
	}else{
		_, _ = fmt.Fprintf(w, "No Existen Años registrados")
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusFound)
}

//Grafico Meses De Arbol De Años
func GraficoMesesArbolAños(w http.ResponseWriter, r *http.Request){
	if añosArbol != nil{
		if añosArbol.raiz != nil{
			vars:=mux.Vars(r)
			b,_:=strconv.Atoi(vars["año"])
			if GraficarMeses(añosArbol.raiz,b,false){
				_, _ = fmt.Fprintf(w, "Arbol Generado")
			}else{
				_, _ = fmt.Fprintf(w, "No Existe Año Solicitado En Arbol")
			}
		}else{
			_, _ = fmt.Fprintf(w, "No Existen Años Registrados")
		}
	}else{
		_, _ = fmt.Fprintf(w, "No Existen Años registrados")
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusFound)
}

func Iniciar(){
	router := mux.NewRouter()
	router.HandleFunc("/guardar", guardar).Methods("GET")
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	router.HandleFunc("/id/{numero}", tiendaN).Methods("GET")
	router.HandleFunc("/GrafoAños", GraficoArbolAños).Methods("GET")
	router.HandleFunc("/GrafoMesesAños/{año}", GraficoMesesArbolAños).Methods("GET")
	router.HandleFunc("/cargartienda", cargar).Methods("POST")
	router.HandleFunc("/TiendaEspecifica", tiendaE).Methods("POST")
	router.HandleFunc("/cargarInventario", inven).Methods("POST")
	router.HandleFunc("/cargarPedido", pedido).Methods("POST")
	router.HandleFunc("/cargarPedidoCarrito", pedidoCarrito).Methods("POST")
	router.HandleFunc("/Eliminar", elim).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000",router))
}