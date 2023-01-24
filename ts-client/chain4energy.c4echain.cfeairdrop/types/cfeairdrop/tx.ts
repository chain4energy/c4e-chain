/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Duration } from "../google/protobuf/duration";
import { Timestamp } from "../google/protobuf/timestamp";
import {
  AirdropCloseAction,
  airdropCloseActionFromJSON,
  airdropCloseActionToJSON,
  AirdropEntry,
  MissionType,
  missionTypeFromJSON,
  missionTypeToJSON,
} from "./airdrop";

export const protobufPackage = "chain4energy.c4echain.cfeairdrop";

export interface MsgClaim {
  claimer: string;
  campaignId: number;
  missionId: number;
}

export interface MsgClaimResponse {
}

export interface MsgInitialClaim {
  claimer: string;
  campaignId: number;
  addressToClaim: string;
}

export interface MsgInitialClaimResponse {
}

export interface MsgCreateAirdropCampaign {
  owner: string;
  name: string;
  description: string;
  allowFeegrant: boolean;
  initialClaimFreeAmount: string;
  startTime: Date | undefined;
  endTime: Date | undefined;
  lockupPeriod: Duration | undefined;
  vestingPeriod: Duration | undefined;
}

export interface MsgCreateAirdropCampaignResponse {
}

export interface MsgAddMissionToAidropCampaign {
  owner: string;
  campaignId: number;
  name: string;
  description: string;
  missionType: MissionType;
  weight: string;
  claimStartDate: Date | undefined;
}

export interface MsgAddMissionToAidropCampaignResponse {
}

export interface MsgAddAirdropEntries {
  owner: string;
  campaignId: number;
  airdropEntries: AirdropEntry[];
}

export interface MsgAddAirdropEntriesResponse {
}

export interface MsgDeleteAirdropEntry {
  owner: string;
  campaignId: number;
  userAddress: string;
}

export interface MsgDeleteAirdropEntryResponse {
}

export interface MsgCloseAirdropCampaign {
  owner: string;
  campaignId: number;
  airdropCloseAction: AirdropCloseAction;
}

export interface MsgCloseAirdropCampaignResponse {
}

export interface MsgStartAirdropCampaign {
  owner: string;
  campaignId: number;
}

export interface MsgStartAirdropCampaignResponse {
}

export interface MsgEditAirdropCampaign {
  owner: string;
  campaignId: number;
  name: string;
  description: string;
  allowFeegrant: boolean;
  initialClaimFreeAmount: string;
  denom: string;
  startTime: Date | undefined;
  endTime: Date | undefined;
  lockupPeriod: Duration | undefined;
  vestingPeriod: Duration | undefined;
}

export interface MsgEditAirdropCampaignResponse {
}

function createBaseMsgClaim(): MsgClaim {
  return { claimer: "", campaignId: 0, missionId: 0 };
}

