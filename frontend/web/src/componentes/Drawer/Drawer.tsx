import * as React from "react"
import './Drawer.scss'
import ProfileBlock from "./ProfileBlock/ProfileBlock";
// import {NavLink} from "react-router-dom";
import { useEffect, useState} from "react";
import useAlgoService from "../../services/UseAlgoService";
import {IMenu} from "../../types/types";
import { NavLink } from "react-router-dom";

const Drawer = ()=>{

    const {getMenuTopics} = useAlgoService();
    const [menu, setMenu] =useState<IMenu[]>([]);

    useEffect(()=>{
        getResources();
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    const getResources = ()=>{
        getMenuTopics().then(onMenuLoaded);
    }

    const onMenuLoaded = (menuList:IMenu[])=>{
        setMenu(menuList);
    }

    // const renderItems=!error? <View menu={menu}/>:<div>Ошибка</div>;

    let  itemRefs: any[] = [];

    const setRef = (ref: any) => {
        itemRefs.push(ref);
    }
    const onFocusItem = (id:number) => {
        itemRefs.forEach(item =>item.classList.remove('active'))
        itemRefs[id].classList.add("active");
        itemRefs[id].focus();
    }

    return(
        <div className={'drawer'}>
            <ul className={"theme-list"}>
                {
                    menu.map((menuItem,index) =>{
                        return(
                            <li  className={"theme-list-item"} key={menuItem.theme_id}  >
                                <span ref={setRef}
                                      className={"theme-list-item "}
                                      onClick={()=> onFocusItem(index)}
                                >{menuItem.title}</span>
                                <ul className={'second-layer-list'}>
                                    {
                                        menuItem.algorithms.map(algorithm=>{
                                            return(
                                                <li className={'second-layer-list-item'} key={algorithm.algorithm_id}>
                                                    <NavLink to={`/topics/${algorithm.algorithm_id}`} className={"theme-list-item"}>{algorithm.title}</NavLink>
                                                </li>
                                            )
                                        })}
                                </ul>
                            </li>
                        )
                    })
                }
            </ul>
            <ProfileBlock/>
        </div>
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