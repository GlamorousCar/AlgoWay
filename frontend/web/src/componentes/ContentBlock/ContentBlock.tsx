import * as React from "react"
import './ContentBlock.scss'
import { Route, Routes} from "react-router-dom";
import TheoryBlock from "../Pages/TheoryBlockPage/TheoryBlock";
import MainContentPage from "../Pages/MainContentPage/MainContentPage";

const ContentBlock = ()=>{
    return(
            <div className={"main-content-page"}>
                <Routes>
                    <Route path={"/topics/:topicId"} element={<TheoryBlock/>}/>
                    <Route path={"/"} element={<MainContentPage/>}/>
                    <Route path={"*"} element={<MainContentPage/>}/>
                </Routes>
            </div>
    )
}

export default ContentBlock