/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { AccountVestingPools } from "./account_vesting_pool";
import { Params } from "./params";
import { VestingAccount } from "./vesting_account";

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

function createBaseGenesisState(): GenesisState {
  return {
    params: undefined,
    vestingTypes: [],
    accountVestingPools: [],
    vestingAccountList: [],
    vestingAccountCount: 0,
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.vestingTypes.push(GenesisVestingType.decode(reader, reader.uint32()));
          break;
        case 3:
          message.accountVestingPools.push(AccountVestingPools.decode(reader, reader.uint32()));
          break;
        case 4:
          message.vestingAccountList.push(VestingAccount.decode(reader, reader.uint32()));
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
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      vestingTypes: Array.isArray(object?.vestingTypes)
        ? object.vestingTypes.map((e: any) => GenesisVestingType.fromJSON(e))
        : [],
      accountVestingPools: Array.isArray(object?.accountVestingPools)
        ? object.accountVestingPools.map((e: any) => AccountVestingPools.fromJSON(e))
        : [],
      vestingAccountList: Array.isArray(object?.vestingAccountList)
        ? object.vestingAccountList.map((e: any) => VestingAccount.fromJSON(e))
        : [],
      vestingAccountCount: isSet(object.vestingAccountCount) ? Number(object.vestingAccountCount) : 0,
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.vestingTypes) {
      obj.vestingTypes = message.vestingTypes.map((e) => e ? GenesisVestingType.toJSON(e) : undefined);
    } else {
      obj.vestingTypes = [];
    }
    if (message.accountVestingPools) {
      obj.accountVestingPools = message.accountVestingPools.map((e) => e ? AccountVestingPools.toJSON(e) : undefined);
    } else {
      obj.accountVestingPools = [];
    }
    if (message.vestingAccountList) {
      obj.vestingAccountList = message.vestingAccountList.map((e) => e ? VestingAccount.toJSON(e) : undefined);
    } else {
      obj.vestingAccountList = [];
    }
    message.vestingAccountCount !== undefined && (obj.vestingAccountCount = Math.round(message.vestingAccountCount));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.vestingTypes = object.vestingTypes?.map((e) => GenesisVestingType.fromPartial(e)) || [];
    message.accountVestingPools = object.accountVestingPools?.map((e) => AccountVestingPools.fromPartial(e)) || [];
    message.vestingAccountList = object.vestingAccountList?.map((e) => VestingAccount.fromPartial(e)) || [];
    message.vestingAccountCount = object.vestingAccountCount ?? 0;
    return message;
  },
};

function createBaseGenesisVestingType(): GenesisVestingType {
  return { name: "", lockupPeriod: 0, lockupPeriodUnit: "", vestingPeriod: 0, vestingPeriodUnit: "" };
}

export const GenesisVestingType = {
  encode(message: GenesisVestingType, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisVestingType {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisVestingType();
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
    return {
      name: isSet(object.name) ? String(object.name) : "",
      lockupPeriod: isSet(object.lockupPeriod) ? Number(object.lockupPeriod) : 0,
      lockupPeriodUnit: isSet(object.lockupPeriodUnit) ? String(object.lockupPeriodUnit) : "",
      vestingPeriod: isSet(object.vestingPeriod) ? Number(object.vestingPeriod) : 0,
      vestingPeriodUnit: isSet(object.vestingPeriodUnit) ? String(object.vestingPeriodUnit) : "",
    };
  },

  toJSON(message: GenesisVestingType): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.lockupPeriod !== undefined && (obj.lockupPeriod = Math.round(message.lockupPeriod));
    message.lockupPeriodUnit !== undefined && (obj.lockupPeriodUnit = message.lockupPeriodUnit);
    message.vestingPeriod !== undefined && (obj.vestingPeriod = Math.round(message.vestingPeriod));
    message.vestingPeriodUnit !== undefined && (obj.vestingPeriodUnit = message.vestingPeriodUnit);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisVestingType>, I>>(object: I): GenesisVestingType {
    const message = createBaseGenesisVestingType();
    message.name = object.name ?? "";
    message.lockupPeriod = object.lockupPeriod ?? 0;
    message.lockupPeriodUnit = object.lockupPeriodUnit ?? "";
    message.vestingPeriod = object.vestingPeriod ?? 0;
    message.vestingPeriodUnit = object.vestingPeriodUnit ?? "";
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
