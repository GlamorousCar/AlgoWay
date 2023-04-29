import * as React from "react"
import './ContentBlock.scss'
import {Outlet} from "react-router-dom";
import Header from "../Header/Header";
import Drawer from "../Drawer/Drawer";
import Footer from "../Footer/Footer";
import { useSelector} from "react-redux";
import {IAppState} from "../../types/store";
import BurgerButton from "../BurgerButton/BurgerButton";

const ContentBlock = ()=>{

    const drawerStatus = useSelector((state:IAppState) =>state.drawerOpeningStatus);

    return(
            <div className={"main-content-page"}>
                <Header/>
                {drawerStatus?<Drawer />:<BurgerButton/>}
                <div style={{paddingLeft:drawerStatus?"15%":'5%'}} className="main">
                    <Outlet/>
                </div>
                <Footer/>
            </div>
    )
}

export default ContentBlock