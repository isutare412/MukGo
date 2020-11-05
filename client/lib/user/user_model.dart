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
  int reviewCount;
  int likeCount;

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
    if (auth.token == null) return;
    var userData = await fetchUserData(auth.token);
    if (userData == null) return;
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
      double sightRadius,
      int reviewCount,
      int likeCount}) {
    this.id = id;
    this.name = name;
    this.level = level;
    this.totalExp = totalExp;
    this.levelExp = levelExp;
    this.curExp = curExp;
    this.expRatio = expRatio;
    this.sightRadius = sightRadius;
    this.reviewCount = reviewCount;
    this.likeCount = likeCount;

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
        sightRadius: userData.sightRadius,
        reviewCount: userData.reviewCount,
        likeCount: userData.likeCount);
  }

  String profileAsset() {
    switch (level) {
      case 1:
        {
          return 'assets/images/onboarding_image_one.svg';
        }
      case 2:
        {
          return 'assets/images/onboarding_image_two.svg';
        }
      case 3:
        {
          return 'assets/images/onboarding_image_three.svg';
        }
      case 4:
        {
          return 'assets/images/onboarding_image_four.svg';
        }
      default:
        {
          return 'assets/images/onboarding_image_five.svg';
        }
    }
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
        sightRadius: null,
        reviewCount: null,
        likeCount: null);
  }
}
