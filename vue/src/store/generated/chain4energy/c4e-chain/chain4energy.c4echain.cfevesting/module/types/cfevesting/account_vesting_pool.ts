/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface AccountVestingPools {
  address: string;
  vesting_pools: VestingPool[];
}

export interface VestingPool {
  name: string;
  vesting_type: string;
  lock_start: Date | undefined;
  lock_end: Date | undefined;
  initially_locked: string;
  withdrawn: string;
  sent: string;
}

const baseAccountVestingPools: object = { address: "" };

export const AccountVestingPools = {
  encode(
    message: AccountVestingPools,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    for (const v of message.vesting_pools) {
      VestingPool.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AccountVestingPools {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAccountVestingPools } as AccountVestingPools;
    message.vesting_pools = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 3:
          message.vesting_pools.push(
            VestingPool.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AccountVestingPools {
    const message = { ...baseAccountVestingPools } as AccountVestingPools;
    message.vesting_pools = [];
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.vesting_pools !== undefined && object.vesting_pools !== null) {
      for (const e of object.vesting_pools) {
        message.vesting_pools.push(VestingPool.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: AccountVestingPools): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    if (message.vesting_pools) {
      obj.vesting_pools = message.vesting_pools.map((e) =>
        e ? VestingPool.toJSON(e) : undefined
      );
    } else {
      obj.vesting_pools = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<AccountVestingPools>): AccountVestingPools {
    const message = { ...baseAccountVestingPools } as AccountVestingPools;
    message.vesting_pools = [];
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.vesting_pools !== undefined && object.vesting_pools !== null) {
      for (const e of object.vesting_pools) {
        message.vesting_pools.push(VestingPool.fromPartial(e));
      }
    }
    return message;
  },
};

const baseVestingPool: object = {
  name: "",
  vesting_type: "",
  initially_locked: "",
  withdrawn: "",
  sent: "",
};

export const VestingPool = {
  encode(message: VestingPool, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.vesting_type !== "") {
      writer.uint32(18).string(message.vesting_type);
    }
    if (message.lock_start !== undefined) {
      Timestamp.encode(
        toTimestamp(message.lock_start),
        writer.uint32(26).fork()
      ).ldelim();
    }
    if (message.lock_end !== undefined) {
      Timestamp.encode(
        toTimestamp(message.lock_end),
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.initially_locked !== "") {
      writer.uint32(42).string(message.initially_locked);
    }
    if (message.withdrawn !== "") {
      writer.uint32(50).string(message.withdrawn);
    }
    if (message.sent !== "") {
      writer.uint32(58).string(message.sent);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): VestingPool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVestingPool } as VestingPool;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.vesting_type = reader.string();
          break;
        case 3:
          message.lock_start = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.lock_end = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.initially_locked = reader.string();
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
    const message = { ...baseVestingPool } as VestingPool;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.vesting_type !== undefined && object.vesting_type !== null) {
      message.vesting_type = String(object.vesting_type);
    } else {
      message.vesting_type = "";
    }
    if (object.lock_start !== undefined && object.lock_start !== null) {
      message.lock_start = fromJsonTimestamp(object.lock_start);
    } else {
      message.lock_start = undefined;
    }
    if (object.lock_end !== undefined && object.lock_end !== null) {
      message.lock_end = fromJsonTimestamp(object.lock_end);
    } else {
      message.lock_end = undefined;
    }
    if (
      object.initially_locked !== undefined &&
      object.initially_locked !== null
    ) {
      message.initially_locked = String(object.initially_locked);
    } else {
      message.initially_locked = "";
    }
    if (object.withdrawn !== undefined && object.withdrawn !== null) {
      message.withdrawn = String(object.withdrawn);
    } else {
      message.withdrawn = "";
    }
    if (object.sent !== undefined && object.sent !== null) {
      message.sent = String(object.sent);
    } else {
      message.sent = "";
    }
    return message;
  },

  toJSON(message: VestingPool): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.vesting_type !== undefined &&
      (obj.vesting_type = message.vesting_type);
    message.lock_start !== undefined &&
      (obj.lock_start =
        message.lock_start !== undefined
          ? message.lock_start.toISOString()
          : null);
    message.lock_end !== undefined &&
      (obj.lock_end =
        message.lock_end !== undefined ? message.lock_end.toISOString() : null);
    message.initially_locked !== undefined &&
      (obj.initially_locked = message.initially_locked);
    message.withdrawn !== undefined && (obj.withdrawn = message.withdrawn);
    message.sent !== undefined && (obj.sent = message.sent);
    return obj;
  },

  fromPartial(object: DeepPartial<VestingPool>): VestingPool {
    const message = { ...baseVestingPool } as VestingPool;
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.vesting_type !== undefined && object.vesting_type !== null) {
      message.vesting_type = object.vesting_type;
    } else {
      message.vesting_type = "";
    }
    if (object.lock_start !== undefined && object.lock_start !== null) {
      message.lock_start = object.lock_start;
    } else {
      message.lock_start = undefined;
    }
    if (object.lock_end !== undefined && object.lock_end !== null) {
      message.lock_end = object.lock_end;
    } else {
      message.lock_end = undefined;
    }
    if (
      object.initially_locked !== undefined &&
      object.initially_locked !== null
    ) {
      message.initially_locked = object.initially_locked;
    } else {
      message.initially_locked = "";
    }
    if (object.withdrawn !== undefined && object.withdrawn !== null) {
      message.withdrawn = object.withdrawn;
    } else {
      message.withdrawn = "";
    }
    if (object.sent !== undefined && object.sent !== null) {
      message.sent = object.sent;
    } else {
      message.sent = "";
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
