import axios from 'axios';

import { RUN } from './types';

export function Run(data) {
    return dispatch => dispatch({
        type: RUN,
        payload: axios.post('/api/v1/run', data)
                   .then(res => res.data)
    });
}
