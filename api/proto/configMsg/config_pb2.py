# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: config.proto
# Protobuf Python Version: 5.28.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    28,
    2,
    '',
    'config.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0c\x63onfig.proto\x12\x06\x43onfig\"q\n\nKeysValues\x12\x32\n\x08KeyValue\x18\x01 \x03(\x0b\x32 .Config.KeysValues.KeyValueEntry\x1a/\n\rKeyValueEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\xfb\x05\n\tConfigMsg\x12\x34\n\tReqClient\x18\x01 \x01(\x0e\x32!.Config.ConfigMsg.ConfigClientCmd\x12\x45\n\x12SectionsKeysValues\x18\x02 \x03(\x0b\x32).Config.ConfigMsg.SectionsKeysValuesEntry\x12\x35\n\nRespServer\x18\x03 \x01(\x0e\x32!.Config.ConfigMsg.ConfigServerMsg\x1aM\n\x17SectionsKeysValuesEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12!\n\x05value\x18\x02 \x01(\x0b\x32\x12.Config.KeysValues:\x02\x38\x01\"\xce\x01\n\x0f\x43onfigClientCmd\x12\x15\n\x11update_mem_config\x10\x00\x12\x18\n\x14update_config_object\x10\x01\x12\x12\n\x0eget_mem_config\x10\x02\x12\x15\n\x11get_config_object\x10\x03\x12\x17\n\x13\x61\x64\x64_config_listener\x10\x04\x12\x13\n\x0f\x64ump_mem_config\x10\x05\x12\x16\n\x12get_notif_loglevel\x10\x06\x12\x19\n\x15update_notif_loglevel\x10\x07\"\x99\x02\n\x0f\x43onfigServerMsg\x12\x18\n\x14propagate_mem_config\x10\x00\x12\x1a\n\x16mem_config_update_done\x10\x01\x12\x14\n\x10propagate_config\x10\x02\x12\x16\n\x12\x63onfig_update_done\x10\x03\x12\x1c\n\x18mem_config_update_failed\x10\x04\x12\x18\n\x14\x63onfig_update_failed\x10\x05\x12\x1c\n\x18propagate_notif_loglevel\x10\x06\x12\x14\n\x10send_config_init\x10\x07\x12\x18\n\x14send_mem_config_init\x10\x08\x12\x1c\n\x18send_notif_loglevel_init\x10\tB\x03Z\x01/b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'config_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\001/'
  _globals['_KEYSVALUES_KEYVALUEENTRY']._loaded_options = None
  _globals['_KEYSVALUES_KEYVALUEENTRY']._serialized_options = b'8\001'
  _globals['_CONFIGMSG_SECTIONSKEYSVALUESENTRY']._loaded_options = None
  _globals['_CONFIGMSG_SECTIONSKEYSVALUESENTRY']._serialized_options = b'8\001'
  _globals['_KEYSVALUES']._serialized_start=24
  _globals['_KEYSVALUES']._serialized_end=137
  _globals['_KEYSVALUES_KEYVALUEENTRY']._serialized_start=90
  _globals['_KEYSVALUES_KEYVALUEENTRY']._serialized_end=137
  _globals['_CONFIGMSG']._serialized_start=140
  _globals['_CONFIGMSG']._serialized_end=903
  _globals['_CONFIGMSG_SECTIONSKEYSVALUESENTRY']._serialized_start=333
  _globals['_CONFIGMSG_SECTIONSKEYSVALUESENTRY']._serialized_end=410
  _globals['_CONFIGMSG_CONFIGCLIENTCMD']._serialized_start=413
  _globals['_CONFIGMSG_CONFIGCLIENTCMD']._serialized_end=619
  _globals['_CONFIGMSG_CONFIGSERVERMSG']._serialized_start=622
  _globals['_CONFIGMSG_CONFIGSERVERMSG']._serialized_end=903
# @@protoc_insertion_point(module_scope)
