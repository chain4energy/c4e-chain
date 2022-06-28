/* eslint-disable */
import { Minter } from "../cfeminter/minter";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeminter";

/** Params defines the parameters for the module. */
export interface Params {
  mint_denom: string;
  minter: Minter | undefined;
}

const baseParams: object = { mint_denom: "" };

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.mint_denom !== "") {
      writer.uint32(10).string(message.mint_denom);
    }
    if (message.minter !== undefined) {
      Minter.encode(message.minter, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.mint_denom = reader.string();
          break;
        case 2:
          message.minter = Minter.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    if (object.mint_denom !== undefined && object.mint_denom !== null) {
      message.mint_denom = String(object.mint_denom);
    } else {
      message.mint_denom = "";
    }
    if (object.minter !== undefined && object.minter !== null) {
      message.minter = Minter.fromJSON(object.minter);
    } else {
      message.minter = undefined;
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.mint_denom !== undefined && (obj.mint_denom = message.mint_denom);
    message.minter !== undefined &&
      (obj.minter = message.minter ? Minter.toJSON(message.minter) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (object.mint_denom !== undefined && object.mint_denom !== null) {
      message.mint_denom = object.mint_denom;
    } else {
      message.mint_denom = "";
    }
    if (object.minter !== undefined && object.minter !== null) {
      message.minter = Minter.fromPartial(object.minter);
    } else {
      message.minter = undefined;
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
