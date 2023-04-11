import React from 'react';
import styles from './About.module.css';
import Table from 'react-bootstrap/Table';


class About extends React.Component {  
    
    render() { 

        return (
        <div class = "about">
            <div class={styles.author}>

            <h1>About author: </h1>
            <Table striped bordered hover variant="dark" className="w-auto">
            <thead>
              <tr>
                
                <th>No</th>
                <th>Nama</th>
                <th>NIM</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>1</td>
                <td>Jeremya Dharmawan R</td>
                <td>13521131</td>
              </tr>
              <tr>
                <td>2</td>
                <td>Antonio Natthan Khrisna</td>
                <td>13521162</td>
              </tr>
            </tbody>
          </Table>
            </div>
            
            <div class = {styles.desc}>
                <p1>Tugas Kecil ini Merupakan Implementasi Algoritma UCS(Uniform Cost Search) dan A* (A-Star) dalam penentuan rute pada peta yang direpresentasikan dalam bentuk graf</p1>
            </div>
            
               
        </div>
          );
     }
}
     
     
  export default About;