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

// interface menuProps{
//     menu:IMenu[]
// }
// export const View:FC<menuProps> = ({menu})=>{
//
//     const showAlgorithmsList = (menuItem: IMenu) =>{
//         return(
//             <ul className={'second-layer-list'}>
//                 {
//                     menuItem.algorithms.map(algorithm=>{
//                         return(
//                             <li className={'second-layer-list-item'} key={algorithm.algorithm_id}>
//                                 <NavLink to={`/topics/${algorithm.theme_id}/${algorithm.algorithm_id}`} className={"theme-list-item"}>{algorithm.title}</NavLink>
//                             </li>
//                         )
//                 })}
//             </ul>
//         )
//     }
//
//     return (
//         <>
//             {showAlgorithmsList}
//         </>
//     )
//
// }

// <li className={"theme-list-item"}>
//     <NavLink to={"/topics/123"} className={"theme-list-item"}>Деревья поиска</NavLink>
//     <ul className={'second-layer-list'}>
//         <li className={'second-layer-list-item'}><span>Деревья в STL</span></li>
//         <li className={'second-layer-list-item active'}><span>Декартово дерево</span></li>
//
//     </ul>
// </li>


// <ul className={"theme-list"}>
//     {
//         menu.map((menuItem,index) =>{
//             return(
//                 <li  className={"theme-list-item"} key={menuItem.theme_id}  >
//                                 <span ref={setRef}
//                                       className={"theme-list-item "}
//                                       onClick={()=> onFocusItem(index)}
//                                 >{menuItem.title}</span>
//                     <ul className={'second-layer-list'}>
//                         {
//                             menuItem.algorithms.map(algorithm=>{
//                                 return(
//                                     <li className={'second-layer-list-item'} key={algorithm.algorithm_id}>
//                                         <NavLink to={`/topics/${algorithm.algorithm_id}`} className={"theme-list-item"}>{algorithm.title}</NavLink>
//                                     </li>
//                                 )
//                             })}
//                     </ul>
//                 </li>
//             )
//         })
//     }
// </ul>