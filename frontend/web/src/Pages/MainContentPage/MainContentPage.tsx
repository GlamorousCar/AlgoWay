import * as React from "react"
import './MainContentPage.scss'
import MainImg from "../../images/main-page-img.png";
import ProgressBar from "../../componentes/ProgressBar/ProgressBar";
import { useNavigate } from "react-router-dom";



const MainContentPage = () => {

    const navigate = useNavigate();
    return (
        <div className={"main-page"}>
            <h2 className={"main-page-title"}>Онлайн-платформа для решения задач по программированию</h2>
            <div className="main-page-content">
                <div className="main-page-content-description">
                    <div className="content-description">
                        <p>Обучайся, тренируйся и подготавливайся к техническим собеседованиям, решая алгоритмические
                            задачи разного уровня сложности с автоматической проверкой решений!</p>
                        <div className="features">
                            <div className="feature">
                                <h3 className={"feature-number"}>5</h3>
                                <h3 className={"feature-text"}>языков</h3>
                            </div>
                            <div className="feature">
                                <h3 className={"feature-number"}>250</h3>
                                <h3 className={"feature-text"}>задач</h3>
                            </div>
                        </div>
                    </div>
                    <button className={"main-content-button"} onClick={()=>navigate('/themeList')}> Начать сейчас!</button>
                </div>
                <img className={'main-image'} src={MainImg} alt={"main"}/>
            </div>
            <ProgressBar/>
        </div>
    )

}

// @ts-ignore
export default MainContentPage