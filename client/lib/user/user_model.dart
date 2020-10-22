import 'dart:async';

import 'package:flutter/cupertino.dart';
import 'package:fixnum/fixnum.dart';

import 'package:mukgo/proto/model.pb.dart';
import 'package:mukgo/auth/auth_model.dart';
import 'package:mukgo/api/api.dart';

class UserModel extends ChangeNotifier {
  // copy of auth model
  AuthModel _auth;

  // user info data
  String id;
  String name;
  int level;
  Int64 totalExp;
  Int64 levelExp;
  Int64 curExp;
  double expRatio;
  double sightRadius;

  // simple auth model accesors for debug usage
  AuthModel get auth => _auth;
  set auth(AuthModel auth) {
    _auth = auth;
    if (_auth == null || _auth.user == null) {
      clear();
    } else {
      fetch();
    }
  }

  // fetch fresh user data from server
  Future<void> fetch() async {
    var userData = await fetchUserData(auth.token);
    updateFromUserData(userData);
  }

  void update(
      {String id,
      String name,
      int level,
      Int64 totalExp,
      Int64 levelExp,
      Int64 curExp,
      double expRatio,
      double sightRadius}) {
    this.id = id;
    this.name = name;
    this.level = level;
    this.totalExp = totalExp;
    this.levelExp = levelExp;
    this.curExp = curExp;
    this.expRatio = expRatio;
    this.sightRadius = sightRadius;

    notifyListeners();
  }

  void updateFromUserData(User userData) {
    update(
        id: userData.id,
        name: userData.name,
        level: userData.level,
        totalExp: userData.totalExp,
        levelExp: userData.levelExp,
        curExp: userData.curExp,
        expRatio: userData.expRatio,
        sightRadius: userData.sightRadius);
  }

  void clear() {
    update(
        id: null,
        name: null,
        level: null,
        totalExp: null,
        levelExp: null,
        curExp: null,
        expRatio: null,
        sightRadius: null);
  }
}
