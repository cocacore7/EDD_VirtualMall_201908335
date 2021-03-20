import React,{useEffect,useState} from 'react'
import MosaicoTiendas from "./MosaicoTiendas"
import "../css/ImportTiendas.css"
const axios = require('axios')

function ImportTiendas() {
    const [tiendas, settiendas] = useState([])
    useEffect(() => {
        async function obtener() {
            if (tiendas.length === 0) {
                const data = await axios.get('https://gorest.co.in/public-api/products');
                console.log(data.data.data)
                settiendas(data.data.data)
            }
        }
        obtener()
    });
    return (
        <div>
            <div className="ImportTiendas">
                <br></br>
                <MosaicoTiendas tiendas={tiendas} />
            </div>
        </div>
    )
}

export default ImportTiendas
