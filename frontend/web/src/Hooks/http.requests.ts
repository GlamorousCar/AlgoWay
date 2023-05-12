import axios from "axios";
import {useCallback} from "react";

const UseHttpRequestHook = ( )=>{
    const postRequest = useCallback(async (url:string, data:object, headers?:object) => {
        try {
            const response = await axios.post(url, data, headers );
            if (response.statusText !== "OK") {
                throw new Error(`Couldn't fetch ${url}, status : ${response.status} `)
            }
            return response.data;
        } catch (e) {
            throw e;
        }
    }, [])


    return {postRequest}
}

export default UseHttpRequestHook;