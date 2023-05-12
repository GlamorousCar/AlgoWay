import * as React from "react"
import './TheoryBlockPage.scss'
import green_book from "../../images/Vector_book_green.svg";
import default_book from "../../images/Vector_book_default.svg";
import default_code from "../../images/Vector_code_default.svg"
import green_code from "../../images/Vector_code_green.svg"
import {NavLink, Outlet, useParams} from "react-router-dom";
import useAlgoService from "../../services/UseAlgoService";
import {IAlgorithm} from "../../types/types";
import { useEffect, useState} from "react";
import {Skeleton} from "@mui/material";
import {useSelector} from "react-redux";
import {IAppState} from "../../types/store";
import {CSSProperties} from "@mui/material/styles/createMixins";

const TheoryBlockPage = () => {
    const {algorithmId} = useParams();

    const {getAlgorithmTheory} = useAlgoService();
    const [algorithm, setAlgorithm] = useState<IAlgorithm>();
    const [loading, setLoading] = useState<boolean>(true)

    const {isAuth} = useSelector((state: IAppState) => state)


    useEffect(() => {
        getResources(Number(algorithmId));
        setLoading(true);
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [algorithmId]);

    const getResources = (algorithmId: number) => {
        getAlgorithmTheory(algorithmId)
            .then(onAlgorithmLoaded)
    }


    const onAlgorithmLoaded = (algorithm: IAlgorithm) => {
        setAlgorithm(algorithm);
        setLoading(false)
    }

    const disableStyle = !isAuth ? "disable" :'';
    return (
        <div className={"algorithm"}>
            <div className="container">
                <h3 className={"algorithm-title"}>{loading ?
                    <Skeleton width={210} sx={{bgcolor: '#272A2D',}}/> : algorithm?.title}</h3>

                <nav className={"nav-switch"}>
                    <NavLink to={"theory"} style={{textDecoration: "none", }} end >
                        {({isActive}) => {
                            let className = isActive ? " active" : ""
                            return (

                                <div className={`nav-item theory ${className}`}>
                                    <img src={isActive ? green_book : default_book} alt=""/>
                                    <span className={"nav-item-text"}> Теория</span>
                                </div>
                            )
                        }}

                    </NavLink>
                    <NavLink to={"practice"} className={disableStyle}>
                        {({isActive, isPending}) => {
                            let className = isActive ? " active" : ""
                            return (
                                <div className={`nav-item practice  ${className}`}>
                                    <img src={isActive ? green_code : default_code} alt=""/>
                                    <span className={"nav-item-text"}>Практика</span>
                                </div>
                            )
                        }}
                    </NavLink>
                </nav>

                <div className="content">
                    <Outlet/>
                </div>
            </div>
        </div>
    );
}


export default TheoryBlockPage