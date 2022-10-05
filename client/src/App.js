import React, { useState } from "react";
import { MantineProvider } from '@mantine/core';

import WorldMap from "./worldMap";
import Websocket from "./websocket";
import TweetvizHeader from "./tweetvizHeader";

const theme = {
  colorScheme: 'dark',
  headings: {
    fontFamily: 'Roboto, sans-serif',
    sizes: {
      h1: { fontSize: 30 },
    },
}}

function App() {
  const [tweets, setTweets] = useState([]);

  return (
    <MantineProvider theme={theme}>
      <div className="App">
        <body>
          <TweetvizHeader />
          <WorldMap tweets={tweets} />
          <Websocket setTweets={setTweets} />
        </body>
      </div>
    </MantineProvider>
  );
}

export default App;
