
import useFetchHook from "../Hooks/UseFetchHook";

const UseAlgoService = () => {

    const {getRequest, error, clearError} = useFetchHook();

    const _baseUrl = "https://bbaunqpcv1t23s9skmhv.containers.yandexcloud.net";

    const getMenuTopics = async () =>{
        const res = await getRequest(`${_baseUrl}/themes/menu`);
        return res;
    }

    const getAlgorithmTheory = async (id: string | undefined)=>{
        const res = await getRequest(`${_baseUrl}/theory?algo_id=${id}`)
        console.log(res)
        return res;
    }

    return {getMenuTopics,getAlgorithmTheory, error, clearError};
};

export default UseAlgoService;