import React from 'react';
import {NavLink} from "react-router-dom";
import {IMenu} from "../../types/types";
import {useSelector} from "react-redux";
import {IAppState} from "../../types/store";

const ThemesList = () => {
    const {menu} = useSelector((state: IAppState) => state);
    return (
        <>
            {
                menu.map(menuItem=> {
                    return (
                        <li className={"theme-list-item"}
                            key={menuItem.theme_id}>
                            <NavLink
                                to={`/topics`}
                                className={"theme-list-item "}>
                                {
                                    ({isActive}) => {
                                       if(isActive) {
                                           return(
                                               <>
                                                   <span>{menuItem.title}</span>
                                                   <ul className={'second-layer-list'}>
                                                       {
                                                           menuItem.algorithms.map(algorithm => {
                                                               return (
                                                                   <>
                                                                       <li className={'second-layer-list-item'}
                                                                           key={algorithm.algorithm_id}>
                                                                           <NavLink
                                                                               to={`/topics/${algorithm.algorithm_id}/theory`}
                                                                               className={"theme-list-item"}>
                                                                               {algorithm.title}
                                                                           </NavLink>
                                                                       </li>
                                                                   </>
                                                               )
                                                           })}
                                                   </ul>
                                               </>
                                           );
                                       }else{
                                           return <span>{menuItem.title}</span>
                                       }
                                    }
                                }
                            </NavLink>
                        </li>
                    )
                })
            }
        </>
    );
};

export default ThemesList;