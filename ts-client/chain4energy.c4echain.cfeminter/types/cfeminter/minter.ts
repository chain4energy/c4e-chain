/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Timestamp } from "../google/protobuf/timestamp";

export const protobufPackage = "chain4energy.c4echain.cfeminter";

export interface Minter {
  start: Date | undefined;
  periods: MintingPeriod[];
}

export interface MintingPeriod {
  position: number;
  periodEnd:
    | Date
    | undefined;
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

function createBaseMinter(): Minter {
  return { start: undefined, periods: [] };
}

export const Minter = {
  encode(message: Minter, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.start !== undefined) {
      Timestamp.encode(toTimestamp(message.start), writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.periods) {
      MintingPeriod.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Minter {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMinter();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.start = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
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
    return {
      start: isSet(object.start) ? fromJsonTimestamp(object.start) : undefined,
      periods: Array.isArray(object?.periods) ? object.periods.map((e: any) => MintingPeriod.fromJSON(e)) : [],
    };
  },

  toJSON(message: Minter): unknown {
    const obj: any = {};
    message.start !== undefined && (obj.start = message.start.toISOString());
    if (message.periods) {
      obj.periods = message.periods.map((e) => e ? MintingPeriod.toJSON(e) : undefined);
    } else {
      obj.periods = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Minter>, I>>(object: I): Minter {
    const message = createBaseMinter();
    message.start = object.start ?? undefined;
    message.periods = object.periods?.map((e) => MintingPeriod.fromPartial(e)) || [];
    return message;
  },
};

function createBaseMintingPeriod(): MintingPeriod {
  return {
    position: 0,
    periodEnd: undefined,
    type: "",
    timeLinearMinter: undefined,
    periodicReductionMinter: undefined,
  };
}

export const MintingPeriod = {
  encode(message: MintingPeriod, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.position !== 0) {
      writer.uint32(8).int32(message.position);
    }
    if (message.periodEnd !== undefined) {
      Timestamp.encode(toTimestamp(message.periodEnd), writer.uint32(18).fork()).ldelim();
    }
    if (message.type !== "") {
      writer.uint32(26).string(message.type);
    }
    if (message.timeLinearMinter !== undefined) {
      TimeLinearMinter.encode(message.timeLinearMinter, writer.uint32(34).fork()).ldelim();
    }
    if (message.periodicReductionMinter !== undefined) {
      PeriodicReductionMinter.encode(message.periodicReductionMinter, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MintingPeriod {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMintingPeriod();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.position = reader.int32();
          break;
        case 2:
          message.periodEnd = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 3:
          message.type = reader.string();
          break;
        case 4:
          message.timeLinearMinter = TimeLinearMinter.decode(reader, reader.uint32());
          break;
        case 5:
          message.periodicReductionMinter = PeriodicReductionMinter.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MintingPeriod {
    return {
      position: isSet(object.position) ? Number(object.position) : 0,
      periodEnd: isSet(object.periodEnd) ? fromJsonTimestamp(object.periodEnd) : undefined,
      type: isSet(object.type) ? String(object.type) : "",
      timeLinearMinter: isSet(object.timeLinearMinter) ? TimeLinearMinter.fromJSON(object.timeLinearMinter) : undefined,
      periodicReductionMinter: isSet(object.periodicReductionMinter)
        ? PeriodicReductionMinter.fromJSON(object.periodicReductionMinter)
        : undefined,
    };
  },

  toJSON(message: MintingPeriod): unknown {
    const obj: any = {};
    message.position !== undefined && (obj.position = Math.round(message.position));
    message.periodEnd !== undefined && (obj.periodEnd = message.periodEnd.toISOString());
    message.type !== undefined && (obj.type = message.type);
    message.timeLinearMinter !== undefined && (obj.timeLinearMinter = message.timeLinearMinter
      ? TimeLinearMinter.toJSON(message.timeLinearMinter)
      : undefined);
    message.periodicReductionMinter !== undefined && (obj.periodicReductionMinter = message.periodicReductionMinter
      ? PeriodicReductionMinter.toJSON(message.periodicReductionMinter)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MintingPeriod>, I>>(object: I): MintingPeriod {
    const message = createBaseMintingPeriod();
    message.position = object.position ?? 0;
    message.periodEnd = object.periodEnd ?? undefined;
    message.type = object.type ?? "";
    message.timeLinearMinter = (object.timeLinearMinter !== undefined && object.timeLinearMinter !== null)
      ? TimeLinearMinter.fromPartial(object.timeLinearMinter)
      : undefined;
    message.periodicReductionMinter =
      (object.periodicReductionMinter !== undefined && object.periodicReductionMinter !== null)
        ? PeriodicReductionMinter.fromPartial(object.periodicReductionMinter)
        : undefined;
    return message;
  },
};

function createBaseTimeLinearMinter(): TimeLinearMinter {
  return { amount: "" };
}

export const TimeLinearMinter = {
  encode(message: TimeLinearMinter, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.amount !== "") {
      writer.uint32(10).string(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): TimeLinearMinter {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTimeLinearMinter();
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
    return { amount: isSet(object.amount) ? String(object.amount) : "" };
  },

  toJSON(message: TimeLinearMinter): unknown {
    const obj: any = {};
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<TimeLinearMinter>, I>>(object: I): TimeLinearMinter {
    const message = createBaseTimeLinearMinter();
    message.amount = object.amount ?? "";
    return message;
  },
};

function createBasePeriodicReductionMinter(): PeriodicReductionMinter {
  return { mintPeriod: 0, mintAmount: "", reductionPeriodLength: 0, reductionFactor: "" };
}

export const PeriodicReductionMinter = {
  encode(message: PeriodicReductionMinter, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): PeriodicReductionMinter {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePeriodicReductionMinter();
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
    return {
      mintPeriod: isSet(object.mintPeriod) ? Number(object.mintPeriod) : 0,
      mintAmount: isSet(object.mintAmount) ? String(object.mintAmount) : "",
      reductionPeriodLength: isSet(object.reductionPeriodLength) ? Number(object.reductionPeriodLength) : 0,
      reductionFactor: isSet(object.reductionFactor) ? String(object.reductionFactor) : "",
    };
  },

  toJSON(message: PeriodicReductionMinter): unknown {
    const obj: any = {};
    message.mintPeriod !== undefined && (obj.mintPeriod = Math.round(message.mintPeriod));
    message.mintAmount !== undefined && (obj.mintAmount = message.mintAmount);
    message.reductionPeriodLength !== undefined
      && (obj.reductionPeriodLength = Math.round(message.reductionPeriodLength));
    message.reductionFactor !== undefined && (obj.reductionFactor = message.reductionFactor);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PeriodicReductionMinter>, I>>(object: I): PeriodicReductionMinter {
    const message = createBasePeriodicReductionMinter();
    message.mintPeriod = object.mintPeriod ?? 0;
    message.mintAmount = object.mintAmount ?? "";
    message.reductionPeriodLength = object.reductionPeriodLength ?? 0;
    message.reductionFactor = object.reductionFactor ?? "";
    return message;
  },
};

function createBaseMinterState(): MinterState {
  return {
    position: 0,
    amountMinted: "",
    remainderToMint: "",
    lastMintBlockTime: undefined,
    remainderFromPreviousPeriod: "",
  };
}

export const MinterState = {
  encode(message: MinterState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
      Timestamp.encode(toTimestamp(message.lastMintBlockTime), writer.uint32(34).fork()).ldelim();
    }
    if (message.remainderFromPreviousPeriod !== "") {
      writer.uint32(42).string(message.remainderFromPreviousPeriod);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MinterState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMinterState();
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
          message.lastMintBlockTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
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
    return {
      position: isSet(object.position) ? Number(object.position) : 0,
      amountMinted: isSet(object.amountMinted) ? String(object.amountMinted) : "",
      remainderToMint: isSet(object.remainderToMint) ? String(object.remainderToMint) : "",
      lastMintBlockTime: isSet(object.lastMintBlockTime) ? fromJsonTimestamp(object.lastMintBlockTime) : undefined,
      remainderFromPreviousPeriod: isSet(object.remainderFromPreviousPeriod)
        ? String(object.remainderFromPreviousPeriod)
        : "",
    };
  },

  toJSON(message: MinterState): unknown {
    const obj: any = {};
    message.position !== undefined && (obj.position = Math.round(message.position));
    message.amountMinted !== undefined && (obj.amountMinted = message.amountMinted);
    message.remainderToMint !== undefined && (obj.remainderToMint = message.remainderToMint);
    message.lastMintBlockTime !== undefined && (obj.lastMintBlockTime = message.lastMintBlockTime.toISOString());
    message.remainderFromPreviousPeriod !== undefined
      && (obj.remainderFromPreviousPeriod = message.remainderFromPreviousPeriod);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MinterState>, I>>(object: I): MinterState {
    const message = createBaseMinterState();
    message.position = object.position ?? 0;
    message.amountMinted = object.amountMinted ?? "";
    message.remainderToMint = object.remainderToMint ?? "";
    message.lastMintBlockTime = object.lastMintBlockTime ?? undefined;
    message.remainderFromPreviousPeriod = object.remainderFromPreviousPeriod ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
