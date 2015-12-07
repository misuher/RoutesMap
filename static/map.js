
		var map;
		var GRUmarker;
		var GRPmarker;
		var GQFmarker;
		var LZRmarker;
		var GRUpos;
		var GRPpos;
		var LZRpos;
		var GQFpos;

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
				center: {lat: 27.6034, lng: -15.6982},
				//center: {lat: 26.6475, lng: -16.1693},
				scrollwheel: false,
				zoom: 8,
				styles: styleArray
			});

			var LPA = {lat: 27.926075, lng: -15.390818};
			var TFN = {lat: 28.040288, lng: -16.572979};
			var ACE = {lat: 28.946344, lng: -13.607218};  //ACE-lanzarote
			var EUN = {lat: 27.142294, lng: -13.225541};  //EUN
			var SPC = {lat: 28.622109, lng: -17.755491};  //SPC - La palma
			var DAC = {lat: 23.717010, lng: -15.933434}; //Dakhla

			var LPATFN = [LPA , TFN];
			var LPAACE = [LPA, ACE];
			var LPAEUN = [LPA, EUN];
			var TFNSPC = [TFN, SPC];
			var LPADAC = [LPA, DAC];

			var lineLPATFN = new google.maps.Polyline({
				path: LPATFN,
				geodesic: true,
				strokeColor: '#FF0000',
				strokeOpacity: 0.3,
				strokeWeight: 2
			});

			var lineLPAACE = new google.maps.Polyline({
				path: LPAACE,
				geodesic: true,
				strokeColor: '#FF0000',
				strokeOpacity: 0.3,
				strokeWeight: 2
			});

			var lineLPAEUN = new google.maps.Polyline({
				path: LPAEUN,
				geodesic: true,
				strokeColor: '#FF0000',
				strokeOpacity: 0.3,
				strokeWeight: 2
			});

			var lineTFNSPC = new google.maps.Polyline({
				path: TFNSPC,
				geodesic: true,
				strokeColor: '#FF0000',
				strokeOpacity: 0.3,
				strokeWeight: 2
			});

			var lineLPADAC = new google.maps.Polyline({
				path: LPADAC,
				geodesic: true,
				strokeColor: '#FF0000',
				strokeOpacity: 0.3,
				strokeWeight: 2
			});

			lineLPATFN.setMap(map);
			lineLPAACE.setMap(map);
			lineLPAEUN.setMap(map);
			lineTFNSPC.setMap(map);
			lineLPADAC.setMap(map);

			GRUmarker = new google.maps.Marker({
				position: LPA,
				map: map,
				title: 'GRU',
				draggable: false,
				label: 'U',
				animation: google.maps.Animation.DROP
			});

			GRPmarker = new google.maps.Marker({
				position: TFN,
				map: map,
				title: 'GRP',
				draggable: false,
				label: 'P',
				animation: google.maps.Animation.DROP
			});

			GQFmarker = new google.maps.Marker({
				position: SPC,
				map: map,
				title: 'GQF',
				draggable: false,
				label: 'F',
				animation: google.maps.Animation.DROP
			});

			LZRmarker = new google.maps.Marker({
				position: DAC,
				map: map,
				title: 'LZR',
				draggable: false,
				label: 'R',
				animation: google.maps.Animation.DROP
			});

			updateMarker();

		}


	setInterval(updateMarker,10000);

	function updateMarker() {
		$.post("http://localhost:4000/getCoords", {}, function(json){
			var pos = JSON.parse(json);
			var LatLngGRU = new google.maps.LatLng(pos.Pos[0].Lat, pos.Pos[0].Lng);
			var LatLngGRP = new google.maps.LatLng(pos.Pos[1].Lat, pos.Pos[1].Lng);
			var LatLngGQF = new google.maps.LatLng(pos.Pos[2].Lat, pos.Pos[2].Lng);
			var LatLngLZR = new google.maps.LatLng(pos.Pos[3].Lat, pos.Pos[3].Lng);
			GRUmarker.setPosition(LatLngGRU);
			GRPmarker.setPosition(LatLngGRP);
			GQFmarker.setPosition(LatLngGQF);
			LZRmarker.setPosition(LatLngLZR);
		});

	/*$.post('/path/to/server/getPosition',{}, function(json) {
		var LatLng = new google.maps.LatLng(json.latitude, json.longitude);
		marker.setPosition(LatLng);
	});*/
	}
