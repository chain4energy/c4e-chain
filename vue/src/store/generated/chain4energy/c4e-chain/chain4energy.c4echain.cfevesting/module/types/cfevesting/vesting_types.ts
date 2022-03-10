/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface VestingTypes {
  vestingTypes: VestingType[];
}

export interface VestingType {
  name: string;
  lockupPeriod: number;
  vestingPeriod: number;
  tokenReleasingPeriod: number;
  delegationsAllowed: boolean;
}

const baseVestingTypes: object = {};

export const VestingTypes = {
  encode(message: VestingTypes, writer: Writer = Writer.create()): Writer {
    for (const v of message.vestingTypes) {
      VestingType.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VestingTypes {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVestingTypes } as VestingTypes;
    message.vestingTypes = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vestingTypes.push(
            VestingType.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VestingTypes {
    const message = { ...baseVestingTypes } as VestingTypes;
    message.vestingTypes = [];
    if (object.vestingTypes !== undefined && object.vestingTypes !== null) {
      for (const e of object.vestingTypes) {
        message.vestingTypes.push(VestingType.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: VestingTypes): unknown {
    const obj: any = {};
    if (message.vestingTypes) {
      obj.vestingTypes = message.vestingTypes.map((e) =>
        e ? VestingType.toJSON(e) : undefined
      );
    } else {
      obj.vestingTypes = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<VestingTypes>): VestingTypes {
    const message = { ...baseVestingTypes } as VestingTypes;
    message.vestingTypes = [];
    if (object.vestingTypes !== undefined && object.vestingTypes !== null) {
      for (const e of object.vestingTypes) {
        message.vestingTypes.push(VestingType.fromPartial(e));
      }
    }
    return message;
  },
};

const baseVestingType: object = {
  name: "",
  lockupPeriod: 0,
  vestingPeriod: 0,
  tokenReleasingPeriod: 0,
  delegationsAllowed: false,
};

export const VestingType = {
  encode(message: VestingType, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.lockupPeriod !== 0) {
      writer.uint32(16).int64(message.lockupPeriod);
    }
    if (message.vestingPeriod !== 0) {
      writer.uint32(24).int64(message.vestingPeriod);
    }
    if (message.tokenReleasingPeriod !== 0) {
      writer.uint32(32).int64(message.tokenReleasingPeriod);
    }
    if (message.delegationsAllowed === true) {
      writer.uint32(40).bool(message.delegationsAllowed);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VestingType {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVestingType } as VestingType;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.lockupPeriod = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.vestingPeriod = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.tokenReleasingPeriod = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.delegationsAllowed = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VestingType {
    const message = { ...baseVestingType } as VestingType;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.lockupPeriod !== undefined && object.lockupPeriod !== null) {
      message.lockupPeriod = Number(object.lockupPeriod);
    } else {
      message.lockupPeriod = 0;
    }
    if (object.vestingPeriod !== undefined && object.vestingPeriod !== null) {
      message.vestingPeriod = Number(object.vestingPeriod);
    } else {
      message.vestingPeriod = 0;
    }
    if (
      object.tokenReleasingPeriod !== undefined &&
      object.tokenReleasingPeriod !== null
    ) {
      message.tokenReleasingPeriod = Number(object.tokenReleasingPeriod);
    } else {
      message.tokenReleasingPeriod = 0;
    }
    if (
      object.delegationsAllowed !== undefined &&
      object.delegationsAllowed !== null
    ) {
      message.delegationsAllowed = Boolean(object.delegationsAllowed);
    } else {
      message.delegationsAllowed = false;
    }
    return message;
  },

  toJSON(message: VestingType): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.lockupPeriod !== undefined &&
      (obj.lockupPeriod = message.lockupPeriod);
    message.vestingPeriod !== undefined &&
      (obj.vestingPeriod = message.vestingPeriod);
    message.tokenReleasingPeriod !== undefined &&
      (obj.tokenReleasingPeriod = message.tokenReleasingPeriod);
    message.delegationsAllowed !== undefined &&
      (obj.delegationsAllowed = message.delegationsAllowed);
    return obj;
  },

  fromPartial(object: DeepPartial<VestingType>): VestingType {
    const message = { ...baseVestingType } as VestingType;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.lockupPeriod !== undefined && object.lockupPeriod !== null) {
      message.lockupPeriod = object.lockupPeriod;
    } else {
      message.lockupPeriod = 0;
    }
    if (object.vestingPeriod !== undefined && object.vestingPeriod !== null) {
      message.vestingPeriod = object.vestingPeriod;
    } else {
      message.vestingPeriod = 0;
    }
    if (
      object.tokenReleasingPeriod !== undefined &&
      object.tokenReleasingPeriod !== null
    ) {
      message.tokenReleasingPeriod = object.tokenReleasingPeriod;
    } else {
      message.tokenReleasingPeriod = 0;
    }
    if (
      object.delegationsAllowed !== undefined &&
      object.delegationsAllowed !== null
    ) {
      message.delegationsAllowed = object.delegationsAllowed;
    } else {
      message.delegationsAllowed = false;
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
