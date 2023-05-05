import UseHttpRequestHook from "../Hooks/http.requests";

const UseAuthService = () => {

    const {getRequest}= UseHttpRequestHook();
    const _baseUrl = "https://bbaunqpcv1t23s9skmhv.containers.yandexcloud.net";

    const login = async (email:string, password:string) =>{
        try{
            return await getRequest(`${_baseUrl}/auth/login`, {
                Email: email, Password: password
            })
        }catch (e){
            console.log('Show error notification!')
            return Promise.reject(e)
        }
    }


    const registration = async (username:string, email:string, password:string) =>{

        try{
            return await getRequest(`${_baseUrl}/auth/register`, {
                Login: username, Email: email, Password: password
            })
        }catch (e){
            console.log('Show error notification!')
            return Promise.reject(e)
        }


    }
    return{registration,login}

};

export default UseAuthService;