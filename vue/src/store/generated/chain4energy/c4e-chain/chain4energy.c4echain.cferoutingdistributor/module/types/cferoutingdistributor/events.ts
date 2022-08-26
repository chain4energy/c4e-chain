/* eslint-disable */
import { DecCoin } from "../cosmos/base/v1beta1/coin";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfedistributor";

export interface EventDistributionFinished {
  subDistributionFinished: SubDistributionFinished[];
}

export interface SubDistributionFinished {
  sources: string[];
  processedDestination: ProcessedDestination[];
}

export interface ProcessedDestination {
  name: string;
  coin_sended: DecCoin[];
}

const baseEventDistributionFinished: object = {};

export const EventDistributionFinished = {
  encode(
    message: EventDistributionFinished,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.subDistributionFinished) {
      SubDistributionFinished.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): EventDistributionFinished {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseEventDistributionFinished,
    } as EventDistributionFinished;
    message.subDistributionFinished = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.subDistributionFinished.push(
            SubDistributionFinished.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventDistributionFinished {
    const message = {
      ...baseEventDistributionFinished,
    } as EventDistributionFinished;
    message.subDistributionFinished = [];
    if (
      object.subDistributionFinished !== undefined &&
      object.subDistributionFinished !== null
    ) {
      for (const e of object.subDistributionFinished) {
        message.subDistributionFinished.push(
          SubDistributionFinished.fromJSON(e)
        );
      }
    }
    return message;
  },

  toJSON(message: EventDistributionFinished): unknown {
    const obj: any = {};
    if (message.subDistributionFinished) {
      obj.subDistributionFinished = message.subDistributionFinished.map((e) =>
        e ? SubDistributionFinished.toJSON(e) : undefined
      );
    } else {
      obj.subDistributionFinished = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<EventDistributionFinished>
  ): EventDistributionFinished {
    const message = {
      ...baseEventDistributionFinished,
    } as EventDistributionFinished;
    message.subDistributionFinished = [];
    if (
      object.subDistributionFinished !== undefined &&
      object.subDistributionFinished !== null
    ) {
      for (const e of object.subDistributionFinished) {
        message.subDistributionFinished.push(
          SubDistributionFinished.fromPartial(e)
        );
      }
    }
    return message;
  },
};

const baseSubDistributionFinished: object = { sources: "" };

export const SubDistributionFinished = {
  encode(
    message: SubDistributionFinished,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.sources) {
      writer.uint32(10).string(v!);
    }
    for (const v of message.processedDestination) {
      ProcessedDestination.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SubDistributionFinished {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseSubDistributionFinished,
    } as SubDistributionFinished;
    message.sources = [];
    message.processedDestination = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.sources.push(reader.string());
          break;
        case 2:
          message.processedDestination.push(
            ProcessedDestination.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SubDistributionFinished {
    const message = {
      ...baseSubDistributionFinished,
    } as SubDistributionFinished;
    message.sources = [];
    message.processedDestination = [];
    if (object.sources !== undefined && object.sources !== null) {
      for (const e of object.sources) {
        message.sources.push(String(e));
      }
    }
    if (
      object.processedDestination !== undefined &&
      object.processedDestination !== null
    ) {
      for (const e of object.processedDestination) {
        message.processedDestination.push(ProcessedDestination.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: SubDistributionFinished): unknown {
    const obj: any = {};
    if (message.sources) {
      obj.sources = message.sources.map((e) => e);
    } else {
      obj.sources = [];
    }
    if (message.processedDestination) {
      obj.processedDestination = message.processedDestination.map((e) =>
        e ? ProcessedDestination.toJSON(e) : undefined
      );
    } else {
      obj.processedDestination = [];
    }
    return obj;
  },

  fromPartial(
    object: DeepPartial<SubDistributionFinished>
  ): SubDistributionFinished {
    const message = {
      ...baseSubDistributionFinished,
    } as SubDistributionFinished;
    message.sources = [];
    message.processedDestination = [];
    if (object.sources !== undefined && object.sources !== null) {
      for (const e of object.sources) {
        message.sources.push(e);
      }
    }
    if (
      object.processedDestination !== undefined &&
      object.processedDestination !== null
    ) {
      for (const e of object.processedDestination) {
        message.processedDestination.push(ProcessedDestination.fromPartial(e));
      }
    }
    return message;
  },
};

const baseProcessedDestination: object = { name: "" };

export const ProcessedDestination = {
  encode(
    message: ProcessedDestination,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    for (const v of message.coin_sended) {
      DecCoin.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ProcessedDestination {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseProcessedDestination } as ProcessedDestination;
    message.coin_sended = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.coin_sended.push(DecCoin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ProcessedDestination {
    const message = { ...baseProcessedDestination } as ProcessedDestination;
    message.coin_sended = [];
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.coin_sended !== undefined && object.coin_sended !== null) {
      for (const e of object.coin_sended) {
        message.coin_sended.push(DecCoin.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: ProcessedDestination): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    if (message.coin_sended) {
      obj.coin_sended = message.coin_sended.map((e) =>
        e ? DecCoin.toJSON(e) : undefined
      );
    } else {
      obj.coin_sended = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<ProcessedDestination>): ProcessedDestination {
    const message = { ...baseProcessedDestination } as ProcessedDestination;
    message.coin_sended = [];
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.coin_sended !== undefined && object.coin_sended !== null) {
      for (const e of object.coin_sended) {
        message.coin_sended.push(DecCoin.fromPartial(e));
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
