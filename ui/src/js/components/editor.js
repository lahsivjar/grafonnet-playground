import React from 'react';
import CodeMirror from 'react-codemirror';
import { connect } from 'react-redux';
import { CodeUpdate } from '../actions/codeUpdate';

import 'codemirror/lib/codemirror.css';
import 'codemirror/mode/javascript/javascript';
import '../../style/components/editor.css';

class Editor extends React.Component {
    updateCode = (newCode) => {
        this.props.CodeUpdate(newCode);
    }

    render() {
        const { code } = this.props;
        const options = {
            mode: 'javascript',
            lineNumbers: true,
            indentUnit: 2,
            matchBrackets: true,
        };
        return(
            <div className='editor'>
                <CodeMirror
                    value={code}
                    onChange={this.updateCode}
                    options={options}
                />
            </div>
        );
    }
}

const mapStateToProps = state => {
    return { ...state.RunReducer };
}

export default connect(mapStateToProps, { CodeUpdate }) (Editor);
