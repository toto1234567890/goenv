# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: Qmsg.proto
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
    'Qmsg.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import any_pb2 as google_dot_protobuf_dot_any__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nQmsg.proto\x12\x04Qmsg\x1a\x19google/protobuf/any.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"@\n\x08HelloMsg\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x13\n\x0blocalServer\x18\x02 \x01(\t\x12\x11\n\tlocalPort\x18\x03 \x01(\t\"\xbe\x01\n\tStreamMsg\x12\x0e\n\x06\x42ROKER\x18\x01 \x01(\t\x12\x33\n\rSTREAM_ACTION\x18\x02 \x01(\x0e\x32\x1c.Qmsg.StreamMsg.StreamAction\x12\x0e\n\x06TICKER\x18\x03 \x01(\t\"\\\n\x0cStreamAction\x12\t\n\x05QUOTE\x10\x00\x12\x0b\n\x07UNQUOTE\x10\x01\x12\x08\n\x04TICK\x10\x02\x12\n\n\x06UNTICK\x10\x03\x12\r\n\tORDERBOOK\x10\x04\x12\x0f\n\x0bUNORDERBOOK\x10\x05\"\xa6\x02\n\x08TradeMsg\x12\x12\n\nACCOUNT_ID\x18\x01 \x01(\x05\x12/\n\x0bTRADEACTION\x18\x02 \x01(\x0e\x32\x1a.Qmsg.TradeMsg.TradeAction\x12\x0e\n\x06TICKER\x18\x03 \x01(\t\x12\x10\n\x08QUANTITY\x18\x04 \x01(\t\x12\x34\n\x0bTradeParams\x18\x05 \x03(\x0b\x32\x1f.Qmsg.TradeMsg.TradeParamsEntry\x1aH\n\x10TradeParamsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12#\n\x05value\x18\x02 \x01(\x0b\x32\x14.google.protobuf.Any:\x02\x38\x01\"3\n\x0bTradeAction\x12\x07\n\x03\x42UY\x10\x00\x12\x08\n\x04SELL\x10\x01\x12\x08\n\x04\x43\x41LL\x10\x02\x12\x07\n\x03PUT\x10\x03\"_\n\x04QMsg\x12\n\n\x02ID\x18\x01 \x01(\t\x12\r\n\x05\x46ROME\x18\x02 \x01(\t\x12\x0b\n\x03TOO\x18\x03 \x01(\t\x12\x0c\n\x04\x41\x43KW\x18\x04 \x01(\x08\x12\x10\n\x08PRIORITY\x18\x05 \x01(\x08\x12\x0f\n\x07MESSAGE\x18\x06 \x01(\x0c\x42\x03Z\x01/b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'Qmsg_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\001/'
  _globals['_TRADEMSG_TRADEPARAMSENTRY']._loaded_options = None
  _globals['_TRADEMSG_TRADEPARAMSENTRY']._serialized_options = b'8\001'
  _globals['_HELLOMSG']._serialized_start=80
  _globals['_HELLOMSG']._serialized_end=144
  _globals['_STREAMMSG']._serialized_start=147
  _globals['_STREAMMSG']._serialized_end=337
  _globals['_STREAMMSG_STREAMACTION']._serialized_start=245
  _globals['_STREAMMSG_STREAMACTION']._serialized_end=337
  _globals['_TRADEMSG']._serialized_start=340
  _globals['_TRADEMSG']._serialized_end=634
  _globals['_TRADEMSG_TRADEPARAMSENTRY']._serialized_start=509
  _globals['_TRADEMSG_TRADEPARAMSENTRY']._serialized_end=581
  _globals['_TRADEMSG_TRADEACTION']._serialized_start=583
  _globals['_TRADEMSG_TRADEACTION']._serialized_end=634
  _globals['_QMSG']._serialized_start=636
  _globals['_QMSG']._serialized_end=731
# @@protoc_insertion_point(module_scope)
