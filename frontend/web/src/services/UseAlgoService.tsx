
import useFetchHook from "../Hooks/UseFetchHook";

const UseAlgoService = () => {

    const {getRequest, error, clearError} = useFetchHook();

    const _baseUrl = "https://bba2ku7v26adaqseufur.containers.yandexcloud.net";

    const getMenuTopics = async () =>{
        const res = await getRequest(`${_baseUrl}/themes/menu`);
        return res;
    }

    const getAlgorithmTheory = async (id: number )=>{
        const res = await getRequest(`${_baseUrl}/theory?algo_id=${id}`)
        return res;
    }

    const getAlgorithmTasks = async (id:number) =>{
        const res = await getRequest(`${_baseUrl}/task?algo_id=${id}`);
        return res;
    }
    return {getMenuTopics,getAlgorithmTheory,getAlgorithmTasks, error, clearError};
};

export default UseAlgoService;