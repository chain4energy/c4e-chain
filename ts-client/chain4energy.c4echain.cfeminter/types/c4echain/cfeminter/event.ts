/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeminter";

export interface EventMint {
  bondedRatio: string;
  inflation: string;
  amount: string;
}

function createBaseEventMint(): EventMint {
  return { bondedRatio: "", inflation: "", amount: "" };
}

export const EventMint = {
  encode(message: EventMint, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): EventMint {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventMint();
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

  fromJSON(object: any): EventMint {
    return {
      bondedRatio: isSet(object.bondedRatio) ? String(object.bondedRatio) : "",
      inflation: isSet(object.inflation) ? String(object.inflation) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
    };
  },

  toJSON(message: EventMint): unknown {
    const obj: any = {};
    message.bondedRatio !== undefined && (obj.bondedRatio = message.bondedRatio);
    message.inflation !== undefined && (obj.inflation = message.inflation);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventMint>, I>>(object: I): EventMint {
    const message = createBaseEventMint();
    message.bondedRatio = object.bondedRatio ?? "";
    message.inflation = object.inflation ?? "";
    message.amount = object.amount ?? "";
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
