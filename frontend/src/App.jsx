import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import Nav from "./components/nav";
import './App.css'
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { Login, Register, List, Home, Detail, Add } from "./pages";


function App() {
  const [count, setCount] = useState(0)

  return (
    <BrowserRouter>
      
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/article/list" element={<List />} />
        <Route path="/article/detail/:id" element={<Detail />} />
        <Route path="/article/add" element={<Add />} />
      </Routes>
    </BrowserRouter>

  
  )
}

export default App
