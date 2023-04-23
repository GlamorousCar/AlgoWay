import * as React from "react"
import './MainContentPage.scss'
import MainImg from "../../images/main-page-img.png";



const MainContentPage = ()=>{

return(
    <div className={"main-page"}>
        <h2 className={"main-page-title"}  >Онлайн-платформа для решения задач по программированию</h2>
        <div className="main-page-content">
            <div className="main-page-content-description">
                <p>Обучайся, тренируйся и подготавливайся к техническим собеседованиям, решая алгоритмические задачи на 5 языках программирования с автоматической проверкой решений и разными уровнями сложности.</p>
                <button className={"main-content-button"}> Начать сейчас! </button>
            </div>
            <img src={MainImg} alt={"main"}/>
        </div>
    </div>

)

}

export default MainContentPage