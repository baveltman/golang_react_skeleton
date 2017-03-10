import React, {Component} from 'react'
import ReactDOM from 'react-dom'


export default class Index extends Component {
  constructor(props) {
    super(props);
    this.state = {

    };

    this.actions = {

    };
  }

  render() {
    return (
      <div className="index">
        <h1>Sample Page</h1>
      </div>
    );
  }
}

Index.propTypes = {

};


window.MyApp = {
  init: function (opts) {
    var mountPoint = opts.mount;
    var config = opts.props;
    ReactDOM.render(React.createFactory(Index)(config), document.getElementById(mountPoint));
  }
};
