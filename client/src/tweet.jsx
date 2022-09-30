import React, { useEffect } from "react";
import {
  MapContainer,
  TileLayer,
  useMap,
  Popup,
  Marker,
  Tooltip,
} from "react-leaflet";

import useTimeout from "./useTimeout";

export default function Tweet(props) {
  const fadeOutCss = {
    visibility: "hidden",
    opacity: 0,
    transition: "visibility 0s 2s, opacity 2s linear",
  };

  useEffect(() => {
    console.log(props.username);
    console.log(props.tweet);
    console.log(props.geo);
  });

  return (
    <Marker
      className={props.username + props.tweet}
      position={props.geo}
      style={fadeOutCss}
    >
      <Tooltip direction="top" offset={[0, 0]} permanent>
        {`${props.username} : ${props.tweet}`}
      </Tooltip>
    </Marker>
  );
}
