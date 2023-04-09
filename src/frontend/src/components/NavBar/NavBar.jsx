import React from 'react';
import styles from './NavBar.module.css';
import { FaHome, FaRoute, FaQuestionCircle} from "react-icons/fa";


class NavBar extends React.Component {  
    
    render() { 

        return (
                <div className={styles.navbar}>
                {/* create navbar variant dark */}
                    <nav class="navbar navbar-expand-lg bg-success navbar-dark">
                        <div class={styles.container_fluid}>
                            <a class="navbar-brand ms-4" href="/" >
                            {/* <img src={logo} width="50" height="50" class="d-inline-block align-top" alt=""/> */}
                                <strong> <FaHome /> Home</strong>
                            </a>
                            <ul class="navbar-nav ms-auto mb-2 mb-lg-0">
                                <li class={styles.nav_items}>
                                    <a class="nav-link" href="/route_planning"><strong><FaRoute/> Route Planning</strong></a>
                                </li>
                                <li class={styles.nav_items}>
                                    <a class="nav-link" href="/about"><strong><FaQuestionCircle/> About Program </strong></a>
                                </li>
                            </ul>
                        </div>
                    </nav>
                </div>
          );
     }
}
     
     
  export default NavBar;