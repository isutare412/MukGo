import 'dart:async';

import 'package:contra/login/contra_text.dart';
import 'package:contra/utils/colors.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import '../src/locations.dart' as locations;
import 'package:geolocator/geolocator.dart';

class MapWidget extends StatefulWidget {
  String distance;
  bool isDetail;

  MapWidget({this.distance, this.isDetail});

  @override
  _MapWidgetState createState() => _MapWidgetState();
}

class _MapWidgetState extends State<MapWidget> {
  GoogleMapController mapController;

  final Set<Marker> _markers = Set<Marker>();
  Future<void> _onMapCreated(GoogleMapController controller) async {
    // Get GPS Location
    var currentLocation =
        await getCurrentPosition(desiredAccuracy: LocationAccuracy.best);

    final googleOffices = await locations.getGoogleOffices();
    setState(() {
      _markers.clear();
      _markers.add(Marker(
        markerId: MarkerId('currLoc'),
        position: LatLng(currentLocation.latitude, currentLocation.longitude),
        infoWindow: InfoWindow(title: 'Your Location'),
      ));
      controller.animateCamera(CameraUpdate.newCameraPosition(CameraPosition(
          target: LatLng(currentLocation.latitude, currentLocation.longitude),
          zoom: 17.0)));
      for (final office in googleOffices.offices) {
        _markers.add(Marker(
          markerId: MarkerId(office.name),
          position: LatLng(office.lat, office.lng),
          infoWindow: InfoWindow(
            title: office.name,
            snippet: office.address,
          ),
        ));
        getPositionStream().listen((Position position) {
          print(position == null
              ? 'Unknown'
              : position.latitude.toString() +
                  ', ' +
                  position.longitude.toString());
          updatePinOnMap(position);
          /*
          controller.animateCamera(CameraUpdate.newCameraPosition(
              CameraPosition(
                  target: LatLng(position.latitude, position.longitude),
                  zoom: 17.0)));
          **/
          print(_markers.where((m) => m.markerId.value == 'currLoc'));
        });
      }
    });
  }

  void updatePinOnMap(Position position) async {
    setState(() {
      _markers.removeWhere((m) => m.markerId.value == 'currLoc');
      _markers.add(Marker(
        markerId: MarkerId('currLoc'),
        position: LatLng(position.latitude, position.longitude),
        infoWindow: InfoWindow(title: 'Your Location'),
      ));
    });
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: Scaffold(
        body: GoogleMap(
          myLocationEnabled: true,
          compassEnabled: true,
          tiltGesturesEnabled: false,
          onMapCreated: _onMapCreated,
          initialCameraPosition: CameraPosition(
              target: const LatLng(37.4654628, 126.9572302), zoom: 11.0),
          markers: _markers,
        ),
      ),
    );
  }
}

/**
class MapWidget extends StatefulWidget {
  String distance;
  bool isDetail;

  MapWidget({this.distance, this.isDetail});

  @override
  _MapWidgetState createState() => _MapWidgetState();
}

class _MapWidgetState extends State<MapWidget> {
  Completer<GoogleMapController> _controller = Completer();
  static const LatLng _center = const LatLng(45.521563, -122.677433);

  void _onMapCreated(GoogleMapController controller) {
    _controller.complete(controller);
  }

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: ShapeDecoration(
          shadows: [
            BoxShadow(
              color: wood_smoke,
              offset: Offset(0, 2),
            )
          ],
          color: white,
          shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.all(Radius.circular(16)),
              side: BorderSide(
                color: wood_smoke,
                width: 2,
              ))),
      child: Stack(
        children: <Widget>[
          Container(
            padding: EdgeInsets.all(4),
            child: GoogleMap(
              onMapCreated: _onMapCreated,
              initialCameraPosition: CameraPosition(
                target: _center,
                zoom: 11.0,
              ),
            ),
          ),
          Positioned(
            right: 24,
            bottom: 24,
            child: Container(
              width: 64,
              padding: EdgeInsets.all(4),
              height: 64,
              decoration: BoxDecoration(
                  color: wood_smoke,
                  borderRadius: BorderRadius.all(Radius.circular(16))),
              child: Column(
                children: <Widget>[
                  ContraText(
                    alignment: Alignment.center,
                    text: "1.5",
                    size: 22,
                    color: white,
                    weight: FontWeight.bold,
                  ),
                  ContraText(
                    alignment: Alignment.center,
                    text: "Kms",
                    size: 15,
                    color: white,
                    weight: FontWeight.bold,
                  )
                ],
              ),
            ),
          )
        ],
      ),
    );
  }
}

*/
