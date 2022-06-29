/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeminter";

/** HalvingMinter represents the inflation parameters. */
export interface HalvingMinter {
  /** the number of coins produced from the first block */
  new_coins_mint: number;
  /** expected blocks per year */
  blocks_per_year: number;
  /** type of coin to mint */
  mint_denom: string;
}

export interface Minter {
  start: Date | undefined;
  periods: MintingPeriod[];
}

export interface MintingPeriod {
  /** TODO change name to position */
  ordering_id: number;
  period_end: Date | undefined;
  type: MintingPeriod_MinterType;
  time_linear_minter: TimeLinearMinter | undefined;
}

export enum MintingPeriod_MinterType {
  NO_MINTING = 0,
  TIME_LINEAR_MINTER = 1,
  UNRECOGNIZED = -1,
}

export function mintingPeriod_MinterTypeFromJSON(
  object: any
): MintingPeriod_MinterType {
  switch (object) {
    case 0:
    case "NO_MINTING":
      return MintingPeriod_MinterType.NO_MINTING;
    case 1:
    case "TIME_LINEAR_MINTER":
      return MintingPeriod_MinterType.TIME_LINEAR_MINTER;
    case -1:
    case "UNRECOGNIZED":
    default:
      return MintingPeriod_MinterType.UNRECOGNIZED;
  }
}

export function mintingPeriod_MinterTypeToJSON(
  object: MintingPeriod_MinterType
): string {
  switch (object) {
    case MintingPeriod_MinterType.NO_MINTING:
      return "NO_MINTING";
    case MintingPeriod_MinterType.TIME_LINEAR_MINTER:
      return "TIME_LINEAR_MINTER";
    default:
      return "UNKNOWN";
  }
}

export interface TimeLinearMinter {
  amount: string;
}

export interface MinterState {
  current_ordering_id: number;
  amount_minted: string;
}

const baseHalvingMinter: object = {
  new_coins_mint: 0,
  blocks_per_year: 0,
  mint_denom: "",
};

export const HalvingMinter = {
  encode(message: HalvingMinter, writer: Writer = Writer.create()): Writer {
    if (message.new_coins_mint !== 0) {
      writer.uint32(8).int64(message.new_coins_mint);
    }
    if (message.blocks_per_year !== 0) {
      writer.uint32(48).int64(message.blocks_per_year);
    }
    if (message.mint_denom !== "") {
      writer.uint32(26).string(message.mint_denom);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): HalvingMinter {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseHalvingMinter } as HalvingMinter;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.new_coins_mint = longToNumber(reader.int64() as Long);
          break;
        case 6:
          message.blocks_per_year = longToNumber(reader.int64() as Long);
          break;
        case 3:
          message.mint_denom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): HalvingMinter {
    const message = { ...baseHalvingMinter } as HalvingMinter;
    if (object.new_coins_mint !== undefined && object.new_coins_mint !== null) {
      message.new_coins_mint = Number(object.new_coins_mint);
    } else {
      message.new_coins_mint = 0;
    }
    if (
      object.blocks_per_year !== undefined &&
      object.blocks_per_year !== null
    ) {
      message.blocks_per_year = Number(object.blocks_per_year);
    } else {
      message.blocks_per_year = 0;
    }
    if (object.mint_denom !== undefined && object.mint_denom !== null) {
      message.mint_denom = String(object.mint_denom);
    } else {
      message.mint_denom = "";
    }
    return message;
  },

  toJSON(message: HalvingMinter): unknown {
    const obj: any = {};
    message.new_coins_mint !== undefined &&
      (obj.new_coins_mint = message.new_coins_mint);
    message.blocks_per_year !== undefined &&
      (obj.blocks_per_year = message.blocks_per_year);
    message.mint_denom !== undefined && (obj.mint_denom = message.mint_denom);
    return obj;
  },

  fromPartial(object: DeepPartial<HalvingMinter>): HalvingMinter {
    const message = { ...baseHalvingMinter } as HalvingMinter;
    if (object.new_coins_mint !== undefined && object.new_coins_mint !== null) {
      message.new_coins_mint = object.new_coins_mint;
    } else {
      message.new_coins_mint = 0;
    }
    if (
      object.blocks_per_year !== undefined &&
      object.blocks_per_year !== null
    ) {
      message.blocks_per_year = object.blocks_per_year;
    } else {
      message.blocks_per_year = 0;
    }
    if (object.mint_denom !== undefined && object.mint_denom !== null) {
      message.mint_denom = object.mint_denom;
    } else {
      message.mint_denom = "";
    }
    return message;
  },
};

