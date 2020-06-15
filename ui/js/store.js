import { createStore, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import logger from 'redux-logger';
import promise from 'redux-promise-middleware';
import { get } from 'idb-keyval';

import { CODE_UPDATE, THEME_UPDATE } from './actions/types';
import rootReducer from './reducers/index';

const middlewares = [promise, thunk];

if (process.env.NODE_ENV === `development`) {
    middlewares.push(logger);
}

const store = createStore(
    rootReducer,
    applyMiddleware(...middlewares)
)

get('code')
    .then(val => {
        if (val !== undefined) {
            store.dispatch({
                type: CODE_UPDATE,
                payload: val,
            })
        }
    })

get('theme')
    .then(val => {
        if (val !== undefined) {
            store.dispatch({
                type: THEME_UPDATE,
                payload: val,
            })
        }
    })

export default store;
