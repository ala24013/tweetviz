import React, { useEffect } from 'react'
import { MapContainer, TileLayer, useMap, Popup, Marker, Tooltip } from 'react-leaflet'

export default function Tweet(props) {

    useEffect(() => {
        setTimeout(
            () => { props.removeItem(props.username, props.tweet) },
            Math.random() * 5000 + 5000 // between 5 and 10 seconds
        )
    })

    return (
        <Marker className={props.username + props.tweet} position={props.geo}>
            <Tooltip direction="top" offset={[20,0]} opacity={1} permanent>
                {`${props.username} : ${props.tweet}`}
            </Tooltip>
        </Marker>
    )
}