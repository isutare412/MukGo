import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:http/http.dart' as http;
import 'package:mukgo/protocol/code.pbenum.dart';

String apiUrl = 'http://10.0.2.2:7777';

// user data response
class UserData {
  String name;
  int level;
  int totalExp;
  int levelExp;
  int curExp;
  double expRatio;
  double sightRadius;

  UserData(
      {this.name,
      this.level,
      this.totalExp,
      this.levelExp,
      this.curExp,
      this.expRatio,
      this.sightRadius});

  // set user info data from json
  factory UserData.fromJSON(Map<String, dynamic> json) {
    return UserData(
        name: json['name'],
        level: json['level'],
        totalExp: json['total_exp'],
        levelExp: json['level_exp'],
        curExp: json['cur_exp'],
        expRatio: json['exp_ratio'],
        sightRadius: json['sight_radius']);
  }
}

// fetch user data from api
Future<UserData> fetchUserData(String token) async {
  try {
    var headers = <String, String>{'Authorization': 'Bearer $token'};
    var res = await http.get('$apiUrl/user', headers: headers);
    if (res.statusCode != 200) {
      return null;
    }
    return UserData.fromJSON(json.decode(res.body));
  } catch (e) {
    return null;
  }

  // serve test data instead
  // await Future.delayed(Duration(seconds: 3));
  // var dummyUser = UserData(
  //     name: '홍길동',
  //     level: 7,
  //     totalExp: 1000,
  //     levelExp: 500,
  //     curExp: 300,
  //     expRatio: 0.6,
  //     sightRadius: 1.0);
  // return dummyUser;
}

// post request to sign up/in
Future<int> trySignUp(String token) async {
  try {
    var headers = getAuthHeader(token);
    var res = await http.post('$apiUrl/user', headers: headers);
    if (res.statusCode != HttpStatus.ok) {
      var data = json.decode(res.body);
      if (data['code'] == Code.USER_EXISTS.value) {
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
