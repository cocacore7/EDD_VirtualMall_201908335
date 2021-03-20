import {React,useState} from 'react'
import { Menu } from "semantic-ui-react";
import { Link } from "react-router-dom";
import "../css/Nav.css"

const colores=["red","green","blue","grey"]
const opciones=["Cargar Archivos","Reportes","Lista De Tiendas","Carrito De Compras"]
const url=["/cargar","/reporte","/listaT","/carrito","/"]
function NavBar() {
    const [activo, setactivo] = useState(colores[0])
    return (
        <Menu inverted className="Nav">
            {colores.map((c,index)=>(
                <Menu.Item as={Link} to={url[index]}
                    key={c} 
                    name={opciones[index]} 
                    active={activo===c} 
                    color={c} 
                    onClick={()=>setactivo(c)} 
                />
            ))}
        </Menu>
    )
}

export default NavBar
