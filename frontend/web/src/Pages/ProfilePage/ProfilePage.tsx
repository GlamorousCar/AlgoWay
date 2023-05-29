import React, {useState} from 'react';
import Avatar from '../../images/Avatar.svg';
import Github from '../../images/GitHub logo.svg';
import Location from '../../images/Location.svg';
import "./ProfilePage.scss"
import ActivityBlock from "../../componentes/ActivityBlock/ActivityBlock";
import ProgressBar from "../../componentes/ProgressBar/ProgressBar";
import ModalSheet from "../../componentes/ModalSheet/ModalSheet";
import Typography from '@mui/joy/Typography';
import { FormControl, FormLabel, Stack} from "@mui/joy";
import Button from "../../componentes/Button/Button";
import {Input} from "@mui/material";




const ProfilePage = () => {

    const[open, setOpen] = useState<boolean>(false)
    const modelHandler = (event:any)=>{
        event.preventDefault();
        setOpen(true);

    }

    return (
        <>
            <div className='profile'>
                <div className="profile-info">
                    <img className='main-profile-photo' src={Avatar} alt=""/>
                    <div className="profile-username">D.Tore.Y</div>
                    <button className="edit-profile-button" onClick={(e)=>modelHandler(e)}> Редактировать профиль</button>
                    <div className="location-block">
                        <img src={Location} alt="" className="location-image"/>
                        <span className='location-text profile-info-text'>Россия</span>
                    </div>
                    <div className="github-link-block">
                        <img src={Github} alt="" className="github-image-profile"/>
                        <span className='github-link profile-info-text'>https://github.com/DToreY</span>
                    </div>
                </div>
                <div className="profile-activity">
                    <span className={'profile-title'}>Активность за год</span>
                    <ActivityBlock/>
                    <ProgressBar/>
                </div>
            </div>
            {open?
                <ModalSheet setOpen={setOpen} open={open}>
                    <Typography id="basic-modal-dialog-title" sx={{textAlign:"center"}} component="h2">
                       Редактировать профиль
                    </Typography>
                    <form
                        onSubmit={(event: React.FormEvent<HTMLFormElement>) => {
                            event.preventDefault();
                            setOpen(false);
                        }}
                    >
                        <Stack spacing={2}>
                            <FormControl>
                                <FormLabel>Никнейм</FormLabel>
                                <Input sx={{
                                    '&.MuiInputBase-root-MuiInput-root:after ': {
                                        border:"2px solid #00E861"
                                    }
                                }} autoFocus required/>
                            </FormControl>
                            <FormControl>
                                <FormLabel>Ссылка на GitHub</FormLabel>
                                <Input  sx={{
                                    '&.MuiInputBase-root-MuiInput-root:after ': {
                                        border:"2px solid #00E861"
                                    }
                                }} required />
                            </FormControl>
                            <FormControl>
                                <FormLabel>Фотография</FormLabel>
                                <Input  sx={{
                                    '& .MuiInputBase-root-MuiInput-root:focus ': {
                                        border:"2px solid #00E861"
                                    }
                                }} required />
                            </FormControl>
                            <Button text={"Сохранить"} ></Button>
                        </Stack>
                    </form>
                </ModalSheet>
                :null}
        </>

    );
};

export default ProfilePage;