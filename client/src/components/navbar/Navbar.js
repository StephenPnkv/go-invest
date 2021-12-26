import React from 'react';
import './Navbar.css';
import Trends from '../trends/Trends';
import { Routes, Route, Link } from "react-router-dom";

const Navbar = () => {

  return (
    <div>
    <nav className="navbar">
      <h1>Precisi</h1>
      <img src="https://img.icons8.com/ios-filled/50/000000/define-location--v1.png"/>
      <h1>n</h1>
      <div className="links">
          <Link to="/stocks">Stocks</Link>
          <Link to="/account">Account</Link>
      </div>
    </nav>
    </div>
  );
}

export default Navbar;
