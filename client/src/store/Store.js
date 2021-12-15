
import React from 'react';

const Store = React.createContext();
Store.displayName='Store';

export const useStore = () => React.useContext(Store);

export const StoreProvider = ({children}) => {

  const redColor = '#FF6666';
  const greenColor= '#28B637';


  return(
    <Store.Provider value={{redColor,greenColor}}>
      {children}
    </Store.Provider>
  );
}
