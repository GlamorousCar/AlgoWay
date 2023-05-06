import React, {useState} from 'react';
import {NavLink} from "react-router-dom";
import {useSelector} from "react-redux";
import {IAppState} from "../../types/store";
import './ThemesList.scss'

const ThemesList = () => {
    const {menu} = useSelector((state: IAppState) => state);
    const [selectedAlgorithm, setSelectedAlgorithm] = useState<number>()
    return (
        <>
            {
                menu.map(menuItem=> {
                    return (
                        <li className={"theme-list-item"}
                            key={menuItem.theme_id}>
                            <NavLink
                                to={`/topics/${menuItem.theme_id}`}
                                className={"theme-list-item"}>
                                {
                                    ({isActive}) => {
                                        if (isActive){
                                            setSelectedAlgorithm(menuItem.theme_id);
                                        }
                                        return <span>{menuItem.title}</span>
                                    }
                                }
                            </NavLink>
                            <ul className={'second-layer-list'}>
                                {
                                    menuItem.algorithms.map(algorithm => {
                                        if(selectedAlgorithm === menuItem.theme_id ){
                                            return (
                                                <>
                                                    <li className={'second-layer-list-item'}
                                                        key={algorithm.algorithm_id}>
                                                        <NavLink
                                                            to={`/topics/${menuItem.theme_id}/${algorithm.algorithm_id}/theory`}
                                                            className={"theme-list-item"}>
                                                            {algorithm.title}
                                                        </NavLink>
                                                    </li>
                                                </>
                                            )
                                        }
                                    })}
                            </ul>
                        </li>
                    )
                })
            }
        </>
    );
};

export default ThemesList;

// return(
//     <>
//         <ul className={'second-layer-list'}>
//             {
//                 menuItem.algorithms.map(algorithm => {
//                     return (
//                         <>
//                             <li className={'second-layer-list-item'}
//                                 key={algorithm.algorithm_id}>
//                                 <NavLink
//                                     to={`/topics/${algorithm.algorithm_id}/theory`}
//                                     className={"theme-list-item"}>
//                                     {algorithm.title}
//                                 </NavLink>
//                             </li>
//                         </>
//                     )
//                 })}
//         </ul>
//     </>
// );