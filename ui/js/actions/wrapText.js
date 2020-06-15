import axios from 'axios';
import { set } from 'idb-keyval';

import { WRAP_TEXT } from './types';

export function WrapText(isWrapEnabled) {
    const payload = {
        wrap: isWrapEnabled,
    };

    set('wrap', payload)
        .catch(err => console.error('Failed to set wrap in local storage'))

    return dispatch => dispatch({
        type: WRAP_TEXT,
        payload: payload,
    });
}
