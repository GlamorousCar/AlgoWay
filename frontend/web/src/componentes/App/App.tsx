import * as React from "react"
import './App.scss';
import ContentBlock from "../ContentBlock/ContentBlock";
import {BrowserRouter as Router, Route, Routes} from "react-router-dom";
import RegistrationPage from "../../Pages/RegistrationPage/RegistrationPage";
import TheoryBlockPage from "../../Pages/TheoryBlockPage/TheoryBlockPage";
import MainContentPage from "../../Pages/MainContentPage/MainContentPage";
import TheoryBlock from "../TheoryBlock/TheoryBlock";
import PracticeBlock from "../PracticeBlock/PracticeBlock";
import useAlgoService from "../../services/UseAlgoService";
import {useEffect, useState} from "react";
import {IMenu} from "../../types/types";
import {useDispatch, useSelector} from "react-redux";
import {IAppState} from "../../types/store";
import {setMenuItems} from "../../actions";



function App() {
    const {getMenuTopics} = useAlgoService();
    const dispatch = useDispatch();


    useEffect(() => {
        getResources();
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    const getResources = () => {
        getMenuTopics().then(onMenuLoaded);
    }

    const onMenuLoaded = (menuList: IMenu[]) => {
        dispatch(setMenuItems(menuList));
    }
    return (
        <Router>
            <div className="App">
                <Routes>
                    <Route  path={"/registration"} element={<RegistrationPage/>}/>
                    <Route element={<ContentBlock/>}>
                        <Route path={"/topics"} element={<TheoryBlockPage/>}/>
                        <Route path={"/topics/:algorithmId"} element={<TheoryBlockPage/>}>
                            <Route index path={'theory'} element={<TheoryBlock/>}/>
                            <Route path={'practice'} element={<PracticeBlock/>}/>
                        </Route>
                        <Route path={"/"} element={<MainContentPage/>}/>
                    </Route>
                </Routes>
            </div>
        </Router>
    );
}

export default App;
