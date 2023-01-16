/* eslint-disable */
import {
  MissionType,
  AirdropEntry,
  missionTypeFromJSON,
  missionTypeToJSON,
} from "../cfeairdrop/airdrop";
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Duration } from "../google/protobuf/duration";

export const protobufPackage = "chain4energy.c4echain.cfeairdrop";

export interface MsgClaim {
  claimer: string;
  campaign_id: string;
  mission_id: string;
}

export interface MsgClaimResponse {}

export interface MsgCreateAirdropCampaign {
  owner: string;
  name: string;
  description: string;
  start_time: number;
  end_time: number;
  lockup_period: Duration | undefined;
  vesting_period: Duration | undefined;
}

export interface MsgCreateAirdropCampaignResponse {}

export interface MsgAddMissionToAidropCampaign {
  owner: string;
  campaignId: number;
  name: string;
  description: string;
  missionType: MissionType;
  weight: string;
}

export interface MsgAddMissionToAidropCampaignResponse {}

export interface MsgAddAirdropEntries {
  owner: string;
  campaign_id: number;
  airdrop_entries: AirdropEntry[];
}

export interface MsgAddAirdropEntriesResponse {}

export interface MsgDeleteAirdropEntry {
  creator: string;
  id: number;
}

export interface MsgDeleteAirdropEntryResponse {}

const baseMsgClaim: object = { claimer: "", campaign_id: "", mission_id: "" };

