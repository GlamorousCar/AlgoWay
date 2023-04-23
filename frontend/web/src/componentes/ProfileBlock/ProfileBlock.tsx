import * as React from "react";
import './ProfileBlock.scss'
import avatar from '../../images/Avatar.svg'

const ProfileBlock = ()=>{
    return(
        <div className={"profile-block"}>
            <div className="content">
                <div className="info-block">
                    <h6 className={'name'}>D.Tore.Y</h6>
                </div>
                <div className="img-block">
                    <img src={avatar} alt=""/>
                </div>

            </div>
        </div>
    )
}
export default ProfileBlock