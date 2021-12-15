import React, {Component} from "react";
import {useDispatch} from "react-redux";
import {useSpring, animated as a} from 'react-spring';
import {AiOutlineMenu, AiOutlineClose} from "react-icons/ai";
import {MdClose} from "react-icons/md";
import useWindowDimensions from './useWindowDimensions';
import { IconContext } from "react-icons";

import {
  BrowserRouter as Router,
  Route,
  Switch,
  Link
} from "react-router-dom";
import "./Navbar.css";

function shouldNavExpand(navStatus){
  if(navStatus == true)
    return false;
  return false;
}

export default function Navbar(){
  {/* navStatus = false => closed navbar */}
  const[navStatus, expandNav] = React.useState(false);
  const dispatch = useDispatch();

  const showNavbar = (showNav) => {
    dispatch({
      type: 'SHOW_NAV',
      showNav: showNav
    });
  };

    const navStyle = {
      height: "6%",
      opacity: "0.95"
    };
    const toggledNavStyle = {
      height: "100%",
      opacity: "0.95",
    };
    const fadeIn = {
      opacity: "0.95"
    };
    const fadeOut = {
      opacity: "0"
    };
    const windowDimensions = useWindowDimensions();
    const menuDimensions = (windowDimensions.width < 800) ? "1.75em" : "2.5em";

    return(


          <a.div className="navbar-main" style={navStatus ? toggledNavStyle : navStyle}>
              <div className="navbar-menu-wrap">
                <div className="icons">
                  <IconContext.Provider value={{color: "var(--green-blue-col)", size: menuDimensions}}>
                    <AiOutlineMenu style={navStatus? fadeOut : fadeIn} className="menu-icon" onClick={() => expandNav(!navStatus)} />
                  </IconContext.Provider>
                </div>
                <div className="home-link">
                  <Link onClick={() => expandNav(shouldNavExpand(navStatus))} to="/">
                    <div className="logo-link">

                    </div>
                  </Link>
                </div>
              </div>

              <IconContext.Provider value={{color: "var(--green-blue-col)", size: menuDimensions}}>
                <MdClose className="close-icon" style={navStatus? fadeIn : fadeOut} onClick={() => expandNav(!navStatus)} />
              </IconContext.Provider>
            <div className="navbar-content" >

              <ul>
                <li >
                <Link onClick={() => expandNav(!navStatus)}  to="/">HOME</Link>

                </li>
                {/*<li>
                <Link onClick={() => expandNav(!navStatus)} to="/about">ABOUT</Link>

                </li>*/}
                <li>
                  <Link onClick={() => expandNav(!navStatus)} to="/contact">CONTACT</Link>
                </li>
                <li>
                  <Link onClick={() => expandNav(!navStatus)} to="/resume">RESUME</Link>
                </li>
              </ul>
            </div>
          </a.div>
        );
}
