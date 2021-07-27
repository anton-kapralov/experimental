import React from 'react';
import {Container, Header, Input, Segment} from "semantic-ui-react";

const style = {
  h1: {
    marginTop: '3em',
  },
}

const Sum = () => {
  return <Container textAlign='center'>
    <Input label='a' placeholder='0' size='huge' />
    +
    <Input label='b' placeholder='0' size='huge' />
    =
    <Input label='c' placeholder='0' size='huge' />
  </Container>
}

class App extends React.Component {
  render() {
    return (
      <div>
        <Header as='h1' content='Calculator' style={style.h1} textAlign='center'/>
        <Sum/>
      </div>
    )
  }
}

export default App
