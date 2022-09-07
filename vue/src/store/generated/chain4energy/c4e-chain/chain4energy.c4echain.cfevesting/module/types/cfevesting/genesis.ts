/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../cfevesting/params";
import { AccountVestingsList } from "../cfevesting/account_vesting";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

/** GenesisState defines the cfevesting module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  /** this line is used by starport scaffolding # genesis/proto/state */
  vesting_types: GenesisVestingType[];
  account_vestings_list: AccountVestingsList | undefined;
}

export interface GenesisVestingType {
  /** vesting type name */
  name: string;
  /** period of locked coins from vesting start */
  lockup_period: number;
  lockup_period_unit: string;
  /** period of veesting coins from lockup period end */
  vesting_period: number;
  vesting_period_unit: string;
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.vesting_types) {
      GenesisVestingType.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.account_vestings_list !== undefined) {
      AccountVestingsList.encode(
        message.account_vestings_list,
        writer.uint32(26).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.vesting_types = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.vesting_types.push(
            GenesisVestingType.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.account_vestings_list = AccountVestingsList.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.vesting_types = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.vesting_types !== undefined && object.vesting_types !== null) {
      for (const e of object.vesting_types) {
        message.vesting_types.push(GenesisVestingType.fromJSON(e));
      }
    }
    if (
      object.account_vestings_list !== undefined &&
      object.account_vestings_list !== null
    ) {
      message.account_vestings_list = AccountVestingsList.fromJSON(
        object.account_vestings_list
      );
    } else {
      message.account_vestings_list = undefined;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.vesting_types) {
      obj.vesting_types = message.vesting_types.map((e) =>
        e ? GenesisVestingType.toJSON(e) : undefined
      );
    } else {
      obj.vesting_types = [];
    }
    message.account_vestings_list !== undefined &&
      (obj.account_vestings_list = message.account_vestings_list
        ? AccountVestingsList.toJSON(message.account_vestings_list)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.vesting_types = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.vesting_types !== undefined && object.vesting_types !== null) {
      for (const e of object.vesting_types) {
        message.vesting_types.push(GenesisVestingType.fromPartial(e));
      }
    }
    if (
      object.account_vestings_list !== undefined &&
      object.account_vestings_list !== null
    ) {
      message.account_vestings_list = AccountVestingsList.fromPartial(
        object.account_vestings_list
      );
    } else {
      message.account_vestings_list = undefined;
    }
    return message;
  },
};

const baseGenesisVestingType: object = {
  name: "",
  lockup_period: 0,
  lockup_period_unit: "",
  vesting_period: 0,
  vesting_period_unit: "",
};

export const GenesisVestingType = {
  encode(
    message: GenesisVestingType,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.lockup_period !== 0) {
      writer.uint32(16).int64(message.lockup_period);
    }
    if (message.lockup_period_unit !== "") {
      writer.uint32(26).string(message.lockup_period_unit);
    }
    if (message.vesting_period !== 0) {
      writer.uint32(32).int64(message.vesting_period);
    }
    if (message.vesting_period_unit !== "") {
      writer.uint32(42).string(message.vesting_period_unit);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisVestingType {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisVestingType } as GenesisVestingType;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.lockup_period = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.lockup_period_unit = reader.string();
          break;
        case 4:
          message.vesting_period = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.vesting_period_unit = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisVestingType {
    const message = { ...baseGenesisVestingType } as GenesisVestingType;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.lockup_period !== undefined && object.lockup_period !== null) {
      message.lockup_period = Number(object.lockup_period);
    } else {
      message.lockup_period = 0;
    }
    if (
      object.lockup_period_unit !== undefined &&
      object.lockup_period_unit !== null
    ) {
      message.lockup_period_unit = String(object.lockup_period_unit);
    } else {
      message.lockup_period_unit = "";
    }
    if (object.vesting_period !== undefined && object.vesting_period !== null) {
      message.vesting_period = Number(object.vesting_period);
    } else {
      message.vesting_period = 0;
    }
    if (
      object.vesting_period_unit !== undefined &&
      object.vesting_period_unit !== null
    ) {
      message.vesting_period_unit = String(object.vesting_period_unit);
    } else {
      message.vesting_period_unit = "";
    }
    return message;
  },

  toJSON(message: GenesisVestingType): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.lockup_period !== undefined &&
      (obj.lockup_period = message.lockup_period);
    message.lockup_period_unit !== undefined &&
      (obj.lockup_period_unit = message.lockup_period_unit);
    message.vesting_period !== undefined &&
      (obj.vesting_period = message.vesting_period);
    message.vesting_period_unit !== undefined &&
      (obj.vesting_period_unit = message.vesting_period_unit);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisVestingType>): GenesisVestingType {
    const message = { ...baseGenesisVestingType } as GenesisVestingType;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.lockup_period !== undefined && object.lockup_period !== null) {
      message.lockup_period = object.lockup_period;
    } else {
      message.lockup_period = 0;
    }
    if (
      object.lockup_period_unit !== undefined &&
      object.lockup_period_unit !== null
    ) {
      message.lockup_period_unit = object.lockup_period_unit;
    } else {
      message.lockup_period_unit = "";
    }
    if (object.vesting_period !== undefined && object.vesting_period !== null) {
      message.vesting_period = object.vesting_period;
    } else {
      message.vesting_period = 0;
    }
    if (
      object.vesting_period_unit !== undefined &&
      object.vesting_period_unit !== null
    ) {
      message.vesting_period_unit = object.vesting_period_unit;
    } else {
      message.vesting_period_unit = "";
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
