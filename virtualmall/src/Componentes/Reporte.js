import React,{useEffect,useState} from 'react'
import "../css/ImportTiendas.css"
import axios from "axios"
import NavBar from "../Componentes/NavBar"

function UserList() {
    const [arbolSC, setArbolSC]=useState('')
    const [arbolCS, setArbolCS]=useState('')
    const [arbolC, setArbolC]=useState('')
    const [grafo, setGrafo]=useState('')
    const [años, setaños]=useState('')
    const [añom, setAñom]=useState('')
    const [mes, setMes]=useState('')
    const [mesm, setMesm]=useState('')
    const [Matriz, setMatriz]=useState('')
    const [categoria, setCategoria]=useState('')
    const [dia, setDia]=useState('')
    const [cola, setCola]=useState('')
    const [llave, setLlave]=useState('')


    const mostrarASC = async()=>{
        const data = await axios.get("http://localhost:3000/ObtenerUsuariosSC")
        setArbolSC("data:image/png;base64,"+data.data)
    }

    const mostrarACS = async()=>{
        if (llave === ""){
            alert("Ingrese una llave")
        }else{
            if (llave.length < 32){
                let pal = ""
                pal = llave
                for (let index = llave.length; index < 32; index++){
                    pal += "a"
                }
                const data = await axios.get("http://localhost:3000/ObtenerUsuariosCS/"+pal)
                setArbolCS("data:image/png;base64,"+data.data)
                const data2 = await axios.get("http://localhost:3000/ObtenerUsuariosC/"+pal)
                setArbolC("data:image/png;base64,"+data2.data)
                alert("Grafos Cargados")
            }else if (llave.length === 32){
                const data = await axios.get("http://localhost:3000/ObtenerUsuariosCS/"+llave)
                setArbolCS("data:image/png;base64,"+data.data)
                const data2 = await axios.get("http://localhost:3000/ObtenerUsuariosC/"+llave)
                setArbolC("data:image/png;base64,"+data2.data)
                alert("Grafos Cargados")
            }
        }
    }

    const mostrarGrafo = async()=>{
        const data = await axios.get("http://localhost:3000/ObtenerGrafo")
        setGrafo("data:image/png;base64,"+data.data)
    }

    useEffect(() => {
        async function mostarAños() {
            const data = await axios.get("http://localhost:3000/GrafoAños") 
            setaños("data:image/png;base64,"+data.data)
        }
        mostarAños()
    },[]);


    const mostrarMes = async()=>{

        const data = await axios.get("http://localhost:3000/GrafoMesesAños/"+añom)
        setMes("data:image/png;base64,"+data.data)
    }

    const mostrarMatriz = async(event)=>{

        const data = await axios.get("http://localhost:3000/GrafoMatrizMesesAños/"+mesm)
        setMatriz("data:image/png;base64,"+data.data)
    }

    const mostrarCola = async(event)=>{

        const data = await axios.get("http://localhost:3000/GrafoColaMatrizMesesAños/"+dia+"/"+categoria)
        setCola("data:image/png;base64,"+data.data)
    }

    return (
        <>
        <NavBar 
            colores={["red","green","yellow","blue","grey"]}
            opciones={["Cargar Archivos","Reportes","Tiendas","Cerrar Sesion", "Eliminar Mi Cuenta"]}
            url={["/cargar","/reporte","/listaTR","/iniciosesion","/iniciosesion"]}
            activo={"green"}
        />
        <div className="ImportTiendas">
            <br/>
            <div className="ui inverted segment container items">

                <div className="item">
                    <div className="ui big segment rounded image">
                        <img src={arbolSC} width={600} height={600}/>
                    </div>
                    <div className="content">
                    <h1 style={{color: '#00FFFF'}}>Arbol Sin Cifrar</h1>
                        <div className="extra">
                            <div className="ui segment green button center" onClick={()=>{mostrarASC()}}>Cargar Arbol</div>
                        </div>
                    </div>
                </div>
                <div className="ui inverted divider" />

                <div className="item">
                    <div className="ui big segment rounded image">
                        <img src={arbolCS} width={600} height={600}/>
                    </div>
                    <div className="content">
                    <h1 style={{color: '#00FFFF'}}>Arbol Cifrado Simple</h1>
                        <div className="meta">
                            <input type="text" placeholder="Ingrese Llave De Cifrado" onChange={e=>setLlave(e.target.value)}></input>
                        </div>
                        <div className="extra">
                            <div className="ui segment green button center" onClick={()=>{mostrarACS()}}>Cargar Arbol</div>
                        </div>
                    </div>
                </div>
                <div className="ui inverted divider" />

                <div className="item">
                    <div className="ui big segment rounded image">
                        <img src={arbolC} width={600} height={600}/>
                    </div>
                    <div className="content">
                    <h1 style={{color: '#00FFFF'}}>Arbol Cifrado</h1>

                    </div>
                </div>
                <div className="ui inverted divider" />

                <div className="item">
                    <div className="ui big segment rounded image">
                        <img src={grafo} width={600} height={600}/>
                    </div>
                    <div className="content">
                    <h1 style={{color: '#00FFFF'}}>Grafo</h1>
                        <div className="extra">
                            <div className="ui segment green button center" onClick={()=>{mostrarGrafo()}}>Cargar Grafo</div>
                        </div>
                    </div>
                </div>
                <div className="ui inverted divider" />

                <div className="item">
                    <div className="ui big segment rounded image">
                        <img src={años} width={600} height={600}/>
                    </div>
                    <div className="content">
                    <h1 style={{color: '#00FFFF'}}>Años</h1>
                        <div className="meta">
                            <input type="text" placeholder="Ingresar Año" onChange={e=>setAñom(e.target.value)}></input>
                        </div>
                        <div className="extra">
                            <div className="ui segment green button center" onClick={()=>{mostrarMes()}}>Cargar Grafica Del Mes Solicitado</div>
                        </div>
                    </div>
                </div>
                <div className="ui inverted divider" />


                <div className="item">
                    <div className="ui big segment rounded image">
                        <img src={mes} width={600} height={600}/>
                    </div>
                    <div className="content">
                    <h1 style={{color: '#00FFFF'}}>Mes</h1>
                    <div className="meta">
                        <input type="text" placeholder="Ingrese Mes" onChange={e=>setMesm(e.target.value)}></input>
                    </div>
                    <div className="extra">
                        <div className="ui segment green button center" onClick={()=>{mostrarMatriz()}}>Cargar Grafica De La Matriz Solicitada</div>
                    </div>
                    </div>
                </div>
                <div className="ui inverted divider" />

                
                <div className="item">
                    <div className="ui big segment rounded image">
                        <img src={Matriz} width={600} height={600}/>
                    </div>
                    <div className="content">
                    <h1 style={{color: '#00FFFF'}}>Matriz De Pedidos</h1>
                        <div className="meta">
                            <input type="text" placeholder="Ingresar Dia" onChange={e=>setDia(e.target.value)}></input><br/><br/>
                            <input type="text" placeholder="Ingresar Categoria" onChange={e=>setCategoria(e.target.value)}></input>
                        </div>
                        <div className="extra">
                            <div className="ui segment green button center" onClick={()=>{mostrarCola()}}>Cargar Grafica De La Cola Solicitada</div>
                        </div>
                    </div>
                </div>
                <div className="ui inverted divider" />
                <div className="item">
                    <div className="ui big segment rounded image">
                        <img src={cola} width={600} height={600}/>
                    </div>
                    <div className="content">
                    <h1 style={{color: '#00FFFF'}}>Cola De Pedidos</h1>
                    </div>
                </div>
            </div>
            <br/>
        </div>
        </>
    )
}

export default UserList
