import React from 'react';
import Img from "../../images/Registration.png"
import './RegistrationPage.scss'

const RegistrationPage = () => {
    return (
        <div className={"registration"}>
            <img src={Img}/>
            <div className="registration-form-content">
                <h4>Добро пожаловать</h4>
                <div className="registration-form">
                    <label htmlFor="email">Логин</label>
                    <input type="text" name="login" />
                    <label htmlFor="email">Адрес электронной почты</label>
                    <input type="email" name="email" />
                    <label htmlFor="email">Пароль</label>
                    <input type="password" name="password" />
                    <button>Создать аккаунт</button>
                </div>
            </div>
        </div>
    );
};

export default RegistrationPage;