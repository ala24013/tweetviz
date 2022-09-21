import React from 'react'
import { MapContainer, TileLayer, useMap, Popup, Marker } from 'react-leaflet'

export default Map = (props) => {
    return (
        <MapContainer center={[45.4, -75.7]} zoom={10} scrollWheelZoom={true}>
          <TileLayer
            url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
          />
        </MapContainer>
      );
}