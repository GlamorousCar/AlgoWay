import React, {useEffect, useState} from 'react';
import "./PracticeBlock.scss"
import {useParams} from "react-router-dom";
import useAlgoService from "../../services/UseAlgoService";
import {ITask} from "../../types/types";

const PracticeBlock = () => {

    const [solution, setSolution] = useState<string>();
    const [selectedTask, setSelectedTask] = useState<number>(1);
    const [tasks, setTasks] = useState<ITask[]>([]);
    const {algorithmId} = useParams();


    const {getAlgorithmTasks} = useAlgoService();


    useEffect(() => {
        getResources(Number(algorithmId));
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, [algorithmId]);

    const getResources = (algorithmId: number) => {
        getAlgorithmTasks(algorithmId)
            .then(onTasksLoaded)
    }


    const onTasksLoaded = (tasks: ITask[]) => {
        setTasks(tasks);
        console.log(algorithmId, tasks);
    }


    function handleChange(value: string) {
        console.log(value)

        setSolution(value);
    }

    function getLineNumbers() {
        const lineCount = solution?.split(/\r\n|\r|\n/).length;
        return Array.from(Array(lineCount), (_, i) => i + 1).join('\n');
    }

    const rows = solution ? solution.split(/\r\n|\r|\n/).length + 1 : 2;

    function handleTaskSelect(event: any) {
        console.log(event.target.value)
        setSelectedTask(event.target.value);
    }

    return (
        <div className='practice-block'>
            <div className="information-block">
                {tasks.length === 0 ? <span>Задач для данного алгоритма еще не составлено</span> : null}
                {tasks.length !== 0
                    ?
                    <>
                        <div className="task-select">
                            {/*<label htmlFor="task-select">Выбрать задачу по {tasks[0].title} :</label>*/}
                            <div className="select-block">
                                <select
                                    id="task-select"
                                    value={selectedTask}
                                    onChange={(event) => handleTaskSelect(event)}
                                >
                                    {tasks.map((task, index) => (
                                        <option key={task.id} value={index}>{task.title}</option>
                                    ))}
                                </select>
                                <div className="task-difficulty easy">
                                    <li className="task-difficulty-item"> Простая</li>
                                </div>
                            </div>
                        </div>
                        <div className="task-content">
                            <p>{tasks[selectedTask]?.content}</p>
                            <h2 className={"input-info-title"}>Формат ввода</h2>
                            <p className={"input-info-content"}>
                                В первой строке входных данных записаны два натуральных числа N и M,
                                не превосходящих 100 — размеры таблицы.
                                Далее идёт N строк, каждая из которых содержит M чисел, разделённых пробелами — описание
                                таблицы.
                                Все числа в клетках таблицы целые и могут принимать значения от 0 до 100.
                            </p>
                            <h2 className={"output-info-title"}>Формат вsвода</h2>
                            <p className={"output-info-content"}>
                                Первая строка выходных данных содержит максимальную возможную сумму, вторая — маршрут, на
                                котором достигается эта сумма. Маршрут выводится в виде последовательности, которая должна
                                содержать N-1 букву D, означающую передвижение вниз и M–1 букву R, означающую передвижение
                                направо.
                                Если таких последовательностей несколько, необходимо вывести ровно одну (любую) из них.
                            </p>
                            <div className="visual-info-block">
                                <div className="visual-info-block-item">
                                    <h2 className="visual-info-block-item-title">
                                        Ввод
                                    </h2>
                                    <p>
                                        6 4 2 1 7 5 9 8 3
                                    </p>
                                </div>
                                <div className="visual-info-block-item">
                                    <h2 className="visual-info-block-item-title">
                                        Вывод
                                    </h2>
                                    <p>
                                        6 <br/>
                                        4 <br/>
                                        2 <br/>
                                        1 <br/>
                                        7 <br/>
                                        5 <br/>
                                        9 <br/>
                                        8 <br/>
                                        3
                                    </p>
                                </div>
                            </div>
                        </div>

                    </>
                    :
                    null}
            </div>
            <div className="block-solve">
                <div className="code-block">
                    <div className="line-numbers">{getLineNumbers()}</div>
                    <textarea
                        value={solution} onChange={(e) => handleChange(e.target.value)}
                        className="code-input"
                        rows={rows}
                        cols={13}/>
                </div>
                <div className="send-solution-block">
                    <div className="solution-result-block">
                        <p className="solution-result-block-title"> Результаты проверки </p>

                    </div>
                    <button className="send-solution">
                        Отправить
                    </button>
                </div>
            </div>
        </div>
    );
};

export default PracticeBlock;