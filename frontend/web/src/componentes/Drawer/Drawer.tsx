import * as React from "react"
import './Drawer.scss'
import ProfileBlock from "./ProfileBlock/ProfileBlock";

const Drawer = ()=>{
    return(
        <div className={'drawer'}>
            <ul className={"theme-list"}>
                <li className={"theme-list-item"}>  <span className={"theme-list-item"}>Вычислительная сложность </span></li>
                <li className={"theme-list-item"}>
                    <span className={"theme-list-item active"}>Деревья поиска</span>
                    <ul className={'second-layer-list'}>
                        <li className={'second-layer-list-item'}><span>Деревья в STL</span></li>
                        <li className={'second-layer-list-item active'}><span>Декартово дерево</span></li>
                        <li className={'second-layer-list-item'}><span>Неявный ключ</span></li>
                    </ul>
                </li>
                <li className={"theme-list-item"}>  <span className={"theme-list-item"}>Вычислительная сложность </span></li>
                <li className={"theme-list-item"}>  <span className={"theme-list-item"}>Вычислительная сложность </span></li>
                <li className={"theme-list-item"}>  <span className={"theme-list-item"}>Вычислительная сложность </span></li>
                <li className={"theme-list-item"}>  <span className={"theme-list-item"}>Вычислительная сложность </span></li>
                <li className={"theme-list-item"}>  <span className={"theme-list-item"}>Вычислительная сложность </span></li>
                <li className={"theme-list-item"}>  <span className={"theme-list-item"}>Вычислительная сложность </span></li>
                <li className={"theme-list-item"}>  <span className={"theme-list-item"}>Вычислительная сложность </span></li>
                <li className={"theme-list-item"}>  <span className={"theme-list-item"}>Вычислительная сложность </span></li>

            </ul>
            <ProfileBlock/>
        </div>
    )
}
export default Drawer