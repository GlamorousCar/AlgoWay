import * as React from "react"
import './App.scss';
import Header from "./componentes/Header/Header";
import Drawer from "./componentes/Drawer/Drawer";
import ContentBlock from "./componentes/ContentBlock/ContentBlock";
import {BrowserRouter as Router} from "react-router-dom";


function App() {
    return (
        <Router>
            <div className="App">
                <Header/>
                <div className="main">
                    <Drawer/>
                    <ContentBlock/>
                </div>
            </div>
        </Router>
    );
}

export default App;
