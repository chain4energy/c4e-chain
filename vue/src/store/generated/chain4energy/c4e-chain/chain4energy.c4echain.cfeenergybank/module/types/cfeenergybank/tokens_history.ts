/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeenergybank";

export interface TokensHistory {
  id: number;
  userAddress: string;
  createdAt: number;
  issuerAddress: string;
  targetAddress: string;
  amount: number;
  tokenName: string;
}

const baseTokensHistory: object = {
  id: 0,
  userAddress: "",
  createdAt: 0,
  issuerAddress: "",
  targetAddress: "",
  amount: 0,
  tokenName: "",
};

export const TokensHistory = {
  encode(message: TokensHistory, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.userAddress !== "") {
      writer.uint32(18).string(message.userAddress);
    }
    if (message.createdAt !== 0) {
      writer.uint32(24).uint64(message.createdAt);
    }
    if (message.issuerAddress !== "") {
      writer.uint32(34).string(message.issuerAddress);
    }
    if (message.targetAddress !== "") {
      writer.uint32(42).string(message.targetAddress);
    }
    if (message.amount !== 0) {
      writer.uint32(48).uint64(message.amount);
    }
    if (message.tokenName !== "") {
      writer.uint32(58).string(message.tokenName);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): TokensHistory {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTokensHistory } as TokensHistory;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.userAddress = reader.string();
          break;
        case 3:
          message.createdAt = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.issuerAddress = reader.string();
          break;
        case 5:
          message.targetAddress = reader.string();
          break;
        case 6:
          message.amount = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.tokenName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TokensHistory {
    const message = { ...baseTokensHistory } as TokensHistory;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.userAddress !== undefined && object.userAddress !== null) {
      message.userAddress = String(object.userAddress);
    } else {
      message.userAddress = "";
    }
    if (object.createdAt !== undefined && object.createdAt !== null) {
      message.createdAt = Number(object.createdAt);
    } else {
      message.createdAt = 0;
    }
    if (object.issuerAddress !== undefined && object.issuerAddress !== null) {
      message.issuerAddress = String(object.issuerAddress);
    } else {
      message.issuerAddress = "";
    }
    if (object.targetAddress !== undefined && object.targetAddress !== null) {
      message.targetAddress = String(object.targetAddress);
    } else {
      message.targetAddress = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = Number(object.amount);
    } else {
      message.amount = 0;
    }
    if (object.tokenName !== undefined && object.tokenName !== null) {
      message.tokenName = String(object.tokenName);
    } else {
      message.tokenName = "";
    }
    return message;
  },

  toJSON(message: TokensHistory): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.userAddress !== undefined &&
      (obj.userAddress = message.userAddress);
    message.createdAt !== undefined && (obj.createdAt = message.createdAt);
    message.issuerAddress !== undefined &&
      (obj.issuerAddress = message.issuerAddress);
    message.targetAddress !== undefined &&
      (obj.targetAddress = message.targetAddress);
    message.amount !== undefined && (obj.amount = message.amount);
    message.tokenName !== undefined && (obj.tokenName = message.tokenName);
    return obj;
  },

  fromPartial(object: DeepPartial<TokensHistory>): TokensHistory {
    const message = { ...baseTokensHistory } as TokensHistory;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.userAddress !== undefined && object.userAddress !== null) {
      message.userAddress = object.userAddress;
    } else {
      message.userAddress = "";
    }
    if (object.createdAt !== undefined && object.createdAt !== null) {
      message.createdAt = object.createdAt;
    } else {
      message.createdAt = 0;
    }
    if (object.issuerAddress !== undefined && object.issuerAddress !== null) {
      message.issuerAddress = object.issuerAddress;
    } else {
      message.issuerAddress = "";
    }
    if (object.targetAddress !== undefined && object.targetAddress !== null) {
      message.targetAddress = object.targetAddress;
    } else {
      message.targetAddress = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = 0;
    }
    if (object.tokenName !== undefined && object.tokenName !== null) {
      message.tokenName = object.tokenName;
    } else {
      message.tokenName = "";
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
