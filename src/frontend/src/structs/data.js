


export class Data{
    constructor() {
      this.places = null // array of place
      this.matrix = null // adjacency matrix
      this.path = null // array of path
      this.dist = 0
    }
  
    // setters
    setPlaces(places) {
      this.places = places
    }
  
    setMatrix(matrix) {
      this.matrix = matrix
    }
  
    setPath(path) {
      this.path = path
    }
  
    setDist(dist) {
      this.dist = dist
    }
  
    // getters
    getPlaces() {
      return this.places
    }
  
    getMatrix() {
      return this.matrix
    }
  
    getPath() {
      return this.path
    }
  
    getDist() {
      return this.dist
    }

    
  }
