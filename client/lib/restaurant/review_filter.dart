import 'package:flutter/material.dart';
import 'package:contra/utils/colors.dart';
import 'package:contra/login/contra_text.dart';
import 'package:contra/custom_widgets/button_plain.dart';
import 'package:contra/custom_widgets/custom_app_bar.dart';
import 'package:contra/custom_widgets/button_round_with_shadow.dart';

class Order {
  String key;
  bool ascending;

  Order({this.key, this.ascending});
}

class Filter {
  int numPeople;
  String wait; // disable, true, false
  Order order;

  Filter({this.numPeople, this.wait, this.order});
}

typedef FilterCallback = void Function(Filter);

class FilterOpenWidget extends StatelessWidget {
  final VoidCallback onTap;

  FilterOpenWidget({this.onTap});

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: onTap,
      child: Container(
        padding: EdgeInsets.all(24),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceBetween,
          children: <Widget>[
            ContraText(
              text: 'Filter',
              alignment: Alignment.centerLeft,
              color: wood_smoke,
              size: 21,
            ),
            Icon(
              Icons.arrow_forward_ios,
              color: wood_smoke,
            )
          ],
        ),
      ),
    );
  }
}

class FilterModal extends StatefulWidget {
  FilterModal({this.callback, this.preFilter});
  final FilterCallback callback;
  final Filter preFilter;

  @override
  _FilterModalState createState() => _FilterModalState();
}

class _FilterModalState extends State<FilterModal> {
  Filter _filter;

  @override
  void initState() {
    super.initState();

    _filter = widget.preFilter;
  }

  @override
  Widget build(BuildContext context) {
    List<bool> waitSelectList = [
      _filter.wait == 'true',
      _filter.wait == 'false'
    ];

    return Scaffold(
      appBar: CustomAppBar(
        height: 80,
        child: Row(
          mainAxisAlignment: MainAxisAlignment.spaceEvenly,
          crossAxisAlignment: CrossAxisAlignment.end,
          children: <Widget>[
            // Expanded(
            //   flex: 1,
            //   child: Padding(
            //     padding: const EdgeInsets.only(left: 24.0),
            //     child: Align(
            //       alignment: Alignment.bottomLeft,
            //       child: ButtonRoundWithShadow(
            //           size: 48,
            //           borderColor: wood_smoke,
            //           color: white,
            //           callback: () {
            //             Navigator.pop(context);
            //           },
            //           shadowColor: wood_smoke,
            //           iconPath: "assets/icons/arrow_back.svg"),
            //     ),
            //   ),
            // ),
            Expanded(
              flex: 2,
              child: ContraText(
                size: 27,
                color: wood_smoke,
                alignment: Alignment.bottomCenter,
                text: "Filter",
              ),
            ),
            // Expanded(
            //   flex: 1,
            //   child: SizedBox(
            //     width: 20,
            //   ),
            // )
          ],
        ),
      ),
      body: Container(
        padding: EdgeInsets.symmetric(horizontal: 24, vertical: 36),
        child: SingleChildScrollView(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.center,
            children: <Widget>[
              Text(
                '방문 인원',
                style: TextStyle(
                    fontSize: 24,
                    fontWeight: FontWeight.bold,
                    color: wood_smoke),
              ),
              SizedBox(
                height: 16,
              ),
              Slider(
                value: _filter.numPeople.toDouble(),
                min: 0,
                max: 10,
                divisions: 10,
                label: _filter.numPeople.toString(),
                onChanged: (double value) {
                  setState(() {
                    _filter.numPeople = value.round();
                  });
                },
              ),
              SizedBox(
                height: 32,
              ),
              Text(
                '대기 유무',
                style: TextStyle(
                    fontSize: 24,
                    fontWeight: FontWeight.bold,
                    color: wood_smoke),
              ),
              SizedBox(
                height: 16,
              ),
              ToggleButtons(
                children: <Widget>[
                  Padding(
                    padding: EdgeInsets.all(12),
                    child: Text(
                      '기다렸어요',
                      style: TextStyle(fontSize: 16),
                    ),
                  ),
                  Padding(
                    padding: EdgeInsets.all(12),
                    child: Text(
                      '바로먹었어요',
                      style: TextStyle(fontSize: 16),
                    ),
                  ),
                ],
                onPressed: (int index) {
                  setState(() {
                    if (index == 0) {
                      if (_filter.wait == 'true') {
                        _filter.wait = 'disable';
                      } else {
                        _filter.wait = 'true';
                      }
                    } else {
                      if (_filter.wait == 'false') {
                        _filter.wait = 'disable';
                      } else {
                        _filter.wait = 'false';
                      }
                    }
                  });
                },
                isSelected: waitSelectList,
                color: trout,
                selectedColor: persian_blue,
                selectedBorderColor: persian_blue,
                borderRadius: BorderRadius.all(Radius.circular(10)),
              ),
              SizedBox(
                height: 32,
              ),
              Text(
                '정렬',
                style: TextStyle(
                    fontSize: 24,
                    fontWeight: FontWeight.bold,
                    color: wood_smoke),
              ),
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: <Widget>[
                  Padding(
                      padding: EdgeInsets.only(right: 24),
                      child: DropdownButton<String>(
                        value: _filter.order.key,
                        elevation: 16,
                        style: TextStyle(
                            fontSize: 21,
                            fontWeight: FontWeight.bold,
                            color: trout),
                        underline: Container(
                          height: 2,
                          color: persian_blue,
                        ),
                        onChanged: (String newValue) {
                          setState(() {
                            _filter.order.key = newValue;
                          });
                        },
                        items: <String>['time', 'score']
                            .map<DropdownMenuItem<String>>((String value) {
                          String text;
                          if (value == 'score') {
                            text = '점수';
                          } else {
                            text = '시간';
                          }
                          return DropdownMenuItem<String>(
                            value: value,
                            child: Text(text),
                          );
                        }).toList(),
                      )),
                  Padding(
                      padding: EdgeInsets.only(left: 24),
                      child: DropdownButton<bool>(
                        value: _filter.order.ascending,
                        elevation: 16,
                        style: TextStyle(
                            fontSize: 21,
                            fontWeight: FontWeight.bold,
                            color: trout),
                        underline: Container(
                          height: 2,
                          color: persian_blue,
                        ),
                        onChanged: (bool newValue) {
                          setState(() {
                            _filter.order.ascending = newValue;
                          });
                        },
                        items: <bool>[true, false]
                            .map<DropdownMenuItem<bool>>((bool value) {
                          var text = value ? '오름차순' : '내림차순';
                          return DropdownMenuItem<bool>(
                            value: value,
                            child: Text(text),
                          );
                        }).toList(),
                      )),
                ],
              ),
              Padding(
                  padding: const EdgeInsets.only(left: 24, right: 24, top: 64),
                  child: ButtonPlain(
                      color: wood_smoke,
                      borderColor: wood_smoke,
                      textColor: white,
                      text: 'Apply filter',
                      textSize: 21,
                      onTap: () {
                        widget.callback(_filter);
                        Navigator.pop(context);
                      })),
            ],
          ),
        ),
      ),
    );
  }
}
