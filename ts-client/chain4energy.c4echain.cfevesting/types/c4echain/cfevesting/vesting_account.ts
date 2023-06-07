/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";
import { BaseVestingAccount } from "../../cosmos/vesting/v1beta1/vesting";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface VestingAccountTrace {
  id: number;
  address: string;
  periodsToTrace: number[];
  genesis: boolean;
  fromGenesisPool: boolean;
  fromGenesisAccount: boolean;
}

/** ContinuousVestingPeriod defines a length of time and amount of coins that will vest. */
export interface ContinuousVestingPeriod {
  startTime: number;
  endTime: number;
  amount: Coin[];
}

/**
 * PeriodicContinuousVestingAccount implements the VestingAccount interface. It
 * periodically vests by unlocking coins during each specified period.
 */
export interface PeriodicContinuousVestingAccount {
  baseVestingAccount: BaseVestingAccount | undefined;
  startTime: number;
  vestingPeriods: ContinuousVestingPeriod[];
}

function createBaseVestingAccountTrace(): VestingAccountTrace {
  return { id: 0, address: "", periodsToTrace: [], genesis: false, fromGenesisPool: false, fromGenesisAccount: false };
}

export const VestingAccountTrace = {
  encode(message: VestingAccountTrace, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    writer.uint32(26).fork();
    for (const v of message.periodsToTrace) {
      writer.uint64(v);
    }
    writer.ldelim();
    if (message.genesis === true) {
      writer.uint32(32).bool(message.genesis);
    }
    if (message.fromGenesisPool === true) {
      writer.uint32(40).bool(message.fromGenesisPool);
    }
    if (message.fromGenesisAccount === true) {
      writer.uint32(48).bool(message.fromGenesisAccount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VestingAccountTrace {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVestingAccountTrace();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.periodsToTrace.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.periodsToTrace.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 4:
          message.genesis = reader.bool();
          break;
        case 5:
          message.fromGenesisPool = reader.bool();
          break;
        case 6:
          message.fromGenesisAccount = reader.bool();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VestingAccountTrace {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      address: isSet(object.address) ? String(object.address) : "",
      periodsToTrace: Array.isArray(object?.periodsToTrace) ? object.periodsToTrace.map((e: any) => Number(e)) : [],
      genesis: isSet(object.genesis) ? Boolean(object.genesis) : false,
      fromGenesisPool: isSet(object.fromGenesisPool) ? Boolean(object.fromGenesisPool) : false,
      fromGenesisAccount: isSet(object.fromGenesisAccount) ? Boolean(object.fromGenesisAccount) : false,
    };
  },

  toJSON(message: VestingAccountTrace): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.address !== undefined && (obj.address = message.address);
    if (message.periodsToTrace) {
      obj.periodsToTrace = message.periodsToTrace.map((e) => Math.round(e));
    } else {
      obj.periodsToTrace = [];
    }
    message.genesis !== undefined && (obj.genesis = message.genesis);
    message.fromGenesisPool !== undefined && (obj.fromGenesisPool = message.fromGenesisPool);
    message.fromGenesisAccount !== undefined && (obj.fromGenesisAccount = message.fromGenesisAccount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VestingAccountTrace>, I>>(object: I): VestingAccountTrace {
    const message = createBaseVestingAccountTrace();
    message.id = object.id ?? 0;
    message.address = object.address ?? "";
    message.periodsToTrace = object.periodsToTrace?.map((e) => e) || [];
    message.genesis = object.genesis ?? false;
    message.fromGenesisPool = object.fromGenesisPool ?? false;
    message.fromGenesisAccount = object.fromGenesisAccount ?? false;
    return message;
  },
};

function createBaseContinuousVestingPeriod(): ContinuousVestingPeriod {
  return { startTime: 0, endTime: 0, amount: [] };
}

export const ContinuousVestingPeriod = {
  encode(message: ContinuousVestingPeriod, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.startTime !== 0) {
      writer.uint32(8).int64(message.startTime);
    }
    if (message.endTime !== 0) {
      writer.uint32(16).int64(message.endTime);
    }
    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ContinuousVestingPeriod {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseContinuousVestingPeriod();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.startTime = longToNumber(reader.int64() as Long);
          break;
        case 2:
          message.endTime = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ContinuousVestingPeriod {
    return {
      startTime: isSet(object.startTime) ? Number(object.startTime) : 0,
      endTime: isSet(object.endTime) ? Number(object.endTime) : 0,
      amount: Array.isArray(object?.amount) ? object.amount.map((e: any) => Coin.fromJSON(e)) : [],
    };
  },

  toJSON(message: ContinuousVestingPeriod): unknown {
    const obj: any = {};
    message.startTime !== undefined && (obj.startTime = Math.round(message.startTime));
    message.endTime !== undefined && (obj.endTime = Math.round(message.endTime));
    if (message.amount) {
      obj.amount = message.amount.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.amount = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ContinuousVestingPeriod>, I>>(object: I): ContinuousVestingPeriod {
    const message = createBaseContinuousVestingPeriod();
    message.startTime = object.startTime ?? 0;
    message.endTime = object.endTime ?? 0;
    message.amount = object.amount?.map((e) => Coin.fromPartial(e)) || [];
    return message;
  },
};

function createBasePeriodicContinuousVestingAccount(): PeriodicContinuousVestingAccount {
  return { baseVestingAccount: undefined, startTime: 0, vestingPeriods: [] };
}

export const PeriodicContinuousVestingAccount = {
  encode(message: PeriodicContinuousVestingAccount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.baseVestingAccount !== undefined) {
      BaseVestingAccount.encode(message.baseVestingAccount, writer.uint32(10).fork()).ldelim();
    }
    if (message.startTime !== 0) {
      writer.uint32(16).int64(message.startTime);
    }
    for (const v of message.vestingPeriods) {
      ContinuousVestingPeriod.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PeriodicContinuousVestingAccount {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePeriodicContinuousVestingAccount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.baseVestingAccount = BaseVestingAccount.decode(reader, reader.uint32());
          break;
        case 2:
          message.startTime = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.vestingPeriods.push(ContinuousVestingPeriod.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PeriodicContinuousVestingAccount {
    return {
      baseVestingAccount: isSet(object.baseVestingAccount)
        ? BaseVestingAccount.fromJSON(object.baseVestingAccount)
        : undefined,
      startTime: isSet(object.startTime) ? Number(object.startTime) : 0,
      vestingPeriods: Array.isArray(object?.vestingPeriods)
        ? object.vestingPeriods.map((e: any) => ContinuousVestingPeriod.fromJSON(e))
        : [],
    };
  },

  toJSON(message: PeriodicContinuousVestingAccount): unknown {
    const obj: any = {};
    message.baseVestingAccount !== undefined && (obj.baseVestingAccount = message.baseVestingAccount
      ? BaseVestingAccount.toJSON(message.baseVestingAccount)
      : undefined);
    message.startTime !== undefined && (obj.startTime = Math.round(message.startTime));
    if (message.vestingPeriods) {
      obj.vestingPeriods = message.vestingPeriods.map((e) => e ? ContinuousVestingPeriod.toJSON(e) : undefined);
    } else {
      obj.vestingPeriods = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PeriodicContinuousVestingAccount>, I>>(
    object: I,
  ): PeriodicContinuousVestingAccount {
    const message = createBasePeriodicContinuousVestingAccount();
    message.baseVestingAccount = (object.baseVestingAccount !== undefined && object.baseVestingAccount !== null)
      ? BaseVestingAccount.fromPartial(object.baseVestingAccount)
      : undefined;
    message.startTime = object.startTime ?? 0;
    message.vestingPeriods = object.vestingPeriods?.map((e) => ContinuousVestingPeriod.fromPartial(e)) || [];
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
