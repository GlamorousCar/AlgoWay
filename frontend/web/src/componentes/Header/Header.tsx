import * as React from "react"
import SearchBar from "../SeadchBar/SearchBar";
import Button from "../Button/Button";
import Logo from '../../images/logo.svg';
import "./Header.scss";
import {Link, NavLink} from "react-router-dom";
import ProfileBlock from "../ProfileBlock/ProfileBlock";

const Header = () => {
    return (
        <header>
            <div className='header'>
                <Link to={"/"}><img src={Logo} alt={'logo'}/></Link>
                <SearchBar/>
                {/*<ProfileBlock/>*/}
                <NavLink to={"/registration"} className={"button"}>
                    Войти
                </NavLink>
            </div>
        </header>
    )

};

export default Header;