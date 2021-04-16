import {React,useState} from 'react'
import { Menu } from "semantic-ui-react";
import { Link } from "react-router-dom";
import "../css/Nav.css"

function NavBar(props) {
    const colores=props.colores
    const opciones=props.opciones
    const url=props.url
    const [activo, setactivo] = useState(props.activo)
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
