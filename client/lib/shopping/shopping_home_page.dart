import 'package:mukgo/shopping/shopping_home_page_one.dart';
import 'package:mukgo/shopping/shopping_home_page_two.dart';
import 'package:mukgo/utils/colors.dart';
import 'package:flutter/material.dart';

class ShoppingHomePage extends StatefulWidget {
  @override
  _ShoppingHomePageState createState() => _ShoppingHomePageState();
}

class _ShoppingHomePageState extends State<ShoppingHomePage> {
  int _currentIndex = 0;
  final List<Widget> _childrenWidgets = [
    ShoppingHomePageTwo(),
    ShoppingHomePageOne(),
    ShoppingHomePageTwo(),
    ShoppingHomePageOne(),
  ];

  void _onItemTapped(int index) {
    setState(() {
      _currentIndex = index;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: _childrenWidgets.elementAt(_currentIndex),
      ),
      bottomNavigationBar: BottomNavigationBar(
        items: [
          BottomNavigationBarItem(icon: Icon(Icons.home), title: Text("Home")),
          BottomNavigationBarItem(
              icon: Icon(Icons.search), title: Text("Search")),
          BottomNavigationBarItem(
              icon: Icon(Icons.chat_bubble), title: Text("Chat")),
          BottomNavigationBarItem(icon: Icon(Icons.info), title: Text("About")),
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
