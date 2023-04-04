import * as React from "react"
import './Button.scss'

interface ButtonProps{
    text:string
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