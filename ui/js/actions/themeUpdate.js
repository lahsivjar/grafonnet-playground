import axios from 'axios';
import { set } from 'idb-keyval';

import { THEME_UPDATE } from './types';

export function ThemeUpdate(theme) {
    const payload = {
        theme: theme,
    };

    set('theme', payload)
        .catch(err => console.error('Failed to set theme in local storage'))

    return dispatch => dispatch({
        type: THEME_UPDATE,
        payload: payload,
    });
}
