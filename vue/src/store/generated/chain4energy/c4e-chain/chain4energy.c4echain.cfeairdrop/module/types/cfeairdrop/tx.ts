/* eslint-disable */
import {
  MissionType,
  AirdropCloseAction,
  ClaimRecord,
  missionTypeFromJSON,
  missionTypeToJSON,
  airdropCloseActionFromJSON,
  airdropCloseActionToJSON,
} from "../cfeairdrop/airdrop";
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import { Timestamp } from "../google/protobuf/timestamp";
import * as Long from "long";
import { Duration } from "../google/protobuf/duration";

export const protobufPackage = "chain4energy.c4echain.cfeairdrop";

export interface MsgClaim {
  claimer: string;
  campaign_id: number;
  mission_id: number;
}

export interface MsgClaimResponse {}

export interface MsgInitialClaim {
  claimer: string;
  campaign_id: number;
  addressToClaim: string;
}

export interface MsgInitialClaimResponse {}

export interface MsgCreateAirdropCampaign {
  owner: string;
  name: string;
  description: string;
  feegrant_amount: string;
  initial_claim_free_amount: string;
  start_time: Date | undefined;
  end_time: Date | undefined;
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
  claim_start_date: Date | undefined;
}

export interface MsgAddMissionToAidropCampaignResponse {}

export interface MsgAddClaimRecords {
  owner: string;
  campaign_id: number;
  claim_records: ClaimRecord[];
}

export interface MsgAddClaimRecordsResponse {}

export interface MsgDeleteClaimRecord {
  owner: string;
  campaignId: number;
  userAddress: string;
}

export interface MsgDeleteClaimRecordResponse {}

export interface MsgCloseAirdropCampaign {
  owner: string;
  campaign_id: number;
  airdrop_close_action: AirdropCloseAction;
}

export interface MsgCloseAirdropCampaignResponse {}

export interface MsgStartAirdropCampaign {
  owner: string;
  campaignId: number;
}

export interface MsgStartAirdropCampaignResponse {}

export interface MsgRemoveAirdropCampaign {
  owner: string;
  campaignId: number;
}

export interface MsgRemoveAirdropCampaignResponse {}

export interface MsgEditAirdropCampaign {
  owner: string;
  campaignId: number;
  name: string;
  description: string;
  feegrant_amount: string;
  initial_claim_free_amount: string;
  start_time: Date | undefined;
  end_time: Date | undefined;
  lockup_period: Duration | undefined;
  vesting_period: Duration | undefined;
}

export interface MsgEditAirdropCampaignResponse {}

const baseMsgClaim: object = { claimer: "", campaign_id: 0, mission_id: 0 };

