import {Box, CircularProgress} from "@mui/material";
import * as React from "react";
import {FC} from "react";


interface EventProps {
    total:number,
    done: number,
}

const Spinner:FC<EventProps>= ({total, done})=> {
    return (

        <Box sx={{ position: 'relative'}}>
            <CircularProgress
                variant="determinate"
                sx={{
                    color: " #00E861",
                    opacity: "0.5" ,
                }}
                size={50}
                thickness={6}
                value={100}
            />
            <CircularProgress
                variant="determinate"
                value={done*100/total}
                sx={{
                    color: "#00E861",
                    position: 'absolute',
                    left: 0,
                    marginLeft: "auto",
                    marginRight: "auto",
                    right: 0,
                    textAlign: "center",
                }}
                size={50}
                thickness={6}
            />
        </Box>
    );
}

export default Spinner