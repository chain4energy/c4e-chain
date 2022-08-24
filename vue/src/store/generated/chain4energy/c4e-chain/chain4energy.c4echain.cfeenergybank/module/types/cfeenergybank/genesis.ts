/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../cfeenergybank/params";
import { EnergyToken } from "../cfeenergybank/energy_token";
import { TokenParams } from "../cfeenergybank/token_params";

export const protobufPackage = "chain4energy.c4echain.cfeenergybank";

/** GenesisState defines the energybank module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  energyTokenList: EnergyToken[];
  energyTokenCount: number;
  /** this line is used by starport scaffolding # genesis/proto/state */
  tokenParamsList: TokenParams[];
}

const baseGenesisState: object = { energyTokenCount: 0 };

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.energyTokenList) {
      EnergyToken.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.energyTokenCount !== 0) {
      writer.uint32(24).uint64(message.energyTokenCount);
    }
    for (const v of message.tokenParamsList) {
      TokenParams.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.energyTokenList = [];
    message.tokenParamsList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.energyTokenList.push(
            EnergyToken.decode(reader, reader.uint32())
          );
          break;
        case 3:
          message.energyTokenCount = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.tokenParamsList.push(
            TokenParams.decode(reader, reader.uint32())
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
    message.energyTokenList = [];
    message.tokenParamsList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (
      object.energyTokenList !== undefined &&
      object.energyTokenList !== null
    ) {
      for (const e of object.energyTokenList) {
        message.energyTokenList.push(EnergyToken.fromJSON(e));
      }
    }
    if (
      object.energyTokenCount !== undefined &&
      object.energyTokenCount !== null
    ) {
      message.energyTokenCount = Number(object.energyTokenCount);
    } else {
      message.energyTokenCount = 0;
    }
    if (
      object.tokenParamsList !== undefined &&
      object.tokenParamsList !== null
    ) {
      for (const e of object.tokenParamsList) {
        message.tokenParamsList.push(TokenParams.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.energyTokenList) {
      obj.energyTokenList = message.energyTokenList.map((e) =>
        e ? EnergyToken.toJSON(e) : undefined
      );
    } else {
      obj.energyTokenList = [];
    }
    message.energyTokenCount !== undefined &&
      (obj.energyTokenCount = message.energyTokenCount);
    if (message.tokenParamsList) {
      obj.tokenParamsList = message.tokenParamsList.map((e) =>
        e ? TokenParams.toJSON(e) : undefined
      );
    } else {
      obj.tokenParamsList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.energyTokenList = [];
    message.tokenParamsList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (
      object.energyTokenList !== undefined &&
      object.energyTokenList !== null
    ) {
      for (const e of object.energyTokenList) {
        message.energyTokenList.push(EnergyToken.fromPartial(e));
      }
    }
    if (
      object.energyTokenCount !== undefined &&
      object.energyTokenCount !== null
    ) {
      message.energyTokenCount = object.energyTokenCount;
    } else {
      message.energyTokenCount = 0;
    }
    if (
      object.tokenParamsList !== undefined &&
      object.tokenParamsList !== null
    ) {
      for (const e of object.tokenParamsList) {
        message.tokenParamsList.push(TokenParams.fromPartial(e));
      }
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
