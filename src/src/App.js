import './App.css';
import React from 'react';
import {BrowserRouter as Router, Routes, Route} from 'react-router-dom';
import Home from './Home';
import About from './components/About';
import HowToUse from './components/HowToUse';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/about" element={<About />} />
        <Route path="/how-to-use" element={<HowToUse />} />
      </Routes>
    </Router>
    // <div>
    //   <Home />
    // </div>
  );
}

export default App;
