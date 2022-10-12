import React, { useState } from "react";
import { MantineProvider } from '@mantine/core';
import useWebSocket from "react-use-websocket";

import WorldMap from "./worldMap";
import TweetvizHeader from "./tweetvizHeader";
import { theme, styles } from "./theme"

function App() {
  const [tweets, setTweets] = useState([]);
  const [loading, setLoading] = useState(false);

  const socketUrl = `ws://${document.location.host}/ws`;

  /* eslint-disable no-unused-vars */
  const {sendMessage} = useWebSocket(socketUrl, {
    onOpen: () => console.log("opened ws"),
    onMessage: (msg) => {
      const data = JSON.parse(msg.data);
      if (data.code === "tweetlist") {
        const tweets = JSON.parse(data.msg)
        setTweets(tweets); 
      } else if (data.code === "doneLoading") {
        setLoading(false)
      } else {
        console.log(`Unexpected code: ${data.code}`)
      }
    },
    shouldReconnect: (closeEvent) => true,
    onClose: () => console.log("closed ws")
  });
  /* eslint-enable no-unused-vars */

  const send = (msg) => {
    setLoading(true);
    sendMessage(msg)
  }

  return (
    <MantineProvider theme={theme} styles={styles}>
      <div className="App">
        <body>
          <TweetvizHeader sendMessage={send} loading={loading} />
          <WorldMap tweets={tweets} loading={loading} />
        </body>
      </div>
    </MantineProvider>
  );
}

export default App;
