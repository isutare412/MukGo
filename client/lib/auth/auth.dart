import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter/foundation.dart';
import 'package:google_sign_in/google_sign_in.dart';

class AuthModel extends ChangeNotifier {
  final GoogleSignIn _googleSignIn = GoogleSignIn(
    scopes: <String>['openid', 'profile', 'email'],
  );
  GoogleSignInAccount _user;

  GoogleSignIn get googleSignIn => _googleSignIn;
  GoogleSignInAccount get user => _user;

  // Set up google sign in
  AuthModel() : super() {
    _googleSignIn.onCurrentUserChanged.listen((GoogleSignInAccount account) {
      // Log for test
      account.authentication.then((value) {
        print(value.toTokenString());
        _user = account;
        notifyListeners();
      });
    });
    _googleSignIn.signInSilently();
  }

  // Sign in handler
  Future<void> signIn() async {
    try {
      await _googleSignIn.signIn();
    } catch (error) {
      print(error);
    }
  }

  // Sign out handler
  Future<void> signOut() async {
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
      return 'accessToken: ${this.accessToken}\nidToken: ${this.idToken}\nserverAuthCode: ${this.serverAuthCode}';
    } else {
      return '';
    }
  }
}
