/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import { Duration } from "../google/protobuf/duration";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface AccountVestingsList {
  vestings: AccountVestings[];
}

export interface AccountVestings {
  address: string;
  delegable_address: string;
  vestings: Vesting[];
}

export interface Vesting {
  id: number;
  vesting_type: string;
  vesting_start: Date | undefined;
  lock_end: Date | undefined;
  vesting_end: Date | undefined;
  vested: string;
  release_period: Duration | undefined;
  delegation_allowed: boolean;
  withdrawn: string;
  sent: string;
  last_modification: Date | undefined;
  last_modification_vested: string;
  last_modification_withdrawn: string;
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

const baseAccountVestings: object = { address: "", delegable_address: "" };

export const AccountVestings = {
  encode(message: AccountVestings, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.delegable_address !== "") {
      writer.uint32(18).string(message.delegable_address);
    }
    for (const v of message.vestings) {
      Vesting.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AccountVestings {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAccountVestings } as AccountVestings;
    message.vestings = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.delegable_address = reader.string();
          break;
        case 3:
          message.vestings.push(Vesting.decode(reader, reader.uint32()));
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
    message.vestings = [];
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (
      object.delegable_address !== undefined &&
      object.delegable_address !== null
    ) {
      message.delegable_address = String(object.delegable_address);
    } else {
      message.delegable_address = "";
    }
    if (object.vestings !== undefined && object.vestings !== null) {
      for (const e of object.vestings) {
        message.vestings.push(Vesting.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: AccountVestings): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.delegable_address !== undefined &&
      (obj.delegable_address = message.delegable_address);
    if (message.vestings) {
      obj.vestings = message.vestings.map((e) =>
        e ? Vesting.toJSON(e) : undefined
      );
    } else {
      obj.vestings = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<AccountVestings>): AccountVestings {
    const message = { ...baseAccountVestings } as AccountVestings;
    message.vestings = [];
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (
      object.delegable_address !== undefined &&
      object.delegable_address !== null
    ) {
      message.delegable_address = object.delegable_address;
    } else {
      message.delegable_address = "";
    }
    if (object.vestings !== undefined && object.vestings !== null) {
      for (const e of object.vestings) {
        message.vestings.push(Vesting.fromPartial(e));
      }
    }
    return message;
  },
};

const baseVesting: object = {
  id: 0,
  vesting_type: "",
  vested: "",
  delegation_allowed: false,
  withdrawn: "",
  sent: "",
  last_modification_vested: "",
  last_modification_withdrawn: "",
};

export const Vesting = {
  encode(message: Vesting, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    if (message.vesting_type !== "") {
      writer.uint32(18).string(message.vesting_type);
    }
    if (message.vesting_start !== undefined) {
      Timestamp.encode(
        toTimestamp(message.vesting_start),
        writer.uint32(26).fork()
      ).ldelim();
    }
    if (message.lock_end !== undefined) {
      Timestamp.encode(
        toTimestamp(message.lock_end),
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.vesting_end !== undefined) {
      Timestamp.encode(
        toTimestamp(message.vesting_end),
        writer.uint32(42).fork()
      ).ldelim();
    }
    if (message.vested !== "") {
      writer.uint32(50).string(message.vested);
    }
    if (message.release_period !== undefined) {
      Duration.encode(
        message.release_period,
        writer.uint32(58).fork()
      ).ldelim();
    }
    if (message.delegation_allowed === true) {
      writer.uint32(64).bool(message.delegation_allowed);
    }
    if (message.withdrawn !== "") {
      writer.uint32(74).string(message.withdrawn);
    }
    if (message.sent !== "") {
      writer.uint32(82).string(message.sent);
    }
    if (message.last_modification !== undefined) {
      Timestamp.encode(
        toTimestamp(message.last_modification),
        writer.uint32(90).fork()
      ).ldelim();
    }
    if (message.last_modification_vested !== "") {
      writer.uint32(98).string(message.last_modification_vested);
    }
    if (message.last_modification_withdrawn !== "") {
      writer.uint32(106).string(message.last_modification_withdrawn);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Vesting {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVesting } as Vesting;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.int32();
          break;
        case 2:
          message.vesting_type = reader.string();
          break;
        case 3:
          message.vesting_start = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 4:
          message.lock_end = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.vesting_end = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 6:
          message.vested = reader.string();
          break;
        case 7:
          message.release_period = Duration.decode(reader, reader.uint32());
          break;
        case 8:
          message.delegation_allowed = reader.bool();
          break;
        case 9:
          message.withdrawn = reader.string();
          break;
        case 10:
          message.sent = reader.string();
          break;
        case 11:
          message.last_modification = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 12:
          message.last_modification_vested = reader.string();
          break;
        case 13:
          message.last_modification_withdrawn = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Vesting {
    const message = { ...baseVesting } as Vesting;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.vesting_type !== undefined && object.vesting_type !== null) {
      message.vesting_type = String(object.vesting_type);
    } else {
      message.vesting_type = "";
    }
    if (object.vesting_start !== undefined && object.vesting_start !== null) {
      message.vesting_start = fromJsonTimestamp(object.vesting_start);
    } else {
      message.vesting_start = undefined;
    }
    if (object.lock_end !== undefined && object.lock_end !== null) {
      message.lock_end = fromJsonTimestamp(object.lock_end);
    } else {
      message.lock_end = undefined;
    }
    if (object.vesting_end !== undefined && object.vesting_end !== null) {
      message.vesting_end = fromJsonTimestamp(object.vesting_end);
    } else {
      message.vesting_end = undefined;
    }
    if (object.vested !== undefined && object.vested !== null) {
      message.vested = String(object.vested);
    } else {
      message.vested = "";
    }
    if (object.release_period !== undefined && object.release_period !== null) {
      message.release_period = Duration.fromJSON(object.release_period);
    } else {
      message.release_period = undefined;
    }
    if (
      object.delegation_allowed !== undefined &&
      object.delegation_allowed !== null
    ) {
      message.delegation_allowed = Boolean(object.delegation_allowed);
    } else {
      message.delegation_allowed = false;
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
    return message;
  },

  toJSON(message: Vesting): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.vesting_type !== undefined &&
      (obj.vesting_type = message.vesting_type);
    message.vesting_start !== undefined &&
      (obj.vesting_start =
        message.vesting_start !== undefined
          ? message.vesting_start.toISOString()
          : null);
    message.lock_end !== undefined &&
      (obj.lock_end =
        message.lock_end !== undefined ? message.lock_end.toISOString() : null);
    message.vesting_end !== undefined &&
      (obj.vesting_end =
        message.vesting_end !== undefined
          ? message.vesting_end.toISOString()
          : null);
    message.vested !== undefined && (obj.vested = message.vested);
    message.release_period !== undefined &&
      (obj.release_period = message.release_period
        ? Duration.toJSON(message.release_period)
        : undefined);
    message.delegation_allowed !== undefined &&
      (obj.delegation_allowed = message.delegation_allowed);
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
    return obj;
  },

  fromPartial(object: DeepPartial<Vesting>): Vesting {
    const message = { ...baseVesting } as Vesting;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.vesting_type !== undefined && object.vesting_type !== null) {
      message.vesting_type = object.vesting_type;
    } else {
      message.vesting_type = "";
    }
    if (object.vesting_start !== undefined && object.vesting_start !== null) {
      message.vesting_start = object.vesting_start;
    } else {
      message.vesting_start = undefined;
    }
    if (object.lock_end !== undefined && object.lock_end !== null) {
      message.lock_end = object.lock_end;
    } else {
      message.lock_end = undefined;
    }
    if (object.vesting_end !== undefined && object.vesting_end !== null) {
      message.vesting_end = object.vesting_end;
    } else {
      message.vesting_end = undefined;
    }
    if (object.vested !== undefined && object.vested !== null) {
      message.vested = object.vested;
    } else {
      message.vested = "";
    }
    if (object.release_period !== undefined && object.release_period !== null) {
      message.release_period = Duration.fromPartial(object.release_period);
    } else {
      message.release_period = undefined;
    }
    if (
      object.delegation_allowed !== undefined &&
      object.delegation_allowed !== null
    ) {
      message.delegation_allowed = object.delegation_allowed;
    } else {
      message.delegation_allowed = false;
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
