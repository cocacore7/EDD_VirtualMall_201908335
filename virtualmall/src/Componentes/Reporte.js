import React from 'react'
import "../css/ImportTiendas.css"

function UserList() {
    return (
        <div className="ImportTiendas">
            <div className="ui inverted segment container items">
                <div className="item">
                    <div className="ui big segment rounded image">
                        <img src="" width={600} height={600}/>
                    </div>
                    <div className="content">
                    <h1 style={{color: '#00FFFF'}}>Años</h1>
                        <div className="meta">
                            <input type="text" placeholder="Ingresar Año"></input>
                        </div>
                        <div className="extra">
                            <div className="ui segment green button center" onClick={()=>{}}>Cargar Grafica Del Mes Solicitado</div>
                        </div>
                    </div>
                </div>
                <div className="ui inverted divider" />
                <div className="item">
                    <div className="ui big segment rounded image">
                        <img src="" width={600} height={600}/>
                    </div>
                    <div className="content">
                    <h1 style={{color: '#00FFFF'}}>Mes</h1>
                        <div className="meta">
                            <input type="text" placeholder="Ingresar Mes"></input>
                        </div>
                        <div className="extra">
                            <div className="ui segment green button center" onClick={()=>{}}>Cargar Grafica De La Matriz Solicitada</div>
                        </div>
                    </div>
                </div>
                <div className="ui inverted divider" />
                <div className="item">
                    <div className="ui big segment rounded image">
                        <img src="" width={600} height={600}/>
                    </div>
                    <div className="content">
                    <h1 style={{color: '#00FFFF'}}>Matriz De Pedidos</h1>
                        <div className="meta">
                            <input type="text" placeholder="Ingresar Dia"></input><br/><br/>
                            <input type="text" placeholder="Ingresar Categoria"></input>
                        </div>
                        <div className="extra">
                            <div className="ui segment green button center" onClick={()=>{}}>Cargar Grafica De La Cola Solicitada</div>
                        </div>
                    </div>
                </div>
                <div className="ui inverted divider" />
                <div className="item">
                    <div className="ui big segment rounded image">
                        <img src="" width={600} height={600}/>
                    </div>
                    <div className="content">
                    <h1 style={{color: '#00FFFF'}}>Cola De Pedidos</h1>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default UserList
