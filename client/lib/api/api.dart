import 'dart:async';
import 'dart:io';
import 'dart:typed_data';

import 'package:fixnum/fixnum.dart';
import 'package:http/http.dart' as http;

import 'package:mukgo/proto/model.pb.dart';
import 'package:mukgo/proto/request.pb.dart';

//String apiUrl = '10.0.2.2:7777';
String apiUrl = 'redshore.asuscomm.com:7777';

// log api error reason code
void printResponseError(String api, Uint8List bytes) {
  var reason = ErrorReason.fromBuffer(bytes);
  var c = reason.code;
  print('Response ERROR: $api: ${c.name} (${c.value})');
}

void printAPIError(String api, dynamic msg) {
  print('API Method ERROR: $api: $msg');
}

// fetch user data from api
Future<User> fetchUserData(String token) async {
  try {
    var headers = getAuthHeader(token);
    var res = await http.get('http://$apiUrl/user', headers: headers);
    if (res.statusCode != HttpStatus.ok) {
      printResponseError('fetchUserData', res.bodyBytes);
      return null;
    }

    return User.fromBuffer(res.bodyBytes);
  } catch (e) {
    printAPIError('fetchUserData', e);
    return null;
  }

  // serve test data instead
  // return getDummyUser();
}

// post request to sign up/in
Future<int> trySignUp(String token) async {
  try {
    var headers = getAuthHeader(token);
    var res = await http.post('http://$apiUrl/user', headers: headers);
    if (res.statusCode != HttpStatus.ok) {
      var reason = ErrorReason.fromBuffer(res.bodyBytes);
      if (reason.code == Code.USER_EXISTS) {
        // already has an account
        return 1;
      }
      printResponseError('trySignUp', res.bodyBytes);
      return null;
    }

    // HttpStatus.ok
    return 0;
  } catch (e) {
    printAPIError('trySignUp', e);
    return null;
  }
}

Future<Restaurant> fetchRestaurantData(String token,
    {String restaurantId}) async {
  try {
    var headers = getAuthHeader(token);
    var uri = Uri.http(apiUrl, '/restaurant', {'restaurant_id': restaurantId});
    var res = await http.get(uri, headers: headers);
    if (res.statusCode != HttpStatus.ok) {
      printResponseError('fetchRestaurantData', res.bodyBytes);
      return null;
    }

    return Restaurant.fromBuffer(res.bodyBytes);
  } catch (e) {
    printAPIError('fetchRestaurantData', e);
    return null;
  }
}

// get restaurants data from api
Future<Restaurants> fetchRestaurantsData(String token,
    {Coordinate coord}) async {
  try {
    var headers = getAuthHeader(token);
    var uri = Uri.http(apiUrl, '/restaurants', {
      'latitude': coord.latitude.toString(),
      'longitude': coord.longitude.toString(),
    });

    var res = await http.get(uri, headers: headers);
    if (res.statusCode != HttpStatus.ok) {
      printResponseError('fetchRestaurantsData', res.bodyBytes);
      return null;
    }

    return Restaurants.fromBuffer(res.bodyBytes);
  } catch (e) {
    printAPIError('fetchRestaurantsData', e);
    return null;
  }
}

// post restaurants data to api
Future<bool> postRestaurantsData(String token, {Restaurants data}) async {
  try {
    var headers = getAuthHeader(token);
    var query = RestaurantsPost()..restaurants.addAll(data.restaurants);
    var res = await http.post('http://$apiUrl/restaurants',
        body: query.writeToBuffer(), headers: headers);

    if (res.statusCode != HttpStatus.ok) {
      printResponseError('postRestaurantsData', res.bodyBytes);
      return false;
    }

    return true;
  } catch (e) {
    printAPIError('postRestaurantsData', e);
    return false;
  }
}

// post restaurant data to api
Future<bool> postRestaurantData(String token, {Restaurant data}) async {
  try {
    var headers = getAuthHeader(token);
    var query = RestaurantPost()..restaurant = data;
    var res = await http.post('http://$apiUrl/restaurants',
        body: query.writeToBuffer(), headers: headers);

    if (res.statusCode != HttpStatus.ok) {
      printResponseError('postRestaurantData', res.bodyBytes);
      return false;
    }

    return true;
  } catch (e) {
    printAPIError('postRestaurantData', e);
    return false;
  }
}

// fetch reviews data from api
Future<Reviews> fetchReviewsData(String token, {String restaurantId}) async {
  try {
    var headers = getAuthHeader(token);
    var uri = Uri.http(apiUrl, '/reviews', {'restaurant_id': restaurantId});
    var res = await http.get(uri, headers: headers);
    if (res.statusCode != HttpStatus.ok) {
      printResponseError('fetchReviewsData', res.bodyBytes);
      return null;
    }

    return Reviews.fromBuffer(res.bodyBytes);
  } catch (e) {
    printAPIError('fetchReviewsData', e);
    return null;
  }
}

// post review data to api
Future<bool> postReviewData(String token, {Review data, String id}) async {
  try {
    var headers = getAuthHeader(token);
    var query = ReviewPost()
      ..restaurantId = id
      ..review = data;
    var res = await http.post('http://$apiUrl/review',
        body: query.writeToBuffer(), headers: headers);

    if (res.statusCode != HttpStatus.ok) {
      printResponseError('postReviewData', res.bodyBytes);
      return false;
    }

    return true;
  } catch (e) {
    printAPIError('postReviewData', e);
    return false;
  }
}

Map<String, String> getAuthHeader(String token) => <String, String>{
      HttpHeaders.authorizationHeader: 'Bearer $token',
      HttpHeaders.contentTypeHeader: 'application/protobuf'
    };

Future<User> getDummyUser() async {
  await Future.delayed(Duration(seconds: 3));
  var dummyUser = User();
  dummyUser.name = '홍길동';
  dummyUser.level = 7;
  dummyUser.totalExp = Int64(1000);
  dummyUser.levelExp = Int64(500);
  dummyUser.curExp = Int64(300);
  dummyUser.expRatio = 0.6;
  dummyUser.sightRadius = 30.0;
  return dummyUser;
}
