import React from 'react'
import CartaCarrito from "./CartaCarrito"

function MosaicoCarrito(props) {
    return (
        <div className="ui inverted segment mosaico container">
            <div className="ui four column link cards row">
            {props.productos.map((c, index) => (
                    <CartaCarrito nombreTienda={c.NombreTienda}
                        departamento={c.Departamento}
                        calificacion={c.Calificacion}
                        nombre={c.Nombre}
                        codigo={c.Codigo}
                        descripcion={c.Descripcion.substring(0,50)}
                        precio={c.Precio}
                        cantidad={c.Cantidad}
                        imagen={c.Imagen}
                        key={c.id}
                    />
                ))}
            </div>
        </div>
    )
}

export default MosaicoCarrito
