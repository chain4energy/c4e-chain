/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

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
  accountId: string;
}

const baseMsgStoreSignature: object = {
  creator: "",
  storageKey: "",
  signatureJSON: "",
};

export const MsgStoreSignature = {
  encode(message: MsgStoreSignature, writer: Writer = Writer.create()): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): MsgStoreSignature {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgStoreSignature } as MsgStoreSignature;
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
    const message = { ...baseMsgStoreSignature } as MsgStoreSignature;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.storageKey !== undefined && object.storageKey !== null) {
      message.storageKey = String(object.storageKey);
    } else {
      message.storageKey = "";
    }
    if (object.signatureJSON !== undefined && object.signatureJSON !== null) {
      message.signatureJSON = String(object.signatureJSON);
    } else {
      message.signatureJSON = "";
    }
    return message;
  },

  toJSON(message: MsgStoreSignature): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.storageKey !== undefined && (obj.storageKey = message.storageKey);
    message.signatureJSON !== undefined &&
      (obj.signatureJSON = message.signatureJSON);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgStoreSignature>): MsgStoreSignature {
    const message = { ...baseMsgStoreSignature } as MsgStoreSignature;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.storageKey !== undefined && object.storageKey !== null) {
      message.storageKey = object.storageKey;
    } else {
      message.storageKey = "";
    }
    if (object.signatureJSON !== undefined && object.signatureJSON !== null) {
      message.signatureJSON = object.signatureJSON;
    } else {
      message.signatureJSON = "";
    }
    return message;
  },
};

const baseMsgStoreSignatureResponse: object = { txId: "", txTimestamp: "" };

export const MsgStoreSignatureResponse = {
  encode(
    message: MsgStoreSignatureResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.txId !== "") {
      writer.uint32(10).string(message.txId);
    }
    if (message.txTimestamp !== "") {
      writer.uint32(18).string(message.txTimestamp);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgStoreSignatureResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgStoreSignatureResponse,
    } as MsgStoreSignatureResponse;
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
    const message = {
      ...baseMsgStoreSignatureResponse,
    } as MsgStoreSignatureResponse;
    if (object.txId !== undefined && object.txId !== null) {
      message.txId = String(object.txId);
    } else {
      message.txId = "";
    }
    if (object.txTimestamp !== undefined && object.txTimestamp !== null) {
      message.txTimestamp = String(object.txTimestamp);
    } else {
      message.txTimestamp = "";
    }
    return message;
  },

  toJSON(message: MsgStoreSignatureResponse): unknown {
    const obj: any = {};
    message.txId !== undefined && (obj.txId = message.txId);
    message.txTimestamp !== undefined &&
      (obj.txTimestamp = message.txTimestamp);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgStoreSignatureResponse>
  ): MsgStoreSignatureResponse {
    const message = {
      ...baseMsgStoreSignatureResponse,
    } as MsgStoreSignatureResponse;
    if (object.txId !== undefined && object.txId !== null) {
      message.txId = object.txId;
    } else {
      message.txId = "";
    }
    if (object.txTimestamp !== undefined && object.txTimestamp !== null) {
      message.txTimestamp = object.txTimestamp;
    } else {
      message.txTimestamp = "";
    }
    return message;
  },
};

const baseMsgPublishReferencePayloadLink: object = {
  creator: "",
  key: "",
  value: "",
};

export const MsgPublishReferencePayloadLink = {
  encode(
    message: MsgPublishReferencePayloadLink,
    writer: Writer = Writer.create()
  ): Writer {
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

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgPublishReferencePayloadLink {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgPublishReferencePayloadLink,
    } as MsgPublishReferencePayloadLink;
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
    const message = {
      ...baseMsgPublishReferencePayloadLink,
    } as MsgPublishReferencePayloadLink;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = String(object.key);
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = String(object.value);
    } else {
      message.value = "";
    }
    return message;
  },

  toJSON(message: MsgPublishReferencePayloadLink): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgPublishReferencePayloadLink>
  ): MsgPublishReferencePayloadLink {
    const message = {
      ...baseMsgPublishReferencePayloadLink,
    } as MsgPublishReferencePayloadLink;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.key !== undefined && object.key !== null) {
      message.key = object.key;
    } else {
      message.key = "";
    }
    if (object.value !== undefined && object.value !== null) {
      message.value = object.value;
    } else {
      message.value = "";
    }
    return message;
  },
};

const baseMsgPublishReferencePayloadLinkResponse: object = { txTimestamp: "" };

