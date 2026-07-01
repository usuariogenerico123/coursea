import { SignedOut, SignIn, SignInButton, UserButton, SignOutButton, SignedIn } from "@clerk/clerk-react"
import { Link } from "react-router-dom"



export function Menu({ boton }) {
    

    return (
        <div className=" relative p-5 border-white border-2 rounded-3xl w-50 m-20 flex flex-col items-center gap-6">

            {/* <a href="" class="text-sm font-bold text-slate-300 hover:text-accent-green mr-6 transition hidden md:inline-block">
                    <Link to="/aula" ><i class="fa-solid fa-right-to-bracket mr-2"></i> */}
            <button onClick={boton} className="text-left m-0 p-0">X</button>
            <div>
                <SignedOut>
                    <i class="fa-solid fa-right-to-bracket mr-2"></i>
                    <span className="text-sm font-bold text-slate-300 hover:text-accent-green mr-6 transition  md:inline-block"><SignInButton mode="modal" forceRedirectUrl="/aula">  Iniciar Sesion </SignInButton></span>

                </SignedOut>
            </div>
            
            <SignedIn >
                <Link to="/aula"><i class="fa-solid fa-right-to-bracket mr-2"></i>
                    <span className=" text-sm font-bold text-slate-300 hover:text-accent-green mr-6 transition  md:inline-block">Area de alumnos </span>
                </Link>

            </SignedIn>
            {/*                             
                    </Link>
                    
                </a> */}
            <a href="https://wa.me/TU_NUMERO" class="mb-6 bg-accent-green/90 hover:bg-accent-green text-crypto-dark font-bold px-6 py-2 rounded-full text-sm transition-all hover:shadow-lg hover:shadow-accent-green/20">
                Empezar Ahora
            </a>
        </div>
    )
}



