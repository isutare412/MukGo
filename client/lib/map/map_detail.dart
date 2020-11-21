import 'dart:async';
import 'dart:ui' as ui;

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:google_maps_flutter/google_maps_flutter.dart';

import 'package:geolocator/geolocator.dart';
import 'package:mukgo/auth/auth_api.dart';
import 'package:mukgo/api/api.dart';
import 'package:mukgo/proto/model.pbserver.dart';
import 'package:mukgo/restaurant/restaurant_detail.dart';
import 'package:provider/provider.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:mukgo/user/user_model.dart';
import 'package:mukgo/proto/model.pb.dart';
import 'package:mukgo/restaurant/restaurant_badge.dart';

class MapDetailPage extends StatefulWidget {
  @override
  _MapDetailPageState createState() => _MapDetailPageState();
}

class _MapDetailPageState extends State<MapDetailPage> {
  GoogleMapController _controller;

  final Set<Marker> _markers = Set<Marker>();
  final Set<Circle> _circles = Set<Circle>();
  final settedLocation = LatLng(37.478206, 126.956936);

  var isRestaurantLoading = true;
  var isUserLoading = true;
  var isLoading = true;
  var _getPositionSubscription;
  var userIcon;
  var floating = false;
  var restaurantIcons;

  UserModel userData;

  Future<void> _onMapCreated(controller) async {
    isRestaurantLoading = isUserLoading = isLoading = true;
    _controller = controller;

    _getPositionSubscription =
        getPositionStream().listen((position) => positionStream(position));

    restaurantIcons = <RestaurantType, BitmapDescriptor>{};
    badgePathMap.forEach((key, svgDir) async {
      var bitmapDescriptorFromSvgAsset =
          _bitmapDescriptorFromSvgAsset(context, svgDir);
      restaurantIcons[key] = await bitmapDescriptorFromSvgAsset;
    });
  }

  void positionStream(position) async {
    var markerShown = false;
    for (var m in _markers) {
      if (markerShown) break;
      var value = await _controller.isMarkerInfoWindowShown(m.markerId);
      if (value) {
        markerShown = true;
      }
    }
    if (!markerShown) {
      locationChanged(position);
    }
  }

  void locationChanged(position) {
    context.read<UserModel>().fetch().then((value) {
      userData = context.read<UserModel>();
      var radius = 100.0;
      if (userData != null) {
        radius = userData.sightRadius;
        var zoom = 19 - ((radius + radius) / 100) / 2;

        _controller.animateCamera(CameraUpdate.newCameraPosition(CameraPosition(
            target: LatLng(position.latitude, position.longitude),
            zoom: zoom)));

        updatePinOnMap(position, radius);
        updateRestaurants(position, radius);
      }
    });
  }

  void updatePinOnMap(position, double radius) async {
    var svgDir = userData.profileAsset();
    var bitmapDescriptorFromSvgAsset =
        _bitmapDescriptorFromSvgAsset(context, svgDir);
    userIcon = await Future.microtask(() {
      return bitmapDescriptorFromSvgAsset;
    });
    setState(() {
      // position of user
      _markers.removeWhere((m) => m.markerId.value == 'currLoc');
      _markers.add(Marker(
        markerId: MarkerId('currLoc'),
        position: LatLng(position.latitude, position.longitude),
        draggable: floating,
        onDragEnd: locationChanged,
        icon: userIcon,
        anchor: Offset(0.5, 0.5),
      ));

      // range of user for review
      _circles.removeWhere((m) => m.circleId.value == 'currLoc');
      _circles.add(Circle(
          circleId: CircleId('currLoc'),
          center: LatLng(position.latitude, position.longitude),
          radius: radius,
          fillColor: Colors.lightBlueAccent.withOpacity(0.5),
          strokeWidth: 3,
          strokeColor: Colors.lightBlueAccent));

      isUserLoading = false;
      isLoading = isUserLoading || isRestaurantLoading;
    });
  }

  // position of restaurants
  void updateRestaurants(position, double radius) async {
    var coord = Coordinate();
    coord.latitude = position.latitude;
    coord.longitude = position.longitude;

    await Future.microtask(() async {
      var auth = readAuth(context);
      var tok = auth.token;
      var restaurantData = await fetchRestaurantsData(tok, coord: coord);
      setState(() {
        _markers.removeWhere((m) => m.markerId.value != 'currLoc');

        restaurantData.restaurants.forEach((r) {
          if (isInMyCircle(position.latitude, position.longitude,
              r.coord.latitude, r.coord.longitude, radius)) {
            var restaurantIcon = restaurantIcons[r.type];
            _markers.add(Marker(
              markerId: MarkerId(r.id),
              position: LatLng(r.coord.latitude, r.coord.longitude),
              icon: restaurantIcon,
              anchor: Offset(0.5, 0.5),
              zIndex: 1.0,
              infoWindow: InfoWindow(
                  title: r.name,
                  onTap: () {
                    Navigator.push(
                        context,
                        MaterialPageRoute(
                            builder: (context) =>
                                RestaurantDetailTestPage(restaurant_id: r.id)));
                  }),
            ));
          }
        });
        isRestaurantLoading = false;
        isLoading = isUserLoading || isRestaurantLoading;
      });
    });

    setState(() {});
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
          mapToolbarEnabled: false,
          tiltGesturesEnabled: floating,
          rotateGesturesEnabled: floating,
          scrollGesturesEnabled: floating,
          zoomControlsEnabled: false,
          //zoomGesturesEnabled: false,
          onMapCreated: _onMapCreated,
          initialCameraPosition:
              CameraPosition(target: settedLocation, zoom: 11.0),
          markers: _markers,
          circles: _circles,
        ),
        floatingActionButton: FloatingActionButton(
          onPressed: () {
            floating = !floating;
            isRestaurantLoading = isUserLoading = isLoading = true;
            if (floating) {
              _getPositionSubscription?.cancel();
              locationChanged(settedLocation);
            } else {
              _getPositionSubscription = getPositionStream()
                  .listen((position) => positionStream(position));
            }
          },
          child: Stack(children: [
            Align(alignment: Alignment.center, child: Icon(Icons.place)),
            Align(
                alignment: Alignment.center,
                child: Visibility(
                    visible: isLoading,
                    child: CircularProgressIndicator(
                        valueColor:
                            AlwaysStoppedAnimation<Color>(Colors.white)))),
          ]),
          backgroundColor: floating ? Colors.red : Colors.blue,
        ),
      ),
    );
  }
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
