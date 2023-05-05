import {IMenu} from "../../types/types";

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
export const menuLoading = ()=>{
    return{
        type:"MENU_FETCHING"
    }
}
export const menuLoaded = ()=>{
    return{
        type:"MENU_FETCHED"
    }
}
export const setAuthTrue = () =>{
    return{
        type:"SET_AUTH_TRUE"
    }
}
export const setAuthFalse = () =>{
    return{
        type:"SET_AUTH_FALSE"
    }
}
