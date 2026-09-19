import { Main } from "./home/Main";
import { Header } from "./home/Header";
import { CtaSection } from "./home/CtaSection";
import { RutaAprendizaje } from "./home/RutaAprendizaje";




export function Home(){

    return (
        <>
        <Header />
        <Main />
        <RutaAprendizaje />
        <CtaSection />
        </>
    )
}