import React from 'react';
import useWebSocket from 'react-use-websocket';

const Websocket = () => {
    const socketUrl = `ws://${document.location.host}/ws`;

    const {
        sendMessage,
        sendJsonMessage,
        lastMessage,
        lastJsonMessage,
        readyState,
        getWebSocket,
    } = useWebSocket(socketUrl, {
        onOpen: () => console.log('opened'),
        onMessage: (msg) => console.log(msg),
        shouldReconnect: (closeEvent) => true,
    })

    return (
        <div></div>
    )
}

export default Websocket