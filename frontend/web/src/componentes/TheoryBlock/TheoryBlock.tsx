import * as React from "react"
import './TheoryBlock.scss'
import green_book from "../../images/Vector_book_green.svg";
import grey_code from "../../images/Vector_code_default.svg"
const TheoryBlock = ()=>{
    return(
        <div className={"theory"}>
            <div className="container">
                <h3 className={"theory-title"}>Декартово дерево</h3>
                <nav>
                    <ul className={"nav-switch"}>
                        <li className={"nav-item theory active"}>
                            <img  src={green_book} alt=""/>
                            <span>Теория</span>
                        </li>
                        <li className={"nav-item practice "}>
                            <img src={grey_code} alt=""/>
                            <span>Практика</span>
                        </li>

                    </ul>
                </nav>

                <div className="content">
                    <p> <span className={"bold"}>Рене Декарт</span> (фр. René Descartes) — великий французский математик и философ XVII века. Рене Декарт не является создателем декартова дерева, однако он является создателем декартовой системы координат, которую мы все знаем и любим. </p>
                    <h6>Декартово дерево же определяется и строится так:</h6>
                    <ul>
                        <li>Нанесём на плоскость набор из n точек. Их x зачем-то назовем ключом, а y приоритетом.</li>
                        <li>Выберем самую верхнюю точку (с наибольшим y, а если таких несколько — любую) и назовём её корнем.</li>
                        <li>От всех вершин, лежащих слева (с меньшим x) от корня, рекурсивно запустим этот же процесс. Если слева была хоть одна вершина, то присоединим корень левой части в качестве левого сына текущего корня.</li>
                        <li>Аналогично, запустимся от правой части и добавим корню правого сына.</li>
                    </ul>
                    <p>Заметим, что если все y и x различны, то дерево строится однозначно. Если нарисовать получившуюся структуру на плоскости, то получится действительно дерево — по традиции, корнем вверх:</p>

                    <p>To be continued....</p>
                </div>
            </div>
        </div>
    );
}


export default TheoryBlock