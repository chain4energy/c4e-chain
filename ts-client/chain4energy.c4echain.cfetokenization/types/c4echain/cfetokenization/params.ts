/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Duration } from "../../google/protobuf/duration";

export const protobufPackage = "chain4energy.c4echain.cfetokenization";

/** Params defines the parameters for the module. */
export interface Params {
  actionTimeWindow: Duration | undefined;
}

function createBaseParams(): Params {
  return { actionTimeWindow: undefined };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.actionTimeWindow !== undefined) {
      Duration.encode(message.actionTimeWindow, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.actionTimeWindow = Duration.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    return {
      actionTimeWindow: isSet(object.actionTimeWindow) ? Duration.fromJSON(object.actionTimeWindow) : undefined,
    };
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.actionTimeWindow !== undefined
      && (obj.actionTimeWindow = message.actionTimeWindow ? Duration.toJSON(message.actionTimeWindow) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params>, I>>(object: I): Params {
    const message = createBaseParams();
    message.actionTimeWindow = (object.actionTimeWindow !== undefined && object.actionTimeWindow !== null)
      ? Duration.fromPartial(object.actionTimeWindow)
      : undefined;
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
