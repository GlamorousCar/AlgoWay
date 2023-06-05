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
// import {GoogleLogin} from "@react-oauth/google";

const LoginForm = () => {
    const [email, setEmail] = useState<string>('');
    const [password, setPassword] = useState<string>('');

    const [loading, setLoading] = useState<boolean>(false);
    const [error, setError] = useState<string>('');

    const {login} = UseAuthService();
    const dispatch = useDispatch();
    const navigate = useNavigate();


    // const responseMessage = (response:any) => {
    //     console.log(response);
    // };
    // const errorGoogleMessage = () => {
    //     console.log("error");
    // };

    const onClickHandler=(e:any)=>{
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
                console.log(response);
                setLoading(false);



                if(response.response?.status === 400){
                    if (response.response.data === "no rows in result set\n"){
                        setError('Введете данные для входа')
                    }
                    if (response.response.data === "crypto/bcrypt: hashedPassword is not the hash of the given password\n"){
                        setError('Введены неверный логин или пароль')
                    }
                    if (response.response.data === "conn closed\n"){
                        setError('Ошибка сервера. Попробуйте войти позже')
                    }
                    if (response.response.data === "write failed: write tcp 172.16.18.131:50428->158.160.66.146:6432: write: connection reset by peer\n"){
                        setError('Ошибка сервера. Попробуйте войти позже')
                    }
                }else {
                    setError('Ошибка сервера')
                }

            });
    }
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
                        onClickHandler(e)
                    }}>Войти</button>
                    <div className="result-block">
                        {spinner}
                        {errorMessage}
                    </div>

                </form>
                <div className="alternative-block">
                    <div className="other-choice">
                        <img src={Google} alt=""/>
                        {/* <GoogleLogin onSuccess={responseMessage} onError={errorGoogleMessage} /> */}
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