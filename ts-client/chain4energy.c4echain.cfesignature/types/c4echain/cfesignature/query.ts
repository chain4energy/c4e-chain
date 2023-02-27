/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Params } from "./params";

export const protobufPackage = "chain4energy.c4echain.cfesignature";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryCreateReferenceIdRequest {
  creator: string;
}

export interface QueryCreateReferenceIdResponse {
  referenceId: string;
}

export interface QueryCreateStorageKeyRequest {
  targetAccAddress: string;
  referenceId: string;
}

export interface QueryCreateStorageKeyResponse {
  storageKey: string;
}

export interface QueryCreateReferencePayloadLinkRequest {
  referenceId: string;
  payloadHash: string;
}

export interface QueryCreateReferencePayloadLinkResponse {
  referenceKey: string;
  referenceValue: string;
}

export interface QueryVerifySignatureRequest {
  referenceId: string;
  targetAccAddress: string;
}

export interface QueryVerifySignatureResponse {
  signature: string;
  algorithm: string;
  certificate: string;
  timestamp: string;
  valid: string;
}

export interface QueryGetAccountInfoRequest {
  accAddressString: string;
}

export interface QueryGetAccountInfoResponse {
  accAddress: string;
  pubKey: string;
}

export interface QueryVerifyReferencePayloadLinkRequest {
  referenceId: string;
  payloadHash: string;
}

export interface QueryVerifyReferencePayloadLinkResponse {
  isValid: boolean;
}

export interface QueryGetReferencePayloadLinkRequest {
  referenceId: string;
}

