import React,{useEffect,useState} from 'react'

import { useHistory } from 'react-router-dom';
import axios from "axios"
import "../css/ImportTiendas.css"
import MosaicoProductos from "./MosaicoProductos"
import NavBar from "../Componentes/NavBar"


function ImportProductos() {
    const [imagenProducto,setImagenP]=useState('')
    const nuevo2 = []
    const [productos, setproductos] = useState([])
    const tienda = []
    let history = useHistory()


    function handleClick() {
        history.push("/listaT");
    }
    

    useEffect(() => {
        let CarTienda = localStorage.getItem('CartaTienda')
        if (CarTienda!=null){
            tienda.push(JSON.parse(CarTienda))
        }
        console.log(tienda)
        async function obtener() {
            const data = await axios.get("http://localhost:3000/guardarProductos");
            console.log(data.data)
            if (typeof(data.data) != "string"){
                    let contador = 1
                    for (let index = 0; index < data.data.Inventarios.length; index++) {
                        if (tienda[0].Tienda === data.data.Inventarios[index].Tienda && tienda[0].Departamento === data.data.Inventarios[index].Departamento && tienda[0].Calificacion === data.data.Inventarios[index].Calificacion){
                            if (data.data.Inventarios[0].Productos != null){
                                for (let index2 = 0; index2 < data.data.Inventarios[index].Productos.length; index2++) {
                                    if(data.data.Inventarios[index].Productos[index2].Cantidad > 0){
                                        const mytienda = {
                                            "NombreTienda": data.data.Inventarios[index].Tienda,
                                            "Departamento": data.data.Inventarios[index].Departamento,
                                            "Calificacion": data.data.Inventarios[index].Calificacion,
                                            "id": contador,
                                            "Nombre": data.data.Inventarios[index].Productos[index2].Nombre,
                                            "Codigo": data.data.Inventarios[index].Productos[index2].Codigo,
                                            "Descripcion": data.data.Inventarios[index].Productos[index2].Descripcion,
                                            "Precio": data.data.Inventarios[index].Productos[index2].Precio,
                                            "Cantidad": data.data.Inventarios[index].Productos[index2].Cantidad,
                                            "Imagen": data.data.Inventarios[index].Productos[index2].Imagen
                                        }
                                        contador++
                                        nuevo2.push(mytienda)
                                    }
                                }
                            }else{
                                alert("No Existen Productos Cargados")
                            }
                        }
                    }
                    setproductos(nuevo2)
                    console.log(nuevo2)
            }else{
                alert("No Existen Productos Cargados")
            }
        }
        obtener()

    },[]);

    return (
        <>
        <NavBar 
            colores={["red","green","blue","grey"]}
            opciones={["Tiendas","Carrito De Compras","Cerrar Sesion", "Eliminar Mi Cuenta"]}
            url={["/listaT","/carrito","/iniciosesion","/iniciosesion"]}
            activo={"red"}
        />
        <div>
            <div className="ImportTiendas">
                <br></br>
                <div className="ui segment red button center container" onClick={()=>{handleClick()}}>Regresar</div>
                <br></br>
                <MosaicoProductos productos={productos} />
                <br></br>
            </div>
        </div>
        </>
    )
}

export default ImportProductos
