import React, { useState } from "react";
import { MantineProvider } from '@mantine/core';

import WorldMap from "./worldMap";
import Websocket from "./websocket";
import TweetvizHeader from "./tweetvizHeader";

const PRIMARY_COLOR = "#1D9BF0"

const styles = {
  Title: (theme) => ({
    root: {
      color: PRIMARY_COLOR
    }
  })
}

const theme = {
  colorScheme: 'dark',
  headings: {
    fontFamily: 'Roboto, sans-serif',
    sizes: {
      h1: { fontSize: 30 },
    },
    colors: {
      primary: [ PRIMARY_COLOR ]
    },
    primaryColor: 'primary',
}}

function App() {
  const [tweets, setTweets] = useState([]);

  return (
    <MantineProvider theme={theme} styles={styles}>
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
