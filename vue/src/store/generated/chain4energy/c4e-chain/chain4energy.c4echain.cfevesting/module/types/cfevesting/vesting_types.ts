/* eslint-disable */
import { Duration } from "../google/protobuf/duration";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface VestingTypes {
  vesting_types: VestingType[];
}

export interface VestingType {
  /** vesting type name */
  name: string;
  /** period of locked coins (minutes) from vesting start */
  lockup_period: Duration | undefined;
  /** period of vesting coins (minutes) from lockup period end */
  vesting_period: Duration | undefined;
}

const baseVestingTypes: object = {};

export const VestingTypes = {
  encode(message: VestingTypes, writer: Writer = Writer.create()): Writer {
    for (const v of message.vesting_types) {
      VestingType.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VestingTypes {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVestingTypes } as VestingTypes;
    message.vesting_types = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vesting_types.push(
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
    message.vesting_types = [];
    if (object.vesting_types !== undefined && object.vesting_types !== null) {
      for (const e of object.vesting_types) {
        message.vesting_types.push(VestingType.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: VestingTypes): unknown {
    const obj: any = {};
    if (message.vesting_types) {
      obj.vesting_types = message.vesting_types.map((e) =>
        e ? VestingType.toJSON(e) : undefined
      );
    } else {
      obj.vesting_types = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<VestingTypes>): VestingTypes {
    const message = { ...baseVestingTypes } as VestingTypes;
    message.vesting_types = [];
    if (object.vesting_types !== undefined && object.vesting_types !== null) {
      for (const e of object.vesting_types) {
        message.vesting_types.push(VestingType.fromPartial(e));
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
    if (message.lockup_period !== undefined) {
      Duration.encode(message.lockup_period, writer.uint32(18).fork()).ldelim();
    }
    if (message.vesting_period !== undefined) {
      Duration.encode(
        message.vesting_period,
        writer.uint32(26).fork()
      ).ldelim();
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
          message.lockup_period = Duration.decode(reader, reader.uint32());
          break;
        case 3:
          message.vesting_period = Duration.decode(reader, reader.uint32());
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
    if (object.lockup_period !== undefined && object.lockup_period !== null) {
      message.lockup_period = Duration.fromJSON(object.lockup_period);
    } else {
      message.lockup_period = undefined;
    }
    if (object.vesting_period !== undefined && object.vesting_period !== null) {
      message.vesting_period = Duration.fromJSON(object.vesting_period);
    } else {
      message.vesting_period = undefined;
    }
    return message;
  },

  toJSON(message: VestingType): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.lockup_period !== undefined &&
      (obj.lockup_period = message.lockup_period
        ? Duration.toJSON(message.lockup_period)
        : undefined);
    message.vesting_period !== undefined &&
      (obj.vesting_period = message.vesting_period
        ? Duration.toJSON(message.vesting_period)
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
    if (object.lockup_period !== undefined && object.lockup_period !== null) {
      message.lockup_period = Duration.fromPartial(object.lockup_period);
    } else {
      message.lockup_period = undefined;
    }
    if (object.vesting_period !== undefined && object.vesting_period !== null) {
      message.vesting_period = Duration.fromPartial(object.vesting_period);
    } else {
      message.vesting_period = undefined;
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
