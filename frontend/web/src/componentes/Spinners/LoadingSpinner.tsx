import React from 'react';
import {Box, CircularProgress} from "@mui/material";
import './LoadingSpinner.scss';

const LoadingSpinner = () => {
    return (
        <Box sx={{ display: 'flex' }}>
            <CircularProgress className={'loader'} sx={{color:"#00E861"}}/>
        </Box>
    );
};

export default LoadingSpinner;