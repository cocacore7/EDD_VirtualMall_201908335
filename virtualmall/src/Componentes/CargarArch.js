import React,{useState} from 'react'
import axios from "axios"
import "../css/ImportTiendas.css"
import NavBar from "../Componentes/NavBar"


function CargarArch() {
    const [usuarios, setUsuarios]=useState(null)
    const [grafo, setGrafo]=useState(null)
    const [tiendas, setTienda]=useState(null)
    const [pedidos, setPedido]=useState(null)
    const [inventario, setInventario]=useState(null)

    
    const cargartiendas = async()=>{
        if (tiendas ===null){
            alert("No Hay Archivos Cargados")
        }else{
            await axios.post("http://localhost:3000/cargartienda",tiendas[0])
        }
    }

    const cargarPedidos = async()=>{
        if (pedidos ===null){
            alert("No Hay Archivos Cargados")
        }else{
            await axios.post("http://localhost:3000/cargarPedido",pedidos[0])
        }
    }

    const cargarinventarios = async()=>{
        if (inventario ===null){
            alert("No Hay Archivos Cargados")
        }else{
            await axios.post("http://localhost:3000/cargarInventario",inventario[0])
        }
    }

    const cargarUsuarios = async()=>{
        if (usuarios ===null){
            alert("No Hay Archivos Cargados")
        }else{
            await axios.post("http://localhost:3000/CargarUsuarios",usuarios[0])
        }
    }

    const cargarGrafo = async()=>{
        if (grafo ===null){
            alert("No Hay Archivos Cargados")
        }else{
            await axios.post("http://localhost:3000/CargarGrafo",grafo[0])
        }
    }
    
    return (
        <>
        <NavBar 
            colores={["red","green","yellow","blue","grey"]}
            opciones={["Cargar Archivos","Reportes","Tiendas","Cerrar Sesion", "Eliminar Mi Cuenta"]}
            url={["/cargar","/reporte","/listaTR","/iniciosesion","/iniciosesion"]}
            activo={"red"}
        />
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
                <div className="ui inverted divider" />

                <h1 style={{color: '#00FFFF'}}>Subir Usuarios </h1>
                <div className="ui inverted input">
                    <input type="file" accept=".json" onChange={(e)=>setUsuarios(e.target.files)}/>
                </div><br/><br/>
                <div className="ui segment green button center" onClick={()=>cargarUsuarios()}>Cargar Pedidos</div>
                <div className="ui inverted divider" />

                <h1 style={{color: '#00FFFF'}}>Subir Grafo </h1>
                <div className="ui inverted input">
                    <input type="file" accept=".json" onChange={(e)=>setGrafo(e.target.files)}/>
                </div><br/><br/>
                <div className="ui segment green button center" onClick={()=>cargarGrafo()}>Cargar Pedidos</div>
            </div>
            <br/>
        </div>
        </>
    )
}

export default CargarArch
