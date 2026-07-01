import { useState } from "react";
import { saveCache } from "../scripts/cache";

export function InfoTarget({ datos, botonVideo, isSelected, onselect}) {
    //const url = "https://www.youtube.com/embed/Z-3su2ov8Vg?si=k7TLC3FcrfFTvSKT"
    const yt_id = datos.url_video.replace("https://www.youtube.com/embed/", "");
    //console.log(datos.url_video)



    const check_circle = "fas fa-check-circle text-green-500 text-sm";
    const no_check = "far fa-circle text-slate-600 group-hover:text-slate-400"
    

    return (
        <a onClick={ ()=>{ 
            botonVideo(datos);
            saveCache(datos); 
            console.log(datos);
            onselect();
            
            }} href={"#"+yt_id} class="flex items-center gap-3 px-5 py-3 border-l-[3px] border-transparent hover:border-slate-600 hover:bg-slate-800/30 transition-colors group opacity-70 hover:opacity-100">
            <div class="flex-shrink-0 w-5 text-center">
                <div class="w-2 h-2 rounded-full bg-accent animate-pulse mx-auto">
                
                </div>
                <i class={isSelected ? check_circle : no_check}></i>
            </div>

            <div class="flex-1 min-w-0">
            <p class="text-sm text-slate-400 truncate group-hover:text-slate-200">{datos.titulo_tema}</p>
                {/* <p class="text-sm text-slate-400 truncate group-hover:text-slate-200">{datos.titulo}</p> */}
                <span class="text-[10px] text-slate-600 flex items-center gap-1 mt-0.5">
                     <i class="far fa-play-circle"></i> {datos.duracion} {datos.status}
                </span>
            </div>
        </a>

            
    )
}