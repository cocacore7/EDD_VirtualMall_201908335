import React from 'react'
import CartaProductos from "./CartaProducto"

function MosaicoProductos(props) {
    return (
        <div className="ui inverted segment mosaico container">
            <div className="ui four column link cards row">
                {props.productos.map((c, index) => (
                    <CartaProductos nombre={c.name}
                        categoria={c.categories[0].name}
                        descripcion={c.description.substring(0,50)}
                        imagen={c.image}
                        precio={c.price}
                        id={c.id}
                        key={c.id}
                    />
                ))}
            </div>
        </div>
    )
}

export default MosaicoProductos
