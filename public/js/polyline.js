function initMap() {
    const map = new google.maps.Map(document.getElementById("map"), {
        zoom: 3,
        center: { lat: 0, lng: -180 },
        mapTypeId: "terrain"
    });
    const flightPlanCoordinates = [];
    const flightPath = new google.maps.Polyline({
        path: flightPlanCoordinates,
        geodesic: true,
        strokeColor: "#FF0000",
        strokeOpacity: 1.0,
        strokeWeight: 2
    });
    flightPath.setMap(map);
}

initMap()