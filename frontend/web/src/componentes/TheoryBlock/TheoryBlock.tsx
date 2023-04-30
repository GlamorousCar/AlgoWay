import React, {useEffect, useState} from 'react';
import {useParams} from "react-router-dom";
import useAlgoService from "../../services/UseAlgoService";
import {IAlgorithm} from "../../types/types";

const TheoryBlock = () => {
    const {algorithmId} = useParams();

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
        <div>
            Теория
            <p>{algorithm?.content}</p>
        </div>
    );
};

export default TheoryBlock;