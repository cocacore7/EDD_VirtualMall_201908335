import React,{useEffect,useState} from 'react'

import { useHistory } from 'react-router-dom';
import axios from "axios"
import "../css/ImportTiendas.css"
import NavBar from "../Componentes/NavBar"

function ImportProductosR() {
    const [imagenProducto,setImagenP]=useState('')
    const tienda = []
    let history = useHistory()


    function handleClick() {
        history.push("/listaTR");
    }
    

    useEffect(() => {
        let CarTienda = localStorage.getItem('CartaTienda')
        if (CarTienda!=null){
            tienda.push(JSON.parse(CarTienda))
        }
        const peticionImagen = {
            "Departamento": tienda[0].Departamento,
            "Nombre": tienda[0].Tienda,
            "Calificacion": tienda[0].Calificacion
        }
        async function obtener() {
            const data2 = await axios.post("http://localhost:3000/graficarArbolProductos",peticionImagen);
            setImagenP("data:image/png;base64,"+data2.data)
        }
        obtener()
    },[]);

    return (
        <>
        <NavBar 
            colores={["red","green","yellow","blue","grey"]}
            opciones={["Cargar Archivos","Reportes","Tiendas","Cerrar Sesion", "Eliminar Mi Cuenta"]}
            url={["/cargar","/reporte","/listaTR","/iniciosesion","/iniciosesion"]}
            activo={"yellow"}
        />
        <div>
            <div className="ImportTiendas">
                <br></br>
                <div className="ui segment red button center container" onClick={()=>{handleClick()}}>Regresar</div>
                <br></br>
                <div className="ui centered container">
                    <img src={imagenProducto} class="ui fluid centered rounded image "/>
                </div>
                <br></br>
            </div>
        </div>
        </>
    )
}

export default ImportProductosR
