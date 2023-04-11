import React, { useState } from "react";
import axios from "axios";

function App() {
  const [message, setMessage] = useState("");

  const handleClick = async () => {
    const response = await axios.get("http://localhost:8000/message");
    setMessage(response.data.message);
  };

  return (
    <div>
      <button onClick={handleClick}>Get Message from Backend</button>
      <p>{message}</p>
    </div>
  );
}

export default App;