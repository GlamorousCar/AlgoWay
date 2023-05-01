import React from 'react';
import MenuIcon from "../MenuIcon/MenuIcon";
import {useDispatch, useSelector} from "react-redux";
import {IAppState} from "../../types/store";
import {drawerOpening} from "../../store/actions";
import "./BurgerButton.scss"

const BurgerButton = () => {

    const dispatch = useDispatch();
    return (
        <div className={"burger"} onClick={()=>dispatch(drawerOpening())}>
            <MenuIcon/>
        </div>
    );
};

export default BurgerButton;