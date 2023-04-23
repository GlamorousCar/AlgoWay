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