
import React, {useState, useEffect} from 'react';

export function useFetch(uri){
  const [data, setData] = useState();
  const [err, setErr] = useState();
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if(!uri) return;
    fetch(uri)
    .then(res => res.json())
    .then(setData)
    .then(() => setLoading(false))
    .catch(setErr)
  },[uri]);

  return {
    data, err, loading
  };
}
