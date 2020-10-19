import 'dart:async';
import 'dart:io';

import 'package:fixnum/fixnum.dart';
import 'package:http/http.dart' as http;

import 'package:mukgo/proto/model.pb.dart';

String apiUrl = 'http://localhost:7777';

// fetch user data from api
Future<User> fetchUserData(String token) async {
  // try {
  //   var headers = <String, String>{'Authorization': 'Bearer $token'};
  //   var res = await http.get('$apiUrl/user', headers: headers);
  //   if (res.statusCode != HttpStatus.ok) {
  //     return null;
  //   }
  //   return User.fromJson(res.body);
  // } catch (e) {
  //   return null;
  // }

  // serve test data instead
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

// post request to sign up/in
Future<int> trySignUp(String token) async {
  try {
    var headers = getAuthHeader(token);
    var res = await http.post('$apiUrl/user', headers: headers);
    if (res.statusCode != HttpStatus.ok) {
      var reason = ErrorReason.fromJson(res.body);
      if (reason.code == Code.USER_EXISTS) {
        // already has an account
        return 1;
      }
    }

    // HttpStatus.ok
    return 0;
  } catch (e) {
    return null;
  }
}

Map<String, String> getAuthHeader(String token) =>
    <String, String>{'Authorization': 'Bearer $token'};
