import 'package:flutter/material.dart';

import 'package:contra/custom_widgets/button_round_with_shadow.dart';
import 'package:contra/utils/colors.dart';

import 'package:mukgo/project/sample_page.dart';
import 'package:mukgo/auth/auth_api.dart';
import 'package:mukgo/map/map_detail.dart';
import 'package:mukgo/user/user_detail.dart';

class MyHomePage extends StatefulWidget {
  MyHomePage({Key key, this.title}) : super(key: key);

  final String title;

  @override
  _MyHomePageState createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  int _currentIndex = 0;
  final List<Widget> _childrenWidgets = [
    MapDetailPage(),
    UserDetail(isBarChart: false),
  ];

   void _onItemTapped(int index) {
    setState(() {
      _currentIndex = index;
    });
  }

  @override
  Widget build(BuildContext context) {
    checkAuth(context);

    return Scaffold(
      appBar: AppBar(
        title: Text(widget.title),

        //-------For Contra Sample Pages (START) ------//
        actions: <Widget>[
          Padding(
            padding: EdgeInsets.only(right: 20.0),
            child: GestureDetector(
              onTap: () {
                Navigator.pushNamed(context, "/sample_page");
              },
              child: Icon(
                Icons.category,
                size: 26.0,
              ),
            )
          ),
        ],
        //-------For Contra Sample Pages (END) ------//
      ),
      body: Center(
        child: _childrenWidgets.elementAt(_currentIndex),
      ),
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
      bottomNavigationBar: BottomNavigationBar(
        items: [
          BottomNavigationBarItem(icon: Icon(Icons.map), title: Text("map")),
          BottomNavigationBarItem(
              icon: Icon(Icons.person), title: Text("My page")),
        ],
        currentIndex: _currentIndex,
        onTap: _onItemTapped,
        selectedItemColor: wood_smoke,
        unselectedItemColor: trout,
        showSelectedLabels: true,
        showUnselectedLabels: true,
        selectedIconTheme: IconThemeData(color: wood_smoke, opacity: 1),
        unselectedIconTheme: IconThemeData(color: trout, opacity: 0.6),
        selectedLabelStyle: TextStyle(
            color: wood_smoke, fontSize: 12, fontWeight: FontWeight.w800),
        unselectedLabelStyle:
            TextStyle(color: trout, fontSize: 12, fontWeight: FontWeight.w800),
      ),
    );
  }
}
