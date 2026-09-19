import { Modulo } from "./components/Modulo"

import { useState } from "react";
import { UserButton } from "@clerk/clerk-react";

export function Sidebar({ botonVideo, cursos }) {



    //console.log(cursos)
    const nombre_curso = cursos.nombre_curso;
    const tutor = cursos.nombre_tutor;
    const modulos = cursos.modulos;



    const [selected, setSelected] = useState(null);
    


    return (

        // <!-- SIDEBAR (IZQUIERDA) - CURRICULUM -->
        <aside class="w-full md:w-80 bg-surface border-r border-slate-800 flex flex-col h-full z-20 flex-shrink-0">

            {/* <!-- Header del Sidebar --> */}
            <div class="p-5 border-b border-slate-800 bg-surface/95 backdrop-blur-sm sticky top-0">
                <div class="flex items-center gap-3 mb-4">
                    <div class="w-10 h-10 rounded-lg bg-gradient-to-br from-blue-500 to-purple-600 flex items-center justify-center text-white font-bold shadow-lg">
                        
                        { <UserButton redirectUrl="/" /> }
                    </div>
                    <div>
                        <h2 class="text-white font-semibold text-sm leading-tight">{nombre_curso}</h2>
                        <p class="text-xs text-slate-500 mt-0.5">{tutor}</p>
                    </div>
                </div>

                {/* <!-- Barra de Progreso --> */}
                <div class="space-y-2">
                    <div class="flex justify-between text-xs font-medium">
                        <span class="text-slate-400">Tu progreso</span>
                        <span class="text-accent">35%</span>
                    </div>
                    <div class="h-1.5 w-full bg-slate-800 rounded-full overflow-hidden">
                        <div class="h-full bg-accent w-[35%] rounded-full shadow-[0_0_10px_rgba(56,189,248,0.5)]"></div>
                    </div>
                </div>
            </div>

           


            {/* <!-- Lista de Módulos (Scrollable) --> */}
            <div class="flex-1 overflow-y-auto" id="curriculum-list">


                {
                    modulos.map((curso, index) => (
                        <Modulo key={index} datosModulo={curso} boton={botonVideo} selected={selected } setSelected={setSelected}/>
                    ))
                }

            </div>



            {/* <!-- Footer Sidebar --> */}
            <div class="p-4 border-t border-slate-800 mt-auto bg-surface">
                <button class="w-full py-2 px-4 rounded-lg border border-slate-700 hover:bg-slate-800 text-xs font-medium text-slate-300 flex items-center justify-center gap-2 transition-colors">
                    <i class="fas fa-certificate text-yellow-500"></i>
                    Solicitar Certificado
                </button>
            </div>
        </aside>
    )
}