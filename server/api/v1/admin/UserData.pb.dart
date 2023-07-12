///
//  Generated code. Do not modify.
//  source: v1/admin/UserData.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;
import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import '../../google/protobuf/empty.pb.dart' as $3;

class UserUpdateDataRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UserUpdateDataRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.admin'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ID', protoName: 'ID')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'FirstName', protoName: 'FirstName')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'LastName', protoName: 'LastName')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'FirstNameKana', protoName: 'FirstNameKana')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'LastNameKana', protoName: 'LastNameKana')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'CompanyName', protoName: 'CompanyName')
    ..aOS(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'BirthDate', protoName: 'BirthDate')
    ..aOS(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ZipCode', protoName: 'ZipCode')
    ..aOS(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Prefecture', protoName: 'Prefecture')
    ..aOS(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'City', protoName: 'City')
    ..aOS(11, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Address', protoName: 'Address')
    ..aOS(12, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Tel', protoName: 'Tel')
    ..aOS(13, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Mail', protoName: 'Mail')
    ..aOB(14, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'AcceptMail', protoName: 'AcceptMail')
    ..aOS(15, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'InnerNote', protoName: 'InnerNote')
    ..aOB(16, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'IsBlackCustomer', protoName: 'IsBlackCustomer')
    ..hasRequiredFields = false
  ;

  UserUpdateDataRequest._() : super();
  factory UserUpdateDataRequest({
    $core.String? iD,
    $core.String? firstName,
    $core.String? lastName,
    $core.String? firstNameKana,
    $core.String? lastNameKana,
    $core.String? companyName,
    $core.String? birthDate,
    $core.String? zipCode,
    $core.String? prefecture,
    $core.String? city,
    $core.String? address,
    $core.String? tel,
    $core.String? mail,
    $core.bool? acceptMail,
    $core.String? innerNote,
    $core.bool? isBlackCustomer,
  }) {
    final _result = create();
    if (iD != null) {
      _result.iD = iD;
    }
    if (firstName != null) {
      _result.firstName = firstName;
    }
    if (lastName != null) {
      _result.lastName = lastName;
    }
    if (firstNameKana != null) {
      _result.firstNameKana = firstNameKana;
    }
    if (lastNameKana != null) {
      _result.lastNameKana = lastNameKana;
    }
    if (companyName != null) {
      _result.companyName = companyName;
    }
    if (birthDate != null) {
      _result.birthDate = birthDate;
    }
    if (zipCode != null) {
      _result.zipCode = zipCode;
    }
    if (prefecture != null) {
      _result.prefecture = prefecture;
    }
    if (city != null) {
      _result.city = city;
    }
    if (address != null) {
      _result.address = address;
    }
    if (tel != null) {
      _result.tel = tel;
    }
    if (mail != null) {
      _result.mail = mail;
    }
    if (acceptMail != null) {
      _result.acceptMail = acceptMail;
    }
    if (innerNote != null) {
      _result.innerNote = innerNote;
    }
    if (isBlackCustomer != null) {
      _result.isBlackCustomer = isBlackCustomer;
    }
    return _result;
  }
  factory UserUpdateDataRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserUpdateDataRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UserUpdateDataRequest clone() => UserUpdateDataRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UserUpdateDataRequest copyWith(void Function(UserUpdateDataRequest) updates) => super.copyWith((message) => updates(message as UserUpdateDataRequest)) as UserUpdateDataRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserUpdateDataRequest create() => UserUpdateDataRequest._();
  UserUpdateDataRequest createEmptyInstance() => create();
  static $pb.PbList<UserUpdateDataRequest> createRepeated() => $pb.PbList<UserUpdateDataRequest>();
  @$core.pragma('dart2js:noInline')
  static UserUpdateDataRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserUpdateDataRequest>(create);
  static UserUpdateDataRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get iD => $_getSZ(0);
  @$pb.TagNumber(1)
  set iD($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasID() => $_has(0);
  @$pb.TagNumber(1)
  void clearID() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get firstName => $_getSZ(1);
  @$pb.TagNumber(2)
  set firstName($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasFirstName() => $_has(1);
  @$pb.TagNumber(2)
  void clearFirstName() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get lastName => $_getSZ(2);
  @$pb.TagNumber(3)
  set lastName($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasLastName() => $_has(2);
  @$pb.TagNumber(3)
  void clearLastName() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get firstNameKana => $_getSZ(3);
  @$pb.TagNumber(4)
  set firstNameKana($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasFirstNameKana() => $_has(3);
  @$pb.TagNumber(4)
  void clearFirstNameKana() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get lastNameKana => $_getSZ(4);
  @$pb.TagNumber(5)
  set lastNameKana($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasLastNameKana() => $_has(4);
  @$pb.TagNumber(5)
  void clearLastNameKana() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get companyName => $_getSZ(5);
  @$pb.TagNumber(6)
  set companyName($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasCompanyName() => $_has(5);
  @$pb.TagNumber(6)
  void clearCompanyName() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get birthDate => $_getSZ(6);
  @$pb.TagNumber(7)
  set birthDate($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasBirthDate() => $_has(6);
  @$pb.TagNumber(7)
  void clearBirthDate() => clearField(7);

  @$pb.TagNumber(8)
  $core.String get zipCode => $_getSZ(7);
  @$pb.TagNumber(8)
  set zipCode($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasZipCode() => $_has(7);
  @$pb.TagNumber(8)
  void clearZipCode() => clearField(8);

  @$pb.TagNumber(9)
  $core.String get prefecture => $_getSZ(8);
  @$pb.TagNumber(9)
  set prefecture($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasPrefecture() => $_has(8);
  @$pb.TagNumber(9)
  void clearPrefecture() => clearField(9);

  @$pb.TagNumber(10)
  $core.String get city => $_getSZ(9);
  @$pb.TagNumber(10)
  set city($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasCity() => $_has(9);
  @$pb.TagNumber(10)
  void clearCity() => clearField(10);

  @$pb.TagNumber(11)
  $core.String get address => $_getSZ(10);
  @$pb.TagNumber(11)
  set address($core.String v) { $_setString(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasAddress() => $_has(10);
  @$pb.TagNumber(11)
  void clearAddress() => clearField(11);

  @$pb.TagNumber(12)
  $core.String get tel => $_getSZ(11);
  @$pb.TagNumber(12)
  set tel($core.String v) { $_setString(11, v); }
  @$pb.TagNumber(12)
  $core.bool hasTel() => $_has(11);
  @$pb.TagNumber(12)
  void clearTel() => clearField(12);

  @$pb.TagNumber(13)
  $core.String get mail => $_getSZ(12);
  @$pb.TagNumber(13)
  set mail($core.String v) { $_setString(12, v); }
  @$pb.TagNumber(13)
  $core.bool hasMail() => $_has(12);
  @$pb.TagNumber(13)
  void clearMail() => clearField(13);

  @$pb.TagNumber(14)
  $core.bool get acceptMail => $_getBF(13);
  @$pb.TagNumber(14)
  set acceptMail($core.bool v) { $_setBool(13, v); }
  @$pb.TagNumber(14)
  $core.bool hasAcceptMail() => $_has(13);
  @$pb.TagNumber(14)
  void clearAcceptMail() => clearField(14);

  @$pb.TagNumber(15)
  $core.String get innerNote => $_getSZ(14);
  @$pb.TagNumber(15)
  set innerNote($core.String v) { $_setString(14, v); }
  @$pb.TagNumber(15)
  $core.bool hasInnerNote() => $_has(14);
  @$pb.TagNumber(15)
  void clearInnerNote() => clearField(15);

  @$pb.TagNumber(16)
  $core.bool get isBlackCustomer => $_getBF(15);
  @$pb.TagNumber(16)
  set isBlackCustomer($core.bool v) { $_setBool(15, v); }
  @$pb.TagNumber(16)
  $core.bool hasIsBlackCustomer() => $_has(15);
  @$pb.TagNumber(16)
  void clearIsBlackCustomer() => clearField(16);
}

class UserDataResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UserDataResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.admin'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ID', protoName: 'ID')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'FirstName', protoName: 'FirstName')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'LastName', protoName: 'LastName')
    ..aOS(4, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'FirstNameKana', protoName: 'FirstNameKana')
    ..aOS(5, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'LastNameKana', protoName: 'LastNameKana')
    ..aOS(6, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'CompanyName', protoName: 'CompanyName')
    ..aOS(7, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'BirthDate', protoName: 'BirthDate')
    ..aOS(8, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ZipCode', protoName: 'ZipCode')
    ..aOS(9, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Prefecture', protoName: 'Prefecture')
    ..aOS(10, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'City', protoName: 'City')
    ..aOS(11, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Address', protoName: 'Address')
    ..aOS(12, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Tel', protoName: 'Tel')
    ..aOS(13, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'Mail', protoName: 'Mail')
    ..aOB(14, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'AcceptMail', protoName: 'AcceptMail')
    ..aOS(15, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'InnerNote', protoName: 'InnerNote')
    ..aOB(16, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'IsBlackCustomer', protoName: 'IsBlackCustomer')
    ..hasRequiredFields = false
  ;

  UserDataResponse._() : super();
  factory UserDataResponse({
    $core.String? iD,
    $core.String? firstName,
    $core.String? lastName,
    $core.String? firstNameKana,
    $core.String? lastNameKana,
    $core.String? companyName,
    $core.String? birthDate,
    $core.String? zipCode,
    $core.String? prefecture,
    $core.String? city,
    $core.String? address,
    $core.String? tel,
    $core.String? mail,
    $core.bool? acceptMail,
    $core.String? innerNote,
    $core.bool? isBlackCustomer,
  }) {
    final _result = create();
    if (iD != null) {
      _result.iD = iD;
    }
    if (firstName != null) {
      _result.firstName = firstName;
    }
    if (lastName != null) {
      _result.lastName = lastName;
    }
    if (firstNameKana != null) {
      _result.firstNameKana = firstNameKana;
    }
    if (lastNameKana != null) {
      _result.lastNameKana = lastNameKana;
    }
    if (companyName != null) {
      _result.companyName = companyName;
    }
    if (birthDate != null) {
      _result.birthDate = birthDate;
    }
    if (zipCode != null) {
      _result.zipCode = zipCode;
    }
    if (prefecture != null) {
      _result.prefecture = prefecture;
    }
    if (city != null) {
      _result.city = city;
    }
    if (address != null) {
      _result.address = address;
    }
    if (tel != null) {
      _result.tel = tel;
    }
    if (mail != null) {
      _result.mail = mail;
    }
    if (acceptMail != null) {
      _result.acceptMail = acceptMail;
    }
    if (innerNote != null) {
      _result.innerNote = innerNote;
    }
    if (isBlackCustomer != null) {
      _result.isBlackCustomer = isBlackCustomer;
    }
    return _result;
  }
  factory UserDataResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserDataResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UserDataResponse clone() => UserDataResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UserDataResponse copyWith(void Function(UserDataResponse) updates) => super.copyWith((message) => updates(message as UserDataResponse)) as UserDataResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserDataResponse create() => UserDataResponse._();
  UserDataResponse createEmptyInstance() => create();
  static $pb.PbList<UserDataResponse> createRepeated() => $pb.PbList<UserDataResponse>();
  @$core.pragma('dart2js:noInline')
  static UserDataResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserDataResponse>(create);
  static UserDataResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get iD => $_getSZ(0);
  @$pb.TagNumber(1)
  set iD($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasID() => $_has(0);
  @$pb.TagNumber(1)
  void clearID() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get firstName => $_getSZ(1);
  @$pb.TagNumber(2)
  set firstName($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasFirstName() => $_has(1);
  @$pb.TagNumber(2)
  void clearFirstName() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get lastName => $_getSZ(2);
  @$pb.TagNumber(3)
  set lastName($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasLastName() => $_has(2);
  @$pb.TagNumber(3)
  void clearLastName() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get firstNameKana => $_getSZ(3);
  @$pb.TagNumber(4)
  set firstNameKana($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasFirstNameKana() => $_has(3);
  @$pb.TagNumber(4)
  void clearFirstNameKana() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get lastNameKana => $_getSZ(4);
  @$pb.TagNumber(5)
  set lastNameKana($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasLastNameKana() => $_has(4);
  @$pb.TagNumber(5)
  void clearLastNameKana() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get companyName => $_getSZ(5);
  @$pb.TagNumber(6)
  set companyName($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasCompanyName() => $_has(5);
  @$pb.TagNumber(6)
  void clearCompanyName() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get birthDate => $_getSZ(6);
  @$pb.TagNumber(7)
  set birthDate($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasBirthDate() => $_has(6);
  @$pb.TagNumber(7)
  void clearBirthDate() => clearField(7);

  @$pb.TagNumber(8)
  $core.String get zipCode => $_getSZ(7);
  @$pb.TagNumber(8)
  set zipCode($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasZipCode() => $_has(7);
  @$pb.TagNumber(8)
  void clearZipCode() => clearField(8);

  @$pb.TagNumber(9)
  $core.String get prefecture => $_getSZ(8);
  @$pb.TagNumber(9)
  set prefecture($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasPrefecture() => $_has(8);
  @$pb.TagNumber(9)
  void clearPrefecture() => clearField(9);

  @$pb.TagNumber(10)
  $core.String get city => $_getSZ(9);
  @$pb.TagNumber(10)
  set city($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasCity() => $_has(9);
  @$pb.TagNumber(10)
  void clearCity() => clearField(10);

  @$pb.TagNumber(11)
  $core.String get address => $_getSZ(10);
  @$pb.TagNumber(11)
  set address($core.String v) { $_setString(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasAddress() => $_has(10);
  @$pb.TagNumber(11)
  void clearAddress() => clearField(11);

  @$pb.TagNumber(12)
  $core.String get tel => $_getSZ(11);
  @$pb.TagNumber(12)
  set tel($core.String v) { $_setString(11, v); }
  @$pb.TagNumber(12)
  $core.bool hasTel() => $_has(11);
  @$pb.TagNumber(12)
  void clearTel() => clearField(12);

  @$pb.TagNumber(13)
  $core.String get mail => $_getSZ(12);
  @$pb.TagNumber(13)
  set mail($core.String v) { $_setString(12, v); }
  @$pb.TagNumber(13)
  $core.bool hasMail() => $_has(12);
  @$pb.TagNumber(13)
  void clearMail() => clearField(13);

  @$pb.TagNumber(14)
  $core.bool get acceptMail => $_getBF(13);
  @$pb.TagNumber(14)
  set acceptMail($core.bool v) { $_setBool(13, v); }
  @$pb.TagNumber(14)
  $core.bool hasAcceptMail() => $_has(13);
  @$pb.TagNumber(14)
  void clearAcceptMail() => clearField(14);

  @$pb.TagNumber(15)
  $core.String get innerNote => $_getSZ(14);
  @$pb.TagNumber(15)
  set innerNote($core.String v) { $_setString(14, v); }
  @$pb.TagNumber(15)
  $core.bool hasInnerNote() => $_has(14);
  @$pb.TagNumber(15)
  void clearInnerNote() => clearField(15);

  @$pb.TagNumber(16)
  $core.bool get isBlackCustomer => $_getBF(15);
  @$pb.TagNumber(16)
  set isBlackCustomer($core.bool v) { $_setBool(15, v); }
  @$pb.TagNumber(16)
  $core.bool hasIsBlackCustomer() => $_has(15);
  @$pb.TagNumber(16)
  void clearIsBlackCustomer() => clearField(16);
}

class UserDeleteRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'UserDeleteRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'server.admin'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'ID', protoName: 'ID')
    ..hasRequiredFields = false
  ;

  UserDeleteRequest._() : super();
  factory UserDeleteRequest({
    $core.String? iD,
  }) {
    final _result = create();
    if (iD != null) {
      _result.iD = iD;
    }
    return _result;
  }
  factory UserDeleteRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory UserDeleteRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  UserDeleteRequest clone() => UserDeleteRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  UserDeleteRequest copyWith(void Function(UserDeleteRequest) updates) => super.copyWith((message) => updates(message as UserDeleteRequest)) as UserDeleteRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static UserDeleteRequest create() => UserDeleteRequest._();
  UserDeleteRequest createEmptyInstance() => create();
  static $pb.PbList<UserDeleteRequest> createRepeated() => $pb.PbList<UserDeleteRequest>();
  @$core.pragma('dart2js:noInline')
  static UserDeleteRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<UserDeleteRequest>(create);
  static UserDeleteRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get iD => $_getSZ(0);
  @$pb.TagNumber(1)
  set iD($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasID() => $_has(0);
  @$pb.TagNumber(1)
  void clearID() => clearField(1);
}

class UserDataControllerApi {
  $pb.RpcClient _client;
  UserDataControllerApi(this._client);

  $async.Future<UserDataResponse> update($pb.ClientContext? ctx, UserUpdateDataRequest request) {
    var emptyResponse = UserDataResponse();
    return _client.invoke<UserDataResponse>(ctx, 'UserDataController', 'Update', request, emptyResponse);
  }
  $async.Future<$3.Empty> delete($pb.ClientContext? ctx, UserDeleteRequest request) {
    var emptyResponse = $3.Empty();
    return _client.invoke<$3.Empty>(ctx, 'UserDataController', 'Delete', request, emptyResponse);
  }
}
