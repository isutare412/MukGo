import 'package:pull_to_refresh/pull_to_refresh.dart';
import 'package:contra/utils/colors.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:provider/provider.dart';
import 'package:mukgo/user/user_model.dart';
import 'package:mukgo/user/user_badge.dart';
import 'package:mukgo/auth/auth_api.dart';

class UserDetailTestPage extends StatefulWidget {
  const UserDetailTestPage({Key key}) : super(key: key);

  @override
  _UserDetailTestPageState createState() => _UserDetailTestPageState();
}

class _UserDetailTestPageState extends State<UserDetailTestPage> {
  RefreshController _refreshController =
      RefreshController(initialRefresh: true);

  void _onRefresh() async {
    // monitor network fetch
    await context.read<UserModel>().fetch(heavy: true);
    // if failed,use refreshFailed()
    _refreshController.refreshCompleted();
  }

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    Provider.of<UserModel>(context, listen: false).fetch(heavy: true);
    return Material(
      child: SmartRefresher(
          enablePullDown: true,
          controller: _refreshController,
          onRefresh: _onRefresh,
          child: SingleChildScrollView(
              child: Container(
            color: white,
            child: Column(
              children: <Widget>[
                Container(
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                    children: <Widget>[
                      SizedBox(
                        height: 30,
                      ),
                      Center(
                        child: Consumer<UserModel>(
                            builder: (context, user, child) {
                          return SvgPicture.asset(
                            user.profileAsset(),
                            height: 320,
                            width: 320,
                          );
                        }),
                      ),
                    ],
                  ),
                ),
                Column(
                  crossAxisAlignment: CrossAxisAlignment.center,
                  children: <Widget>[
                    Padding(
                      padding: const EdgeInsets.only(
                          left: 24.0, right: 24.0, top: 12.0, bottom: 12.0),
                      child:
                          Consumer<UserModel>(builder: (context, user, child) {
                        return Text(
                          user.name ?? '로딩중..',
                          textAlign: TextAlign.center,
                          style: TextStyle(
                              fontSize: 36,
                              color: wood_smoke,
                              fontWeight: FontWeight.w800),
                        );
                      }),
                    ),
                    Padding(
                      padding: const EdgeInsets.only(
                          left: 24.0, right: 24.0, top: 12.0, bottom: 12.0),
                      child: Column(children: <Widget>[
                        Consumer<UserModel>(builder: (context, user, child) {
                          return Text(
                            user.level != null ? 'Lv.${user.level}' : '',
                            textAlign: TextAlign.center,
                            style: TextStyle(
                                fontSize: 30,
                                color: trout,
                                fontWeight: FontWeight.w500),
                          );
                        }),
                        Padding(
                          padding: const EdgeInsets.only(
                              top: 12.0, left: 40.0, right: 40.0),
                          child: Consumer<UserModel>(
                              builder: (context, user, child) {
                            return LinearProgressIndicator(
                              value: user.expRatio,
                              minHeight: 6.0,
                            );
                          }),
                        ),
                        Padding(
                          padding: const EdgeInsets.only(top: 4.0),
                          child: Consumer<UserModel>(
                              builder: (context, user, child) {
                            return Text(
                              user.level != null
                                  ? '${(user.expRatio * 100).toStringAsFixed(1)}%'
                                  : '',
                              textAlign: TextAlign.center,
                              style: TextStyle(
                                  fontSize: 16,
                                  color: trout,
                                  fontWeight: FontWeight.w500),
                            );
                          }),
                        )
                      ]),
                    ),
                    Padding(
                      padding: const EdgeInsets.only(bottom: 6),
                      child: Row(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          Consumer<UserModel>(builder: (context, user, child) {
                            return Text(
                              user.level != null
                                  ? '${user.likeCount} Likes '
                                  : '',
                              textAlign: TextAlign.center,
                              style: TextStyle(fontSize: 20),
                            );
                          }),
                          Icon(
                            Icons.favorite_border,
                            color: wood_smoke,
                          ),
                        ],
                      ),
                    ),
                    Padding(
                      padding: const EdgeInsets.only(bottom: 24),
                      child: Row(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          Consumer<UserModel>(builder: (context, user, child) {
                            return Text(
                              user.level != null
                                  ? '${user.reviewCount} Reviews '
                                  : '',
                              textAlign: TextAlign.center,
                              style: TextStyle(fontSize: 20),
                            );
                          }),
                          Icon(
                            Icons.rate_review,
                            color: wood_smoke,
                          ),
                        ],
                      ),
                    ),
                    Divider(
                      color: santas_gray_10,
                      thickness: 1,
                      indent: 24,
                      endIndent: 24,
                    ),
                    Consumer<UserModel>(builder: (context, user, child) {
                      var badges = user.restaurantTypeCounts != null
                          ? user.restaurantTypeCounts.map((e) {
                              return UserBadge(type: e.type, count: e.count);
                            }).toList()
                          : <UserBadge>[];
                      return BadgeGrid(
                        badges: badges,
                      );
                    }),
                    Divider(
                      color: santas_gray_10,
                      thickness: 1,
                      indent: 24,
                      endIndent: 24,
                    ),
                    Padding(
                      padding: const EdgeInsets.only(bottom: 12, top: 24),
                      child: RaisedButton(
                        padding: EdgeInsets.all(16),
                        color: lightening_yellow,
                        textColor: white,
                        onPressed: () {
                          googleSignOut(context);
                        },
                        child: Text(
                          'Logout',
                          style: TextStyle(
                              fontSize: 21.0, fontWeight: FontWeight.bold),
                        ),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(16.0),
                          // side: BorderSide(color: black, width: 2.0)
                        ),
                      ),
                    ),
                  ],
                ),
              ],
            ),
          ))),
    );
  }
}
