import React from 'react';
import './ActivityBlock.scss'


const ActivityBlock = () => {
    const months = ["Янв.", "Фев.", "Март", "Апр.", "Май", "Июнь", "Июль", "Авг.", "Сен.", "Окт.", "Ноя.", "Дек."];
    const week_days = ["Пн.", "Вт.", "Ср.", "Чт.", "Пт.", "Сб.", "Вс."];
    return (
        <div className='activity-block'>
            <div className="month-row">
                {months.map((month) => {
                    return (<span className={"month-header"}>{month}</span>)
                })}
            </div>
            <div className="second-block">
                <div className="week-column">
                    {week_days.map((day) => {
                        return (<p className={"week-header"}>{day}</p>)
                    })}
                </div>
                <div className="activity-blocks">
                    {Array.from({length:51}).map((item, index) => {
                        return(
                            <div className="activity-week" key={index}>
                                {Array.from({length:7}).map((item , index) => {
                                    return <div className="activity-week-day" key={index}></div>
                                })}
                            </div>
                        )
                    })}
                </div>
            </div>
            <div className="color-description">
                <span>Меньше </span>
                <div className="activity-week-day zero-solution" key={"0"} ></div>
                <div className="activity-week-day two-solution" key={"2"}></div>
                <div className="activity-week-day four-solution" key={"4"}></div>
                <div className="activity-week-day six-solution" key={"6"}></div>
                <div className="activity-week-day eight-solution" key={"8"}></div>
                <span>Больше </span>
            </div>
        </div>
    )
};
export default ActivityBlock;