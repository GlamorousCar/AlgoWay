import { IAppState} from "../../types/store";
import {IMenu} from "../../types/types";

const initialState = {
    drawerOpeningStatus : false,
    menu:[],
    algorithms:[],
    isAuth:false,
    menuLoading:true
}

const reducer = (state:IAppState = initialState,action:any): IAppState  => {
    switch (action.type) {
        case 'DRAWER_OPEN':
            return {...state, drawerOpeningStatus: true}
        case 'DRAWER_CLOSE':
            return {...state, drawerOpeningStatus: false}
        case "SET_MENU":
            return {...state, menu:action.payload }
        case "MENU_FETCHING":
            return {...state, menuLoading:true }
        case "MENU_FETCHED":
            return {...state, menuLoading:false }
        case "SET_AUTH_TRUE":
            return {...state, isAuth: true}
        case "SET_AUTH_FALSE":
            return {...state, isAuth: false}
        default: return <IAppState>state
    }
}

export default reducer;
