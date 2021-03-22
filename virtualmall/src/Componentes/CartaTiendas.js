import React,{useState} from 'react'
import '../css/Carta.css'
import { useHistory } from 'react-router-dom';


function CartaTiendas(props) {
    let history = useHistory()

    const enviar = () =>{
        const mytienda = {
            "Tienda": props.nombre,
            "Departamento": props.departamento,
            "Calificacion": props.calificacion
        }
        localStorage.setItem('CartaTienda',JSON.stringify(mytienda))
        history.push("/listaP");
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
                    <div className="ui segment green button center fluid" onClick={enviar}>Ver Productos</div>
                </div>
                <div className="extra content">
                    <a>Calificacion: {props.calificacion}</a>
                </div>
            </div>
        </div>
    )
}

export default CartaTiendas
