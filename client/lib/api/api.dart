import 'dart:async';
import 'dart:convert';
import 'dart:io';

import 'package:fixnum/fixnum.dart';
import 'package:http/http.dart' as http;

import 'package:mukgo/proto/model.pb.dart';
import 'package:mukgo/proto/request.pb.dart';

String apiUrl = 'http://10.0.2.2:7777';

// log api error reason code
void printResponseError(String api, String body) {
  var reason = ErrorReason()..mergeFromProto3Json(jsonDecode(body));
  var c = reason.code;
  print('Response ERROR: $api: ${c.name} (${c.value})');
}

void printAPIError(String api, String msg) {
  print('API Method ERROR: $api: $msg');
}

// fetch user data from api
Future<User> fetchUserData(String token) async {
  try {
    var headers = getAuthHeader(token);
    var res = await http.get('$apiUrl/user', headers: headers);
    var body = utf8.decode(res.bodyBytes);

    if (res.statusCode != HttpStatus.ok) {
      printResponseError('fetchUserData', body);
      return null;
    }

    print('body: $body');
    return User()..mergeFromProto3Json(jsonDecode(body));
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
    var res = await http.post('$apiUrl/user', headers: headers);
    var body = utf8.decode(res.bodyBytes);

    if (res.statusCode != HttpStatus.ok) {
      var reason = ErrorReason()..mergeFromProto3Json(jsonDecode(body));
      if (reason.code == Code.USER_EXISTS) {
        // already has an account
        return 1;
      }
      printResponseError('trySignUp', body);
      return null;
    }

    // HttpStatus.ok
    return 0;
  } catch (e) {
    printAPIError('trySignUp', e);
    return null;
  }
}

// get restaurants data from api
Future<Restaurants> fetchRestaurantsData(String token,
    {Coordinate coord}) async {
  try {
    var headers = getAuthHeader(token);
    var query = RestaurantsGet()..coord = coord;
    var uri = Uri.http(apiUrl, '/restaurants', query.toProto3Json());
    var res = await http.get(uri, headers: headers);
    var body = utf8.decode(res.bodyBytes);

    if (res.statusCode != HttpStatus.ok) {
      printResponseError('fetchRestaurantsData', body);
      return null;
    }

    print('body: ${utf8.decode(res.bodyBytes)}');
    return Restaurants()..mergeFromProto3Json(jsonDecode(body));
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
    var res = await http.post('$apiUrl/restaurants',
        body: query.toProto3Json(), headers: headers);
    var body = utf8.decode(res.bodyBytes);

    if (res.statusCode != HttpStatus.ok) {
      printResponseError('postRestaurantsData', body);
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
    var res = await http.post('$apiUrl/restaurants',
        body: query.toProto3Json(), headers: headers);
    var body = utf8.decode(res.bodyBytes);

    if (res.statusCode != HttpStatus.ok) {
      printResponseError('postRestaurantData', body);
      return false;
    }

    return true;
  } catch (e) {
    printAPIError('postRestaurantData', e);
    return false;
  }
}

// fetch reviews data from api
Future<Reviews> fetchReviewsData(String token, {String id}) async {
  try {
    var headers = getAuthHeader(token);
    var query = ReviewsGet()..restaurantId = id;
    var uri = Uri.http(apiUrl, '/reviews', query.toProto3Json());
    var res = await http.get(uri, headers: headers);
    var body = utf8.decode(res.bodyBytes);

    if (res.statusCode != HttpStatus.ok) {
      printResponseError('fetchReviewsData', body);
      return null;
    }

    print('body: $body');
    return Reviews()..mergeFromProto3Json(jsonDecode(body));
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
    var res = await http.post('$apiUrl/review',
        body: query.toProto3Json(), headers: headers);
    var body = utf8.decode(res.bodyBytes);

    if (res.statusCode != HttpStatus.ok) {
      printResponseError('postReviewData', body);
      return false;
    }

    return true;
  } catch (e) {
    printAPIError('postReviewData', e);
    return false;
  }
}

Map<String, String> getAuthHeader(String token) =>
    <String, String>{'Authorization': 'Bearer $token'};

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
