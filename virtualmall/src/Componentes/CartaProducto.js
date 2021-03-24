import React from 'react'

function CartaProducto(props) {
    
    const enviar = () =>{
        var llave = localStorage.getItem('CPedido')
        if (llave === null || llave === undefined){
            llave=0
            localStorage.setItem('CPedido',0)
        }else{
            llave++
            localStorage.setItem('CPedido',llave)
        }
        const myproducto = {
            "NombreTienda": props.nombreTienda,
            "Departamento": props.departamento,
            "Calificacion": props.calificacion,
            "Nombre": props.nombre,
            "Codigo": props.codigo,
            "Descripcion": props.descripcion,
            "Precio": props.precio,
            "Cantidad": props.cantidad,
            "Imagen": props.imagen,
            "id":llave
        }
        var datos = localStorage.getItem('Carrito')
        if (datos === null || datos === undefined){
            localStorage.setItem('Carrito',JSON.stringify([myproducto]))
        }else{
            datos = JSON.parse(datos)
            datos.push(myproducto)
            console.log(datos)
            localStorage.setItem('Carrito',JSON.stringify(datos))
        }
        alert("Producto " + props.nombre + "Añadido")
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
                    </div>
                    <div className="description">Descipcion: {props.descripcion}</div>
                    <div className="ui segment green button center fluid" onClick={enviar}>Comprar</div>
                </div>
                <div className="extra content">
                    <span className="right floated">Cantidad: {props.cantidad}</span>
                    <span><i className="dollar sign icon" />{props.precio}</span>
                </div>
            </div>
        </div>
    )
}

export default CartaProducto
