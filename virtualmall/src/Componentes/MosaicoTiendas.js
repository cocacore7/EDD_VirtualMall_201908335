import React from 'react'
import CartaTiendas from "./CartaTiendas"

function MosaicoTiendas(props) {
    return (
        <div className="ui inverted segment mosaico container">
            <div className="ui four column link cards row">
                {props.tiendas.map((c, index) => (
                    <CartaTiendas nombre={c.Nombre}
                        departamento={c.Departamento}
                        descripcion={c.Descripcion.substring(0,50)}
                        imagen={c.Logo}
                        contacto={c.Contacto}
                        calificacion={c.Calificacion}
                        key={c.id}
                    />
                ))}
            </div>
        </div>
    )
}

export default MosaicoTiendas
