/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.energybank";

export interface TokenParams {
  index: string;
  name: string;
  tradingCompany: string;
  burningTime: number;
  burningType: string;
  sendPrice: number;
}

const baseTokenParams: object = {
  index: "",
  name: "",
  tradingCompany: "",
  burningTime: 0,
  burningType: "",
  sendPrice: 0,
};

export const TokenParams = {
  encode(message: TokenParams, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.tradingCompany !== "") {
      writer.uint32(26).string(message.tradingCompany);
    }
    if (message.burningTime !== 0) {
      writer.uint32(32).uint64(message.burningTime);
    }
    if (message.burningType !== "") {
      writer.uint32(42).string(message.burningType);
    }
    if (message.sendPrice !== 0) {
      writer.uint32(48).uint64(message.sendPrice);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): TokenParams {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTokenParams } as TokenParams;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.tradingCompany = reader.string();
          break;
        case 4:
          message.burningTime = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.burningType = reader.string();
          break;
        case 6:
          message.sendPrice = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TokenParams {
    const message = { ...baseTokenParams } as TokenParams;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.tradingCompany !== undefined && object.tradingCompany !== null) {
      message.tradingCompany = String(object.tradingCompany);
    } else {
      message.tradingCompany = "";
    }
    if (object.burningTime !== undefined && object.burningTime !== null) {
      message.burningTime = Number(object.burningTime);
    } else {
      message.burningTime = 0;
    }
    if (object.burningType !== undefined && object.burningType !== null) {
      message.burningType = String(object.burningType);
    } else {
      message.burningType = "";
    }
    if (object.sendPrice !== undefined && object.sendPrice !== null) {
      message.sendPrice = Number(object.sendPrice);
    } else {
      message.sendPrice = 0;
    }
    return message;
  },

  toJSON(message: TokenParams): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.name !== undefined && (obj.name = message.name);
    message.tradingCompany !== undefined &&
      (obj.tradingCompany = message.tradingCompany);
    message.burningTime !== undefined &&
      (obj.burningTime = message.burningTime);
    message.burningType !== undefined &&
      (obj.burningType = message.burningType);
    message.sendPrice !== undefined && (obj.sendPrice = message.sendPrice);
    return obj;
  },

  fromPartial(object: DeepPartial<TokenParams>): TokenParams {
    const message = { ...baseTokenParams } as TokenParams;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.tradingCompany !== undefined && object.tradingCompany !== null) {
      message.tradingCompany = object.tradingCompany;
    } else {
      message.tradingCompany = "";
    }
    if (object.burningTime !== undefined && object.burningTime !== null) {
      message.burningTime = object.burningTime;
    } else {
      message.burningTime = 0;
    }
    if (object.burningType !== undefined && object.burningType !== null) {
      message.burningType = object.burningType;
    } else {
      message.burningType = "";
    }
    if (object.sendPrice !== undefined && object.sendPrice !== null) {
      message.sendPrice = object.sendPrice;
    } else {
      message.sendPrice = 0;
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
