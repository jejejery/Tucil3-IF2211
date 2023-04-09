import React from 'react';
import styles from './Route_Planning.module.css';
import { Form, Button, Card } from "react-bootstrap";



class Route_Planning extends React.Component {  
    constructor(){
        this.target
    };


    
    handleRadioButton(e){
        setMetode(e.target.value);
    };

    render() { 
        const axios = require("axios");
        return (
                <div className="Route_Planning">
                {/* create Route_Planning variant dark */}
                    
                     {/* title "Cek DNA Pasien" */}
          <div className={styles.container}>
            <div class="d-grid gap-3">
              <h1>Route Planning!</h1>
              <Form onSubmit>
                <div className={styles.form}>
                  <div class="d-grid gap-3">
                    
                    <Form.Group controlId="formBasicNamaPasien">
                      <Form.Label>Nama Pasien</Form.Label>
                      <Form.Control
                        type="text"
                        placeholder="contoh: Gagas Praharsa Bahar"
                        
                      />
                    </Form.Group>

                    <Form.Group controlId="formBasicSequenceDNA">
                      <Form.Label>Upload Sequence DNA</Form.Label>
                      <Form.Control
                        type="file"
                        placeholder="Enter file"
                       
                      />
                    </Form.Group>

                    <Form.Group controlId="formBasicNamaPenyakit">
                      <Form.Label>Nama Penyakit</Form.Label>
                      <Form.Control
                        type="text"
                        placeholder="contoh: HIV"
                        
                      />
                    </Form.Group>
                  </div>
                </div>

                {/* create radio */}
                <div className={styles.radio}>
                  <div class="d-grid gap-3">
                    <div class="form-check">
                      <input
                        type="radio"
                        class="form-check-input"
                        id="radio1"
                        name="optradio"
                        value="KMP"
                        onClick={handleRadioButton}
                        checked
                      />
                      KMP
                      <label class="form-check-label" for="radio1"></label>
                    </div>
                    <div class="form-check">
                      <input
                        type="radio"
                        class="form-check-input"
                        id="radio2"
                        name="optradio"
                        value="BM"
                        onClick={handleRadioButton}
                      />
                      Boyer Moore
                      <label class="form-check-label" for="radio2"></label>
                    </div>
                  </div>
                </div>

                <div className={styles.submit}>
                  <div class="col-md-12 text-center">
                    <Button variant="success" type="submit">
                      <strong>Submit</strong>
                    </Button>
                  </div>
                </div>
              </Form>
            </div>
          </div>

                </div>
          );
     }
}
     
     
  export default Route_Planning;