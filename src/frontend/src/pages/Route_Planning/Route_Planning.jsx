import React, {useEffect} from 'react';
import styles from './Route_Planning.module.css';
import { Form, Button } from "react-bootstrap";
import axios from 'axios';
import {Data} from '../../structs/data'
import Graph from "react-graph-vis";
import {v4 as uuidv4} from 'uuid';
import { getOverlayDirection } from 'react-bootstrap/esm/helpers';



class Route_Planning extends React.Component {  
    constructor(props){
      super(props);
      this.state = {
        submit: false,
        validinput: false,
        isLoaded : false,
        isUCS : true,
        isClicked : false,
        text : "",
        start : "",
        dest : "",
        the_data : new Data(),
        graph : {
          nodes :[],
          edges : []
        }
      };
      
    }


    handleSubmit = async (e) =>{
      e.preventDefault(); 
       
        let x = this.state.the_data.getPlaces()
        if(!(x.includes(this.state.start) && x.includes(this.state.dest))){
          alert("Titik awal atau tujuan yang anda masukkan salah!")
          console.log("ngaco")
        }
        else{
          this.setState(
            {
              submit: true
            }
          )
             
          const res =  await axios.post("http://localhost:8000/algo",{start:this.state.start,dest:this.state.dest,isUCS:this.state.isUCS})
          .then(response => {
            console.log(response.data);
          })
          .catch(error => {
            console.error(error);
          });

          const respons = await axios.get("http://localhost:8000/algo");

          let data_dum = this.state.the_data
          data_dum.setPath(respons.data.path)
          data_dum.setDist(respons.data.id)

          let graph = this.state.graph
          let path = respons.data.path
          console.log(path)
          for(let i = 0; i < path.length-1; i++){
            graph.edges.forEach(
              edge =>{
                if(edge.from == path[i] && edge.to == path[i+1]){
                  graph.nodes[path[i]].color =  "#7be041"
                  graph.nodes[path[i+1]].color =  "#7be041"
                  edge.color = "#7be041"
                  edge.arrows = {to: true, from: false}
                  edge.width = 3
                }
              }
            )
          }

          console.log(graph)
          
        
          this.setState(
            {
              the_data : data_dum,
              graph : graph
            }
          )
        }
     
    }

    handleSubmitButton(){
      if(this.state.validinput && this.state.isClicked){
        return false;
      }
      else{
        return true;
      }
    }
  
    
    handleReadText = async (e) =>{
     
      
      const reader = new FileReader();
      reader.onload = async (e) => {
        this.setState((state, props) => ({
          text : e.target.result,
          validinput : this.validateText(e.target.result),
          isLoaded : true,
          submit : false
        }));
      };
      
      reader.readAsText(e.target.files[0]);
     
      // Upload the file to the server
      const formData = new FormData();
      formData.append("file", e.target.files[0]);
      const response = axios.post("http://localhost:8000/upload", formData)
       
    }

    handleStart = (e) =>{
      this.setState({
        start : e.target.value
      })
    }

    handleDest = (e) =>{
      this.setState({
        dest : e.target.value
      })
    }

    handleRadio= (e) =>{
      this.setState({
        isUCS : (e.target.value == "UCS"),
        isClicked : true
      })
    }

   

    validateText(txt){
      const buff = txt.split(/\s+/);

      buff.pop() //delete empty string
  
      //containts the buffer of the data
      let text_data =  new Data()
      let places = null
      let matrix = null

      //algorithm

      let incr = 0;
      let node_ctr = 0;
      let matrix_ctr = 0;
      let num = NaN;
      const node_atr = 3;

      for(const kata of buff){
        
       
        if(incr == 0){
          
          num = parseInt(kata,10);
          
          if(isNaN(num)){
            console.log("bukan number")
            return false;
          } 

          incr++;
          //initiate matrix and places
          matrix = new Array(num)
          for (let i = 0; i < num; i++) {
            matrix[i] = new Array(num).fill(0);
          }
          places = new Array(num)
          
          
        }

        else{
            if(node_ctr < num*node_atr){ //read each element to retrieve heruistic information
              if(node_ctr % node_atr == 0){
                //Read string
                const buffnode =  parseFloat(kata);
                if(!isNaN(buffnode)) return false;
                places[node_ctr/node_atr] = kata;
              }
              else{
                const coordinate = parseFloat(kata);
                if(isNaN(coordinate)) return false;
              }
              node_ctr++;
            }
            else{ //read the adjacency matrix
              if(matrix_ctr < num*num){
                const buffmtx = parseFloat(kata);
                if(isNaN(buffmtx)) return false;
                matrix[Math.floor(matrix_ctr/num)][matrix_ctr % num] = buffmtx
                matrix_ctr++;
              }
              else{
                return false;
              }
            }
        }
        
      }

      //if matrix is fulllfiled
      if(matrix_ctr == num*num){
        //check the matrix symetry
        for (let i = 0; i < num; i++) {
          for(let j = i; j < num; j++){
              if(i == j && matrix[i][j]!= 0) return false;
              else if(matrix[i][j] != matrix[j][i]) return false;
          }
        }

        
        //finally
        text_data.setPlaces(places)
        text_data.setMatrix(matrix)

        let graf = {
          nodes :[],
          edges : []
        }

        for(let i = 0; i < matrix.length; i++){
          graf.nodes.push({id: i, label: places[i],color :"#e04141"})
        }

        for(let i = 0; i < matrix.length; i++){
          for(let j = i+1; j < matrix.length; j++){
            if(matrix[i][j] != 0){
              graf.edges.push({from:i,to:j, color :"#e04141", label:matrix[i][j].toString(), arrows: { to: false, from:false }, width:1})
              graf.edges.push({from:j,to:i, color :"#e04141", arrows: { to: false, from: false}, width:1})
            }
          }
        }

        this.setState({
          graph :graf,
          the_data : text_data
        })
        
        return true;
      }
      else{
        return false;
      }
     
    }

