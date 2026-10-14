
import { BrowserRouter, Routes, Route } from 'react-router-dom'

import { Home } from './page/Home'
import { Aula } from './page/Aula'
import { ProtectedRoute } from './security/ProtectedRoute'

function App() {
 
  


  return (
    <BrowserRouter>
      <Routes>
        <Route path='/' element={<Home />} />
        
        <Route element={<ProtectedRoute />}>
          <Route path="/aula" element={ <Aula />}/>
        </Route>
      
        
      </Routes>
    </BrowserRouter>
    
  )
  
}

export default App
