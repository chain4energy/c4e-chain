/* eslint-disable */
import { SubDistributor } from "../cferoutingdistributor/sub_distributor";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cferoutingdistributor";

/** Params defines the parameters for the module. */
export interface Params {
  sub_distributors: SubDistributor[];
}

const baseParams: object = {};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    for (const v of message.sub_distributors) {
      SubDistributor.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    message.sub_distributors = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sub_distributors.push(
            SubDistributor.decode(reader, reader.uint32())
          );
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
    message.sub_distributors = [];
    if (
      object.sub_distributors !== undefined &&
      object.sub_distributors !== null
    ) {
      for (const e of object.sub_distributors) {
        message.sub_distributors.push(SubDistributor.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    if (message.sub_distributors) {
      obj.sub_distributors = message.sub_distributors.map((e) =>
        e ? SubDistributor.toJSON(e) : undefined
      );
    } else {
      obj.sub_distributors = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    message.sub_distributors = [];
    if (
      object.sub_distributors !== undefined &&
      object.sub_distributors !== null
    ) {
      for (const e of object.sub_distributors) {
        message.sub_distributors.push(SubDistributor.fromPartial(e));
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
