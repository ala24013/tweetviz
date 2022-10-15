import { Loader, Container, Title, Text } from "@mantine/core";
import React from "react";
import { MapContainer, TileLayer } from "react-leaflet";

import Tweet from "./tweet";
import { HEADER_HEIGHT } from "./tweetvizHeader";
import { PRIMARY_COLOR } from "./theme"

export default function WorldMap(props) {
  const height = window.innerHeight - HEADER_HEIGHT;
  const opacity = props.loading ? 0.25 : 1
  return (
    <MapContainer style={{height: height}} center={[41.9028, 12.4964]} zoom={5} scrollWheelZoom={true}>
      <TileLayer
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
        opacity={opacity}
      />
      {props.loading ?
        <div style={{display: "flex", flexDirection: "row", alignContent: 'center', height: '100%', justifyContent: 'space-around'}} >
          <Container style={{display: "flex", flexDirection: "column", justifyContent: "space-around"}}>
            <Container style={{marginTop: "15vh"}}>
              <Container style={{display: "flex", flexDirection: "row", justifyContent: "space-around"}}>
                <Loader size="xl" variant="bars" />
              </Container>
              <Container style={{display: "flex", flexDirection: "row", justifyContent: "space-around", marginTop: '2vh'}}>
                <Title order={2} color={PRIMARY_COLOR}>Resetting Stream...</Title>
              </Container>
              <Container style={{display: "flex", flexDirection: "row", justifyContent: "space-around", marginTop: '15vh'}}>
                <Text size="md" color={PRIMARY_COLOR}>Seem like it's taking a while? Twitter's streaming API is very slow to allow new streaming connections.</Text>
              </Container>
            </Container>
          </Container>
        </div> :
        <></>
      }
      {props.tweets.map((t) => (
        <Tweet username={t.Username} geo={t.Geo} tweet={t.Tweet} />
      ))}
    </MapContainer>
  );
};
