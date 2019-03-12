import React from "react";
import ReactDOM from "react-dom";
import BasicForm from "./components/BasicForm.js"

export default class App extends React.Component {

  constructor() {
    super();

    this.state = {
      listObj: '',
      getObj: '',
      closestObj: ''
    };

    this.postSubmit = this.postSubmit.bind(this);
    this.listSubmit = this.listSubmit.bind(this);
    this.closestSubmit = this.closestSubmit.bind(this);
  };

  postSubmit(data) {
    let url = "http://localhost:8000/playground/GET"
   
    return fetch(url, {
        method: "POST",
        headers: {
          'Accept': 'application/json',
          "Content-Type": "application/x-www-form-urlencoded",
        },
        body: JSON.stringify(data)
      })
      .then(response => {
        return response.json()
      })
      .then((json) => {
        this.setState({ getObj: json })
      })      
  }

  listSubmit() {
    let url = "http://localhost:8000/playground/LIST"
    
    return fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
      })
      .then(response => {
        return response.json()
      })
      .then((json) => {
        this.setState({ listObj: json })
      })        
  }

  closestSubmit(data) {
    
    let url = "http://localhost:8000/playground/CLOSEST"
  
    return fetch(url, {
        method: "POST",
        headers: {
          'Accept': 'application/json',
          "Content-Type": "application/x-www-form-urlencoded",
        },
        body: JSON.stringify(data)
      })
      .then(response => {
        return response.json()
      })
      .then((json) => {
        this.setState({ closestObj: json })
      })        
  }

  renderObjectlist () {
    if (this.state.listObj) {
      return (
        <ul>
          {this.state.listObj.map((obj, i) => <li key={i}>
          Name: {obj.Name}, Address: {obj.Address}, Latitude: {obj.Latitude }, Longitude: {obj.Longitude}</li>)}
          
        </ul>
      )
    }   
  }

  renderClosestlist () {
    if (this.state.closestObj) {
      return (
        <ul>
          {this.state.closestObj.map((obj, i) => <li key={i}>
          Name: {obj.Name}, Address: {obj.Address}, Latitude: {obj.Latitude }, Longitude: {obj.Longitude}</li>)}
          
        </ul>
      )
    }
    
  }

  renderGetPlayground() {
    if (this.state.getObj) {
      return (
        <ul>
          Name : {this.state.getObj.Name}
          Address : {this.state.getObj.Address}
          Latitude: {this.state.getObj.Latitude}
          Longitude: {this.state.getObj.Longitude}
        </ul>
      )
    }
   
  }



  render() {
    return (
       <div> 
        <BasicForm postSubmit={this.postSubmit} listSubmit={this.listSubmit} closestSubmit={this.closestSubmit}/>
        <div>
          <h1> Specific Playground </h1>
          <div>{this.renderGetPlayground()}</div>
          <br/>
          <h1> List of Playground </h1>
          <div>{this.renderObjectlist()}</div>
          <h1> Closest Playground </h1>
          <div>{this.renderClosestlist()}</div>
        </div>

      </div>
    );
  }

}

const rootElement = document.getElementById("root");
ReactDOM.render( < App / > , rootElement);