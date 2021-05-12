import React,{useEffect,useState} from 'react'
import { Button, Header, Image, Modal,Comment, Form } from 'semantic-ui-react'
import Comentario from './Comentario'
import axios from "axios"

function CartaProducto(props) {
    const [open, setOpen] = React.useState(false) 
    const [mensaje, setMensaje] = useState("")
    const [comentarios, setComentarios] = useState([])
    const nuevo2 = []
    
    const enviar = () =>{
        var llave = localStorage.getItem('CPedido')
        if (llave === null || llave === undefined){
            llave=0
            localStorage.setItem('CPedido',0)
        }else{
            llave++
            localStorage.setItem('CPedido',llave)
        }
        const myproducto = {
            "NombreTienda": props.nombreTienda,
            "Departamento": props.departamento,
            "Calificacion": props.calificacion,
            "Nombre": props.nombre,
            "Codigo": props.codigo,
            "Descripcion": props.descripcion,
            "Precio": props.precio,
            "Cantidad": props.cantidad,
            "Imagen": props.imagen,
            "id":llave
        }
        var datos = localStorage.getItem('Carrito')
        if (datos === null || datos === undefined){
            localStorage.setItem('Carrito',JSON.stringify([myproducto]))
        }else{
            datos = JSON.parse(datos)
            datos.push(myproducto)
            localStorage.setItem('Carrito',JSON.stringify(datos))
        }
        alert("Producto " + props.nombre + "Añadido")
    }

    

    const enviarComentario = async()=> {
        let Usuario = localStorage.getItem('Usuario')
        Usuario = JSON.parse(Usuario)
        let comentario = {
            Tipo: "Producto",
            Tienda: props.nombreTienda,
            Departamento: props.departamento,
            Calificacion: parseInt(props.calificacion),
            Codigo: parseInt(props.codigo),
            Dpi: parseInt(Usuario.Dpi) ,
            Comentario: mensaje
        }
        await axios.post("http://localhost:3000/IngresarComentario",comentario)
        const comentariosprueba = []
        const resultado = await axios.post("http://localhost:3000/MostrarComentario",comentario)
        console.log(resultado.data)
        if(resultado.data == undefined || resultado.data == null ){
            setComentarios(comentariosprueba)
        }else{
            if (resultado.data.Comentarios.length != 0){
                for (let i = 0; i < resultado.data.Comentarios.length; i++) {
                    const miComentario = {
                        Tipo: "Producto",
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

    const cargarComentarios = async()=> {
        let Usuario = localStorage.getItem('Usuario')
        Usuario = JSON.parse(Usuario)
        const comentariosprueba = []
        let comentario = {
            Tipo: "Producto",
            Tienda: props.nombreTienda,
            Departamento: props.departamento,
            Calificacion: parseInt(props.calificacion),
            Codigo: parseInt(props.codigo),
            Dpi: parseInt(Usuario.Dpi),
            Comentario: ""
        }
        const resultado = await axios.post("http://localhost:3000/MostrarComentario",comentario)
        console.log(resultado.data)
        if(resultado.data == undefined || resultado.data == null ){
            setComentarios(comentariosprueba)
        }else{
            if (resultado.data.Comentarios.length != 0){
                for (let i = 0; i < resultado.data.Comentarios.length; i++) {
                    const miComentario = {
                        Tipo: "Producto",
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
                        <p>Codigo: {props.codigo}</p>
                    </div>
                    <div className="description">Descipcion: {props.descripcion}</div>
                    <div className="ui segment green button center fluid" onClick={enviar}>Comprar</div>
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
                    <span className="right floated">Cantidad: {props.cantidad}</span>
                    <span><i className="dollar sign icon" />{props.precio}</span>
                </div>
            </div>
        </div>
    )
}

export default CartaProducto
