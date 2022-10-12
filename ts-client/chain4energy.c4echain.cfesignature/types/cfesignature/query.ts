/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../cfesignature/params";

export const protobufPackage = "chain4energy.c4echain.cfesignature";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

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

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
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
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryCreateReferenceIdRequest: object = { creator: "" };

export const QueryCreateReferenceIdRequest = {
  encode(
    message: QueryCreateReferenceIdRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryCreateReferenceIdRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCreateReferenceIdRequest,
    } as QueryCreateReferenceIdRequest;
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
    const message = {
      ...baseQueryCreateReferenceIdRequest,
    } as QueryCreateReferenceIdRequest;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    return message;
  },

  toJSON(message: QueryCreateReferenceIdRequest): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCreateReferenceIdRequest>
  ): QueryCreateReferenceIdRequest {
    const message = {
      ...baseQueryCreateReferenceIdRequest,
    } as QueryCreateReferenceIdRequest;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    return message;
  },
};

const baseQueryCreateReferenceIdResponse: object = { referenceId: "" };

export const QueryCreateReferenceIdResponse = {
  encode(
    message: QueryCreateReferenceIdResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.referenceId !== "") {
      writer.uint32(10).string(message.referenceId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryCreateReferenceIdResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCreateReferenceIdResponse,
    } as QueryCreateReferenceIdResponse;
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
    const message = {
      ...baseQueryCreateReferenceIdResponse,
    } as QueryCreateReferenceIdResponse;
    if (object.referenceId !== undefined && object.referenceId !== null) {
      message.referenceId = String(object.referenceId);
    } else {
      message.referenceId = "";
    }
    return message;
  },

  toJSON(message: QueryCreateReferenceIdResponse): unknown {
    const obj: any = {};
    message.referenceId !== undefined &&
      (obj.referenceId = message.referenceId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCreateReferenceIdResponse>
  ): QueryCreateReferenceIdResponse {
    const message = {
      ...baseQueryCreateReferenceIdResponse,
    } as QueryCreateReferenceIdResponse;
    if (object.referenceId !== undefined && object.referenceId !== null) {
      message.referenceId = object.referenceId;
    } else {
      message.referenceId = "";
    }
    return message;
  },
};

const baseQueryCreateStorageKeyRequest: object = {
  targetAccAddress: "",
  referenceId: "",
};

export const QueryCreateStorageKeyRequest = {
  encode(
    message: QueryCreateStorageKeyRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.targetAccAddress !== "") {
      writer.uint32(10).string(message.targetAccAddress);
    }
    if (message.referenceId !== "") {
      writer.uint32(18).string(message.referenceId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryCreateStorageKeyRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCreateStorageKeyRequest,
    } as QueryCreateStorageKeyRequest;
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
    const message = {
      ...baseQueryCreateStorageKeyRequest,
    } as QueryCreateStorageKeyRequest;
    if (
      object.targetAccAddress !== undefined &&
      object.targetAccAddress !== null
    ) {
      message.targetAccAddress = String(object.targetAccAddress);
    } else {
      message.targetAccAddress = "";
    }
    if (object.referenceId !== undefined && object.referenceId !== null) {
      message.referenceId = String(object.referenceId);
    } else {
      message.referenceId = "";
    }
    return message;
  },

  toJSON(message: QueryCreateStorageKeyRequest): unknown {
    const obj: any = {};
    message.targetAccAddress !== undefined &&
      (obj.targetAccAddress = message.targetAccAddress);
    message.referenceId !== undefined &&
      (obj.referenceId = message.referenceId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCreateStorageKeyRequest>
  ): QueryCreateStorageKeyRequest {
    const message = {
      ...baseQueryCreateStorageKeyRequest,
    } as QueryCreateStorageKeyRequest;
    if (
      object.targetAccAddress !== undefined &&
      object.targetAccAddress !== null
    ) {
      message.targetAccAddress = object.targetAccAddress;
    } else {
      message.targetAccAddress = "";
    }
    if (object.referenceId !== undefined && object.referenceId !== null) {
      message.referenceId = object.referenceId;
    } else {
      message.referenceId = "";
    }
    return message;
  },
};

const baseQueryCreateStorageKeyResponse: object = { storageKey: "" };

export const QueryCreateStorageKeyResponse = {
  encode(
    message: QueryCreateStorageKeyResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.storageKey !== "") {
      writer.uint32(10).string(message.storageKey);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryCreateStorageKeyResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCreateStorageKeyResponse,
    } as QueryCreateStorageKeyResponse;
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
    const message = {
      ...baseQueryCreateStorageKeyResponse,
    } as QueryCreateStorageKeyResponse;
    if (object.storageKey !== undefined && object.storageKey !== null) {
      message.storageKey = String(object.storageKey);
    } else {
      message.storageKey = "";
    }
    return message;
  },

  toJSON(message: QueryCreateStorageKeyResponse): unknown {
    const obj: any = {};
    message.storageKey !== undefined && (obj.storageKey = message.storageKey);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCreateStorageKeyResponse>
  ): QueryCreateStorageKeyResponse {
    const message = {
      ...baseQueryCreateStorageKeyResponse,
    } as QueryCreateStorageKeyResponse;
    if (object.storageKey !== undefined && object.storageKey !== null) {
      message.storageKey = object.storageKey;
    } else {
      message.storageKey = "";
    }
    return message;
  },
};

const baseQueryCreateReferencePayloadLinkRequest: object = {
  referenceId: "",
  payloadHash: "",
};

export const QueryCreateReferencePayloadLinkRequest = {
  encode(
    message: QueryCreateReferencePayloadLinkRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.referenceId !== "") {
      writer.uint32(10).string(message.referenceId);
    }
    if (message.payloadHash !== "") {
      writer.uint32(18).string(message.payloadHash);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryCreateReferencePayloadLinkRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCreateReferencePayloadLinkRequest,
    } as QueryCreateReferencePayloadLinkRequest;
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
    const message = {
      ...baseQueryCreateReferencePayloadLinkRequest,
    } as QueryCreateReferencePayloadLinkRequest;
    if (object.referenceId !== undefined && object.referenceId !== null) {
      message.referenceId = String(object.referenceId);
    } else {
      message.referenceId = "";
    }
    if (object.payloadHash !== undefined && object.payloadHash !== null) {
      message.payloadHash = String(object.payloadHash);
    } else {
      message.payloadHash = "";
    }
    return message;
  },

  toJSON(message: QueryCreateReferencePayloadLinkRequest): unknown {
    const obj: any = {};
    message.referenceId !== undefined &&
      (obj.referenceId = message.referenceId);
    message.payloadHash !== undefined &&
      (obj.payloadHash = message.payloadHash);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCreateReferencePayloadLinkRequest>
  ): QueryCreateReferencePayloadLinkRequest {
    const message = {
      ...baseQueryCreateReferencePayloadLinkRequest,
    } as QueryCreateReferencePayloadLinkRequest;
    if (object.referenceId !== undefined && object.referenceId !== null) {
      message.referenceId = object.referenceId;
    } else {
      message.referenceId = "";
    }
    if (object.payloadHash !== undefined && object.payloadHash !== null) {
      message.payloadHash = object.payloadHash;
    } else {
      message.payloadHash = "";
    }
    return message;
  },
};

