export interface IMenu{
    theme_id:number,
    title:string,
    position:number,
    algorithms:IAlgorithm[]
}
export interface IAlgorithm{
    algorithm_id:number,
    title:string,
    description:string,
    position:number,
    theme_id: number,
    content:string
}
export interface IUser{
    email:string,
    isActivated:boolean,
    id:string
}
export interface ITask{
    id:number,
    is_solved:boolean,
    title:string,
    content:string
}