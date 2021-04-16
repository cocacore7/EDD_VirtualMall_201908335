import React from 'react'
import { BrowserRouter as Router,Redirect,Route } from "react-router-dom";
import Reporte from "./Componentes/Reporte"
import ImportTiendas from "./Componentes/ImportTiendas"
import ImportProductos from "./Componentes/ImportProductos"
import CarritoCompras from "./Componentes/CarritoCompras"
import CargarArch from "./Componentes/CargarArch"
import Iniciosesion from "./Componentes/iniciosesion"
import Crearusuario from "./Componentes/CreateUser"


function App() {
  return (
    <Router>
      <Route exact path="/">
        <Redirect to="/iniciosesion"/>
      </Route>
      <Route path="/reporte" component={Reporte} />
      <Route path="/listaT" component={ImportTiendas} />
      <Route path="/listaP" component={ImportProductos} />
      <Route path="/carrito" component={CarritoCompras} />
      <Route path="/cargar" component={CargarArch} />
      <Route path="/iniciosesion" component={Iniciosesion} />
      <Route path="/crearusuario" component={Crearusuario} />
    </Router>
  )
}

export default App
