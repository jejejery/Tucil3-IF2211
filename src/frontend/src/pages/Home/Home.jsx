import React from 'react';
import styles from './Home.module.css';
import img_map from '../../attribute/map.gif'
import { FaHome, FaRoute, FaQuestionCircle} from "react-icons/fa";

console.log(img_map);

class Home extends React.Component {  
    render() {   
        return (
          
          <div class="home">

            <div class = {styles.title}> 
            <h1 id="h1" >Selamat datang di Aplikasi Route Planning!</h1> 
            </div>
            

            <div class={styles.img}>
              <img id = "gambar" alt="logo aplikasi" src={img_map}/>
            </div>
            

            <div class = {styles.intro}>

            
            <h3 id="p0">Beberapa fitur dari aplikasi ini antara lain:</h3>
            <p id="p1">1. Menerima input file graf ketetanggaan</p>
            <p id="p2">2. Melakukan pencarian rute</p>
            <p id="p3">3. Menampilkan hasil pencarian</p>
            <div class = {styles.outro}>
            <h2 id="h2">Pilih fitur yang tersedia pada navigation bar di atas</h2>
            <h2 id="h3">Semoga aplikasi ini bermanfaat! 😃</h2>
            </div>
            
          
            </div>

            
            

           
            
        </div>
          );
     }
}
     
     
  export default Home;