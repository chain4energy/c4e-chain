/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeminter";

export interface Minter {
  start: Date | undefined;
  periods: MintingPeriod[];
}

export interface MintingPeriod {
  position: number;
  period_end: Date | undefined;
  /**
   * types:
   *   NO_MINTING;
   *   TIME_LINEAR_MINTER;
   *   PERIODIC_REDUCTION_MINTER;
   */
  type: string;
  time_linear_minter: TimeLinearMinter | undefined;
  periodic_reduction_minter: PeriodicReductionMinter | undefined;
}

export interface TimeLinearMinter {
  amount: string;
}

export interface PeriodicReductionMinter {
  /** mint_period in seconds */
  mint_period: number;
  mint_amount: string;
  reduction_period_length: number;
  reduction_factor: string;
}

export interface MinterState {
  position: number;
  amount_minted: string;
  remainder_to_mint: string;
  last_mint_block_time: Date | undefined;
  remainder_from_previous_period: string;
}

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

const baseMintingPeriod: object = { position: 0, type: "" };

export const MintingPeriod = {
  encode(message: MintingPeriod, writer: Writer = Writer.create()): Writer {
    if (message.position !== 0) {
      writer.uint32(8).int32(message.position);
    }
    if (message.period_end !== undefined) {
      Timestamp.encode(
        toTimestamp(message.period_end),
        writer.uint32(18).fork()
      ).ldelim();
    }
    if (message.type !== "") {
      writer.uint32(26).string(message.type);
    }
    if (message.time_linear_minter !== undefined) {
      TimeLinearMinter.encode(
        message.time_linear_minter,
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.periodic_reduction_minter !== undefined) {
      PeriodicReductionMinter.encode(
        message.periodic_reduction_minter,
        writer.uint32(42).fork()
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
          message.position = reader.int32();
          break;
        case 2:
          message.period_end = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.type = reader.string();
          break;
        case 4:
          message.time_linear_minter = TimeLinearMinter.decode(
            reader,
            reader.uint32()
          );
          break;
        case 5:
          message.periodic_reduction_minter = PeriodicReductionMinter.decode(
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
    if (object.position !== undefined && object.position !== null) {
      message.position = Number(object.position);
    } else {
      message.position = 0;
    }
    if (object.period_end !== undefined && object.period_end !== null) {
      message.period_end = fromJsonTimestamp(object.period_end);
    } else {
      message.period_end = undefined;
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = String(object.type);
    } else {
      message.type = "";
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
    if (
      object.periodic_reduction_minter !== undefined &&
      object.periodic_reduction_minter !== null
    ) {
      message.periodic_reduction_minter = PeriodicReductionMinter.fromJSON(
        object.periodic_reduction_minter
      );
    } else {
      message.periodic_reduction_minter = undefined;
    }
    return message;
  },

  toJSON(message: MintingPeriod): unknown {
    const obj: any = {};
    message.position !== undefined && (obj.position = message.position);
    message.period_end !== undefined &&
      (obj.period_end =
        message.period_end !== undefined
          ? message.period_end.toISOString()
          : null);
    message.type !== undefined && (obj.type = message.type);
    message.time_linear_minter !== undefined &&
      (obj.time_linear_minter = message.time_linear_minter
        ? TimeLinearMinter.toJSON(message.time_linear_minter)
        : undefined);
    message.periodic_reduction_minter !== undefined &&
      (obj.periodic_reduction_minter = message.periodic_reduction_minter
        ? PeriodicReductionMinter.toJSON(message.periodic_reduction_minter)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<MintingPeriod>): MintingPeriod {
    const message = { ...baseMintingPeriod } as MintingPeriod;
    if (object.position !== undefined && object.position !== null) {
      message.position = object.position;
    } else {
      message.position = 0;
    }
    if (object.period_end !== undefined && object.period_end !== null) {
      message.period_end = object.period_end;
    } else {
      message.period_end = undefined;
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = "";
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
    if (
      object.periodic_reduction_minter !== undefined &&
      object.periodic_reduction_minter !== null
    ) {
      message.periodic_reduction_minter = PeriodicReductionMinter.fromPartial(
        object.periodic_reduction_minter
      );
    } else {
      message.periodic_reduction_minter = undefined;
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

const basePeriodicReductionMinter: object = {
  mint_period: 0,
  mint_amount: "",
  reduction_period_length: 0,
  reduction_factor: "",
};

export const PeriodicReductionMinter = {
  encode(
    message: PeriodicReductionMinter,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.mint_period !== 0) {
      writer.uint32(8).int32(message.mint_period);
    }
    if (message.mint_amount !== "") {
      writer.uint32(18).string(message.mint_amount);
    }
    if (message.reduction_period_length !== 0) {
      writer.uint32(24).int32(message.reduction_period_length);
    }
    if (message.reduction_factor !== "") {
      writer.uint32(34).string(message.reduction_factor);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): PeriodicReductionMinter {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...basePeriodicReductionMinter,
    } as PeriodicReductionMinter;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.mint_period = reader.int32();
          break;
        case 2:
          message.mint_amount = reader.string();
          break;
        case 3:
          message.reduction_period_length = reader.int32();
          break;
        case 4:
          message.reduction_factor = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PeriodicReductionMinter {
    const message = {
      ...basePeriodicReductionMinter,
    } as PeriodicReductionMinter;
    if (object.mint_period !== undefined && object.mint_period !== null) {
      message.mint_period = Number(object.mint_period);
    } else {
      message.mint_period = 0;
    }
    if (object.mint_amount !== undefined && object.mint_amount !== null) {
      message.mint_amount = String(object.mint_amount);
    } else {
      message.mint_amount = "";
    }
    if (
      object.reduction_period_length !== undefined &&
      object.reduction_period_length !== null
    ) {
      message.reduction_period_length = Number(object.reduction_period_length);
    } else {
      message.reduction_period_length = 0;
    }
    if (
      object.reduction_factor !== undefined &&
      object.reduction_factor !== null
    ) {
      message.reduction_factor = String(object.reduction_factor);
    } else {
      message.reduction_factor = "";
    }
    return message;
  },

  toJSON(message: PeriodicReductionMinter): unknown {
    const obj: any = {};
    message.mint_period !== undefined &&
      (obj.mint_period = message.mint_period);
    message.mint_amount !== undefined &&
      (obj.mint_amount = message.mint_amount);
    message.reduction_period_length !== undefined &&
      (obj.reduction_period_length = message.reduction_period_length);
    message.reduction_factor !== undefined &&
      (obj.reduction_factor = message.reduction_factor);
    return obj;
  },

  fromPartial(
    object: DeepPartial<PeriodicReductionMinter>
  ): PeriodicReductionMinter {
    const message = {
      ...basePeriodicReductionMinter,
    } as PeriodicReductionMinter;
    if (object.mint_period !== undefined && object.mint_period !== null) {
      message.mint_period = object.mint_period;
    } else {
      message.mint_period = 0;
    }
    if (object.mint_amount !== undefined && object.mint_amount !== null) {
      message.mint_amount = object.mint_amount;
    } else {
      message.mint_amount = "";
    }
    if (
      object.reduction_period_length !== undefined &&
      object.reduction_period_length !== null
    ) {
      message.reduction_period_length = object.reduction_period_length;
    } else {
      message.reduction_period_length = 0;
    }
    if (
      object.reduction_factor !== undefined &&
      object.reduction_factor !== null
    ) {
      message.reduction_factor = object.reduction_factor;
    } else {
      message.reduction_factor = "";
    }
    return message;
  },
};

const baseMinterState: object = {
  position: 0,
  amount_minted: "",
  remainder_to_mint: "",
  remainder_from_previous_period: "",
};

export const MinterState = {
  encode(message: MinterState, writer: Writer = Writer.create()): Writer {
    if (message.position !== 0) {
      writer.uint32(8).int32(message.position);
    }
    if (message.amount_minted !== "") {
      writer.uint32(18).string(message.amount_minted);
    }
    if (message.remainder_to_mint !== "") {
      writer.uint32(26).string(message.remainder_to_mint);
    }
    if (message.last_mint_block_time !== undefined) {
      Timestamp.encode(
        toTimestamp(message.last_mint_block_time),
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.remainder_from_previous_period !== "") {
      writer.uint32(42).string(message.remainder_from_previous_period);
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
          message.position = reader.int32();
          break;
        case 2:
          message.amount_minted = reader.string();
          break;
        case 3:
          message.remainder_to_mint = reader.string();
          break;
        case 4:
          message.last_mint_block_time = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.remainder_from_previous_period = reader.string();
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
    if (object.position !== undefined && object.position !== null) {
      message.position = Number(object.position);
    } else {
      message.position = 0;
    }
    if (object.amount_minted !== undefined && object.amount_minted !== null) {
      message.amount_minted = String(object.amount_minted);
    } else {
      message.amount_minted = "";
    }
    if (
      object.remainder_to_mint !== undefined &&
      object.remainder_to_mint !== null
    ) {
      message.remainder_to_mint = String(object.remainder_to_mint);
    } else {
      message.remainder_to_mint = "";
    }
    if (
      object.last_mint_block_time !== undefined &&
      object.last_mint_block_time !== null
    ) {
      message.last_mint_block_time = fromJsonTimestamp(
        object.last_mint_block_time
      );
    } else {
      message.last_mint_block_time = undefined;
    }
    if (
      object.remainder_from_previous_period !== undefined &&
      object.remainder_from_previous_period !== null
    ) {
      message.remainder_from_previous_period = String(
        object.remainder_from_previous_period
      );
    } else {
      message.remainder_from_previous_period = "";
    }
    return message;
  },

  toJSON(message: MinterState): unknown {
    const obj: any = {};
    message.position !== undefined && (obj.position = message.position);
    message.amount_minted !== undefined &&
      (obj.amount_minted = message.amount_minted);
    message.remainder_to_mint !== undefined &&
      (obj.remainder_to_mint = message.remainder_to_mint);
    message.last_mint_block_time !== undefined &&
      (obj.last_mint_block_time =
        message.last_mint_block_time !== undefined
          ? message.last_mint_block_time.toISOString()
          : null);
    message.remainder_from_previous_period !== undefined &&
      (obj.remainder_from_previous_period =
        message.remainder_from_previous_period);
    return obj;
  },

  fromPartial(object: DeepPartial<MinterState>): MinterState {
    const message = { ...baseMinterState } as MinterState;
    if (object.position !== undefined && object.position !== null) {
      message.position = object.position;
    } else {
      message.position = 0;
    }
    if (object.amount_minted !== undefined && object.amount_minted !== null) {
      message.amount_minted = object.amount_minted;
    } else {
      message.amount_minted = "";
    }
    if (
      object.remainder_to_mint !== undefined &&
      object.remainder_to_mint !== null
    ) {
      message.remainder_to_mint = object.remainder_to_mint;
    } else {
      message.remainder_to_mint = "";
    }
    if (
      object.last_mint_block_time !== undefined &&
      object.last_mint_block_time !== null
    ) {
      message.last_mint_block_time = object.last_mint_block_time;
    } else {
      message.last_mint_block_time = undefined;
    }
    if (
      object.remainder_from_previous_period !== undefined &&
      object.remainder_from_previous_period !== null
    ) {
      message.remainder_from_previous_period =
        object.remainder_from_previous_period;
    } else {
      message.remainder_from_previous_period = "";
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
