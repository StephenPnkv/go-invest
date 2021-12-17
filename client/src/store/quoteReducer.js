export const GET_QUOTE = '/api/quote?symbol='

export const initialState = {
  ticker: "",
  res: {},
};

export const getQuote = (res) =>({
  type: 'GET_QUOTE',
  res,
});

export const getQuoteReducer = (state=initialState, action) => {
  if(action.type === GET_QUOTE){
    return {
      ...state,
      ticker: state.ticker,
      res: state.res.concat(action.res),
    };
  }
};
