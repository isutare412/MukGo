///
//  Generated code. Do not modify.
//  source: proto/request.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'model.pb.dart' as $0;

class ReviewPost extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ReviewPost', package: const $pb.PackageName('proto'), createEmptyInstance: create)
    ..aOS(1, 'restaurantId')
    ..aOM<$0.Review>(2, 'review', subBuilder: $0.Review.create)
    ..hasRequiredFields = false
  ;

  ReviewPost._() : super();
  factory ReviewPost() => create();
  factory ReviewPost.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ReviewPost.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ReviewPost clone() => ReviewPost()..mergeFromMessage(this);
  ReviewPost copyWith(void Function(ReviewPost) updates) => super.copyWith((message) => updates(message as ReviewPost));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ReviewPost create() => ReviewPost._();
  ReviewPost createEmptyInstance() => create();
  static $pb.PbList<ReviewPost> createRepeated() => $pb.PbList<ReviewPost>();
  @$core.pragma('dart2js:noInline')
  static ReviewPost getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ReviewPost>(create);
  static ReviewPost _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get restaurantId => $_getSZ(0);
  @$pb.TagNumber(1)
  set restaurantId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasRestaurantId() => $_has(0);
  @$pb.TagNumber(1)
  void clearRestaurantId() => clearField(1);

  @$pb.TagNumber(2)
  $0.Review get review => $_getN(1);
  @$pb.TagNumber(2)
  set review($0.Review v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasReview() => $_has(1);
  @$pb.TagNumber(2)
  void clearReview() => clearField(2);
  @$pb.TagNumber(2)
  $0.Review ensureReview() => $_ensure(1);
}

class RestaurantPost extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RestaurantPost', package: const $pb.PackageName('proto'), createEmptyInstance: create)
    ..aOM<$0.Restaurant>(1, 'restaurant', subBuilder: $0.Restaurant.create)
    ..hasRequiredFields = false
  ;

  RestaurantPost._() : super();
  factory RestaurantPost() => create();
  factory RestaurantPost.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RestaurantPost.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RestaurantPost clone() => RestaurantPost()..mergeFromMessage(this);
  RestaurantPost copyWith(void Function(RestaurantPost) updates) => super.copyWith((message) => updates(message as RestaurantPost));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RestaurantPost create() => RestaurantPost._();
  RestaurantPost createEmptyInstance() => create();
  static $pb.PbList<RestaurantPost> createRepeated() => $pb.PbList<RestaurantPost>();
  @$core.pragma('dart2js:noInline')
  static RestaurantPost getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RestaurantPost>(create);
  static RestaurantPost _defaultInstance;

  @$pb.TagNumber(1)
  $0.Restaurant get restaurant => $_getN(0);
  @$pb.TagNumber(1)
  set restaurant($0.Restaurant v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasRestaurant() => $_has(0);
  @$pb.TagNumber(1)
  void clearRestaurant() => clearField(1);
  @$pb.TagNumber(1)
  $0.Restaurant ensureRestaurant() => $_ensure(0);
}

class RestaurantsPost extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RestaurantsPost', package: const $pb.PackageName('proto'), createEmptyInstance: create)
    ..pc<$0.Restaurant>(1, 'restaurants', $pb.PbFieldType.PM, subBuilder: $0.Restaurant.create)
    ..hasRequiredFields = false
  ;

  RestaurantsPost._() : super();
  factory RestaurantsPost() => create();
  factory RestaurantsPost.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RestaurantsPost.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RestaurantsPost clone() => RestaurantsPost()..mergeFromMessage(this);
  RestaurantsPost copyWith(void Function(RestaurantsPost) updates) => super.copyWith((message) => updates(message as RestaurantsPost));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RestaurantsPost create() => RestaurantsPost._();
  RestaurantsPost createEmptyInstance() => create();
  static $pb.PbList<RestaurantsPost> createRepeated() => $pb.PbList<RestaurantsPost>();
  @$core.pragma('dart2js:noInline')
  static RestaurantsPost getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RestaurantsPost>(create);
  static RestaurantsPost _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$0.Restaurant> get restaurants => $_getList(0);
}

class LikePost extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('LikePost', package: const $pb.PackageName('proto'), createEmptyInstance: create)
    ..aOS(1, 'reviewId')
    ..hasRequiredFields = false
  ;

  LikePost._() : super();
  factory LikePost() => create();
  factory LikePost.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory LikePost.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  LikePost clone() => LikePost()..mergeFromMessage(this);
  LikePost copyWith(void Function(LikePost) updates) => super.copyWith((message) => updates(message as LikePost));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static LikePost create() => LikePost._();
  LikePost createEmptyInstance() => create();
  static $pb.PbList<LikePost> createRepeated() => $pb.PbList<LikePost>();
  @$core.pragma('dart2js:noInline')
  static LikePost getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<LikePost>(create);
  static LikePost _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get reviewId => $_getSZ(0);
  @$pb.TagNumber(1)
  set reviewId($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasReviewId() => $_has(0);
  @$pb.TagNumber(1)
  void clearReviewId() => clearField(1);
}

