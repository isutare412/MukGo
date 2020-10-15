import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class ProjectPageMain extends StatefulWidget {
  @override
  _ProjectPageMainState createState() => _ProjectPageMainState();
}

class _ProjectPageMainState extends State<ProjectPageMain> {
  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Project Pages"),
      ),
      body: SingleChildScrollView(
        child: Column(
          children: <Widget>[
            ListTile(
              contentPadding: EdgeInsets.all(20),
              trailing: Icon(Icons.navigate_next),
              title: Text("Map"),
              onTap: () {
                Navigator.pushNamed(context, "/project_map");
              },
            ),
            ListTile(
              contentPadding: EdgeInsets.all(20),
              trailing: Icon(Icons.navigate_next),
              title: Text("Restaurant"),
              onTap: () {
                Navigator.pushNamed(context, "/project_restaurant");
              },
            ),
            ListTile(
              contentPadding: EdgeInsets.all(20),
              trailing: Icon(Icons.navigate_next),
              title: Text("Review"),
              onTap: () {
                Navigator.pushNamed(context, "/project_review");
              },
            ),
            Container(
              child: ListTile(
                contentPadding: EdgeInsets.all(20),
                trailing: Icon(Icons.navigate_next),
                title: Text("User"),
                onTap: () {
                  Navigator.pushNamed(context, "/project_user");
                },
              ),
            ),
            Container(
              child: ListTile(
                contentPadding: EdgeInsets.all(20),
                trailing: Icon(Icons.navigate_next),
                title: Text("Login"),
                onTap: () {
                  Navigator.pushNamed(context, "/project_login");
                },
              ),
            ),
          ],
        ),
      ),
    );
  }
}
