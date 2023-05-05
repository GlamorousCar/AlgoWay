import React, {useState} from 'react';
import "./LoginForm.scss"
import Google from "../../images/GoogleLogo.svg";
import GitHub from "../../images/GitHub.svg";
import {Navigate, NavLink, useNavigate} from "react-router-dom";
import UseAuthService from "../../services/UseAuthService";
import spinner from "../Spinners/Spinner";
import LoadingSpinner from "../Spinners/LoadingSpinner";
import {setAuthTrue} from "../../store/actions";
import {useDispatch} from "react-redux";

const LoginForm = () => {
    const [email, setEmail] = useState<string>('');
    const [password, setPassword] = useState<string>('');

    const [loading, setLoading] = useState<boolean>(false);
    const [error, setError] = useState<string>('');

    const {login} = UseAuthService();
    const dispatch = useDispatch();
    const navigate = useNavigate();

    const spinner = loading ? <LoadingSpinner/>:null;
    const errorMessage = error ? <span className={'login-error-message'}>{error}</span> :null;
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
                    <button className={"button"} onClick={ (e)=> {
                        e.preventDefault();
                        setError('');
                        setLoading(true);
                        login(email, password)
                            .then(response => {
                                localStorage.setItem("token",response.token);
                                setLoading(false);
                                dispatch(setAuthTrue());
                                navigate('/');
                            })
                            .catch(response=> {
                                console.log(response)
                                setLoading(false);
                                if(response.response.status === 400){
                                    if (response.response.data === "no rows in result set\n"){
                                        setError('Введете данные для входа')
                                    }
                                    if (response.response.data === "crypto/bcrypt: hashedPassword is not the hash of the given password\n"){
                                        setError('Введены неверный логин или пароль')
                                    }
                                }else {
                                    setError('Ошибка сервера')
                                }

                            });
                    }} >Войти</button>
                    <div className="result-block">
                        {spinner}
                        {errorMessage}
                    </div>

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