import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      health: "",
      user: {}
    };
  }

  componentDidMount() {
    fetch("http://localhost:3001/health")
    .then(result => result.text())
    .then(data => {
       this.setState({ health: data });
    });

    fetch("http://localhost:3001/user")
    .then(result => result.json())
    .then(data => {
       this.setState({ user: data});
    });
  }

  render() {
    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <h1 className="App-title">Server status: {this.state.health}</h1>
        </header>
        <h2>Username: {this.state.user.username}</h2>
          <h2>Email: {this.state.user.email}</h2>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
      </div>
    );
  }
}

export default App;
