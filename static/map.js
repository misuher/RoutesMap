
		var map;
		var marker;
		
		var styleArray = [
			{
			featureType: "all",
			stylers: [
			{ saturation: -80 }
			]
			},{
			featureType: "road.arterial",
			elementType: "geometry",
			stylers: [
				{ hue: "#00ffee" },
				{ saturation: 50 }
			]
			},{
			featureType: "poi.business",
			elementType: "labels",
			stylers: [
				{ visibility: "off" }
			]
			}
		];

		function initMap() {
			$('#map-canvas').height($(window).height());
		// Create a map object and specify the DOM element for display.
		map = new google.maps.Map(document.getElementById('map-canvas'), {
			center: {lat: 28.6034, lng: -15.6982},
			scrollwheel: false,
			zoom: 8,
			styles: styleArray
		});

			var path1Coor = [
				{lat: 27.926075, lng: -15.390818},
				{lat: 28.040288, lng: -16.572979}
			];

			var path1 = new google.maps.Polyline({
				path: path1Coor,
				geodesic: true,
				strokeColor: '#FF0000',
				strokeOpacity: 1.0,
				strokeWeight: 2
			});

			var GRUpos = {lat: 27.926075, lng: -15.390818};
			marker = new google.maps.Marker({
				position: GRUpos,
				map: map,
				title: 'GRU',
				draggable: false,
				label: 'GRU',
				animation: google.maps.Animation.DROP
			});

			path1.setMap(map);
		}



		// every 10 seconds
		setInterval(updateMarker,10000);

		function updateMarker() {
			marker.setPosition({lat: 28.040288, lng: -16.572979});
		/*$.post('/path/to/server/getPosition',{}, function(json) {
			var LatLng = new google.maps.LatLng(json.latitude, json.longitude);
			marker.setPosition(LatLng);
		});*/
		}
