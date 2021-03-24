import React,{useEffect,useState} from 'react'
import "../css/ImportTiendas.css"
import axios from "axios"

function UserList() {
    const [años, setaños]=useState('')
    const [añom, setAñom]=useState('')
    const [mes, setMes]=useState('')
    const [mesm, setMesm]=useState('')
    const [Matriz, setMatriz]=useState('')
    const [categoria, setCategoria]=useState('')
    const [dia, setDia]=useState('')
    const [cola, setCola]=useState('')


    useEffect(() => {
        async function mostarAños() {
            const data = await axios.get("http://localhost:3000/GrafoAños") 
            alert(data.data)
            setaños(data.data)
        }
        mostarAños()
    },[]);


    const mostrarMes = async()=>{
        console.log(añom)
        const data = await axios.get("http://localhost:3000/GrafoMesesAños/"+añom)
        setMes(data.data)
        alert(data.data)
    }

    const mostrarMatriz = async(event)=>{
        console.log(mesm)
        const data = await axios.get("http://localhost:3000/GrafoMatrizMesesAños/"+mesm)
        setMatriz(data.data)
        alert(data.data)
    }

    const mostrarCola = async(event)=>{
        console.log(categoria)
        console.log(dia)
        const data = await axios.get("http://localhost:3000/GrafoColaMatrizMesesAños/"+dia+"/"+categoria)
        setCola(data.data)
        alert(data.data)
    }

    return (
        <div className="ImportTiendas">
            <br/>
            <div className="ui inverted segment container items">
                <div className="item">
                    <div className="ui big segment rounded image">
                        <img src={`data:años/jpeg;base64,${años}`} width={600} height={600}/>
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
                        <img src={`data:Mes/jpeg;base64,${mes}`} width={600} height={600}/>
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
                        <img src={`data:Mes/jpeg;base64,${Matriz}`} width={600} height={600}/>
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
                        <img src={`data:Mes/jpeg;base64,${cola}`} width={600} height={600}/>
                    </div>
                    <div className="content">
                    <h1 style={{color: '#00FFFF'}}>Cola De Pedidos</h1>
                    </div>
                </div>
            </div>
            <br/>
        </div>
    )
}

export default UserList
