/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Timestamp } from "../google/protobuf/timestamp";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface AccountVestingPools {
  address: string;
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
}

function createBaseAccountVestingPools(): AccountVestingPools {
  return { address: "", vestingPools: [] };
}

export const AccountVestingPools = {
  encode(message: AccountVestingPools, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    for (const v of message.vestingPools) {
      VestingPool.encode(v!, writer.uint32(26).fork()).ldelim();
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
          message.address = reader.string();
          break;
        case 3:
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
      address: isSet(object.address) ? String(object.address) : "",
      vestingPools: Array.isArray(object?.vestingPools)
        ? object.vestingPools.map((e: any) => VestingPool.fromJSON(e))
        : [],
    };
  },

  toJSON(message: AccountVestingPools): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    if (message.vestingPools) {
      obj.vestingPools = message.vestingPools.map((e) => e ? VestingPool.toJSON(e) : undefined);
    } else {
      obj.vestingPools = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AccountVestingPools>, I>>(object: I): AccountVestingPools {
    const message = createBaseAccountVestingPools();
    message.address = object.address ?? "";
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
    return message;
  },
};

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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
