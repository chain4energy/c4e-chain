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
  vesting_pools: VestingPool[];
}

export interface VestingPool {
  id: number;
  name: string;
  vesting_type: string;
  lock_start: Date | undefined;
  lock_end: Date | undefined;
  /** google.protobuf.Timestamp vesting_end = 5 [(gogoproto.nullable) = false, (gogoproto.stdtime) = true]; */
  vested: string;
  /**
   * google.protobuf.Duration release_period = 7 [(gogoproto.nullable) = false, (gogoproto.stdduration) = true];
   * bool delegation_allowed = 8;
   */
  withdrawn: string;
  sent: string;
  last_modification: Date | undefined;
  last_modification_vested: string;
  last_modification_withdrawn: string;
  transfer_allowed: boolean;
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
    for (const v of message.vesting_pools) {
      VestingPool.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AccountVestings {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAccountVestings } as AccountVestings;
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

  fromJSON(object: any): AccountVestings {
    const message = { ...baseAccountVestings } as AccountVestings;
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

  toJSON(message: AccountVestings): unknown {
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

  fromPartial(object: DeepPartial<AccountVestings>): AccountVestings {
    const message = { ...baseAccountVestings } as AccountVestings;
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
  id: 0,
  name: "",
  vesting_type: "",
  vested: "",
  withdrawn: "",
  sent: "",
  last_modification_vested: "",
  last_modification_withdrawn: "",
  transfer_allowed: false,
};

export const VestingPool = {
  encode(message: VestingPool, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.vesting_type !== "") {
      writer.uint32(26).string(message.vesting_type);
    }
    if (message.lock_start !== undefined) {
      Timestamp.encode(
        toTimestamp(message.lock_start),
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.lock_end !== undefined) {
      Timestamp.encode(
        toTimestamp(message.lock_end),
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
    if (message.last_modification !== undefined) {
      Timestamp.encode(
        toTimestamp(message.last_modification),
        writer.uint32(74).fork()
      ).ldelim();
    }
    if (message.last_modification_vested !== "") {
      writer.uint32(82).string(message.last_modification_vested);
    }
    if (message.last_modification_withdrawn !== "") {
      writer.uint32(90).string(message.last_modification_withdrawn);
    }
    if (message.transfer_allowed === true) {
      writer.uint32(96).bool(message.transfer_allowed);
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
          message.vesting_type = reader.string();
          break;
        case 4:
          message.lock_start = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.lock_end = fromTimestamp(
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
          message.last_modification = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 10:
          message.last_modification_vested = reader.string();
          break;
        case 11:
          message.last_modification_withdrawn = reader.string();
          break;
        case 12:
          message.transfer_allowed = reader.bool();
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
      object.last_modification !== undefined &&
      object.last_modification !== null
    ) {
      message.last_modification = fromJsonTimestamp(object.last_modification);
    } else {
      message.last_modification = undefined;
    }
    if (
      object.last_modification_vested !== undefined &&
      object.last_modification_vested !== null
    ) {
      message.last_modification_vested = String(
        object.last_modification_vested
      );
    } else {
      message.last_modification_vested = "";
    }
    if (
      object.last_modification_withdrawn !== undefined &&
      object.last_modification_withdrawn !== null
    ) {
      message.last_modification_withdrawn = String(
        object.last_modification_withdrawn
      );
    } else {
      message.last_modification_withdrawn = "";
    }
    if (
      object.transfer_allowed !== undefined &&
      object.transfer_allowed !== null
    ) {
      message.transfer_allowed = Boolean(object.transfer_allowed);
    } else {
      message.transfer_allowed = false;
    }
    return message;
  },

  toJSON(message: VestingPool): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
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
    message.vested !== undefined && (obj.vested = message.vested);
    message.withdrawn !== undefined && (obj.withdrawn = message.withdrawn);
    message.sent !== undefined && (obj.sent = message.sent);
    message.last_modification !== undefined &&
      (obj.last_modification =
        message.last_modification !== undefined
          ? message.last_modification.toISOString()
          : null);
    message.last_modification_vested !== undefined &&
      (obj.last_modification_vested = message.last_modification_vested);
    message.last_modification_withdrawn !== undefined &&
      (obj.last_modification_withdrawn = message.last_modification_withdrawn);
    message.transfer_allowed !== undefined &&
      (obj.transfer_allowed = message.transfer_allowed);
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
      object.last_modification !== undefined &&
      object.last_modification !== null
    ) {
      message.last_modification = object.last_modification;
    } else {
      message.last_modification = undefined;
    }
    if (
      object.last_modification_vested !== undefined &&
      object.last_modification_vested !== null
    ) {
      message.last_modification_vested = object.last_modification_vested;
    } else {
      message.last_modification_vested = "";
    }
    if (
      object.last_modification_withdrawn !== undefined &&
      object.last_modification_withdrawn !== null
    ) {
      message.last_modification_withdrawn = object.last_modification_withdrawn;
    } else {
      message.last_modification_withdrawn = "";
    }
    if (
      object.transfer_allowed !== undefined &&
      object.transfer_allowed !== null
    ) {
      message.transfer_allowed = object.transfer_allowed;
    } else {
      message.transfer_allowed = false;
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
