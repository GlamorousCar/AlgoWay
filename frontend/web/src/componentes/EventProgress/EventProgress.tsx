import React, {FC} from 'react';
import "./EventProgress.scss"
import Spinner from '../Spinners/Spinner';

interface EventProps {
    total:number,
    done: number,
    description:string
}

const EventProgress:FC<EventProps>= ({total, done, description}) => {
    return (
        <div className="event-progress">
            <Spinner total={total} done={done} />
            <div className="progress">
                <h4 className={"done-event"}>{done}</h4>
                <h4>/</h4>
                <h4 className={"total-event"}>{total}</h4>
            </div>
            <div className="progress-description">
                <p>{description}</p>
            </div>

        </div>
    );
};

export default EventProgress;