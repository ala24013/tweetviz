import React from 'react';
//import World from './world';
import Map from './map'
import Websocket from './websocket';

function App() {
  return (
    <div className="App">
      <body>
        <Map />
        <Websocket />
      </body>
    </div>
  );
}

export default App;
