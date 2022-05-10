/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfesignature";

export interface Signature {
  signature: string;
  algorithm: string;
  certificate: string;
  timestamp: string;
}

const baseSignature: object = {
  signature: "",
  algorithm: "",
  certificate: "",
  timestamp: "",
};

export const Signature = {
  encode(message: Signature, writer: Writer = Writer.create()): Writer {
    if (message.signature !== "") {
      writer.uint32(10).string(message.signature);
    }
    if (message.algorithm !== "") {
      writer.uint32(18).string(message.algorithm);
    }
    if (message.certificate !== "") {
      writer.uint32(26).string(message.certificate);
    }
    if (message.timestamp !== "") {
      writer.uint32(34).string(message.timestamp);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Signature {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSignature } as Signature;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.signature = reader.string();
          break;
        case 2:
          message.algorithm = reader.string();
          break;
        case 3:
          message.certificate = reader.string();
          break;
        case 4:
          message.timestamp = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Signature {
    const message = { ...baseSignature } as Signature;
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = String(object.signature);
    } else {
      message.signature = "";
    }
    if (object.algorithm !== undefined && object.algorithm !== null) {
      message.algorithm = String(object.algorithm);
    } else {
      message.algorithm = "";
    }
    if (object.certificate !== undefined && object.certificate !== null) {
      message.certificate = String(object.certificate);
    } else {
      message.certificate = "";
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = String(object.timestamp);
    } else {
      message.timestamp = "";
    }
    return message;
  },

  toJSON(message: Signature): unknown {
    const obj: any = {};
    message.signature !== undefined && (obj.signature = message.signature);
    message.algorithm !== undefined && (obj.algorithm = message.algorithm);
    message.certificate !== undefined &&
      (obj.certificate = message.certificate);
    message.timestamp !== undefined && (obj.timestamp = message.timestamp);
    return obj;
  },

  fromPartial(object: DeepPartial<Signature>): Signature {
    const message = { ...baseSignature } as Signature;
    if (object.signature !== undefined && object.signature !== null) {
      message.signature = object.signature;
    } else {
      message.signature = "";
    }
    if (object.algorithm !== undefined && object.algorithm !== null) {
      message.algorithm = object.algorithm;
    } else {
      message.algorithm = "";
    }
    if (object.certificate !== undefined && object.certificate !== null) {
      message.certificate = object.certificate;
    } else {
      message.certificate = "";
    }
    if (object.timestamp !== undefined && object.timestamp !== null) {
      message.timestamp = object.timestamp;
    } else {
      message.timestamp = "";
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