const baseQueryCreateReferencePayloadLinkResponse: object = {
  referenceKey: "",
  referenceValue: "",
};

export const QueryCreateReferencePayloadLinkResponse = {
  encode(
    message: QueryCreateReferencePayloadLinkResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.referenceKey !== "") {
      writer.uint32(10).string(message.referenceKey);
    }
    if (message.referenceValue !== "") {
      writer.uint32(18).string(message.referenceValue);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryCreateReferencePayloadLinkResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCreateReferencePayloadLinkResponse,
    } as QueryCreateReferencePayloadLinkResponse;
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
    const message = {
      ...baseQueryCreateReferencePayloadLinkResponse,
    } as QueryCreateReferencePayloadLinkResponse;
    if (object.referenceKey !== undefined && object.referenceKey !== null) {
      message.referenceKey = String(object.referenceKey);
    } else {
      message.referenceKey = "";
    }
    if (object.referenceValue !== undefined && object.referenceValue !== null) {
      message.referenceValue = String(object.referenceValue);
    } else {
      message.referenceValue = "";
    }
    return message;
  },

  toJSON(message: QueryCreateReferencePayloadLinkResponse): unknown {
    const obj: any = {};
    message.referenceKey !== undefined &&
      (obj.referenceKey = message.referenceKey);
    message.referenceValue !== undefined &&
      (obj.referenceValue = message.referenceValue);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCreateReferencePayloadLinkResponse>
  ): QueryCreateReferencePayloadLinkResponse {
    const message = {
      ...baseQueryCreateReferencePayloadLinkResponse,
    } as QueryCreateReferencePayloadLinkResponse;
    if (object.referenceKey !== undefined && object.referenceKey !== null) {
      message.referenceKey = object.referenceKey;
    } else {
      message.referenceKey = "";
    }
    if (object.referenceValue !== undefined && object.referenceValue !== null) {
      message.referenceValue = object.referenceValue;
    } else {
      message.referenceValue = "";
    }
    return message;
  },
};

