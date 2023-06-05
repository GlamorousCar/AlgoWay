import * as React from "react"
import './Button.scss'
import { fontGrid } from "@mui/material/styles/cssUtils";

interface ButtonProps{
    text:string,
    onClick?:()=>void;
    textSize?:string;
    height?:string;
    
}

const Button = (props:ButtonProps)=>{


    return(
        <>
            <button style={{fontSize:props.textSize, height:props.height}} onClick={props.onClick}>
                {props.text}
            </button>

        </>
    )
}

export default Button