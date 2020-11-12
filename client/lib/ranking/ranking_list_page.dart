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
  List<RankingUser> users = [];

  @override
  void initState() {
    super.initState();
    Future.microtask(() async {
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
    });
  }

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
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
    );
  }
}
