/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../cfevesting/params";
import { AccountVestingPools } from "../cfevesting/account_vesting_pool";
import { VestingAccount } from "../cfevesting/vesting_account";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

/** GenesisState defines the cfevesting module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  vesting_types: GenesisVestingType[];
  account_vesting_pools: AccountVestingPools[];
  vesting_account_list: VestingAccount[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  vesting_account_count: number;
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
  /**
   * units to select:
   * days
   * hours
   * minutes
   * seconds
   */
  initial_bonus: string;
}

const baseGenesisState: object = { vesting_account_count: 0 };

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.vesting_types) {
      GenesisVestingType.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.account_vesting_pools) {
      AccountVestingPools.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.vesting_account_list) {
      VestingAccount.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.vesting_account_count !== 0) {
      writer.uint32(40).uint64(message.vesting_account_count);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.vesting_types = [];
    message.account_vesting_pools = [];
    message.vesting_account_list = [];
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
          message.account_vesting_pools.push(
            AccountVestingPools.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.vesting_account_list.push(
            VestingAccount.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.vesting_account_count = longToNumber(reader.uint64() as Long);
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
    message.account_vesting_pools = [];
    message.vesting_account_list = [];
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
      object.account_vesting_pools !== undefined &&
      object.account_vesting_pools !== null
    ) {
      for (const e of object.account_vesting_pools) {
        message.account_vesting_pools.push(AccountVestingPools.fromJSON(e));
      }
    }
    if (
      object.vesting_account_list !== undefined &&
      object.vesting_account_list !== null
    ) {
      for (const e of object.vesting_account_list) {
        message.vesting_account_list.push(VestingAccount.fromJSON(e));
      }
    }
    if (
      object.vesting_account_count !== undefined &&
      object.vesting_account_count !== null
    ) {
      message.vesting_account_count = Number(object.vesting_account_count);
    } else {
      message.vesting_account_count = 0;
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
    if (message.account_vesting_pools) {
      obj.account_vesting_pools = message.account_vesting_pools.map((e) =>
        e ? AccountVestingPools.toJSON(e) : undefined
      );
    } else {
      obj.account_vesting_pools = [];
    }
    if (message.vesting_account_list) {
      obj.vesting_account_list = message.vesting_account_list.map((e) =>
        e ? VestingAccount.toJSON(e) : undefined
      );
    } else {
      obj.vesting_account_list = [];
    }
    message.vesting_account_count !== undefined &&
      (obj.vesting_account_count = message.vesting_account_count);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.vesting_types = [];
    message.account_vesting_pools = [];
    message.vesting_account_list = [];
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
      object.account_vesting_pools !== undefined &&
      object.account_vesting_pools !== null
    ) {
      for (const e of object.account_vesting_pools) {
        message.account_vesting_pools.push(AccountVestingPools.fromPartial(e));
      }
    }
    if (
      object.vesting_account_list !== undefined &&
      object.vesting_account_list !== null
    ) {
      for (const e of object.vesting_account_list) {
        message.vesting_account_list.push(VestingAccount.fromPartial(e));
      }
    }
    if (
      object.vesting_account_count !== undefined &&
      object.vesting_account_count !== null
    ) {
      message.vesting_account_count = object.vesting_account_count;
    } else {
      message.vesting_account_count = 0;
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
  initial_bonus: "",
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
    if (message.initial_bonus !== "") {
      writer.uint32(50).string(message.initial_bonus);
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
        case 6:
          message.initial_bonus = reader.string();
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
    if (object.initial_bonus !== undefined && object.initial_bonus !== null) {
      message.initial_bonus = String(object.initial_bonus);
    } else {
      message.initial_bonus = "";
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
    message.initial_bonus !== undefined &&
      (obj.initial_bonus = message.initial_bonus);
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
    if (object.initial_bonus !== undefined && object.initial_bonus !== null) {
      message.initial_bonus = object.initial_bonus;
    } else {
      message.initial_bonus = "";
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
