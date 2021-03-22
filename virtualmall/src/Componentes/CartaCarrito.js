import React,{useEffect,useState} from 'react'
import { useHistory } from 'react-router-dom';

function CartaCarrito(props) {
    const [productos3, setproductos3] = useState([])
    const [productos4, setproductos4] = useState([])
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
                        <a>Departamento: {props.departamento}</a><br/>
                        <a>Contacto: {props.contacto}</a>
                    </div>
                    <div className="description">Descripcion: {props.descripcion}</div>
                    <div className="ui segment red button center container" onClick={enviar}>Eliminar</div>
                </div>
                <div className="extra content">
                    <a>Calificacion: {props.calificacion}</a>
                </div>
            </div>
        </div>
    )
}

export default CartaCarrito
