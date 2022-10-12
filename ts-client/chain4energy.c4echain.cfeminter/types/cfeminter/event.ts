/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeminter";

export interface Mint {
  bondedRatio: string;
  inflation: string;
  amount: string;
}

const baseMint: object = { bondedRatio: "", inflation: "", amount: "" };

export const Mint = {
  encode(message: Mint, writer: Writer = Writer.create()): Writer {
    if (message.bondedRatio !== "") {
      writer.uint32(10).string(message.bondedRatio);
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
          message.bondedRatio = reader.string();
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
    if (object.bondedRatio !== undefined && object.bondedRatio !== null) {
      message.bondedRatio = String(object.bondedRatio);
    } else {
      message.bondedRatio = "";
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
    message.bondedRatio !== undefined &&
      (obj.bondedRatio = message.bondedRatio);
    message.inflation !== undefined && (obj.inflation = message.inflation);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<Mint>): Mint {
    const message = { ...baseMint } as Mint;
    if (object.bondedRatio !== undefined && object.bondedRatio !== null) {
      message.bondedRatio = object.bondedRatio;
    } else {
      message.bondedRatio = "";
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
