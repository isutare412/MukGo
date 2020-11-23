import 'package:flutter/material.dart';
import 'package:contra/utils/colors.dart';

enum TabItem { map, ranking, user }

Map<TabItem, String> tabName = {
  TabItem.map: 'map',
  TabItem.ranking: 'ranking',
  TabItem.user: 'my page',
};

Map<TabItem, int> tabIndex = {
  TabItem.map: 0,
  TabItem.ranking: 1,
  TabItem.user: 2,
};

class BottomNavigation extends StatelessWidget {
  BottomNavigation({this.currentTab, this.onSelectTab});
  final TabItem currentTab;
  final ValueChanged<TabItem> onSelectTab;

  @override
  Widget build(BuildContext context) {
    return BottomNavigationBar(
      items: [
        BottomNavigationBarItem(
            icon: Icon(Icons.map), title: Text(tabName[TabItem.map])),
        BottomNavigationBarItem(
            icon: Icon(Icons.insert_chart),
            title: Text(tabName[TabItem.ranking])),
        BottomNavigationBarItem(
            icon: Icon(Icons.person), title: Text(tabName[TabItem.user])),
      ],
      currentIndex: tabIndex[currentTab],
      onTap: (index) => onSelectTab(
        TabItem.values[index],
      ),
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
    );
  }
}
