import React from 'react';

class Sum extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      a: 0,
      b: 0,
      c: 0,
    }
  }

  render() {
    const {a, b, c} = this.state
    return (
      <div className="ui stacked">
        <div className="ui labeled huge input">
          <div className="ui label">a</div>
          <input type="text" placeholder="0" name="a" value={a} onChange={this.handleChange}/>
        </div>
        <div className="ui circular huge label">+</div>
        <div className="ui labeled huge input">
          <div className="ui label">b</div>
          <input type="text" placeholder="0" name="b" value={b} onChange={this.handleChange}/>
        </div>
        <button className="ui secondary huge button" style={{margin: "8px"}} onClick={this.calculate}>=</button>
        <div className="ui labeled huge input">
          <div className="ui label">c</div>
          <input type="text" placeholder="0" readOnly={true} value={c}/>
        </div>
      </div>
    )
  }

  handleChange = (event) => {
    const { name, value } = event.target;

    this.setState({
      [name] : parseInt(value)
    });
  }

  calculate = () => {
    const {a, b} = this.state
    this.setState({
      a: a,
      b: b,
      c: a + b
    })
  }
}

class App extends React.Component {
  render() {
    return (
      <div className="main ui container">
        <h1 className="ui dividing centered header">Sum</h1>
        <div id="content">
          <Sum/>
        </div>
      </div>
    )
  }
}

export default App
