import React from 'react';
import {Modal} from "@mui/material";
import ModalDialog from '@mui/joy/ModalDialog';


interface Props{
    setOpen: (b: boolean)=>void,
    children:React.ReactNode,
    open:boolean,

}
const ModalSheet = (props:Props) => {
    return (
        <>
            <Modal open={props.open} onClose={() => props.setOpen(false)}>
                <ModalDialog
                    aria-labelledby="basic-modal-dialog-title"
                    aria-describedby="basic-modal-dialog-description"
                    sx={{ maxWidth: 500 }}
                >
                    {props.children}
                </ModalDialog>
            </Modal>
        </>
    );
};

export default ModalSheet;