import React, { useEffect, useState } from 'react'
import { MapContainer, TileLayer, useMap, Popup, Marker, Tooltip } from 'react-leaflet'

import Tweet from './tweet'

const INITIAL_STATE = [
    {username: "hello", geo: [0,0], tweet: "thetweet"},
    {username: "hi", geo: [50,50], tweet: "thetweet2"},
    {username: "hi2", geo: [45,45], tweet: "thetweet22"},
    {username: "hi3", geo: [55,55], tweet: "thetweet23"},
    {username: "hi4", geo: [60,60], tweet: "thetweet24"}
]

export default Map = (props) => {
    const [tweets, setTweets] = useState(INITIAL_STATE);

    const removeItem = (username, tweet) => {
        console.log("REMOVING")
        setTweets(
            tweets.filter((t) => (
                t.username != username ||
                t.tweet != tweet
            ))
        )
    }

    return (
        <MapContainer center={[45.4, -75.7]} zoom={10} scrollWheelZoom={true}>
            <TileLayer
                url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
            />
            { tweets.map((t) => <Tweet username={t.username} geo={t.geo} tweet={t.tweet} removeItem={removeItem} />) }
        </MapContainer>
    );
}