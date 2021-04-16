import React,{useState} from 'react'
import NavBar from "./NavBar"
import "../css/CreateUser.css"
import axios from "axios"
import { useHistory } from 'react-router-dom';

function CreateUser() {
    const [dpi, setdpi]=useState('')
    const [Nombre, setNombre] = useState("")
    const [Correo, setCorreo] = useState("")
    const [Password, setPassword] = useState("")
    const [Usuario, setUsuario] = useState("")
    let history = useHistory()

    const enviar = () =>{
        var usuario = {
            "Dpi": parseInt(dpi),
            "Nombre": Nombre,
            "Correo": Correo,
            "Password": Password,
            "Cuenta": Usuario
        }
        async function obtener() {
            const data = await axios.post("http://localhost:3000/CrearUsuario",usuario);
            alert( JSON.stringify(data.data));
            history.push("/iniciosesion")
        }
        obtener()
    }

    return (
        <>
        <NavBar 
            colores={["red","green"]}
            opciones={["Crear Usuario","Iniciar Sesion"]}
            url={["/crearusuario","/iniciosesion","/"]}
            activo={"red"}
        />
        <div className="UserList">
            <br></br>
            <div className="ui inverted segment container formulario form">
                <div className="field">
                    <label>DPI</label>
                    <input type="text" name="dpi" placeholder="123456789" onChange={e=>setdpi(e.target.value)} />
                </div>
                <div className="field">
                    <label>Nombre</label>
                    <input type="text" name="nombre" placeholder="Ejemplo777" onChange={e=>setNombre(e.target.value)} />
                </div>
                <div className="field">
                    <label>Correo</label>
                    <input type="text" name="correo" placeholder="Ejemplo@gmail.com" onChange={e=>setCorreo(e.target.value)} />
                </div>
                <div className="field">
                    <label>Contraseña</label>
                    <input type="text" name="contraseña" placeholder="a1a1a1a1****" onChange={e=>setPassword(e.target.value)} />
                </div>
                <div className="field">
                    <label>Usuario</label>
                    <input type="text" name="usuario" placeholder="Admin / Usuario" onChange={e=>setUsuario(e.target.value)} />
                </div>
                <button className="ui green button" onClick={enviar}>Enviar</button>
            </div>
        </div>
        </>
    )
}

export default CreateUser
