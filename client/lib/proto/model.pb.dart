///
//  Generated code. Do not modify.
//  source: proto/model.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import 'model.pbenum.dart';

export 'model.pbenum.dart';

class User extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('User', package: const $pb.PackageName('proto'), createEmptyInstance: create)
    ..aOS(1, 'id')
    ..aOS(2, 'name')
    ..a<$core.int>(3, 'level', $pb.PbFieldType.O3)
    ..aInt64(4, 'totalExp')
    ..aInt64(5, 'levelExp')
    ..aInt64(6, 'curExp')
    ..a<$core.double>(7, 'expRatio', $pb.PbFieldType.OD)
    ..a<$core.double>(8, 'sightRadius', $pb.PbFieldType.OD)
    ..a<$core.int>(9, 'reviewCount', $pb.PbFieldType.O3)
    ..a<$core.int>(10, 'likeCount', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  User._() : super();
  factory User() => create();
  factory User.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory User.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  User clone() => User()..mergeFromMessage(this);
  User copyWith(void Function(User) updates) => super.copyWith((message) => updates(message as User));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static User create() => User._();
  User createEmptyInstance() => create();
  static $pb.PbList<User> createRepeated() => $pb.PbList<User>();
  @$core.pragma('dart2js:noInline')
  static User getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<User>(create);
  static User _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get id => $_getSZ(0);
  @$pb.TagNumber(1)
  set id($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);

  @$pb.TagNumber(3)
  $core.int get level => $_getIZ(2);
  @$pb.TagNumber(3)
  set level($core.int v) { $_setSignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasLevel() => $_has(2);
  @$pb.TagNumber(3)
  void clearLevel() => clearField(3);

  @$pb.TagNumber(4)
  $fixnum.Int64 get totalExp => $_getI64(3);
  @$pb.TagNumber(4)
  set totalExp($fixnum.Int64 v) { $_setInt64(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasTotalExp() => $_has(3);
  @$pb.TagNumber(4)
  void clearTotalExp() => clearField(4);

  @$pb.TagNumber(5)
  $fixnum.Int64 get levelExp => $_getI64(4);
  @$pb.TagNumber(5)
  set levelExp($fixnum.Int64 v) { $_setInt64(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasLevelExp() => $_has(4);
  @$pb.TagNumber(5)
  void clearLevelExp() => clearField(5);

  @$pb.TagNumber(6)
  $fixnum.Int64 get curExp => $_getI64(5);
  @$pb.TagNumber(6)
  set curExp($fixnum.Int64 v) { $_setInt64(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasCurExp() => $_has(5);
  @$pb.TagNumber(6)
  void clearCurExp() => clearField(6);

  @$pb.TagNumber(7)
  $core.double get expRatio => $_getN(6);
  @$pb.TagNumber(7)
  set expRatio($core.double v) { $_setDouble(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasExpRatio() => $_has(6);
  @$pb.TagNumber(7)
  void clearExpRatio() => clearField(7);

  @$pb.TagNumber(8)
  $core.double get sightRadius => $_getN(7);
  @$pb.TagNumber(8)
  set sightRadius($core.double v) { $_setDouble(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasSightRadius() => $_has(7);
  @$pb.TagNumber(8)
  void clearSightRadius() => clearField(8);

  @$pb.TagNumber(9)
  $core.int get reviewCount => $_getIZ(8);
  @$pb.TagNumber(9)
  set reviewCount($core.int v) { $_setSignedInt32(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasReviewCount() => $_has(8);
  @$pb.TagNumber(9)
  void clearReviewCount() => clearField(9);

  @$pb.TagNumber(10)
  $core.int get likeCount => $_getIZ(9);
  @$pb.TagNumber(10)
  set likeCount($core.int v) { $_setSignedInt32(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasLikeCount() => $_has(9);
  @$pb.TagNumber(10)
  void clearLikeCount() => clearField(10);
}

class Coordinate extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Coordinate', package: const $pb.PackageName('proto'), createEmptyInstance: create)
    ..a<$core.double>(1, 'latitude', $pb.PbFieldType.OD)
    ..a<$core.double>(2, 'longitude', $pb.PbFieldType.OD)
    ..hasRequiredFields = false
  ;

  Coordinate._() : super();
  factory Coordinate() => create();
  factory Coordinate.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Coordinate.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Coordinate clone() => Coordinate()..mergeFromMessage(this);
  Coordinate copyWith(void Function(Coordinate) updates) => super.copyWith((message) => updates(message as Coordinate));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Coordinate create() => Coordinate._();
  Coordinate createEmptyInstance() => create();
  static $pb.PbList<Coordinate> createRepeated() => $pb.PbList<Coordinate>();
  @$core.pragma('dart2js:noInline')
  static Coordinate getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Coordinate>(create);
  static Coordinate _defaultInstance;

  @$pb.TagNumber(1)
  $core.double get latitude => $_getN(0);
  @$pb.TagNumber(1)
  set latitude($core.double v) { $_setDouble(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasLatitude() => $_has(0);
  @$pb.TagNumber(1)
  void clearLatitude() => clearField(1);

  @$pb.TagNumber(2)
  $core.double get longitude => $_getN(1);
  @$pb.TagNumber(2)
  set longitude($core.double v) { $_setDouble(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLongitude() => $_has(1);
  @$pb.TagNumber(2)
  void clearLongitude() => clearField(2);
}

class Restaurant extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Restaurant', package: const $pb.PackageName('proto'), createEmptyInstance: create)
    ..aOS(1, 'id')
    ..aOS(2, 'name')
    ..aOM<Coordinate>(3, 'coord', subBuilder: Coordinate.create)
    ..hasRequiredFields = false
  ;

  Restaurant._() : super();
  factory Restaurant() => create();
  factory Restaurant.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Restaurant.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Restaurant clone() => Restaurant()..mergeFromMessage(this);
  Restaurant copyWith(void Function(Restaurant) updates) => super.copyWith((message) => updates(message as Restaurant));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Restaurant create() => Restaurant._();
  Restaurant createEmptyInstance() => create();
  static $pb.PbList<Restaurant> createRepeated() => $pb.PbList<Restaurant>();
  @$core.pragma('dart2js:noInline')
  static Restaurant getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Restaurant>(create);
  static Restaurant _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get id => $_getSZ(0);
  @$pb.TagNumber(1)
  set id($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);

  @$pb.TagNumber(3)
  Coordinate get coord => $_getN(2);
  @$pb.TagNumber(3)
  set coord(Coordinate v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasCoord() => $_has(2);
  @$pb.TagNumber(3)
  void clearCoord() => clearField(3);
  @$pb.TagNumber(3)
  Coordinate ensureCoord() => $_ensure(2);
}

class Review extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Review', package: const $pb.PackageName('proto'), createEmptyInstance: create)
    ..aOS(1, 'reviewId')
    ..aOS(2, 'userId')
    ..aOS(3, 'userName')
    ..a<$core.int>(4, 'score', $pb.PbFieldType.O3)
    ..aOS(5, 'comment')
    ..pPS(6, 'menus')
    ..aOB(7, 'wait')
    ..a<$core.int>(8, 'numPeople', $pb.PbFieldType.O3)
    ..aInt64(9, 'timestamp')
    ..a<$core.int>(10, 'userLevel', $pb.PbFieldType.O3)
    ..hasRequiredFields = false
  ;

  Review._() : super();
  factory Review() => create();
  factory Review.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Review.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Review clone() => Review()..mergeFromMessage(this);
  Review copyWith(void Function(Review) updates) => super.copyWith((message) => updates(message as Review));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Review create() => Review._();
  Review createEmptyInstance() => create();
  static $pb.PbList<Review> createRepeated() => $pb.PbList<Review>();
  @$core.pragma('dart2js:noInline')
  static Review getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Review>(create);
  static Review _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get reviewId => $_getSZ(0);
  @$pb.TagNumber(1)
  set reviewId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasReviewId() => $_has(0);
  @$pb.TagNumber(1)
  void clearReviewId() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get userId => $_getSZ(1);
  @$pb.TagNumber(2)
  set userId($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasUserId() => $_has(1);
  @$pb.TagNumber(2)
  void clearUserId() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get userName => $_getSZ(2);
  @$pb.TagNumber(3)
  set userName($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasUserName() => $_has(2);
  @$pb.TagNumber(3)
  void clearUserName() => clearField(3);

  @$pb.TagNumber(4)
  $core.int get score => $_getIZ(3);
  @$pb.TagNumber(4)
  set score($core.int v) { $_setSignedInt32(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasScore() => $_has(3);
  @$pb.TagNumber(4)
  void clearScore() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get comment => $_getSZ(4);
  @$pb.TagNumber(5)
  set comment($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasComment() => $_has(4);
  @$pb.TagNumber(5)
  void clearComment() => clearField(5);

  @$pb.TagNumber(6)
  $core.List<$core.String> get menus => $_getList(5);

  @$pb.TagNumber(7)
  $core.bool get wait => $_getBF(6);
  @$pb.TagNumber(7)
  set wait($core.bool v) { $_setBool(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasWait() => $_has(6);
  @$pb.TagNumber(7)
  void clearWait() => clearField(7);

  @$pb.TagNumber(8)
  $core.int get numPeople => $_getIZ(7);
  @$pb.TagNumber(8)
  set numPeople($core.int v) { $_setSignedInt32(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasNumPeople() => $_has(7);
  @$pb.TagNumber(8)
  void clearNumPeople() => clearField(8);

  @$pb.TagNumber(9)
  $fixnum.Int64 get timestamp => $_getI64(8);
  @$pb.TagNumber(9)
  set timestamp($fixnum.Int64 v) { $_setInt64(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasTimestamp() => $_has(8);
  @$pb.TagNumber(9)
  void clearTimestamp() => clearField(9);

  @$pb.TagNumber(10)
  $core.int get userLevel => $_getIZ(9);
  @$pb.TagNumber(10)
  set userLevel($core.int v) { $_setSignedInt32(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasUserLevel() => $_has(9);
  @$pb.TagNumber(10)
  void clearUserLevel() => clearField(10);
}

class Users extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Users', package: const $pb.PackageName('proto'), createEmptyInstance: create)
    ..pc<User>(1, 'users', $pb.PbFieldType.PM, subBuilder: User.create)
    ..hasRequiredFields = false
  ;

  Users._() : super();
  factory Users() => create();
  factory Users.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Users.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Users clone() => Users()..mergeFromMessage(this);
  Users copyWith(void Function(Users) updates) => super.copyWith((message) => updates(message as Users));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Users create() => Users._();
  Users createEmptyInstance() => create();
  static $pb.PbList<Users> createRepeated() => $pb.PbList<Users>();
  @$core.pragma('dart2js:noInline')
  static Users getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Users>(create);
  static Users _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<User> get users => $_getList(0);
}

class Restaurants extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Restaurants', package: const $pb.PackageName('proto'), createEmptyInstance: create)
    ..pc<Restaurant>(1, 'restaurants', $pb.PbFieldType.PM, subBuilder: Restaurant.create)
    ..hasRequiredFields = false
  ;

  Restaurants._() : super();
  factory Restaurants() => create();
  factory Restaurants.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Restaurants.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Restaurants clone() => Restaurants()..mergeFromMessage(this);
  Restaurants copyWith(void Function(Restaurants) updates) => super.copyWith((message) => updates(message as Restaurants));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Restaurants create() => Restaurants._();
  Restaurants createEmptyInstance() => create();
  static $pb.PbList<Restaurants> createRepeated() => $pb.PbList<Restaurants>();
  @$core.pragma('dart2js:noInline')
  static Restaurants getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Restaurants>(create);
  static Restaurants _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Restaurant> get restaurants => $_getList(0);
}

class Reviews extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Reviews', package: const $pb.PackageName('proto'), createEmptyInstance: create)
    ..pc<Review>(1, 'reviews', $pb.PbFieldType.PM, subBuilder: Review.create)
    ..hasRequiredFields = false
  ;

  Reviews._() : super();
  factory Reviews() => create();
  factory Reviews.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Reviews.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Reviews clone() => Reviews()..mergeFromMessage(this);
  Reviews copyWith(void Function(Reviews) updates) => super.copyWith((message) => updates(message as Reviews));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Reviews create() => Reviews._();
  Reviews createEmptyInstance() => create();
  static $pb.PbList<Reviews> createRepeated() => $pb.PbList<Reviews>();
  @$core.pragma('dart2js:noInline')
  static Reviews getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Reviews>(create);
  static Reviews _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<Review> get reviews => $_getList(0);
}

class ErrorReason extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ErrorReason', package: const $pb.PackageName('proto'), createEmptyInstance: create)
    ..e<Code>(1, 'code', $pb.PbFieldType.OE, defaultOrMaker: Code.SUCCESS, valueOf: Code.valueOf, enumValues: Code.values)
    ..hasRequiredFields = false
  ;

  ErrorReason._() : super();
  factory ErrorReason() => create();
  factory ErrorReason.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ErrorReason.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ErrorReason clone() => ErrorReason()..mergeFromMessage(this);
  ErrorReason copyWith(void Function(ErrorReason) updates) => super.copyWith((message) => updates(message as ErrorReason));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ErrorReason create() => ErrorReason._();
  ErrorReason createEmptyInstance() => create();
  static $pb.PbList<ErrorReason> createRepeated() => $pb.PbList<ErrorReason>();
  @$core.pragma('dart2js:noInline')
  static ErrorReason getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ErrorReason>(create);
  static ErrorReason _defaultInstance;

  @$pb.TagNumber(1)
  Code get code => $_getN(0);
  @$pb.TagNumber(1)
  set code(Code v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasCode() => $_has(0);
  @$pb.TagNumber(1)
  void clearCode() => clearField(1);
}

