import React from "react";

export default class BasicForm extends React.Component {
  constructor(props) {
    super(props);
    this.state = {Name: '', Address: ''};
    

    this.handleChange = this.handleChange.bind(this);
    
    this.handleSubmit = this.handleSubmit.bind(this);
   
    this.handleListSubmit = this.handleListSubmit.bind(this);

    this.handleClosestSubmit = this.handleClosestSubmit.bind(this);
  }

  handleChange(event) {
    this.setState({[event.target.name]: event.target.value});
  }

  handleSubmit(event) {
    let data = this.state
    this.props.postSubmit(data);
    event.preventDefault();
  }

  handleClosestSubmit(event) {
    let data = this.state
    this.props.closestSubmit(data);
    event.preventDefault();
   
  }

  handleListSubmit(event) {
    event.preventDefault();
    this.props.listSubmit();
  }

 


  render() {
    return (
      <div >
        < form >
          <label htmlFor = "name" > Name:
            <input type = "text" name = "Name" value ={this.state.Name || ''} onChange ={this.handleChange} / >
          </label> 
          <label htmlFor = "address" > Address:
            < input type ="text" name= "Address" value ={this.state.Address || ''} onChange ={this.handleChange}/ >
          </label> 
         
          <input type = "submit" value = "Get Playground" onClick={this.handleSubmit} / >
          <input type = "submit" value = "List Playground" onClick={this.handleListSubmit} / >
           <input type = "submit" value = "Closest Playground" onClick={this.handleClosestSubmit} / >
          
        </form> 
      </div>
    )
  }
}