const baseQueryVerifySignatureRequest: object = {
  referenceId: "",
  targetAccAddress: "",
};

export const QueryVerifySignatureRequest = {
  encode(
    message: QueryVerifySignatureRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.referenceId !== "") {
      writer.uint32(10).string(message.referenceId);
    }
    if (message.targetAccAddress !== "") {
      writer.uint32(18).string(message.targetAccAddress);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryVerifySignatureRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryVerifySignatureRequest,
    } as QueryVerifySignatureRequest;
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
    const message = {
      ...baseQueryVerifySignatureRequest,
    } as QueryVerifySignatureRequest;
    if (object.referenceId !== undefined && object.referenceId !== null) {
      message.referenceId = String(object.referenceId);
    } else {
      message.referenceId = "";
    }
    if (
      object.targetAccAddress !== undefined &&
      object.targetAccAddress !== null
    ) {
      message.targetAccAddress = String(object.targetAccAddress);
    } else {
      message.targetAccAddress = "";
    }
    return message;
  },

  toJSON(message: QueryVerifySignatureRequest): unknown {
    const obj: any = {};
    message.referenceId !== undefined &&
      (obj.referenceId = message.referenceId);
    message.targetAccAddress !== undefined &&
      (obj.targetAccAddress = message.targetAccAddress);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryVerifySignatureRequest>
  ): QueryVerifySignatureRequest {
    const message = {
      ...baseQueryVerifySignatureRequest,
    } as QueryVerifySignatureRequest;
    if (object.referenceId !== undefined && object.referenceId !== null) {
      message.referenceId = object.referenceId;
    } else {
      message.referenceId = "";
    }
    if (
      object.targetAccAddress !== undefined &&
      object.targetAccAddress !== null
    ) {
      message.targetAccAddress = object.targetAccAddress;
    } else {
      message.targetAccAddress = "";
    }
    return message;
  },
};

const baseQueryVerifySignatureResponse: object = {
  signature: "",
  algorithm: "",
  certificate: "",
  timestamp: "",
  valid: "",
};

export const QueryVerifySignatureResponse = {
  encode(
    message: QueryVerifySignatureResponse,
    writer: Writer = Writer.create()
  ): Writer {
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

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryVerifySignatureResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryVerifySignatureResponse,
    } as QueryVerifySignatureResponse;
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
    const message = {
      ...baseQueryVerifySignatureResponse,
    } as QueryVerifySignatureResponse;
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = String(object.signature);
    } else {
      message.signature = "";
    }
    if (object.algorithm !== undefined && object.algorithm !== null) {
      message.algorithm = String(object.algorithm);
    } else {
      message.algorithm = "";
    }
    if (object.certificate !== undefined && object.certificate !== null) {
      message.certificate = String(object.certificate);
    } else {
      message.certificate = "";
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = String(object.timestamp);
    } else {
      message.timestamp = "";
    }
    if (object.valid !== undefined && object.valid !== null) {
      message.valid = String(object.valid);
    } else {
      message.valid = "";
    }
    return message;
  },

  toJSON(message: QueryVerifySignatureResponse): unknown {
    const obj: any = {};
    message.signature !== undefined && (obj.signature = message.signature);
    message.algorithm !== undefined && (obj.algorithm = message.algorithm);
    message.certificate !== undefined &&
      (obj.certificate = message.certificate);
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    message.valid !== undefined && (obj.valid = message.valid);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryVerifySignatureResponse>
  ): QueryVerifySignatureResponse {
    const message = {
      ...baseQueryVerifySignatureResponse,
    } as QueryVerifySignatureResponse;
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = object.signature;
    } else {
      message.signature = "";
    }
    if (object.algorithm !== undefined && object.algorithm !== null) {
      message.algorithm = object.algorithm;
    } else {
      message.algorithm = "";
    }
    if (object.certificate !== undefined && object.certificate !== null) {
      message.certificate = object.certificate;
    } else {
      message.certificate = "";
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = object.timestamp;
    } else {
      message.timestamp = "";
    }
    if (object.valid !== undefined && object.valid !== null) {
      message.valid = object.valid;
    } else {
      message.valid = "";
    }
    return message;
  },
};

