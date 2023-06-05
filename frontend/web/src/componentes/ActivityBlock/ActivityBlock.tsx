import React, { useEffect, useRef } from 'react';
import './ActivityBlock.scss'
import {activityData} from './mockData'


const ActivityBlock = () => {
    const months = ["Янв.", "Фев.", "Март", "Апр.", "Май", "Июнь", "Июль", "Авг.", "Сен.", "Окт.", "Ноя.", "Дек."];
    const week_days = ["Пн.", "Вт.", "Ср.", "Чт.", "Пт.", "Сб.", "Вс."];


    const itemsRef = useRef<Array<HTMLDivElement | null>>([]);

    useEffect(()=>{
        itemsRef.current = itemsRef.current.slice(0, activityData.map((item)=>item).length)
    },[activityData])

   let itemClasses = '';
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
                    {activityData.map((week, index) => {
                        return(
                            <div className="activity-week" key={index}>
                                {week.map((day , index) => {
                                    if(day.contributions == 0 ){
                                        itemClasses = "activity-week-day zero";
                                    }else if(day.contributions <=2 ){
                                        itemClasses = "activity-week-day two";
                                    }else if(day.contributions <=4 ){
                                        itemClasses = "activity-week-day four";
                                    }else if(day.contributions <=6 ){
                                        itemClasses = "activity-week-day six";
                                    }
                                    else if(day.contributions >=7 ){
                                        itemClasses = "activity-week-day eight";
                                    }
                                    return <div 
                                    className={itemClasses }
                                    key={index}
                                    ref={el => itemsRef.current[day.id] = el} 
                                    >
                                    </div>
                                })}
                            </div>
                        )
                    })}
                </div>
            </div>
            <div className="color-description">
                <span>Меньше </span>
                <div className="activity-week-day zero-solution " key={"0"} ></div>
                <div className="activity-week-day two-solution " key={"2"}></div>
                <div className="activity-week-day four-solution " key={"4"}></div>
                <div className="activity-week-day six-solution " key={"6"}></div>
                <div className="activity-week-day eight-solution " key={"8"}></div>
                <span>Больше </span>
            </div>
        </div>
    )
};
export default ActivityBlock;