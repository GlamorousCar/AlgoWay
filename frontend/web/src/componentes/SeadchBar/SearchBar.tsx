import * as React from "react"
import "./SearchBar.scss"

const SearchBar = ()=>{
    return(
        <div className={"search-bar"}>
            <input type="text" placeholder={"Найти на сайте"}/>
        </div>
    )
}
export default SearchBar