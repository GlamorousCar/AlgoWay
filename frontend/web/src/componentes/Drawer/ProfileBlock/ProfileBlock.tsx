import * as React from "react";
import './ProfileBlock.scss'
import avatar from  '../../../images/Avatar.svg'
import dots from "../../../images/Dots.svg"

const ProfileBlock = ()=>{
    return(
        <div className={"profile-block"}>
            <div className="content">
                <div className="img-block">
                    <img src={avatar} alt=""/>
                </div>
                <div className="info-block">
                    <h6 className={'name'}>D.Tore.Y</h6>
                    <p className={'email'}>d.tore.y@yandex.ru</p>
                </div>
                <div className="show-more-block">
                    <img src={dots} alt=""/>
                </div>
            </div>
        </div>
    )
}
export default ProfileBlock