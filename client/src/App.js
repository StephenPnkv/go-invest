import React from 'react'
import './index.css';
import Navbar from './components/navbar/Navbar';
import Quote from './components/quotes/Quote';
import Chart from './components/charts/Chart';

import Trends from './components/trends/Trends';
import {createContext} from 'react';

import {StoreProvider} from './store/Store';

function App() {
  return (
    <div className="App">
      <Navbar />

      <div className="content">
        <StoreProvider>
          <Trends />
          <Quote />
        </StoreProvider>
      </div>
    </div>
  );
}

export default App;
