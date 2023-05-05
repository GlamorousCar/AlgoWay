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
        localStorage.setItem('solution', value);
        setSolution(value);
    }

    function getLineNumbers() {
        const lineCount = solution?.split(/\r\n|\r|\n/).length;
        return Array.from(Array(lineCount), (_, i) => i + 1).join('\n');
    }

    const rows = solution ? solution.split(/\r\n|\r|\n/).length + 1 : 2;

    function handleTaskSelect(event:any) {
        console.log(event.target.value)
        setSelectedTask(event.target.value);
    }
    return (
        <div className='practice-block'>
            <div className="information-block">
                {tasks.length ===0 ? <span>Задач для данного алгоритма еще не составлено</span>:<span>{tasks[0].title}</span>}
                <div className="task-select">
                    {tasks.length !==0
                        ?
                        <>
                            <label htmlFor="task-select">Выбрать задачу по данному алгоритму :</label>
                            <select
                                id="task-select"
                                value={selectedTask}
                                onChange={(event)=>handleTaskSelect(event)}
                            >
                                {tasks.map((task,index)=> (
                                    <option key={task.id} value={index}>{task.title}</option>
                                ))}
                            </select>
                        </>
                        :
                        null}

                </div>

                <div className="task-content">
                   <p>{tasks[selectedTask]?.content}</p>
                </div>
            </div>
            <div className="block-solve">
                <div className="code-block">
                    <div className="line-numbers">{getLineNumbers()}</div>
                    <textarea
                        value={localStorage.getItem('solution')!} onChange={(e) => handleChange(e.target.value)}
                        className="code-input"
                        rows={rows}
                        cols={10}/>
                </div>
            </div>
        </div>
    );
};

export default PracticeBlock;