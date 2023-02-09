/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfesignature";

export interface MsgStoreSignature {
  creator: string;
  storageKey: string;
  signatureJSON: string;
}

export interface MsgStoreSignatureResponse {
  txId: string;
  txTimestamp: string;
}

export interface MsgPublishReferencePayloadLink {
  creator: string;
  key: string;
  value: string;
}

export interface MsgPublishReferencePayloadLinkResponse {
  txTimestamp: string;
}

export interface MsgCreateAccount {
  creator: string;
  accAddressString: string;
  pubKeyString: string;
}

export interface MsgCreateAccountResponse {
  accountNumber: string;
}

function createBaseMsgStoreSignature(): MsgStoreSignature {
  return { creator: "", storageKey: "", signatureJSON: "" };
}

export const MsgStoreSignature = {
  encode(message: MsgStoreSignature, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.storageKey !== "") {
      writer.uint32(18).string(message.storageKey);
    }
    if (message.signatureJSON !== "") {
      writer.uint32(26).string(message.signatureJSON);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgStoreSignature {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgStoreSignature();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.storageKey = reader.string();
          break;
        case 3:
          message.signatureJSON = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgStoreSignature {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      storageKey: isSet(object.storageKey) ? String(object.storageKey) : "",
      signatureJSON: isSet(object.signatureJSON) ? String(object.signatureJSON) : "",
    };
  },

  toJSON(message: MsgStoreSignature): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.storageKey !== undefined && (obj.storageKey = message.storageKey);
    message.signatureJSON !== undefined && (obj.signatureJSON = message.signatureJSON);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgStoreSignature>, I>>(object: I): MsgStoreSignature {
    const message = createBaseMsgStoreSignature();
    message.creator = object.creator ?? "";
    message.storageKey = object.storageKey ?? "";
    message.signatureJSON = object.signatureJSON ?? "";
    return message;
  },
};

function createBaseMsgStoreSignatureResponse(): MsgStoreSignatureResponse {
  return { txId: "", txTimestamp: "" };
}

export const MsgStoreSignatureResponse = {
  encode(message: MsgStoreSignatureResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.txId !== "") {
      writer.uint32(10).string(message.txId);
    }
    if (message.txTimestamp !== "") {
      writer.uint32(18).string(message.txTimestamp);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgStoreSignatureResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgStoreSignatureResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.txId = reader.string();
          break;
        case 2:
          message.txTimestamp = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgStoreSignatureResponse {
    return {
      txId: isSet(object.txId) ? String(object.txId) : "",
      txTimestamp: isSet(object.txTimestamp) ? String(object.txTimestamp) : "",
    };
  },

  toJSON(message: MsgStoreSignatureResponse): unknown {
    const obj: any = {};
    message.txId !== undefined && (obj.txId = message.txId);
    message.txTimestamp !== undefined && (obj.txTimestamp = message.txTimestamp);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgStoreSignatureResponse>, I>>(object: I): MsgStoreSignatureResponse {
    const message = createBaseMsgStoreSignatureResponse();
    message.txId = object.txId ?? "";
    message.txTimestamp = object.txTimestamp ?? "";
    return message;
  },
};

function createBaseMsgPublishReferencePayloadLink(): MsgPublishReferencePayloadLink {
  return { creator: "", key: "", value: "" };
}

export const MsgPublishReferencePayloadLink = {
  encode(message: MsgPublishReferencePayloadLink, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.key !== "") {
      writer.uint32(18).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(26).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPublishReferencePayloadLink {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPublishReferencePayloadLink();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.key = reader.string();
          break;
        case 3:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPublishReferencePayloadLink {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: MsgPublishReferencePayloadLink): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgPublishReferencePayloadLink>, I>>(
    object: I,
  ): MsgPublishReferencePayloadLink {
    const message = createBaseMsgPublishReferencePayloadLink();
    message.creator = object.creator ?? "";
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseMsgPublishReferencePayloadLinkResponse(): MsgPublishReferencePayloadLinkResponse {
  return { txTimestamp: "" };
}

export const MsgPublishReferencePayloadLinkResponse = {
  encode(message: MsgPublishReferencePayloadLinkResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.txTimestamp !== "") {
      writer.uint32(10).string(message.txTimestamp);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPublishReferencePayloadLinkResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPublishReferencePayloadLinkResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.txTimestamp = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgPublishReferencePayloadLinkResponse {
    return { txTimestamp: isSet(object.txTimestamp) ? String(object.txTimestamp) : "" };
  },

  toJSON(message: MsgPublishReferencePayloadLinkResponse): unknown {
    const obj: any = {};
    message.txTimestamp !== undefined && (obj.txTimestamp = message.txTimestamp);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgPublishReferencePayloadLinkResponse>, I>>(
    object: I,
  ): MsgPublishReferencePayloadLinkResponse {
    const message = createBaseMsgPublishReferencePayloadLinkResponse();
    message.txTimestamp = object.txTimestamp ?? "";
    return message;
  },
};

function createBaseMsgCreateAccount(): MsgCreateAccount {
  return { creator: "", accAddressString: "", pubKeyString: "" };
}

export const MsgCreateAccount = {
  encode(message: MsgCreateAccount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.accAddressString !== "") {
      writer.uint32(18).string(message.accAddressString);
    }
    if (message.pubKeyString !== "") {
      writer.uint32(26).string(message.pubKeyString);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateAccount {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateAccount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.accAddressString = reader.string();
          break;
        case 3:
          message.pubKeyString = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateAccount {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      accAddressString: isSet(object.accAddressString) ? String(object.accAddressString) : "",
      pubKeyString: isSet(object.pubKeyString) ? String(object.pubKeyString) : "",
    };
  },

  toJSON(message: MsgCreateAccount): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.accAddressString !== undefined && (obj.accAddressString = message.accAddressString);
    message.pubKeyString !== undefined && (obj.pubKeyString = message.pubKeyString);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateAccount>, I>>(object: I): MsgCreateAccount {
    const message = createBaseMsgCreateAccount();
    message.creator = object.creator ?? "";
    message.accAddressString = object.accAddressString ?? "";
    message.pubKeyString = object.pubKeyString ?? "";
    return message;
  },
};

function createBaseMsgCreateAccountResponse(): MsgCreateAccountResponse {
  return { accountNumber: "" };
}

export const MsgCreateAccountResponse = {
  encode(message: MsgCreateAccountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.accountNumber !== "") {
      writer.uint32(10).string(message.accountNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateAccountResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateAccountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.accountNumber = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateAccountResponse {
    return { accountNumber: isSet(object.accountNumber) ? String(object.accountNumber) : "" };
  },

  toJSON(message: MsgCreateAccountResponse): unknown {
    const obj: any = {};
    message.accountNumber !== undefined && (obj.accountNumber = message.accountNumber);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateAccountResponse>, I>>(object: I): MsgCreateAccountResponse {
    const message = createBaseMsgCreateAccountResponse();
    message.accountNumber = object.accountNumber ?? "";
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  StoreSignature(request: MsgStoreSignature): Promise<MsgStoreSignatureResponse>;
  PublishReferencePayloadLink(request: MsgPublishReferencePayloadLink): Promise<MsgPublishReferencePayloadLinkResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateAccount(request: MsgCreateAccount): Promise<MsgCreateAccountResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.StoreSignature = this.StoreSignature.bind(this);
    this.PublishReferencePayloadLink = this.PublishReferencePayloadLink.bind(this);
    this.CreateAccount = this.CreateAccount.bind(this);
  }
  StoreSignature(request: MsgStoreSignature): Promise<MsgStoreSignatureResponse> {
    const data = MsgStoreSignature.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfesignature.Msg", "StoreSignature", data);
    return promise.then((data) => MsgStoreSignatureResponse.decode(new _m0.Reader(data)));
  }

  PublishReferencePayloadLink(
    request: MsgPublishReferencePayloadLink,
  ): Promise<MsgPublishReferencePayloadLinkResponse> {
    const data = MsgPublishReferencePayloadLink.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfesignature.Msg", "PublishReferencePayloadLink", data);
    return promise.then((data) => MsgPublishReferencePayloadLinkResponse.decode(new _m0.Reader(data)));
  }

  CreateAccount(request: MsgCreateAccount): Promise<MsgCreateAccountResponse> {
    const data = MsgCreateAccount.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfesignature.Msg", "CreateAccount", data);
    return promise.then((data) => MsgCreateAccountResponse.decode(new _m0.Reader(data)));
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
