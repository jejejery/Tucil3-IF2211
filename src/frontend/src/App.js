import React from 'react';
import "./App.css"
import NavBar from './components/NavBar/NavBar';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './pages/Home/Home';
import About from './pages/About/About';
import Route_Planning from './pages/Route_Planning/Route_Planning';

  
function App() {
  return (
      <Router>
      <NavBar />
      <Routes>
        <Route path='/' active element={<Home/>} />
        <Route path='/about' element={<About/>} />
        <Route path='/route_planning' element={<Route_Planning/>} />
        
      </Routes>
    </Router>
  
    
  );
}
  
export default App;