export interface QueryGetReferencePayloadLinkResponse {
  referencePayloadLinkValue: string;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryCreateReferenceIdRequest(): QueryCreateReferenceIdRequest {
  return { creator: "" };
}

export const QueryCreateReferenceIdRequest = {
  encode(message: QueryCreateReferenceIdRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCreateReferenceIdRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCreateReferenceIdRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCreateReferenceIdRequest {
    return { creator: isSet(object.creator) ? String(object.creator) : "" };
  },

  toJSON(message: QueryCreateReferenceIdRequest): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCreateReferenceIdRequest>, I>>(
    object: I,
  ): QueryCreateReferenceIdRequest {
    const message = createBaseQueryCreateReferenceIdRequest();
    message.creator = object.creator ?? "";
    return message;
  },
};

function createBaseQueryCreateReferenceIdResponse(): QueryCreateReferenceIdResponse {
  return { referenceId: "" };
}

export const QueryCreateReferenceIdResponse = {
  encode(message: QueryCreateReferenceIdResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.referenceId !== "") {
      writer.uint32(10).string(message.referenceId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCreateReferenceIdResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCreateReferenceIdResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.referenceId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCreateReferenceIdResponse {
    return { referenceId: isSet(object.referenceId) ? String(object.referenceId) : "" };
  },

  toJSON(message: QueryCreateReferenceIdResponse): unknown {
    const obj: any = {};
    message.referenceId !== undefined && (obj.referenceId = message.referenceId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCreateReferenceIdResponse>, I>>(
    object: I,
  ): QueryCreateReferenceIdResponse {
    const message = createBaseQueryCreateReferenceIdResponse();
    message.referenceId = object.referenceId ?? "";
    return message;
  },
};

function createBaseQueryCreateStorageKeyRequest(): QueryCreateStorageKeyRequest {
  return { targetAccAddress: "", referenceId: "" };
}

export const QueryCreateStorageKeyRequest = {
  encode(message: QueryCreateStorageKeyRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.targetAccAddress !== "") {
      writer.uint32(10).string(message.targetAccAddress);
    }
    if (message.referenceId !== "") {
      writer.uint32(18).string(message.referenceId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCreateStorageKeyRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCreateStorageKeyRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.targetAccAddress = reader.string();
          break;
        case 2:
          message.referenceId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCreateStorageKeyRequest {
    return {
      targetAccAddress: isSet(object.targetAccAddress) ? String(object.targetAccAddress) : "",
      referenceId: isSet(object.referenceId) ? String(object.referenceId) : "",
    };
  },

  toJSON(message: QueryCreateStorageKeyRequest): unknown {
    const obj: any = {};
    message.targetAccAddress !== undefined && (obj.targetAccAddress = message.targetAccAddress);
    message.referenceId !== undefined && (obj.referenceId = message.referenceId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCreateStorageKeyRequest>, I>>(object: I): QueryCreateStorageKeyRequest {
    const message = createBaseQueryCreateStorageKeyRequest();
    message.targetAccAddress = object.targetAccAddress ?? "";
    message.referenceId = object.referenceId ?? "";
    return message;
  },
};

function createBaseQueryCreateStorageKeyResponse(): QueryCreateStorageKeyResponse {
  return { storageKey: "" };
}

export const QueryCreateStorageKeyResponse = {
  encode(message: QueryCreateStorageKeyResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.storageKey !== "") {
      writer.uint32(10).string(message.storageKey);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCreateStorageKeyResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCreateStorageKeyResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.storageKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCreateStorageKeyResponse {
    return { storageKey: isSet(object.storageKey) ? String(object.storageKey) : "" };
  },

  toJSON(message: QueryCreateStorageKeyResponse): unknown {
    const obj: any = {};
    message.storageKey !== undefined && (obj.storageKey = message.storageKey);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCreateStorageKeyResponse>, I>>(
    object: I,
  ): QueryCreateStorageKeyResponse {
    const message = createBaseQueryCreateStorageKeyResponse();
    message.storageKey = object.storageKey ?? "";
    return message;
  },
};

function createBaseQueryCreateReferencePayloadLinkRequest(): QueryCreateReferencePayloadLinkRequest {
  return { referenceId: "", payloadHash: "" };
}

export const QueryCreateReferencePayloadLinkRequest = {
  encode(message: QueryCreateReferencePayloadLinkRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.referenceId !== "") {
      writer.uint32(10).string(message.referenceId);
    }
    if (message.payloadHash !== "") {
      writer.uint32(18).string(message.payloadHash);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCreateReferencePayloadLinkRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCreateReferencePayloadLinkRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.referenceId = reader.string();
          break;
        case 2:
          message.payloadHash = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCreateReferencePayloadLinkRequest {
    return {
      referenceId: isSet(object.referenceId) ? String(object.referenceId) : "",
      payloadHash: isSet(object.payloadHash) ? String(object.payloadHash) : "",
    };
  },

  toJSON(message: QueryCreateReferencePayloadLinkRequest): unknown {
    const obj: any = {};
    message.referenceId !== undefined && (obj.referenceId = message.referenceId);
    message.payloadHash !== undefined && (obj.payloadHash = message.payloadHash);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCreateReferencePayloadLinkRequest>, I>>(
    object: I,
  ): QueryCreateReferencePayloadLinkRequest {
    const message = createBaseQueryCreateReferencePayloadLinkRequest();
    message.referenceId = object.referenceId ?? "";
    message.payloadHash = object.payloadHash ?? "";
    return message;
  },
};

function createBaseQueryCreateReferencePayloadLinkResponse(): QueryCreateReferencePayloadLinkResponse {
  return { referenceKey: "", referenceValue: "" };
}

export const QueryCreateReferencePayloadLinkResponse = {
  encode(message: QueryCreateReferencePayloadLinkResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.referenceKey !== "") {
      writer.uint32(10).string(message.referenceKey);
    }
    if (message.referenceValue !== "") {
      writer.uint32(18).string(message.referenceValue);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCreateReferencePayloadLinkResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCreateReferencePayloadLinkResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.referenceKey = reader.string();
          break;
        case 2:
          message.referenceValue = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCreateReferencePayloadLinkResponse {
    return {
      referenceKey: isSet(object.referenceKey) ? String(object.referenceKey) : "",
      referenceValue: isSet(object.referenceValue) ? String(object.referenceValue) : "",
    };
  },

  toJSON(message: QueryCreateReferencePayloadLinkResponse): unknown {
    const obj: any = {};
    message.referenceKey !== undefined && (obj.referenceKey = message.referenceKey);
    message.referenceValue !== undefined && (obj.referenceValue = message.referenceValue);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCreateReferencePayloadLinkResponse>, I>>(
    object: I,
  ): QueryCreateReferencePayloadLinkResponse {
    const message = createBaseQueryCreateReferencePayloadLinkResponse();
    message.referenceKey = object.referenceKey ?? "";
    message.referenceValue = object.referenceValue ?? "";
    return message;
  },
};

function createBaseQueryVerifySignatureRequest(): QueryVerifySignatureRequest {
  return { referenceId: "", targetAccAddress: "" };
}

export const QueryVerifySignatureRequest = {
  encode(message: QueryVerifySignatureRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.referenceId !== "") {
      writer.uint32(10).string(message.referenceId);
    }
    if (message.targetAccAddress !== "") {
      writer.uint32(18).string(message.targetAccAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryVerifySignatureRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryVerifySignatureRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.referenceId = reader.string();
          break;
        case 2:
          message.targetAccAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryVerifySignatureRequest {
    return {
      referenceId: isSet(object.referenceId) ? String(object.referenceId) : "",
      targetAccAddress: isSet(object.targetAccAddress) ? String(object.targetAccAddress) : "",
    };
  },

  toJSON(message: QueryVerifySignatureRequest): unknown {
    const obj: any = {};
    message.referenceId !== undefined && (obj.referenceId = message.referenceId);
    message.targetAccAddress !== undefined && (obj.targetAccAddress = message.targetAccAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryVerifySignatureRequest>, I>>(object: I): QueryVerifySignatureRequest {
    const message = createBaseQueryVerifySignatureRequest();
    message.referenceId = object.referenceId ?? "";
    message.targetAccAddress = object.targetAccAddress ?? "";
    return message;
  },
};

function createBaseQueryVerifySignatureResponse(): QueryVerifySignatureResponse {
  return { signature: "", algorithm: "", certificate: "", timestamp: "", valid: "" };
}

export const QueryVerifySignatureResponse = {
  encode(message: QueryVerifySignatureResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.signature !== "") {
      writer.uint32(10).string(message.signature);
    }
    if (message.algorithm !== "") {
      writer.uint32(18).string(message.algorithm);
    }
    if (message.certificate !== "") {
      writer.uint32(26).string(message.certificate);
    }
    if (message.timestamp !== "") {
      writer.uint32(34).string(message.timestamp);
    }
    if (message.valid !== "") {
      writer.uint32(42).string(message.valid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryVerifySignatureResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryVerifySignatureResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.signature = reader.string();
          break;
        case 2:
          message.algorithm = reader.string();
          break;
        case 3:
          message.certificate = reader.string();
          break;
        case 4:
          message.timestamp = reader.string();
          break;
        case 5:
          message.valid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryVerifySignatureResponse {
    return {
      signature: isSet(object.signature) ? String(object.signature) : "",
      algorithm: isSet(object.algorithm) ? String(object.algorithm) : "",
      certificate: isSet(object.certificate) ? String(object.certificate) : "",
      timestamp: isSet(object.timestamp) ? String(object.timestamp) : "",
      valid: isSet(object.valid) ? String(object.valid) : "",
    };
  },

  toJSON(message: QueryVerifySignatureResponse): unknown {
    const obj: any = {};
    message.signature !== undefined && (obj.signature = message.signature);
    message.algorithm !== undefined && (obj.algorithm = message.algorithm);
    message.certificate !== undefined && (obj.certificate = message.certificate);
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    message.valid !== undefined && (obj.valid = message.valid);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryVerifySignatureResponse>, I>>(object: I): QueryVerifySignatureResponse {
    const message = createBaseQueryVerifySignatureResponse();
    message.signature = object.signature ?? "";
    message.algorithm = object.algorithm ?? "";
    message.certificate = object.certificate ?? "";
    message.timestamp = object.timestamp ?? "";
    message.valid = object.valid ?? "";
    return message;
  },
};

function createBaseQueryGetAccountInfoRequest(): QueryGetAccountInfoRequest {
  return { accAddressString: "" };
}

export const QueryGetAccountInfoRequest = {
  encode(message: QueryGetAccountInfoRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.accAddressString !== "") {
      writer.uint32(10).string(message.accAddressString);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetAccountInfoRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetAccountInfoRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.accAddressString = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAccountInfoRequest {
    return { accAddressString: isSet(object.accAddressString) ? String(object.accAddressString) : "" };
  },

  toJSON(message: QueryGetAccountInfoRequest): unknown {
    const obj: any = {};
    message.accAddressString !== undefined && (obj.accAddressString = message.accAddressString);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetAccountInfoRequest>, I>>(object: I): QueryGetAccountInfoRequest {
    const message = createBaseQueryGetAccountInfoRequest();
    message.accAddressString = object.accAddressString ?? "";
    return message;
  },
};

function createBaseQueryGetAccountInfoResponse(): QueryGetAccountInfoResponse {
  return { accAddress: "", pubKey: "" };
}

export const QueryGetAccountInfoResponse = {
  encode(message: QueryGetAccountInfoResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.accAddress !== "") {
      writer.uint32(10).string(message.accAddress);
    }
    if (message.pubKey !== "") {
      writer.uint32(18).string(message.pubKey);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetAccountInfoResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetAccountInfoResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.accAddress = reader.string();
          break;
        case 2:
          message.pubKey = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetAccountInfoResponse {
    return {
      accAddress: isSet(object.accAddress) ? String(object.accAddress) : "",
      pubKey: isSet(object.pubKey) ? String(object.pubKey) : "",
    };
  },

  toJSON(message: QueryGetAccountInfoResponse): unknown {
    const obj: any = {};
    message.accAddress !== undefined && (obj.accAddress = message.accAddress);
    message.pubKey !== undefined && (obj.pubKey = message.pubKey);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetAccountInfoResponse>, I>>(object: I): QueryGetAccountInfoResponse {
    const message = createBaseQueryGetAccountInfoResponse();
    message.accAddress = object.accAddress ?? "";
    message.pubKey = object.pubKey ?? "";
    return message;
  },
};

function createBaseQueryVerifyReferencePayloadLinkRequest(): QueryVerifyReferencePayloadLinkRequest {
  return { referenceId: "", payloadHash: "" };
}

export const QueryVerifyReferencePayloadLinkRequest = {
  encode(message: QueryVerifyReferencePayloadLinkRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.referenceId !== "") {
      writer.uint32(10).string(message.referenceId);
    }
    if (message.payloadHash !== "") {
      writer.uint32(18).string(message.payloadHash);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryVerifyReferencePayloadLinkRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryVerifyReferencePayloadLinkRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.referenceId = reader.string();
          break;
        case 2:
          message.payloadHash = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryVerifyReferencePayloadLinkRequest {
    return {
      referenceId: isSet(object.referenceId) ? String(object.referenceId) : "",
      payloadHash: isSet(object.payloadHash) ? String(object.payloadHash) : "",
    };
  },

  toJSON(message: QueryVerifyReferencePayloadLinkRequest): unknown {
    const obj: any = {};
    message.referenceId !== undefined && (obj.referenceId = message.referenceId);
    message.payloadHash !== undefined && (obj.payloadHash = message.payloadHash);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryVerifyReferencePayloadLinkRequest>, I>>(
    object: I,
  ): QueryVerifyReferencePayloadLinkRequest {
    const message = createBaseQueryVerifyReferencePayloadLinkRequest();
    message.referenceId = object.referenceId ?? "";
    message.payloadHash = object.payloadHash ?? "";
    return message;
  },
};

function createBaseQueryVerifyReferencePayloadLinkResponse(): QueryVerifyReferencePayloadLinkResponse {
  return { isValid: false };
}

export const QueryVerifyReferencePayloadLinkResponse = {
  encode(message: QueryVerifyReferencePayloadLinkResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.isValid === true) {
      writer.uint32(8).bool(message.isValid);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryVerifyReferencePayloadLinkResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryVerifyReferencePayloadLinkResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.isValid = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryVerifyReferencePayloadLinkResponse {
    return { isValid: isSet(object.isValid) ? Boolean(object.isValid) : false };
  },

  toJSON(message: QueryVerifyReferencePayloadLinkResponse): unknown {
    const obj: any = {};
    message.isValid !== undefined && (obj.isValid = message.isValid);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryVerifyReferencePayloadLinkResponse>, I>>(
    object: I,
  ): QueryVerifyReferencePayloadLinkResponse {
    const message = createBaseQueryVerifyReferencePayloadLinkResponse();
    message.isValid = object.isValid ?? false;
    return message;
  },
};

function createBaseQueryGetReferencePayloadLinkRequest(): QueryGetReferencePayloadLinkRequest {
  return { referenceId: "" };
}

export const QueryGetReferencePayloadLinkRequest = {
  encode(message: QueryGetReferencePayloadLinkRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.referenceId !== "") {
      writer.uint32(10).string(message.referenceId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetReferencePayloadLinkRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetReferencePayloadLinkRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.referenceId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetReferencePayloadLinkRequest {
    return { referenceId: isSet(object.referenceId) ? String(object.referenceId) : "" };
  },

  toJSON(message: QueryGetReferencePayloadLinkRequest): unknown {
    const obj: any = {};
    message.referenceId !== undefined && (obj.referenceId = message.referenceId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetReferencePayloadLinkRequest>, I>>(
    object: I,
  ): QueryGetReferencePayloadLinkRequest {
    const message = createBaseQueryGetReferencePayloadLinkRequest();
    message.referenceId = object.referenceId ?? "";
    return message;
  },
};

function createBaseQueryGetReferencePayloadLinkResponse(): QueryGetReferencePayloadLinkResponse {
  return { referencePayloadLinkValue: "" };
}

export const QueryGetReferencePayloadLinkResponse = {
  encode(message: QueryGetReferencePayloadLinkResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.referencePayloadLinkValue !== "") {
      writer.uint32(10).string(message.referencePayloadLinkValue);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetReferencePayloadLinkResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetReferencePayloadLinkResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.referencePayloadLinkValue = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetReferencePayloadLinkResponse {
    return {
      referencePayloadLinkValue: isSet(object.referencePayloadLinkValue)
        ? String(object.referencePayloadLinkValue)
        : "",
    };
  },

  toJSON(message: QueryGetReferencePayloadLinkResponse): unknown {
    const obj: any = {};
    message.referencePayloadLinkValue !== undefined
      && (obj.referencePayloadLinkValue = message.referencePayloadLinkValue);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetReferencePayloadLinkResponse>, I>>(
    object: I,
  ): QueryGetReferencePayloadLinkResponse {
    const message = createBaseQueryGetReferencePayloadLinkResponse();
    message.referencePayloadLinkValue = object.referencePayloadLinkValue ?? "";
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of CreateReferenceId items. */
  CreateReferenceId(request: QueryCreateReferenceIdRequest): Promise<QueryCreateReferenceIdResponse>;
  /** Queries a list of CreateStorageKey items. */
  CreateStorageKey(request: QueryCreateStorageKeyRequest): Promise<QueryCreateStorageKeyResponse>;
  /** Queries a list of CreateReferencePayloadLink items. */
  CreateReferencePayloadLink(
    request: QueryCreateReferencePayloadLinkRequest,
  ): Promise<QueryCreateReferencePayloadLinkResponse>;
  /** Queries a list of VerifySignature items. */
  VerifySignature(request: QueryVerifySignatureRequest): Promise<QueryVerifySignatureResponse>;
  /** Queries a list of GetAccountInfo items. */
  GetAccountInfo(request: QueryGetAccountInfoRequest): Promise<QueryGetAccountInfoResponse>;
  /** Queries a list of VerifyReferencePayloadLink items. */
  VerifyReferencePayloadLink(
    request: QueryVerifyReferencePayloadLinkRequest,
  ): Promise<QueryVerifyReferencePayloadLinkResponse>;
  /** Queries a list of GetReferencePayloadLink items. */
  GetReferencePayloadLink(request: QueryGetReferencePayloadLinkRequest): Promise<QueryGetReferencePayloadLinkResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.CreateReferenceId = this.CreateReferenceId.bind(this);
    this.CreateStorageKey = this.CreateStorageKey.bind(this);
    this.CreateReferencePayloadLink = this.CreateReferencePayloadLink.bind(this);
    this.VerifySignature = this.VerifySignature.bind(this);
    this.GetAccountInfo = this.GetAccountInfo.bind(this);
    this.VerifyReferencePayloadLink = this.VerifyReferencePayloadLink.bind(this);
    this.GetReferencePayloadLink = this.GetReferencePayloadLink.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfesignature.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  CreateReferenceId(request: QueryCreateReferenceIdRequest): Promise<QueryCreateReferenceIdResponse> {
    const data = QueryCreateReferenceIdRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfesignature.Query", "CreateReferenceId", data);
    return promise.then((data) => QueryCreateReferenceIdResponse.decode(new _m0.Reader(data)));
  }

  CreateStorageKey(request: QueryCreateStorageKeyRequest): Promise<QueryCreateStorageKeyResponse> {
    const data = QueryCreateStorageKeyRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfesignature.Query", "CreateStorageKey", data);
    return promise.then((data) => QueryCreateStorageKeyResponse.decode(new _m0.Reader(data)));
  }

  CreateReferencePayloadLink(
    request: QueryCreateReferencePayloadLinkRequest,
  ): Promise<QueryCreateReferencePayloadLinkResponse> {
    const data = QueryCreateReferencePayloadLinkRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfesignature.Query", "CreateReferencePayloadLink", data);
    return promise.then((data) => QueryCreateReferencePayloadLinkResponse.decode(new _m0.Reader(data)));
  }

  VerifySignature(request: QueryVerifySignatureRequest): Promise<QueryVerifySignatureResponse> {
    const data = QueryVerifySignatureRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfesignature.Query", "VerifySignature", data);
    return promise.then((data) => QueryVerifySignatureResponse.decode(new _m0.Reader(data)));
  }

  GetAccountInfo(request: QueryGetAccountInfoRequest): Promise<QueryGetAccountInfoResponse> {
    const data = QueryGetAccountInfoRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfesignature.Query", "GetAccountInfo", data);
    return promise.then((data) => QueryGetAccountInfoResponse.decode(new _m0.Reader(data)));
  }

  VerifyReferencePayloadLink(
    request: QueryVerifyReferencePayloadLinkRequest,
  ): Promise<QueryVerifyReferencePayloadLinkResponse> {
    const data = QueryVerifyReferencePayloadLinkRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfesignature.Query", "VerifyReferencePayloadLink", data);
    return promise.then((data) => QueryVerifyReferencePayloadLinkResponse.decode(new _m0.Reader(data)));
  }

  GetReferencePayloadLink(request: QueryGetReferencePayloadLinkRequest): Promise<QueryGetReferencePayloadLinkResponse> {
    const data = QueryGetReferencePayloadLinkRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfesignature.Query", "GetReferencePayloadLink", data);
    return promise.then((data) => QueryGetReferencePayloadLinkResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
