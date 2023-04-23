import React from 'react';
import Logo from "../../images/logo.svg";
import "./Footer.scss"

const Footer = () => {
    return (
        <footer>
            <div className="footer">
                <img src={Logo} alt={'logo'}/>
                <span>© 2023 algoway.</span>
            </div>
        </footer>
    );
};

export default Footer;