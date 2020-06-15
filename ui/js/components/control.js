import React from 'react';

import { withStyles } from '@material-ui/core/styles';
import CloudDownloadIcon from '@material-ui/icons/CloudDownload';
import SendIcon from '@material-ui/icons/Send';
import IconButton from '@material-ui/core/IconButton';
import Select from '@material-ui/core/Select';
import MenuItem from '@material-ui/core/MenuItem';
import InputLabel from '@material-ui/core/InputLabel';
import Grid from '@material-ui/core/Grid';
import Icon from '@material-ui/core/Icon';
import { connect } from 'react-redux';
import { Run } from '../actions/run';
import { ThemeUpdate } from '../actions/themeUpdate';
import fileDownload from 'js-file-download';

import 'codemirror/theme/ambiance.css';
import 'codemirror/theme/ayu-mirage.css';
import 'codemirror/theme/cobalt.css';
import 'codemirror/theme/darcula.css';
import 'codemirror/theme/material.css';
import 'codemirror/theme/monokai.css';
import 'codemirror/theme/solarized.css';
import '../../style/components/control.css';

const styles = theme => ({
    button: {
        marginRight: '8px',
    }
});

const themes = [
    'default',
    'ambiance',
    'ayu-mirage', 'cobalt',
    'darcula',
    'material',
    'monokai',
    'solarized dark',
    'solarized light',
];

class Control extends React.Component {
    runCode = () => {
        const data = {
            code: this.props.code,
        }
        this.props.Run(data)
    }

    themeUpdate = (event) => {
        this.props.ThemeUpdate(event.target.value)
    }

    download = () => {
        fileDownload(
            this.props.code,
            'grafonnet-playground.jsonnet',
            'text/plain;charset=utf-8'
        )
    }

    render() {
        const { classes } = this.props;
        return(
            <div className='control'>
                <Grid container spacing={0}>
                    <Grid item xs={12}>
                        <Grid container justify='flex-end'>
                            <Grid item className='control-select-container'>
                                <span className='control-select-theme'>Theme</span>
                                <Select
                                    value={this.props.theme}
                                    onChange={this.themeUpdate}
                                    label='Theme'
                                >
                                    {themes.map((value, idx) => {
                                        return <MenuItem key={idx} value={value}>{value}</MenuItem>
                                    })}
                                </Select>
                            </Grid>
                            <Grid item>
                                <IconButton
                                    className={classes.button}
                                    aria-label='Download'
                                    onClick={this.download}
                                >
                                    <CloudDownloadIcon />
                                </IconButton>
                            </Grid>
                            <Grid item>
                                <IconButton
                                    color='primary'
                                    className={classes.button}
                                    onClick={this.runCode}
                                    aria-label='Run'
                                    disabled={this.props.loading}
                                >
                                    <SendIcon />
                                </IconButton>
                            </Grid>
                        </Grid>
                    </Grid>
                </Grid>
            </div>
        );
    }
}

const mapStateToProps = state => {
    return { ...state.RunReducer };
}

export default connect(mapStateToProps, { Run, ThemeUpdate }) (withStyles(styles) (Control));
