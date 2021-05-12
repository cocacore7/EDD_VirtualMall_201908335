import React from 'react'
import { useHistory } from 'react-router-dom';
import { Rating } from 'semantic-ui-react'
import '../css/Carta.css'

function CartaTiendaR(props) {
    let history = useHistory()

    
    const enviar = () =>{
        const mytienda = {
            "Tienda": props.nombre,
            "Departamento": props.departamento,
            "Calificacion": props.calificacion
        }
        localStorage.setItem('CartaTienda',JSON.stringify(mytienda))
        history.push("/listaPR");
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
                        <p>Departamento: {props.departamento}</p>
                        <p>Contacto: {props.contacto}</p>
                    </div>
                    <div className="description">Descripcion: {props.descripcion}</div>
                    <div className="ui segment green button center fluid" onClick={enviar}>Ver Productos</div>
                </div>
                <div className="extra content">
                    <Rating defaultRating={props.calificacion} maxRating={5} disabled />
                </div>
            </div>
        </div>
    )
}

export default CartaTiendaR
