import React from 'react'
import CartaProductos from "./CartaProducto"

function MosaicoProductos(props) {
    return (
        <div className="ui inverted segment mosaico container">
            <div className="ui four column link cards row">
                {props.productos.map((c, index) => (
                    <CartaProductos nombreTienda={c.NombreTienda}
                        departamento={c.Departamento}
                        calificacion={c.Calificacion}
                        nombre={c.Nombre}
                        codigo={c.Codigo}
                        descripcion={c.Descripcion.substring(0,50)}
                        precio={c.Precio}
                        cantidad={c.Cantidad}
                        imagen={c.Imagen}
                        id={c.id}
                        key={c.id}
                    />
                ))}
            </div>
        </div>
    )
}

export default MosaicoProductos
