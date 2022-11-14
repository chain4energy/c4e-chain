/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeairdrop";

export interface CampaignRecord {
  campaign_id: number;
  claimable: string;
  completedMissions: number[];
}

export interface ClaimRecord {
  address: string;
  campaign_records: CampaignRecord[];
}

export interface Campaign {
  campaign_id: number;
  description: string;
}

export interface InitialClaimff {
  enabled: boolean;
  campaign_id: number;
  mission_id: number;
}

export interface Mission {
  campaign_id: number;
  mission_id: number;
  description: string;
  weight: string;
}

const baseCampaignRecord: object = {
  campaign_id: 0,
  claimable: "",
  completedMissions: 0,
};

export const CampaignRecord = {
  encode(message: CampaignRecord, writer: Writer = Writer.create()): Writer {
    if (message.campaign_id !== 0) {
      writer.uint32(8).uint64(message.campaign_id);
    }
    if (message.claimable !== "") {
      writer.uint32(18).string(message.claimable);
    }
    writer.uint32(26).fork();
    for (const v of message.completedMissions) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): CampaignRecord {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCampaignRecord } as CampaignRecord;
    message.completedMissions = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaign_id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.claimable = reader.string();
          break;
        case 3:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.completedMissions.push(
                longToNumber(reader.uint64() as Long)
              );
            }
          } else {
            message.completedMissions.push(
              longToNumber(reader.uint64() as Long)
            );
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CampaignRecord {
    const message = { ...baseCampaignRecord } as CampaignRecord;
    message.completedMissions = [];
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = Number(object.campaign_id);
    } else {
      message.campaign_id = 0;
    }
    if (object.claimable !== undefined && object.claimable !== null) {
      message.claimable = String(object.claimable);
    } else {
      message.claimable = "";
    }
    if (
      object.completedMissions !== undefined &&
      object.completedMissions !== null
    ) {
      for (const e of object.completedMissions) {
        message.completedMissions.push(Number(e));
      }
    }
    return message;
  },

  toJSON(message: CampaignRecord): unknown {
    const obj: any = {};
    message.campaign_id !== undefined &&
      (obj.campaign_id = message.campaign_id);
    message.claimable !== undefined && (obj.claimable = message.claimable);
    if (message.completedMissions) {
      obj.completedMissions = message.completedMissions.map((e) => e);
    } else {
      obj.completedMissions = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<CampaignRecord>): CampaignRecord {
    const message = { ...baseCampaignRecord } as CampaignRecord;
    message.completedMissions = [];
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = object.campaign_id;
    } else {
      message.campaign_id = 0;
    }
    if (object.claimable !== undefined && object.claimable !== null) {
      message.claimable = object.claimable;
    } else {
      message.claimable = "";
    }
    if (
      object.completedMissions !== undefined &&
      object.completedMissions !== null
    ) {
      for (const e of object.completedMissions) {
        message.completedMissions.push(e);
      }
    }
    return message;
  },
};

const baseClaimRecord: object = { address: "" };

