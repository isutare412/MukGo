import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:contra/utils/colors.dart';
import 'package:mukgo/auth/auth_model.dart';
import 'package:mukgo/auth/login_page.dart';
import 'package:mukgo/user/user_model.dart';
import 'package:mukgo/loading.dart';
import 'package:mukgo/app.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    var child = MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Contra Flutter Kit',
      theme: ThemeData(
          fontFamily: 'Montserrat',
          primarySwatch: Colors.blue,
          primaryColor: persian_blue),
      home: LoadingScreen(),
      routes: {
        '/project': (context) => App(title: 'Mukgo Project'),
        '/project_login': (context) => LoginForm(),
        '/loading': (context) => LoadingScreen(),
      },
    );

    return MultiProvider(
      providers: [
        ChangeNotifierProvider(create: (context) => AuthModel(), child: child),
        ChangeNotifierProxyProvider<AuthModel, UserModel>(
          create: (context) => UserModel(),
          update: (context, auth, user) {
            user.auth = auth;
            user.fetch();
            return user;
          },
        )
      ],
      child: child,
    );
  }
}
