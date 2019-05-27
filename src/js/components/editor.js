import React from 'react';
import CodeMirror from 'react-codemirror';

import 'codemirror/lib/codemirror.css';
import 'codemirror/mode/javascript/javascript';
import '../../style/components/editor.css';

class Editor extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            code: '',
        };
    }

    updateCode = (newCode) => {
        this.setState(
            {
                code: newCode,
            }
        );
    }

    render() {
        var options = {
            mode: 'javascript',
            lineNumbers: true,
            indentUnit: 2,
            matchBrackets: true,
        }
        return(
            <div className="editor">
                <CodeMirror
                    value={this.state.code}
                    onChange={this.updateCode}
                    options={options}
                />
            </div>
        );
    }
}

export default Editor;
