import React,{useEffect,useState} from 'react'
import { useHistory } from 'react-router-dom';
import axios from "axios"
import MosaicoCarrito from "./MosaicoCarrito"
import "../css/ImportTiendas.css"


function CarritoCompras() {
    const [productos2, setproductos2] = useState([])
    let history = useHistory()


    useEffect(() => {
        let Carr = localStorage.getItem('Carrito')
        if (Carr!=null){
            setproductos2(JSON.parse(Carr))
        }
        console.log(JSON.parse(Carr))
    },[]);

    
    const enviar = () =>{
        const regreso = []
        for (let index = 0; index < productos2.length; index++) {
            var fecha = new Date()
            const pedido = {
                "Fecha": fecha.getDate() + '-' + (fecha.getMonth() + 1) + '-' + fecha.getFullYear(),
                "Tienda": productos2[index].NombreTienda,
                "Departamento": productos2[index].Departamento,
                "Calificacion": productos2[index].Calificacion,
                "Productos": [
                {
                    "Codigo": productos2[index].Codigo
                }]
            }
            regreso.push(pedido)
        }
        const Solicitud = {"Pedidos":regreso}
        console.log(Solicitud)
        const cargartiendas = async()=>{
            const data = await axios.post("http://localhost:3000/cargarPedidoCarrito",Solicitud)

            if (typeof data != "string"){
                localStorage.clear('Carrito')
                localStorage.clear('CPedido')
                alert(JSON.stringify(data.data))
            }else{
                alert(data.data)
            }
            
        }
        cargartiendas()
        history.go("/carrito");   
    }
    
    return (
        <div>
            <div className="ImportTiendas">
                <br></br>
                <MosaicoCarrito productos={productos2} />
                <br></br>
                <div className="ui segment blue button center container" onClick={enviar}>Ingresar Pedido</div>
                <br></br>
            </div>
        </div>
    )
}

export default CarritoCompras
