import * as React from "react"
import SearchBar from "../SeadchBar/SearchBar";
import Button from "../Button/Button";
import Logo from '../../images/Logotype.svg';
import "./Header.scss";
import {Link} from "react-router-dom";

const Header = () => {
    return (
        <header>
            <div className='header'>
                <Link to={"/"}><img src={Logo} alt={'logo'}/></Link>
                <SearchBar/>
                <Button text={"Выйти"}/>
            </div>
        </header>
    )

};

export default Header;