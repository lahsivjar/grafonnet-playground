import { combineReducers } from 'redux';
import RunReducer from './runReducer'

const reducers = combineReducers({
    RunReducer: RunReducer,
});

export default reducers;
