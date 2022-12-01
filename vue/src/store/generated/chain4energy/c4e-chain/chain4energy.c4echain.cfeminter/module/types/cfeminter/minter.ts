/* eslint-disable */
import { Timestamp } from "../google/protobuf/timestamp";
import { Duration } from "../google/protobuf/duration";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeminter";

export interface Minter {
  /** option (gogoproto.goproto_getters) = false; */
  sequence_id: number;
  end_time: Date | undefined;
  /**
   * types:
   *   NO_MINTING;
   *   TIME_LINEAR_MINTER;
   *   PERIODIC_REDUCTION_MINTER;
   */
  type: string;
  linear_minting: LinearMinting | undefined;
  exponential_step_minting: ExponentialStepMinting | undefined;
}

export interface LinearMinting {
  amount: string;
}

export interface ExponentialStepMinting {
  /** mint_period in seconds */
  amount: string;
  step_duration: Duration | undefined;
  amount_multiplier: string;
}

export interface MinterState {
  SequenceId: number;
  amount_minted: string;
  remainder_to_mint: string;
  last_mint_block_time: Date | undefined;
  remainder_from_previous_period: string;
}

const baseMinter: object = { sequence_id: 0, type: "" };

export const Minter = {
  encode(message: Minter, writer: Writer = Writer.create()): Writer {
    if (message.sequence_id !== 0) {
      writer.uint32(8).int32(message.sequence_id);
    }
    if (message.end_time !== undefined) {
      Timestamp.encode(
        toTimestamp(message.end_time),
        writer.uint32(18).fork()
      ).ldelim();
    }
    if (message.type !== "") {
      writer.uint32(26).string(message.type);
    }
    if (message.linear_minting !== undefined) {
      LinearMinting.encode(
        message.linear_minting,
        writer.uint32(34).fork()
      ).ldelim();
    }
    if (message.exponential_step_minting !== undefined) {
      ExponentialStepMinting.encode(
        message.exponential_step_minting,
        writer.uint32(42).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Minter {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMinter } as Minter;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sequence_id = reader.int32();
          break;
        case 2:
          message.end_time = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.type = reader.string();
          break;
        case 4:
          message.linear_minting = LinearMinting.decode(
            reader,
            reader.uint32()
          );
          break;
        case 5:
          message.exponential_step_minting = ExponentialStepMinting.decode(
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

  fromJSON(object: any): Minter {
    const message = { ...baseMinter } as Minter;
    if (object.sequence_id !== undefined && object.sequence_id !== null) {
      message.sequence_id = Number(object.sequence_id);
    } else {
      message.sequence_id = 0;
    }
    if (object.end_time !== undefined && object.end_time !== null) {
      message.end_time = fromJsonTimestamp(object.end_time);
    } else {
      message.end_time = undefined;
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = String(object.type);
    } else {
      message.type = "";
    }
    if (object.linear_minting !== undefined && object.linear_minting !== null) {
      message.linear_minting = LinearMinting.fromJSON(object.linear_minting);
    } else {
      message.linear_minting = undefined;
    }
    if (
      object.exponential_step_minting !== undefined &&
      object.exponential_step_minting !== null
    ) {
      message.exponential_step_minting = ExponentialStepMinting.fromJSON(
        object.exponential_step_minting
      );
    } else {
      message.exponential_step_minting = undefined;
    }
    return message;
  },

  toJSON(message: Minter): unknown {
    const obj: any = {};
    message.sequence_id !== undefined &&
      (obj.sequence_id = message.sequence_id);
    message.end_time !== undefined &&
      (obj.end_time =
        message.end_time !== undefined ? message.end_time.toISOString() : null);
    message.type !== undefined && (obj.type = message.type);
    message.linear_minting !== undefined &&
      (obj.linear_minting = message.linear_minting
        ? LinearMinting.toJSON(message.linear_minting)
        : undefined);
    message.exponential_step_minting !== undefined &&
      (obj.exponential_step_minting = message.exponential_step_minting
        ? ExponentialStepMinting.toJSON(message.exponential_step_minting)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Minter>): Minter {
    const message = { ...baseMinter } as Minter;
    if (object.sequence_id !== undefined && object.sequence_id !== null) {
      message.sequence_id = object.sequence_id;
    } else {
      message.sequence_id = 0;
    }
    if (object.end_time !== undefined && object.end_time !== null) {
      message.end_time = object.end_time;
    } else {
      message.end_time = undefined;
    }
    if (object.type !== undefined && object.type !== null) {
      message.type = object.type;
    } else {
      message.type = "";
    }
    if (object.linear_minting !== undefined && object.linear_minting !== null) {
      message.linear_minting = LinearMinting.fromPartial(object.linear_minting);
    } else {
      message.linear_minting = undefined;
    }
    if (
      object.exponential_step_minting !== undefined &&
      object.exponential_step_minting !== null
    ) {
      message.exponential_step_minting = ExponentialStepMinting.fromPartial(
        object.exponential_step_minting
      );
    } else {
      message.exponential_step_minting = undefined;
    }
    return message;
  },
};

const baseLinearMinting: object = { amount: "" };

export const LinearMinting = {
  encode(message: LinearMinting, writer: Writer = Writer.create()): Writer {
    if (message.amount !== "") {
      writer.uint32(10).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): LinearMinting {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLinearMinting } as LinearMinting;
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
    const message = { ...baseLinearMinting } as LinearMinting;
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: LinearMinting): unknown {
    const obj: any = {};
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<LinearMinting>): LinearMinting {
    const message = { ...baseLinearMinting } as LinearMinting;
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    return message;
  },
};

const baseExponentialStepMinting: object = {
  amount: "",
  amount_multiplier: "",
};

export const ExponentialStepMinting = {
  encode(
    message: ExponentialStepMinting,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    if (message.step_duration !== undefined) {
      Duration.encode(message.step_duration, writer.uint32(10).fork()).ldelim();
    }
    if (message.amount_multiplier !== "") {
      writer.uint32(34).string(message.amount_multiplier);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ExponentialStepMinting {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseExponentialStepMinting } as ExponentialStepMinting;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.amount = reader.string();
          break;
        case 1:
          message.step_duration = Duration.decode(reader, reader.uint32());
          break;
        case 4:
          message.amount_multiplier = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ExponentialStepMinting {
    const message = { ...baseExponentialStepMinting } as ExponentialStepMinting;
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.step_duration !== undefined && object.step_duration !== null) {
      message.step_duration = Duration.fromJSON(object.step_duration);
    } else {
      message.step_duration = undefined;
    }
    if (
      object.amount_multiplier !== undefined &&
      object.amount_multiplier !== null
    ) {
      message.amount_multiplier = String(object.amount_multiplier);
    } else {
      message.amount_multiplier = "";
    }
    return message;
  },

  toJSON(message: ExponentialStepMinting): unknown {
    const obj: any = {};
    message.amount !== undefined && (obj.amount = message.amount);
    message.step_duration !== undefined &&
      (obj.step_duration = message.step_duration
        ? Duration.toJSON(message.step_duration)
        : undefined);
    message.amount_multiplier !== undefined &&
      (obj.amount_multiplier = message.amount_multiplier);
    return obj;
  },

  fromPartial(
    object: DeepPartial<ExponentialStepMinting>
  ): ExponentialStepMinting {
    const message = { ...baseExponentialStepMinting } as ExponentialStepMinting;
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.step_duration !== undefined && object.step_duration !== null) {
      message.step_duration = Duration.fromPartial(object.step_duration);
    } else {
      message.step_duration = undefined;
    }
    if (
      object.amount_multiplier !== undefined &&
      object.amount_multiplier !== null
    ) {
      message.amount_multiplier = object.amount_multiplier;
    } else {
      message.amount_multiplier = "";
    }
    return message;
  },
};

const baseMinterState: object = {
  SequenceId: 0,
  amount_minted: "",
  remainder_to_mint: "",
  remainder_from_previous_period: "",
};

export const MinterState = {
  encode(message: MinterState, writer: Writer = Writer.create()): Writer {
    if (message.SequenceId !== 0) {
      writer.uint32(8).int32(message.SequenceId);
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
          message.SequenceId = reader.int32();
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
    if (object.SequenceId !== undefined && object.SequenceId !== null) {
      message.SequenceId = Number(object.SequenceId);
    } else {
      message.SequenceId = 0;
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
    message.SequenceId !== undefined && (obj.SequenceId = message.SequenceId);
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
    if (object.SequenceId !== undefined && object.SequenceId !== null) {
      message.SequenceId = object.SequenceId;
    } else {
      message.SequenceId = 0;
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
