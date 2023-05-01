import React, {FC, useState} from 'react';
import Img from "../../images/Registration.png"
import './RegistrationPage.scss'
import {Outlet} from "react-router-dom";
import Hand from "../../images/hand.svg"

const RegistrationPage:FC = () => {



    return (
        <div className={"registration"}>
            <img src={Img}/>
            <div className="registration-form-content">
                <div className="title-block">
                    <h4 className={"registration-title"}>Добро пожаловать</h4>
                    <img src={Hand} alt="" style={{width:"32px", height:"32px"}}/>
                </div>
                <Outlet/>
            </div>
        </div>
    );
};

export default RegistrationPage;