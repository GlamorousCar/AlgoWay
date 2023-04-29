import * as React from "react"
import './TheoryBlockPage.scss'
import green_book from "../../images/Vector_book_green.svg";
import grey_code from "../../images/Vector_code_default.svg"
import {NavLink, useParams} from "react-router-dom";
import useAlgoService from "../../services/UseAlgoService";
import {IAlgorithm} from "../../types/types";
import {useEffect, useState} from "react";

const TheoryBlockPage = () => {
    const {algorithmId} = useParams();
    console.log(algorithmId)


    const {getAlgorithmTheory} = useAlgoService();
    const [algorithm, setAlgorithm] = useState<IAlgorithm>();


    useEffect(() => {
        getResources(algorithmId);
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [algorithmId]);

    const getResources = (algorithmId: string | undefined) => {
        getAlgorithmTheory(algorithmId)
            .then(onAlgorithmLoaded)
    }


    const onAlgorithmLoaded = (algorithm: IAlgorithm) => {
        setAlgorithm(algorithm);
    }
    return (
        <div className={"theory"}>
            <div className="container">
                <h3 className={"theory-title"}>{algorithm?.title ? algorithm.title : "В скором времени появится заголовок"}</h3>
                <nav>
                    <ul className={"nav-switch"}>
                        <li className={"nav-item theory active"}>

                            <img src={green_book} alt=""/>
                            <span>Теория</span>
                        </li>
                        <li className={"nav-item practice "}>
                            <img src={grey_code} alt=""/>
                            <span>Практика</span>
                        </li>

                    </ul>
                </nav>

                <div className="content">
                    <p>{algorithm?.content}</p>

                </div>
            </div>
        </div>
    );
}


export default TheoryBlockPage