import React from "react";
import { MapContainer, TileLayer } from "react-leaflet";

import Tweet from "./tweet";

export default function WorldMap(props) {
  return (
    <MapContainer center={[41.9028, 12.4964]} zoom={5} scrollWheelZoom={true}>
      <TileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
      />
      {props.tweets.map((t) => (
        <Tweet username={t.Username} geo={t.Geo} tweet={t.Tweet} />
      ))}
    </MapContainer>
  );
};
