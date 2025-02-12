// source: Qmsg.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

goog.provide('proto.Qmsg.QMsg');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.Qmsg.QMsg = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.Qmsg.QMsg, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.Qmsg.QMsg.displayName = 'proto.Qmsg.QMsg';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.Qmsg.QMsg.prototype.toObject = function(opt_includeInstance) {
  return proto.Qmsg.QMsg.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.Qmsg.QMsg} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Qmsg.QMsg.toObject = function(includeInstance, msg) {
  var f, obj = {
id: jspb.Message.getFieldWithDefault(msg, 1, ""),
frome: jspb.Message.getFieldWithDefault(msg, 2, ""),
too: jspb.Message.getFieldWithDefault(msg, 3, ""),
ackw: jspb.Message.getBooleanFieldWithDefault(msg, 4, false),
priority: jspb.Message.getBooleanFieldWithDefault(msg, 5, false),
message: msg.getMessage_asB64()
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.Qmsg.QMsg}
 */
proto.Qmsg.QMsg.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.Qmsg.QMsg;
  return proto.Qmsg.QMsg.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.Qmsg.QMsg} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.Qmsg.QMsg}
 */
proto.Qmsg.QMsg.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setFrome(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setToo(value);
      break;
    case 4:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setAckw(value);
      break;
    case 5:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setPriority(value);
      break;
    case 6:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setMessage(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.Qmsg.QMsg.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.Qmsg.QMsg.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.Qmsg.QMsg} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Qmsg.QMsg.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getFrome();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getToo();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getAckw();
  if (f) {
    writer.writeBool(
      4,
      f
    );
  }
  f = message.getPriority();
  if (f) {
    writer.writeBool(
      5,
      f
    );
  }
  f = message.getMessage_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      6,
      f
    );
  }
};


/**
 * optional string ID = 1;
 * @return {string}
 */
proto.Qmsg.QMsg.prototype.getId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.Qmsg.QMsg} returns this
 */
proto.Qmsg.QMsg.prototype.setId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string FROME = 2;
 * @return {string}
 */
proto.Qmsg.QMsg.prototype.getFrome = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.Qmsg.QMsg} returns this
 */
proto.Qmsg.QMsg.prototype.setFrome = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string TOO = 3;
 * @return {string}
 */
proto.Qmsg.QMsg.prototype.getToo = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.Qmsg.QMsg} returns this
 */
proto.Qmsg.QMsg.prototype.setToo = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional bool ACKW = 4;
 * @return {boolean}
 */
proto.Qmsg.QMsg.prototype.getAckw = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 4, false));
};


/**
 * @param {boolean} value
 * @return {!proto.Qmsg.QMsg} returns this
 */
proto.Qmsg.QMsg.prototype.setAckw = function(value) {
  return jspb.Message.setProto3BooleanField(this, 4, value);
};


/**
 * optional bool PRIORITY = 5;
 * @return {boolean}
 */
proto.Qmsg.QMsg.prototype.getPriority = function() {
  return /** @type {boolean} */ (jspb.Message.getBooleanFieldWithDefault(this, 5, false));
};


/**
 * @param {boolean} value
 * @return {!proto.Qmsg.QMsg} returns this
 */
proto.Qmsg.QMsg.prototype.setPriority = function(value) {
  return jspb.Message.setProto3BooleanField(this, 5, value);
};


/**
 * optional bytes MESSAGE = 6;
 * @return {string}
 */
proto.Qmsg.QMsg.prototype.getMessage = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 6, ""));
};


/**
 * optional bytes MESSAGE = 6;
 * This is a type-conversion wrapper around `getMessage()`
 * @return {string}
 */
proto.Qmsg.QMsg.prototype.getMessage_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getMessage()));
};


/**
 * optional bytes MESSAGE = 6;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getMessage()`
 * @return {!Uint8Array}
 */
proto.Qmsg.QMsg.prototype.getMessage_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getMessage()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.Qmsg.QMsg} returns this
 */
proto.Qmsg.QMsg.prototype.setMessage = function(value) {
  return jspb.Message.setProto3BytesField(this, 6, value);
};


