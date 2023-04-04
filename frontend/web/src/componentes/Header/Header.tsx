import * as React from "react"
import SearchBar from "../SeadchBar/SearchBar";
import Button from "../Button/Button";
import Logo from '../../images/Logotype.svg';
import "./Header.scss";

const Header = ()=> {
    return(
        <header>
            <div className='header'>
                <img src={Logo} alt={'logo'}/>
                <SearchBar/>
                <Button text={"Выйти"}/>
            </div>
        </header>
    )

};

export default Header ;