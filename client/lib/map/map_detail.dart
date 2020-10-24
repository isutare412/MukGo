//import 'package:contra/custom_widgets/button_round_with_shadow.dart';
//import 'package:contra/custom_widgets/button_solid_with_icon.dart';
//import 'package:contra/custom_widgets/custom_app_bar.dart';
//import 'package:contra/login/contra_text.dart';
//import 'package:contra/utils/colors.dart';
import 'dart:async';
import 'dart:ui' as ui;
import 'dart:math';
//import 'dart:convert';
//import 'dart:io';

//import 'package:fixnum/fixnum.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';
//import '../src/locations.dart' as locations;
//import '../api/api.dart';
import 'package:geolocator/geolocator.dart';
import 'package:mukgo/auth/auth_api.dart';
import 'package:fixnum/fixnum.dart';
import 'package:mukgo/api/api.dart';
import 'package:mukgo/proto/model.pbserver.dart';
import 'package:mukgo/proto/request.pbserver.dart';
import 'package:mukgo/restaurant/restaurant_detail_test.dart';
import 'package:provider/provider.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'map_widget.dart';

class MapDetailPage extends StatefulWidget {
  @override
  _MapDetailPageState createState() => _MapDetailPageState();
}

class _MapDetailPageState extends State<MapDetailPage> {
  GoogleMapController mapController;

  final Set<Marker> _markers = Set<Marker>();
  final Set<Circle> _circles = Set<Circle>();

  var _getPositionSubscription;
  // var tok = 'your token';
  var userIcon;
  var userData;
  var initLocation;

  Future<void> _onMapCreated(controller) async {
    var bitmapDescriptorFromSvgAsset = _bitmapDescriptorFromSvgAsset(
        context, 'assets/images/onboarding_image_five.svg');
    userIcon = await Future.microtask(() {
      return bitmapDescriptorFromSvgAsset;
    });
    userData = await Future.microtask(() {
      var auth = readAuth(context);
      var tok = auth.token;
      return fetchUserData(tok);
    });
    initLocation =
        await getCurrentPosition(desiredAccuracy: LocationAccuracy.best);

    setState(() {
      controller.animateCamera(CameraUpdate.newCameraPosition(CameraPosition(
          target: LatLng(initLocation.latitude, initLocation.longitude),
          zoom: 17.0)));
      var radius = userData.sightRadius;
      _getPositionSubscription =
          getPositionStream().listen((Position position) {
        updatePinOnMap(position, radius);
        updateRestaurants(position, radius);
      });
    });
  }

  void updatePinOnMap(Position position, double radius) async {
    setState(() {
      // position of user
      _markers.removeWhere((m) => m.markerId.value == 'currLoc');
      _markers.add(Marker(
        markerId: MarkerId('currLoc'),
        position: LatLng(position.latitude, position.longitude),
        icon: userIcon,
        anchor: Offset(0.5, 0.5),
      ));

      // range of user for review
      _circles.add(Circle(
          circleId: CircleId('currLoc'),
          center: LatLng(position.latitude, position.longitude),
          radius: radius,
          fillColor: Colors.lightBlueAccent.withOpacity(0.5),
          strokeWidth: 3,
          strokeColor: Colors.lightBlueAccent));
    });
  }

  // position of restaurants
  void updateRestaurants(Position position, double radius) async {
    var coord = Coordinate();
    coord.latitude = position.latitude;
    coord.longitude = position.longitude;

    var restaurantData = await Future.microtask(() {
      var auth = readAuth(context);
      var tok = auth.token;
      return fetchRestaurantsData(tok, coord: coord);
    });

    setState(() {
      _markers.removeWhere((m) => m.markerId.value != 'currLoc');
      restaurantData.restaurants.forEach((r) {
        if (isInMyCircle(position.latitude, position.longitude,
            r.coord.latitude, r.coord.longitude, radius)) {
          _markers.add(Marker(
            markerId: MarkerId(r.id),
            position: LatLng(r.coord.latitude, r.coord.longitude),
            zIndex: 1.0,
            infoWindow: InfoWindow(
                title: r.name,
                snippet: 'This is ' + r.name,
                onTap: () {
                  //Navigator.pushNamed(context, '/project_restaurant',
                  //    arguments: r.id);
                  Navigator.push(
                      context,
                      MaterialPageRoute(
                          builder: (context) =>
                              RestaurantDetailTestPage(restaurant_id: r.id)));
                }),
          ));
        }
      });
    });
  }

  bool isInMyCircle(x1, y1, x2, y2, r) {
    return distanceBetween(x1, y1, x2, y2) < r;
  }

  @override
  void dispose() {
    _getPositionSubscription?.cancel();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      home: Scaffold(
        body: GoogleMap(
          myLocationEnabled: false,
          compassEnabled: true,
          tiltGesturesEnabled: false,
          onMapCreated: _onMapCreated,
          initialCameraPosition:
              CameraPosition(target: const LatLng(37, 126), zoom: 11.0),
          markers: _markers,
          circles: _circles,
        ),
      ),
    );
  }
}

Future<Restaurants> getDummyRestaurants() async {
  await Future.delayed(Duration(microseconds: 100));
  var dummyRestaurants = Restaurants();
  for (var i = 0; i < 40; i++) {
    var dummyRestaurant = Restaurant();
    dummyRestaurant.id = "5f8e9eafcc0ad2855f7c158" + i.toString();
    dummyRestaurant.name = "restaurant" + i.toString();
    dummyRestaurant.coord = Coordinate();
    dummyRestaurant.coord.latitude = 37.4654628 + (i - 20) / 5000;
    dummyRestaurant.coord.longitude = 126.9572302 + (i - 20) / 5000;
    dummyRestaurants.restaurants.add(dummyRestaurant);
  }
  return dummyRestaurants;
}

Future<BitmapDescriptor> _bitmapDescriptorFromSvgAsset(
    BuildContext context, String assetName) async {
  var svgString = await DefaultAssetBundle.of(context).loadString(assetName);
  var svgDrawableRoot = await svg.fromSvgString(svgString, null);

  var queryData = MediaQuery.of(context);
  var devicePixelRatio = queryData.devicePixelRatio;
  var width = 32 * devicePixelRatio;
  var height = 32 * devicePixelRatio;

  var picture = svgDrawableRoot.toPicture(size: Size(width, height));

  var image = await picture.toImage(width.toInt(), height.toInt());
  var bytes = await image.toByteData(format: ui.ImageByteFormat.png);
  return BitmapDescriptor.fromBytes(bytes.buffer.asUint8List());
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