const baseQueryGetAccountInfoRequest: object = { accAddressString: "" };

export const QueryGetAccountInfoRequest = {
  encode(
    message: QueryGetAccountInfoRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.accAddressString !== "") {
      writer.uint32(10).string(message.accAddressString);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetAccountInfoRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAccountInfoRequest,
    } as QueryGetAccountInfoRequest;
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
    const message = {
      ...baseQueryGetAccountInfoRequest,
    } as QueryGetAccountInfoRequest;
    if (
      object.accAddressString !== undefined &&
      object.accAddressString !== null
    ) {
      message.accAddressString = String(object.accAddressString);
    } else {
      message.accAddressString = "";
    }
    return message;
  },

  toJSON(message: QueryGetAccountInfoRequest): unknown {
    const obj: any = {};
    message.accAddressString !== undefined &&
      (obj.accAddressString = message.accAddressString);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAccountInfoRequest>
  ): QueryGetAccountInfoRequest {
    const message = {
      ...baseQueryGetAccountInfoRequest,
    } as QueryGetAccountInfoRequest;
    if (
      object.accAddressString !== undefined &&
      object.accAddressString !== null
    ) {
      message.accAddressString = object.accAddressString;
    } else {
      message.accAddressString = "";
    }
    return message;
  },
};

const baseQueryGetAccountInfoResponse: object = { accAddress: "", pubKey: "" };

export const QueryGetAccountInfoResponse = {
  encode(
    message: QueryGetAccountInfoResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.accAddress !== "") {
      writer.uint32(10).string(message.accAddress);
    }
    if (message.pubKey !== "") {
      writer.uint32(18).string(message.pubKey);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetAccountInfoResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetAccountInfoResponse,
    } as QueryGetAccountInfoResponse;
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
    const message = {
      ...baseQueryGetAccountInfoResponse,
    } as QueryGetAccountInfoResponse;
    if (object.accAddress !== undefined && object.accAddress !== null) {
      message.accAddress = String(object.accAddress);
    } else {
      message.accAddress = "";
    }
    if (object.pubKey !== undefined && object.pubKey !== null) {
      message.pubKey = String(object.pubKey);
    } else {
      message.pubKey = "";
    }
    return message;
  },

  toJSON(message: QueryGetAccountInfoResponse): unknown {
    const obj: any = {};
    message.accAddress !== undefined && (obj.accAddress = message.accAddress);
    message.pubKey !== undefined && (obj.pubKey = message.pubKey);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetAccountInfoResponse>
  ): QueryGetAccountInfoResponse {
    const message = {
      ...baseQueryGetAccountInfoResponse,
    } as QueryGetAccountInfoResponse;
    if (object.accAddress !== undefined && object.accAddress !== null) {
      message.accAddress = object.accAddress;
    } else {
      message.accAddress = "";
    }
    if (object.pubKey !== undefined && object.pubKey !== null) {
      message.pubKey = object.pubKey;
    } else {
      message.pubKey = "";
    }
    return message;
  },
};

const baseQueryVerifyReferencePayloadLinkRequest: object = {
  referenceId: "",
  payloadHash: "",
};

export const QueryVerifyReferencePayloadLinkRequest = {
  encode(
    message: QueryVerifyReferencePayloadLinkRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.referenceId !== "") {
      writer.uint32(10).string(message.referenceId);
    }
    if (message.payloadHash !== "") {
      writer.uint32(18).string(message.payloadHash);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryVerifyReferencePayloadLinkRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryVerifyReferencePayloadLinkRequest,
    } as QueryVerifyReferencePayloadLinkRequest;
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
    const message = {
      ...baseQueryVerifyReferencePayloadLinkRequest,
    } as QueryVerifyReferencePayloadLinkRequest;
    if (object.referenceId !== undefined && object.referenceId !== null) {
      message.referenceId = String(object.referenceId);
    } else {
      message.referenceId = "";
    }
    if (object.payloadHash !== undefined && object.payloadHash !== null) {
      message.payloadHash = String(object.payloadHash);
    } else {
      message.payloadHash = "";
    }
    return message;
  },

  toJSON(message: QueryVerifyReferencePayloadLinkRequest): unknown {
    const obj: any = {};
    message.referenceId !== undefined &&
      (obj.referenceId = message.referenceId);
    message.payloadHash !== undefined &&
      (obj.payloadHash = message.payloadHash);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryVerifyReferencePayloadLinkRequest>
  ): QueryVerifyReferencePayloadLinkRequest {
    const message = {
      ...baseQueryVerifyReferencePayloadLinkRequest,
    } as QueryVerifyReferencePayloadLinkRequest;
    if (object.referenceId !== undefined && object.referenceId !== null) {
      message.referenceId = object.referenceId;
    } else {
      message.referenceId = "";
    }
    if (object.payloadHash !== undefined && object.payloadHash !== null) {
      message.payloadHash = object.payloadHash;
    } else {
      message.payloadHash = "";
    }
    return message;
  },
};

