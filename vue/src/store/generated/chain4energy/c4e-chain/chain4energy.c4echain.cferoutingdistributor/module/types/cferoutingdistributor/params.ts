/* eslint-disable */
import { RoutingDistributor } from "../cferoutingdistributor/sub_distributor";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cferoutingdistributor";

/** Params defines the parameters for the module. */
export interface Params {
  routing_distributor: RoutingDistributor | undefined;
}

const baseParams: object = {};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.routing_distributor !== undefined) {
      RoutingDistributor.encode(
        message.routing_distributor,
        writer.uint32(18).fork()
      ).ldelim();
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
        case 2:
          message.routing_distributor = RoutingDistributor.decode(
            reader,
            reader.uint32()
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
    if (
      object.routing_distributor !== undefined &&
      object.routing_distributor !== null
    ) {
      message.routing_distributor = RoutingDistributor.fromJSON(
        object.routing_distributor
      );
    } else {
      message.routing_distributor = undefined;
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.routing_distributor !== undefined &&
      (obj.routing_distributor = message.routing_distributor
        ? RoutingDistributor.toJSON(message.routing_distributor)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (
      object.routing_distributor !== undefined &&
      object.routing_distributor !== null
    ) {
      message.routing_distributor = RoutingDistributor.fromPartial(
        object.routing_distributor
      );
    } else {
      message.routing_distributor = undefined;
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
