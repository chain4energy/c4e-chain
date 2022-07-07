/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import { Duration } from "../google/protobuf/duration";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeminter";

export interface Minter {
  start: Date | undefined;
  periods: MintingPeriod[];
}

export interface MintingPeriod {
  position: number;
  period_end: Date | undefined;
  type: MintingPeriod_MinterType;
  time_linear_minter: TimeLinearMinter | undefined;
  periodic_reduction_minter: PeriodicReductionMinter | undefined;
}

export enum MintingPeriod_MinterType {
  NO_MINTING = 0,
  TIME_LINEAR_MINTER = 1,
  PERIODIC_REDUCTION_MINTER = 2,
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
    case 2:
    case "PERIODIC_REDUCTION_MINTER":
      return MintingPeriod_MinterType.PERIODIC_REDUCTION_MINTER;
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
    case MintingPeriod_MinterType.PERIODIC_REDUCTION_MINTER:
      return "PERIODIC_REDUCTION_MINTER";
    default:
      return "UNKNOWN";
  }
}

export interface TimeLinearMinter {
  amount: string;
}

export interface PeriodicReductionMinter {
  mint_period: Duration | undefined;
  mint_amount: string;
  reduction_period_length: number;
  reduction_factor: string;
}

export interface MinterState {
  current_position: number;
  amount_minted: string;
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

const baseMintingPeriod: object = { position: 0, type: 0 };

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
    if (message.type !== 0) {
      writer.uint32(24).int32(message.type);
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
          message.type = reader.int32() as any;
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
    message.type !== undefined &&
      (obj.type = mintingPeriod_MinterTypeToJSON(message.type));
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
  mint_amount: "",
  reduction_period_length: 0,
  reduction_factor: "",
};

export const PeriodicReductionMinter = {
  encode(
    message: PeriodicReductionMinter,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.mint_period !== undefined) {
      Duration.encode(message.mint_period, writer.uint32(10).fork()).ldelim();
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
          message.mint_period = Duration.decode(reader, reader.uint32());
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
      message.mint_period = Duration.fromJSON(object.mint_period);
    } else {
      message.mint_period = undefined;
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
      (obj.mint_period = message.mint_period
        ? Duration.toJSON(message.mint_period)
        : undefined);
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
      message.mint_period = Duration.fromPartial(object.mint_period);
    } else {
      message.mint_period = undefined;
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

const baseMinterState: object = { current_position: 0, amount_minted: "" };

export const MinterState = {
  encode(message: MinterState, writer: Writer = Writer.create()): Writer {
    if (message.current_position !== 0) {
      writer.uint32(8).int32(message.current_position);
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
          message.current_position = reader.int32();
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
      object.current_position !== undefined &&
      object.current_position !== null
    ) {
      message.current_position = Number(object.current_position);
    } else {
      message.current_position = 0;
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
    message.current_position !== undefined &&
      (obj.current_position = message.current_position);
    message.amount_minted !== undefined &&
      (obj.amount_minted = message.amount_minted);
    return obj;
  },

  fromPartial(object: DeepPartial<MinterState>): MinterState {
    const message = { ...baseMinterState } as MinterState;
    if (
      object.current_position !== undefined &&
      object.current_position !== null
    ) {
      message.current_position = object.current_position;
    } else {
      message.current_position = 0;
    }
    if (object.amount_minted !== undefined && object.amount_minted !== null) {
      message.amount_minted = object.amount_minted;
    } else {
      message.amount_minted = "";
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
