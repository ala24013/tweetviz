import React from "react"
import { Map, Marker } from "pigeon-maps"

export default function World() {
  return (
    <Map height={window.innerHeight-4} defaultCenter={[47.879, 8.6997]} defaultZoom={11} onBoundsChanged={({center, zoom, bounds, initial}) => {console.log(bounds)}}>
      <Marker width={20} color={"red"} anchor={[47.879, 8.6997]} />
    </Map>
  )
}