/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import { Minter } from "../cfeminter/minter";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeminter";

/** Params defines the parameters for the module. */
export interface Params {
  mint_denom: string;
  start_time: Date | undefined;
  minters: Minter[];
}

const baseParams: object = { mint_denom: "" };

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.mint_denom !== "") {
      writer.uint32(10).string(message.mint_denom);
    }
    if (message.start_time !== undefined) {
      Timestamp.encode(
        toTimestamp(message.start_time),
        writer.uint32(18).fork()
      ).ldelim();
    }
    for (const v of message.minters) {
      Minter.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    message.minters = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.mint_denom = reader.string();
          break;
        case 2:
          message.start_time = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.minters.push(Minter.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    message.minters = [];
    if (object.mint_denom !== undefined && object.mint_denom !== null) {
      message.mint_denom = String(object.mint_denom);
    } else {
      message.mint_denom = "";
    }
    if (object.start_time !== undefined && object.start_time !== null) {
      message.start_time = fromJsonTimestamp(object.start_time);
    } else {
      message.start_time = undefined;
    }
    if (object.minters !== undefined && object.minters !== null) {
      for (const e of object.minters) {
        message.minters.push(Minter.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.mint_denom !== undefined && (obj.mint_denom = message.mint_denom);
    message.start_time !== undefined &&
      (obj.start_time =
        message.start_time !== undefined
          ? message.start_time.toISOString()
          : null);
    if (message.minters) {
      obj.minters = message.minters.map((e) =>
        e ? Minter.toJSON(e) : undefined
      );
    } else {
      obj.minters = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    message.minters = [];
    if (object.mint_denom !== undefined && object.mint_denom !== null) {
      message.mint_denom = object.mint_denom;
    } else {
      message.mint_denom = "";
    }
    if (object.start_time !== undefined && object.start_time !== null) {
      message.start_time = object.start_time;
    } else {
      message.start_time = undefined;
    }
    if (object.minters !== undefined && object.minters !== null) {
      for (const e of object.minters) {
        message.minters.push(Minter.fromPartial(e));
      }
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
