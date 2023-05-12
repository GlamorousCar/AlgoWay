import * as React from "react"
import './Drawer.scss'
import Exit from '../../images/Exit-icon.svg';
import {Transition} from "react-transition-group";
import {useDispatch, useSelector} from "react-redux";
import {drawerClosing} from "../../store/actions";
import {IAppState} from "../../types/store";
import ThemesList from "../ThemesList/ThemesList";
import {Skeleton} from "@mui/material";
import SkeletonList from "../skeletonList/SkeletonList";

const Drawer = () => {
    const {drawerOpeningStatus, menuLoading } = useSelector((state: IAppState) => state);
    const dispatch = useDispatch();

    const duration = 30000;
    const defaultStyles = {
        transition: `opacity ${duration}ms ease-in`,
        opacity: 0
    }

    const transitionStyles = {
        entering: {opacity: 1},
        entered: {opacity: 1},
        exiting: {opacity: 0},
        exited: {opacity: 0}
    };

    function onEnterHandler() {
        console.log("enter")
    }

    return (
        <Transition in={drawerOpeningStatus} timeout={duration} unmountOnExit onEnter={onEnterHandler}>
            {state => (
                <div className={'drawer'}
                     style={{...defaultStyles, ...transitionStyles[state as keyof typeof transitionStyles]}}>
                    <div className="exit-icon-block">
                        <img onClick={() => dispatch(drawerClosing())} src={Exit} alt="exit icon"/>
                    </div>
                    <h4 className={"exit-icon-title"}>Алгоритмы</h4>
                    <ul className={"theme-list"}>
                        {menuLoading?<SkeletonList/>:<ThemesList />}
                    </ul>
                </div>
            )}
        </Transition>

    )
}
export default Drawer