export const MsgClaim = {
  encode(message: MsgClaim, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.claimer !== "") {
      writer.uint32(10).string(message.claimer);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    if (message.missionId !== 0) {
      writer.uint32(24).uint64(message.missionId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaim {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaim();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.claimer = reader.string();
          break;
        case 2:
          message.campaignId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.missionId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClaim {
    return {
      claimer: isSet(object.claimer) ? String(object.claimer) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      missionId: isSet(object.missionId) ? Number(object.missionId) : 0,
    };
  },

  toJSON(message: MsgClaim): unknown {
    const obj: any = {};
    message.claimer !== undefined && (obj.claimer = message.claimer);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.missionId !== undefined && (obj.missionId = Math.round(message.missionId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaim>, I>>(object: I): MsgClaim {
    const message = createBaseMsgClaim();
    message.claimer = object.claimer ?? "";
    message.campaignId = object.campaignId ?? 0;
    message.missionId = object.missionId ?? 0;
    return message;
  },
};

function createBaseMsgClaimResponse(): MsgClaimResponse {
  return {};
}

export const MsgClaimResponse = {
  encode(_: MsgClaimResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimResponse();
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
    return {};
  },

  toJSON(_: MsgClaimResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimResponse>, I>>(_: I): MsgClaimResponse {
    const message = createBaseMsgClaimResponse();
    return message;
  },
};

function createBaseMsgInitialClaim(): MsgInitialClaim {
  return { claimer: "", campaignId: 0, addressToClaim: "" };
}

export const MsgInitialClaim = {
  encode(message: MsgInitialClaim, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.claimer !== "") {
      writer.uint32(10).string(message.claimer);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    if (message.addressToClaim !== "") {
      writer.uint32(34).string(message.addressToClaim);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgInitialClaim {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgInitialClaim();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.claimer = reader.string();
          break;
        case 2:
          message.campaignId = longToNumber(reader.uint64() as Long);
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
    return {
      claimer: isSet(object.claimer) ? String(object.claimer) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      addressToClaim: isSet(object.addressToClaim) ? String(object.addressToClaim) : "",
    };
  },

  toJSON(message: MsgInitialClaim): unknown {
    const obj: any = {};
    message.claimer !== undefined && (obj.claimer = message.claimer);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.addressToClaim !== undefined && (obj.addressToClaim = message.addressToClaim);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgInitialClaim>, I>>(object: I): MsgInitialClaim {
    const message = createBaseMsgInitialClaim();
    message.claimer = object.claimer ?? "";
    message.campaignId = object.campaignId ?? 0;
    message.addressToClaim = object.addressToClaim ?? "";
    return message;
  },
};

function createBaseMsgInitialClaimResponse(): MsgInitialClaimResponse {
  return {};
}

export const MsgInitialClaimResponse = {
  encode(_: MsgInitialClaimResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgInitialClaimResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgInitialClaimResponse();
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
    return {};
  },

  toJSON(_: MsgInitialClaimResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgInitialClaimResponse>, I>>(_: I): MsgInitialClaimResponse {
    const message = createBaseMsgInitialClaimResponse();
    return message;
  },
};

function createBaseMsgCreateAirdropCampaign(): MsgCreateAirdropCampaign {
  return {
    owner: "",
    name: "",
    description: "",
    allowFeegrant: false,
    initialClaimFreeAmount: "",
    startTime: undefined,
    endTime: undefined,
    lockupPeriod: undefined,
    vestingPeriod: undefined,
  };
}

export const MsgCreateAirdropCampaign = {
  encode(message: MsgCreateAirdropCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.allowFeegrant === true) {
      writer.uint32(32).bool(message.allowFeegrant);
    }
    if (message.initialClaimFreeAmount !== "") {
      writer.uint32(42).string(message.initialClaimFreeAmount);
    }
    if (message.startTime !== undefined) {
      Timestamp.encode(toTimestamp(message.startTime), writer.uint32(50).fork()).ldelim();
    }
    if (message.endTime !== undefined) {
      Timestamp.encode(toTimestamp(message.endTime), writer.uint32(58).fork()).ldelim();
    }
    if (message.lockupPeriod !== undefined) {
      Duration.encode(message.lockupPeriod, writer.uint32(66).fork()).ldelim();
    }
    if (message.vestingPeriod !== undefined) {
      Duration.encode(message.vestingPeriod, writer.uint32(74).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateAirdropCampaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateAirdropCampaign();
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
          message.allowFeegrant = reader.bool();
          break;
        case 5:
          message.initialClaimFreeAmount = reader.string();
          break;
        case 6:
          message.startTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 7:
          message.endTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 8:
          message.lockupPeriod = Duration.decode(reader, reader.uint32());
          break;
        case 9:
          message.vestingPeriod = Duration.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateAirdropCampaign {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      allowFeegrant: isSet(object.allowFeegrant) ? Boolean(object.allowFeegrant) : false,
      initialClaimFreeAmount: isSet(object.initialClaimFreeAmount) ? String(object.initialClaimFreeAmount) : "",
      startTime: isSet(object.startTime) ? fromJsonTimestamp(object.startTime) : undefined,
      endTime: isSet(object.endTime) ? fromJsonTimestamp(object.endTime) : undefined,
      lockupPeriod: isSet(object.lockupPeriod) ? Duration.fromJSON(object.lockupPeriod) : undefined,
      vestingPeriod: isSet(object.vestingPeriod) ? Duration.fromJSON(object.vestingPeriod) : undefined,
    };
  },

  toJSON(message: MsgCreateAirdropCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.allowFeegrant !== undefined && (obj.allowFeegrant = message.allowFeegrant);
    message.initialClaimFreeAmount !== undefined && (obj.initialClaimFreeAmount = message.initialClaimFreeAmount);
    message.startTime !== undefined && (obj.startTime = message.startTime.toISOString());
    message.endTime !== undefined && (obj.endTime = message.endTime.toISOString());
    message.lockupPeriod !== undefined
      && (obj.lockupPeriod = message.lockupPeriod ? Duration.toJSON(message.lockupPeriod) : undefined);
    message.vestingPeriod !== undefined
      && (obj.vestingPeriod = message.vestingPeriod ? Duration.toJSON(message.vestingPeriod) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateAirdropCampaign>, I>>(object: I): MsgCreateAirdropCampaign {
    const message = createBaseMsgCreateAirdropCampaign();
    message.owner = object.owner ?? "";
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.allowFeegrant = object.allowFeegrant ?? false;
    message.initialClaimFreeAmount = object.initialClaimFreeAmount ?? "";
    message.startTime = object.startTime ?? undefined;
    message.endTime = object.endTime ?? undefined;
    message.lockupPeriod = (object.lockupPeriod !== undefined && object.lockupPeriod !== null)
      ? Duration.fromPartial(object.lockupPeriod)
      : undefined;
    message.vestingPeriod = (object.vestingPeriod !== undefined && object.vestingPeriod !== null)
      ? Duration.fromPartial(object.vestingPeriod)
      : undefined;
    return message;
  },
};

function createBaseMsgCreateAirdropCampaignResponse(): MsgCreateAirdropCampaignResponse {
  return {};
}

export const MsgCreateAirdropCampaignResponse = {
  encode(_: MsgCreateAirdropCampaignResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateAirdropCampaignResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateAirdropCampaignResponse();
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
    return {};
  },

  toJSON(_: MsgCreateAirdropCampaignResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateAirdropCampaignResponse>, I>>(
    _: I,
  ): MsgCreateAirdropCampaignResponse {
    const message = createBaseMsgCreateAirdropCampaignResponse();
    return message;
  },
};

function createBaseMsgAddMissionToAidropCampaign(): MsgAddMissionToAidropCampaign {
  return { owner: "", campaignId: 0, name: "", description: "", missionType: 0, weight: "", claimStartDate: undefined };
}

export const MsgAddMissionToAidropCampaign = {
  encode(message: MsgAddMissionToAidropCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
    if (message.claimStartDate !== undefined) {
      Timestamp.encode(toTimestamp(message.claimStartDate), writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddMissionToAidropCampaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddMissionToAidropCampaign();
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
          message.claimStartDate = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddMissionToAidropCampaign {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      missionType: isSet(object.missionType) ? missionTypeFromJSON(object.missionType) : 0,
      weight: isSet(object.weight) ? String(object.weight) : "",
      claimStartDate: isSet(object.claimStartDate) ? fromJsonTimestamp(object.claimStartDate) : undefined,
    };
  },

  toJSON(message: MsgAddMissionToAidropCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.missionType !== undefined && (obj.missionType = missionTypeToJSON(message.missionType));
    message.weight !== undefined && (obj.weight = message.weight);
    message.claimStartDate !== undefined && (obj.claimStartDate = message.claimStartDate.toISOString());
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddMissionToAidropCampaign>, I>>(
    object: I,
  ): MsgAddMissionToAidropCampaign {
    const message = createBaseMsgAddMissionToAidropCampaign();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? 0;
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.missionType = object.missionType ?? 0;
    message.weight = object.weight ?? "";
    message.claimStartDate = object.claimStartDate ?? undefined;
    return message;
  },
};

function createBaseMsgAddMissionToAidropCampaignResponse(): MsgAddMissionToAidropCampaignResponse {
  return {};
}

export const MsgAddMissionToAidropCampaignResponse = {
  encode(_: MsgAddMissionToAidropCampaignResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddMissionToAidropCampaignResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddMissionToAidropCampaignResponse();
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
    return {};
  },

  toJSON(_: MsgAddMissionToAidropCampaignResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddMissionToAidropCampaignResponse>, I>>(
    _: I,
  ): MsgAddMissionToAidropCampaignResponse {
    const message = createBaseMsgAddMissionToAidropCampaignResponse();
    return message;
  },
};

function createBaseMsgAddAirdropEntries(): MsgAddAirdropEntries {
  return { owner: "", campaignId: 0, airdropEntries: [] };
}

export const MsgAddAirdropEntries = {
  encode(message: MsgAddAirdropEntries, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    for (const v of message.airdropEntries) {
      AirdropEntry.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddAirdropEntries {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddAirdropEntries();
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
          message.airdropEntries.push(AirdropEntry.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddAirdropEntries {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      airdropEntries: Array.isArray(object?.airdropEntries)
        ? object.airdropEntries.map((e: any) => AirdropEntry.fromJSON(e))
        : [],
    };
  },

  toJSON(message: MsgAddAirdropEntries): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    if (message.airdropEntries) {
      obj.airdropEntries = message.airdropEntries.map((e) => e ? AirdropEntry.toJSON(e) : undefined);
    } else {
      obj.airdropEntries = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddAirdropEntries>, I>>(object: I): MsgAddAirdropEntries {
    const message = createBaseMsgAddAirdropEntries();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? 0;
    message.airdropEntries = object.airdropEntries?.map((e) => AirdropEntry.fromPartial(e)) || [];
    return message;
  },
};

function createBaseMsgAddAirdropEntriesResponse(): MsgAddAirdropEntriesResponse {
  return {};
}

export const MsgAddAirdropEntriesResponse = {
  encode(_: MsgAddAirdropEntriesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddAirdropEntriesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddAirdropEntriesResponse();
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
    return {};
  },

  toJSON(_: MsgAddAirdropEntriesResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddAirdropEntriesResponse>, I>>(_: I): MsgAddAirdropEntriesResponse {
    const message = createBaseMsgAddAirdropEntriesResponse();
    return message;
  },
};

function createBaseMsgDeleteAirdropEntry(): MsgDeleteAirdropEntry {
  return { owner: "", campaignId: 0, userAddress: "" };
}

export const MsgDeleteAirdropEntry = {
  encode(message: MsgDeleteAirdropEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteAirdropEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteAirdropEntry();
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

  fromJSON(object: any): MsgDeleteAirdropEntry {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      userAddress: isSet(object.userAddress) ? String(object.userAddress) : "",
    };
  },

  toJSON(message: MsgDeleteAirdropEntry): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.userAddress !== undefined && (obj.userAddress = message.userAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteAirdropEntry>, I>>(object: I): MsgDeleteAirdropEntry {
    const message = createBaseMsgDeleteAirdropEntry();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? 0;
    message.userAddress = object.userAddress ?? "";
    return message;
  },
};

function createBaseMsgDeleteAirdropEntryResponse(): MsgDeleteAirdropEntryResponse {
  return {};
}

export const MsgDeleteAirdropEntryResponse = {
  encode(_: MsgDeleteAirdropEntryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteAirdropEntryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteAirdropEntryResponse();
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
    return {};
  },

  toJSON(_: MsgDeleteAirdropEntryResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteAirdropEntryResponse>, I>>(_: I): MsgDeleteAirdropEntryResponse {
    const message = createBaseMsgDeleteAirdropEntryResponse();
    return message;
  },
};

function createBaseMsgCloseAirdropCampaign(): MsgCloseAirdropCampaign {
  return { owner: "", campaignId: 0, airdropCloseAction: 0 };
}

export const MsgCloseAirdropCampaign = {
  encode(message: MsgCloseAirdropCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    if (message.airdropCloseAction !== 0) {
      writer.uint32(24).int32(message.airdropCloseAction);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCloseAirdropCampaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCloseAirdropCampaign();
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
          message.airdropCloseAction = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCloseAirdropCampaign {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      airdropCloseAction: isSet(object.airdropCloseAction) ? airdropCloseActionFromJSON(object.airdropCloseAction) : 0,
    };
  },

  toJSON(message: MsgCloseAirdropCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.airdropCloseAction !== undefined
      && (obj.airdropCloseAction = airdropCloseActionToJSON(message.airdropCloseAction));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCloseAirdropCampaign>, I>>(object: I): MsgCloseAirdropCampaign {
    const message = createBaseMsgCloseAirdropCampaign();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? 0;
    message.airdropCloseAction = object.airdropCloseAction ?? 0;
    return message;
  },
};

function createBaseMsgCloseAirdropCampaignResponse(): MsgCloseAirdropCampaignResponse {
  return {};
}

export const MsgCloseAirdropCampaignResponse = {
  encode(_: MsgCloseAirdropCampaignResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCloseAirdropCampaignResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCloseAirdropCampaignResponse();
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
    return {};
  },

  toJSON(_: MsgCloseAirdropCampaignResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCloseAirdropCampaignResponse>, I>>(_: I): MsgCloseAirdropCampaignResponse {
    const message = createBaseMsgCloseAirdropCampaignResponse();
    return message;
  },
};

function createBaseMsgStartAirdropCampaign(): MsgStartAirdropCampaign {
  return { owner: "", campaignId: 0 };
}

export const MsgStartAirdropCampaign = {
  encode(message: MsgStartAirdropCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgStartAirdropCampaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgStartAirdropCampaign();
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
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
    };
  },

  toJSON(message: MsgStartAirdropCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgStartAirdropCampaign>, I>>(object: I): MsgStartAirdropCampaign {
    const message = createBaseMsgStartAirdropCampaign();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? 0;
    return message;
  },
};

function createBaseMsgStartAirdropCampaignResponse(): MsgStartAirdropCampaignResponse {
  return {};
}

export const MsgStartAirdropCampaignResponse = {
  encode(_: MsgStartAirdropCampaignResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgStartAirdropCampaignResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgStartAirdropCampaignResponse();
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
    return {};
  },

  toJSON(_: MsgStartAirdropCampaignResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgStartAirdropCampaignResponse>, I>>(_: I): MsgStartAirdropCampaignResponse {
    const message = createBaseMsgStartAirdropCampaignResponse();
    return message;
  },
};

function createBaseMsgEditAirdropCampaign(): MsgEditAirdropCampaign {
  return {
    owner: "",
    campaignId: 0,
    name: "",
    description: "",
    allowFeegrant: false,
    initialClaimFreeAmount: "",
    denom: "",
    startTime: undefined,
    endTime: undefined,
    lockupPeriod: undefined,
    vestingPeriod: undefined,
  };
}

export const MsgEditAirdropCampaign = {
  encode(message: MsgEditAirdropCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
    if (message.allowFeegrant === true) {
      writer.uint32(40).bool(message.allowFeegrant);
    }
    if (message.initialClaimFreeAmount !== "") {
      writer.uint32(50).string(message.initialClaimFreeAmount);
    }
    if (message.denom !== "") {
      writer.uint32(58).string(message.denom);
    }
    if (message.startTime !== undefined) {
      Timestamp.encode(toTimestamp(message.startTime), writer.uint32(66).fork()).ldelim();
    }
    if (message.endTime !== undefined) {
      Timestamp.encode(toTimestamp(message.endTime), writer.uint32(74).fork()).ldelim();
    }
    if (message.lockupPeriod !== undefined) {
      Duration.encode(message.lockupPeriod, writer.uint32(82).fork()).ldelim();
    }
    if (message.vestingPeriod !== undefined) {
      Duration.encode(message.vestingPeriod, writer.uint32(90).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgEditAirdropCampaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgEditAirdropCampaign();
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
          message.allowFeegrant = reader.bool();
          break;
        case 6:
          message.initialClaimFreeAmount = reader.string();
          break;
        case 7:
          message.denom = reader.string();
          break;
        case 8:
          message.startTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 9:
          message.endTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 10:
          message.lockupPeriod = Duration.decode(reader, reader.uint32());
          break;
        case 11:
          message.vestingPeriod = Duration.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEditAirdropCampaign {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      allowFeegrant: isSet(object.allowFeegrant) ? Boolean(object.allowFeegrant) : false,
      initialClaimFreeAmount: isSet(object.initialClaimFreeAmount) ? String(object.initialClaimFreeAmount) : "",
      denom: isSet(object.denom) ? String(object.denom) : "",
      startTime: isSet(object.startTime) ? fromJsonTimestamp(object.startTime) : undefined,
      endTime: isSet(object.endTime) ? fromJsonTimestamp(object.endTime) : undefined,
      lockupPeriod: isSet(object.lockupPeriod) ? Duration.fromJSON(object.lockupPeriod) : undefined,
      vestingPeriod: isSet(object.vestingPeriod) ? Duration.fromJSON(object.vestingPeriod) : undefined,
    };
  },

  toJSON(message: MsgEditAirdropCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.allowFeegrant !== undefined && (obj.allowFeegrant = message.allowFeegrant);
    message.initialClaimFreeAmount !== undefined && (obj.initialClaimFreeAmount = message.initialClaimFreeAmount);
    message.denom !== undefined && (obj.denom = message.denom);
    message.startTime !== undefined && (obj.startTime = message.startTime.toISOString());
    message.endTime !== undefined && (obj.endTime = message.endTime.toISOString());
    message.lockupPeriod !== undefined
      && (obj.lockupPeriod = message.lockupPeriod ? Duration.toJSON(message.lockupPeriod) : undefined);
    message.vestingPeriod !== undefined
      && (obj.vestingPeriod = message.vestingPeriod ? Duration.toJSON(message.vestingPeriod) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgEditAirdropCampaign>, I>>(object: I): MsgEditAirdropCampaign {
    const message = createBaseMsgEditAirdropCampaign();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? 0;
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.allowFeegrant = object.allowFeegrant ?? false;
    message.initialClaimFreeAmount = object.initialClaimFreeAmount ?? "";
    message.denom = object.denom ?? "";
    message.startTime = object.startTime ?? undefined;
    message.endTime = object.endTime ?? undefined;
    message.lockupPeriod = (object.lockupPeriod !== undefined && object.lockupPeriod !== null)
      ? Duration.fromPartial(object.lockupPeriod)
      : undefined;
    message.vestingPeriod = (object.vestingPeriod !== undefined && object.vestingPeriod !== null)
      ? Duration.fromPartial(object.vestingPeriod)
      : undefined;
    return message;
  },
};

function createBaseMsgEditAirdropCampaignResponse(): MsgEditAirdropCampaignResponse {
  return {};
}

export const MsgEditAirdropCampaignResponse = {
  encode(_: MsgEditAirdropCampaignResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgEditAirdropCampaignResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgEditAirdropCampaignResponse();
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
    return {};
  },

  toJSON(_: MsgEditAirdropCampaignResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgEditAirdropCampaignResponse>, I>>(_: I): MsgEditAirdropCampaignResponse {
    const message = createBaseMsgEditAirdropCampaignResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Claim(request: MsgClaim): Promise<MsgClaimResponse>;
  InitialClaim(request: MsgInitialClaim): Promise<MsgInitialClaimResponse>;
  CreateAirdropCampaign(request: MsgCreateAirdropCampaign): Promise<MsgCreateAirdropCampaignResponse>;
  EditAirdropCampaign(request: MsgEditAirdropCampaign): Promise<MsgEditAirdropCampaignResponse>;
  AddMissionToAidropCampaign(request: MsgAddMissionToAidropCampaign): Promise<MsgAddMissionToAidropCampaignResponse>;
  AddAirdropEntries(request: MsgAddAirdropEntries): Promise<MsgAddAirdropEntriesResponse>;
  DeleteAirdropEntry(request: MsgDeleteAirdropEntry): Promise<MsgDeleteAirdropEntryResponse>;
  CloseAirdropCampaign(request: MsgCloseAirdropCampaign): Promise<MsgCloseAirdropCampaignResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  StartAirdropCampaign(request: MsgStartAirdropCampaign): Promise<MsgStartAirdropCampaignResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Claim = this.Claim.bind(this);
    this.InitialClaim = this.InitialClaim.bind(this);
    this.CreateAirdropCampaign = this.CreateAirdropCampaign.bind(this);
    this.EditAirdropCampaign = this.EditAirdropCampaign.bind(this);
    this.AddMissionToAidropCampaign = this.AddMissionToAidropCampaign.bind(this);
    this.AddAirdropEntries = this.AddAirdropEntries.bind(this);
    this.DeleteAirdropEntry = this.DeleteAirdropEntry.bind(this);
    this.CloseAirdropCampaign = this.CloseAirdropCampaign.bind(this);
    this.StartAirdropCampaign = this.StartAirdropCampaign.bind(this);
  }
  Claim(request: MsgClaim): Promise<MsgClaimResponse> {
    const data = MsgClaim.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Msg", "Claim", data);
    return promise.then((data) => MsgClaimResponse.decode(new _m0.Reader(data)));
  }

  InitialClaim(request: MsgInitialClaim): Promise<MsgInitialClaimResponse> {
    const data = MsgInitialClaim.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Msg", "InitialClaim", data);
    return promise.then((data) => MsgInitialClaimResponse.decode(new _m0.Reader(data)));
  }

  CreateAirdropCampaign(request: MsgCreateAirdropCampaign): Promise<MsgCreateAirdropCampaignResponse> {
    const data = MsgCreateAirdropCampaign.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Msg", "CreateAirdropCampaign", data);
    return promise.then((data) => MsgCreateAirdropCampaignResponse.decode(new _m0.Reader(data)));
  }

  EditAirdropCampaign(request: MsgEditAirdropCampaign): Promise<MsgEditAirdropCampaignResponse> {
    const data = MsgEditAirdropCampaign.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Msg", "EditAirdropCampaign", data);
    return promise.then((data) => MsgEditAirdropCampaignResponse.decode(new _m0.Reader(data)));
  }

  AddMissionToAidropCampaign(request: MsgAddMissionToAidropCampaign): Promise<MsgAddMissionToAidropCampaignResponse> {
    const data = MsgAddMissionToAidropCampaign.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Msg", "AddMissionToAidropCampaign", data);
    return promise.then((data) => MsgAddMissionToAidropCampaignResponse.decode(new _m0.Reader(data)));
  }

  AddAirdropEntries(request: MsgAddAirdropEntries): Promise<MsgAddAirdropEntriesResponse> {
    const data = MsgAddAirdropEntries.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Msg", "AddAirdropEntries", data);
    return promise.then((data) => MsgAddAirdropEntriesResponse.decode(new _m0.Reader(data)));
  }

  DeleteAirdropEntry(request: MsgDeleteAirdropEntry): Promise<MsgDeleteAirdropEntryResponse> {
    const data = MsgDeleteAirdropEntry.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Msg", "DeleteAirdropEntry", data);
    return promise.then((data) => MsgDeleteAirdropEntryResponse.decode(new _m0.Reader(data)));
  }

  CloseAirdropCampaign(request: MsgCloseAirdropCampaign): Promise<MsgCloseAirdropCampaignResponse> {
    const data = MsgCloseAirdropCampaign.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Msg", "CloseAirdropCampaign", data);
    return promise.then((data) => MsgCloseAirdropCampaignResponse.decode(new _m0.Reader(data)));
  }

  StartAirdropCampaign(request: MsgStartAirdropCampaign): Promise<MsgStartAirdropCampaignResponse> {
    const data = MsgStartAirdropCampaign.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeairdrop.Msg", "StartAirdropCampaign", data);
    return promise.then((data) => MsgStartAirdropCampaignResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

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

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
