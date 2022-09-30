import React from "react";
import useWebSocket from "react-use-websocket";

const Websocket = (props) => {
  const socketUrl = `ws://${document.location.host}/ws`;

  const {
    sendMessage,
    sendJsonMessage,
    lastMessage,
    lastJsonMessage,
    readyState,
    getWebSocket,
  } = useWebSocket(socketUrl, {
    onOpen: () => console.log("opened"),
    onMessage: (msg) => {
      const data = JSON.parse(msg.data);
      props.setTweets(data);
    },
    shouldReconnect: (closeEvent) => true,
  });

  return <div></div>;
};

export default Websocket;
