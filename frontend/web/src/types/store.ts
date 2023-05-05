
import {IAlgorithm, IMenu} from "./types";


export interface IAppState {
    drawerOpeningStatus:boolean;
    menu : IMenu[];
    algorithms:IAlgorithm[];
    isAuth:boolean;
    menuLoading:boolean;
}
