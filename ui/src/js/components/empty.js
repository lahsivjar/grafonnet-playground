import React from 'react';
import CircularProgress from '@material-ui/core/CircularProgress';
import WarningIcon from '@material-ui/icons/Warning';
import Grid from '@material-ui/core/Grid';

import '../../style/components/empty.css';

class Empty extends React.Component {
    getContent = () => {
        if (this.props.loading) {
            return (
                <div className="loading-content">
                    <CircularProgress />
                </div>
            )
        }
        if (this.props.errorMsg) {
            return (
                <div className="error-container">
                    <Grid container direction="row" alignItems="center">
                        <Grid item>
                            <WarningIcon
                                className="warning-icon"
                                style={{ fontSize: 35 }}
                            />
                        </Grid>
                        <Grid item>
                            <span className="error-container-header">
                                Error occurred while evaluating jsonnet code
                            </span>
                        </Grid>
                    </Grid>
                    <pre className="error-container-body">
                        {this.props.errorMsg}
                    </pre>
                </div>
            )
        }
        return (
            <div className="empty-container">
                <span className="empty-container-header">
                    Grafonnet Playground
                </span>
                <span className="empty-container-body">
                    Edit grafonnet code in editor and view the rendered graph
                </span>
            </div>
        )
    }

    render() {
        return(
            <div className='empty'>
                {this.getContent()}
            </div>
        );
    }
}

const mapStateToProps = state => {
    return { ...state.RunReducer };
}

export default Empty;
