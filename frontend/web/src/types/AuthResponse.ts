import {IUser} from "./types";

export interface AuthResponse{
    accessToken: string;
    refreshToken:string;
    user:IUser;
}