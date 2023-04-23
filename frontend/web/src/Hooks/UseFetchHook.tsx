import {useCallback, useState} from "react";
import axios from "axios";

const UseFetchHook = () => {

    const [error, setError] = useState('');

    const getRequest = useCallback(async (url: string) => {
        try {
            const response = await axios.get(url);
            console.log(response);
            if (response.statusText !== "OK") {
                throw new Error(`Couldn't fetch ${url}, status : ${response.status} `)
            }
            return response.data;
        } catch (e) {
            setError("Error message");
            throw e;
        }
    }, [])


    const clearError = useCallback(()=>{
        setError("")
    }, [])

    return {getRequest, error, clearError}
};

export default UseFetchHook;