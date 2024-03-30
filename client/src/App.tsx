import { useState, useEffect } from 'react';
import reactLogo from './assets/react.svg';
import viteLogo from '/vite.svg';
import './App.css';

function App() {
  const [count, setCount] = useState(0);

  useEffect(() => {
    // Check if the user is already logged in when the component mounts
    // This could be a simple check to a backend endpoint that validates the session
    // For this example, we are not implementing it, but it's a recommended approach
  }, []);

  const handleLogin = () => {
    // Redirect to the backend endpoint for Google login
    // This is a simplification, your actual login flow might need to handle
    // redirections and responses differently
    window.location.href = "http://localhost:3000/api/auth/google";
  };

  const handleApiCall = async () => {
    try {
      const response = await fetch('http://localhost:3000/api', {
        credentials: 'include' // Important: Include cookies with the request
      });
      const data = await response.json();
      console.log(data);
    } catch (error) {
      console.error('Error fetching API:', error);
    }
  };

  return (
    <>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>Edit <code>src/App.tsx</code> and save to test HMR</p>
      </div>
      <div>
        <button onClick={handleLogin}>Google Login</button>
      </div>
      <div>
        <button onClick={handleApiCall}>Call API</button>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App;
