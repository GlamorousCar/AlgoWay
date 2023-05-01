import {IMenu} from "../../types/types";
import AuthService from "../../services/AuthService";

export const drawerOpening = () => {
    return {
        type: 'DRAWER_OPEN'
    }
}

export const drawerClosing = () => {
    return {
        type: 'DRAWER_CLOSE'
    }
}
export const setMenuItems = (menus:IMenu[]) =>{
    return{
        type: "SET_MENU",
        payload:menus
    }
}
export const login = async (email:string, password:string)=>{
    try {
        const response = await AuthService.login(email, password);
        console.log(response)
        localStorage.setItem("token", response.data.accessToken);
        return{
            type:"SET_AUTH_TRUE"
        }
    }catch(e:any){
        console.log(e.response?.data?.message);
    }
}
export const registration = async (login:string, email:string, password:string)=>{
    try {
        const response = await AuthService.registration(login, email, password);
        console.log(response)
        localStorage.setItem("token", response.data.accessToken);
        return{
            type:"SET_AUTH_TRUE"
        }
    }catch(e:any){
        console.log(e.response?.data?.message);
    }
}
