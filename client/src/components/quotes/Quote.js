import React, {useEffect, useState, useReducer,createContext} from 'react';
import axios from 'axios';
import {useFecth} from '../utility/UseFetch';
import './quotes.css';
import {formatNum} from '../utility/Utility';
import Chart from '../charts/Chart';
import {useStore} from '../../store/Store';
import Options from '../options/Options';


const Quote = (props) => {
  const {redColor, greenColor} = useStore();

  const [data, setData] = useState({});
  const [siData, setSiData] = useState([]);
  const [url, setUrl] = useState("");
  const [renderTable, setRenderTable] = useState(false)
  const [colorAnimation,setColorAnimation] = useState({});

  useEffect(()=> {
    setRenderTable(false);
    setColorAnimation({animation: 'greenToWhiteBackGround 1s linear forwards'});
  },[url]);



  const handleSubmit = e => {
    e.preventDefault();
    fetchQuote(url);
  }


  const fetchQuote = () =>{
    axios.get(url)
      .then(res => {
        setData(res.data.quoteResponse.result[0]);
        axios.get('http://localhost:8080/api/si?symbol=' + res.data.quoteResponse.result[0].symbol)
          .then(siRes => {
            let d = [];
            for(let i = 0; i < siRes.data.nsdq.dataset_data.data.length; i++){
              let dataset = {};
              dataset.nsdq = siRes.data.nsdq.dataset_data.data[i];
              dataset.nyse = siRes.data.nyse.dataset_data.data[i];
              d[i] = dataset;
            }
            setSiData(d);
            setRenderTable(true);
          })
          .catch(err => console.log(err));
      })
      .catch(err => console.log(err));

  }

  const getStyle = () =>{
    switch(data.marketState){
      case "REGULAR":
        return (data.regularMarketChangePercent < 0) ? {color: redColor} : {color: greenColor};
      case "POST":
        return (data.postMarketChangePercent < 0) ? {color: redColor} : {color: greenColor};
      default:

    }

  }

  return(
    <div className="wrapper">
    <div>
      <form>
        <fieldset style={colorAnimation}>
          <label>
          <div className="search-container">
              <button
                style={colorAnimation}
                onClick={handleSubmit}
                type="submit">
                <img src="https://img.icons8.com/ios-filled/50/000000/define-location--v1.png"/>
              </button>
              <input
                style={colorAnimation}
                type="text"
                name="ticker"
                onChange={event => setUrl('http://localhost:8080/api/quote?symbol=' + event.target.value)}
                placeholder="search"/>
            </div>
          </label>
        </fieldset>
      </form>
    </div>

    <div>
      { renderTable &&
        <div className="display-name-price">
          <h1> ${data.symbol} </h1>
          <p> {data.longName} </p>
          <p style={getStyle()}>{parseFloat(data.regularMarketPrice.toPrecision(3))} ({parseFloat(data.regularMarketChangePercent.toPrecision(3))}%) Today</p>

          {
            (data.marketState === "POST") &&
            <p style={getStyle()}>{parseFloat(data.postMarketPrice.toPrecision(3))} ({parseFloat(data.postMarketChangePercent.toPrecision(3))}%) After Hours</p>
          }
        </div>
      }
      { renderTable &&
        <table>
        <thead className="quote-head">
          <tr>
            <th>Average Vol.</th>
            <th>Vol.</th>
            <th>MarketCap</th>
            <th>Dividend</th>
            <th>P/E</th>
          </tr>
        </thead>
        <tbody className="quote-body">
          <tr>
            <td>{data.averageDailyVolume10Day}</td>
            <td>{data.regularMarketVolume}</td>
            <td>{data.marketCap}</td>
            <td>{data.trailingAnnualDividendYield}</td>
            <td>{parseFloat(data.priceEpsCurrentYear.toPrecision(3))}</td>
          </tr>
        </tbody>
          <thead>
            <tr>
              <th>Open</th>
              <th>High</th>
              <th>Low</th>
              <th>52-week high</th>
              <th>52-week low</th>
            </tr>
          </thead>

        <tbody>
          <tr>
            <td>{parseFloat(data.regularMarketOpen.toPrecision(4))}</td>
            <td>{parseFloat(data.regularMarketDayHigh.toPrecision(4))}</td>
            <td>{parseFloat(data.regularMarketDayLow.toPrecision(4))}</td>
            <td>{parseFloat(data.fiftyTwoWeekHigh.toPrecision(4))}</td>
            <td>{parseFloat(data.fiftyTwoWeekLow.toPrecision(4))}</td>
          </tr>
        </tbody>
      </table>
    }

    </div>

    {renderTable &&
      <table>
      <thead >
        <tr>
          <th>Date</th>
          <th>Short Vol.</th>
          <th>Short Exempt Vol.</th>
          <th>Total Vol.</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          {siData[0].nsdq.map((ele,index) => {
            if(index == 0){
              ele = "NSDQ " + ele;
            }
            return  <td key={index}>{ele}</td>
          })}
        </tr>
        <tr>
          {siData[0].nyse.map((ele,index) => {
            if(index == 0){
              ele = "NYSE " + ele;
            }
            return  <td key={index}>{ele}</td>
          })}
        </tr>
        {
          (siData.length > 1) &&
          <tr>
          {siData[1].nsdq.map((ele,index) => {
            if(index == 0){
              ele = "NSDQ " + ele;
            }
            return  <td key={index}>{ele}</td>
          })}
          </tr>
        }
        { (siData.length > 1) &&

          <tr>
            {siData[1].nyse.map((ele,index) => {
              if(index == 0){
                ele = "NYSE " + ele;
              }
              return  <td key={index}>{ele}</td>
            })}
          </tr>
        }
      </tbody>
      </table>

      }
      <div className="chart">
      {
        renderTable &&
        <Chart symbol={data.symbol}/>
      }
      </div>
      {
        renderTable &&
        <Options symbol={data.symbol}/>
      }
    </div>
  );
}

export default Quote;
