import { IAppState} from "../types/store";
import {IMenu} from "../types/types";

const initialState = {
    drawerOpeningStatus : false,
    menu:[],
    algorithms:[],
}

const reducer = (state:IAppState = initialState,action:any): IAppState  => {
    switch (action.type) {
        case 'DRAWER_OPEN':
            return {...state, drawerOpeningStatus: true}
        case 'DRAWER_CLOSE':
            return {...state, drawerOpeningStatus: false}
        case "SET_MENU":
            return {...state, menu:action.payload }
        default: return <IAppState>state
    }
}

export default reducer;
