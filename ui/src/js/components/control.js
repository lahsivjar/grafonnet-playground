import React from 'react';

import { withStyles } from '@material-ui/core/styles';
import CloudDownloadIcon from '@material-ui/icons/CloudDownload';
import SendIcon from '@material-ui/icons/Send';
import IconButton from '@material-ui/core/IconButton';
import Grid from '@material-ui/core/Grid';
import Icon from '@material-ui/core/Icon';
import { connect } from 'react-redux';
import { Run } from '../actions/run';

import '../../style/components/control.css';

const styles = theme => ({
    button: {
        marginRight: '8px',
    }
});


class Control extends React.Component {
    runCode = () => {
        const data = {
            code: this.props.code,
        }
        this.props.Run(data)
    }

    render() {
        const { classes } = this.props;
        return(
            <div className='control'>
                <Grid container spacing={0}>
                    <Grid item xs={12}>
                        <Grid container justify='flex-end'>
                            <Grid item>
                                <IconButton
                                    className={classes.button}
                                    aria-label='Download'
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

export default connect(mapStateToProps, { Run }) (withStyles(styles) (Control));
