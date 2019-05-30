import React from 'react';
import { connect } from 'react-redux';

import Empty from './empty';
import '../../style/components/graph.css';

class Graph extends React.Component {
    getContent = () => {
        if (this.props.url) {
            return <iframe className='graph-frame' src={this.props.url} />;
        }

        return (
            <Empty
                loading={this.props.loading}
                errorMsg={this.props.errorMsg}
            />
        )
    }

    render() {
        return(
            <div className='graph'>
                {this.getContent()}
            </div>
        );
    }
}

const mapStateToProps = state => {
    return { ...state.RunReducer };
}

export default connect(mapStateToProps, {}) (Graph);
