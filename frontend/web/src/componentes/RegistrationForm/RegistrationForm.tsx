import React, {useState} from 'react';
import { registration} from "../../store/actions";
import Google from "../../images/GoogleLogo.svg";
import GitHub from "../../images/GitHub.svg";
import {NavLink} from "react-router-dom";
import './RegistrationForm.scss'

const RegistrationForm = () => {
    const [email, setEmail] = useState<string>('');
    const [username, setUsername] = useState<string>('');
    const [password, setPassword] = useState<string>('');

    return (
        <div className="registration-form">
            <form className="registration-form">
                <label htmlFor="email">Логин</label>
                <input type="text"
                       onChange={e => setUsername(e.target.value)}
                       value={username}
                       name="login"
                       placeholder={"Login"}/>
                <label htmlFor="email">Адрес электронной почты</label>
                <input onChange={e => setEmail(e.target.value)}
                       value={email}
                       type="email"
                       name="email"
                       placeholder={"Email"}/>
                <label htmlFor="password">Пароль</label>
                <input onChange={e => setPassword(e.target.value)}
                       value={password}
                       type="password"
                       name="password"
                       placeholder={"Password"}
                />
                <button className={"button"} onClick={() => registration(username, email, password)}>
                    Создать аккаунт
                </button>
            </form>
            <div className="alternative-block">
            <div className="other-choices">
                <img src={Google} alt=""/>
                <img src={GitHub} alt=""/>
            </div>
            <div className="alternative">
                <p>Уже есть аккаунта?   <NavLink to={'/login'} className="selected">Войти</NavLink></p>

            </div>
        </div>

        </div>
    )
};

export default RegistrationForm;