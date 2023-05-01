import React, {useState} from 'react';
import {login} from "../../store/actions";
import "./LoginForm.scss"
import Google from "../../images/GoogleLogo.svg";
import GitHub from "../../images/GitHub.svg";
import {NavLink} from "react-router-dom";

const LoginForm = () => {
    const [email, setEmail] = useState<string>('');
    const [password, setPassword] = useState<string>('');

    return (
        <>
            <div className="login-form">
                <form className="login-form">
                    <label htmlFor="email">Адрес электронной почты</label>
                    <input onChange={e=>setEmail(e.target.value)}
                           value={email}
                           type="email"
                           name="email"
                           placeholder={"Email"}/>
                    <label htmlFor="password">Пароль</label>
                    <input onChange={e=>setPassword(e.target.value)}
                           value={password}
                           type="password"
                           name="password"
                           placeholder={"Password"}
                    />
                    <button className={"button"} onClick={ ()=>login(email,password)} >Войти</button>
                </form>
                <div className="alternative-block">
                    <div className="other-choices">
                        <img src={Google} alt=""/>
                        <img src={GitHub} alt=""/>
                    </div>
                    <div className="alternative">
                        <span>Еще нет аккаунта?   <NavLink to={'/registration'} className="selected">Зарегистрироваться</NavLink></span>

                    </div>
                </div>
            </div>

        </>
    );
};

export default LoginForm;