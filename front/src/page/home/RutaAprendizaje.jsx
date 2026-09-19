import { Card } from "./components/Card"

export function RutaAprendizaje() {

    const fundamentos = {
        modulo: 1,
        titulo: "Fundamentos del Dinero Digital",
        descripcion: `Olvida la teoría aburrida. Entiende por qué USDT vale $1, qué es una "Red" de envío (ERC20/TRC20) y por qué no es lo mismo que Bitcoin`,
        lista: ["Stablecoins vs. Volátiles", "Cómo elegir la red correcta para no perder dinero."],
        posicion: "izquierda"//left-right
    }

    const operatoria = {
        modulo: 2,
        titulo: "La Operatoria: Comprar sin Miedo",
        descripcion: `La parte práctica. Entramos a los Exchanges (Bybit/Binance) y aprendemos a usar el mercado P2P para cambiar tu moneda local por dólares.`,
        lista: ["Filtrar comerciantes P2P verificados.", "El protocolo para evitar bloqueos bancarios."],
        posicion: "derecha"//left-right
    }

    const autocustodia = {
        modulo: 3,
        titulo: "Autocustodia: Sé tu Propio Banco",
        descripcion: `Si el dinero está en el Exchange, no es tuyo. Aprende a instalar y usar Wallets privadas (MetaMask/Trust) correctamente.`,
        lista: ["Creación de Frase Semilla.", "Primera transferencia de prueba a tu bóveda."],
        posicion: "izquierda"//left-right
    }

    const utilidad = {
        modulo: 4,
        titulo: "Utilidad: Gastar y Usar",
        descripcion: `¿De qué sirve tener dólares si no los puedes usar? Aprende a conectarlos con el mundo real.`,
        lista: ["Tarjetas Visa/Mastercard recargables con Crypto.", "Off-Ramp: Volver a moneda local cuando lo necesites."],
        posicion: "derecha"//left-right
    }

    return (

        <section id="ruta-aprendizaje" class="py-24 px-6 relative">
            <div class="max-w-4xl mx-auto">
                <div class="text-center mb-20">
                    <h2 class="text-4xl font-black text-white uppercase tracking-wide mb-4">El Mapa del Curso</h2>
                    <p class="text-lg text-slate-400">5 pasos directos. Desde cero absoluto hasta la seguridad total.</p>
                </div>



                <div class="relative">
                    <div class="absolute left-4 md:left-1/2 h-full w-1 bg-slate-800/50 -ml-0.5 transform md:-translate-x-1/2"></div>


                    <Card data={fundamentos} />

                    <Card data={operatoria} />

                    <Card data={autocustodia} />

                    <Card data={utilidad} />


                    <div class="relative z-10">
                        <div class="max-w-3xl mx-auto bg-crypto-card/50 p-8 rounded-3xl border-2 border-security-red/50 shadow-[0_0_30px_rgba(239,68,68,0.15)] relative overflow-hidden">
                            <div class="absolute inset-0 opacity-10 bg-[radial-gradient(circle_at_center,_var(--tw-gradient-stops))] from-security-red via-transparent to-transparent"></div>

                            <div class="relative z-20 text-center">
                                <div class="w-20 h-20 bg-security-red mx-auto rounded-full flex items-center justify-center text-4xl text-white mb-6 shadow-lg shadow-security-red/30 animate-pulse">
                                    <i class="fa-solid fa-shield-virus"></i>
                                </div>

                                <h3 class="text-3xl font-black text-white mb-2 uppercase tracking-wider">Módulo 5: El Escudo Final</h3>
                                <p class="text-security-red font-bold mb-6">Aspectos Críticos de Seguridad y Prevención de Fraudes</p>

                                <p class="text-slate-300 mb-8 text-lg leading-relaxed">
                                    Este es el módulo más importante. Aquí aprenderás a blindarte contra los errores humanos y las estafas que hacen que la gente pierda su dinero. <strong class="text-white">Sin esto, nada de lo anterior importa.</strong>
                                </p>

                                <div class="grid sm:grid-cols-3 gap-4 text-left">
                                    <div class="bg-crypto-dark p-4 rounded-xl border border-slate-800 flex items-start gap-3">
                                        <i class="fa-solid fa-triangle-exclamation text-security-red mt-1"></i>
                                        <div>
                                            <h4 class="font-bold text-white">Anti-Phishing</h4>
                                            <p class="text-xs text-slate-500">Detectar sitios y correos falsos al instante.</p>
                                        </div>
                                    </div>
                                    <div class="bg-crypto-dark p-4 rounded-xl border border-slate-800 flex items-start gap-3">
                                        <i class="fa-solid fa-mobile-screen-button text-security-red mt-1"></i>
                                        <div>
                                            <h4 class="font-bold text-white">Higiene Digital</h4>
                                            <p class="text-xs text-slate-500">Autenticación 2FA correcta (No SMS) y gestores.</p>
                                        </div>
                                    </div>
                                    <div class="bg-crypto-dark p-4 rounded-xl border border-slate-800 flex items-start gap-3">
                                        <i class="fa-solid fa-user-secret text-security-red mt-1"></i>
                                        <div>
                                            <h4 class="font-bold text-white">Anti-Estafas</h4>
                                            <p class="text-xs text-slate-500">Identificar esquemas Ponzi y promesas falsas.</p>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    )
}