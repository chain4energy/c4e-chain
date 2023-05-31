/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Timestamp } from "../../google/protobuf/timestamp";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface AccountVestingPools {
  owner: string;
  vestingPools: VestingPool[];
}

export interface VestingPool {
  name: string;
  vestingType: string;
  lockStart: Date | undefined;
  lockEnd: Date | undefined;
  initiallyLocked: string;
  withdrawn: string;
  sent: string;
  genesisPool: boolean;
  reservations: VestingPoolReservation[];
}

export interface VestingPoolReservation {
  id: number;
  amount: string;
}

function createBaseAccountVestingPools(): AccountVestingPools {
  return { owner: "", vestingPools: [] };
}

export const AccountVestingPools = {
  encode(message: AccountVestingPools, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    for (const v of message.vestingPools) {
      VestingPool.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AccountVestingPools {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAccountVestingPools();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.vestingPools.push(VestingPool.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AccountVestingPools {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      vestingPools: Array.isArray(object?.vestingPools)
        ? object.vestingPools.map((e: any) => VestingPool.fromJSON(e))
        : [],
    };
  },

  toJSON(message: AccountVestingPools): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    if (message.vestingPools) {
      obj.vestingPools = message.vestingPools.map((e) => e ? VestingPool.toJSON(e) : undefined);
    } else {
      obj.vestingPools = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AccountVestingPools>, I>>(object: I): AccountVestingPools {
    const message = createBaseAccountVestingPools();
    message.owner = object.owner ?? "";
    message.vestingPools = object.vestingPools?.map((e) => VestingPool.fromPartial(e)) || [];
    return message;
  },
};

function createBaseVestingPool(): VestingPool {
  return {
    name: "",
    vestingType: "",
    lockStart: undefined,
    lockEnd: undefined,
    initiallyLocked: "",
    withdrawn: "",
    sent: "",
    genesisPool: false,
    reservations: [],
  };
}

export const VestingPool = {
  encode(message: VestingPool, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.vestingType !== "") {
      writer.uint32(18).string(message.vestingType);
    }
    if (message.lockStart !== undefined) {
      Timestamp.encode(toTimestamp(message.lockStart), writer.uint32(26).fork()).ldelim();
    }
    if (message.lockEnd !== undefined) {
      Timestamp.encode(toTimestamp(message.lockEnd), writer.uint32(34).fork()).ldelim();
    }
    if (message.initiallyLocked !== "") {
      writer.uint32(42).string(message.initiallyLocked);
    }
    if (message.withdrawn !== "") {
      writer.uint32(50).string(message.withdrawn);
    }
    if (message.sent !== "") {
      writer.uint32(58).string(message.sent);
    }
    if (message.genesisPool === true) {
      writer.uint32(64).bool(message.genesisPool);
    }
    for (const v of message.reservations) {
      VestingPoolReservation.encode(v!, writer.uint32(74).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VestingPool {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVestingPool();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.vestingType = reader.string();
          break;
        case 3:
          message.lockStart = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 4:
          message.lockEnd = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 5:
          message.initiallyLocked = reader.string();
          break;
        case 6:
          message.withdrawn = reader.string();
          break;
        case 7:
          message.sent = reader.string();
          break;
        case 8:
          message.genesisPool = reader.bool();
          break;
        case 9:
          message.reservations.push(VestingPoolReservation.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VestingPool {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      vestingType: isSet(object.vestingType) ? String(object.vestingType) : "",
      lockStart: isSet(object.lockStart) ? fromJsonTimestamp(object.lockStart) : undefined,
      lockEnd: isSet(object.lockEnd) ? fromJsonTimestamp(object.lockEnd) : undefined,
      initiallyLocked: isSet(object.initiallyLocked) ? String(object.initiallyLocked) : "",
      withdrawn: isSet(object.withdrawn) ? String(object.withdrawn) : "",
      sent: isSet(object.sent) ? String(object.sent) : "",
      genesisPool: isSet(object.genesisPool) ? Boolean(object.genesisPool) : false,
      reservations: Array.isArray(object?.reservations)
        ? object.reservations.map((e: any) => VestingPoolReservation.fromJSON(e))
        : [],
    };
  },

  toJSON(message: VestingPool): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.vestingType !== undefined && (obj.vestingType = message.vestingType);
    message.lockStart !== undefined && (obj.lockStart = message.lockStart.toISOString());
    message.lockEnd !== undefined && (obj.lockEnd = message.lockEnd.toISOString());
    message.initiallyLocked !== undefined && (obj.initiallyLocked = message.initiallyLocked);
    message.withdrawn !== undefined && (obj.withdrawn = message.withdrawn);
    message.sent !== undefined && (obj.sent = message.sent);
    message.genesisPool !== undefined && (obj.genesisPool = message.genesisPool);
    if (message.reservations) {
      obj.reservations = message.reservations.map((e) => e ? VestingPoolReservation.toJSON(e) : undefined);
    } else {
      obj.reservations = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VestingPool>, I>>(object: I): VestingPool {
    const message = createBaseVestingPool();
    message.name = object.name ?? "";
    message.vestingType = object.vestingType ?? "";
    message.lockStart = object.lockStart ?? undefined;
    message.lockEnd = object.lockEnd ?? undefined;
    message.initiallyLocked = object.initiallyLocked ?? "";
    message.withdrawn = object.withdrawn ?? "";
    message.sent = object.sent ?? "";
    message.genesisPool = object.genesisPool ?? false;
    message.reservations = object.reservations?.map((e) => VestingPoolReservation.fromPartial(e)) || [];
    return message;
  },
};

function createBaseVestingPoolReservation(): VestingPoolReservation {
  return { id: 0, amount: "" };
}

export const VestingPoolReservation = {
  encode(message: VestingPoolReservation, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): VestingPoolReservation {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseVestingPoolReservation();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): VestingPoolReservation {
    return { id: isSet(object.id) ? Number(object.id) : 0, amount: isSet(object.amount) ? String(object.amount) : "" };
  },

  toJSON(message: VestingPoolReservation): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<VestingPoolReservation>, I>>(object: I): VestingPoolReservation {
    const message = createBaseVestingPoolReservation();
    message.id = object.id ?? 0;
    message.amount = object.amount ?? "";
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

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
