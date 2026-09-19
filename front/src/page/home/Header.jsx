import { Link } from "react-router-dom"

import { SignedOut, SignIn, SignInButton, UserButton, SignOutButton, SignedIn } from "@clerk/clerk-react"
import { useState } from "react";
import { Menu } from "./Menu";



export function Header() {

    const [verMenu, cerrarMenu] = useState(false);
    const boton = ()=>{
        cerrarMenu(!verMenu);
        console.log("siu");
    }

    return (
        <nav class="fixed w-full top-0 z-50 bg-crypto-dark/80 backdrop-blur-lg border-b border-slate-800/50 py-4 ">
            
            <div class=" max-w-6xl mx-auto px-6 flex justify-between items-center">
                <div class="flex justify-center items-center gap-2 font-black sm:text-xl text-white tracking-wider">
                    <i class="fa-solid fa-shield-halved text-accent-green"></i>
                    RUTA<span class="text-accent-green">DOLARIZADA</span>
                </div>
                <div className="hidden sm:flex flex-col sm:flex-row items-center sm:flex-row">

                    {/* <a href="" class="text-sm font-bold text-slate-300 hover:text-accent-green mr-6 transition hidden md:inline-block">
                    <Link to="/aula" ><i class="fa-solid fa-right-to-bracket mr-2"></i> */}
                    <SignedOut>
                        <i class="fa-solid fa-right-to-bracket mr-2"></i>
                        <span className="text-sm font-bold text-slate-300 hover:text-accent-green mr-6 transition  md:inline-block"><SignInButton mode="modal" forceRedirectUrl="/aula">  Iniciar Sesion </SignInButton></span>

                    </SignedOut>
                    <SignedIn >
                        <Link to="/aula"><i class="fa-solid fa-right-to-bracket mr-2"></i>
                            <span className=" text-sm font-bold text-slate-300 hover:text-accent-green mr-6 transition  md:inline-block">Area de alumnos </span>
                        </Link>

                    </SignedIn>
                    {/*                             
                    </Link>
                    
                </a> */}
                    <a href="https://wa.me/TU_NUMERO" class=" bg-accent-green/90 hover:bg-accent-green text-crypto-dark font-bold px-6 py-2 rounded-full text-sm transition-all hover:shadow-lg hover:shadow-accent-green/20">
                        Empezar Ahora
                    </a>
                </div>
                <button className="sm:hidden" onClick={boton}>
                    <svg class="size-6 " xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" >
                        <path stroke-linecap="round" stroke-linejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
                    </svg>
                
                </button>
            </div>
            
            {verMenu && <Menu boton={boton}/>}
        </nav>
    )
}