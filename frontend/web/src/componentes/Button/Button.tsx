import * as React from "react"
import './Button.scss'

interface ButtonProps{
    text:string,
    onClick?:()=>void;
}

const Button = (props:ButtonProps)=>{


    return(
        <>
            <button onClick={props.onClick}>
                {props.text}
            </button>

        </>
    )
}

export default Button