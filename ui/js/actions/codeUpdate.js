import axios from 'axios';
import { throttle } from 'throttle-debounce';
import { set } from 'idb-keyval';

import { CODE_UPDATE } from './types';

export function CodeUpdate(code) {
    const payload = {
        code: code,
    };

    throttledSave(payload);

    return dispatch => dispatch({
        type: CODE_UPDATE,
        payload: payload,
    });
}

const throttledSave = throttle(100, (payload) => {
    set('code', payload)
        .catch(err => console.error('Failed to set code in local storage'))
})
