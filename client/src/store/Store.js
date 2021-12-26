
import React from 'react';

const Store = React.createContext();
Store.displayName='Store';

export const useStore = () => React.useContext(Store);

export const StoreProvider = ({children, initialState,reducer}) => {

  const redColor = '#FF6666';
  const greenColor= '#28B637';
  const [globalState,dispatch] = React.useReducer(reducer,initialState);

  return(
    <Store.Provider
      value={[globalState,dispatch],{redColor,greenColor}}>
      {children}
    </Store.Provider>
  );
}
