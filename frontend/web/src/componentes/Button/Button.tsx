import * as React from "react"
import './Button.scss'
import {ReactNode} from "react";

interface ButtonProps{
    text:string,
}

const Button = (props:ButtonProps)=>{
    return(
        <>
            <button>
                {props.text}
            </button>

        </>
    )
}

export default Button