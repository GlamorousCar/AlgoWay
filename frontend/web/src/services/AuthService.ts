import { AxiosResponse } from "axios";
import $api from "../Hooks/http.requests";
import {AuthResponse} from "../types/AuthResponse";

export default class AuthService{
    static async login(email:string, password:string):Promise<AxiosResponse<AuthResponse>>{
        return $api.post<AuthResponse>("/auth/login", {email, password});
    }

    static async registration(login:string, email:string, password:string):Promise<AxiosResponse<AuthResponse>>{
        return $api.post<AuthResponse>("/auth/register", {login, email, password});
    }
}