const baseMinter: object = {};

export const Minter = {
  encode(message: Minter, writer: Writer = Writer.create()): Writer {
    if (message.start !== undefined) {
      Timestamp.encode(
        toTimestamp(message.start),
        writer.uint32(10).fork()
      ).ldelim();
    }
    for (const v of message.periods) {
      MintingPeriod.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Minter {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMinter } as Minter;
    message.periods = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.start = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 2:
          message.periods.push(MintingPeriod.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Minter {
    const message = { ...baseMinter } as Minter;
    message.periods = [];
    if (object.start !== undefined && object.start !== null) {
      message.start = fromJsonTimestamp(object.start);
    } else {
      message.start = undefined;
    }
    if (object.periods !== undefined && object.periods !== null) {
      for (const e of object.periods) {
        message.periods.push(MintingPeriod.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: Minter): unknown {
    const obj: any = {};
    message.start !== undefined &&
      (obj.start =
        message.start !== undefined ? message.start.toISOString() : null);
    if (message.periods) {
      obj.periods = message.periods.map((e) =>
        e ? MintingPeriod.toJSON(e) : undefined
      );
    } else {
      obj.periods = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Minter>): Minter {
    const message = { ...baseMinter } as Minter;
    message.periods = [];
    if (object.start !== undefined && object.start !== null) {
      message.start = object.start;
    } else {
      message.start = undefined;
    }
    if (object.periods !== undefined && object.periods !== null) {
      for (const e of object.periods) {
        message.periods.push(MintingPeriod.fromPartial(e));
      }
    }
    return message;
  },
};

const baseMintingPeriod: object = { ordering_id: 0, type: 0 };

export const MintingPeriod = {
  encode(message: MintingPeriod, writer: Writer = Writer.create()): Writer {
    if (message.ordering_id !== 0) {
      writer.uint32(8).int32(message.ordering_id);
    }
    if (message.period_end !== undefined) {
      Timestamp.encode(
        toTimestamp(message.period_end),
        writer.uint32(18).fork()
      ).ldelim();
    }
    if (message.type !== 0) {
      writer.uint32(24).int32(message.type);
    }
    if (message.time_linear_minter !== undefined) {
      TimeLinearMinter.encode(
        message.time_linear_minter,
        writer.uint32(34).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MintingPeriod {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMintingPeriod } as MintingPeriod;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ordering_id = reader.int32();
          break;
        case 2:
          message.period_end = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.type = reader.int32() as any;
          break;
        case 4:
          message.time_linear_minter = TimeLinearMinter.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MintingPeriod {
    const message = { ...baseMintingPeriod } as MintingPeriod;
    if (object.ordering_id !== undefined && object.ordering_id !== null) {
      message.ordering_id = Number(object.ordering_id);
    } else {
      message.ordering_id = 0;
    }
    if (object.period_end !== undefined && object.period_end !== null) {
      message.period_end = fromJsonTimestamp(object.period_end);
    } else {
      message.period_end = undefined;
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = mintingPeriod_MinterTypeFromJSON(object.type);
    } else {
      message.type = 0;
    }
    if (
      object.time_linear_minter !== undefined &&
      object.time_linear_minter !== null
    ) {
      message.time_linear_minter = TimeLinearMinter.fromJSON(
        object.time_linear_minter
      );
    } else {
      message.time_linear_minter = undefined;
    }
    return message;
  },

  toJSON(message: MintingPeriod): unknown {
    const obj: any = {};
    message.ordering_id !== undefined &&
      (obj.ordering_id = message.ordering_id);
    message.period_end !== undefined &&
      (obj.period_end =
        message.period_end !== undefined
          ? message.period_end.toISOString()
          : null);
    message.type !== undefined &&
      (obj.type = mintingPeriod_MinterTypeToJSON(message.type));
    message.time_linear_minter !== undefined &&
      (obj.time_linear_minter = message.time_linear_minter
        ? TimeLinearMinter.toJSON(message.time_linear_minter)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MintingPeriod>): MintingPeriod {
    const message = { ...baseMintingPeriod } as MintingPeriod;
    if (object.ordering_id !== undefined && object.ordering_id !== null) {
      message.ordering_id = object.ordering_id;
    } else {
      message.ordering_id = 0;
    }
    if (object.period_end !== undefined && object.period_end !== null) {
      message.period_end = object.period_end;
    } else {
      message.period_end = undefined;
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = 0;
    }
    if (
      object.time_linear_minter !== undefined &&
      object.time_linear_minter !== null
    ) {
      message.time_linear_minter = TimeLinearMinter.fromPartial(
        object.time_linear_minter
      );
    } else {
      message.time_linear_minter = undefined;
    }
    return message;
  },
};

const baseTimeLinearMinter: object = { amount: "" };

export const TimeLinearMinter = {
  encode(message: TimeLinearMinter, writer: Writer = Writer.create()): Writer {
    if (message.amount !== "") {
      writer.uint32(10).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): TimeLinearMinter {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseTimeLinearMinter } as TimeLinearMinter;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): TimeLinearMinter {
    const message = { ...baseTimeLinearMinter } as TimeLinearMinter;
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: TimeLinearMinter): unknown {
    const obj: any = {};
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<TimeLinearMinter>): TimeLinearMinter {
    const message = { ...baseTimeLinearMinter } as TimeLinearMinter;
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const baseMinterState: object = { current_ordering_id: 0, amount_minted: "" };

export const MinterState = {
  encode(message: MinterState, writer: Writer = Writer.create()): Writer {
    if (message.current_ordering_id !== 0) {
      writer.uint32(8).int32(message.current_ordering_id);
    }
    if (message.amount_minted !== "") {
      writer.uint32(18).string(message.amount_minted);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MinterState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMinterState } as MinterState;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.current_ordering_id = reader.int32();
          break;
        case 2:
          message.amount_minted = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MinterState {
    const message = { ...baseMinterState } as MinterState;
    if (
      object.current_ordering_id !== undefined &&
      object.current_ordering_id !== null
    ) {
      message.current_ordering_id = Number(object.current_ordering_id);
    } else {
      message.current_ordering_id = 0;
    }
    if (object.amount_minted !== undefined && object.amount_minted !== null) {
      message.amount_minted = String(object.amount_minted);
    } else {
      message.amount_minted = "";
    }
    return message;
  },

  toJSON(message: MinterState): unknown {
    const obj: any = {};
    message.current_ordering_id !== undefined &&
      (obj.current_ordering_id = message.current_ordering_id);
    message.amount_minted !== undefined &&
      (obj.amount_minted = message.amount_minted);
    return obj;
  },

  fromPartial(object: DeepPartial<MinterState>): MinterState {
    const message = { ...baseMinterState } as MinterState;
    if (
      object.current_ordering_id !== undefined &&
      object.current_ordering_id !== null
    ) {
      message.current_ordering_id = object.current_ordering_id;
    } else {
      message.current_ordering_id = 0;
    }
    if (object.amount_minted !== undefined && object.amount_minted !== null) {
      message.amount_minted = object.amount_minted;
    } else {
      message.amount_minted = "";
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

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
