import React from 'react';

import '../../style/components/graph.css';

class Graph extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            source: "",
        }
    }

    render() {
        return(
            <div className="graph">
                <iframe className="graph-frame"
                    src={this.state.source}
                />
            </div>
        );
    }
}

export default Graph;
