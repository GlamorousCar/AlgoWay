import UseHttpRequestHook from "../Hooks/http.requests";

const UseAuthService = () => {

    const {postRequest} = UseHttpRequestHook();
    const _baseUrl = "https://bbaunqpcv1t23s9skmhv.containers.yandexcloud.net";

    const login = async (email: string, password: string) => {
        try {
            return await postRequest(`${_baseUrl}/auth/login`, {
                Email: email, Password: password
            })
        } catch (e) {
            console.log('Show error notification!')
            return Promise.reject(e)
        }
    }


    const registration = async (username: string, email: string, password: string) => {

        try {
            return await postRequest(`${_baseUrl}/auth/register`, {
                Login: username, Email: email, Password: password
            })
        } catch (e) {
            console.log('Show error notification!')
            return Promise.reject(e)
        }


    }

    const checkTask = async (token: string|null, language: string, taskId: number, code: string) => {

        try {
            return await postRequest(`${_baseUrl}/check_task`, {
                lang: language, source_code: code, task_id: taskId
            }, {
                headers: {
                    'user_token': token
                }
            })
        } catch (e) {
            console.log('Show error notification!')
            return Promise.reject(e)
        }


    }
    return {registration, login, checkTask}

};

export default UseAuthService;