import React, {useState} from 'react';
import {NavLink} from "react-router-dom";
import {useDispatch, useSelector} from "react-redux";
import {IAppState} from "../../types/store";
import './ThemesList.scss'
import { drawerClosing } from '../../store/actions';

const ThemesList = () => {
    const {menu} = useSelector((state: IAppState) => state);
    const dispatch = useDispatch();
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
                                key={menuItem.theme_id}
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
                                                            to={`/topics/${menuItem.theme_id}/${algorithm.algorithm_id}`}
                                                            className={"second-layer-list-item-link"}
                                                            onClick={()=>dispatch(drawerClosing())}
                                                            key={algorithm.algorithm_id}>
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