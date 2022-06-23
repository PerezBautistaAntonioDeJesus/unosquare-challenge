
const initialState = {
    services: [],
    active: null
}
export const serviceReducer = (state = initialState, action) => {
    switch (action.type) {
        case "active-service":
            return{
                ...state,
                active: {...action.payload}
            }
        case "load-services":
            return{
                ...state,
                services: [...action.payload]
            }
    
        default:
            return state
    }
}