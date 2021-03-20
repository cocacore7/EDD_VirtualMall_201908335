import React from 'react'
import '../css/Carta.css'
import { useHistory } from 'react-router-dom';


function CartaTiendas(props) {
    let history = useHistory()
    function handleClick() {
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
                        <a>{props.categoria}</a>
                    </div>
                    <div className="description">{props.descripcion}</div>
                    <div className="ui segment green button center fluid" onClick={handleClick}>Ver Productos</div>
                </div>
                <div className="extra content">
                    <span className="right floated">Joined in 2021</span>
                    <span><i className="dollar sign icon" />{props.precio}</span>
                </div>
            </div>
        </div>
    )
}

export default CartaTiendas
