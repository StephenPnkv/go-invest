import Big from 'big.js';


export const formatNum = (num, precision) => {
  let x = new Big(num)
  return x.toPrecision(precision)
}
