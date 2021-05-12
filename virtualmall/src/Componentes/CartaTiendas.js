import React,{useEffect,useState} from 'react'
import { useHistory } from 'react-router-dom';
import { Rating } from 'semantic-ui-react'
import '../css/Carta.css'
import { Button, Header, Image, Modal,Comment, Form } from 'semantic-ui-react'
import Comentario from './Comentario'
import axios from "axios"


function CartaTiendas(props) {
    let history = useHistory()
    const [open, setOpen] = React.useState(false)
    const [mensaje, setMensaje] = useState("")
    const [comentarios, setComentarios] = useState([])
    const nuevo2 = []

    
    const enviar = () =>{
        const mytienda = {
            "Tienda": props.nombre,
            "Departamento": props.departamento,
            "Calificacion": props.calificacion
        }
        localStorage.setItem('CartaTienda',JSON.stringify(mytienda))
        history.push("/listaP");
    }

    const enviarComentario = async()=> {
        let Usuario = localStorage.getItem('Usuario')
        Usuario = JSON.parse(Usuario)
        let comentario = {
            Tipo: "Tienda",
            Tienda: props.nombre,
            Departamento: props.departamento,
            Calificacion: parseInt(props.calificacion),
            Codigo: parseInt(0),
            Dpi: parseInt(Usuario.Dpi) ,
            Comentario: mensaje
        }
        await axios.post("http://localhost:3000/IngresarComentario",comentario)
    }

    const cargarComentarios = async()=> {
        const comentariosprueba = []
        let comentario = {
            Tipo: "Tienda",
            Tienda: props.nombre,
            Departamento: props.departamento,
            Calificacion: parseInt(props.calificacion),
            Codigo: parseInt(0),
            Dpi: parseInt(0),
            Comentario: ""
        }
        const resultado = await axios.post("http://localhost:3000/MostrarComentario",comentario)
        console.log(resultado.data)
        if(resultado.data == undefined || resultado.data == null ){
            setComentarios(comentariosprueba)
        }else{
            if (resultado.data.Comentarios.length != 0){
                let Usuario = localStorage.getItem('Usuario')
                Usuario = JSON.parse(Usuario)
                for (let i = 0; i < resultado.data.Comentarios.length; i++) {
                    const miComentario = {
                        Tipo: "Tienda",
                        NombreTienda: resultado.data.Comentarios[i].Tienda,
                        Departamento: resultado.data.Comentarios[i].Departamento,
                        Calificacion: parseInt(resultado.data.Comentarios[i].Calificacion),
                        Codigo: parseInt(resultado.data.Comentarios[i].Codigo),
                        Nombre:'Usuario',
                        Dpi: parseInt(resultado.data.Comentarios[i].Dpi),
                        Fecha: resultado.data.Comentarios[i].Fecha,
                        Hora: resultado.data.Comentarios[i].Hora,
                        Mensaje: resultado.data.Comentarios[i].Comentario,
                        SubComments:[]
                    }
                    nuevo2.push(miComentario)
                    
                }
                setComentarios(nuevo2)
                
            }else{
                setComentarios(comentariosprueba)
            }
        }
    }

    return (
        <div className="column carta">
            <div className="ui card">
                <div className="image">
                    <img src={props.imagen} />
                </div>
                <div className="content">
                    <div className="header">{props.nombre}</div>
                    <div className="meta">
                        <p>Departamento: {props.departamento}</p>
                        <p>Contacto: {props.contacto}</p>
                    </div>
                    <div className="description">Descripcion: {props.descripcion}</div>
                    <div className="ui segment green button center fluid" onClick={enviar}>Ver Productos</div>
                    <Modal
                    onClose={() => setOpen(false)}
                    onOpen={() => setOpen(true)}
                    open={open}
                    trigger={<Button className="ui segment green button center fluid" onClick={cargarComentarios}>Comentarios</Button>}
                    >
                    <Modal.Header>Comentarios</Modal.Header>
                    <Modal.Content image scrolling>
                        <Modal.Description>
                        <Comment.Group>
                            {comentarios.map((c, index) => (
                                <Comentario 
                                    Tipo={c.Tipo}
                                    NombreTienda = {c.NombreTienda}
                                    Departamento={c.Departamento}
                                    Calificacion={c.Calificacion}
                                    Codigo = {c.Codigo}
                                    Nombre={c.Nombre}
                                    Dpi={c.Dpi}
                                    Fecha={c.Fecha}
                                    Hora={c.Hora}
                                    Mensaje={c.Mensaje}
                                    SubComments={c.SubComments}
                                    key={index}
                                />
                            ))}

                            <Form reply>
                            <Form.TextArea onChange={(e)=>setMensaje(e.target.value)}/>
                            <Button content='Agregar Comentario' labelPosition='left' icon='edit' primary onClick={enviarComentario} />
                            </Form>
                        </Comment.Group>
                        </Modal.Description>
                    </Modal.Content>
                    <Modal.Actions>
                        <Button
                        content="Salir"
                        labelPosition='right'
                        icon='checkmark'
                        onClick={() => setOpen(false)}
                        positive
                        />
                    </Modal.Actions>
                    </Modal>
                </div>
                <div className="extra content">
                    <Rating defaultRating={props.calificacion} maxRating={5} disabled />
                </div>
            </div>
        </div>
    )
}

export default CartaTiendas
