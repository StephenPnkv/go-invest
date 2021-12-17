import React, {useEffect, useState, useReducer,createContext} from 'react';
import axios from 'axios';
import {useFecth} from '../utility/UseFetch';
import './options.css';


const Options = (props) => {

  const [data, setData] = useState({});
  const [url, setUrl] = useState("");
  const [renderTable, setRenderTable] = useState(false)

  useEffect(()=> {
    fetchQuote();
  },[]);



  const fetchQuote = () =>{
    if(props.symbol === null) return
    axios.get('http://localhost:8080/api/options?symbol=',props.symbol)
      .then(res => {
        setData(res.data.optionChain.result[0]);
      })
      .catch(err => console.log(err));
      console.log(data);
      setRenderTable(true);
  }


  return(
    <div className="wrapper">


    <div>

      { renderTable &&
        <h1>Options Table</h1>
        &&
        <table>
        <thead className="quote-head">
          <tr>
            <th> </th>
          </tr>
        </thead>
        <tbody className="quote-body">
          <tr>
            <td> </td>
          </tr>
        </tbody>
          <thead>
            <tr>
              <th> </th>
            </tr>
          </thead>

        <tbody>
          <tr>
            <td> </td>
          </tr>
        </tbody>
      </table>
    }

    </div>

    {renderTable &&
      <table>
      <thead >
        <tr>
          <th></th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td></td>
        </tr>
        <tr>
          <td></td>
        </tr>
        <tr>
          <td></td>
        </tr>
        <tr>

          <td></td>
        </tr>
      </tbody>
      </table>


    }
    </div>
  );
}

export default Options;