export const MsgClaim = {
  encode(message: MsgClaim, writer: Writer = Writer.create()): Writer {
    if (message.claimer !== "") {
      writer.uint32(10).string(message.claimer);
    }
    if (message.campaign_id !== "") {
      writer.uint32(18).string(message.campaign_id);
    }
    if (message.mission_id !== "") {
      writer.uint32(26).string(message.mission_id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgClaim {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgClaim } as MsgClaim;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.claimer = reader.string();
          break;
        case 2:
          message.campaign_id = reader.string();
          break;
        case 3:
          message.mission_id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClaim {
    const message = { ...baseMsgClaim } as MsgClaim;
    if (object.claimer !== undefined && object.claimer !== null) {
      message.claimer = String(object.claimer);
    } else {
      message.claimer = "";
    }
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = String(object.campaign_id);
    } else {
      message.campaign_id = "";
    }
    if (object.mission_id !== undefined && object.mission_id !== null) {
      message.mission_id = String(object.mission_id);
    } else {
      message.mission_id = "";
    }
    return message;
  },

  toJSON(message: MsgClaim): unknown {
    const obj: any = {};
    message.claimer !== undefined && (obj.claimer = message.claimer);
    message.campaign_id !== undefined &&
      (obj.campaign_id = message.campaign_id);
    message.mission_id !== undefined && (obj.mission_id = message.mission_id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgClaim>): MsgClaim {
    const message = { ...baseMsgClaim } as MsgClaim;
    if (object.claimer !== undefined && object.claimer !== null) {
      message.claimer = object.claimer;
    } else {
      message.claimer = "";
    }
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = object.campaign_id;
    } else {
      message.campaign_id = "";
    }
    if (object.mission_id !== undefined && object.mission_id !== null) {
      message.mission_id = object.mission_id;
    } else {
      message.mission_id = "";
    }
    return message;
  },
};

const baseMsgClaimResponse: object = {};

export const MsgClaimResponse = {
  encode(_: MsgClaimResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgClaimResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgClaimResponse } as MsgClaimResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgClaimResponse {
    const message = { ...baseMsgClaimResponse } as MsgClaimResponse;
    return message;
  },

  toJSON(_: MsgClaimResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgClaimResponse>): MsgClaimResponse {
    const message = { ...baseMsgClaimResponse } as MsgClaimResponse;
    return message;
  },
};

const baseMsgCreateAirdropCampaign: object = {
  owner: "",
  name: "",
  description: "",
  start_time: 0,
  end_time: 0,
};

export const MsgCreateAirdropCampaign = {
  encode(
    message: MsgCreateAirdropCampaign,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.start_time !== 0) {
      writer.uint32(32).int64(message.start_time);
    }
    if (message.end_time !== 0) {
      writer.uint32(40).int64(message.end_time);
    }
    if (message.lockup_period !== undefined) {
      Duration.encode(message.lockup_period, writer.uint32(50).fork()).ldelim();
    }
    if (message.vesting_period !== undefined) {
      Duration.encode(
        message.vesting_period,
        writer.uint32(58).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateAirdropCampaign {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateAirdropCampaign,
    } as MsgCreateAirdropCampaign;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.description = reader.string();
          break;
        case 4:
          message.start_time = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.end_time = longToNumber(reader.int64() as Long);
          break;
        case 6:
          message.lockup_period = Duration.decode(reader, reader.uint32());
          break;
        case 7:
          message.vesting_period = Duration.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateAirdropCampaign {
    const message = {
      ...baseMsgCreateAirdropCampaign,
    } as MsgCreateAirdropCampaign;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.start_time !== undefined && object.start_time !== null) {
      message.start_time = Number(object.start_time);
    } else {
      message.start_time = 0;
    }
    if (object.end_time !== undefined && object.end_time !== null) {
      message.end_time = Number(object.end_time);
    } else {
      message.end_time = 0;
    }
    if (object.lockup_period !== undefined && object.lockup_period !== null) {
      message.lockup_period = Duration.fromJSON(object.lockup_period);
    } else {
      message.lockup_period = undefined;
    }
    if (object.vesting_period !== undefined && object.vesting_period !== null) {
      message.vesting_period = Duration.fromJSON(object.vesting_period);
    } else {
      message.vesting_period = undefined;
    }
    return message;
  },

  toJSON(message: MsgCreateAirdropCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined &&
      (obj.description = message.description);
    message.start_time !== undefined && (obj.start_time = message.start_time);
    message.end_time !== undefined && (obj.end_time = message.end_time);
    message.lockup_period !== undefined &&
      (obj.lockup_period = message.lockup_period
        ? Duration.toJSON(message.lockup_period)
        : undefined);
    message.vesting_period !== undefined &&
      (obj.vesting_period = message.vesting_period
        ? Duration.toJSON(message.vesting_period)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateAirdropCampaign>
  ): MsgCreateAirdropCampaign {
    const message = {
      ...baseMsgCreateAirdropCampaign,
    } as MsgCreateAirdropCampaign;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.start_time !== undefined && object.start_time !== null) {
      message.start_time = object.start_time;
    } else {
      message.start_time = 0;
    }
    if (object.end_time !== undefined && object.end_time !== null) {
      message.end_time = object.end_time;
    } else {
      message.end_time = 0;
    }
    if (object.lockup_period !== undefined && object.lockup_period !== null) {
      message.lockup_period = Duration.fromPartial(object.lockup_period);
    } else {
      message.lockup_period = undefined;
    }
    if (object.vesting_period !== undefined && object.vesting_period !== null) {
      message.vesting_period = Duration.fromPartial(object.vesting_period);
    } else {
      message.vesting_period = undefined;
    }
    return message;
  },
};

const baseMsgCreateAirdropCampaignResponse: object = {};

export const MsgCreateAirdropCampaignResponse = {
  encode(
    _: MsgCreateAirdropCampaignResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateAirdropCampaignResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateAirdropCampaignResponse,
    } as MsgCreateAirdropCampaignResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCreateAirdropCampaignResponse {
    const message = {
      ...baseMsgCreateAirdropCampaignResponse,
    } as MsgCreateAirdropCampaignResponse;
    return message;
  },

  toJSON(_: MsgCreateAirdropCampaignResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCreateAirdropCampaignResponse>
  ): MsgCreateAirdropCampaignResponse {
    const message = {
      ...baseMsgCreateAirdropCampaignResponse,
    } as MsgCreateAirdropCampaignResponse;
    return message;
  },
};

const baseMsgAddMissionToAidropCampaign: object = {
  owner: "",
  campaignId: 0,
  name: "",
  description: "",
  missionType: 0,
  weight: "",
};

export const MsgAddMissionToAidropCampaign = {
  encode(
    message: MsgAddMissionToAidropCampaign,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(34).string(message.description);
    }
    if (message.missionType !== 0) {
      writer.uint32(40).int32(message.missionType);
    }
    if (message.weight !== "") {
      writer.uint32(50).string(message.weight);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgAddMissionToAidropCampaign {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAddMissionToAidropCampaign,
    } as MsgAddMissionToAidropCampaign;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.campaignId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.name = reader.string();
          break;
        case 4:
          message.description = reader.string();
          break;
        case 5:
          message.missionType = reader.int32() as any;
          break;
        case 6:
          message.weight = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddMissionToAidropCampaign {
    const message = {
      ...baseMsgAddMissionToAidropCampaign,
    } as MsgAddMissionToAidropCampaign;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.campaignId !== undefined && object.campaignId !== null) {
      message.campaignId = Number(object.campaignId);
    } else {
      message.campaignId = 0;
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = String(object.description);
    } else {
      message.description = "";
    }
    if (object.missionType !== undefined && object.missionType !== null) {
      message.missionType = missionTypeFromJSON(object.missionType);
    } else {
      message.missionType = 0;
    }
    if (object.weight !== undefined && object.weight !== null) {
      message.weight = String(object.weight);
    } else {
      message.weight = "";
    }
    return message;
  },

  toJSON(message: MsgAddMissionToAidropCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined &&
      (obj.description = message.description);
    message.missionType !== undefined &&
      (obj.missionType = missionTypeToJSON(message.missionType));
    message.weight !== undefined && (obj.weight = message.weight);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgAddMissionToAidropCampaign>
  ): MsgAddMissionToAidropCampaign {
    const message = {
      ...baseMsgAddMissionToAidropCampaign,
    } as MsgAddMissionToAidropCampaign;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.campaignId !== undefined && object.campaignId !== null) {
      message.campaignId = object.campaignId;
    } else {
      message.campaignId = 0;
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.description !== undefined && object.description !== null) {
      message.description = object.description;
    } else {
      message.description = "";
    }
    if (object.missionType !== undefined && object.missionType !== null) {
      message.missionType = object.missionType;
    } else {
      message.missionType = 0;
    }
    if (object.weight !== undefined && object.weight !== null) {
      message.weight = object.weight;
    } else {
      message.weight = "";
    }
    return message;
  },
};

const baseMsgAddMissionToAidropCampaignResponse: object = {};

export const MsgAddMissionToAidropCampaignResponse = {
  encode(
    _: MsgAddMissionToAidropCampaignResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgAddMissionToAidropCampaignResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAddMissionToAidropCampaignResponse,
    } as MsgAddMissionToAidropCampaignResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgAddMissionToAidropCampaignResponse {
    const message = {
      ...baseMsgAddMissionToAidropCampaignResponse,
    } as MsgAddMissionToAidropCampaignResponse;
    return message;
  },

  toJSON(_: MsgAddMissionToAidropCampaignResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgAddMissionToAidropCampaignResponse>
  ): MsgAddMissionToAidropCampaignResponse {
    const message = {
      ...baseMsgAddMissionToAidropCampaignResponse,
    } as MsgAddMissionToAidropCampaignResponse;
    return message;
  },
};

const baseMsgAddAirdropEntries: object = { owner: "", campaign_id: 0 };

export const MsgAddAirdropEntries = {
  encode(
    message: MsgAddAirdropEntries,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaign_id !== 0) {
      writer.uint32(16).uint64(message.campaign_id);
    }
    for (const v of message.airdrop_entries) {
      AirdropEntry.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddAirdropEntries {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddAirdropEntries } as MsgAddAirdropEntries;
    message.airdrop_entries = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.campaign_id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.airdrop_entries.push(
            AirdropEntry.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddAirdropEntries {
    const message = { ...baseMsgAddAirdropEntries } as MsgAddAirdropEntries;
    message.airdrop_entries = [];
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = Number(object.campaign_id);
    } else {
      message.campaign_id = 0;
    }
    if (
      object.airdrop_entries !== undefined &&
      object.airdrop_entries !== null
    ) {
      for (const e of object.airdrop_entries) {
        message.airdrop_entries.push(AirdropEntry.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: MsgAddAirdropEntries): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaign_id !== undefined &&
      (obj.campaign_id = message.campaign_id);
    if (message.airdrop_entries) {
      obj.airdrop_entries = message.airdrop_entries.map((e) =>
        e ? AirdropEntry.toJSON(e) : undefined
      );
    } else {
      obj.airdrop_entries = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddAirdropEntries>): MsgAddAirdropEntries {
    const message = { ...baseMsgAddAirdropEntries } as MsgAddAirdropEntries;
    message.airdrop_entries = [];
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = object.campaign_id;
    } else {
      message.campaign_id = 0;
    }
    if (
      object.airdrop_entries !== undefined &&
      object.airdrop_entries !== null
    ) {
      for (const e of object.airdrop_entries) {
        message.airdrop_entries.push(AirdropEntry.fromPartial(e));
      }
    }
    return message;
  },
};

const baseMsgAddAirdropEntriesResponse: object = {};

export const MsgAddAirdropEntriesResponse = {
  encode(
    _: MsgAddAirdropEntriesResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgAddAirdropEntriesResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAddAirdropEntriesResponse,
    } as MsgAddAirdropEntriesResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgAddAirdropEntriesResponse {
    const message = {
      ...baseMsgAddAirdropEntriesResponse,
    } as MsgAddAirdropEntriesResponse;
    return message;
  },

  toJSON(_: MsgAddAirdropEntriesResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgAddAirdropEntriesResponse>
  ): MsgAddAirdropEntriesResponse {
    const message = {
      ...baseMsgAddAirdropEntriesResponse,
    } as MsgAddAirdropEntriesResponse;
    return message;
  },
};

const baseMsgDeleteAirdropEntry: object = { creator: "", id: 0 };

export const MsgDeleteAirdropEntry = {
  encode(
    message: MsgDeleteAirdropEntry,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteAirdropEntry {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteAirdropEntry } as MsgDeleteAirdropEntry;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteAirdropEntry {
    const message = { ...baseMsgDeleteAirdropEntry } as MsgDeleteAirdropEntry;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgDeleteAirdropEntry): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgDeleteAirdropEntry>
  ): MsgDeleteAirdropEntry {
    const message = { ...baseMsgDeleteAirdropEntry } as MsgDeleteAirdropEntry;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgDeleteAirdropEntryResponse: object = {};

export const MsgDeleteAirdropEntryResponse = {
  encode(
    _: MsgDeleteAirdropEntryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteAirdropEntryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteAirdropEntryResponse,
    } as MsgDeleteAirdropEntryResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgDeleteAirdropEntryResponse {
    const message = {
      ...baseMsgDeleteAirdropEntryResponse,
    } as MsgDeleteAirdropEntryResponse;
    return message;
  },

  toJSON(_: MsgDeleteAirdropEntryResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteAirdropEntryResponse>
  ): MsgDeleteAirdropEntryResponse {
    const message = {
      ...baseMsgDeleteAirdropEntryResponse,
    } as MsgDeleteAirdropEntryResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Claim(request: MsgClaim): Promise<MsgClaimResponse>;
  /** rpc InitialClaim(MsgInitialClaim) returns (MsgInitialClaimResponse); */
  CreateAirdropCampaign(
    request: MsgCreateAirdropCampaign
  ): Promise<MsgCreateAirdropCampaignResponse>;
  AddMissionToAidropCampaign(
    request: MsgAddMissionToAidropCampaign
  ): Promise<MsgAddMissionToAidropCampaignResponse>;
  AddAirdropEntries(
    request: MsgAddAirdropEntries
  ): Promise<MsgAddAirdropEntriesResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  DeleteAirdropEntry(
    request: MsgDeleteAirdropEntry
  ): Promise<MsgDeleteAirdropEntryResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Claim(request: MsgClaim): Promise<MsgClaimResponse> {
    const data = MsgClaim.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Msg",
      "Claim",
      data
    );
    return promise.then((data) => MsgClaimResponse.decode(new Reader(data)));
  }

  CreateAirdropCampaign(
    request: MsgCreateAirdropCampaign
  ): Promise<MsgCreateAirdropCampaignResponse> {
    const data = MsgCreateAirdropCampaign.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Msg",
      "CreateAirdropCampaign",
      data
    );
    return promise.then((data) =>
      MsgCreateAirdropCampaignResponse.decode(new Reader(data))
    );
  }

  AddMissionToAidropCampaign(
    request: MsgAddMissionToAidropCampaign
  ): Promise<MsgAddMissionToAidropCampaignResponse> {
    const data = MsgAddMissionToAidropCampaign.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Msg",
      "AddMissionToAidropCampaign",
      data
    );
    return promise.then((data) =>
      MsgAddMissionToAidropCampaignResponse.decode(new Reader(data))
    );
  }

  AddAirdropEntries(
    request: MsgAddAirdropEntries
  ): Promise<MsgAddAirdropEntriesResponse> {
    const data = MsgAddAirdropEntries.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Msg",
      "AddAirdropEntries",
      data
    );
    return promise.then((data) =>
      MsgAddAirdropEntriesResponse.decode(new Reader(data))
    );
  }

  DeleteAirdropEntry(
    request: MsgDeleteAirdropEntry
  ): Promise<MsgDeleteAirdropEntryResponse> {
    const data = MsgDeleteAirdropEntry.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Msg",
      "DeleteAirdropEntry",
      data
    );
    return promise.then((data) =>
      MsgDeleteAirdropEntryResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
