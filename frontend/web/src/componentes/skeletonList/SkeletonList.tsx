import React from 'react';
import {Skeleton} from "@mui/material";

const SkeletonList = () => {
    return (
        <div className={"skeleton-block"}>
            {Array.from({ length: 9 }).map((item, index) => (
                <Skeleton height={35} width={"200px"} sx={{mt:"2px"}} key={index} animation="wave" />
            ))}
        </div>
    );
};

export default SkeletonList;