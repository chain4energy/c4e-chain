/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Any } from "../../google/protobuf/any";
import { Duration } from "../../google/protobuf/duration";
import { Timestamp } from "../../google/protobuf/timestamp";

export const protobufPackage = "chain4energy.c4echain.cfeminter";

export interface Minter {
  sequenceId: number;
  endTime: Date | undefined;
  config: Any | undefined;
}

export interface NoMinting {
}

export interface LinearMinting {
  amount: string;
}

export interface ExponentialStepMinting {
  stepDuration: Duration | undefined;
  amount: string;
  amountMultiplier: string;
}

export interface MinterState {
  sequenceId: number;
  amountMinted: string;
  remainderToMint: string;
  lastMintBlockTime: Date | undefined;
  remainderFromPreviousMinter: string;
}

function createBaseMinter(): Minter {
  return { sequenceId: 0, endTime: undefined, config: undefined };
}

export const Minter = {
  encode(message: Minter, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sequenceId !== 0) {
      writer.uint32(8).uint32(message.sequenceId);
    }
    if (message.endTime !== undefined) {
      Timestamp.encode(toTimestamp(message.endTime), writer.uint32(18).fork()).ldelim();
    }
    if (message.config !== undefined) {
      Any.encode(message.config, writer.uint32(26).fork()).ldelim();
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
          message.sequenceId = reader.uint32();
          break;
        case 2:
          message.endTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 3:
          message.config = Any.decode(reader, reader.uint32());
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
      sequenceId: isSet(object.sequenceId) ? Number(object.sequenceId) : 0,
      endTime: isSet(object.endTime) ? fromJsonTimestamp(object.endTime) : undefined,
      config: isSet(object.config) ? Any.fromJSON(object.config) : undefined,
    };
  },

  toJSON(message: Minter): unknown {
    const obj: any = {};
    message.sequenceId !== undefined && (obj.sequenceId = Math.round(message.sequenceId));
    message.endTime !== undefined && (obj.endTime = message.endTime.toISOString());
    message.config !== undefined && (obj.config = message.config ? Any.toJSON(message.config) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Minter>, I>>(object: I): Minter {
    const message = createBaseMinter();
    message.sequenceId = object.sequenceId ?? 0;
    message.endTime = object.endTime ?? undefined;
    message.config = (object.config !== undefined && object.config !== null)
      ? Any.fromPartial(object.config)
      : undefined;
    return message;
  },
};

function createBaseNoMinting(): NoMinting {
  return {};
}

export const NoMinting = {
  encode(_: NoMinting, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): NoMinting {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNoMinting();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): NoMinting {
    return {};
  },

  toJSON(_: NoMinting): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<NoMinting>, I>>(_: I): NoMinting {
    const message = createBaseNoMinting();
    return message;
  },
};

function createBaseLinearMinting(): LinearMinting {
  return { amount: "" };
}

export const LinearMinting = {
  encode(message: LinearMinting, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.amount !== "") {
      writer.uint32(10).string(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): LinearMinting {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLinearMinting();
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

  fromJSON(object: any): LinearMinting {
    return { amount: isSet(object.amount) ? String(object.amount) : "" };
  },

  toJSON(message: LinearMinting): unknown {
    const obj: any = {};
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<LinearMinting>, I>>(object: I): LinearMinting {
    const message = createBaseLinearMinting();
    message.amount = object.amount ?? "";
    return message;
  },
};

function createBaseExponentialStepMinting(): ExponentialStepMinting {
  return { stepDuration: undefined, amount: "", amountMultiplier: "" };
}

export const ExponentialStepMinting = {
  encode(message: ExponentialStepMinting, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.stepDuration !== undefined) {
      Duration.encode(message.stepDuration, writer.uint32(10).fork()).ldelim();
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    if (message.amountMultiplier !== "") {
      writer.uint32(26).string(message.amountMultiplier);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ExponentialStepMinting {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseExponentialStepMinting();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.stepDuration = Duration.decode(reader, reader.uint32());
          break;
        case 2:
          message.amount = reader.string();
          break;
        case 3:
          message.amountMultiplier = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ExponentialStepMinting {
    return {
      stepDuration: isSet(object.stepDuration) ? Duration.fromJSON(object.stepDuration) : undefined,
      amount: isSet(object.amount) ? String(object.amount) : "",
      amountMultiplier: isSet(object.amountMultiplier) ? String(object.amountMultiplier) : "",
    };
  },

  toJSON(message: ExponentialStepMinting): unknown {
    const obj: any = {};
    message.stepDuration !== undefined
      && (obj.stepDuration = message.stepDuration ? Duration.toJSON(message.stepDuration) : undefined);
    message.amount !== undefined && (obj.amount = message.amount);
    message.amountMultiplier !== undefined && (obj.amountMultiplier = message.amountMultiplier);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ExponentialStepMinting>, I>>(object: I): ExponentialStepMinting {
    const message = createBaseExponentialStepMinting();
    message.stepDuration = (object.stepDuration !== undefined && object.stepDuration !== null)
      ? Duration.fromPartial(object.stepDuration)
      : undefined;
    message.amount = object.amount ?? "";
    message.amountMultiplier = object.amountMultiplier ?? "";
    return message;
  },
};

function createBaseMinterState(): MinterState {
  return {
    sequenceId: 0,
    amountMinted: "",
    remainderToMint: "",
    lastMintBlockTime: undefined,
    remainderFromPreviousMinter: "",
  };
}

export const MinterState = {
  encode(message: MinterState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.sequenceId !== 0) {
      writer.uint32(8).uint32(message.sequenceId);
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
    if (message.remainderFromPreviousMinter !== "") {
      writer.uint32(42).string(message.remainderFromPreviousMinter);
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
          message.sequenceId = reader.uint32();
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
          message.remainderFromPreviousMinter = reader.string();
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
      sequenceId: isSet(object.sequenceId) ? Number(object.sequenceId) : 0,
      amountMinted: isSet(object.amountMinted) ? String(object.amountMinted) : "",
      remainderToMint: isSet(object.remainderToMint) ? String(object.remainderToMint) : "",
      lastMintBlockTime: isSet(object.lastMintBlockTime) ? fromJsonTimestamp(object.lastMintBlockTime) : undefined,
      remainderFromPreviousMinter: isSet(object.remainderFromPreviousMinter)
        ? String(object.remainderFromPreviousMinter)
        : "",
    };
  },

  toJSON(message: MinterState): unknown {
    const obj: any = {};
    message.sequenceId !== undefined && (obj.sequenceId = Math.round(message.sequenceId));
    message.amountMinted !== undefined && (obj.amountMinted = message.amountMinted);
    message.remainderToMint !== undefined && (obj.remainderToMint = message.remainderToMint);
    message.lastMintBlockTime !== undefined && (obj.lastMintBlockTime = message.lastMintBlockTime.toISOString());
    message.remainderFromPreviousMinter !== undefined
      && (obj.remainderFromPreviousMinter = message.remainderFromPreviousMinter);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MinterState>, I>>(object: I): MinterState {
    const message = createBaseMinterState();
    message.sequenceId = object.sequenceId ?? 0;
    message.amountMinted = object.amountMinted ?? "";
    message.remainderToMint = object.remainderToMint ?? "";
    message.lastMintBlockTime = object.lastMintBlockTime ?? undefined;
    message.remainderFromPreviousMinter = object.remainderFromPreviousMinter ?? "";
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
