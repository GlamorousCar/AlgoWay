import * as React from "react"
import './Drawer.scss'
import ProfileBlock from "./ProfileBlock/ProfileBlock";
import {NavLink} from "react-router-dom";

const Drawer = ()=>{
    return(
        <div className={'drawer'}>
            <ul className={"theme-list"}>
                <li className={"theme-list-item"}>
                    <NavLink to={"/topics/122"} className={"theme-list-item"}>Вычислительная сложность</NavLink>
                </li>
                <li className={"theme-list-item"}>
                    <NavLink to={"/topics/123"} className={"theme-list-item"}>Деревья поиска</NavLink>
                    <ul className={'second-layer-list'}>
                        <li className={'second-layer-list-item'}><span>Деревья в STL</span></li>
                        <li className={'second-layer-list-item active'}><span>Декартово дерево</span></li>
                        <li className={'second-layer-list-item'}><span>Неявный ключ</span></li>
                    </ul>
                </li>
                <li className={"theme-list-item"}>
                    <NavLink to={"/topics/124"} className={"theme-list-item"}>Вычислительная сложность</NavLink>
                </li>
                <li className={"theme-list-item"}>
                    <NavLink to={"/topics/125"} className={"theme-list-item"}>Вычислительная сложность</NavLink>
                </li>
                <li className={"theme-list-item"}>
                    <NavLink to={"/topics/126"} className={"theme-list-item"}>Вычислительная сложность</NavLink>
                </li>
                <li className={"theme-list-item"}>
                    <NavLink to={"/topics/127"} className={"theme-list-item"}>Вычислительная сложность</NavLink>
                </li>
                <li className={"theme-list-item"}>
                    <NavLink to={"/topics/128"} className={"theme-list-item"}>Вычислительная сложность</NavLink>
                </li>
                <li className={"theme-list-item"}>
                    <NavLink to={"/topics/129"} className={"theme-list-item"}>Вычислительная сложность</NavLink>
                </li>
                <li className={"theme-list-item"}>
                    <NavLink to={"/topics/130"} className={"theme-list-item"}>Вычислительная сложность</NavLink>
                </li>
                <li className={"theme-list-item"}>
                    <NavLink to={"/topics/131"} className={"theme-list-item"}>Вычислительная сложность</NavLink>
                </li>
                <li className={"theme-list-item"}>
                    <NavLink to={"/topics/132"} className={"theme-list-item"}>Вычислительная сложность</NavLink>
                </li>
                <li className={"theme-list-item"}>
                    <NavLink to={"/topics/132"} className={"theme-list-item"}>Вычислительная сложность</NavLink>
                </li>
                <li className={"theme-list-item"}>
                    <NavLink to={"/topics/133"} className={"theme-list-item"}>Вычислительная сложность</NavLink>
                </li>
                <li className={"theme-list-item"}>
                    <NavLink to={"/topics/134"} className={"theme-list-item"}>Вычислительная сложность</NavLink>
                </li>
                <li className={"theme-list-item"}>
                    <NavLink to={"/topics/135"} className={"theme-list-item"}>Вычислительная сложность</NavLink>
                </li>
                <li className={"theme-list-item"}>
                    <NavLink to={"/topics/128"} className={"theme-list-item"}>Вычислительная сложность</NavLink>
                </li>


            </ul>
            <ProfileBlock/>
        </div>
    )
}
export default Drawer