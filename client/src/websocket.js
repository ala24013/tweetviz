import React from "react";
import useWebSocket from "react-use-websocket";

const Websocket = (props) => {
  const socketUrl = `ws://${document.location.host}/ws`;

  /* eslint-disable no-unused-vars */
  const {
    sendMessage,
    sendJsonMessage,
    lastMessage,
    lastJsonMessage,
    readyState,
    getWebSocket,
  } = useWebSocket(socketUrl, {
    onOpen: () => console.log("opened ws"),
    onMessage: (msg) => {
      const data = JSON.parse(msg.data);
      props.setTweets(data);
    },
    shouldReconnect: (closeEvent) => true,
    onClose: () => console.log("closed ws")
  });
  /* eslint-enable no-unused-vars */

  return <div></div>;
};

export default Websocket;
