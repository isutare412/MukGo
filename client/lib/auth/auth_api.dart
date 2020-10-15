import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:google_sign_in/google_sign_in.dart';
import 'package:mukgo/auth/auth.dart';

/// Retrieving
///
/// state 값을 가져오기만 하는 함수들

// build 바깥에서 호출되게!
Future<void> googleSignIn(BuildContext context) async {
  var auth = context.read<AuthModel>();
  return auth.signIn();
}

// build 바깥에서 호출되게!
Future<void> googleSignOut(BuildContext context) async {
  var auth = context.read<AuthModel>();
  return auth.signOut();
}

// build 바깥에서 호출되게!
AuthModel readAuth(BuildContext context) {
  return context.read<AuthModel>();
}

// build 안에서 호출되게!
AuthModel getAuth(BuildContext context) {
  return Provider.of<AuthModel>(context, listen: false);
}

/// Listening
///
/// state의 변화가 생기면 build를 재호출시키는 함수들

// build 안에서 호출되게!
// watch의 경우 model의 일부분이라도 변경된다면 build가 재호출됨
AuthModel watchAuth(BuildContext context) {
  return context.watch<AuthModel>();
}

// build 안에서 호출되게!
// select의 경우 user(필드의 값)이 변경된다면 build가 재호출됨
GoogleSignInAccount selectUser(BuildContext context) {
  var user =
      context.select<AuthModel, GoogleSignInAccount>((auth) => auth.user);
  return user;
}

// 인증이 필요한 widget의 *build* method안에서 실행시켜야 함
void checkAuth(BuildContext context) {
  // user가 변경되면 build가 재실행됨
  if (selectUser(context) == null) {
    Future.microtask(() {
      // login page로 강제 이동
      return Navigator.pushNamed(context, '/project_login');
    });
  }
}
