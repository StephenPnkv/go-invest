import React, {useEffect, useState, useReducer,createContext} from 'react';
import axios from 'axios';
import {useFecth} from '../utility/UseFetch';
import './options.css';


const Options = (props) => {

  const [data, setData] = useState({});
  const [renderTable, setRenderTable] = useState(false);

  useEffect(()=> {
    fetchQuote();
  },[]);

/*
ask: 2.6
bid: 1
change: 0
contractSize: "REGULAR"
contractSymbol: "PROG211223C00000500"
currency: "USD"
expiration: 1640217600
impliedVolatility: 30.250000546875
inTheMoney: true
lastPrice: 1.75
lastTradeDate: 1640205412
openInterest: 11
strike: 0.5
volume: 3
*/

const getHeader = () => {
    return (
      <tr>
        <th>Strike</th>
        <th>Last</th>
        <th>Change</th>
        <th>Bid</th>
        <th>Ask</th>
        <th>Volume</th>
        <th>Open Int.</th>
        <th>Imp. Vol.</th>
      </tr>
  );
}

const getCalls = () => {
  return data.calls.map((obj,index) => {
    return (<tr key={index}>
      <td>{obj.strike}</td>
      <td>{obj.lastPrice}</td>
      <td>{obj.change}</td>
      <td>{obj.bid}</td>
      <td>{obj.ask}</td>
      <td>{obj.volume}</td>
      <td>{obj.openInterest}</td>
      <td>{obj.impliedVolatility}</td>
    </tr>);
  })
}

const getPuts = () => {
  return data.puts.map((obj,index) => {
    return (<tr key={index}>
      <td>{obj.strike}</td>
      <td>{obj.lastPrice}</td>
      <td>{obj.change}</td>
      <td>{obj.bid}</td>
      <td>{obj.ask}</td>
      <td>{obj.volume}</td>
      <td>{obj.openInterest}</td>
      <td>{obj.impliedVolatility}</td>
    </tr>);
  })
}

  const fetchQuote = () =>{
    axios.get('http://localhost:8080/api/options?symbol=' + props.symbol)
      .then(res => {
        let optionsData = {};
        optionsData.calls = res.data.optionChain.result[0].options[0].calls;
        optionsData.puts = res.data.optionChain.result[0].options[0].puts;
        setData(optionsData);
        setRenderTable(true);
      })
      .catch(err => console.log(err));
  }


  return(
      <div className="options-wrapper">
      <h1>{props.symbol ? (props.symbol + " Calls") : "Not Found" }</h1>

          <table>
          <thead className="quote-head">
            {
              renderTable && getHeader()
            }
          </thead>
          <tbody className="quote-body">
              {
                renderTable && getCalls()
              }
          </tbody>
            <thead>
            {
              renderTable && getHeader()
            }
            </thead>

          <tbody>
            {
              renderTable && getCalls()
            }
          </tbody>
        </table>

      </div>
  );
}

export default Options;
