import React,{useEffect,useState} from 'react'
import MosaicoTiendas from "./MosaicoTiendas"
import "../css/ImportTiendas.css"
const axios = require('axios')

function ImportTiendas() {
    const [nuevo, setnuevo] = useState([])
    const [tiendas, settiendas] = useState([])
    
    useEffect(() => {
        async function obtener() {
            const data = await axios.get("http://localhost:3000/guardar");
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
        }
        obtener()
    },[]);
    return (
        <div>
            <div className="ImportTiendas">
                <br></br>
                <MosaicoTiendas tiendas={tiendas} />
                <br></br>
            </div>
        </div>
    )
}

export default ImportTiendas