import axios from "axios";


export const _baseUrl = "https://bbaunqpcv1t23s9skmhv.containers.yandexcloud.net" ;
const $api = axios.create({

    withCredentials:true,
    baseURL:_baseUrl,

})

$api.interceptors.request.use((config) =>{
    config.headers.Authorization = `Bearer ${localStorage.getItem('token')}`;
    return config;
})

export default $api