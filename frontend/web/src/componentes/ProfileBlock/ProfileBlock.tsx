import * as React from "react";
import './ProfileBlock.scss'
import avatar from '../../images/Avatar.svg'
import {Link} from "react-router-dom";

const ProfileBlock = ()=>{
    return(

        <>
            <Link className={"profile-link"} to={'/profile'}>
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
            </Link>
        </>

    )
}
export default ProfileBlock