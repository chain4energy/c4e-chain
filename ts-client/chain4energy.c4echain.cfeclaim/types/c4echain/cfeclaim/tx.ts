/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";
import { Duration } from "../../google/protobuf/duration";
import { Timestamp } from "../../google/protobuf/timestamp";
import { CampaignType, campaignTypeFromJSON, campaignTypeToJSON } from "./campaign";
import { ClaimRecordEntry } from "./claim_record";
import { MissionType, missionTypeFromJSON, missionTypeToJSON } from "./mission";

export const protobufPackage = "chain4energy.c4echain.cfeclaim";

export interface MsgClaim {
  claimer: string;
  campaignId: number;
  missionId: number;
}

export interface MsgClaimResponse {
  amount: Coin[];
}

export interface MsgInitialClaim {
  claimer: string;
  campaignId: number;
  destinationAddress: string;
}

export interface MsgInitialClaimResponse {
  amount: Coin[];
}

export interface MsgCreateCampaign {
  owner: string;
  name: string;
  description: string;
  campaignType: CampaignType;
  removableClaimRecords: boolean;
  feegrantAmount: string;
  initialClaimFreeAmount: string;
  free: string;
  startTime: Date | undefined;
  endTime: Date | undefined;
  lockupPeriod: Duration | undefined;
  vestingPeriod: Duration | undefined;
  vestingPoolName: string;
}

export interface MsgCreateCampaignResponse {
  campaignId: number;
}

export interface MsgAddMission {
  owner: string;
  campaignId: number;
  name: string;
  description: string;
  missionType: MissionType;
  weight: string;
  claimStartDate: Date | undefined;
}

export interface MsgAddMissionResponse {
  missionId: number;
}

export interface MsgAddClaimRecords {
  owner: string;
  campaignId: number;
  claimRecordEntries: ClaimRecordEntry[];
}

export interface MsgAddClaimRecordsResponse {
}

export interface MsgDeleteClaimRecord {
  owner: string;
  campaignId: number;
  userAddress: string;
}

export interface MsgDeleteClaimRecordResponse {
}

export interface MsgCloseCampaign {
  owner: string;
  campaignId: number;
}

export interface MsgCloseCampaignResponse {
}

export interface MsgEnableCampaign {
  owner: string;
  campaignId: number;
  startTime: Date | undefined;
  endTime: Date | undefined;
}

export interface MsgEnableCampaignResponse {
}

export interface MsgRemoveCampaign {
  owner: string;
  campaignId: number;
}

export interface MsgRemoveCampaignResponse {
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
  return { amount: [] };
}

