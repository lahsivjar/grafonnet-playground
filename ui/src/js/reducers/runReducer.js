import { RUN_PENDING, RUN_FULFILLED, RUN_REJECTED } from '../actions/types';
import { CODE_UPDATE } from '../actions/types';

const initialState ={
    url: '',
    errorMsg: '',
    code: '',
    loading: false,
    error: false,
}

export default function RunReducer(state = initialState, action) {
    switch(action.type) {
        case RUN_PENDING:
            return {
                ...state,
                loading: true,
                error: false,
            }
        case RUN_FULFILLED:
            return {
                ...state,
                ...action.payload,
                loading: false,
                error: false,
            }
        case RUN_REJECTED:
            if (action.payload.response && action.payload.response.data) {
                state.errorMsg = action.payload.response.data.errorMsg;
            } else if (action.payload.message) {
                state.errorMsg = action.payload.message
            }
            return {
                ...state,
                loading: false,
                error: true,
            }
        case CODE_UPDATE:
            return {
                ...state,
                ...action.payload,
            }
        default:
            return state
    }
}
