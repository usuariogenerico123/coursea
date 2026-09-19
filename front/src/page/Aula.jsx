import { Sidebar } from "./aula/Sidebar"
import { MainContent } from "./aula/MainContent"
import { useState } from "react";
import { useLocation } from "react-router-dom";
import { Protect } from "@clerk/clerk-react";
import cursos from "./aula/data/cursos.json";
import { getDataFromCache } from "./aula/scripts/cache";



export function Aula(){

    
    const location = useLocation();
    const path = location.hash;
    //const pathVideo = path ? "https://www.youtube.com/embed/"+path.replace("#", "") : "https://www.youtube.com/embed/dQw4w9WgXcQ?si=kOLNeXiAi5U8mNvw"
    //console.log(pathVideo);

    
    const presentacion = {
            "nombre_curso": cursos.nombre_curso,
            "nombre_tutor": cursos.nombre_tutor,
            "video_presentacion": cursos.video_presentacion,
            "descripcion_curso": cursos.descripcion_curso,
            "metas_aprendizaje": cursos.metas_aprendizaje
        }
    const [tema, setTema] = useState((path ? getDataFromCache() : presentacion));
    const verTema = (data) =>{
        console.log("---------")
        console.log(data);
        setTema(data); 
    }

    



    return (
        <Protect>
                <div className="text-slate-300 md:h-screen overflow-hidden flex flex-col md:flex-row">
                    
                    <Sidebar botonVideo={verTema} cursos={cursos}/>

 
                    <MainContent tema={tema}/>

                </div>
        </Protect>
        
        
        
            
        
    )
}