export const MsgClaimResponse = {
  encode(message: MsgClaimResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgClaimResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgClaimResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClaimResponse {
    return { amount: Array.isArray(object?.amount) ? object.amount.map((e: any) => Coin.fromJSON(e)) : [] };
  },

  toJSON(message: MsgClaimResponse): unknown {
    const obj: any = {};
    if (message.amount) {
      obj.amount = message.amount.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.amount = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgClaimResponse>, I>>(object: I): MsgClaimResponse {
    const message = createBaseMsgClaimResponse();
    message.amount = object.amount?.map((e) => Coin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseMsgInitialClaim(): MsgInitialClaim {
  return { claimer: "", campaignId: 0, destinationAddress: "" };
}

export const MsgInitialClaim = {
  encode(message: MsgInitialClaim, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.claimer !== "") {
      writer.uint32(10).string(message.claimer);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    if (message.destinationAddress !== "") {
      writer.uint32(26).string(message.destinationAddress);
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
        case 3:
          message.destinationAddress = reader.string();
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
      destinationAddress: isSet(object.destinationAddress) ? String(object.destinationAddress) : "",
    };
  },

  toJSON(message: MsgInitialClaim): unknown {
    const obj: any = {};
    message.claimer !== undefined && (obj.claimer = message.claimer);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.destinationAddress !== undefined && (obj.destinationAddress = message.destinationAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgInitialClaim>, I>>(object: I): MsgInitialClaim {
    const message = createBaseMsgInitialClaim();
    message.claimer = object.claimer ?? "";
    message.campaignId = object.campaignId ?? 0;
    message.destinationAddress = object.destinationAddress ?? "";
    return message;
  },
};

function createBaseMsgInitialClaimResponse(): MsgInitialClaimResponse {
  return { amount: [] };
}

export const MsgInitialClaimResponse = {
  encode(message: MsgInitialClaimResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.amount) {
      Coin.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgInitialClaimResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgInitialClaimResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.amount.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgInitialClaimResponse {
    return { amount: Array.isArray(object?.amount) ? object.amount.map((e: any) => Coin.fromJSON(e)) : [] };
  },

  toJSON(message: MsgInitialClaimResponse): unknown {
    const obj: any = {};
    if (message.amount) {
      obj.amount = message.amount.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.amount = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgInitialClaimResponse>, I>>(object: I): MsgInitialClaimResponse {
    const message = createBaseMsgInitialClaimResponse();
    message.amount = object.amount?.map((e) => Coin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseMsgCreateCampaign(): MsgCreateCampaign {
  return {
    owner: "",
    name: "",
    description: "",
    campaignType: 0,
    removableClaimRecords: false,
    feegrantAmount: "",
    initialClaimFreeAmount: "",
    free: "",
    startTime: undefined,
    endTime: undefined,
    lockupPeriod: undefined,
    vestingPeriod: undefined,
    vestingPoolName: "",
  };
}

export const MsgCreateCampaign = {
  encode(message: MsgCreateCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(26).string(message.description);
    }
    if (message.campaignType !== 0) {
      writer.uint32(32).int32(message.campaignType);
    }
    if (message.removableClaimRecords === true) {
      writer.uint32(40).bool(message.removableClaimRecords);
    }
    if (message.feegrantAmount !== "") {
      writer.uint32(50).string(message.feegrantAmount);
    }
    if (message.initialClaimFreeAmount !== "") {
      writer.uint32(58).string(message.initialClaimFreeAmount);
    }
    if (message.free !== "") {
      writer.uint32(66).string(message.free);
    }
    if (message.startTime !== undefined) {
      Timestamp.encode(toTimestamp(message.startTime), writer.uint32(74).fork()).ldelim();
    }
    if (message.endTime !== undefined) {
      Timestamp.encode(toTimestamp(message.endTime), writer.uint32(82).fork()).ldelim();
    }
    if (message.lockupPeriod !== undefined) {
      Duration.encode(message.lockupPeriod, writer.uint32(90).fork()).ldelim();
    }
    if (message.vestingPeriod !== undefined) {
      Duration.encode(message.vestingPeriod, writer.uint32(98).fork()).ldelim();
    }
    if (message.vestingPoolName !== "") {
      writer.uint32(106).string(message.vestingPoolName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateCampaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateCampaign();
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
          message.campaignType = reader.int32() as any;
          break;
        case 5:
          message.removableClaimRecords = reader.bool();
          break;
        case 6:
          message.feegrantAmount = reader.string();
          break;
        case 7:
          message.initialClaimFreeAmount = reader.string();
          break;
        case 8:
          message.free = reader.string();
          break;
        case 9:
          message.startTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 10:
          message.endTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 11:
          message.lockupPeriod = Duration.decode(reader, reader.uint32());
          break;
        case 12:
          message.vestingPeriod = Duration.decode(reader, reader.uint32());
          break;
        case 13:
          message.vestingPoolName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateCampaign {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      campaignType: isSet(object.campaignType) ? campaignTypeFromJSON(object.campaignType) : 0,
      removableClaimRecords: isSet(object.removableClaimRecords) ? Boolean(object.removableClaimRecords) : false,
      feegrantAmount: isSet(object.feegrantAmount) ? String(object.feegrantAmount) : "",
      initialClaimFreeAmount: isSet(object.initialClaimFreeAmount) ? String(object.initialClaimFreeAmount) : "",
      free: isSet(object.free) ? String(object.free) : "",
      startTime: isSet(object.startTime) ? fromJsonTimestamp(object.startTime) : undefined,
      endTime: isSet(object.endTime) ? fromJsonTimestamp(object.endTime) : undefined,
      lockupPeriod: isSet(object.lockupPeriod) ? Duration.fromJSON(object.lockupPeriod) : undefined,
      vestingPeriod: isSet(object.vestingPeriod) ? Duration.fromJSON(object.vestingPeriod) : undefined,
      vestingPoolName: isSet(object.vestingPoolName) ? String(object.vestingPoolName) : "",
    };
  },

  toJSON(message: MsgCreateCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.campaignType !== undefined && (obj.campaignType = campaignTypeToJSON(message.campaignType));
    message.removableClaimRecords !== undefined && (obj.removableClaimRecords = message.removableClaimRecords);
    message.feegrantAmount !== undefined && (obj.feegrantAmount = message.feegrantAmount);
    message.initialClaimFreeAmount !== undefined && (obj.initialClaimFreeAmount = message.initialClaimFreeAmount);
    message.free !== undefined && (obj.free = message.free);
    message.startTime !== undefined && (obj.startTime = message.startTime.toISOString());
    message.endTime !== undefined && (obj.endTime = message.endTime.toISOString());
    message.lockupPeriod !== undefined
      && (obj.lockupPeriod = message.lockupPeriod ? Duration.toJSON(message.lockupPeriod) : undefined);
    message.vestingPeriod !== undefined
      && (obj.vestingPeriod = message.vestingPeriod ? Duration.toJSON(message.vestingPeriod) : undefined);
    message.vestingPoolName !== undefined && (obj.vestingPoolName = message.vestingPoolName);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateCampaign>, I>>(object: I): MsgCreateCampaign {
    const message = createBaseMsgCreateCampaign();
    message.owner = object.owner ?? "";
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.campaignType = object.campaignType ?? 0;
    message.removableClaimRecords = object.removableClaimRecords ?? false;
    message.feegrantAmount = object.feegrantAmount ?? "";
    message.initialClaimFreeAmount = object.initialClaimFreeAmount ?? "";
    message.free = object.free ?? "";
    message.startTime = object.startTime ?? undefined;
    message.endTime = object.endTime ?? undefined;
    message.lockupPeriod = (object.lockupPeriod !== undefined && object.lockupPeriod !== null)
      ? Duration.fromPartial(object.lockupPeriod)
      : undefined;
    message.vestingPeriod = (object.vestingPeriod !== undefined && object.vestingPeriod !== null)
      ? Duration.fromPartial(object.vestingPeriod)
      : undefined;
    message.vestingPoolName = object.vestingPoolName ?? "";
    return message;
  },
};

function createBaseMsgCreateCampaignResponse(): MsgCreateCampaignResponse {
  return { campaignId: 0 };
}

export const MsgCreateCampaignResponse = {
  encode(message: MsgCreateCampaignResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.campaignId !== 0) {
      writer.uint32(8).uint64(message.campaignId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateCampaignResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateCampaignResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaignId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateCampaignResponse {
    return { campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0 };
  },

  toJSON(message: MsgCreateCampaignResponse): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateCampaignResponse>, I>>(object: I): MsgCreateCampaignResponse {
    const message = createBaseMsgCreateCampaignResponse();
    message.campaignId = object.campaignId ?? 0;
    return message;
  },
};

function createBaseMsgAddMission(): MsgAddMission {
  return { owner: "", campaignId: 0, name: "", description: "", missionType: 0, weight: "", claimStartDate: undefined };
}

export const MsgAddMission = {
  encode(message: MsgAddMission, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddMission {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddMission();
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

  fromJSON(object: any): MsgAddMission {
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

  toJSON(message: MsgAddMission): unknown {
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

  fromPartial<I extends Exact<DeepPartial<MsgAddMission>, I>>(object: I): MsgAddMission {
    const message = createBaseMsgAddMission();
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

function createBaseMsgAddMissionResponse(): MsgAddMissionResponse {
  return { missionId: 0 };
}

export const MsgAddMissionResponse = {
  encode(message: MsgAddMissionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.missionId !== 0) {
      writer.uint32(8).uint64(message.missionId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddMissionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddMissionResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.missionId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddMissionResponse {
    return { missionId: isSet(object.missionId) ? Number(object.missionId) : 0 };
  },

  toJSON(message: MsgAddMissionResponse): unknown {
    const obj: any = {};
    message.missionId !== undefined && (obj.missionId = Math.round(message.missionId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddMissionResponse>, I>>(object: I): MsgAddMissionResponse {
    const message = createBaseMsgAddMissionResponse();
    message.missionId = object.missionId ?? 0;
    return message;
  },
};

function createBaseMsgAddClaimRecords(): MsgAddClaimRecords {
  return { owner: "", campaignId: 0, claimRecordEntries: [] };
}

export const MsgAddClaimRecords = {
  encode(message: MsgAddClaimRecords, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    for (const v of message.claimRecordEntries) {
      ClaimRecordEntry.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddClaimRecords {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddClaimRecords();
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
          message.claimRecordEntries.push(ClaimRecordEntry.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddClaimRecords {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      claimRecordEntries: Array.isArray(object?.claimRecordEntries)
        ? object.claimRecordEntries.map((e: any) => ClaimRecordEntry.fromJSON(e))
        : [],
    };
  },

  toJSON(message: MsgAddClaimRecords): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    if (message.claimRecordEntries) {
      obj.claimRecordEntries = message.claimRecordEntries.map((e) => e ? ClaimRecordEntry.toJSON(e) : undefined);
    } else {
      obj.claimRecordEntries = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddClaimRecords>, I>>(object: I): MsgAddClaimRecords {
    const message = createBaseMsgAddClaimRecords();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? 0;
    message.claimRecordEntries = object.claimRecordEntries?.map((e) => ClaimRecordEntry.fromPartial(e)) || [];
    return message;
  },
};

function createBaseMsgAddClaimRecordsResponse(): MsgAddClaimRecordsResponse {
  return {};
}

export const MsgAddClaimRecordsResponse = {
  encode(_: MsgAddClaimRecordsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddClaimRecordsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddClaimRecordsResponse();
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
    return {};
  },

  toJSON(_: MsgAddClaimRecordsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddClaimRecordsResponse>, I>>(_: I): MsgAddClaimRecordsResponse {
    const message = createBaseMsgAddClaimRecordsResponse();
    return message;
  },
};

function createBaseMsgDeleteClaimRecord(): MsgDeleteClaimRecord {
  return { owner: "", campaignId: 0, userAddress: "" };
}

export const MsgDeleteClaimRecord = {
  encode(message: MsgDeleteClaimRecord, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteClaimRecord {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteClaimRecord();
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
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      userAddress: isSet(object.userAddress) ? String(object.userAddress) : "",
    };
  },

  toJSON(message: MsgDeleteClaimRecord): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.userAddress !== undefined && (obj.userAddress = message.userAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteClaimRecord>, I>>(object: I): MsgDeleteClaimRecord {
    const message = createBaseMsgDeleteClaimRecord();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? 0;
    message.userAddress = object.userAddress ?? "";
    return message;
  },
};

function createBaseMsgDeleteClaimRecordResponse(): MsgDeleteClaimRecordResponse {
  return {};
}

export const MsgDeleteClaimRecordResponse = {
  encode(_: MsgDeleteClaimRecordResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgDeleteClaimRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgDeleteClaimRecordResponse();
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
    return {};
  },

  toJSON(_: MsgDeleteClaimRecordResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgDeleteClaimRecordResponse>, I>>(_: I): MsgDeleteClaimRecordResponse {
    const message = createBaseMsgDeleteClaimRecordResponse();
    return message;
  },
};

function createBaseMsgCloseCampaign(): MsgCloseCampaign {
  return { owner: "", campaignId: 0 };
}

export const MsgCloseCampaign = {
  encode(message: MsgCloseCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCloseCampaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCloseCampaign();
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

  fromJSON(object: any): MsgCloseCampaign {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
    };
  },

  toJSON(message: MsgCloseCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCloseCampaign>, I>>(object: I): MsgCloseCampaign {
    const message = createBaseMsgCloseCampaign();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? 0;
    return message;
  },
};

function createBaseMsgCloseCampaignResponse(): MsgCloseCampaignResponse {
  return {};
}

export const MsgCloseCampaignResponse = {
  encode(_: MsgCloseCampaignResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCloseCampaignResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCloseCampaignResponse();
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

  fromJSON(_: any): MsgCloseCampaignResponse {
    return {};
  },

  toJSON(_: MsgCloseCampaignResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCloseCampaignResponse>, I>>(_: I): MsgCloseCampaignResponse {
    const message = createBaseMsgCloseCampaignResponse();
    return message;
  },
};

function createBaseMsgEnableCampaign(): MsgEnableCampaign {
  return { owner: "", campaignId: 0, startTime: undefined, endTime: undefined };
}

export const MsgEnableCampaign = {
  encode(message: MsgEnableCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    if (message.startTime !== undefined) {
      Timestamp.encode(toTimestamp(message.startTime), writer.uint32(26).fork()).ldelim();
    }
    if (message.endTime !== undefined) {
      Timestamp.encode(toTimestamp(message.endTime), writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgEnableCampaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgEnableCampaign();
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
          message.startTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 4:
          message.endTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEnableCampaign {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      startTime: isSet(object.startTime) ? fromJsonTimestamp(object.startTime) : undefined,
      endTime: isSet(object.endTime) ? fromJsonTimestamp(object.endTime) : undefined,
    };
  },

  toJSON(message: MsgEnableCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.startTime !== undefined && (obj.startTime = message.startTime.toISOString());
    message.endTime !== undefined && (obj.endTime = message.endTime.toISOString());
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgEnableCampaign>, I>>(object: I): MsgEnableCampaign {
    const message = createBaseMsgEnableCampaign();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? 0;
    message.startTime = object.startTime ?? undefined;
    message.endTime = object.endTime ?? undefined;
    return message;
  },
};

function createBaseMsgEnableCampaignResponse(): MsgEnableCampaignResponse {
  return {};
}

export const MsgEnableCampaignResponse = {
  encode(_: MsgEnableCampaignResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgEnableCampaignResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgEnableCampaignResponse();
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

  fromJSON(_: any): MsgEnableCampaignResponse {
    return {};
  },

  toJSON(_: MsgEnableCampaignResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgEnableCampaignResponse>, I>>(_: I): MsgEnableCampaignResponse {
    const message = createBaseMsgEnableCampaignResponse();
    return message;
  },
};

function createBaseMsgRemoveCampaign(): MsgRemoveCampaign {
  return { owner: "", campaignId: 0 };
}

export const MsgRemoveCampaign = {
  encode(message: MsgRemoveCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRemoveCampaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRemoveCampaign();
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

  fromJSON(object: any): MsgRemoveCampaign {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
    };
  },

  toJSON(message: MsgRemoveCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRemoveCampaign>, I>>(object: I): MsgRemoveCampaign {
    const message = createBaseMsgRemoveCampaign();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? 0;
    return message;
  },
};

function createBaseMsgRemoveCampaignResponse(): MsgRemoveCampaignResponse {
  return {};
}

export const MsgRemoveCampaignResponse = {
  encode(_: MsgRemoveCampaignResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRemoveCampaignResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRemoveCampaignResponse();
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

  fromJSON(_: any): MsgRemoveCampaignResponse {
    return {};
  },

  toJSON(_: MsgRemoveCampaignResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRemoveCampaignResponse>, I>>(_: I): MsgRemoveCampaignResponse {
    const message = createBaseMsgRemoveCampaignResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Claim(request: MsgClaim): Promise<MsgClaimResponse>;
  InitialClaim(request: MsgInitialClaim): Promise<MsgInitialClaimResponse>;
  CreateCampaign(request: MsgCreateCampaign): Promise<MsgCreateCampaignResponse>;
  AddMission(request: MsgAddMission): Promise<MsgAddMissionResponse>;
  AddClaimRecords(request: MsgAddClaimRecords): Promise<MsgAddClaimRecordsResponse>;
  DeleteClaimRecord(request: MsgDeleteClaimRecord): Promise<MsgDeleteClaimRecordResponse>;
  CloseCampaign(request: MsgCloseCampaign): Promise<MsgCloseCampaignResponse>;
  EnableCampaign(request: MsgEnableCampaign): Promise<MsgEnableCampaignResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RemoveCampaign(request: MsgRemoveCampaign): Promise<MsgRemoveCampaignResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Claim = this.Claim.bind(this);
    this.InitialClaim = this.InitialClaim.bind(this);
    this.CreateCampaign = this.CreateCampaign.bind(this);
    this.AddMission = this.AddMission.bind(this);
    this.AddClaimRecords = this.AddClaimRecords.bind(this);
    this.DeleteClaimRecord = this.DeleteClaimRecord.bind(this);
    this.CloseCampaign = this.CloseCampaign.bind(this);
    this.EnableCampaign = this.EnableCampaign.bind(this);
    this.RemoveCampaign = this.RemoveCampaign.bind(this);
  }
  Claim(request: MsgClaim): Promise<MsgClaimResponse> {
    const data = MsgClaim.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeclaim.Msg", "Claim", data);
    return promise.then((data) => MsgClaimResponse.decode(new _m0.Reader(data)));
  }

  InitialClaim(request: MsgInitialClaim): Promise<MsgInitialClaimResponse> {
    const data = MsgInitialClaim.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeclaim.Msg", "InitialClaim", data);
    return promise.then((data) => MsgInitialClaimResponse.decode(new _m0.Reader(data)));
  }

  CreateCampaign(request: MsgCreateCampaign): Promise<MsgCreateCampaignResponse> {
    const data = MsgCreateCampaign.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeclaim.Msg", "CreateCampaign", data);
    return promise.then((data) => MsgCreateCampaignResponse.decode(new _m0.Reader(data)));
  }

  AddMission(request: MsgAddMission): Promise<MsgAddMissionResponse> {
    const data = MsgAddMission.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeclaim.Msg", "AddMission", data);
    return promise.then((data) => MsgAddMissionResponse.decode(new _m0.Reader(data)));
  }

  AddClaimRecords(request: MsgAddClaimRecords): Promise<MsgAddClaimRecordsResponse> {
    const data = MsgAddClaimRecords.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeclaim.Msg", "AddClaimRecords", data);
    return promise.then((data) => MsgAddClaimRecordsResponse.decode(new _m0.Reader(data)));
  }

  DeleteClaimRecord(request: MsgDeleteClaimRecord): Promise<MsgDeleteClaimRecordResponse> {
    const data = MsgDeleteClaimRecord.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeclaim.Msg", "DeleteClaimRecord", data);
    return promise.then((data) => MsgDeleteClaimRecordResponse.decode(new _m0.Reader(data)));
  }

  CloseCampaign(request: MsgCloseCampaign): Promise<MsgCloseCampaignResponse> {
    const data = MsgCloseCampaign.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeclaim.Msg", "CloseCampaign", data);
    return promise.then((data) => MsgCloseCampaignResponse.decode(new _m0.Reader(data)));
  }

  EnableCampaign(request: MsgEnableCampaign): Promise<MsgEnableCampaignResponse> {
    const data = MsgEnableCampaign.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeclaim.Msg", "EnableCampaign", data);
    return promise.then((data) => MsgEnableCampaignResponse.decode(new _m0.Reader(data)));
  }

  RemoveCampaign(request: MsgRemoveCampaign): Promise<MsgRemoveCampaignResponse> {
    const data = MsgRemoveCampaign.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeclaim.Msg", "RemoveCampaign", data);
    return promise.then((data) => MsgRemoveCampaignResponse.decode(new _m0.Reader(data)));
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