    getRoute(path,dist){
      if(path != null && dist != 0){
        return "Rute: " + path.toString() + ", Jarak tempuh: " + dist.toString()
      }
      
    }

    rendermap(){
      if(!this.state.isLoaded){ //Not loaded anything
        return(
          <div></div>
        );
      }

      else if(this.state.validinput && !this.state.submit){  
        return(
          <div class = {styles.map}>
            <div class = {styles.mapborder}>
              
              <Graph graph = {this.state.graph} style={{ height: "400px" }}/>
            </div>
          </div>
          
          
        )
      }
      else if(this.state.validinput && this.state.submit){
       
        let graf = {
          nodes: [],
          edges : []
        }

        for(let x = 0; x < this.state.graph.nodes.length; x++){
          graf.nodes[x] = {id: this.state.graph.nodes[x].id, label: this.state.graph.nodes[x].label, color: this.state.graph.nodes[x].color}
        }
        for(let x = 0; x < this.state.graph.edges.length; x++){
          graf.edges[x] = { from:this.state.graph.edges[x].from, 
                            to:this.state.graph.edges[x].to, 
                            label: this.state.graph.edges[x].label, 
                            color: this.state.graph.edges[x].color,
                            arrows : this.state.graph.edges[x].arrows,
                            width : this.state.graph.edges[x].width
                          }
        }



        return(
          <div class = {styles.map}>
            <div class = {styles.mapborder}>
              
              <Graph graph = {graf} style={{ height: "400px" }}/>
             
            </div>
            <p1>{this.getRoute(this.state.the_data.path,this.state.the_data.dist)}</p1>
          </div>
          
          
        )
      }
      else{
        return(
          <div class = {styles.invalid}>
            <h1>Input file tidak valid!</h1>
          </div>
        )
      }
    }
    
    

    render() { 
      return (
        
        <div className="Route_Planning">

          {/* create Route_Planning title  */}
          <div className={styles.title}>
            <h1>Route Planning!</h1>
          </div>
                  
          {/* Container */}
          <div className={styles.container}>
            
            <div class="d-grid gap-3">
              <Form onSubmit = {this.handleSubmit}>
                <div className="Form">
                  
                    
                    
                    <Form.Group  controlId="formMap">
                      <Form.Label className>Upload File</Form.Label>
                      <Form.Control
                        type="file"
                        placeholder="Enter file"
                        onChange = {this.handleReadText}
                      />
                    </Form.Group>
                    
                    <Form.Group controlId="formStartNode">
                      <Form.Label >Start </Form.Label>
                      <Form.Control
                        type="text"
                        placeholder="contoh: 'Ciumbuleuit'"
                        onChange = {this.handleStart}
                      />
                    </Form.Group>


                    <Form.Group controlId="formDestinationNode" >
                      <Form.Label >Destination: </Form.Label>
                      <Form.Control
                        type="text"
                        placeholder="contoh: 'Dago'"
                        onChange = {this.handleDest}                    
                      />
                    </Form.Group>
                  
                </div>

                {/* create radio */}
                <div className={styles.radio}>
                  <div class="d-grid gap-1">
                    <div class="form-check">
                      <input
                        type="radio"
                        class="form-check-input"
                        id="radio1"
                        name="optradio"
                        value="UCS"
                        onClick={this.handleRadio}
                        
                      />
                      UCS (Uniform Cost Search) 
                      <label class="form-check-label" for="radio1"></label>
                    </div>
                    <div class="form-check">
                      <input
                        type="radio"
                        class="form-check-input"
                        id="radio2"
                        name="optradio"
                        value="A*"
                        onClick={this.handleRadio}
                      />
                      A* 
                      <label class="form-check-label" for="radio2"></label>
                    </div>
                  </div>
                </div>

                <div className={styles.submit}>
                  <div class="col-md-12 text-center">
                    <Button variant="success" type="submit" disabled = {this.handleSubmitButton()}>
                      <strong>Get route!</strong>
                    </Button>
                  </div>
                </div>
              </Form>
            </div>
          </div>
        {this.rendermap()}
        </div>
          );
     }
}
     
     
  export default Route_Planning;