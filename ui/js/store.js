import { createStore, applyMiddleware } from 'redux';
import thunk from 'redux-thunk';
import logger from 'redux-logger';
import promise from 'redux-promise-middleware';

import rootReducer from './reducers/index';

const middlewares = [promise, thunk];

if (process.env.NODE_ENV === `development`) {
    middlewares.push(logger);
}

const store = createStore(
    rootReducer,
    applyMiddleware(...middlewares)
);

export default store;
