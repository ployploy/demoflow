import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      health: "",
      user: {},
      createResult: ""
    };

    this.createUser = this.createUser.bind(this);
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

  createUser(){
    fetch('http://localhost:3001/user', {
      method: 'POST',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        username: 'ploynaka',
        email: 'ploynaka@gmail.com',
      })
    })
    .then(result => result.text())
    .then(data => {
      console.log(data);
      this.setState({ createResult: data });
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
        <button onClick={this.createUser}>Create user</button>
        <h2>{this.state.createResult}</h2>
        <p className="App-intro">
          To get started, edit <code>src/App.js</code> and save to reload.
        </p>
      </div>
    );
  }
}

export default App;
