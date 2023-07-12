///
//  Generated code. Do not modify.
//  source: v1/user/UserData.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,deprecated_member_use_from_same_package,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'package:protobuf/protobuf.dart' as $pb;

import 'dart:core' as $core;
import 'UserData.pb.dart' as $6;
import 'UserData.pbjson.dart';

export 'UserData.pb.dart';

abstract class UserDataControllerServiceBase extends $pb.GeneratedService {
  $async.Future<$6.UserDataResponse> register($pb.ServerContext ctx, $6.UserRegisterRequest request);
  $async.Future<$6.UserDataResponse> update($pb.ServerContext ctx, $6.UserUpdateDataRequest request);

  $pb.GeneratedMessage createRequest($core.String method) {
    switch (method) {
      case 'Register': return $6.UserRegisterRequest();
      case 'Update': return $6.UserUpdateDataRequest();
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $async.Future<$pb.GeneratedMessage> handleCall($pb.ServerContext ctx, $core.String method, $pb.GeneratedMessage request) {
    switch (method) {
      case 'Register': return this.register(ctx, request as $6.UserRegisterRequest);
      case 'Update': return this.update(ctx, request as $6.UserUpdateDataRequest);
      default: throw $core.ArgumentError('Unknown method: $method');
    }
  }

  $core.Map<$core.String, $core.dynamic> get $json => UserDataControllerServiceBase$json;
  $core.Map<$core.String, $core.Map<$core.String, $core.dynamic>> get $messageJson => UserDataControllerServiceBase$messageJson;
}
