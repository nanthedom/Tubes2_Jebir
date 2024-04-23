// Header.js
import React from 'react';
import { Link, useLocation } from 'react-router-dom';
import './Header.css';

function Header() {
    const location = useLocation();

    return (
        <div className="header">
            <Link to="/" className={`logo`}>Jebir Wikirace</Link>
            <div className="header-right">
                <Link to="/" className={` ${location.pathname === '/' ? 'active' : ''}`}>Home</Link>
                <Link to="/how-to-use" className={`${location.pathname === '/how-to-use' ? 'active' : ''}`}>How-To Use</Link>
                <Link to="/about" className={`${location.pathname === '/about' ? 'active' : ''}`}>About</Link>
            </div>
        </div>
    );
}

export default Header;
