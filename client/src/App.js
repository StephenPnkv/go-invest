import * as React from "react";
import { Routes, Route, Link } from "react-router-dom";
import './index.css';
import Navbar from './components/navbar/Navbar';
import Quote from './components/quotes/Quote';
import Chart from './components/charts/Chart';
import Trends from './components/trends/Trends';
import {createContext} from 'react';
import {initialState, getQuoteReducer} from './store/quoteReducer.js';

import {StoreProvider} from './store/Store';

function App() {
  return (
    <div className="App">
      <Navbar />
      <Trends />
      <div className="content">
        <StoreProvider initialState={initialState} reducer={getQuoteReducer}>
          <Routes>
            <Route path="/stocks" element={<Quote />} />
          </Routes>

        </StoreProvider>
      </div>
    </div>
  );
}

export default App;
