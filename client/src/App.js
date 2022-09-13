import React from 'react';
import './App.css';
import World from './world';
import Websocket from './websocket';

function App() {
  return (
    <div className="App">
      <body>
        <World />
        <Websocket />
      </body>
    </div>
  );
}

export default App;
