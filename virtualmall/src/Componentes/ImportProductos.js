import React,{useEffect,useState} from 'react'
import MosaicoProductos from "./MosaicoProductos"
import { useHistory } from 'react-router-dom';
import "../css/ImportTiendas.css"
const axios = require('axios')

function ImportProductos() {
    const [productos, setproductos] = useState([])
    let history = useHistory()
    function handleClick() {
        history.push("/listaT");
    }
    useEffect(() => {
        async function obtener() {
            if (productos.length === 0) {
                const data = await axios.get('https://gorest.co.in/public-api/products');
                console.log(data.data.data)
                setproductos(data.data.data)
            }
        }
        obtener()
    });
    return (
        <div>
            <div className="ImportTiendas">
                <br></br>
                <div className="ui segment red button center container" onClick={handleClick}>Regresar</div>
                <br></br>
                <MosaicoProductos productos={productos} />
            </div>
        </div>
    )
}

export default ImportProductos
