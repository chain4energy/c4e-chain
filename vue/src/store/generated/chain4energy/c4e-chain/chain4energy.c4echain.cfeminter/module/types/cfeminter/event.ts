/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeminter";

export interface Mint {
  bonded_ratio: string;
  inflation: string;
  amount: string;
}

const baseMint: object = { bonded_ratio: "", inflation: "", amount: "" };

export const Mint = {
  encode(message: Mint, writer: Writer = Writer.create()): Writer {
    if (message.bonded_ratio !== "") {
      writer.uint32(10).string(message.bonded_ratio);
    }
    if (message.inflation !== "") {
      writer.uint32(18).string(message.inflation);
    }
    if (message.amount !== "") {
      writer.uint32(26).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Mint {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMint } as Mint;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.bonded_ratio = reader.string();
          break;
        case 2:
          message.inflation = reader.string();
          break;
        case 3:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Mint {
    const message = { ...baseMint } as Mint;
    if (object.bonded_ratio !== undefined && object.bonded_ratio !== null) {
      message.bonded_ratio = String(object.bonded_ratio);
    } else {
      message.bonded_ratio = "";
    }
    if (object.inflation !== undefined && object.inflation !== null) {
      message.inflation = String(object.inflation);
    } else {
      message.inflation = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: Mint): unknown {
    const obj: any = {};
    message.bonded_ratio !== undefined &&
      (obj.bonded_ratio = message.bonded_ratio);
    message.inflation !== undefined && (obj.inflation = message.inflation);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<Mint>): Mint {
    const message = { ...baseMint } as Mint;
    if (object.bonded_ratio !== undefined && object.bonded_ratio !== null) {
      message.bonded_ratio = object.bonded_ratio;
    } else {
      message.bonded_ratio = "";
    }
    if (object.inflation !== undefined && object.inflation !== null) {
      message.inflation = object.inflation;
    } else {
      message.inflation = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
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
