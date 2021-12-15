
import React, {useEffect, useState} from 'react';
import {useFetch} from '../utility/UseFetch';
import axios from 'axios';
import './trends.css';



const Trends = () => {

  const [data, setData] = useState([]);
  const [err, setErr] = useState("");
  const url = "http://localhost:8080/api/trends";

  useEffect(() => {
    fetchQuote();
  },[]);

  const fetchQuote = () =>{
    axios.get(url)
      .then(res => {
        let results = []
        res.data.finance.result['0'].quotes.forEach((item, i) => {
          results.push(item['symbol']);
        });
        setData(results);
      })
      .catch(err => console.log(err));
  }

  const items = () =>{
    return data.map((element,index) => {
      return(
        <ul
          key={index}
          >{element}</ul>
      );
    });
  }

  return (
    <div className="trend-feed">
      <div className="hwrap">
        <div className="hmove">
          {items()}
        </div>
      </div>
      <div className="trend-header">
        <img src="https://img.icons8.com/ios-glyphs/30/000000/fire-element--v1.png"/>
      </div>
    </div>
  );

}

export default Trends;
