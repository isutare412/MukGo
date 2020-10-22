import 'dart:async';
import 'package:flutter/material.dart';
import 'package:contra/onboarding/onboard_main.dart';

import 'package:mukgo/review/review_detail_test.dart';

class SamplePage extends StatefulWidget {
  @override
  _SamplePageState createState() => _SamplePageState();
}

class _SamplePageState extends State<SamplePage> {
  @override
  Widget build(BuildContext context) {
    //checkAuth(context);

    return Scaffold(
      appBar: AppBar(
        title: Text('contra sample page'),
      ),
      body: SingleChildScrollView(
        child: Container(
          padding: EdgeInsets.all(10),
          child: Column(
            children: <Widget>[
              ListTile(
                contentPadding: EdgeInsets.all(20),
                trailing: Icon(Icons.navigate_next),
                title: Text("Review"),
                onTap: () {
                  Navigator.pushNamed(context, "/project_review",
                      arguments: ReviewPageArguments(
                          id: '5f915970fc1495705932c25a', name: 'my home'));
                },
              ),
              ListTile(
                contentPadding: EdgeInsets.all(20),
                trailing: Icon(Icons.navigate_next),
                title: Text("Onboarding and Splash"),
                onTap: () {
                  Navigator.pushNamed(context, "/onboard_all");
                  /*
                  Navigator.push(
                      context,
                      MaterialPageRoute(builder: (context) => OnboardPageMain()),
                  );
                  */
                },
              ),
              ListTile(
                contentPadding: EdgeInsets.all(20),
                trailing: Icon(Icons.navigate_next),
                title: Text("Forms, Login, Signup"),
                onTap: () {
                  Navigator.pushNamed(context, "/login_all");
                },
              ),
              ListTile(
                contentPadding: EdgeInsets.all(20),
                trailing: Icon(Icons.navigate_next),
                title: Text("Chatting Screens"),
                onTap: () {
                  Navigator.pushNamed(context, "/chat_home");
                },
              ),
              ListTile(
                contentPadding: EdgeInsets.all(20),
                trailing: Icon(Icons.navigate_next),
                title: Text("Shopping Screens"),
                onTap: () {
                  Navigator.pushNamed(context, "/shopping_main_page");
                },
              ),
              ListTile(
                contentPadding: EdgeInsets.all(20),
                trailing: Icon(Icons.navigate_next),
                title: Text("Blog Screens"),
                onTap: () {
                  Navigator.pushNamed(context, "/blog_main_page");
                },
              ),
              Container(
                child: ListTile(
                  contentPadding: EdgeInsets.all(20),
                  trailing: Icon(Icons.navigate_next),
                  title: Text("Payment"),
                  onTap: () {
                    Navigator.pushNamed(context, "/payment_main_page");
                  },
                ),
              ),
              Container(
                child: ListTile(
                  contentPadding: EdgeInsets.all(20),
                  trailing: Icon(Icons.navigate_next),
                  title: Text("Alarm, Clock, Weather"),
                  onTap: () {
                    Navigator.pushNamed(context, "/alarm_main_page");
                  },
                ),
              ),
              Container(
                child: ListTile(
                  contentPadding: EdgeInsets.all(20),
                  trailing: Icon(Icons.navigate_next),
                  title: Text("Data and Statistics"),
                  onTap: () {
                    Navigator.pushNamed(context, "/chart_main_page");
                  },
                ),
              ),
              Container(
                child: ListTile(
                  contentPadding: EdgeInsets.all(20),
                  trailing: Icon(Icons.navigate_next),
                  title: Text("Location And Map"),
                  onTap: () {
                    Navigator.pushNamed(context, "/map_main_page");
                  },
                ),
              ),
              Container(
                child: ListTile(
                  contentPadding: EdgeInsets.all(20),
                  trailing: Icon(Icons.navigate_next),
                  title: Text("Content, Text Details"),
                  onTap: () {
                    Navigator.pushNamed(context, "/content_text_main_page");
                  },
                ),
              ),
              Container(
                child: ListTile(
                  contentPadding: EdgeInsets.all(20),
                  trailing: Icon(Icons.navigate_next),
                  title: Text("Menu and Settings"),
                  onTap: () {
                    Navigator.pushNamed(context, "/menu_settings_main_page");
                  },
                ),
              ),
/*              Container(
                child: ListTile(
                  contentPadding: EdgeInsets.all(20),
                  trailing: Icon(Icons.navigate_next),
                  title: Text("Dialogs, Filters, Toasts"),
                  onTap: () {
                    Navigator.pushNamed(context, "/empty_state");
                  },
                ),
              ),
              Container(
                child: ListTile(
                  contentPadding: EdgeInsets.all(20),
                  trailing: Icon(Icons.navigate_next),
                  title: Text("Profile"),
                  onTap: () {
                    Navigator.pushNamed(context, "/empty_state");
                  },
                ),
              ),
              Container(
                child: ListTile(
                  contentPadding: EdgeInsets.all(20),
                  trailing: Icon(Icons.navigate_next),
                  title: Text("Menus"),
                  onTap: () {
                    Navigator.pushNamed(context, "/empty_state");
                  },
                ),
              ),
              ListTile(
                contentPadding: EdgeInsets.all(20),
                trailing: Icon(Icons.navigate_next),
                title: Text("Profile"),
                onTap: () {
                  Navigator.pushNamed(context, "/empty_state");
                },
              )*/
            ],
          ),
        ),
      ),
      /*
      floatingActionButton: Align(
        alignment: Alignment.bottomRight,
        child: Padding(
          padding: const EdgeInsets.all(24.0),
          child: ButtonRoundWithShadow(
              size: 60,
              borderColor: wood_smoke,
              color: white,
              callback: () {
                googleSignOut(context);
              },
              shadowColor: wood_smoke,
              iconPath: "assets/icons/ic_add.svg"),
        ),
      ),
      */
    );
  }
}
