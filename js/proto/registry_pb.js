// source: registry.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var gogoproto_gogo_pb = require('./gogoproto/gogo_pb.js');
goog.object.extend(proto, gogoproto_gogo_pb);
goog.exportSymbol('proto.registry.NodeIdentity', null, global);
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
proto.registry.NodeIdentity = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.registry.NodeIdentity, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.registry.NodeIdentity.displayName = 'proto.registry.NodeIdentity';
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
proto.registry.NodeIdentity.prototype.toObject = function(opt_includeInstance) {
  return proto.registry.NodeIdentity.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.registry.NodeIdentity} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.registry.NodeIdentity.toObject = function(includeInstance, msg) {
  var f, obj = {
    moniker: jspb.Message.getFieldWithDefault(msg, 1, ""),
    networkaddress: jspb.Message.getFieldWithDefault(msg, 2, ""),
    tendermintnodeid: msg.getTendermintnodeid_asB64(),
    validatorpublickey: msg.getValidatorpublickey_asB64()
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
 * @return {!proto.registry.NodeIdentity}
 */
proto.registry.NodeIdentity.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.registry.NodeIdentity;
  return proto.registry.NodeIdentity.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.registry.NodeIdentity} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.registry.NodeIdentity}
 */
proto.registry.NodeIdentity.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setMoniker(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setNetworkaddress(value);
      break;
    case 3:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setTendermintnodeid(value);
      break;
    case 4:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setValidatorpublickey(value);
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
proto.registry.NodeIdentity.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.registry.NodeIdentity.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.registry.NodeIdentity} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.registry.NodeIdentity.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMoniker();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getNetworkaddress();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getTendermintnodeid_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      3,
      f
    );
  }
  f = message.getValidatorpublickey_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      4,
      f
    );
  }
};


/**
 * optional string Moniker =  1;
 * @return {string}
 */
proto.registry.NodeIdentity.prototype.getMoniker =  function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.registry.NodeIdentity} returns this
 */
proto.registry.NodeIdentity.prototype.setMoniker =  function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string NetworkAddress = 2;
 * @return {string}
 */
proto.registry.NodeIdentity.prototype.getNetworkaddress = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.registry.NodeIdentity} returns this
 */
proto.registry.NodeIdentity.prototype.setNetworkaddress = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional bytes TendermintNodeID = 3;
 * @return {!(string|Uint8Array)}
 */
proto.registry.NodeIdentity.prototype.getTendermintnodeid = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * optional bytes TendermintNodeID = 3;
 * This is a type-conversion wrapper around `getTendermintnodeid()`
 * @return {string}
 */
proto.registry.NodeIdentity.prototype.getTendermintnodeid_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getTendermintnodeid()));
};


/**
 * optional bytes TendermintNodeID = 3;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getTendermintnodeid()`
 * @return {!Uint8Array}
 */
proto.registry.NodeIdentity.prototype.getTendermintnodeid_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getTendermintnodeid()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.registry.NodeIdentity} returns this
 */
proto.registry.NodeIdentity.prototype.setTendermintnodeid = function(value) {
  return jspb.Message.setProto3BytesField(this, 3, value);
};


/**
 * optional bytes ValidatorPublicKey = 4;
 * @return {!(string|Uint8Array)}
 */
proto.registry.NodeIdentity.prototype.getValidatorpublickey = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * optional bytes ValidatorPublicKey = 4;
 * This is a type-conversion wrapper around `getValidatorpublickey()`
 * @return {string}
 */
proto.registry.NodeIdentity.prototype.getValidatorpublickey_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getValidatorpublickey()));
};


/**
 * optional bytes ValidatorPublicKey = 4;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getValidatorpublickey()`
 * @return {!Uint8Array}
 */
proto.registry.NodeIdentity.prototype.getValidatorpublickey_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getValidatorpublickey()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.registry.NodeIdentity} returns this
 */
proto.registry.NodeIdentity.prototype.setValidatorpublickey = function(value) {
  return jspb.Message.setProto3BytesField(this, 4, value);
};


goog.object.extend(exports, proto.registry);
