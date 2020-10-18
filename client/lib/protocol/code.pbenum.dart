///
//  Generated code. Do not modify.
//  source: code.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' as $core;
import 'package:protobuf/protobuf.dart' as $pb;

class Code extends $pb.ProtobufEnum {
  static const Code SUCCESS = Code._(0, 'SUCCESS');
  static const Code INTERNAL_ERROR = Code._(1, 'INTERNAL_ERROR');
  static const Code METHOD_NOT_ALLOWED = Code._(2, 'METHOD_NOT_ALLOWED');
  static const Code AUTH_FAILED = Code._(3, 'AUTH_FAILED');
  static const Code PROTOCOL_MISMATCH = Code._(4, 'PROTOCOL_MISMATCH');
  static const Code INVALID_DATA = Code._(5, 'INVALID_DATA');
  static const Code USER_EXISTS = Code._(6, 'USER_EXISTS');
  static const Code USER_NOT_EXISTS = Code._(7, 'USER_NOT_EXISTS');
  static const Code RESTAURANT_NOT_EXISTS = Code._(8, 'RESTAURANT_NOT_EXISTS');

  static const $core.List<Code> values = <Code> [
    SUCCESS,
    INTERNAL_ERROR,
    METHOD_NOT_ALLOWED,
    AUTH_FAILED,
    PROTOCOL_MISMATCH,
    INVALID_DATA,
    USER_EXISTS,
    USER_NOT_EXISTS,
    RESTAURANT_NOT_EXISTS,
  ];

  static final $core.Map<$core.int, Code> _byValue = $pb.ProtobufEnum.initByValue(values);
  static Code valueOf($core.int value) => _byValue[value];

  const Code._($core.int v, $core.String n) : super(v, n);
}

