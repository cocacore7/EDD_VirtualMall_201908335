import React,{useEffect,useState} from 'react'
import MosaicoCarrito from "./MosaicoCarrito"
import "../css/ImportTiendas.css"
import { useHistory } from 'react-router-dom';

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
        localStorage.clear('Carrito')
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
