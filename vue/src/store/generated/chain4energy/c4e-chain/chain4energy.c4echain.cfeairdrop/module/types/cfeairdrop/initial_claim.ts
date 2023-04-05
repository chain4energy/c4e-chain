/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeairdrop";

export interface InitialClaim {
  campaignId: string;
  enabled: string;
  missionId: string;
}

const baseInitialClaim: object = { campaignId: "", enabled: "", missionId: "" };

export const InitialClaim = {
  encode(message: InitialClaim, writer: Writer = Writer.create()): Writer {
    if (message.campaignId !== "") {
      writer.uint32(10).string(message.campaignId);
    }
    if (message.enabled !== "") {
      writer.uint32(18).string(message.enabled);
    }
    if (message.missionId !== "") {
      writer.uint32(26).string(message.missionId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): InitialClaim {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseInitialClaim } as InitialClaim;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaignId = reader.string();
          break;
        case 2:
          message.enabled = reader.string();
          break;
        case 3:
          message.missionId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InitialClaim {
    const message = { ...baseInitialClaim } as InitialClaim;
    if (object.campaignId !== undefined && object.campaignId !== null) {
      message.campaignId = String(object.campaignId);
    } else {
      message.campaignId = "";
    }
    if (object.enabled !== undefined && object.enabled !== null) {
      message.enabled = String(object.enabled);
    } else {
      message.enabled = "";
    }
    if (object.missionId !== undefined && object.missionId !== null) {
      message.missionId = String(object.missionId);
    } else {
      message.missionId = "";
    }
    return message;
  },

  toJSON(message: InitialClaim): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    message.enabled !== undefined && (obj.enabled = message.enabled);
    message.missionId !== undefined && (obj.missionId = message.missionId);
    return obj;
  },

  fromPartial(object: DeepPartial<InitialClaim>): InitialClaim {
    const message = { ...baseInitialClaim } as InitialClaim;
    if (object.campaignId !== undefined && object.campaignId !== null) {
      message.campaignId = object.campaignId;
    } else {
      message.campaignId = "";
    }
    if (object.enabled !== undefined && object.enabled !== null) {
      message.enabled = object.enabled;
    } else {
      message.enabled = "";
    }
    if (object.missionId !== undefined && object.missionId !== null) {
      message.missionId = object.missionId;
    } else {
      message.missionId = "";
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
