import React,{useEffect,useState} from 'react'
import { Button, Header, Image, Modal,Comment, Form } from 'semantic-ui-react'
import axios from "axios"

function Comentario(props) {
    const [mensaje, setMensaje] = useState("")
    const [activo, setActivo] = useState(false)

    const enviarComentario = async()=> {
        let Usuario = localStorage.getItem('Usuario')
        Usuario = JSON.parse(Usuario)
        let comentario = {
            Tipo: props.Tipo,
            Tienda: props.NombreTienda,
            Departamento: props.Departamento,
            Calificacion: parseInt(props.Calificacion),
            Codigo: parseInt(props.Codigo),
            Dpi: parseInt(Usuario.Dpi) ,
            Comentario: mensaje
        }
        console.log(props.key)
        await axios.post("http://localhost:3000/IngresarComentario",comentario)
    }

    return (
        <Comment>
        <Comment.Content>
        <Comment.Author as='a'>{props.Dpi} - {props.Nombre}</Comment.Author>
            <Comment.Metadata>
            <div>{props.Fecha} - {props.Hora}</div>
            </Comment.Metadata>
            <Comment.Text>{props.Mensaje}</Comment.Text>
            <Comment.Actions>
            <Comment.Action as = 'a' onClick={()=>{setActivo(!activo)}}>Responder</Comment.Action>
            </Comment.Actions>
        </Comment.Content>
        {activo?
        <Form reply>
        <Form.TextArea style={{minHeight:50,maxHeight:50}} onChange={(e)=>setMensaje(e.target.value)}/>
        <Button content='Agregar Comentario' labelPosition='left' icon='edit' primary onClick={enviarComentario} />
        </Form>
        :
        <></>
        }
        
        {((props.SubComments != null) && (props.SubComments.length>0))?
        <Comment.Group>
            {props.SubComments.map((c, index) => (
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
        </Comment.Group>
        :
        <></>
        }
        </Comment>
    )
}

export default Comentario
