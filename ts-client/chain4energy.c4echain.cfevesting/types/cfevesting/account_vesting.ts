/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface AccountVestingsList {
  vestings: AccountVestings[];
}

export interface AccountVestings {
  address: string;
  /** string delegable_address = 2; */
  vestingPools: VestingPool[];
}

export interface VestingPool {
  id: number;
  name: string;
  vestingType: string;
  lockStart: Date | undefined;
  lockEnd: Date | undefined;
  vested: string;
  withdrawn: string;
  sent: string;
  lastModification: Date | undefined;
  lastModificationVested: string;
  lastModificationWithdrawn: string;
}

const baseAccountVestingsList: object = {};

export const AccountVestingsList = {
  encode(
    message: AccountVestingsList,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.vestings) {
      AccountVestings.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AccountVestingsList {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAccountVestingsList } as AccountVestingsList;
    message.vestings = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vestings.push(
            AccountVestings.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AccountVestingsList {
    const message = { ...baseAccountVestingsList } as AccountVestingsList;
    message.vestings = [];
    if (object.vestings !== undefined && object.vestings !== null) {
      for (const e of object.vestings) {
        message.vestings.push(AccountVestings.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: AccountVestingsList): unknown {
    const obj: any = {};
    if (message.vestings) {
      obj.vestings = message.vestings.map((e) =>
        e ? AccountVestings.toJSON(e) : undefined
      );
    } else {
      obj.vestings = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<AccountVestingsList>): AccountVestingsList {
    const message = { ...baseAccountVestingsList } as AccountVestingsList;
    message.vestings = [];
    if (object.vestings !== undefined && object.vestings !== null) {
      for (const e of object.vestings) {
        message.vestings.push(AccountVestings.fromPartial(e));
      }
    }
    return message;
  },
};

const baseAccountVestings: object = { address: "" };

export const AccountVestings = {
  encode(message: AccountVestings, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    for (const v of message.vestingPools) {
      VestingPool.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AccountVestings {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAccountVestings } as AccountVestings;
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

  fromJSON(object: any): AccountVestings {
    const message = { ...baseAccountVestings } as AccountVestings;
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

  toJSON(message: AccountVestings): unknown {
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

  fromPartial(object: DeepPartial<AccountVestings>): AccountVestings {
    const message = { ...baseAccountVestings } as AccountVestings;
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
  id: 0,
  name: "",
  vestingType: "",
  vested: "",
  withdrawn: "",
  sent: "",
  lastModificationVested: "",
  lastModificationWithdrawn: "",
};

export const VestingPool = {
  encode(message: VestingPool, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.vestingType !== "") {
      writer.uint32(26).string(message.vestingType);
    }
    if (message.lockStart !== undefined) {
      Timestamp.encode(
        toTimestamp(message.lockStart),
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.lockEnd !== undefined) {
      Timestamp.encode(
        toTimestamp(message.lockEnd),
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.vested !== "") {
      writer.uint32(50).string(message.vested);
    }
    if (message.withdrawn !== "") {
      writer.uint32(58).string(message.withdrawn);
    }
    if (message.sent !== "") {
      writer.uint32(66).string(message.sent);
    }
    if (message.lastModification !== undefined) {
      Timestamp.encode(
        toTimestamp(message.lastModification),
        writer.uint32(74).fork()
      ).ldelim();
    }
    if (message.lastModificationVested !== "") {
      writer.uint32(82).string(message.lastModificationVested);
    }
    if (message.lastModificationWithdrawn !== "") {
      writer.uint32(90).string(message.lastModificationWithdrawn);
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
          message.id = reader.int32();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.vestingType = reader.string();
          break;
        case 4:
          message.lockStart = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.lockEnd = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 6:
          message.vested = reader.string();
          break;
        case 7:
          message.withdrawn = reader.string();
          break;
        case 8:
          message.sent = reader.string();
          break;
        case 9:
          message.lastModification = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 10:
          message.lastModificationVested = reader.string();
          break;
        case 11:
          message.lastModificationWithdrawn = reader.string();
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
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
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
    if (object.vested !== undefined && object.vested !== null) {
      message.vested = String(object.vested);
    } else {
      message.vested = "";
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
    if (
      object.lastModification !== undefined &&
      object.lastModification !== null
    ) {
      message.lastModification = fromJsonTimestamp(object.lastModification);
    } else {
      message.lastModification = undefined;
    }
    if (
      object.lastModificationVested !== undefined &&
      object.lastModificationVested !== null
    ) {
      message.lastModificationVested = String(object.lastModificationVested);
    } else {
      message.lastModificationVested = "";
    }
    if (
      object.lastModificationWithdrawn !== undefined &&
      object.lastModificationWithdrawn !== null
    ) {
      message.lastModificationWithdrawn = String(
        object.lastModificationWithdrawn
      );
    } else {
      message.lastModificationWithdrawn = "";
    }
    return message;
  },

  toJSON(message: VestingPool): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
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
    message.vested !== undefined && (obj.vested = message.vested);
    message.withdrawn !== undefined && (obj.withdrawn = message.withdrawn);
    message.sent !== undefined && (obj.sent = message.sent);
    message.lastModification !== undefined &&
      (obj.lastModification =
        message.lastModification !== undefined
          ? message.lastModification.toISOString()
          : null);
    message.lastModificationVested !== undefined &&
      (obj.lastModificationVested = message.lastModificationVested);
    message.lastModificationWithdrawn !== undefined &&
      (obj.lastModificationWithdrawn = message.lastModificationWithdrawn);
    return obj;
  },

  fromPartial(object: DeepPartial<VestingPool>): VestingPool {
    const message = { ...baseVestingPool } as VestingPool;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
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
    if (object.vested !== undefined && object.vested !== null) {
      message.vested = object.vested;
    } else {
      message.vested = "";
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
    if (
      object.lastModification !== undefined &&
      object.lastModification !== null
    ) {
      message.lastModification = object.lastModification;
    } else {
      message.lastModification = undefined;
    }
    if (
      object.lastModificationVested !== undefined &&
      object.lastModificationVested !== null
    ) {
      message.lastModificationVested = object.lastModificationVested;
    } else {
      message.lastModificationVested = "";
    }
    if (
      object.lastModificationWithdrawn !== undefined &&
      object.lastModificationWithdrawn !== null
    ) {
      message.lastModificationWithdrawn = object.lastModificationWithdrawn;
    } else {
      message.lastModificationWithdrawn = "";
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
