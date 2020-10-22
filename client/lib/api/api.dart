import 'dart:async';
import 'dart:convert';
import 'dart:io';

import 'package:fixnum/fixnum.dart';
import 'package:http/http.dart' as http;

import 'package:mukgo/proto/model.pb.dart';

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
    if (res.statusCode != HttpStatus.ok) {
      printResponseError('fetchUserData', res.body);
      return null;
    }
    print('body: ${res.body}');
    return User()..mergeFromProto3Json(jsonDecode(res.body));
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
    if (res.statusCode != HttpStatus.ok) {
      var reason = ErrorReason()..mergeFromProto3Json(jsonDecode(res.body));
      if (reason.code == Code.USER_EXISTS) {
        // already has an account
        return 1;
      }
      printResponseError('trySignUp', res.body);
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
Future<Restaurants> fetchRestaurantsData(String token) async {
  try {
    var headers = getAuthHeader(token);
    var res = await http.get('$apiUrl/restaurants', headers: headers);
    if (res.statusCode != HttpStatus.ok) {
      printResponseError('fetchRestaurantsData', res.body);
      return null;
    }
    print('body: ${res.body}');
    return Restaurants()..mergeFromProto3Json(jsonDecode(res.body));
  } catch (e) {
    printAPIError('fetchRestaurantsData', e);
    return null;
  }
}

// post restaurants data to api
Future<bool> postRestaurantsData(String token, {Restaurants data}) async {
  try {
    var headers = getAuthHeader(token);
    var body = data.writeToJsonMap();
    var res =
        await http.post('$apiUrl/restaurants', body: body, headers: headers);
    if (res.statusCode != HttpStatus.ok) {
      printResponseError('postRestaurantsData', res.body);
      return false;
    }
    return true;
  } catch (e) {
    printAPIError('postRestaurantsData', e);
    return false;
  }
}

// post restaurant data to api
Future<bool> postRestaurantData(String token, {Restaurants data}) async {
  try {
    var headers = getAuthHeader(token);
    var body = data.writeToJsonMap();
    var res =
        await http.post('$apiUrl/restaurants', body: body, headers: headers);
    if (res.statusCode != HttpStatus.ok) {
      printResponseError('postRestaurantData', res.body);
      return false;
    }
    return true;
  } catch (e) {
    printAPIError('postRestaurantData', e);
    return false;
  }
}

// fetch reviews data from api
Future<Reviews> fetchReviewsData(String token) async {
  try {
    var headers = getAuthHeader(token);
    var res = await http.get('$apiUrl/review', headers: headers);
    if (res.statusCode != HttpStatus.ok) {
      printResponseError('fetchReviewsData', res.body);
      return null;
    }
    print('body: ${res.body}');
    return Reviews()..mergeFromProto3Json(jsonDecode(res.body));
  } catch (e) {
    printAPIError('fetchReviewsData', e);
    return null;
  }
}

// post review data to api
Future<bool> postReviewData(String token, {Review data}) async {
  try {
    var headers = getAuthHeader(token);
    var body = data.writeToJsonMap();
    var res = await http.post('$apiUrl/review', body: body, headers: headers);
    if (res.statusCode != HttpStatus.ok) {
      printResponseError('postReviewData', res.body);
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
