import React from 'react';
import EventProgress from "../EventProgress/EventProgress";
import "./ProgressBar.scss"
import {useSelector} from "react-redux";
import {IAppState} from "../../types/store";
import {containerClasses} from "@mui/material";
import Button from "../Button/Button";

const ProgressBar = () => {
    const {isAuth} = useSelector((state: IAppState) => state);
    const blurStyle = !isAuth ? " progress-container-blur" : "";
    return (
        <div className="progress-container-box">
            <div className={blurStyle}>
                <div className={`progress-container `}>
                    <h3 className={"progress-title"}> Прогресс изучения</h3>
                    <div className="progress-events">
                        <EventProgress total={25} done={13} description={"теорий прочитано"}/>
                        <EventProgress total={25} done={1} description={"задач выполнено"}/>
                    </div>
                </div>
            </div>
            {!isAuth
                ?
                <div className="hidden-block">
                    <p>Прогресс сохраняется только у авторизированных пользователей</p>
                    <Button text={"Создать аккаунт"}/>
                </div>
                :
                null
            }
        </div>



    );
};

export default ProgressBar;