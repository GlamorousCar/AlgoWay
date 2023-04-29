import { Action } from "redux";
import {IMenu} from "./types";


export interface IAppState {
    drawerOpeningStatus:boolean;
}

export interface IAppAction extends Action {
    payload?: Partial<IAppState>;
}