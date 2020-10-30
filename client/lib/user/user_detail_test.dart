import 'package:contra/utils/colors.dart';
import 'package:contra/custom_widgets/button_plain.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';
import 'package:provider/provider.dart';
import 'package:mukgo/user/user_model.dart';
import 'package:mukgo/review/review_detail_test.dart';

class UserDetailTestPage extends StatefulWidget {
  const UserDetailTestPage({Key key}) : super(key: key);

  @override
  _UserDetailTestPageState createState() => _UserDetailTestPageState();
}

class _UserDetailTestPageState extends State<UserDetailTestPage> {
  @override
  void initState() {
    super.initState();
    Future.microtask(() {
      // fetch user info after randering
      return context.read<UserModel>().fetch();
    });
  }

  @override
  Widget build(BuildContext context) {
    return Material(
      child: Container(
        color: white,
        child: Column(
          children: <Widget>[
            Expanded(
              flex: 10,
              child: Container(
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                  children: <Widget>[
                    SizedBox(
                      height: 30,
                    ),
                    Center(
                      child:
                          Consumer<UserModel>(builder: (context, user, child) {
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
            ),
            Expanded(
              flex: 8,
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.center,
                children: <Widget>[
                  Padding(
                    padding: const EdgeInsets.only(
                        left: 24.0, right: 24.0, top: 12.0, bottom: 12.0),
                    child: Consumer<UserModel>(builder: (context, user, child) {
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
                  SizedBox(
                    width: 120.0,
                    child: ButtonPlain(
                      color: google_red,
                      textColor: white,
                      text: 'Refresh',
                      onTap: () async {
                        var user = context.read<UserModel>();
                        user.clear();
                        await user.fetch();
                        print('reloaded');
                      },
                    ),
                  ),
                  Padding(
                    padding: const EdgeInsets.only(top: 10.0),
                    child: SizedBox(
                      width: 120.0,
                      child: ButtonPlain(
                        color: lightening_yellow,
                        textColor: white,
                        text: 'Test Review',
                        onTap: () {
                          Navigator.pushNamed(context, "/review-list",
                              arguments: ReviewPageArguments(
                                  id: '5f915970fc1495705932c25a',
                                  name: 'my home'));
                        },
                      ),
                    ),
                  )
                ],
              ),
            ),
            Expanded(
              flex: 1,
              child: Container(),
            )
          ],
        ),
      ),
    );
  }
}
