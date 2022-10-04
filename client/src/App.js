import React, { useState } from "react";
import { MantineProvider } from '@mantine/core';

import Map from "./map";
import Websocket from "./websocket";
import SearchBar from "./searchBar";

function App() {
  const [tweets, setTweets] = useState([]);

  return (
    <MantineProvider withGlobalStyles withNormalizeCSS>
      <div className="App">
        <body>
          <SearchBar />
          <Map tweets={tweets} />
          <Websocket setTweets={setTweets} />
        </body>
      </div>
    </MantineProvider>
  );
}

export default App;
