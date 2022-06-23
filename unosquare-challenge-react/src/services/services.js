import { host } from "../constants/constants";


export const getServices = async() => {

    const url = host + "/cluster/services"

    const res = await fetch( url );
    const data = await res.json();

    return data
}