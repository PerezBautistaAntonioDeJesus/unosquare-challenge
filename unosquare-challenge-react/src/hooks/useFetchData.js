import { useEffect, useState } from "react";
import { getServices } from "../services/services"




export const useFetchServices = () => {

    const [services, setServices] = useState([])
    const [isLoading, setIsLoading] = useState(true)

    const getData = async() => {
        const services = await getServices();
        setServices( services )
        setIsLoading( false )
    }
    useEffect(() => {
        getData()
        
    }, [])

    return {
        services,
        isLoading
    }
}