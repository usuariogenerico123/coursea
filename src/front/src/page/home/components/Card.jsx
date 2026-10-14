export function Card({data}){
    console.log(data.titulo)

    const izquierda = "ml-16 md:ml-0 md:mr-auto md:w-5/12 bg-crypto-card p-6 rounded-2xl border border-slate-700/50 shadow-lg relative group hover:border-accent-blue/50 transition-all cursor-pointer";
    const derecha = "ml-16 md:ml-auto md:w-5/12 bg-crypto-card p-6 rounded-2xl border border-slate-700/50 shadow-lg relative group hover:border-accent-green/50 transition-all cursor-pointer"
    

    const iconoAzul = "absolute -top-4 -left-4 w-12 h-12 bg-accent-blue flex items-center justify-center font-black text-xl text-white rounded-xl rotate-3 group-hover:rotate-0 transition"
    const iconoVerde = "absolute -top-4 -right-4 w-12 h-12 bg-accent-green flex items-center justify-center font-black text-xl text-crypto-dark rounded-xl -rotate-3 group-hover:rotate-0 transition"
    
    
    const circuloVerde = "absolute left-4 md:left-1/2 -ml-2 md:-ml-3 w-6 h-6 bg-accent-green border-4 border-crypto-dark rounded-full shadow-[0_0_15px_rgba(16,185,129,0.5)]";
    const circuloAzul = "absolute left-4 md:left-1/2 -ml-2 md:-ml-3 w-6 h-6 bg-accent-blue border-4 border-crypto-dark rounded-full shadow-[0_0_15px_rgba(59,130,246,0.5)]"
    
    return (
        <div class="relative z-10 mb-12 md:mb-24">
                <div class="flex items-center">
                     <div class={(data.posicion == "izquierda")? circuloAzul : circuloVerde}></div>
                    
                    <div class={(data.posicion === "izquierda")? izquierda : derecha}>
                         <div class={(data.posicion == "izquierda")? iconoAzul : iconoVerde}>M{data.modulo}</div>
                        
                        <h3 class="text-2xl font-bold text-white mb-3 mt-2">{data.titulo}</h3>
                        <p class="text-slate-400 mb-4">{data.descripcion}.</p>
                        <ul class="text-sm text-slate-500 space-y-2">
                            {
                                data.lista.map((item)=>(
                                    <li><i class={(data.posicion == "izquierda")? "fa-solid fa-check text-accent-blue mr-2": "fa-solid fa-check text-accent-green mr-2"}></i>{item}</li>
                                ))
                            }                           
                        </ul>
                    </div>
                </div>
            </div>
    )

}