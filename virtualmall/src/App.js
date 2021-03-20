import React from 'react'
import { BrowserRouter as Router,Route } from "react-router-dom";
import Reporte from "./Componentes/Reporte"
import ImportTiendas from "./Componentes/ImportTiendas"
import ImportProductos from "./Componentes/ImportProductos"
import CarritoCompras from "./Componentes/CarritoCompras"
import CargarArch from "./Componentes/CargarArch"
import NavBar from "./Componentes/NavBar"

function App() {
  return (
    <Router>
    <NavBar />
      <Route path="/reporte" component={Reporte} />
      <Route path="/listaT" component={ImportTiendas} />
      <Route path="/listaP" component={ImportProductos} />
      <Route path="/carrito" component={CarritoCompras} />
      <Route path="/cargar" component={CargarArch} />
    </Router>
  )
}

export default App
