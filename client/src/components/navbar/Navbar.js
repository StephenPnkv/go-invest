import React from 'react';
import './Navbar.css';
import Trends from '../trends/Trends';

const Navbar = () => {


  return (
    <div>
    <nav className="navbar">
      <h1>Precisi</h1>
      <img src="https://img.icons8.com/ios-filled/50/000000/define-location--v1.png"/>
      <h1>nTrades</h1>
      <div className="links">
          <a href="/dashboard">Dashboard</a>
          <a href="/stocks">Account</a>
      </div>
    </nav>
    </div>
  );
}

export default Navbar;
