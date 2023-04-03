/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Duration } from "../../google/protobuf/duration";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface VestingTypes {
  vestingTypes: VestingType[];
}

export interface VestingType {
  /** vesting type name */
  name: string;
  /** period of locked coins (minutes) from vesting start */
  lockupPeriod:
    | Duration
    | undefined;
  /** period of vesting coins (minutes) from lockup period end */
  vestingPeriod:
    | Duration
    | undefined;
  /** the percentage of tokens that are released initially */
  free: string;
}

function createBaseVestingTypes(): VestingTypes {
  return { vestingTypes: [] };
}

export const VestingTypes = {
  encode(message: VestingTypes, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.vestingTypes) {
      VestingType.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VestingTypes {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVestingTypes();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vestingTypes.push(VestingType.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VestingTypes {
    return {
      vestingTypes: Array.isArray(object?.vestingTypes)
        ? object.vestingTypes.map((e: any) => VestingType.fromJSON(e))
        : [],
    };
  },

  toJSON(message: VestingTypes): unknown {
    const obj: any = {};
    if (message.vestingTypes) {
      obj.vestingTypes = message.vestingTypes.map((e) => e ? VestingType.toJSON(e) : undefined);
    } else {
      obj.vestingTypes = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VestingTypes>, I>>(object: I): VestingTypes {
    const message = createBaseVestingTypes();
    message.vestingTypes = object.vestingTypes?.map((e) => VestingType.fromPartial(e)) || [];
    return message;
  },
};

function createBaseVestingType(): VestingType {
  return { name: "", lockupPeriod: undefined, vestingPeriod: undefined, free: "" };
}

export const VestingType = {
  encode(message: VestingType, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.lockupPeriod !== undefined) {
      Duration.encode(message.lockupPeriod, writer.uint32(18).fork()).ldelim();
    }
    if (message.vestingPeriod !== undefined) {
      Duration.encode(message.vestingPeriod, writer.uint32(26).fork()).ldelim();
    }
    if (message.free !== "") {
      writer.uint32(34).string(message.free);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VestingType {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVestingType();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.lockupPeriod = Duration.decode(reader, reader.uint32());
          break;
        case 3:
          message.vestingPeriod = Duration.decode(reader, reader.uint32());
          break;
        case 4:
          message.free = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VestingType {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      lockupPeriod: isSet(object.lockupPeriod) ? Duration.fromJSON(object.lockupPeriod) : undefined,
      vestingPeriod: isSet(object.vestingPeriod) ? Duration.fromJSON(object.vestingPeriod) : undefined,
      free: isSet(object.free) ? String(object.free) : "",
    };
  },

  toJSON(message: VestingType): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.lockupPeriod !== undefined
      && (obj.lockupPeriod = message.lockupPeriod ? Duration.toJSON(message.lockupPeriod) : undefined);
    message.vestingPeriod !== undefined
      && (obj.vestingPeriod = message.vestingPeriod ? Duration.toJSON(message.vestingPeriod) : undefined);
    message.free !== undefined && (obj.free = message.free);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VestingType>, I>>(object: I): VestingType {
    const message = createBaseVestingType();
    message.name = object.name ?? "";
    message.lockupPeriod = (object.lockupPeriod !== undefined && object.lockupPeriod !== null)
      ? Duration.fromPartial(object.lockupPeriod)
      : undefined;
    message.vestingPeriod = (object.vestingPeriod !== undefined && object.vestingPeriod !== null)
      ? Duration.fromPartial(object.vestingPeriod)
      : undefined;
    message.free = object.free ?? "";
    return message;
  },
};

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
