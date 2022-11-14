/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeairdrop";

export interface Mission {
  campaignId: string;
  missionId: string;
  description: string;
  weight: string;
}

const baseMission: object = {
  campaignId: "",
  missionId: "",
  description: "",
  weight: "",
};

export const Mission = {
  encode(message: Mission, writer: Writer = Writer.create()): Writer {
    if (message.campaignId !== "") {
      writer.uint32(10).string(message.campaignId);
    }
    if (message.missionId !== "") {
      writer.uint32(18).string(message.missionId);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.weight !== "") {
      writer.uint32(34).string(message.weight);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Mission {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMission } as Mission;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaignId = reader.string();
          break;
        case 2:
          message.missionId = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.weight = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Mission {
    const message = { ...baseMission } as Mission;
    if (object.campaignId !== undefined && object.campaignId !== null) {
      message.campaignId = String(object.campaignId);
    } else {
      message.campaignId = "";
    }
    if (object.missionId !== undefined && object.missionId !== null) {
      message.missionId = String(object.missionId);
    } else {
      message.missionId = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.weight !== undefined && object.weight !== null) {
      message.weight = String(object.weight);
    } else {
      message.weight = "";
    }
    return message;
  },

  toJSON(message: Mission): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    message.missionId !== undefined && (obj.missionId = message.missionId);
    message.description !== undefined &&
      (obj.description = message.description);
    message.weight !== undefined && (obj.weight = message.weight);
    return obj;
  },

  fromPartial(object: DeepPartial<Mission>): Mission {
    const message = { ...baseMission } as Mission;
    if (object.campaignId !== undefined && object.campaignId !== null) {
      message.campaignId = object.campaignId;
    } else {
      message.campaignId = "";
    }
    if (object.missionId !== undefined && object.missionId !== null) {
      message.missionId = object.missionId;
    } else {
      message.missionId = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.weight !== undefined && object.weight !== null) {
      message.weight = object.weight;
    } else {
      message.weight = "";
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
