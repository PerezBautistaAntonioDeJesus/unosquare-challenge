import { applyMiddleware, combineReducers, compose, createStore } from "redux";
import thunk from "redux-thunk";
import { serviceReducer } from "../reducers/serviceReducer";


const composeEnhancers = (typeof window !== 'undefined' && window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__) || compose;

const reducers =  combineReducers({
    services: serviceReducer
});


export const store = createStore(
    reducers,  
    composeEnhancers( applyMiddleware( thunk )))