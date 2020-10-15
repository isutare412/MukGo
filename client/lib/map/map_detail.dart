import 'package:contra/custom_widgets/button_round_with_shadow.dart';
import 'package:contra/custom_widgets/button_solid_with_icon.dart';
import 'package:contra/custom_widgets/custom_app_bar.dart';
import 'package:contra/login/contra_text.dart';
import 'package:contra/utils/colors.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
import '../src/locations.dart' as locations;
import 'package:geolocator/geolocator.dart';

class MapDetailPage extends StatefulWidget {
  @override
  _MapDetailPageState createState() => _MapDetailPageState();
}

class _MapDetailPageState extends State<MapDetailPage> {
  GoogleMapController mapController;

  final Set<Marker> _markers = Set<Marker>();
  Future<void> _onMapCreated(GoogleMapController controller) async {
    // Get GPS Location
    var currentLocation =
        await getCurrentPosition(desiredAccuracy: LocationAccuracy.best);

    //final googleOffices = await locations.getGoogleOffices();
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
      /**for (final office in googleOffices.offices) {
        _markers.add(Marker(
          markerId: MarkerId(office.name),
          position: LatLng(office.lat, office.lng),
          infoWindow: InfoWindow(
            title: office.name,
            snippet: office.address,
          ),
        ));
      }
      */
      getPositionStream().listen((Position position) async {
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
          initialCameraPosition:
              CameraPosition(target: const LatLng(37, 126), zoom: 11.0),
          markers: _markers,
        ),
      ),
    );
  }
}

/**
class _MapDetailPageState extends State<MapDetailPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: persian_blue,
      appBar: CustomAppBar(
        height: 120,
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceEvenly,
          crossAxisAlignment: CrossAxisAlignment.end,
          children: <Widget>[
            Expanded(
              flex: 1,
              child: Padding(
                padding: const EdgeInsets.only(left: 24.0),
                child: Align(
                  alignment: Alignment.bottomLeft,
                  child: ButtonRoundWithShadow(
                      size: 48,
                      borderColor: lightening_yellow,
                      color: lightening_yellow,
                      callback: () {
                        Navigator.pop(context);
                      },
                      shadowColor: wood_smoke,
                      iconPath: "assets/icons/arrow_back_white.svg"),
                ),
              ),
            ),
            Expanded(
              flex: 2,
              child: ContraText(
                size: 27,
                color: white,
                alignment: Alignment.bottomCenter,
                text: "Directions",
              ),
            ),
            Expanded(
              flex: 1,
              child: SizedBox(
                width: 20,
              ),
            )
          ],
        ),
      ),
      body: Column(
        children: <Widget>[
          SizedBox(
            height: 24,
          ),
          Expanded(
            flex: 4,
            child: Container(
              padding: EdgeInsets.all(24),
              child: MapWidget(
                distance: "2.5",
              ),
            ),
          ),
          Expanded(
            flex: 3,
            child: Container(
              padding: EdgeInsets.symmetric(horizontal: 24),
              child: Column(
                children: <Widget>[
                  Padding(
                    padding: const EdgeInsets.symmetric(vertical: 4.0),
                    child: ContraText(
                      color: white,
                      size: 44,
                      weight: FontWeight.w800,
                      alignment: Alignment.centerLeft,
                      text: "Space 8",
                    ),
                  ),
                  Padding(
                    padding: const EdgeInsets.symmetric(vertical: 4.0),
                    child: ContraText(
                      color: white,
                      size: 21,
                      weight: FontWeight.w500,
                      alignment: Alignment.centerLeft,
                      text: "Wolf Crater, 897, \n New Milkyway Mars",
                    ),
                  ),
                  Padding(
                    padding: const EdgeInsets.symmetric(vertical: 12),
                    child: ButtonPlainWithIcon(
                      text: "Get Direction",
                      color: wood_smoke,
                      callback: () {},
                      size: 48,
                      isPrefix: false,
                      isSuffix: true,
                      textColor: white,
                      iconPath: "assets/icons/ic_navigation_white.svg",
                    ),
                  )
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}

 */
