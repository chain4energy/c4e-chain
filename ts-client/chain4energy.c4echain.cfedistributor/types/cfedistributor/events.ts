/* eslint-disable */
import { Account } from "../cfedistributor/sub_distributor";
import { DecCoin } from "../cosmos/base/v1beta1/coin";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfedistributor";

export interface DistributionResult {
  source: Account[];
  destination: Account | undefined;
  coinSend: DecCoin[];
}

export interface DistributionsResult {
  distributionResult: DistributionResult[];
}

const baseDistributionResult: object = {};

export const DistributionResult = {
  encode(
    message: DistributionResult,
    writer: Writer = Writer.create()
  ): Writer {
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

  decode(input: Reader | Uint8Array, length?: number): DistributionResult {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDistributionResult } as DistributionResult;
    message.source = [];
    message.coinSend = [];
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
    const message = { ...baseDistributionResult } as DistributionResult;
    message.source = [];
    message.coinSend = [];
    if (object.source !== undefined && object.source !== null) {
      for (const e of object.source) {
        message.source.push(Account.fromJSON(e));
      }
    }
    if (object.destination !== undefined && object.destination !== null) {
      message.destination = Account.fromJSON(object.destination);
    } else {
      message.destination = undefined;
    }
    if (object.coinSend !== undefined && object.coinSend !== null) {
      for (const e of object.coinSend) {
        message.coinSend.push(DecCoin.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: DistributionResult): unknown {
    const obj: any = {};
    if (message.source) {
      obj.source = message.source.map((e) =>
        e ? Account.toJSON(e) : undefined
      );
    } else {
      obj.source = [];
    }
    message.destination !== undefined &&
      (obj.destination = message.destination
        ? Account.toJSON(message.destination)
        : undefined);
    if (message.coinSend) {
      obj.coinSend = message.coinSend.map((e) =>
        e ? DecCoin.toJSON(e) : undefined
      );
    } else {
      obj.coinSend = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<DistributionResult>): DistributionResult {
    const message = { ...baseDistributionResult } as DistributionResult;
    message.source = [];
    message.coinSend = [];
    if (object.source !== undefined && object.source !== null) {
      for (const e of object.source) {
        message.source.push(Account.fromPartial(e));
      }
    }
    if (object.destination !== undefined && object.destination !== null) {
      message.destination = Account.fromPartial(object.destination);
    } else {
      message.destination = undefined;
    }
    if (object.coinSend !== undefined && object.coinSend !== null) {
      for (const e of object.coinSend) {
        message.coinSend.push(DecCoin.fromPartial(e));
      }
    }
    return message;
  },
};

const baseDistributionsResult: object = {};

export const DistributionsResult = {
  encode(
    message: DistributionsResult,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.distributionResult) {
      DistributionResult.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): DistributionsResult {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDistributionsResult } as DistributionsResult;
    message.distributionResult = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.distributionResult.push(
            DistributionResult.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DistributionsResult {
    const message = { ...baseDistributionsResult } as DistributionsResult;
    message.distributionResult = [];
    if (
      object.distributionResult !== undefined &&
      object.distributionResult !== null
    ) {
      for (const e of object.distributionResult) {
        message.distributionResult.push(DistributionResult.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: DistributionsResult): unknown {
    const obj: any = {};
    if (message.distributionResult) {
      obj.distributionResult = message.distributionResult.map((e) =>
        e ? DistributionResult.toJSON(e) : undefined
      );
    } else {
      obj.distributionResult = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<DistributionsResult>): DistributionsResult {
    const message = { ...baseDistributionsResult } as DistributionsResult;
    message.distributionResult = [];
    if (
      object.distributionResult !== undefined &&
      object.distributionResult !== null
    ) {
      for (const e of object.distributionResult) {
        message.distributionResult.push(DistributionResult.fromPartial(e));
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