const baseQueryVerifyReferencePayloadLinkResponse: object = { isValid: false };

export const QueryVerifyReferencePayloadLinkResponse = {
  encode(
    message: QueryVerifyReferencePayloadLinkResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.isValid === true) {
      writer.uint32(8).bool(message.isValid);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryVerifyReferencePayloadLinkResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryVerifyReferencePayloadLinkResponse,
    } as QueryVerifyReferencePayloadLinkResponse;
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
    const message = {
      ...baseQueryVerifyReferencePayloadLinkResponse,
    } as QueryVerifyReferencePayloadLinkResponse;
    if (object.isValid !== undefined && object.isValid !== null) {
      message.isValid = Boolean(object.isValid);
    } else {
      message.isValid = false;
    }
    return message;
  },

  toJSON(message: QueryVerifyReferencePayloadLinkResponse): unknown {
    const obj: any = {};
    message.isValid !== undefined && (obj.isValid = message.isValid);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryVerifyReferencePayloadLinkResponse>
  ): QueryVerifyReferencePayloadLinkResponse {
    const message = {
      ...baseQueryVerifyReferencePayloadLinkResponse,
    } as QueryVerifyReferencePayloadLinkResponse;
    if (object.isValid !== undefined && object.isValid !== null) {
      message.isValid = object.isValid;
    } else {
      message.isValid = false;
    }
    return message;
  },
};

const baseQueryGetReferencePayloadLinkRequest: object = { referenceId: "" };

