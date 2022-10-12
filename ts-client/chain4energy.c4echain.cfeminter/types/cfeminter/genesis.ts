/* eslint-disable */
import { Params } from "../cfeminter/params";
import { MinterState } from "../cfeminter/minter";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeminter";

/** GenesisState defines the cfeminter module's genesis state. */
export interface GenesisState {
  /** this line is used by starport scaffolding # genesis/proto/state */
  params: Params | undefined;
  minterState: MinterState | undefined;
  stateHistory: MinterState[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.minterState !== undefined) {
      MinterState.encode(
        message.minterState,
        writer.uint32(18).fork()
      ).ldelim();
    }
    for (const v of message.stateHistory) {
      MinterState.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.stateHistory = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.minterState = MinterState.decode(reader, reader.uint32());
          break;
        case 3:
          message.stateHistory.push(
            MinterState.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.stateHistory = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.minterState !== undefined && object.minterState !== null) {
      message.minterState = MinterState.fromJSON(object.minterState);
    } else {
      message.minterState = undefined;
    }
    if (object.stateHistory !== undefined && object.stateHistory !== null) {
      for (const e of object.stateHistory) {
        message.stateHistory.push(MinterState.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.minterState !== undefined &&
      (obj.minterState = message.minterState
        ? MinterState.toJSON(message.minterState)
        : undefined);
    if (message.stateHistory) {
      obj.stateHistory = message.stateHistory.map((e) =>
        e ? MinterState.toJSON(e) : undefined
      );
    } else {
      obj.stateHistory = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.stateHistory = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.minterState !== undefined && object.minterState !== null) {
      message.minterState = MinterState.fromPartial(object.minterState);
    } else {
      message.minterState = undefined;
    }
    if (object.stateHistory !== undefined && object.stateHistory !== null) {
      for (const e of object.stateHistory) {
        message.stateHistory.push(MinterState.fromPartial(e));
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
