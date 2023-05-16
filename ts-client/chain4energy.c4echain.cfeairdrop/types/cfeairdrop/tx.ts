/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Duration } from "../google/protobuf/duration";
import { Timestamp } from "../google/protobuf/timestamp";
import {
  CloseAction,
  CloseActionFromJSON,
  CloseActionToJSON,
  ClaimRecord,
  MissionType,
  missionTypeFromJSON,
  missionTypeToJSON,
} from "./claim";

export const protobufPackage = "chain4energy.c4echain.cfeclaim";

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

export interface MsgCreateCampaign {
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

export interface MsgCreateCampaignResponse {
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

export interface MsgAddClaimRecords {
  owner: string;
  campaignId: number;
  claimEntries: ClaimRecord[];
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
  CloseAction: CloseAction;
}

export interface MsgCloseCampaignResponse {
}

export interface MsgEnableCampaign {
  owner: string;
  campaignId: number;
}

export interface MsgEnableCampaignResponse {
}

export interface MsgEditCampaign {
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

export interface MsgEditCampaignResponse {
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

function createBaseMsgCreateCampaign(): MsgCreateCampaign {
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

  fromJSON(object: any): MsgCreateCampaign {
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

  toJSON(message: MsgCreateCampaign): unknown {
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

  fromPartial<I extends Exact<DeepPartial<MsgCreateCampaign>, I>>(object: I): MsgCreateCampaign {
    const message = createBaseMsgCreateCampaign();
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

function createBaseMsgCreateCampaignResponse(): MsgCreateCampaignResponse {
  return {};
}

export const MsgCreateCampaignResponse = {
  encode(_: MsgCreateCampaignResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateCampaignResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateCampaignResponse();
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

  fromJSON(_: any): MsgCreateCampaignResponse {
    return {};
  },

  toJSON(_: MsgCreateCampaignResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateCampaignResponse>, I>>(
    _: I,
  ): MsgCreateCampaignResponse {
    const message = createBaseMsgCreateCampaignResponse();
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

function createBaseMsgAddClaimRecords(): MsgAddClaimRecords {
  return { owner: "", campaignId: 0, claimEntries: [] };
}

export const MsgAddClaimRecords = {
  encode(message: MsgAddClaimRecords, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    for (const v of message.claimEntries) {
      ClaimRecord.encode(v!, writer.uint32(26).fork()).ldelim();
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
          message.claimEntries.push(ClaimRecord.decode(reader, reader.uint32()));
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
      claimEntries: Array.isArray(object?.claimEntries)
        ? object.claimEntries.map((e: any) => ClaimRecord.fromJSON(e))
        : [],
    };
  },

  toJSON(message: MsgAddClaimRecords): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    if (message.claimEntries) {
      obj.claimEntries = message.claimEntries.map((e) => e ? ClaimRecord.toJSON(e) : undefined);
    } else {
      obj.claimEntries = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddClaimRecords>, I>>(object: I): MsgAddClaimRecords {
    const message = createBaseMsgAddClaimRecords();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? 0;
    message.claimEntries = object.claimEntries?.map((e) => ClaimRecord.fromPartial(e)) || [];
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
  return { owner: "", campaignId: 0, CloseAction: 0 };
}

export const MsgCloseCampaign = {
  encode(message: MsgCloseCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    if (message.CloseAction !== 0) {
      writer.uint32(24).int32(message.CloseAction);
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
        case 3:
          message.CloseAction = reader.int32() as any;
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
      CloseAction: isSet(object.CloseAction) ? CloseActionFromJSON(object.CloseAction) : 0,
    };
  },

  toJSON(message: MsgCloseCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.CloseAction !== undefined
      && (obj.CloseAction = CloseActionToJSON(message.CloseAction));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCloseCampaign>, I>>(object: I): MsgCloseCampaign {
    const message = createBaseMsgCloseCampaign();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? 0;
    message.CloseAction = object.CloseAction ?? 0;
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
  return { owner: "", campaignId: 0 };
}

export const MsgEnableCampaign = {
  encode(message: MsgEnableCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
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
    };
  },

  toJSON(message: MsgEnableCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgEnableCampaign>, I>>(object: I): MsgEnableCampaign {
    const message = createBaseMsgEnableCampaign();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? 0;
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

function createBaseMsgEditCampaign(): MsgEditCampaign {
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

export const MsgEditCampaign = {
  encode(message: MsgEditCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgEditCampaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgEditCampaign();
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

  fromJSON(object: any): MsgEditCampaign {
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

  toJSON(message: MsgEditCampaign): unknown {
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

  fromPartial<I extends Exact<DeepPartial<MsgEditCampaign>, I>>(object: I): MsgEditCampaign {
    const message = createBaseMsgEditCampaign();
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

function createBaseMsgEditCampaignResponse(): MsgEditCampaignResponse {
  return {};
}

export const MsgEditCampaignResponse = {
  encode(_: MsgEditCampaignResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgEditCampaignResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgEditCampaignResponse();
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

  fromJSON(_: any): MsgEditCampaignResponse {
    return {};
  },

  toJSON(_: MsgEditCampaignResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgEditCampaignResponse>, I>>(_: I): MsgEditCampaignResponse {
    const message = createBaseMsgEditCampaignResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Claim(request: MsgClaim): Promise<MsgClaimResponse>;
  InitialClaim(request: MsgInitialClaim): Promise<MsgInitialClaimResponse>;
  CreateCampaign(request: MsgCreateCampaign): Promise<MsgCreateCampaignResponse>;
  EditCampaign(request: MsgEditCampaign): Promise<MsgEditCampaignResponse>;
  AddMissionToAidropCampaign(request: MsgAddMissionToAidropCampaign): Promise<MsgAddMissionToAidropCampaignResponse>;
  AddClaimRecords(request: MsgAddClaimRecords): Promise<MsgAddClaimRecordsResponse>;
  DeleteClaimRecord(request: MsgDeleteClaimRecord): Promise<MsgDeleteClaimRecordResponse>;
  CloseCampaign(request: MsgCloseCampaign): Promise<MsgCloseCampaignResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  EnableCampaign(request: MsgEnableCampaign): Promise<MsgEnableCampaignResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Claim = this.Claim.bind(this);
    this.InitialClaim = this.InitialClaim.bind(this);
    this.CreateCampaign = this.CreateCampaign.bind(this);
    this.EditCampaign = this.EditCampaign.bind(this);
    this.AddMissionToAidropCampaign = this.AddMissionToAidropCampaign.bind(this);
    this.AddClaimRecords = this.AddClaimRecords.bind(this);
    this.DeleteClaimRecord = this.DeleteClaimRecord.bind(this);
    this.CloseCampaign = this.CloseCampaign.bind(this);
    this.EnableCampaign = this.EnableCampaign.bind(this);
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

  EditCampaign(request: MsgEditCampaign): Promise<MsgEditCampaignResponse> {
    const data = MsgEditCampaign.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeclaim.Msg", "EditCampaign", data);
    return promise.then((data) => MsgEditCampaignResponse.decode(new _m0.Reader(data)));
  }

  AddMissionToAidropCampaign(request: MsgAddMissionToAidropCampaign): Promise<MsgAddMissionToAidropCampaignResponse> {
    const data = MsgAddMissionToAidropCampaign.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeclaim.Msg", "AddMissionToAidropCampaign", data);
    return promise.then((data) => MsgAddMissionToAidropCampaignResponse.decode(new _m0.Reader(data)));
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
