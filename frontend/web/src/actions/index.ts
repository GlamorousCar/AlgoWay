import {IMenu} from "../types/types";

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
