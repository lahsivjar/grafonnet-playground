import React from 'react';
import ReactDOM from 'react-dom';

import CssBaseline from '@material-ui/core/CssBaseline';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';

import Editor from './components/editor';
import Graph from './components/graph';
import Control from './components/control';
import '../style/app.css';

function App() {
    return (
        <Grid container component="main" className="app">
            <Grid item xs={12} sm={6} component={Paper}>
                <Control />
                <Editor />
            </Grid>
            <Grid item xs={12} sm={6} component={Paper}>
                <Graph />
            </Grid>
        </Grid>
    );
}

ReactDOM.render(
    <App />,
    document.getElementById('container'),
);
