import React, { useState, useCallback } from "react";
import { MantineProvider } from '@mantine/core';
import useWebSocket from "react-use-websocket";

import WorldMap from "./worldMap";
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

  const socketUrl = `ws://${document.location.host}/ws`;

  /* eslint-disable no-unused-vars */
  const {sendMessage} = useWebSocket(socketUrl, {
    onOpen: () => console.log("opened ws"),
    onMessage: (msg) => {
      const data = JSON.parse(msg.data);
      setTweets(data);
    },
    shouldReconnect: (closeEvent) => true,
    onClose: () => console.log("closed ws")
  });
  /* eslint-enable no-unused-vars */

  return (
    <MantineProvider theme={theme} styles={styles}>
      <div className="App">
        <body>
          <TweetvizHeader sendMessage={(msg) => sendMessage(msg)} />
          <WorldMap tweets={tweets} />
        </body>
      </div>
    </MantineProvider>
  );
}

export default App;
