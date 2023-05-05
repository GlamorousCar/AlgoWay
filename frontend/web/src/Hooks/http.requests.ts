import axios from "axios";
import {useCallback} from "react";

const UseHttpRequestHook = ( )=>{
    const getRequest = useCallback(async (url:string, data:object) => {
        try {
            const response = await axios.post(url, data);
            if (response.statusText !== "OK") {
                throw new Error(`Couldn't fetch ${url}, status : ${response.status} `)
            }
            return response.data;
        } catch (e) {
            throw e;
        }
    }, [])


    return {getRequest}
}

export default UseHttpRequestHook;