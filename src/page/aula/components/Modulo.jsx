import { InfoTarget } from "./InfoTarget"
import { useState } from "react"

export function Modulo({datosModulo, boton, selected , setSelected}) {

    const [mostrar, setMostrar] = useState(false);
    const verClases = () =>{
        setMostrar(!mostrar);
    }


    
    const temas = datosModulo.temas



    return (

        <div class="border-b border-slate-800/50">
            <button onClick={verClases} class="w-full px-5 py-4 flex items-center justify-between hover:bg-slate-800/30 transition-colors group opacity-80 hover:opacity-100">
                <div class="text-left">
                    <p class="text-xs font-bold text-slate-500 uppercase tracking-wider mb-1">Modulo {datosModulo.numero_modulo}</p>
                    <h3 class="text-sm font-semibold text-slate-300">{datosModulo.titulo_modulo}</h3>
                </div>
                
                <i class=  {(mostrar ? "fas fa-chevron-down text-slate-500 group-hover:text-white transition-transform duration-300" : "fas fa-chevron-right text-slate-600 group-hover:text-slate-400")}></i>
            </button>

            <div className="bg-crypto-ark/50 " style={{ display:(mostrar ? "block" : "none") }}>
                {/* contenedor para los temas  */}
                
                {
                    temas.map((dato, index) => (
                        //console.log(dato.id),
                        <InfoTarget key={index} datos={dato} botonVideo={boton} isSelected={selected === dato.id} onselect={()=>{setSelected(dato.id)}}/>
                    ))
                }
            </div>
        </div>

    )
}