export const MsgClaim = {
  encode(message: MsgClaim, writer: Writer = Writer.create()): Writer {
    if (message.claimer !== "") {
      writer.uint32(10).string(message.claimer);
    }
    if (message.campaign_id !== 0) {
      writer.uint32(16).uint64(message.campaign_id);
    }
    if (message.mission_id !== 0) {
      writer.uint32(24).uint64(message.mission_id);
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

  fromJSON(object: any): MsgClaim {
    const message = { ...baseMsgClaim } as MsgClaim;
    if (object.claimer !== undefined && object.claimer !== null) {
      message.claimer = String(object.claimer);
    } else {
      message.claimer = "";
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

const baseMsgInitialClaim: object = {
  claimer: "",
  campaign_id: 0,
  addressToClaim: "",
};

export const MsgInitialClaim = {
  encode(message: MsgInitialClaim, writer: Writer = Writer.create()): Writer {
    if (message.claimer !== "") {
      writer.uint32(10).string(message.claimer);
    }
    if (message.campaign_id !== 0) {
      writer.uint32(16).uint64(message.campaign_id);
    }
    if (message.addressToClaim !== "") {
      writer.uint32(34).string(message.addressToClaim);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgInitialClaim {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgInitialClaim } as MsgInitialClaim;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.claimer = reader.string();
          break;
        case 2:
          message.campaign_id = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.addressToClaim = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInitialClaim {
    const message = { ...baseMsgInitialClaim } as MsgInitialClaim;
    if (object.claimer !== undefined && object.claimer !== null) {
      message.claimer = String(object.claimer);
    } else {
      message.claimer = "";
    }
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = Number(object.campaign_id);
    } else {
      message.campaign_id = 0;
    }
    if (object.addressToClaim !== undefined && object.addressToClaim !== null) {
      message.addressToClaim = String(object.addressToClaim);
    } else {
      message.addressToClaim = "";
    }
    return message;
  },

  toJSON(message: MsgInitialClaim): unknown {
    const obj: any = {};
    message.claimer !== undefined && (obj.claimer = message.claimer);
    message.campaign_id !== undefined &&
      (obj.campaign_id = message.campaign_id);
    message.addressToClaim !== undefined &&
      (obj.addressToClaim = message.addressToClaim);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgInitialClaim>): MsgInitialClaim {
    const message = { ...baseMsgInitialClaim } as MsgInitialClaim;
    if (object.claimer !== undefined && object.claimer !== null) {
      message.claimer = object.claimer;
    } else {
      message.claimer = "";
    }
    if (object.campaign_id !== undefined && object.campaign_id !== null) {
      message.campaign_id = object.campaign_id;
    } else {
      message.campaign_id = 0;
    }
    if (object.addressToClaim !== undefined && object.addressToClaim !== null) {
      message.addressToClaim = object.addressToClaim;
    } else {
      message.addressToClaim = "";
    }
    return message;
  },
};

const baseMsgInitialClaimResponse: object = {};

export const MsgInitialClaimResponse = {
  encode(_: MsgInitialClaimResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgInitialClaimResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgInitialClaimResponse,
    } as MsgInitialClaimResponse;
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

  fromJSON(_: any): MsgInitialClaimResponse {
    const message = {
      ...baseMsgInitialClaimResponse,
    } as MsgInitialClaimResponse;
    return message;
  },

  toJSON(_: MsgInitialClaimResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgInitialClaimResponse>
  ): MsgInitialClaimResponse {
    const message = {
      ...baseMsgInitialClaimResponse,
    } as MsgInitialClaimResponse;
    return message;
  },
};

const baseMsgCreateAirdropCampaign: object = {
  owner: "",
  name: "",
  description: "",
  feegrant_amount: "",
  initial_claim_free_amount: "",
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
    if (message.feegrant_amount !== "") {
      writer.uint32(34).string(message.feegrant_amount);
    }
    if (message.initial_claim_free_amount !== "") {
      writer.uint32(42).string(message.initial_claim_free_amount);
    }
    if (message.start_time !== undefined) {
      Timestamp.encode(
        toTimestamp(message.start_time),
        writer.uint32(50).fork()
      ).ldelim();
    }
    if (message.end_time !== undefined) {
      Timestamp.encode(
        toTimestamp(message.end_time),
        writer.uint32(58).fork()
      ).ldelim();
    }
    if (message.lockup_period !== undefined) {
      Duration.encode(message.lockup_period, writer.uint32(66).fork()).ldelim();
    }
    if (message.vesting_period !== undefined) {
      Duration.encode(
        message.vesting_period,
        writer.uint32(74).fork()
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
          message.feegrant_amount = reader.string();
          break;
        case 5:
          message.initial_claim_free_amount = reader.string();
          break;
        case 6:
          message.start_time = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 7:
          message.end_time = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 8:
          message.lockup_period = Duration.decode(reader, reader.uint32());
          break;
        case 9:
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
    if (
      object.feegrant_amount !== undefined &&
      object.feegrant_amount !== null
    ) {
      message.feegrant_amount = String(object.feegrant_amount);
    } else {
      message.feegrant_amount = "";
    }
    if (
      object.initial_claim_free_amount !== undefined &&
      object.initial_claim_free_amount !== null
    ) {
      message.initial_claim_free_amount = String(
        object.initial_claim_free_amount
      );
    } else {
      message.initial_claim_free_amount = "";
    }
    if (object.start_time !== undefined && object.start_time !== null) {
      message.start_time = fromJsonTimestamp(object.start_time);
    } else {
      message.start_time = undefined;
    }
    if (object.end_time !== undefined && object.end_time !== null) {
      message.end_time = fromJsonTimestamp(object.end_time);
    } else {
      message.end_time = undefined;
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
    message.feegrant_amount !== undefined &&
      (obj.feegrant_amount = message.feegrant_amount);
    message.initial_claim_free_amount !== undefined &&
      (obj.initial_claim_free_amount = message.initial_claim_free_amount);
    message.start_time !== undefined &&
      (obj.start_time =
        message.start_time !== undefined
          ? message.start_time.toISOString()
          : null);
    message.end_time !== undefined &&
      (obj.end_time =
        message.end_time !== undefined ? message.end_time.toISOString() : null);
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
    if (
      object.feegrant_amount !== undefined &&
      object.feegrant_amount !== null
    ) {
      message.feegrant_amount = object.feegrant_amount;
    } else {
      message.feegrant_amount = "";
    }
    if (
      object.initial_claim_free_amount !== undefined &&
      object.initial_claim_free_amount !== null
    ) {
      message.initial_claim_free_amount = object.initial_claim_free_amount;
    } else {
      message.initial_claim_free_amount = "";
    }
    if (object.start_time !== undefined && object.start_time !== null) {
      message.start_time = object.start_time;
    } else {
      message.start_time = undefined;
    }
    if (object.end_time !== undefined && object.end_time !== null) {
      message.end_time = object.end_time;
    } else {
      message.end_time = undefined;
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
    if (message.claim_start_date !== undefined) {
      Timestamp.encode(
        toTimestamp(message.claim_start_date),
        writer.uint32(58).fork()
      ).ldelim();
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
        case 7:
          message.claim_start_date = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
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
    if (
      object.claim_start_date !== undefined &&
      object.claim_start_date !== null
    ) {
      message.claim_start_date = fromJsonTimestamp(object.claim_start_date);
    } else {
      message.claim_start_date = undefined;
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
    message.claim_start_date !== undefined &&
      (obj.claim_start_date =
        message.claim_start_date !== undefined
          ? message.claim_start_date.toISOString()
          : null);
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
    if (
      object.claim_start_date !== undefined &&
      object.claim_start_date !== null
    ) {
      message.claim_start_date = object.claim_start_date;
    } else {
      message.claim_start_date = undefined;
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

const baseMsgAddClaimRecords: object = { owner: "", campaign_id: 0 };

export const MsgAddClaimRecords = {
  encode(
    message: MsgAddClaimRecords,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaign_id !== 0) {
      writer.uint32(16).uint64(message.campaign_id);
    }
    for (const v of message.claim_records) {
      ClaimRecord.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAddClaimRecords {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgAddClaimRecords } as MsgAddClaimRecords;
    message.claim_records = [];
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
          message.claim_records.push(
            ClaimRecord.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddClaimRecords {
    const message = { ...baseMsgAddClaimRecords } as MsgAddClaimRecords;
    message.claim_records = [];
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
    if (object.claim_records !== undefined && object.claim_records !== null) {
      for (const e of object.claim_records) {
        message.claim_records.push(ClaimRecord.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: MsgAddClaimRecords): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaign_id !== undefined &&
      (obj.campaign_id = message.campaign_id);
    if (message.claim_records) {
      obj.claim_records = message.claim_records.map((e) =>
        e ? ClaimRecord.toJSON(e) : undefined
      );
    } else {
      obj.claim_records = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgAddClaimRecords>): MsgAddClaimRecords {
    const message = { ...baseMsgAddClaimRecords } as MsgAddClaimRecords;
    message.claim_records = [];
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
    if (object.claim_records !== undefined && object.claim_records !== null) {
      for (const e of object.claim_records) {
        message.claim_records.push(ClaimRecord.fromPartial(e));
      }
    }
    return message;
  },
};

const baseMsgAddClaimRecordsResponse: object = {};

export const MsgAddClaimRecordsResponse = {
  encode(
    _: MsgAddClaimRecordsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgAddClaimRecordsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgAddClaimRecordsResponse,
    } as MsgAddClaimRecordsResponse;
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

  fromJSON(_: any): MsgAddClaimRecordsResponse {
    const message = {
      ...baseMsgAddClaimRecordsResponse,
    } as MsgAddClaimRecordsResponse;
    return message;
  },

  toJSON(_: MsgAddClaimRecordsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgAddClaimRecordsResponse>
  ): MsgAddClaimRecordsResponse {
    const message = {
      ...baseMsgAddClaimRecordsResponse,
    } as MsgAddClaimRecordsResponse;
    return message;
  },
};

const baseMsgDeleteClaimRecord: object = {
  owner: "",
  campaignId: 0,
  userAddress: "",
};

export const MsgDeleteClaimRecord = {
  encode(
    message: MsgDeleteClaimRecord,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    if (message.userAddress !== "") {
      writer.uint32(26).string(message.userAddress);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgDeleteClaimRecord {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgDeleteClaimRecord } as MsgDeleteClaimRecord;
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
          message.userAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgDeleteClaimRecord {
    const message = { ...baseMsgDeleteClaimRecord } as MsgDeleteClaimRecord;
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
    if (object.userAddress !== undefined && object.userAddress !== null) {
      message.userAddress = String(object.userAddress);
    } else {
      message.userAddress = "";
    }
    return message;
  },

  toJSON(message: MsgDeleteClaimRecord): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    message.userAddress !== undefined &&
      (obj.userAddress = message.userAddress);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgDeleteClaimRecord>): MsgDeleteClaimRecord {
    const message = { ...baseMsgDeleteClaimRecord } as MsgDeleteClaimRecord;
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
    if (object.userAddress !== undefined && object.userAddress !== null) {
      message.userAddress = object.userAddress;
    } else {
      message.userAddress = "";
    }
    return message;
  },
};

const baseMsgDeleteClaimRecordResponse: object = {};

export const MsgDeleteClaimRecordResponse = {
  encode(
    _: MsgDeleteClaimRecordResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgDeleteClaimRecordResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgDeleteClaimRecordResponse,
    } as MsgDeleteClaimRecordResponse;
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

  fromJSON(_: any): MsgDeleteClaimRecordResponse {
    const message = {
      ...baseMsgDeleteClaimRecordResponse,
    } as MsgDeleteClaimRecordResponse;
    return message;
  },

  toJSON(_: MsgDeleteClaimRecordResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgDeleteClaimRecordResponse>
  ): MsgDeleteClaimRecordResponse {
    const message = {
      ...baseMsgDeleteClaimRecordResponse,
    } as MsgDeleteClaimRecordResponse;
    return message;
  },
};

const baseMsgCloseAirdropCampaign: object = {
  owner: "",
  campaign_id: 0,
  airdrop_close_action: 0,
};

export const MsgCloseAirdropCampaign = {
  encode(
    message: MsgCloseAirdropCampaign,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaign_id !== 0) {
      writer.uint32(16).uint64(message.campaign_id);
    }
    if (message.airdrop_close_action !== 0) {
      writer.uint32(24).int32(message.airdrop_close_action);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCloseAirdropCampaign {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCloseAirdropCampaign,
    } as MsgCloseAirdropCampaign;
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
          message.airdrop_close_action = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCloseAirdropCampaign {
    const message = {
      ...baseMsgCloseAirdropCampaign,
    } as MsgCloseAirdropCampaign;
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
      object.airdrop_close_action !== undefined &&
      object.airdrop_close_action !== null
    ) {
      message.airdrop_close_action = airdropCloseActionFromJSON(
        object.airdrop_close_action
      );
    } else {
      message.airdrop_close_action = 0;
    }
    return message;
  },

  toJSON(message: MsgCloseAirdropCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaign_id !== undefined &&
      (obj.campaign_id = message.campaign_id);
    message.airdrop_close_action !== undefined &&
      (obj.airdrop_close_action = airdropCloseActionToJSON(
        message.airdrop_close_action
      ));
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCloseAirdropCampaign>
  ): MsgCloseAirdropCampaign {
    const message = {
      ...baseMsgCloseAirdropCampaign,
    } as MsgCloseAirdropCampaign;
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
      object.airdrop_close_action !== undefined &&
      object.airdrop_close_action !== null
    ) {
      message.airdrop_close_action = object.airdrop_close_action;
    } else {
      message.airdrop_close_action = 0;
    }
    return message;
  },
};

const baseMsgCloseAirdropCampaignResponse: object = {};

export const MsgCloseAirdropCampaignResponse = {
  encode(
    _: MsgCloseAirdropCampaignResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCloseAirdropCampaignResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCloseAirdropCampaignResponse,
    } as MsgCloseAirdropCampaignResponse;
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

  fromJSON(_: any): MsgCloseAirdropCampaignResponse {
    const message = {
      ...baseMsgCloseAirdropCampaignResponse,
    } as MsgCloseAirdropCampaignResponse;
    return message;
  },

  toJSON(_: MsgCloseAirdropCampaignResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCloseAirdropCampaignResponse>
  ): MsgCloseAirdropCampaignResponse {
    const message = {
      ...baseMsgCloseAirdropCampaignResponse,
    } as MsgCloseAirdropCampaignResponse;
    return message;
  },
};

const baseMsgStartAirdropCampaign: object = { owner: "", campaignId: 0 };

export const MsgStartAirdropCampaign = {
  encode(
    message: MsgStartAirdropCampaign,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgStartAirdropCampaign {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgStartAirdropCampaign,
    } as MsgStartAirdropCampaign;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.campaignId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgStartAirdropCampaign {
    const message = {
      ...baseMsgStartAirdropCampaign,
    } as MsgStartAirdropCampaign;
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
    return message;
  },

  toJSON(message: MsgStartAirdropCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgStartAirdropCampaign>
  ): MsgStartAirdropCampaign {
    const message = {
      ...baseMsgStartAirdropCampaign,
    } as MsgStartAirdropCampaign;
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
    return message;
  },
};

const baseMsgStartAirdropCampaignResponse: object = {};

export const MsgStartAirdropCampaignResponse = {
  encode(
    _: MsgStartAirdropCampaignResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgStartAirdropCampaignResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgStartAirdropCampaignResponse,
    } as MsgStartAirdropCampaignResponse;
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

  fromJSON(_: any): MsgStartAirdropCampaignResponse {
    const message = {
      ...baseMsgStartAirdropCampaignResponse,
    } as MsgStartAirdropCampaignResponse;
    return message;
  },

  toJSON(_: MsgStartAirdropCampaignResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgStartAirdropCampaignResponse>
  ): MsgStartAirdropCampaignResponse {
    const message = {
      ...baseMsgStartAirdropCampaignResponse,
    } as MsgStartAirdropCampaignResponse;
    return message;
  },
};

const baseMsgRemoveAirdropCampaign: object = { owner: "", campaignId: 0 };

export const MsgRemoveAirdropCampaign = {
  encode(
    message: MsgRemoveAirdropCampaign,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRemoveAirdropCampaign {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRemoveAirdropCampaign,
    } as MsgRemoveAirdropCampaign;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.campaignId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRemoveAirdropCampaign {
    const message = {
      ...baseMsgRemoveAirdropCampaign,
    } as MsgRemoveAirdropCampaign;
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
    return message;
  },

  toJSON(message: MsgRemoveAirdropCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgRemoveAirdropCampaign>
  ): MsgRemoveAirdropCampaign {
    const message = {
      ...baseMsgRemoveAirdropCampaign,
    } as MsgRemoveAirdropCampaign;
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
    return message;
  },
};

const baseMsgRemoveAirdropCampaignResponse: object = {};

export const MsgRemoveAirdropCampaignResponse = {
  encode(
    _: MsgRemoveAirdropCampaignResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRemoveAirdropCampaignResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRemoveAirdropCampaignResponse,
    } as MsgRemoveAirdropCampaignResponse;
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

  fromJSON(_: any): MsgRemoveAirdropCampaignResponse {
    const message = {
      ...baseMsgRemoveAirdropCampaignResponse,
    } as MsgRemoveAirdropCampaignResponse;
    return message;
  },

  toJSON(_: MsgRemoveAirdropCampaignResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRemoveAirdropCampaignResponse>
  ): MsgRemoveAirdropCampaignResponse {
    const message = {
      ...baseMsgRemoveAirdropCampaignResponse,
    } as MsgRemoveAirdropCampaignResponse;
    return message;
  },
};

const baseMsgEditAirdropCampaign: object = {
  owner: "",
  campaignId: 0,
  name: "",
  description: "",
  feegrant_amount: "",
  initial_claim_free_amount: "",
};

export const MsgEditAirdropCampaign = {
  encode(
    message: MsgEditAirdropCampaign,
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
    if (message.feegrant_amount !== "") {
      writer.uint32(42).string(message.feegrant_amount);
    }
    if (message.initial_claim_free_amount !== "") {
      writer.uint32(50).string(message.initial_claim_free_amount);
    }
    if (message.start_time !== undefined) {
      Timestamp.encode(
        toTimestamp(message.start_time),
        writer.uint32(58).fork()
      ).ldelim();
    }
    if (message.end_time !== undefined) {
      Timestamp.encode(
        toTimestamp(message.end_time),
        writer.uint32(66).fork()
      ).ldelim();
    }
    if (message.lockup_period !== undefined) {
      Duration.encode(message.lockup_period, writer.uint32(74).fork()).ldelim();
    }
    if (message.vesting_period !== undefined) {
      Duration.encode(
        message.vesting_period,
        writer.uint32(82).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgEditAirdropCampaign {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgEditAirdropCampaign } as MsgEditAirdropCampaign;
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
          message.feegrant_amount = reader.string();
          break;
        case 6:
          message.initial_claim_free_amount = reader.string();
          break;
        case 7:
          message.start_time = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 8:
          message.end_time = fromTimestamp(
            Timestamp.decode(reader, reader.uint32())
          );
          break;
        case 9:
          message.lockup_period = Duration.decode(reader, reader.uint32());
          break;
        case 10:
          message.vesting_period = Duration.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEditAirdropCampaign {
    const message = { ...baseMsgEditAirdropCampaign } as MsgEditAirdropCampaign;
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
    if (
      object.feegrant_amount !== undefined &&
      object.feegrant_amount !== null
    ) {
      message.feegrant_amount = String(object.feegrant_amount);
    } else {
      message.feegrant_amount = "";
    }
    if (
      object.initial_claim_free_amount !== undefined &&
      object.initial_claim_free_amount !== null
    ) {
      message.initial_claim_free_amount = String(
        object.initial_claim_free_amount
      );
    } else {
      message.initial_claim_free_amount = "";
    }
    if (object.start_time !== undefined && object.start_time !== null) {
      message.start_time = fromJsonTimestamp(object.start_time);
    } else {
      message.start_time = undefined;
    }
    if (object.end_time !== undefined && object.end_time !== null) {
      message.end_time = fromJsonTimestamp(object.end_time);
    } else {
      message.end_time = undefined;
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

  toJSON(message: MsgEditAirdropCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined &&
      (obj.description = message.description);
    message.feegrant_amount !== undefined &&
      (obj.feegrant_amount = message.feegrant_amount);
    message.initial_claim_free_amount !== undefined &&
      (obj.initial_claim_free_amount = message.initial_claim_free_amount);
    message.start_time !== undefined &&
      (obj.start_time =
        message.start_time !== undefined
          ? message.start_time.toISOString()
          : null);
    message.end_time !== undefined &&
      (obj.end_time =
        message.end_time !== undefined ? message.end_time.toISOString() : null);
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
    object: DeepPartial<MsgEditAirdropCampaign>
  ): MsgEditAirdropCampaign {
    const message = { ...baseMsgEditAirdropCampaign } as MsgEditAirdropCampaign;
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
    if (
      object.feegrant_amount !== undefined &&
      object.feegrant_amount !== null
    ) {
      message.feegrant_amount = object.feegrant_amount;
    } else {
      message.feegrant_amount = "";
    }
    if (
      object.initial_claim_free_amount !== undefined &&
      object.initial_claim_free_amount !== null
    ) {
      message.initial_claim_free_amount = object.initial_claim_free_amount;
    } else {
      message.initial_claim_free_amount = "";
    }
    if (object.start_time !== undefined && object.start_time !== null) {
      message.start_time = object.start_time;
    } else {
      message.start_time = undefined;
    }
    if (object.end_time !== undefined && object.end_time !== null) {
      message.end_time = object.end_time;
    } else {
      message.end_time = undefined;
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

const baseMsgEditAirdropCampaignResponse: object = {};

export const MsgEditAirdropCampaignResponse = {
  encode(
    _: MsgEditAirdropCampaignResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgEditAirdropCampaignResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgEditAirdropCampaignResponse,
    } as MsgEditAirdropCampaignResponse;
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

  fromJSON(_: any): MsgEditAirdropCampaignResponse {
    const message = {
      ...baseMsgEditAirdropCampaignResponse,
    } as MsgEditAirdropCampaignResponse;
    return message;
  },

  toJSON(_: MsgEditAirdropCampaignResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgEditAirdropCampaignResponse>
  ): MsgEditAirdropCampaignResponse {
    const message = {
      ...baseMsgEditAirdropCampaignResponse,
    } as MsgEditAirdropCampaignResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Claim(request: MsgClaim): Promise<MsgClaimResponse>;
  InitialClaim(request: MsgInitialClaim): Promise<MsgInitialClaimResponse>;
  CreateAirdropCampaign(
    request: MsgCreateAirdropCampaign
  ): Promise<MsgCreateAirdropCampaignResponse>;
  EditAirdropCampaign(
    request: MsgEditAirdropCampaign
  ): Promise<MsgEditAirdropCampaignResponse>;
  AddMissionToAidropCampaign(
    request: MsgAddMissionToAidropCampaign
  ): Promise<MsgAddMissionToAidropCampaignResponse>;
  AddClaimRecords(
    request: MsgAddClaimRecords
  ): Promise<MsgAddClaimRecordsResponse>;
  DeleteClaimRecord(
    request: MsgDeleteClaimRecord
  ): Promise<MsgDeleteClaimRecordResponse>;
  CloseAirdropCampaign(
    request: MsgCloseAirdropCampaign
  ): Promise<MsgCloseAirdropCampaignResponse>;
  StartAirdropCampaign(
    request: MsgStartAirdropCampaign
  ): Promise<MsgStartAirdropCampaignResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RemoveAirdropCampaign(
    request: MsgRemoveAirdropCampaign
  ): Promise<MsgRemoveAirdropCampaignResponse>;
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

  InitialClaim(request: MsgInitialClaim): Promise<MsgInitialClaimResponse> {
    const data = MsgInitialClaim.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Msg",
      "InitialClaim",
      data
    );
    return promise.then((data) =>
      MsgInitialClaimResponse.decode(new Reader(data))
    );
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

  EditAirdropCampaign(
    request: MsgEditAirdropCampaign
  ): Promise<MsgEditAirdropCampaignResponse> {
    const data = MsgEditAirdropCampaign.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Msg",
      "EditAirdropCampaign",
      data
    );
    return promise.then((data) =>
      MsgEditAirdropCampaignResponse.decode(new Reader(data))
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

  AddClaimRecords(
    request: MsgAddClaimRecords
  ): Promise<MsgAddClaimRecordsResponse> {
    const data = MsgAddClaimRecords.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Msg",
      "AddClaimRecords",
      data
    );
    return promise.then((data) =>
      MsgAddClaimRecordsResponse.decode(new Reader(data))
    );
  }

  DeleteClaimRecord(
    request: MsgDeleteClaimRecord
  ): Promise<MsgDeleteClaimRecordResponse> {
    const data = MsgDeleteClaimRecord.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Msg",
      "DeleteClaimRecord",
      data
    );
    return promise.then((data) =>
      MsgDeleteClaimRecordResponse.decode(new Reader(data))
    );
  }

  CloseAirdropCampaign(
    request: MsgCloseAirdropCampaign
  ): Promise<MsgCloseAirdropCampaignResponse> {
    const data = MsgCloseAirdropCampaign.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Msg",
      "CloseAirdropCampaign",
      data
    );
    return promise.then((data) =>
      MsgCloseAirdropCampaignResponse.decode(new Reader(data))
    );
  }

  StartAirdropCampaign(
    request: MsgStartAirdropCampaign
  ): Promise<MsgStartAirdropCampaignResponse> {
    const data = MsgStartAirdropCampaign.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Msg",
      "StartAirdropCampaign",
      data
    );
    return promise.then((data) =>
      MsgStartAirdropCampaignResponse.decode(new Reader(data))
    );
  }

  RemoveAirdropCampaign(
    request: MsgRemoveAirdropCampaign
  ): Promise<MsgRemoveAirdropCampaignResponse> {
    const data = MsgRemoveAirdropCampaign.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Msg",
      "RemoveAirdropCampaign",
      data
    );
    return promise.then((data) =>
      MsgRemoveAirdropCampaignResponse.decode(new Reader(data))
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

function toTimestamp(date: Date): Timestamp {
  const seconds = date.getTime() / 1_000;
  const nanos = (date.getTime() % 1_000) * 1_000_000;
  return { seconds, nanos };
}

function fromTimestamp(t: Timestamp): Date {
  let millis = t.seconds * 1_000;
  millis += t.nanos / 1_000_000;
  return new Date(millis);
}

function fromJsonTimestamp(o: any): Date {
  if (o instanceof Date) {
    return o;
  } else if (typeof o === "string") {
    return new Date(o);
  } else {
    return fromTimestamp(Timestamp.fromJSON(o));
  }
}

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