export const MsgPublishReferencePayloadLinkResponse = {
  encode(
    message: MsgPublishReferencePayloadLinkResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.txTimestamp !== "") {
      writer.uint32(10).string(message.txTimestamp);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgPublishReferencePayloadLinkResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgPublishReferencePayloadLinkResponse,
    } as MsgPublishReferencePayloadLinkResponse;
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
    const message = {
      ...baseMsgPublishReferencePayloadLinkResponse,
    } as MsgPublishReferencePayloadLinkResponse;
    if (object.txTimestamp !== undefined && object.txTimestamp !== null) {
      message.txTimestamp = String(object.txTimestamp);
    } else {
      message.txTimestamp = "";
    }
    return message;
  },

  toJSON(message: MsgPublishReferencePayloadLinkResponse): unknown {
    const obj: any = {};
    message.txTimestamp !== undefined &&
      (obj.txTimestamp = message.txTimestamp);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgPublishReferencePayloadLinkResponse>
  ): MsgPublishReferencePayloadLinkResponse {
    const message = {
      ...baseMsgPublishReferencePayloadLinkResponse,
    } as MsgPublishReferencePayloadLinkResponse;
    if (object.txTimestamp !== undefined && object.txTimestamp !== null) {
      message.txTimestamp = object.txTimestamp;
    } else {
      message.txTimestamp = "";
    }
    return message;
  },
};

const baseMsgCreateAccount: object = {
  creator: "",
  accAddressString: "",
  pubKeyString: "",
};

export const MsgCreateAccount = {
  encode(message: MsgCreateAccount, writer: Writer = Writer.create()): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): MsgCreateAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateAccount } as MsgCreateAccount;
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
    const message = { ...baseMsgCreateAccount } as MsgCreateAccount;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (
      object.accAddressString !== undefined &&
      object.accAddressString !== null
    ) {
      message.accAddressString = String(object.accAddressString);
    } else {
      message.accAddressString = "";
    }
    if (object.pubKeyString !== undefined && object.pubKeyString !== null) {
      message.pubKeyString = String(object.pubKeyString);
    } else {
      message.pubKeyString = "";
    }
    return message;
  },

  toJSON(message: MsgCreateAccount): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.accAddressString !== undefined &&
      (obj.accAddressString = message.accAddressString);
    message.pubKeyString !== undefined &&
      (obj.pubKeyString = message.pubKeyString);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateAccount>): MsgCreateAccount {
    const message = { ...baseMsgCreateAccount } as MsgCreateAccount;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (
      object.accAddressString !== undefined &&
      object.accAddressString !== null
    ) {
      message.accAddressString = object.accAddressString;
    } else {
      message.accAddressString = "";
    }
    if (object.pubKeyString !== undefined && object.pubKeyString !== null) {
      message.pubKeyString = object.pubKeyString;
    } else {
      message.pubKeyString = "";
    }
    return message;
  },
};

const baseMsgCreateAccountResponse: object = { accountId: "" };

export const MsgCreateAccountResponse = {
  encode(
    message: MsgCreateAccountResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.accountId !== "") {
      writer.uint32(10).string(message.accountId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateAccountResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateAccountResponse,
    } as MsgCreateAccountResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.accountId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateAccountResponse {
    const message = {
      ...baseMsgCreateAccountResponse,
    } as MsgCreateAccountResponse;
    if (object.accountId !== undefined && object.accountId !== null) {
      message.accountId = String(object.accountId);
    } else {
      message.accountId = "";
    }
    return message;
  },

  toJSON(message: MsgCreateAccountResponse): unknown {
    const obj: any = {};
    message.accountId !== undefined && (obj.accountId = message.accountId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateAccountResponse>
  ): MsgCreateAccountResponse {
    const message = {
      ...baseMsgCreateAccountResponse,
    } as MsgCreateAccountResponse;
    if (object.accountId !== undefined && object.accountId !== null) {
      message.accountId = object.accountId;
    } else {
      message.accountId = "";
    }
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  StoreSignature(
    request: MsgStoreSignature
  ): Promise<MsgStoreSignatureResponse>;
  PublishReferencePayloadLink(
    request: MsgPublishReferencePayloadLink
  ): Promise<MsgPublishReferencePayloadLinkResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateAccount(request: MsgCreateAccount): Promise<MsgCreateAccountResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  StoreSignature(
    request: MsgStoreSignature
  ): Promise<MsgStoreSignatureResponse> {
    const data = MsgStoreSignature.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfesignature.Msg",
      "StoreSignature",
      data
    );
    return promise.then((data) =>
      MsgStoreSignatureResponse.decode(new Reader(data))
    );
  }

  PublishReferencePayloadLink(
    request: MsgPublishReferencePayloadLink
  ): Promise<MsgPublishReferencePayloadLinkResponse> {
    const data = MsgPublishReferencePayloadLink.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfesignature.Msg",
      "PublishReferencePayloadLink",
      data
    );
    return promise.then((data) =>
      MsgPublishReferencePayloadLinkResponse.decode(new Reader(data))
    );
  }

  CreateAccount(request: MsgCreateAccount): Promise<MsgCreateAccountResponse> {
    const data = MsgCreateAccount.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfesignature.Msg",
      "CreateAccount",
      data
    );
    return promise.then((data) =>
      MsgCreateAccountResponse.decode(new Reader(data))
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
