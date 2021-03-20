import React from 'react'
import CartaTiendas from "./CartaTiendas"

function MosaicoTiendas(props) {
    return (
        <div className="ui inverted segment mosaico container">
            <div className="ui four column link cards row">
                {props.tiendas.map((c, index) => (
                    <CartaTiendas nombre={c.name}
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

export default MosaicoTiendas
