import React, { useState } from 'react';
//import World from './world';
import Map from './map'
import Websocket from './websocket';

const INITIAL_STATE = [
  {Username: "hello", Geo: [0,0], Tweet: "thetweet"},
  {Username: "hi", Geo: [50,50], Tweet: "thetweet2"},
  {Username: "hi2", Geo: [45,45], Tweet: "thetweet22"},
  {Username: "hi3", Geo: [55,55], Tweet: "thetweet23"},
  {Username: "hi4", Geo: [60,60], Tweet: "thetweet24"}
]

function App() {
  const [tweets, setTweets] = useState(INITIAL_STATE);

  return (
    <div className="App">
      <body>
        <Map tweets={tweets} />
        <Websocket setTweets={setTweets} />
      </body>
    </div>
  );
}

export default App;
