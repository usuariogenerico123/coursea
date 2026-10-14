//import { useEffect, useState } from "react"
import { UserButton } from "@clerk/clerk-react"

export function MainContent( {tema} ){
    console.log("-----------------------------------------------------------")
    //console.log(tema)
    //console.log(tema.titulo_tema);

    const loQueAprenderas = tema.metas_aprendizaje;
    //console.log(loQueAprenderas);


    //const [videoUrl, setVideoUrl] = useState(video);
    const video = tema.url_video ? tema.url_video : tema.video_presentacion;
    


    return (
        
    // <!-- MAIN CONTENT (DERECHA) - VIDEO & INFO -->
    <main class="flex-1 flex flex-col h-full relative overflow-y-auto bg-midnight">
        
        {/* <!-- Top Nav (Flotante) --> */}
        <header class="w-full h-16 px-6 flex items-center justify-between border-b border-slate-800/50 bg-midnight/80 backdrop-blur-md sticky top-0 z-10 p-2">
            {/* <!-- Breadcrumbs --> */}
            <nav class="flex items-center text-sm text-slate-400 gap-2">
                <a href="#" class="hover:text-white transition-colors">Curso</a>
                <i class="fas fa-chevron-right text-[10px] opacity-50"></i>
                <a href="#" class="hover:text-white transition-colors">Módulo {(tema.id ? tema.id.replace(/\..*/, ""):"")}</a>
                <i class="fas fa-chevron-right text-[10px] opacity-50"></i>
                <span class="text-accent font-medium">{tema.id} {tema.titulo_tema}</span>
            </nav>

            {/* <!-- Acciones Derecha --> */}
            <div class="flex items-center gap-4">
                {/* <button class="w-8 h-8 rounded-full bg-slate-800 flex items-center justify-center text-slate-400 hover:text-white hover:bg-slate-700 transition-all" title="Modo Oscuro/Claro">
                    <i class="fas fa-moon"></i>
                </button> */}
                <div class="h-8 w-[1px] bg-slate-700"></div>
                {/* <button class="flex items-center gap-2 text-sm font-medium text-slate-300 hover:text-white transition-colors">
                   
                   
                    <img src="https://i.pravatar.cc/150?img=11" alt="User" class="w-8 h-8 rounded-full border-2 border-slate-700"/>
                    
                    <span class="hidden md:inline w-8 h-8 rounded-full border-2 border-slate-700">{<UserButton redirectUrl="/" />}</span>
                
                </button> */}
            </div>
        </header>

        {/* <!-- Contenido Central --> */}
        <div className="flex-1 max-w-5xl mx-auto w-full p-6 md:p-8">
            
            {/* <!-- Video Player Container (Aspect Ratio 16:9) --> */}
            <div class="relative w-full aspect-video bg-black rounded-xl overflow-hidden shadow-2xl video-glow mb-8 group border border-slate-800">
                
                <iframe width="100%" height="100%" src={video} title="YouTube video player" frameBorder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerPolicy="strict-origin-when-cross-origin" allowFullScreen></iframe>
                {/* <!-- Simulación de Thumbnail/Video --> */}
                {/* <div class="absolute inset-0 flex items-center justify-center bg-gradient-to-t from-black/80 via-transparent to-black/20 z-10">
                    <button class="w-16 h-16 md:w-20 md:h-20 bg-accent hover:bg-accent_hover text-white rounded-full flex items-center justify-center shadow-[0_0_30px_rgba(56,189,248,0.6)] transform group-hover:scale-110 transition-all duration-300 backdrop-blur-sm bg-opacity-90">
                        <i class="fas fa-play ml-1 text-2xl md:text-3xl"></i>
                    </button>
                </div> */}
                {/* <!-- Imagen de fondo fake --> */}
                {/* <img src="https://images.unsplash.com/photo-1633356122544-f134324a6cee?q=80&w=2070&auto=format&fit=crop" class="w-full h-full object-cover opacity-60 group-hover:opacity-40 transition-opacity duration-500" alt="Video cover"/>
                 */}
                {/* <!-- Barra de control simulada --> */}
                {/* <div class="absolute bottom-0 left-0 right-0 h-14 bg-gradient-to-t from-black to-transparent z-20 px-4 flex items-center gap-4 opacity-0 group-hover:opacity-100 transition-opacity duration-300">
                    <div class="h-1 flex-1 bg-slate-700 rounded-full cursor-pointer">
                        <div class="h-full bg-red-600 w-1/3 relative"></div>
                    </div>
                </div> */}
            </div>

            {/* <!-- Título y Botones de Acción --> */}
            <div class="flex flex-col md:flex-row md:items-start md:justify-between gap-4 mb-8">
                <div>
                    <h1 class="text-2xl font-bold text-white mb-2">{tema.titulo_tema ? tema.titulo_tema : tema.nombre_curso}</h1>
                    <p class="text-slate-400 text-sm">Publicado el 12 Oct, 2024 • Última actualización ayer</p>
                </div>
                <div class="flex gap-3">
                    <button class="px-4 py-2 rounded-lg bg-slate-800 hover:bg-slate-700 text-slate-300 text-sm font-medium transition-colors border border-slate-700">
                        <i class="far fa-heart mr-2 text-pink-500"></i> Favorito
                    </button>
                    <button class="px-4 py-2 rounded-lg bg-accent/10 hover:bg-accent/20 text-accent text-sm font-medium transition-colors border border-accent/20">
                        <i class="fas fa-check mr-2"></i> Completado
                    </button>
                </div>
            </div>

            {/* <!-- Tabs de Contenido --> */}
            <div class="border-b border-slate-800 mb-6">
                <nav class="flex gap-8" id="tabs-nav">
                    <button class="pb-3 border-b-2 border-accent text-accent font-medium text-sm transition-colors tab-btn" data-target="desc">
                        Descripción
                    </button>
                    <button class="pb-3 border-b-2 border-transparent text-slate-400 hover:text-slate-200 font-medium text-sm transition-colors tab-btn" data-target="recursos">
                        Recursos <span class="ml-1 px-1.5 py-0.5 bg-slate-800 rounded text-xs text-slate-500">3</span>
                    </button>
                    {/* <button class="pb-3 border-b-2 border-transparent text-slate-400 hover:text-slate-200 font-medium text-sm transition-colors tab-btn" data-target="comentarios">
                        Comentarios
                    </button> */}
                </nav>
            </div>

            {/* <!-- Contenido de las Tabs --> */}
            <div class="text-slate-300 leading-relaxed text-sm md:text-base min-h-[200px]">
                
                {/* <!-- Tab: Descripción --> */}
                <div id="desc" class="tab-content animate-fade-in">
                    <p class="mb-4">{tema.descripcion ? tema.descripcion : tema.descripcion_curso}</p>
                    
                    <h3 class="text-white font-semibold mt-6 mb-3 text-lg">Lo que aprenderás:</h3>
                    
                    <ul class="space-y-2 mb-6">
                    {
                        loQueAprenderas.map((item, index) => (
                            <li class="flex items-start gap-3">
                                <i class="fas fa-check text-green-500 mt-1"></i>
                                <span>{item}</span>
                            </li>
                           
                        ))
                    }
                        
                        {/* <li class="flex items-start gap-3">
                            <i class="fas fa-check text-green-500 mt-1"></i>
                            <span>Cómo reemplazar <code>componentDidMount</code> y <code>componentWillUnmount</code>.</span>
                        </li>
                        <li class="flex items-start gap-3">
                            <i class="fas fa-check text-green-500 mt-1"></i>
                            <span>El arreglo de dependencias y sus trampas comunes.</span>
                        </li>
                        <li class="flex items-start gap-3">
                            <i class="fas fa-check text-green-500 mt-1"></i>
                            <span>Funciones de limpieza (cleanup functions).</span>
                        </li> */}
                    </ul>
                </div>

                {/* <!-- Tab: Recursos (Oculto por defecto) --> */}
                <div id="recursos" class="hidden tab-content">
                    <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                        <a href="#" class="flex items-center p-4 rounded-xl border border-slate-700 bg-slate-800/50 hover:bg-slate-800 transition-colors group">
                            <div class="w-10 h-10 rounded-lg bg-blue-500/20 text-blue-400 flex items-center justify-center text-xl mr-4 group-hover:bg-blue-500 group-hover:text-white transition-all">
                                <i class="fab fa-github"></i>
                            </div>
                            <div>
                                <h4 class="text-white font-medium">Código Fuente (GitHub)</h4>
                                <p class="text-xs text-slate-500">Repositorio del proyecto final</p>
                            </div>
                            <i class="fas fa-external-link-alt ml-auto text-slate-600 group-hover:text-white"></i>
                        </a>

                        <a href="#" class="flex items-center p-4 rounded-xl border border-slate-700 bg-slate-800/50 hover:bg-slate-800 transition-colors group">
                            <div class="w-10 h-10 rounded-lg bg-red-500/20 text-red-400 flex items-center justify-center text-xl mr-4 group-hover:bg-red-500 group-hover:text-white transition-all">
                                <i class="fas fa-file-pdf"></i>
                            </div>
                            <div>
                                <h4 class="text-white font-medium">Slides de la clase</h4>
                                <p class="text-xs text-slate-500">PDF • 2.4 MB</p>
                            </div>
                            <i class="fas fa-download ml-auto text-slate-600 group-hover:text-white"></i>
                        </a>
                    </div>
                </div>
                
                {/* <!-- Tab: Comentarios (Oculto) --> */}
                <div id="comentarios" class="hidden tab-content">
                    <div class="flex gap-4 mb-6">
                        <img src="https://i.pravatar.cc/150?img=11" alt="User" class="w-10 h-10 rounded-full border border-slate-700"/>
                        <div class="flex-1">
                            <textarea class="w-full bg-slate-900 border border-slate-700 rounded-lg p-3 text-sm focus:border-accent focus:ring-1 focus:ring-accent outline-none resize-none h-24" placeholder="Escribe una pregunta o comentario..."></textarea>
                            <div class="flex justify-end mt-2">
                                <button class="bg-accent hover:bg-accent_hover text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors">Publicar</button>
                            </div>
                        </div>
                    </div>
                    {/* <!-- Comentario Ejemplo --> */}
                    <div class="flex gap-4 p-4 rounded-lg hover:bg-slate-800/30 transition-colors">
                        <img src="https://i.pravatar.cc/150?img=33" alt="User" class="w-10 h-10 rounded-full"/>
                        <div>
                            <div class="flex items-center gap-2 mb-1">
                                <span class="font-semibold text-white text-sm">Sarah Connor</span>
                                <span class="text-xs text-slate-500">hace 2 horas</span>
                            </div>
                            <p class="text-slate-300 text-sm">Excelente explicación sobre el array de dependencias, siempre me causaba loops infinitos. ¡Gracias!</p>
                        </div>
                    </div>
                </div> 

            </div>
        </div>
    </main>


    )
}