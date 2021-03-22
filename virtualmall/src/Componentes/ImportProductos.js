import React,{useEffect,useState} from 'react'
import MosaicoProductos from "./MosaicoProductos"
import { useHistory } from 'react-router-dom';
import "../css/ImportTiendas.css"
const axios = require('axios')

function ImportProductos() {
    const [nuevo2, setnuevo2] = useState([])
    const [productos, setproductos] = useState([])
    const [tienda,settienda] = useState([])
    let history = useHistory()

    function handleClick() {
        history.push("/listaT");
    }

    useEffect(() => {
        let CarTienda = localStorage.getItem('CartaTienda')
        if (CarTienda!=null){
            tienda.push(JSON.parse(CarTienda))
        }
        async function obtener() {
            const data = await axios.get("http://localhost:3000/guardarProductos");
            let contador = 1
            for (let index = 0; index < data.data.Inventarios.length; index++) {
                for (let index2 = 0; index2 < data.data.Inventarios[index].Productos.length; index2++) {
                    if (tienda[0].Tienda == data.data.Inventarios[index].Tienda && tienda[0].Departamento == data.data.Inventarios[index].Departamento && tienda[0].Calificacion == data.data.Inventarios[index].Calificacion && data.data.Inventarios[index].Productos[index2].Cantidad!=0){
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
            }
            setproductos(nuevo2)
            console.log(nuevo2)
        }
        obtener()
        
    },[]);

    return (
        <div>
            <div className="ImportTiendas">
                <br></br>
                <div className="ui segment red button center container" onClick={()=>{handleClick()}}>Regresar</div>
                <br></br>
                <MosaicoProductos productos={productos} />
                <br></br>
            </div>
        </div>
    )
}

export default ImportProductos
