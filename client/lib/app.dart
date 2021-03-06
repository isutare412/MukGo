import 'package:flutter/material.dart';
import 'package:mukgo/auth/auth_api.dart';
import 'package:mukgo/navigator/bottom_navigation.dart';
import 'package:mukgo/navigator/tab_navigator.dart';

// pages
import 'package:mukgo/map/map_detail.dart';
import 'package:mukgo/restaurant/restaurant_detail.dart';
import 'package:mukgo/user/user_detail.dart';
import 'package:mukgo/ranking/ranking_list_page.dart';

Map<String, WidgetBuilder> routes = {
  '/map': (context) => MapDetailPage(),
  '/user': (context) => UserDetailTestPage(),
  '/restaurant': (context) => RestaurantDetailTestPage(),
  '/ranking': (context) => RankingListPage(),
};

class App extends StatefulWidget {
  App({Key key, this.title = 'Mukgo Project'}) : super(key: key);

  final String title;

  @override
  State<StatefulWidget> createState() => AppState();
}

class AppState extends State<App> {
  TabItem _currentTab = TabItem.map;
  final Map<TabItem, GlobalKey<NavigatorState>> _navigatorKeys = {
    TabItem.map: GlobalKey<NavigatorState>(),
    TabItem.user: GlobalKey<NavigatorState>(),
  };

  void _selectTab(TabItem tabItem) {
    if (tabItem == _currentTab) {
      // pop to first route
      _navigatorKeys[tabItem].currentState.popUntil((route) => route.isFirst);
    } else {
      setState(() => _currentTab = tabItem);
    }
  }

  @override
  Widget build(BuildContext context) {
    checkAuth(context);

    return WillPopScope(
      onWillPop: () async {
        final isFirstRouteInCurrentTab =
            !await _navigatorKeys[_currentTab].currentState.maybePop();
        if (isFirstRouteInCurrentTab) {
          // if not on the 'map' tab
          if (_currentTab != TabItem.map) {
            // select 'map' tab
            _selectTab(TabItem.map);
            // back button handled by app
            return false;
          }
        }
        // let system handle back button if we're on the first route
        return isFirstRouteInCurrentTab;
      },
      child: Scaffold(
        appBar: AppBar(
          title: Text(widget.title),
        ),
        body: Stack(children: <Widget>[
          _buildOffstageNavigator(TabItem.map, '/map'),
          _buildOffstageNavigator(TabItem.ranking, '/ranking'),
          _buildOffstageNavigator(TabItem.user, '/user'),
        ]),
        bottomNavigationBar: BottomNavigation(
          currentTab: _currentTab,
          onSelectTab: _selectTab,
        ),
      ),
    );
  }

  Widget _buildOffstageNavigator(TabItem tabItem, String initialRoute) {
    var newRoutes = Map<String, WidgetBuilder>.from(routes);
    newRoutes['/'] = routes[initialRoute];
    return Offstage(
      offstage: _currentTab != tabItem,
      child: TabNavigator(
        navigatorKey: _navigatorKeys[tabItem],
        tabItem: tabItem,
        initialRoute: '/',
        routes: newRoutes,
      ),
    );
  }
}