export const QueryGetReferencePayloadLinkRequest = {
  encode(
    message: QueryGetReferencePayloadLinkRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.referenceId !== "") {
      writer.uint32(10).string(message.referenceId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetReferencePayloadLinkRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetReferencePayloadLinkRequest,
    } as QueryGetReferencePayloadLinkRequest;
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
    const message = {
      ...baseQueryGetReferencePayloadLinkRequest,
    } as QueryGetReferencePayloadLinkRequest;
    if (object.referenceId !== undefined && object.referenceId !== null) {
      message.referenceId = String(object.referenceId);
    } else {
      message.referenceId = "";
    }
    return message;
  },

  toJSON(message: QueryGetReferencePayloadLinkRequest): unknown {
    const obj: any = {};
    message.referenceId !== undefined &&
      (obj.referenceId = message.referenceId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetReferencePayloadLinkRequest>
  ): QueryGetReferencePayloadLinkRequest {
    const message = {
      ...baseQueryGetReferencePayloadLinkRequest,
    } as QueryGetReferencePayloadLinkRequest;
    if (object.referenceId !== undefined && object.referenceId !== null) {
      message.referenceId = object.referenceId;
    } else {
      message.referenceId = "";
    }
    return message;
  },
};

const baseQueryGetReferencePayloadLinkResponse: object = {
  referencePayloadLinkValue: "",
};

export const QueryGetReferencePayloadLinkResponse = {
  encode(
    message: QueryGetReferencePayloadLinkResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.referencePayloadLinkValue !== "") {
      writer.uint32(10).string(message.referencePayloadLinkValue);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetReferencePayloadLinkResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetReferencePayloadLinkResponse,
    } as QueryGetReferencePayloadLinkResponse;
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
    const message = {
      ...baseQueryGetReferencePayloadLinkResponse,
    } as QueryGetReferencePayloadLinkResponse;
    if (
      object.referencePayloadLinkValue !== undefined &&
      object.referencePayloadLinkValue !== null
    ) {
      message.referencePayloadLinkValue = String(
        object.referencePayloadLinkValue
      );
    } else {
      message.referencePayloadLinkValue = "";
    }
    return message;
  },

  toJSON(message: QueryGetReferencePayloadLinkResponse): unknown {
    const obj: any = {};
    message.referencePayloadLinkValue !== undefined &&
      (obj.referencePayloadLinkValue = message.referencePayloadLinkValue);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetReferencePayloadLinkResponse>
  ): QueryGetReferencePayloadLinkResponse {
    const message = {
      ...baseQueryGetReferencePayloadLinkResponse,
    } as QueryGetReferencePayloadLinkResponse;
    if (
      object.referencePayloadLinkValue !== undefined &&
      object.referencePayloadLinkValue !== null
    ) {
      message.referencePayloadLinkValue = object.referencePayloadLinkValue;
    } else {
      message.referencePayloadLinkValue = "";
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of CreateReferenceId items. */
  CreateReferenceId(
    request: QueryCreateReferenceIdRequest
  ): Promise<QueryCreateReferenceIdResponse>;
  /** Queries a list of CreateStorageKey items. */
  CreateStorageKey(
    request: QueryCreateStorageKeyRequest
  ): Promise<QueryCreateStorageKeyResponse>;
  /** Queries a list of CreateReferencePayloadLink items. */
  CreateReferencePayloadLink(
    request: QueryCreateReferencePayloadLinkRequest
  ): Promise<QueryCreateReferencePayloadLinkResponse>;
  /** Queries a list of VerifySignature items. */
  VerifySignature(
    request: QueryVerifySignatureRequest
  ): Promise<QueryVerifySignatureResponse>;
  /** Queries a list of GetAccountInfo items. */
  GetAccountInfo(
    request: QueryGetAccountInfoRequest
  ): Promise<QueryGetAccountInfoResponse>;
  /** Queries a list of VerifyReferencePayloadLink items. */
  VerifyReferencePayloadLink(
    request: QueryVerifyReferencePayloadLinkRequest
  ): Promise<QueryVerifyReferencePayloadLinkResponse>;
  /** Queries a list of GetReferencePayloadLink items. */
  GetReferencePayloadLink(
    request: QueryGetReferencePayloadLinkRequest
  ): Promise<QueryGetReferencePayloadLinkResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfesignature.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  CreateReferenceId(
    request: QueryCreateReferenceIdRequest
  ): Promise<QueryCreateReferenceIdResponse> {
    const data = QueryCreateReferenceIdRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfesignature.Query",
      "CreateReferenceId",
      data
    );
    return promise.then((data) =>
      QueryCreateReferenceIdResponse.decode(new Reader(data))
    );
  }

  CreateStorageKey(
    request: QueryCreateStorageKeyRequest
  ): Promise<QueryCreateStorageKeyResponse> {
    const data = QueryCreateStorageKeyRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfesignature.Query",
      "CreateStorageKey",
      data
    );
    return promise.then((data) =>
      QueryCreateStorageKeyResponse.decode(new Reader(data))
    );
  }

  CreateReferencePayloadLink(
    request: QueryCreateReferencePayloadLinkRequest
  ): Promise<QueryCreateReferencePayloadLinkResponse> {
    const data = QueryCreateReferencePayloadLinkRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfesignature.Query",
      "CreateReferencePayloadLink",
      data
    );
    return promise.then((data) =>
      QueryCreateReferencePayloadLinkResponse.decode(new Reader(data))
    );
  }

  VerifySignature(
    request: QueryVerifySignatureRequest
  ): Promise<QueryVerifySignatureResponse> {
    const data = QueryVerifySignatureRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfesignature.Query",
      "VerifySignature",
      data
    );
    return promise.then((data) =>
      QueryVerifySignatureResponse.decode(new Reader(data))
    );
  }

  GetAccountInfo(
    request: QueryGetAccountInfoRequest
  ): Promise<QueryGetAccountInfoResponse> {
    const data = QueryGetAccountInfoRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfesignature.Query",
      "GetAccountInfo",
      data
    );
    return promise.then((data) =>
      QueryGetAccountInfoResponse.decode(new Reader(data))
    );
  }

  VerifyReferencePayloadLink(
    request: QueryVerifyReferencePayloadLinkRequest
  ): Promise<QueryVerifyReferencePayloadLinkResponse> {
    const data = QueryVerifyReferencePayloadLinkRequest.encode(
      request
    ).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfesignature.Query",
      "VerifyReferencePayloadLink",
      data
    );
    return promise.then((data) =>
      QueryVerifyReferencePayloadLinkResponse.decode(new Reader(data))
    );
  }

  GetReferencePayloadLink(
    request: QueryGetReferencePayloadLinkRequest
  ): Promise<QueryGetReferencePayloadLinkResponse> {
    const data = QueryGetReferencePayloadLinkRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfesignature.Query",
      "GetReferencePayloadLink",
      data
    );
    return promise.then((data) =>
      QueryGetReferencePayloadLinkResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
