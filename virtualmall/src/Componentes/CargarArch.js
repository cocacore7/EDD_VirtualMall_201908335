import {React,usestate} from 'react'
import "../css/ImportTiendas.css"
const axios = require('axios')

function CargarArch() {
    return (
        <div className="ImportTiendas">
            <div className="ui inverted segment container">
                <h1 style={{color: '#00FFFF'}}>Subir Tinedas </h1>
                <div className="ui inverted input">
                    <input type="file" accept=".json"/>
                </div><br/><br/>
                <div className="ui segment green button center" onClick={()=>{}}>Cargar Tiendas</div>
                <div className="ui inverted divider" />
                    <h1 style={{color: '#00FFFF'}}>Subir Inventario </h1>
                <div className="ui inverted input">
                    <input type="file" accept=".json"/>
                </div><br/><br/>
                <div className="ui segment green button center" onClick={()=>{}}>Cargar Inventario</div>
                <div className="ui inverted divider" />
                <h1 style={{color: '#00FFFF'}}>Subir Pedidos </h1>
                <div className="ui inverted input">
                    <input type="file" accept=".json"/>
                </div><br/><br/>
                <div className="ui segment green button center" onClick={()=>{}}>Cargar Pedidos</div>
            </div>
        </div>
    )
}

export default CargarArch
