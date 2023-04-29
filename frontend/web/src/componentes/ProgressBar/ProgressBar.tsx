import React from 'react';
import EventProgress from "../EventProgress/EventProgress";
import "./ProgressBar.scss"

const ProgressBar = () => {
    return (
        <div className={"progress-container"}>
            <h3 className={"progress-title"}> Прогресс изучения</h3>
            <div className="progress-events">
                <EventProgress total={25} done={13} description={"теорий прочитано"}/>
                <EventProgress total={25} done={1} description={"задач выполнено"}/>
            </div>
        </div>
    );
};

export default ProgressBar;