/* eslint-disable */
import { Params } from "../cfeminter/params";
import { MinterState } from "../cfeminter/minter";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeminter";

/** GenesisState defines the cfeminter module's genesis state. */
export interface GenesisState {
  /** this line is used by starport scaffolding # genesis/proto/state */
  params: Params | undefined;
  minter_state: MinterState | undefined;
  state_history: MinterState[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.minter_state !== undefined) {
      MinterState.encode(
        message.minter_state,
        writer.uint32(18).fork()
      ).ldelim();
    }
    for (const v of message.state_history) {
      MinterState.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.state_history = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.minter_state = MinterState.decode(reader, reader.uint32());
          break;
        case 3:
          message.state_history.push(
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
    message.state_history = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.minter_state !== undefined && object.minter_state !== null) {
      message.minter_state = MinterState.fromJSON(object.minter_state);
    } else {
      message.minter_state = undefined;
    }
    if (object.state_history !== undefined && object.state_history !== null) {
      for (const e of object.state_history) {
        message.state_history.push(MinterState.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.minter_state !== undefined &&
      (obj.minter_state = message.minter_state
        ? MinterState.toJSON(message.minter_state)
        : undefined);
    if (message.state_history) {
      obj.state_history = message.state_history.map((e) =>
        e ? MinterState.toJSON(e) : undefined
      );
    } else {
      obj.state_history = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.state_history = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.minter_state !== undefined && object.minter_state !== null) {
      message.minter_state = MinterState.fromPartial(object.minter_state);
    } else {
      message.minter_state = undefined;
    }
    if (object.state_history !== undefined && object.state_history !== null) {
      for (const e of object.state_history) {
        message.state_history.push(MinterState.fromPartial(e));
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
