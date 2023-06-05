import * as React from "react"
import './App.scss';
import ContentBlock from "../ContentBlock/ContentBlock";
import {BrowserRouter as Router, Route, Routes, useNavigate, Navigate} from "react-router-dom";
import RegistrationPage from "../../Pages/RegistrationPage/RegistrationPage";
import TheoryBlockPage from "../../Pages/TheoryBlockPage/TheoryBlockPage";
import MainContentPage from "../../Pages/MainContentPage/MainContentPage";
import TheoryBlock from "../TheoryBlock/TheoryBlock";
import PracticeBlock from "../PracticeBlock/PracticeBlock";
import useAlgoService from "../../services/UseAlgoService";
import {useEffect} from "react";
import {IMenu} from "../../types/types";
import {useDispatch, useSelector} from "react-redux";
import {menuLoaded, menuLoading, setMenuItems} from "../../store/actions";
import LoginForm from "../LoginForm/LoginForm";
import RegistrationForm from "../RegistrationForm/RegistrationForm";
import ProfilePage from "../../Pages/ProfilePage/ProfilePage";
import { IAppState } from "../../types/store";
import { setAuthTrue } from "../../store/actions";
import { setAuthFalse } from "../../store/actions";
import StartUpPage from "../../Pages/StartUpPage/StartUpPage";



function App() {
    const {getMenuTopics} = useAlgoService();
    const dispatch = useDispatch();
    const {isAuth} = useSelector((state: IAppState) => state);

    useEffect(() => {
        getMenus();
        localStorage.getItem('token') ? dispatch(setAuthTrue()) : dispatch(setAuthFalse());
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    const getMenus = () => {
        dispatch(menuLoading());
        getMenuTopics().then(onMenuLoaded).catch(()=>{
            setTimeout(getMenus, 10000);

        });
    }

    const onMenuLoaded = (menuList: IMenu[]) => {
        console.log("получение данных ")
        console.log(menuList)
        if(menuList.length === 0){
            setTimeout(getMenus, 10000);
        }else{
            dispatch(menuLoaded());
            dispatch(setMenuItems(menuList));
        }

    }
    return (
        <Router>
            <div className="App">
                <Routes>
                    <Route   element={ <RegistrationPage/> }>
                        <Route  path={"/login"} element={  isAuth? ( <Navigate replace to={"/"} />) :(<LoginForm/> )}/>
                        <Route  path={"/registration"} element={ isAuth? ( <Navigate replace to={"/"} />) :(<RegistrationForm/> )}/>
                    </Route>
                    <Route element={<ContentBlock/>}>
                        <Route path={"/topics/:topicId"} element={<MainContentPage/>}/>
                        <Route path={"/topics/:topicId/:algorithmId"} element={<TheoryBlockPage/>}>
                            <Route index path={'theory'} element={<TheoryBlock/>}/>
                            <Route path={'practice'} element={<PracticeBlock/>}/>
                        </Route>
                        <Route path={"/profile"} element={<ProfilePage/>}/>
                        <Route path={"/themeList"}  element={<StartUpPage/>}/>
                        <Route path={"/"} element={<MainContentPage/>}/>
                    </Route>
                </Routes>
            </div>
        </Router>
    );
}

export default App;
