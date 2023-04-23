import * as React from "react"
import './App.scss';
import Header from "./componentes/Header/Header";
import Drawer from "./componentes/Drawer/Drawer";
import ContentBlock from "./componentes/ContentBlock/ContentBlock";
import {BrowserRouter as Router} from "react-router-dom";
import Footer from "./componentes/Footer/Footer";


function App() {
    return (
        <Router>
            <div className="App">
                <Header/>
                <Drawer/>
                <div className="main">
                    <ContentBlock/>
                </div>
                <Footer/>
            </div>
        </Router>
    );
}

export default App;
