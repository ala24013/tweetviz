import React from "react";
import { Map, Marker } from "pigeon-maps";

function CustomMarker(props) {
  return <Marker props />;
}

export default function World() {
  return (
    <Map
      height={window.innerHeight - 4}
      defaultCenter={[47.879, 8.6997]}
      defaultZoom={11}
      onBoundsChanged={({ center, zoom, bounds, initial }) => {
        console.log(bounds);
      }}
    >
      <CustomMarker
        width={20}
        color={"red"}
        anchor={[47.879, 8.6997]}
        onMouseOver={(e) => console.log(e)}
        payload={1}
      />
    </Map>
  );
}
