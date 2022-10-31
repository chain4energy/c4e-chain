/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import { Writer, Reader } from "protobufjs/minimal";

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

const baseAccountVestingPools: object = { address: "" };

export const AccountVestingPools = {
  encode(
    message: AccountVestingPools,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    for (const v of message.vestingPools) {
      VestingPool.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AccountVestingPools {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAccountVestingPools } as AccountVestingPools;
    message.vestingPools = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 3:
          message.vestingPools.push(
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
    message.vestingPools = [];
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.vestingPools !== undefined && object.vestingPools !== null) {
      for (const e of object.vestingPools) {
        message.vestingPools.push(VestingPool.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: AccountVestingPools): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    if (message.vestingPools) {
      obj.vestingPools = message.vestingPools.map((e) =>
        e ? VestingPool.toJSON(e) : undefined
      );
    } else {
      obj.vestingPools = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<AccountVestingPools>): AccountVestingPools {
    const message = { ...baseAccountVestingPools } as AccountVestingPools;
    message.vestingPools = [];
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.vestingPools !== undefined && object.vestingPools !== null) {
      for (const e of object.vestingPools) {
        message.vestingPools.push(VestingPool.fromPartial(e));
      }
    }
    return message;
  },
};

const baseVestingPool: object = {
  name: "",
  vestingType: "",
  initiallyLocked: "",
  withdrawn: "",
  sent: "",
};

export const VestingPool = {
  encode(message: VestingPool, writer: Writer = Writer.create()): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.vestingType !== "") {
      writer.uint32(18).string(message.vestingType);
    }
    if (message.lockStart !== undefined) {
      Timestamp.encode(
        toTimestamp(message.lockStart),
        writer.uint32(26).fork()
      ).ldelim();
    }
    if (message.lockEnd !== undefined) {
      Timestamp.encode(
        toTimestamp(message.lockEnd),
        writer.uint32(34).fork()
      ).ldelim();
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
          message.vestingType = reader.string();
          break;
        case 3:
          message.lockStart = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.lockEnd = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
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
    const message = { ...baseVestingPool } as VestingPool;
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.vestingType !== undefined && object.vestingType !== null) {
      message.vestingType = String(object.vestingType);
    } else {
      message.vestingType = "";
    }
    if (object.lockStart !== undefined && object.lockStart !== null) {
      message.lockStart = fromJsonTimestamp(object.lockStart);
    } else {
      message.lockStart = undefined;
    }
    if (object.lockEnd !== undefined && object.lockEnd !== null) {
      message.lockEnd = fromJsonTimestamp(object.lockEnd);
    } else {
      message.lockEnd = undefined;
    }
    if (
      object.initiallyLocked !== undefined &&
      object.initiallyLocked !== null
    ) {
      message.initiallyLocked = String(object.initiallyLocked);
    } else {
      message.initiallyLocked = "";
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
    message.vestingType !== undefined &&
      (obj.vestingType = message.vestingType);
    message.lockStart !== undefined &&
      (obj.lockStart =
        message.lockStart !== undefined
          ? message.lockStart.toISOString()
          : null);
    message.lockEnd !== undefined &&
      (obj.lockEnd =
        message.lockEnd !== undefined ? message.lockEnd.toISOString() : null);
    message.initiallyLocked !== undefined &&
      (obj.initiallyLocked = message.initiallyLocked);
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
    if (object.vestingType !== undefined && object.vestingType !== null) {
      message.vestingType = object.vestingType;
    } else {
      message.vestingType = "";
    }
    if (object.lockStart !== undefined && object.lockStart !== null) {
      message.lockStart = object.lockStart;
    } else {
      message.lockStart = undefined;
    }
    if (object.lockEnd !== undefined && object.lockEnd !== null) {
      message.lockEnd = object.lockEnd;
    } else {
      message.lockEnd = undefined;
    }
    if (
      object.initiallyLocked !== undefined &&
      object.initiallyLocked !== null
    ) {
      message.initiallyLocked = object.initiallyLocked;
    } else {
      message.initiallyLocked = "";
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
