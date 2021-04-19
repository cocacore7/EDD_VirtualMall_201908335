import React,{useState} from 'react'
import NavBar from "../Componentes/NavBar"
import "../css/CreateUser.css"
import axios from "axios"
import { useHistory } from 'react-router-dom';
import sha256 from 'crypto-js/sha256';

function InicioSesion() {
    const [dpi, setdpi]=useState("")
    const [password, setpassword]=useState("")
    let history = useHistory()


    const enviar = () =>{
        var usuario = {
            "Dpi": parseInt(dpi),
            "Nombre": "",
            "Correo": "",
            "Password": sha256(password).toString(),
            "Cuenta": ""
        }
        async function obtener() {
            const data = await axios.post("http://localhost:3000/BuscarUsuario",usuario);
            alert( JSON.stringify(data.data));
            if (data.data.Cuenta === "Admin"){
                history.push("/cargar")
            }else if(data.data.Cuenta === "Usuario"){
                history.push("/listaT")
            }
        }
        obtener()
    }
    
    return (
        <>
        <NavBar 
            colores={["red","green"]}
            opciones={["Crear Usuario","Iniciar Sesion"]}
            url={["/crearusuario","/iniciosesion","/"]}
            activo={"green"}
        />
        <div className="UserList">
            <br></br>
            <div className="ui inverted segment container formulario form">
                <div className="field">
                    <label>DPI</label>
                    <input type="text" name="dpi" placeholder="123456789" onChange={e=>setdpi(e.target.value)} />
                </div>
                <div className="field">
                    <label>Contraseña</label>
                    <input type="password" name="Contraseña" placeholder="a1a1a1a1****" onChange={e=>setpassword(e.target.value)} />
                </div>
                <button className="ui green button" onClick={enviar}>Enviar</button>
            </div>
        </div>
        </>
    )
}

export default InicioSesion
