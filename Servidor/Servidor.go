package Servidor

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var Cod int
var Gaño int
var Gmes string

//---------------------------------------------------------------------------------------------
//FASE1

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
	if data.Datos != nil{
		Crear(data)
		if vec != nil{
			_ = json.NewEncoder(w).Encode(data)
		}else{
			_ = json.NewEncoder(w).Encode("Archivo Vacio")
		}
	}else{
		_ = json.NewEncoder(w).Encode("Estructura Del Archivo Incorrecta")
	}



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

//FASE1
//---------------------------------------------------------------------------------------------

//---------------------------------------------------------------------------------------------
//FASE2

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
		if v.Inventarios == nil{
			_ = json.NewEncoder(w).Encode("Archivo Con Estructura Incorrecta")
		}else{
			_ = json.NewEncoder(w).Encode(v)
		}
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
		if v.Pedidos == nil{
			_ = json.NewEncoder(w).Encode("Archivo Con Estructura Incorrecta")
		}else{
			_ = json.NewEncoder(w).Encode(v)
		}

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
			if v.Pedidos == nil{
				_ = json.NewEncoder(w).Encode("Archivo Con Estructura Incorrecta")
			}else{
				_ = json.NewEncoder(w).Encode(v)
			}
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
	Gaño=0
	Gmes=""
	if añosArbol != nil{
		if añosArbol.raiz != nil{
			graficarAño(añosArbol)
			f, _ := os.Open("./ArbolAños.png")
			reader := bufio.NewReader(f)
			contenido, _ := ioutil.ReadAll(reader)
			encoded := base64.StdEncoding.EncodeToString(contenido)
			_, _ = fmt.Fprintf(w, encoded)
			f.Close()
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
			Gaño,_=strconv.Atoi(vars["año"])
			if GraficarMeses(añosArbol.raiz,Gaño,false){
				f, _ := os.Open("./GraficoMeses.png")
				reader := bufio.NewReader(f)
				contenido, _ := ioutil.ReadAll(reader)
				encoded := base64.StdEncoding.EncodeToString(contenido)
				_, _ = fmt.Fprintf(w, encoded)
				f.Close()
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

//Grafico Meses De Arbol De Años

func GraficoMatrizMesesArbolAños(w http.ResponseWriter, r *http.Request){
	if añosArbol != nil{
		if añosArbol.raiz != nil{
			vars:=mux.Vars(r)
			Gmes=vars["Mes"]
			if Gaño != 0{
				if GraficarMatrizMeses(añosArbol.raiz,Gaño,Gmes,false){
					f, _ := os.Open("./GraficoMatriz.png")
					reader := bufio.NewReader(f)
					contenido, _ := ioutil.ReadAll(reader)
					encoded := base64.StdEncoding.EncodeToString(contenido)
					_, _ = fmt.Fprintf(w, encoded)
					f.Close()
				}else{
					_, _ = fmt.Fprintf(w, "No Existen Matriz Solicitada, Ya que No Habian Productos Cargados"+"\n"+"Carga Productos Y Volver a Cargar Pedido")
				}
			}else {
				_, _ = fmt.Fprintf(w, "No Existe Grafica De Mes Creada")
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

//Grafico Meses De Arbol De Años

func GraficoColaMatrizMesesArbolAños(w http.ResponseWriter, r *http.Request){
	if añosArbol != nil{
		if añosArbol.raiz != nil{
			vars:=mux.Vars(r)
			a,_:=strconv.Atoi(vars["dia"])
			vars2:=mux.Vars(r)
			b:=vars2["cat"]
			if Gaño != 0{
				if Gmes != ""{
					if GraficarColaMatrizMeses(añosArbol.raiz,Gaño,Gmes,a,b,false){
						f, _ := os.Open("./GraficoColaMatriz.png")
						reader := bufio.NewReader(f)
						contenido, _ := ioutil.ReadAll(reader)
						encoded := base64.StdEncoding.EncodeToString(contenido)
						_, _ = fmt.Fprintf(w, encoded)
						f.Close()
					}else{
						_, _ = fmt.Fprintf(w, "Grafico No Generado, Dato Erroneos")
					}
				}else{
					_, _ = fmt.Fprintf(w, "No Existe Grafico Matriz")
				}
			}else {
				_, _ = fmt.Fprintf(w, "No Existe Grafica De Mes Creada")
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

//Grafico Arbol De Productos

func GraficoArbolProductos(w http.ResponseWriter, r *http.Request){
	var t unico
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	if vec!=nil{
		_ = json.Unmarshal(body, &t)
		if graficarArbolP(t,false){
			f, _ := os.Open("./ArbolProducto.jpg")
			reader := bufio.NewReader(f)
			contenido, _ := ioutil.ReadAll(reader)
			encoded := base64.StdEncoding.EncodeToString(contenido)
			_ = json.NewEncoder(w).Encode(encoded)
			f.Close()
		}else {
			_ = json.NewEncoder(w).Encode("No Existen Productos Cargados En Tienda Solicitada")
		}
	}else {
		_ = json.NewEncoder(w).Encode("No Hay Tiendas Cargadas")
	}
	//w.Header().Set("Content-Type","applicattion/json")
	//w.WriteHeader(http.StatusFound)
}

//GuardarProductos
func guardarProductos(w http.ResponseWriter, _ *http.Request){
	if vec != nil{
		var regreso Inventarios
		a := RegresoProductos()
		w.Header().Set("Content-Type","applicattion/json")
		w.WriteHeader(http.StatusAccepted)
		_ = json.Unmarshal(a, &regreso)
		_ = json.NewEncoder(w).Encode(regreso)
	}else{
		_ = json.NewEncoder(w).Encode("No Hay Tiendas Cargadas")
	}
}

//FASE2
//---------------------------------------------------------------------------------------------

//---------------------------------------------------------------------------------------------
//FASE3

//Usuarios
func usua(w http.ResponseWriter, r *http.Request){
	var t Usuarios
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusOK)
	_ = json.Unmarshal(body, &t)
	a:= AgregarUsuarios(t)
	var v Usuarios
	_ = json.Unmarshal(a, &v)
	if v.Usuarios == nil{
		_ = json.NewEncoder(w).Encode("Archivo Con Estructura Incorrecta")
	}else{
		_ = json.NewEncoder(w).Encode(v)
	}
}

func crearusua(w http.ResponseWriter, r *http.Request){
	var t Usuario
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusOK)
	_ = json.Unmarshal(body, &t)
	a:= CrearUsuario(t)
	if a{
		_ = json.NewEncoder(w).Encode("Usuario Ingresado Con Exito")
	}else{
		_ = json.NewEncoder(w).Encode("El DPI Ya Existe, Ingrese uno distinto")
	}
}

func buscarusua(w http.ResponseWriter, r *http.Request) {
	var t Usuario
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type", "applicattion/json")
	w.WriteHeader(http.StatusOK)
	_ = json.Unmarshal(body, &t)
	a := BuscarUsuario(t)
	if a.Nombre != "" && a.Nombre != "NoContra"{
		crearJson, _ := json.Marshal(a)
		var Usu Usuario
		_ = json.Unmarshal(crearJson, &Usu)
		_ = json.NewEncoder(w).Encode(Usu)
	}else if a.Nombre=="NoContra"{
		_ = json.NewEncoder(w).Encode("Contraseña Incorrecta")
	}else{
		_ = json.NewEncoder(w).Encode("Usuario No Encontrado")
	}
}

func elimusua(w http.ResponseWriter, r *http.Request){
	var t Usuario
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusOK)
	_ = json.Unmarshal(body, &t)
	a:= EliminarUsuario(t)
	var v Usuario
	_ = json.Unmarshal(a, &v)
	if v.Dpi == 0{
		_ = json.NewEncoder(w).Encode("Archivo Con Estructura Incorrecta")
	}else{
		_ = json.NewEncoder(w).Encode(v)
	}
}

func grafarbSC(w http.ResponseWriter, r *http.Request){
	if usuarios !=nil{
		usuarios.GraficarABSC()
		f, _ := os.Open("./ArbolSC.png")
		reader := bufio.NewReader(f)
		contenido, _ := ioutil.ReadAll(reader)
		encoded := base64.StdEncoding.EncodeToString(contenido)
		_, _ = fmt.Fprintf(w, encoded)
		f.Close()
	}else{
		_, _ = fmt.Fprintf(w, "No Exisen Usuarios Creados")
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusFound)
}

func grafarbCS(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	llave:=vars["Llave"]

	if usuarios !=nil{
		usuarios.GraficarABCS(llave)
		f, _ := os.Open("./ArbolCS.png")
		reader := bufio.NewReader(f)
		contenido, _ := ioutil.ReadAll(reader)
		encoded := base64.StdEncoding.EncodeToString(contenido)
		_, _ = fmt.Fprintf(w, encoded)
		f.Close()
	}else{
		_, _ = fmt.Fprintf(w, "No Exisen Usuarios Creados")
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusFound)
}

func grafarbC(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	llave:=vars["Llave"]

	if usuarios !=nil{
		usuarios.GraficarABC(llave)
		f, _ := os.Open("./ArbolC.png")
		reader := bufio.NewReader(f)
		contenido, _ := ioutil.ReadAll(reader)
		encoded := base64.StdEncoding.EncodeToString(contenido)
		_, _ = fmt.Fprintf(w, encoded)
		f.Close()
	}else{
		_, _ = fmt.Fprintf(w, "No Exisen Usuarios Creados")
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusFound)
}

//Grafo
func graf(w http.ResponseWriter, r *http.Request){
	var t Grafo
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusOK)
	_ = json.Unmarshal(body, &t)
	a:= definirGrafo(t)
	var v Grafo
	_ = json.Unmarshal(a, &v)
	if v.PosicionInicialRobot == ""{
		_ = json.NewEncoder(w).Encode("Archivo Con Estructura Incorrecta")
	}else{
		_ = json.NewEncoder(w).Encode(v)
	}
}

//Retornar Imagen Grafo
func retgraf(w http.ResponseWriter, r *http.Request){
	if dibujarGrafo() !=false{
		f, _ := os.Open("./Grafo.png")
		reader := bufio.NewReader(f)
		contenido, _ := ioutil.ReadAll(reader)
		encoded := base64.StdEncoding.EncodeToString(contenido)
		_, _ = fmt.Fprintf(w, encoded)
		f.Close()
	}else{
		_, _ = fmt.Fprintf(w, "No Existe Una Estructura De Grafo Cargada")
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusFound)
}

//FASE3
//---------------------------------------------------------------------------------------------


//---------------------------------------------------------------------------------------------
//FASE4

//Comentarios
func comen(w http.ResponseWriter, r *http.Request){
	var t coment
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusOK)
	_ = json.Unmarshal(body, &t)
	if vec != nil{
		a:= Comenta(t)
		var v coment
		_ = json.Unmarshal(a, &v)
		_ = json.NewEncoder(w).Encode(v)
	}else{
		_ = json.NewEncoder(w).Encode("No Hay Tiendas Cargadas")
	}
}

//Mostrar Comentarios
func mostrarcomen(w http.ResponseWriter, r *http.Request){
	var t coment
	body, err := ioutil.ReadAll(r.Body)
	if err != nil{
		_, _ = fmt.Fprintf(w, "Error al insertar")
	}
	w.Header().Set("Content-Type","applicattion/json")
	w.WriteHeader(http.StatusOK)
	_ = json.Unmarshal(body, &t)
	if vec != nil{
		a:= MostrarComenta(t)
		var v Comentarios
		_ = json.Unmarshal(a, &v)
		_ = json.NewEncoder(w).Encode(v)
	}else{
		_ = json.NewEncoder(w).Encode("No Hay Tiendas Cargadas")
	}
}

//FASE4
//---------------------------------------------------------------------------------------------

//ENDPOINTS

func Iniciar(){
	router := mux.NewRouter()
	//FASE1
	router.HandleFunc("/guardar", guardar).Methods("GET")
	router.HandleFunc("/getArreglo", getArreglo).Methods("GET")
	router.HandleFunc("/id/{numero}", tiendaN).Methods("GET")
	router.HandleFunc("/cargartienda", cargar).Methods("POST")
	router.HandleFunc("/TiendaEspecifica", tiendaE).Methods("POST")
	router.HandleFunc("/Eliminar", elim).Methods("DELETE")

	//FASE2
	router.HandleFunc("/guardarProductos", guardarProductos).Methods("GET")
	router.HandleFunc("/GrafoAños", GraficoArbolAños).Methods("GET")
	router.HandleFunc("/GrafoMesesAños/{año}", GraficoMesesArbolAños).Methods("GET")
	router.HandleFunc("/GrafoMatrizMesesAños/{Mes}", GraficoMatrizMesesArbolAños).Methods("GET")
	router.HandleFunc("/GrafoColaMatrizMesesAños/{dia}/{cat}", GraficoColaMatrizMesesArbolAños).Methods("GET")
	router.HandleFunc("/cargarInventario", inven).Methods("POST")
	router.HandleFunc("/cargarPedido", pedido).Methods("POST")
	router.HandleFunc("/cargarPedidoCarrito", pedidoCarrito).Methods("POST")
	router.HandleFunc("/graficarArbolProductos", GraficoArbolProductos).Methods("POST")

	//FASE3
	router.HandleFunc("/ObtenerUsuariosSC", grafarbSC).Methods("GET")
	router.HandleFunc("/ObtenerUsuariosCS/{Llave}", grafarbCS).Methods("GET")
	router.HandleFunc("/ObtenerUsuariosC/{Llave}", grafarbC).Methods("GET")
	router.HandleFunc("/ObtenerGrafo", retgraf).Methods("GET")
	router.HandleFunc("/CargarUsuarios", usua).Methods("POST")
	router.HandleFunc("/CrearUsuario", crearusua).Methods("POST")
	router.HandleFunc("/BuscarUsuario", buscarusua).Methods("POST")
	router.HandleFunc("/EliminarUsuario", elimusua).Methods("POST")
	router.HandleFunc("/CargarGrafo", graf).Methods("POST")

	//FASE4
	router.HandleFunc("/IngresarComentario", comen).Methods("POST")
	router.HandleFunc("/MostrarComentario", mostrarcomen).Methods("POST")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":3000",handler))
}