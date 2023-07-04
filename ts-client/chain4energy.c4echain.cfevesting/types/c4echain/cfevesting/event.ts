/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface EventNewVestingAccount {
  address: string;
}

export interface EventNewVestingPool {
  owner: string;
  name: string;
  amount: string;
  duration: string;
  vestingType: string;
}

export interface EventNewVestingPeriodFromVestingPool {
  owner: string;
  address: string;
  vestingPoolName: string;
  amount: string;
  restartVesting: string;
  periodId: number;
}

export interface EventWithdrawAvailable {
  owner: string;
  vestingPoolName: string;
  amount: string;
}

export interface EventVestingSplit {
  source: string;
  destination: string;
}

function createBaseEventNewVestingAccount(): EventNewVestingAccount {
  return { address: "" };
}

export const EventNewVestingAccount = {
  encode(message: EventNewVestingAccount, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EventNewVestingAccount {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventNewVestingAccount();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventNewVestingAccount {
    return { address: isSet(object.address) ? String(object.address) : "" };
  },

  toJSON(message: EventNewVestingAccount): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventNewVestingAccount>, I>>(object: I): EventNewVestingAccount {
    const message = createBaseEventNewVestingAccount();
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseEventNewVestingPool(): EventNewVestingPool {
  return { owner: "", name: "", amount: "", duration: "", vestingType: "" };
}

export const EventNewVestingPool = {
  encode(message: EventNewVestingPool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    if (message.duration !== "") {
      writer.uint32(34).string(message.duration);
    }
    if (message.vestingType !== "") {
      writer.uint32(42).string(message.vestingType);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EventNewVestingPool {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventNewVestingPool();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.amount = reader.string();
          break;
        case 4:
          message.duration = reader.string();
          break;
        case 5:
          message.vestingType = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventNewVestingPool {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      name: isSet(object.name) ? String(object.name) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      duration: isSet(object.duration) ? String(object.duration) : "",
      vestingType: isSet(object.vestingType) ? String(object.vestingType) : "",
    };
  },

  toJSON(message: EventNewVestingPool): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.name !== undefined && (obj.name = message.name);
    message.amount !== undefined && (obj.amount = message.amount);
    message.duration !== undefined && (obj.duration = message.duration);
    message.vestingType !== undefined && (obj.vestingType = message.vestingType);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventNewVestingPool>, I>>(object: I): EventNewVestingPool {
    const message = createBaseEventNewVestingPool();
    message.owner = object.owner ?? "";
    message.name = object.name ?? "";
    message.amount = object.amount ?? "";
    message.duration = object.duration ?? "";
    message.vestingType = object.vestingType ?? "";
    return message;
  },
};

function createBaseEventNewVestingPeriodFromVestingPool(): EventNewVestingPeriodFromVestingPool {
  return { owner: "", address: "", vestingPoolName: "", amount: "", restartVesting: "", periodId: 0 };
}

export const EventNewVestingPeriodFromVestingPool = {
  encode(message: EventNewVestingPeriodFromVestingPool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    if (message.vestingPoolName !== "") {
      writer.uint32(26).string(message.vestingPoolName);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    if (message.restartVesting !== "") {
      writer.uint32(42).string(message.restartVesting);
    }
    if (message.periodId !== 0) {
      writer.uint32(48).uint64(message.periodId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EventNewVestingPeriodFromVestingPool {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventNewVestingPeriodFromVestingPool();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.vestingPoolName = reader.string();
          break;
        case 4:
          message.amount = reader.string();
          break;
        case 5:
          message.restartVesting = reader.string();
          break;
        case 6:
          message.periodId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventNewVestingPeriodFromVestingPool {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      address: isSet(object.address) ? String(object.address) : "",
      vestingPoolName: isSet(object.vestingPoolName) ? String(object.vestingPoolName) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      restartVesting: isSet(object.restartVesting) ? String(object.restartVesting) : "",
      periodId: isSet(object.periodId) ? Number(object.periodId) : 0,
    };
  },

  toJSON(message: EventNewVestingPeriodFromVestingPool): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.address !== undefined && (obj.address = message.address);
    message.vestingPoolName !== undefined && (obj.vestingPoolName = message.vestingPoolName);
    message.amount !== undefined && (obj.amount = message.amount);
    message.restartVesting !== undefined && (obj.restartVesting = message.restartVesting);
    message.periodId !== undefined && (obj.periodId = Math.round(message.periodId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventNewVestingPeriodFromVestingPool>, I>>(
    object: I,
  ): EventNewVestingPeriodFromVestingPool {
    const message = createBaseEventNewVestingPeriodFromVestingPool();
    message.owner = object.owner ?? "";
    message.address = object.address ?? "";
    message.vestingPoolName = object.vestingPoolName ?? "";
    message.amount = object.amount ?? "";
    message.restartVesting = object.restartVesting ?? "";
    message.periodId = object.periodId ?? 0;
    return message;
  },
};

function createBaseEventWithdrawAvailable(): EventWithdrawAvailable {
  return { owner: "", vestingPoolName: "", amount: "" };
}

export const EventWithdrawAvailable = {
  encode(message: EventWithdrawAvailable, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.vestingPoolName !== "") {
      writer.uint32(18).string(message.vestingPoolName);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EventWithdrawAvailable {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventWithdrawAvailable();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.vestingPoolName = reader.string();
          break;
        case 3:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventWithdrawAvailable {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      vestingPoolName: isSet(object.vestingPoolName) ? String(object.vestingPoolName) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
    };
  },

  toJSON(message: EventWithdrawAvailable): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.vestingPoolName !== undefined && (obj.vestingPoolName = message.vestingPoolName);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventWithdrawAvailable>, I>>(object: I): EventWithdrawAvailable {
    const message = createBaseEventWithdrawAvailable();
    message.owner = object.owner ?? "";
    message.vestingPoolName = object.vestingPoolName ?? "";
    message.amount = object.amount ?? "";
    return message;
  },
};

function createBaseEventVestingSplit(): EventVestingSplit {
  return { source: "", destination: "" };
}

export const EventVestingSplit = {
  encode(message: EventVestingSplit, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.source !== "") {
      writer.uint32(10).string(message.source);
    }
    if (message.destination !== "") {
      writer.uint32(18).string(message.destination);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EventVestingSplit {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventVestingSplit();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.source = reader.string();
          break;
        case 2:
          message.destination = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventVestingSplit {
    return {
      source: isSet(object.source) ? String(object.source) : "",
      destination: isSet(object.destination) ? String(object.destination) : "",
    };
  },

  toJSON(message: EventVestingSplit): unknown {
    const obj: any = {};
    message.source !== undefined && (obj.source = message.source);
    message.destination !== undefined && (obj.destination = message.destination);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventVestingSplit>, I>>(object: I): EventVestingSplit {
    const message = createBaseEventVestingSplit();
    message.source = object.source ?? "";
    message.destination = object.destination ?? "";
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
