/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { DecCoin } from "../../cosmos/base/v1beta1/coin";
import { Account } from "./sub_distributor";

export const protobufPackage = "chain4energy.c4echain.cfedistributor";

export interface DistributionResult {
  source: Account[];
  destination: Account | undefined;
  coinSend: DecCoin[];
}

export interface DistributionsResult {
  distributionResult: DistributionResult[];
}

function createBaseDistributionResult(): DistributionResult {
  return { source: [], destination: undefined, coinSend: [] };
}

export const DistributionResult = {
  encode(message: DistributionResult, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.source) {
      Account.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.destination !== undefined) {
      Account.encode(message.destination, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.coinSend) {
      DecCoin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DistributionResult {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDistributionResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.source.push(Account.decode(reader, reader.uint32()));
          break;
        case 2:
          message.destination = Account.decode(reader, reader.uint32());
          break;
        case 3:
          message.coinSend.push(DecCoin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DistributionResult {
    return {
      source: Array.isArray(object?.source) ? object.source.map((e: any) => Account.fromJSON(e)) : [],
      destination: isSet(object.destination) ? Account.fromJSON(object.destination) : undefined,
      coinSend: Array.isArray(object?.coinSend) ? object.coinSend.map((e: any) => DecCoin.fromJSON(e)) : [],
    };
  },

  toJSON(message: DistributionResult): unknown {
    const obj: any = {};
    if (message.source) {
      obj.source = message.source.map((e) => e ? Account.toJSON(e) : undefined);
    } else {
      obj.source = [];
    }
    message.destination !== undefined
      && (obj.destination = message.destination ? Account.toJSON(message.destination) : undefined);
    if (message.coinSend) {
      obj.coinSend = message.coinSend.map((e) => e ? DecCoin.toJSON(e) : undefined);
    } else {
      obj.coinSend = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DistributionResult>, I>>(object: I): DistributionResult {
    const message = createBaseDistributionResult();
    message.source = object.source?.map((e) => Account.fromPartial(e)) || [];
    message.destination = (object.destination !== undefined && object.destination !== null)
      ? Account.fromPartial(object.destination)
      : undefined;
    message.coinSend = object.coinSend?.map((e) => DecCoin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseDistributionsResult(): DistributionsResult {
  return { distributionResult: [] };
}

export const DistributionsResult = {
  encode(message: DistributionsResult, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.distributionResult) {
      DistributionResult.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DistributionsResult {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDistributionsResult();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.distributionResult.push(DistributionResult.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DistributionsResult {
    return {
      distributionResult: Array.isArray(object?.distributionResult)
        ? object.distributionResult.map((e: any) => DistributionResult.fromJSON(e))
        : [],
    };
  },

  toJSON(message: DistributionsResult): unknown {
    const obj: any = {};
    if (message.distributionResult) {
      obj.distributionResult = message.distributionResult.map((e) => e ? DistributionResult.toJSON(e) : undefined);
    } else {
      obj.distributionResult = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DistributionsResult>, I>>(object: I): DistributionsResult {
    const message = createBaseDistributionsResult();
    message.distributionResult = object.distributionResult?.map((e) => DistributionResult.fromPartial(e)) || [];
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
