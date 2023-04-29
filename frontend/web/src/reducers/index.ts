import { IAppState, IAppAction } from "../types/store";
import {IMenu} from "../types/types";

const initialState = {
    drawerOpeningStatus : true,
}

const reducer = (state = initialState, action: IAppAction): IAppState  => {
    switch (action.type) {
        case 'DRAWER_OPEN':
            return {...state, drawerOpeningStatus: true}
        case 'DRAWER_CLOSE':
            return {...state, drawerOpeningStatus: false}
        default: return <IAppState>state
    }
}

export default reducer;
