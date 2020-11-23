import 'package:flutter/material.dart';
import 'package:mukgo/navigator/bottom_navigation.dart';

class TabNavigator extends StatelessWidget {
  TabNavigator(
      {this.navigatorKey, this.tabItem, this.routes, this.initialRoute});

  final GlobalKey<NavigatorState> navigatorKey;
  final TabItem tabItem;
  final Map<String, WidgetBuilder> routes;
  final String initialRoute;

  @override
  Widget build(BuildContext context) {
    return Navigator(
      key: navigatorKey,
      initialRoute: initialRoute,
      onGenerateRoute: (routeSettings) {
        return MaterialPageRoute(
          settings: routeSettings,
          builder: (context) => routes[routeSettings.name](context),
        );
      },
    );
  }
}
