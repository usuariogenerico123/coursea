import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server : {
    allowedHosts:[
        "hltef-34-46-161-55.a.free.pinggy.link"
    ]
  }
})
