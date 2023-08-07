import React, { useState, useEffect } from 'react';
import './App.css';

function App() {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch("http://localhost:8000/temperature")
      .then(response => {
        if (!response.ok) {
          throw new Error("Failed to fetch data");
        }
        return response.json();
      })
      .then(data => setData(data))
      .catch(error => setError(error));
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        {error && <p>Error: {error.message}</p>}
        {data && (
          <div>
            <h1>Temperature Data</h1>
            <p>City: {data.city}</p>
            <p>Timestamp: {data.timestamp}</p>
            <p>Degree: {data.degree}°C</p>
          </div>
        )}
      </header>
    </div>
  );
}

export default App;