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
  periodEnd: Date | undefined;
  /**
   * types:
   *   NO_MINTING;
   *   TIME_LINEAR_MINTER;
   *   PERIODIC_REDUCTION_MINTER;
   */
  type: string;
  timeLinearMinter: TimeLinearMinter | undefined;
  periodicReductionMinter: PeriodicReductionMinter | undefined;
}

export interface TimeLinearMinter {
  amount: string;
}

export interface PeriodicReductionMinter {
  /** mint_period in seconds */
  mintPeriod: number;
  mintAmount: string;
  reductionPeriodLength: number;
  reductionFactor: string;
}

export interface MinterState {
  position: number;
  amountMinted: string;
  remainderToMint: string;
  lastMintBlockTime: Date | undefined;
  remainderFromPreviousPeriod: string;
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
    if (message.periodEnd !== undefined) {
      Timestamp.encode(
        toTimestamp(message.periodEnd),
        writer.uint32(18).fork()
      ).ldelim();
    }
    if (message.type !== "") {
      writer.uint32(26).string(message.type);
    }
    if (message.timeLinearMinter !== undefined) {
      TimeLinearMinter.encode(
        message.timeLinearMinter,
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.periodicReductionMinter !== undefined) {
      PeriodicReductionMinter.encode(
        message.periodicReductionMinter,
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
          message.periodEnd = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.type = reader.string();
          break;
        case 4:
          message.timeLinearMinter = TimeLinearMinter.decode(
            reader,
            reader.uint32()
          );
          break;
        case 5:
          message.periodicReductionMinter = PeriodicReductionMinter.decode(
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
    if (object.periodEnd !== undefined && object.periodEnd !== null) {
      message.periodEnd = fromJsonTimestamp(object.periodEnd);
    } else {
      message.periodEnd = undefined;
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = String(object.type);
    } else {
      message.type = "";
    }
    if (
      object.timeLinearMinter !== undefined &&
      object.timeLinearMinter !== null
    ) {
      message.timeLinearMinter = TimeLinearMinter.fromJSON(
        object.timeLinearMinter
      );
    } else {
      message.timeLinearMinter = undefined;
    }
    if (
      object.periodicReductionMinter !== undefined &&
      object.periodicReductionMinter !== null
    ) {
      message.periodicReductionMinter = PeriodicReductionMinter.fromJSON(
        object.periodicReductionMinter
      );
    } else {
      message.periodicReductionMinter = undefined;
    }
    return message;
  },

  toJSON(message: MintingPeriod): unknown {
    const obj: any = {};
    message.position !== undefined && (obj.position = message.position);
    message.periodEnd !== undefined &&
      (obj.periodEnd =
        message.periodEnd !== undefined
          ? message.periodEnd.toISOString()
          : null);
    message.type !== undefined && (obj.type = message.type);
    message.timeLinearMinter !== undefined &&
      (obj.timeLinearMinter = message.timeLinearMinter
        ? TimeLinearMinter.toJSON(message.timeLinearMinter)
        : undefined);
    message.periodicReductionMinter !== undefined &&
      (obj.periodicReductionMinter = message.periodicReductionMinter
        ? PeriodicReductionMinter.toJSON(message.periodicReductionMinter)
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
    if (object.periodEnd !== undefined && object.periodEnd !== null) {
      message.periodEnd = object.periodEnd;
    } else {
      message.periodEnd = undefined;
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = "";
    }
    if (
      object.timeLinearMinter !== undefined &&
      object.timeLinearMinter !== null
    ) {
      message.timeLinearMinter = TimeLinearMinter.fromPartial(
        object.timeLinearMinter
      );
    } else {
      message.timeLinearMinter = undefined;
    }
    if (
      object.periodicReductionMinter !== undefined &&
      object.periodicReductionMinter !== null
    ) {
      message.periodicReductionMinter = PeriodicReductionMinter.fromPartial(
        object.periodicReductionMinter
      );
    } else {
      message.periodicReductionMinter = undefined;
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
  mintPeriod: 0,
  mintAmount: "",
  reductionPeriodLength: 0,
  reductionFactor: "",
};

export const PeriodicReductionMinter = {
  encode(
    message: PeriodicReductionMinter,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.mintPeriod !== 0) {
      writer.uint32(8).int32(message.mintPeriod);
    }
    if (message.mintAmount !== "") {
      writer.uint32(18).string(message.mintAmount);
    }
    if (message.reductionPeriodLength !== 0) {
      writer.uint32(24).int32(message.reductionPeriodLength);
    }
    if (message.reductionFactor !== "") {
      writer.uint32(34).string(message.reductionFactor);
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
          message.mintPeriod = reader.int32();
          break;
        case 2:
          message.mintAmount = reader.string();
          break;
        case 3:
          message.reductionPeriodLength = reader.int32();
          break;
        case 4:
          message.reductionFactor = reader.string();
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
    if (object.mintPeriod !== undefined && object.mintPeriod !== null) {
      message.mintPeriod = Number(object.mintPeriod);
    } else {
      message.mintPeriod = 0;
    }
    if (object.mintAmount !== undefined && object.mintAmount !== null) {
      message.mintAmount = String(object.mintAmount);
    } else {
      message.mintAmount = "";
    }
    if (
      object.reductionPeriodLength !== undefined &&
      object.reductionPeriodLength !== null
    ) {
      message.reductionPeriodLength = Number(object.reductionPeriodLength);
    } else {
      message.reductionPeriodLength = 0;
    }
    if (
      object.reductionFactor !== undefined &&
      object.reductionFactor !== null
    ) {
      message.reductionFactor = String(object.reductionFactor);
    } else {
      message.reductionFactor = "";
    }
    return message;
  },

  toJSON(message: PeriodicReductionMinter): unknown {
    const obj: any = {};
    message.mintPeriod !== undefined && (obj.mintPeriod = message.mintPeriod);
    message.mintAmount !== undefined && (obj.mintAmount = message.mintAmount);
    message.reductionPeriodLength !== undefined &&
      (obj.reductionPeriodLength = message.reductionPeriodLength);
    message.reductionFactor !== undefined &&
      (obj.reductionFactor = message.reductionFactor);
    return obj;
  },

  fromPartial(
    object: DeepPartial<PeriodicReductionMinter>
  ): PeriodicReductionMinter {
    const message = {
      ...basePeriodicReductionMinter,
    } as PeriodicReductionMinter;
    if (object.mintPeriod !== undefined && object.mintPeriod !== null) {
      message.mintPeriod = object.mintPeriod;
    } else {
      message.mintPeriod = 0;
    }
    if (object.mintAmount !== undefined && object.mintAmount !== null) {
      message.mintAmount = object.mintAmount;
    } else {
      message.mintAmount = "";
    }
    if (
      object.reductionPeriodLength !== undefined &&
      object.reductionPeriodLength !== null
    ) {
      message.reductionPeriodLength = object.reductionPeriodLength;
    } else {
      message.reductionPeriodLength = 0;
    }
    if (
      object.reductionFactor !== undefined &&
      object.reductionFactor !== null
    ) {
      message.reductionFactor = object.reductionFactor;
    } else {
      message.reductionFactor = "";
    }
    return message;
  },
};

const baseMinterState: object = {
  position: 0,
  amountMinted: "",
  remainderToMint: "",
  remainderFromPreviousPeriod: "",
};

export const MinterState = {
  encode(message: MinterState, writer: Writer = Writer.create()): Writer {
    if (message.position !== 0) {
      writer.uint32(8).int32(message.position);
    }
    if (message.amountMinted !== "") {
      writer.uint32(18).string(message.amountMinted);
    }
    if (message.remainderToMint !== "") {
      writer.uint32(26).string(message.remainderToMint);
    }
    if (message.lastMintBlockTime !== undefined) {
      Timestamp.encode(
        toTimestamp(message.lastMintBlockTime),
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.remainderFromPreviousPeriod !== "") {
      writer.uint32(42).string(message.remainderFromPreviousPeriod);
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
          message.amountMinted = reader.string();
          break;
        case 3:
          message.remainderToMint = reader.string();
          break;
        case 4:
          message.lastMintBlockTime = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 5:
          message.remainderFromPreviousPeriod = reader.string();
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
    if (object.amountMinted !== undefined && object.amountMinted !== null) {
      message.amountMinted = String(object.amountMinted);
    } else {
      message.amountMinted = "";
    }
    if (
      object.remainderToMint !== undefined &&
      object.remainderToMint !== null
    ) {
      message.remainderToMint = String(object.remainderToMint);
    } else {
      message.remainderToMint = "";
    }
    if (
      object.lastMintBlockTime !== undefined &&
      object.lastMintBlockTime !== null
    ) {
      message.lastMintBlockTime = fromJsonTimestamp(object.lastMintBlockTime);
    } else {
      message.lastMintBlockTime = undefined;
    }
    if (
      object.remainderFromPreviousPeriod !== undefined &&
      object.remainderFromPreviousPeriod !== null
    ) {
      message.remainderFromPreviousPeriod = String(
        object.remainderFromPreviousPeriod
      );
    } else {
      message.remainderFromPreviousPeriod = "";
    }
    return message;
  },

  toJSON(message: MinterState): unknown {
    const obj: any = {};
    message.position !== undefined && (obj.position = message.position);
    message.amountMinted !== undefined &&
      (obj.amountMinted = message.amountMinted);
    message.remainderToMint !== undefined &&
      (obj.remainderToMint = message.remainderToMint);
    message.lastMintBlockTime !== undefined &&
      (obj.lastMintBlockTime =
        message.lastMintBlockTime !== undefined
          ? message.lastMintBlockTime.toISOString()
          : null);
    message.remainderFromPreviousPeriod !== undefined &&
      (obj.remainderFromPreviousPeriod = message.remainderFromPreviousPeriod);
    return obj;
  },

  fromPartial(object: DeepPartial<MinterState>): MinterState {
    const message = { ...baseMinterState } as MinterState;
    if (object.position !== undefined && object.position !== null) {
      message.position = object.position;
    } else {
      message.position = 0;
    }
    if (object.amountMinted !== undefined && object.amountMinted !== null) {
      message.amountMinted = object.amountMinted;
    } else {
      message.amountMinted = "";
    }
    if (
      object.remainderToMint !== undefined &&
      object.remainderToMint !== null
    ) {
      message.remainderToMint = object.remainderToMint;
    } else {
      message.remainderToMint = "";
    }
    if (
      object.lastMintBlockTime !== undefined &&
      object.lastMintBlockTime !== null
    ) {
      message.lastMintBlockTime = object.lastMintBlockTime;
    } else {
      message.lastMintBlockTime = undefined;
    }
    if (
      object.remainderFromPreviousPeriod !== undefined &&
      object.remainderFromPreviousPeriod !== null
    ) {
      message.remainderFromPreviousPeriod = object.remainderFromPreviousPeriod;
    } else {
      message.remainderFromPreviousPeriod = "";
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
