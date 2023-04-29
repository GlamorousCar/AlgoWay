import * as React from "react"
import './App.scss';
import ContentBlock from "../ContentBlock/ContentBlock";
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import RegistrationPage from "../../Pages/RegistrationPage/RegistrationPage";
import TheoryBlockPage from "../../Pages/TheoryBlockPage/TheoryBlockPage";
import MainContentPage from "../../Pages/MainContentPage/MainContentPage";



function App() {
    return (
        <Router>
            <div className="App">
                <Routes>
                    <Route  path={"/registration"} element={<RegistrationPage/>}/>
                    <Route element={<ContentBlock/>}>
                        <Route path={"/topics/:algorithmId"} element={<TheoryBlockPage/>}>
                            <Route />
                        </Route>
                        <Route path={"/"} element={<MainContentPage/>}/>
                    </Route>
                </Routes>
            </div>
        </Router>
    );
}

export default App;
