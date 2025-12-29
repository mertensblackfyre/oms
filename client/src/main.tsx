import React from "react"
import ReactDOM from "react-dom/client"
import { BrowserRouter, Routes, Route } from "react-router-dom"
import Login from "./pages/Login"
import Register from "./pages/Register"
import Order from "./pages/Order"
import './index.css';

ReactDOM.createRoot(document.getElementById("root")!).render(


  <BrowserRouter>
    <Routes>
      <Route path="/login" element={<Login />} />
      <Route path="/register" element={<Register />} />
   
      <Route path="/order" element={<Order />} />
      </Routes>
  </BrowserRouter>

)

