import React from 'react';

import { connect } from 'react-redux';
import '../../style/components/graph.css';

class Graph extends React.Component {
    render() {
        return(
            <div className='graph'>
                <iframe className='graph-frame'
                    src={this.props.url}
                />
            </div>
        );
    }
}

const mapStateToProps = state => {
    return { ...state.RunReducer };
}

export default connect(mapStateToProps, {}) (Graph);
