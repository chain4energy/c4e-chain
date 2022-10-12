/* eslint-disable */
import { SubDistributor } from "../cfedistributor/sub_distributor";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfedistributor";

/** Params defines the parameters for the module. */
export interface Params {
  subDistributors: SubDistributor[];
}

const baseParams: object = {};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    for (const v of message.subDistributors) {
      SubDistributor.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    message.subDistributors = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.subDistributors.push(
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
    message.subDistributors = [];
    if (
      object.subDistributors !== undefined &&
      object.subDistributors !== null
    ) {
      for (const e of object.subDistributors) {
        message.subDistributors.push(SubDistributor.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    if (message.subDistributors) {
      obj.subDistributors = message.subDistributors.map((e) =>
        e ? SubDistributor.toJSON(e) : undefined
      );
    } else {
      obj.subDistributors = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    message.subDistributors = [];
    if (
      object.subDistributors !== undefined &&
      object.subDistributors !== null
    ) {
      for (const e of object.subDistributors) {
        message.subDistributors.push(SubDistributor.fromPartial(e));
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
