import React,{useState} from 'react'
import "../css/ImportTiendas.css"
import axios from "axios"

function CargarArch() {
    const [tiendas, setTienda]=useState(null)
    const [pedidos, setPedido]=useState(null)
    const [inventario, setInventario]=useState(null)

    const cargartiendas = async()=>{
        if (tiendas ===null){
            console.log("No Hay Archivos Cargados")
        }else{
            const data = await axios.post("http://localhost:3000/cargartienda",tiendas[0])
            console.log(data.data)
        }
        
    }

    const cargarPedidos = async()=>{
        if (pedidos ===null){
            console.log("No Hay Archivos Cargados")
        }else{
            const data = await axios.post("http://localhost:3000/cargarPedido",pedidos[0])
            console.log(data.data)
        }
    }

    const cargarinventarios = async()=>{
        if (inventario ===null){
            console.log("No Hay Archivos Cargados")
        }else{
            const data = await axios.post("http://localhost:3000/cargarInventario",inventario[0])
            console.log(data.data)
        }
    }
    
    return (
        <div className="ImportTiendas">
            <br/>
            <div className="ui inverted segment container">
                <h1 style={{color: '#00FFFF'}}>Subir Tinedas </h1>
                <div className="ui inverted input">
                    <input type="file" accept=".json" onChange={(e)=>setTienda(e.target.files)}/>
                </div><br/><br/>
                <div className="ui segment green button center" onClick={()=>cargartiendas()}>Cargar Tiendas</div>
                <div className="ui inverted divider" />
                    <h1 style={{color: '#00FFFF'}}>Subir Inventario </h1>
                <div className="ui inverted input">
                    <input type="file" accept=".json" onChange={(e)=>setInventario(e.target.files)}/>
                </div><br/><br/>
                <div className="ui segment green button center" onClick={()=>cargarinventarios()}>Cargar Inventario</div>
                <div className="ui inverted divider" />
                <h1 style={{color: '#00FFFF'}}>Subir Pedidos </h1>
                <div className="ui inverted input">
                    <input type="file" accept=".json" onChange={(e)=>setPedido(e.target.files)}/>
                </div><br/><br/>
                <div className="ui segment green button center" onClick={()=>cargarPedidos()}>Cargar Pedidos</div>
            </div>
            <br/>
        </div>
    )
}

export default CargarArch
