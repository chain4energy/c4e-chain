/* eslint-disable */
import { Duration } from "../google/protobuf/duration";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface VestingTypes {
  vestingTypes: VestingType[];
}

export interface VestingType {
  /** vesting type name */
  name: string;
  /** period of locked coins (minutes) from vesting start */
  lockupPeriod: Duration | undefined;
  /** period of veesting coins (minutes) from lockup period end */
  vestingPeriod: Duration | undefined;
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

const baseVestingType: object = { name: "" };

export const VestingType = {
  encode(message: VestingType, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.lockupPeriod !== undefined) {
      Duration.encode(message.lockupPeriod, writer.uint32(18).fork()).ldelim();
    }
    if (message.vestingPeriod !== undefined) {
      Duration.encode(message.vestingPeriod, writer.uint32(26).fork()).ldelim();
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
          message.lockupPeriod = Duration.decode(reader, reader.uint32());
          break;
        case 3:
          message.vestingPeriod = Duration.decode(reader, reader.uint32());
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
      message.lockupPeriod = Duration.fromJSON(object.lockupPeriod);
    } else {
      message.lockupPeriod = undefined;
    }
    if (object.vestingPeriod !== undefined && object.vestingPeriod !== null) {
      message.vestingPeriod = Duration.fromJSON(object.vestingPeriod);
    } else {
      message.vestingPeriod = undefined;
    }
    return message;
  },

  toJSON(message: VestingType): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.lockupPeriod !== undefined &&
      (obj.lockupPeriod = message.lockupPeriod
        ? Duration.toJSON(message.lockupPeriod)
        : undefined);
    message.vestingPeriod !== undefined &&
      (obj.vestingPeriod = message.vestingPeriod
        ? Duration.toJSON(message.vestingPeriod)
        : undefined);
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
      message.lockupPeriod = Duration.fromPartial(object.lockupPeriod);
    } else {
      message.lockupPeriod = undefined;
    }
    if (object.vestingPeriod !== undefined && object.vestingPeriod !== null) {
      message.vestingPeriod = Duration.fromPartial(object.vestingPeriod);
    } else {
      message.vestingPeriod = undefined;
    }
    return message;
  },
};

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
