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
import {useEffect} from "react";
import {IMenu} from "../../types/types";
import {useDispatch} from "react-redux";
import {setMenuItems} from "../../store/actions";
import LoginForm from "../LoginForm/LoginForm";
import RegistrationForm from "../RegistrationForm/RegistrationForm";



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
        console.log("получение данных ")
        console.log(menuList)
        if(menuList.length === 0){
            getResources();
        }
        dispatch(setMenuItems(menuList));
    }
    return (
        <Router>
            <div className="App">
                <Routes>
                    <Route   element={<RegistrationPage/>}>
                        <Route  path={"/login"} element={<LoginForm/>}/>
                        <Route  path={"/registration"} element={<RegistrationForm/>}/>
                    </Route>
                    <Route element={<ContentBlock/>}>
                        <Route path={"/topics"} element={<MainContentPage/>}/>
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
