import * as React from "react"
import './App.scss';
import Header from "./componentes/Header/Header";
import Drawer from "./componentes/Drawer/Drawer";
import TheoryBlock from "./componentes/TheoryBlock/TheoryBlock";

function App() {
  return (
    <div className="App">
        <Header/>
        <div className="main">
            <Drawer/>
            <TheoryBlock/>
        </div>


    </div>
  );
}

export default App;