export const ClaimRecord = {
  encode(message: ClaimRecord, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    for (const v of message.campaign_records) {
      CampaignRecord.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): ClaimRecord {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseClaimRecord } as ClaimRecord;
    message.campaign_records = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.campaign_records.push(
            CampaignRecord.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ClaimRecord {
    const message = { ...baseClaimRecord } as ClaimRecord;
    message.campaign_records = [];
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (
      object.campaign_records !== undefined &&
      object.campaign_records !== null
    ) {
      for (const e of object.campaign_records) {
        message.campaign_records.push(CampaignRecord.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: ClaimRecord): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    if (message.campaign_records) {
      obj.campaign_records = message.campaign_records.map((e) =>
        e ? CampaignRecord.toJSON(e) : undefined
      );
    } else {
      obj.campaign_records = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<ClaimRecord>): ClaimRecord {
    const message = { ...baseClaimRecord } as ClaimRecord;
    message.campaign_records = [];
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (
      object.campaign_records !== undefined &&
      object.campaign_records !== null
    ) {
      for (const e of object.campaign_records) {
        message.campaign_records.push(CampaignRecord.fromPartial(e));
      }
    }
    return message;
  },
};

const baseCampaign: object = { campaign_id: 0, description: "" };

export const Campaign = {
  encode(message: Campaign, writer: Writer = Writer.create()): Writer {
    if (message.campaign_id !== 0) {
      writer.uint32(8).uint64(message.campaign_id);
    }
    if (message.description !== "") {
      writer.uint32(18).string(message.description);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Campaign {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCampaign } as Campaign;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaign_id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.description = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Campaign {
    const message = { ...baseCampaign } as Campaign;
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = Number(object.campaign_id);
    } else {
      message.campaign_id = 0;
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    return message;
  },

  toJSON(message: Campaign): unknown {
    const obj: any = {};
    message.campaign_id !== undefined &&
      (obj.campaign_id = message.campaign_id);
    message.description !== undefined &&
      (obj.description = message.description);
    return obj;
  },

  fromPartial(object: DeepPartial<Campaign>): Campaign {
    const message = { ...baseCampaign } as Campaign;
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = object.campaign_id;
    } else {
      message.campaign_id = 0;
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    return message;
  },
};

const baseInitialClaimff: object = {
  enabled: false,
  campaign_id: 0,
  mission_id: 0,
};

export const InitialClaimff = {
  encode(message: InitialClaimff, writer: Writer = Writer.create()): Writer {
    if (message.enabled === true) {
      writer.uint32(8).bool(message.enabled);
    }
    if (message.campaign_id !== 0) {
      writer.uint32(16).uint64(message.campaign_id);
    }
    if (message.mission_id !== 0) {
      writer.uint32(24).uint64(message.mission_id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): InitialClaimff {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseInitialClaimff } as InitialClaimff;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.enabled = reader.bool();
          break;
        case 2:
          message.campaign_id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.mission_id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): InitialClaimff {
    const message = { ...baseInitialClaimff } as InitialClaimff;
    if (object.enabled !== undefined && object.enabled !== null) {
      message.enabled = Boolean(object.enabled);
    } else {
      message.enabled = false;
    }
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = Number(object.campaign_id);
    } else {
      message.campaign_id = 0;
    }
    if (object.mission_id !== undefined && object.mission_id !== null) {
      message.mission_id = Number(object.mission_id);
    } else {
      message.mission_id = 0;
    }
    return message;
  },

  toJSON(message: InitialClaimff): unknown {
    const obj: any = {};
    message.enabled !== undefined && (obj.enabled = message.enabled);
    message.campaign_id !== undefined &&
      (obj.campaign_id = message.campaign_id);
    message.mission_id !== undefined && (obj.mission_id = message.mission_id);
    return obj;
  },

  fromPartial(object: DeepPartial<InitialClaimff>): InitialClaimff {
    const message = { ...baseInitialClaimff } as InitialClaimff;
    if (object.enabled !== undefined && object.enabled !== null) {
      message.enabled = object.enabled;
    } else {
      message.enabled = false;
    }
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = object.campaign_id;
    } else {
      message.campaign_id = 0;
    }
    if (object.mission_id !== undefined && object.mission_id !== null) {
      message.mission_id = object.mission_id;
    } else {
      message.mission_id = 0;
    }
    return message;
  },
};

const baseMission: object = {
  campaign_id: 0,
  mission_id: 0,
  description: "",
  weight: "",
};

export const Mission = {
  encode(message: Mission, writer: Writer = Writer.create()): Writer {
    if (message.campaign_id !== 0) {
      writer.uint32(8).uint64(message.campaign_id);
    }
    if (message.mission_id !== 0) {
      writer.uint32(16).uint64(message.mission_id);
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
          message.campaign_id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.mission_id = longToNumber(reader.uint64() as Long);
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
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = Number(object.campaign_id);
    } else {
      message.campaign_id = 0;
    }
    if (object.mission_id !== undefined && object.mission_id !== null) {
      message.mission_id = Number(object.mission_id);
    } else {
      message.mission_id = 0;
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
    message.campaign_id !== undefined &&
      (obj.campaign_id = message.campaign_id);
    message.mission_id !== undefined && (obj.mission_id = message.mission_id);
    message.description !== undefined &&
      (obj.description = message.description);
    message.weight !== undefined && (obj.weight = message.weight);
    return obj;
  },

  fromPartial(object: DeepPartial<Mission>): Mission {
    const message = { ...baseMission } as Mission;
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = object.campaign_id;
    } else {
      message.campaign_id = 0;
    }
    if (object.mission_id !== undefined && object.mission_id !== null) {
      message.mission_id = object.mission_id;
    } else {
      message.mission_id = 0;
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

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
