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
  vestingTypes: GenesisVestingType[];
  accountVestingPools: AccountVestingPools[];
  vestingAccountList: VestingAccount[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  vestingAccountCount: number;
}

export interface GenesisVestingType {
  /** vesting type name */
  name: string;
  /** period of locked coins from vesting start */
  lockupPeriod: number;
  lockupPeriodUnit: string;
  /** period of veesting coins from lockup period end */
  vestingPeriod: number;
  vestingPeriodUnit: string;
}

const baseGenesisState: object = { vestingAccountCount: 0 };

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.vestingTypes) {
      GenesisVestingType.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.accountVestingPools) {
      AccountVestingPools.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.vestingAccountList) {
      VestingAccount.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    if (message.vestingAccountCount !== 0) {
      writer.uint32(40).uint64(message.vestingAccountCount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.vestingTypes = [];
    message.accountVestingPools = [];
    message.vestingAccountList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.vestingTypes.push(
            GenesisVestingType.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.accountVestingPools.push(
            AccountVestingPools.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.vestingAccountList.push(
            VestingAccount.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.vestingAccountCount = longToNumber(reader.uint64() as Long);
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
    message.vestingTypes = [];
    message.accountVestingPools = [];
    message.vestingAccountList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.vestingTypes !== undefined && object.vestingTypes !== null) {
      for (const e of object.vestingTypes) {
        message.vestingTypes.push(GenesisVestingType.fromJSON(e));
      }
    }
    if (
      object.accountVestingPools !== undefined &&
      object.accountVestingPools !== null
    ) {
      for (const e of object.accountVestingPools) {
        message.accountVestingPools.push(AccountVestingPools.fromJSON(e));
      }
    }
    if (
      object.vestingAccountList !== undefined &&
      object.vestingAccountList !== null
    ) {
      for (const e of object.vestingAccountList) {
        message.vestingAccountList.push(VestingAccount.fromJSON(e));
      }
    }
    if (
      object.vestingAccountCount !== undefined &&
      object.vestingAccountCount !== null
    ) {
      message.vestingAccountCount = Number(object.vestingAccountCount);
    } else {
      message.vestingAccountCount = 0;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.vestingTypes) {
      obj.vestingTypes = message.vestingTypes.map((e) =>
        e ? GenesisVestingType.toJSON(e) : undefined
      );
    } else {
      obj.vestingTypes = [];
    }
    if (message.accountVestingPools) {
      obj.accountVestingPools = message.accountVestingPools.map((e) =>
        e ? AccountVestingPools.toJSON(e) : undefined
      );
    } else {
      obj.accountVestingPools = [];
    }
    if (message.vestingAccountList) {
      obj.vestingAccountList = message.vestingAccountList.map((e) =>
        e ? VestingAccount.toJSON(e) : undefined
      );
    } else {
      obj.vestingAccountList = [];
    }
    message.vestingAccountCount !== undefined &&
      (obj.vestingAccountCount = message.vestingAccountCount);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.vestingTypes = [];
    message.accountVestingPools = [];
    message.vestingAccountList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.vestingTypes !== undefined && object.vestingTypes !== null) {
      for (const e of object.vestingTypes) {
        message.vestingTypes.push(GenesisVestingType.fromPartial(e));
      }
    }
    if (
      object.accountVestingPools !== undefined &&
      object.accountVestingPools !== null
    ) {
      for (const e of object.accountVestingPools) {
        message.accountVestingPools.push(AccountVestingPools.fromPartial(e));
      }
    }
    if (
      object.vestingAccountList !== undefined &&
      object.vestingAccountList !== null
    ) {
      for (const e of object.vestingAccountList) {
        message.vestingAccountList.push(VestingAccount.fromPartial(e));
      }
    }
    if (
      object.vestingAccountCount !== undefined &&
      object.vestingAccountCount !== null
    ) {
      message.vestingAccountCount = object.vestingAccountCount;
    } else {
      message.vestingAccountCount = 0;
    }
    return message;
  },
};

const baseGenesisVestingType: object = {
  name: "",
  lockupPeriod: 0,
  lockupPeriodUnit: "",
  vestingPeriod: 0,
  vestingPeriodUnit: "",
};

export const GenesisVestingType = {
  encode(
    message: GenesisVestingType,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.lockupPeriod !== 0) {
      writer.uint32(16).int64(message.lockupPeriod);
    }
    if (message.lockupPeriodUnit !== "") {
      writer.uint32(26).string(message.lockupPeriodUnit);
    }
    if (message.vestingPeriod !== 0) {
      writer.uint32(32).int64(message.vestingPeriod);
    }
    if (message.vestingPeriodUnit !== "") {
      writer.uint32(42).string(message.vestingPeriodUnit);
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
          message.lockupPeriod = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.lockupPeriodUnit = reader.string();
          break;
        case 4:
          message.vestingPeriod = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.vestingPeriodUnit = reader.string();
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
    if (object.lockupPeriod !== undefined && object.lockupPeriod !== null) {
      message.lockupPeriod = Number(object.lockupPeriod);
    } else {
      message.lockupPeriod = 0;
    }
    if (
      object.lockupPeriodUnit !== undefined &&
      object.lockupPeriodUnit !== null
    ) {
      message.lockupPeriodUnit = String(object.lockupPeriodUnit);
    } else {
      message.lockupPeriodUnit = "";
    }
    if (object.vestingPeriod !== undefined && object.vestingPeriod !== null) {
      message.vestingPeriod = Number(object.vestingPeriod);
    } else {
      message.vestingPeriod = 0;
    }
    if (
      object.vestingPeriodUnit !== undefined &&
      object.vestingPeriodUnit !== null
    ) {
      message.vestingPeriodUnit = String(object.vestingPeriodUnit);
    } else {
      message.vestingPeriodUnit = "";
    }
    return message;
  },

  toJSON(message: GenesisVestingType): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.lockupPeriod !== undefined &&
      (obj.lockupPeriod = message.lockupPeriod);
    message.lockupPeriodUnit !== undefined &&
      (obj.lockupPeriodUnit = message.lockupPeriodUnit);
    message.vestingPeriod !== undefined &&
      (obj.vestingPeriod = message.vestingPeriod);
    message.vestingPeriodUnit !== undefined &&
      (obj.vestingPeriodUnit = message.vestingPeriodUnit);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisVestingType>): GenesisVestingType {
    const message = { ...baseGenesisVestingType } as GenesisVestingType;
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
    if (
      object.lockupPeriodUnit !== undefined &&
      object.lockupPeriodUnit !== null
    ) {
      message.lockupPeriodUnit = object.lockupPeriodUnit;
    } else {
      message.lockupPeriodUnit = "";
    }
    if (object.vestingPeriod !== undefined && object.vestingPeriod !== null) {
      message.vestingPeriod = object.vestingPeriod;
    } else {
      message.vestingPeriod = 0;
    }
    if (
      object.vestingPeriodUnit !== undefined &&
      object.vestingPeriodUnit !== null
    ) {
      message.vestingPeriodUnit = object.vestingPeriodUnit;
    } else {
      message.vestingPeriodUnit = "";
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
