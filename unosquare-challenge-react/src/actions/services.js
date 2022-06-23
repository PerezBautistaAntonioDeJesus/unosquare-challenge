import { getServices } from "../services/services";


export const activeService = (service ) => ({
    type: 'active-service',
    payload: {
        ...service
    }
});

export const startServices = () => {
    return async ( dispatch ) => {
        const services = await getServices();
        dispatch( setServices( services ) )
    }
}

export const setServices = ( services ) => ({
    type: 'load-services',
    payload: services,
})