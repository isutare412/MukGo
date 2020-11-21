import 'package:pull_to_refresh/pull_to_refresh.dart';
import 'package:contra/login/contra_text.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

import 'package:mukgo/auth/auth_api.dart';
import 'package:mukgo/ranking/ranking_user.dart';
import 'package:mukgo/ranking/ranking_list_item.dart';
import 'package:mukgo/api/api.dart';
import 'package:mukgo/user/user_model.dart';

class RankingListPage extends StatefulWidget {
  @override
  _RankingListPageState createState() => _RankingListPageState();
}

class _RankingListPageState extends State<RankingListPage> {
  RefreshController _refreshController =
      RefreshController(initialRefresh: false);

  List<RankingUser> users = [];

  void _update() async {
    var rawUsers = await fetchRankingData(readAuth(context).token);
    setState(() {
      users = rawUsers.users.asMap().entries.map((entry) {
        var idx = entry.key;
        var user = entry.value;
        return RankingUser(
            name: user.name,
            level: user.level,
            reviewCount: user.reviewCount,
            asset: levelToProfileAsset(user.level),
            rank: idx + 1);
      }).toList();
    });
  }

  void _onRefresh() async {
    // monitor network fetch
    await _update();
    // if failed,use refreshFailed()
    _refreshController.refreshCompleted();
  }

  @override
  void initState() {
    super.initState();
    Future.microtask(() {
      _update();
    });
  }

  @override
  Widget build(BuildContext context) {
    return SmartRefresher(
        enablePullDown: true,
        controller: _refreshController,
        onRefresh: _onRefresh,
        child: SingleChildScrollView(
          child: Column(
            children: <Widget>[
              Padding(
                  padding: const EdgeInsets.all(
                    24,
                  ),
                  child: ContraText(
                    size: 30.0,
                    text: 'Ranking',
                    alignment: Alignment.centerLeft,
                  )),
              ListView.builder(
                  shrinkWrap: true,
                  padding: EdgeInsets.all(24),
                  physics: NeverScrollableScrollPhysics(),
                  itemCount: users.length,
                  itemBuilder: (BuildContext context, int index) {
                    return RankingListItem(
                      user: users[index],
                    );
                  })
            ],
          ),
        ));
  }
}
