import React,{useEffect,useState} from 'react'
import MosaicoTiendas from "./MosaicoTiendasR"
import "../css/ImportTiendas.css"
import axios from "axios"
import NavBar from "../Componentes/NavBar"

function ImportTiendasR() {
    const nuevo = []
    const [tiendas, settiendas] = useState([])
    
    useEffect(() => {
        async function obtener() {
            const data = await axios.get("http://localhost:3000/guardar");
            if (typeof(data.data) != "string"){
                let contador = 1
                for (let index = 0; index < data.data.Datos.length; index++) {
                    for (let index2 = 0; index2 < data.data.Datos[index].Departamentos.length; index2++) {
                        for (let index3 = 0; index3 < data.data.Datos[index].Departamentos[index2].Tiendas.length; index3++) {
                            const mytienda = {
                                "Nombre": data.data.Datos[index].Departamentos[index2].Tiendas[index3].Nombre,
                                "Departamento": data.data.Datos[index].Departamentos[index2].Nombre,
                                "Descripcion": data.data.Datos[index].Departamentos[index2].Tiendas[index3].Descripcion,
                                "Contacto": data.data.Datos[index].Departamentos[index2].Tiendas[index3].Contacto,
                                "Calificacion": data.data.Datos[index].Departamentos[index2].Tiendas[index3].Calificacion,
                                "Logo": data.data.Datos[index].Departamentos[index2].Tiendas[index3].Logo,
                                "id": contador
                            }
                            contador++
                            nuevo.push(mytienda)
                            }
                    }
                }
                settiendas(nuevo)
                console.log(nuevo)
            }else{
                alert("No Existen Tiendas Cargadas")
            }
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
                <MosaicoTiendas tiendas={tiendas} />
                <br></br>
            </div>
        </div>
        </>
    )
}

export default ImportTiendasR
