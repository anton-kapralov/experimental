import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import 'semantic-ui-css/semantic.min.css'
import {Container, Header, Input, Segment} from "semantic-ui-react";

const style = {
  h1: {
    marginTop: '3em',
  },
}

class App extends React.Component {
  render() {
    return (
      <div>
        <Header as='h1' content='Calculator' style={style.h1} textAlign='center'/>

        <Container textAlign='center'>
          <Input label='a' placeholder='0' size='huge' className='number'/>
          +
          <Input label='b' placeholder='0' size='huge'/>
          =
          <Input label='c' placeholder='0' size='huge'/>
        </Container>
      </div>
    )
  }
}

ReactDOM.render(
  <React.StrictMode>
    <App/>
  </React.StrictMode>,
  document.getElementById('root')
);
