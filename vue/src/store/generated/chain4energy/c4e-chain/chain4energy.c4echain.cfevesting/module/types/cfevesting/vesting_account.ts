/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";
import { BaseVestingAccount } from "../cosmos/vesting/v1beta1/vesting";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface VestingAccount {
  id: number;
  address: string;
}

/** ContinuousVestingPeriod defines a length of time and amount of coins that will vest. */
export interface ContinuousVestingPeriod {
  start_time: number;
  end_time: number;
  amount: Coin[];
}

/**
 * RepeatedContinuousVestingAccount implements the VestingAccount interface. It
 * periodically vests by unlocking coins during each specified period.
 */
export interface RepeatedContinuousVestingAccount {
  base_vesting_account: BaseVestingAccount | undefined;
  start_time: number;
  vesting_periods: ContinuousVestingPeriod[];
}

const baseVestingAccount: object = { id: 0, address: "" };

export const VestingAccount = {
  encode(message: VestingAccount, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VestingAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVestingAccount } as VestingAccount;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VestingAccount {
    const message = { ...baseVestingAccount } as VestingAccount;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: VestingAccount): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(object: DeepPartial<VestingAccount>): VestingAccount {
    const message = { ...baseVestingAccount } as VestingAccount;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseContinuousVestingPeriod: object = { start_time: 0, end_time: 0 };

export const ContinuousVestingPeriod = {
  encode(
    message: ContinuousVestingPeriod,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.start_time !== 0) {
      writer.uint32(8).int64(message.start_time);
    }
    if (message.end_time !== 0) {
      writer.uint32(16).int64(message.end_time);
    }
    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ContinuousVestingPeriod {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseContinuousVestingPeriod,
    } as ContinuousVestingPeriod;
    message.amount = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.start_time = longToNumber(reader.int64() as Long);
          break;
        case 2:
          message.end_time = longToNumber(reader.int64() as Long);
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
    const message = {
      ...baseContinuousVestingPeriod,
    } as ContinuousVestingPeriod;
    message.amount = [];
    if (object.start_time !== undefined && object.start_time !== null) {
      message.start_time = Number(object.start_time);
    } else {
      message.start_time = 0;
    }
    if (object.end_time !== undefined && object.end_time !== null) {
      message.end_time = Number(object.end_time);
    } else {
      message.end_time = 0;
    }
    if (object.amount !== undefined && object.amount !== null) {
      for (const e of object.amount) {
        message.amount.push(Coin.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: ContinuousVestingPeriod): unknown {
    const obj: any = {};
    message.start_time !== undefined && (obj.start_time = message.start_time);
    message.end_time !== undefined && (obj.end_time = message.end_time);
    if (message.amount) {
      obj.amount = message.amount.map((e) => (e ? Coin.toJSON(e) : undefined));
    } else {
      obj.amount = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<ContinuousVestingPeriod>
  ): ContinuousVestingPeriod {
    const message = {
      ...baseContinuousVestingPeriod,
    } as ContinuousVestingPeriod;
    message.amount = [];
    if (object.start_time !== undefined && object.start_time !== null) {
      message.start_time = object.start_time;
    } else {
      message.start_time = 0;
    }
    if (object.end_time !== undefined && object.end_time !== null) {
      message.end_time = object.end_time;
    } else {
      message.end_time = 0;
    }
    if (object.amount !== undefined && object.amount !== null) {
      for (const e of object.amount) {
        message.amount.push(Coin.fromPartial(e));
      }
    }
    return message;
  },
};

const baseRepeatedContinuousVestingAccount: object = { start_time: 0 };

export const RepeatedContinuousVestingAccount = {
  encode(
    message: RepeatedContinuousVestingAccount,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.base_vesting_account !== undefined) {
      BaseVestingAccount.encode(
        message.base_vesting_account,
        writer.uint32(10).fork()
      ).ldelim();
    }
    if (message.start_time !== 0) {
      writer.uint32(16).int64(message.start_time);
    }
    for (const v of message.vesting_periods) {
      ContinuousVestingPeriod.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): RepeatedContinuousVestingAccount {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseRepeatedContinuousVestingAccount,
    } as RepeatedContinuousVestingAccount;
    message.vesting_periods = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.base_vesting_account = BaseVestingAccount.decode(
            reader,
            reader.uint32()
          );
          break;
        case 2:
          message.start_time = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.vesting_periods.push(
            ContinuousVestingPeriod.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RepeatedContinuousVestingAccount {
    const message = {
      ...baseRepeatedContinuousVestingAccount,
    } as RepeatedContinuousVestingAccount;
    message.vesting_periods = [];
    if (
      object.base_vesting_account !== undefined &&
      object.base_vesting_account !== null
    ) {
      message.base_vesting_account = BaseVestingAccount.fromJSON(
        object.base_vesting_account
      );
    } else {
      message.base_vesting_account = undefined;
    }
    if (object.start_time !== undefined && object.start_time !== null) {
      message.start_time = Number(object.start_time);
    } else {
      message.start_time = 0;
    }
    if (
      object.vesting_periods !== undefined &&
      object.vesting_periods !== null
    ) {
      for (const e of object.vesting_periods) {
        message.vesting_periods.push(ContinuousVestingPeriod.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: RepeatedContinuousVestingAccount): unknown {
    const obj: any = {};
    message.base_vesting_account !== undefined &&
      (obj.base_vesting_account = message.base_vesting_account
        ? BaseVestingAccount.toJSON(message.base_vesting_account)
        : undefined);
    message.start_time !== undefined && (obj.start_time = message.start_time);
    if (message.vesting_periods) {
      obj.vesting_periods = message.vesting_periods.map((e) =>
        e ? ContinuousVestingPeriod.toJSON(e) : undefined
      );
    } else {
      obj.vesting_periods = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<RepeatedContinuousVestingAccount>
  ): RepeatedContinuousVestingAccount {
    const message = {
      ...baseRepeatedContinuousVestingAccount,
    } as RepeatedContinuousVestingAccount;
    message.vesting_periods = [];
    if (
      object.base_vesting_account !== undefined &&
      object.base_vesting_account !== null
    ) {
      message.base_vesting_account = BaseVestingAccount.fromPartial(
        object.base_vesting_account
      );
    } else {
      message.base_vesting_account = undefined;
    }
    if (object.start_time !== undefined && object.start_time !== null) {
      message.start_time = object.start_time;
    } else {
      message.start_time = 0;
    }
    if (
      object.vesting_periods !== undefined &&
      object.vesting_periods !== null
    ) {
      for (const e of object.vesting_periods) {
        message.vesting_periods.push(ContinuousVestingPeriod.fromPartial(e));
      }
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
