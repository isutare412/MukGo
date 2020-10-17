import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter/foundation.dart';
import 'package:google_sign_in/google_sign_in.dart';

import 'package:mukgo/api/api.dart';

class AuthModel extends ChangeNotifier {
  final GoogleSignIn _googleSignIn = GoogleSignIn(
    scopes: <String>['openid', 'profile', 'email'],
  );
  GoogleSignInAccount _user;
  String _token;
  bool debug;

  // Set up google sign in
  AuthModel({this.debug = false}) : super() {
    _googleSignIn.onCurrentUserChanged.listen(signIn);
    _googleSignIn.signInSilently();
  }

  GoogleSignInAccount get user => _user;
  String get token => _token;

  // Sign in listner
  Future<void> signIn(GoogleSignInAccount account) async {
    var authentication = await account.authentication;
    print(authentication.toTokenString());
    _token = authentication.accessToken;
    _user = account;
    // var code = await trySignUp(token);
    // print('try sign up response code $code');

    notifyListeners();
  }

  // Google sign in handler
  Future<void> googleSignIn() async {
    try {
      await _googleSignIn.signIn();
    } catch (error) {
      print(error);
      rethrow;
    }
  }

  // Google sign out handler
  Future<void> googleSignOut() async {
    try {
      await _googleSignIn.signOut();
      _user = null;
      notifyListeners();
    } catch (error) {
      print(error);
    }
  }
}

extension TokenToString on GoogleSignInAuthentication {
  String toTokenString() {
    if (this != null) {
      return 'accessToken: $accessToken\nidToken: $idToken\nserverAuthCode: $serverAuthCode';
    } else {
      return '';
    }
  }
}
