import React, {useEffect, useState} from 'react';
import ReactDOM from 'react-dom';
import { VictoryBar, VictoryChart, VictoryAxis, VictoryTheme,VictoryLine,VictoryCandlestick   } from 'victory';
import axios from 'axios';
import {formatNum} from '../utility/Utility';
import './chart.css';

const Chart = (props) => {

  const [chartData, setChartData] = useState([]);
  const [renderChart, setRenderChart] = useState(false);
  const [url, setUrl] = useState("");

  useEffect(() => {
    setUrl("http://localhost:8080/api/chart?symbol=" + props.symbol)
    fetchQuote(url);
  },[url]);

  const fetchQuote = () =>{
    if(!url) return;
    axios.get(url)
    .then(res => {
      let data = [];
      for(let i = 0; i < res.data.chart.result[0].timestamp.length; i++){
        let obj = {};
        obj.x = new Date(res.data.chart.result[0].timestamp[i] * 1000);
        obj.open = parseFloat(res.data.chart.result[0].indicators.quote[0].open[i].toPrecision(4));
        obj.high = parseFloat(res.data.chart.result[0].indicators.quote[0].high[i].toPrecision(4));
        obj.low = parseFloat(res.data.chart.result[0].indicators.quote[0].low[i].toPrecision(4));
        obj.close = parseFloat(res.data.chart.result[0].indicators.adjclose[0].adjclose[i].toPrecision(4));
        data[i] = obj;
      }
      setChartData(data);
      setRenderChart(true);
    })
    .catch(err => console.log(err));
  }

  return (
    <div className="chart-wrapper">
    {
      renderChart &&

        <VictoryChart
          theme={VictoryTheme.material}
          domainPadding={10}
          scale={{ x: "time" }}
          origin={{ x: 0, y: 0 }}
          >

        <VictoryCandlestick
          candleColors={{ positive: "#28B637", negative: "#c43a31" }}
          data={chartData}
        />

        </VictoryChart>
    }
    </div>
  );
}

export default Chart;
