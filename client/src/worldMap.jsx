import { Loader } from "@mantine/core";
import React from "react";
import { MapContainer, TileLayer } from "react-leaflet";

import Tweet from "./tweet";
import { HEADER_HEIGHT } from "./tweetvizHeader";

export default function WorldMap(props) {
  const height = window.innerHeight - HEADER_HEIGHT;
  const opacity = props.loading ? 0.5 : 1
  return (
    <MapContainer style={{height: height}} center={[41.9028, 12.4964]} zoom={5} scrollWheelZoom={true}>
      <TileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
        opacity={opacity}
      />
      {props.loading ?
        <div style={{display: "flex", flexDirection: "row", alignContent: 'center', height: '100%', justifyContent: 'space-around'}} >
          <Loader size="xl" variant="bars" />
        </div> :
        <></>
      }
      {props.tweets.map((t) => (
        <Tweet username={t.Username} geo={t.Geo} tweet={t.Tweet} />
      ))}
    </MapContainer>
  );
};
