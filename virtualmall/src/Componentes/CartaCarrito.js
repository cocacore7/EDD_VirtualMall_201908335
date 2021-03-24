import React,{useEffect,useState} from 'react'
import { useHistory } from 'react-router-dom';

function CartaCarrito(props) {
    const [productos3, setproductos3] = useState([])
    const productos4 = []
    let history = useHistory()
    

    useEffect(() => {
        let data = localStorage.getItem('Carrito')
        if (data!=null){
            setproductos3(JSON.parse(data))
        }
    },[]);

    const enviar = () =>{
        let contador=0
        for (let index = 0; index < productos3.length; index++) {
            if (productos3[index].NombreTienda===props.nombreTienda && productos3[index].Departamento===props.departamento && productos3[index].Calificacion===props.calificacion && productos3[index].Nombre===props.nombre && productos3[index].Codigo===props.codigo){
                if (contador !== 0){
                    productos4.push(productos3[index])
                }
                contador++
            }else{
                productos4.push(productos3[index])
            }
        }
        setproductos3(productos4)
        localStorage.clear("Carrito")
        localStorage.setItem('Carrito',JSON.stringify(productos4))
        history.go("/carrito");
        alert("Producto Eliminado")
    }
    
    return (
        <div className="column carta">
            <div className="ui card">
                <div className="image">
                    <img src={props.imagen} />
                </div>
                <div className="content">
                    <div className="header">{props.nombre}</div>
                    <div className="meta">
                        <p>Codigo: {props.codigo}</p>
                        <p>Departamento: {props.departamento}</p>
                        <p>Tienda: {props.nombreTienda}</p>
                    </div>
                    <div className="description">Descripcion: {props.descripcion}</div>
                    <div className="ui segment red button center container" onClick={enviar}>Eliminar</div>
                </div>
                <div className="extra content">
                    <span><i className="dollar sign icon" />{props.precio}</span>
                </div>
            </div>
        </div>
    )
}

export default CartaCarrito
