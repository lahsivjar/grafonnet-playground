import axios from 'axios';

import { CODE_UPDATE } from './types';

export function CodeUpdate(code) {
    const payload = {
        code: code,
    };

    return dispatch => dispatch({
        type: CODE_UPDATE,
        payload: payload,
    });
}
