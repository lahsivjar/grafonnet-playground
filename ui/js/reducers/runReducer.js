import querystring from 'query-string';
import random from 'math-random';

import { RUN_PENDING, RUN_FULFILLED, RUN_REJECTED } from '../actions/types';
import { CODE_UPDATE } from '../actions/types';
import { THEME_UPDATE } from '../actions/types';
import { WRAP_TEXT } from '../actions/types';

const initialState ={
    url: '',
    errorMsg: '',
    code: '',
    theme: 'default',
    wrap: true,
    loading: false,
    error: false,
}

function getUrl(baseurl) {
    const parsed = querystring.parse(baseurl);
    parsed.kiosk = 1;
    parsed.buster = random();

    return baseurl + '?' + querystring.stringify(parsed);
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
                url: getUrl(action.payload.url),
                loading: false,
                error: false,
            }
        case RUN_REJECTED:
            var errMsg = 'Unknown error occurred while attempting to run jsonnet';
            if (action.payload.response && action.payload.response.data) {
                errMsg = action.payload.response.data.errorMsg;
            } else if (action.payload.message) {
                errMsg = action.payload.message
            }
            return {
                ...state,
                errorMsg: errMsg,
                url: '',
                loading: false,
                error: true,
            }
        case CODE_UPDATE:
            return {
                ...state,
                ...action.payload,
            }
        case THEME_UPDATE:
            return {
                ...state,
                ...action.payload,
            }
        case WRAP_TEXT:
            return {
                ...state,
                ...action.payload,
            }
        default:
            return state
    }
}
