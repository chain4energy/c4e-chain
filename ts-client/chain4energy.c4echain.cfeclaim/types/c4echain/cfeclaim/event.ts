/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Duration } from "../../google/protobuf/duration";
import { Timestamp } from "../../google/protobuf/timestamp";
import { CampaignType, campaignTypeFromJSON, campaignTypeToJSON } from "./campaign";
import { MissionType, missionTypeFromJSON, missionTypeToJSON } from "./mission";

export const protobufPackage = "chain4energy.c4echain.cfeclaim";

export interface EventNewCampaign {
  id: number;
  owner: string;
  name: string;
  description: string;
  campaignType: CampaignType;
  removableClaimRecords: boolean;
  feegrantAmount: string;
  initialClaimFreeAmount: string;
  free: string;
  enabled: boolean;
  startTime: Date | undefined;
  endTime: Date | undefined;
  lockupPeriod: Duration | undefined;
  vestingPeriod: Duration | undefined;
  vestingPoolName: string;
}

export interface EventCloseCampaign {
  owner: string;
  campaignId: number;
  campaignCloseAction: string;
}

export interface EventRemoveCampaign {
  owner: string;
  campaignId: number;
}

export interface EventEnableCampaign {
  owner: string;
  campaignId: number;
}

export interface EventAddMission {
  id: number;
  owner: string;
  campaignId: number;
  name: string;
  description: string;
  missionType: MissionType;
  weight: string;
  claimStartDate: Date | undefined;
}

export interface EventClaim {
  claimer: string;
  campaignId: number;
  missionId: number;
  amount: string;
}

export interface EventInitialClaim {
  claimer: string;
  campaignId: number;
  addressToClaim: string;
  amount: string;
}

export interface EventAddClaimRecords {
  owner: string;
  campaignId: number;
  claimRecordsTotalAmount: string;
  claimRecordsNumber: number;
}

export interface EventDeleteClaimRecord {
  owner: string;
  campaignId: number;
  userAddress: string;
  claimRecordAmount: string;
}

export interface EventCompleteMission {
  campaignId: number;
  missionId: number;
  userAddress: string;
}

function createBaseEventNewCampaign(): EventNewCampaign {
  return {
    id: 0,
    owner: "",
    name: "",
    description: "",
    campaignType: 0,
    removableClaimRecords: false,
    feegrantAmount: "",
    initialClaimFreeAmount: "",
    free: "",
    enabled: false,
    startTime: undefined,
    endTime: undefined,
    lockupPeriod: undefined,
    vestingPeriod: undefined,
    vestingPoolName: "",
  };
}

export const EventNewCampaign = {
  encode(message: EventNewCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    if (message.name !== "") {
      writer.uint32(26).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(34).string(message.description);
    }
    if (message.campaignType !== 0) {
      writer.uint32(40).int32(message.campaignType);
    }
    if (message.removableClaimRecords === true) {
      writer.uint32(48).bool(message.removableClaimRecords);
    }
    if (message.feegrantAmount !== "") {
      writer.uint32(58).string(message.feegrantAmount);
    }
    if (message.initialClaimFreeAmount !== "") {
      writer.uint32(66).string(message.initialClaimFreeAmount);
    }
    if (message.free !== "") {
      writer.uint32(74).string(message.free);
    }
    if (message.enabled === true) {
      writer.uint32(80).bool(message.enabled);
    }
    if (message.startTime !== undefined) {
      Timestamp.encode(toTimestamp(message.startTime), writer.uint32(90).fork()).ldelim();
    }
    if (message.endTime !== undefined) {
      Timestamp.encode(toTimestamp(message.endTime), writer.uint32(98).fork()).ldelim();
    }
    if (message.lockupPeriod !== undefined) {
      Duration.encode(message.lockupPeriod, writer.uint32(106).fork()).ldelim();
    }
    if (message.vestingPeriod !== undefined) {
      Duration.encode(message.vestingPeriod, writer.uint32(114).fork()).ldelim();
    }
    if (message.vestingPoolName !== "") {
      writer.uint32(122).string(message.vestingPoolName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EventNewCampaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventNewCampaign();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.owner = reader.string();
          break;
        case 3:
          message.name = reader.string();
          break;
        case 4:
          message.description = reader.string();
          break;
        case 5:
          message.campaignType = reader.int32() as any;
          break;
        case 6:
          message.removableClaimRecords = reader.bool();
          break;
        case 7:
          message.feegrantAmount = reader.string();
          break;
        case 8:
          message.initialClaimFreeAmount = reader.string();
          break;
        case 9:
          message.free = reader.string();
          break;
        case 10:
          message.enabled = reader.bool();
          break;
        case 11:
          message.startTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 12:
          message.endTime = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 13:
          message.lockupPeriod = Duration.decode(reader, reader.uint32());
          break;
        case 14:
          message.vestingPeriod = Duration.decode(reader, reader.uint32());
          break;
        case 15:
          message.vestingPoolName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventNewCampaign {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      owner: isSet(object.owner) ? String(object.owner) : "",
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      campaignType: isSet(object.campaignType) ? campaignTypeFromJSON(object.campaignType) : 0,
      removableClaimRecords: isSet(object.removableClaimRecords) ? Boolean(object.removableClaimRecords) : false,
      feegrantAmount: isSet(object.feegrantAmount) ? String(object.feegrantAmount) : "",
      initialClaimFreeAmount: isSet(object.initialClaimFreeAmount) ? String(object.initialClaimFreeAmount) : "",
      free: isSet(object.free) ? String(object.free) : "",
      enabled: isSet(object.enabled) ? Boolean(object.enabled) : false,
      startTime: isSet(object.startTime) ? fromJsonTimestamp(object.startTime) : undefined,
      endTime: isSet(object.endTime) ? fromJsonTimestamp(object.endTime) : undefined,
      lockupPeriod: isSet(object.lockupPeriod) ? Duration.fromJSON(object.lockupPeriod) : undefined,
      vestingPeriod: isSet(object.vestingPeriod) ? Duration.fromJSON(object.vestingPeriod) : undefined,
      vestingPoolName: isSet(object.vestingPoolName) ? String(object.vestingPoolName) : "",
    };
  },

  toJSON(message: EventNewCampaign): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.owner !== undefined && (obj.owner = message.owner);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.campaignType !== undefined && (obj.campaignType = campaignTypeToJSON(message.campaignType));
    message.removableClaimRecords !== undefined && (obj.removableClaimRecords = message.removableClaimRecords);
    message.feegrantAmount !== undefined && (obj.feegrantAmount = message.feegrantAmount);
    message.initialClaimFreeAmount !== undefined && (obj.initialClaimFreeAmount = message.initialClaimFreeAmount);
    message.free !== undefined && (obj.free = message.free);
    message.enabled !== undefined && (obj.enabled = message.enabled);
    message.startTime !== undefined && (obj.startTime = message.startTime.toISOString());
    message.endTime !== undefined && (obj.endTime = message.endTime.toISOString());
    message.lockupPeriod !== undefined
      && (obj.lockupPeriod = message.lockupPeriod ? Duration.toJSON(message.lockupPeriod) : undefined);
    message.vestingPeriod !== undefined
      && (obj.vestingPeriod = message.vestingPeriod ? Duration.toJSON(message.vestingPeriod) : undefined);
    message.vestingPoolName !== undefined && (obj.vestingPoolName = message.vestingPoolName);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventNewCampaign>, I>>(object: I): EventNewCampaign {
    const message = createBaseEventNewCampaign();
    message.id = object.id ?? 0;
    message.owner = object.owner ?? "";
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.campaignType = object.campaignType ?? 0;
    message.removableClaimRecords = object.removableClaimRecords ?? false;
    message.feegrantAmount = object.feegrantAmount ?? "";
    message.initialClaimFreeAmount = object.initialClaimFreeAmount ?? "";
    message.free = object.free ?? "";
    message.enabled = object.enabled ?? false;
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

function createBaseEventCloseCampaign(): EventCloseCampaign {
  return { owner: "", campaignId: 0, campaignCloseAction: "" };
}

export const EventCloseCampaign = {
  encode(message: EventCloseCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    if (message.campaignCloseAction !== "") {
      writer.uint32(26).string(message.campaignCloseAction);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EventCloseCampaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventCloseCampaign();
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
          message.campaignCloseAction = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventCloseCampaign {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      campaignCloseAction: isSet(object.campaignCloseAction) ? String(object.campaignCloseAction) : "",
    };
  },

  toJSON(message: EventCloseCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.campaignCloseAction !== undefined && (obj.campaignCloseAction = message.campaignCloseAction);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventCloseCampaign>, I>>(object: I): EventCloseCampaign {
    const message = createBaseEventCloseCampaign();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? 0;
    message.campaignCloseAction = object.campaignCloseAction ?? "";
    return message;
  },
};

function createBaseEventRemoveCampaign(): EventRemoveCampaign {
  return { owner: "", campaignId: 0 };
}

export const EventRemoveCampaign = {
  encode(message: EventRemoveCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EventRemoveCampaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventRemoveCampaign();
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

  fromJSON(object: any): EventRemoveCampaign {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
    };
  },

  toJSON(message: EventRemoveCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventRemoveCampaign>, I>>(object: I): EventRemoveCampaign {
    const message = createBaseEventRemoveCampaign();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? 0;
    return message;
  },
};

function createBaseEventEnableCampaign(): EventEnableCampaign {
  return { owner: "", campaignId: 0 };
}

export const EventEnableCampaign = {
  encode(message: EventEnableCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EventEnableCampaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventEnableCampaign();
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

  fromJSON(object: any): EventEnableCampaign {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
    };
  },

  toJSON(message: EventEnableCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventEnableCampaign>, I>>(object: I): EventEnableCampaign {
    const message = createBaseEventEnableCampaign();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? 0;
    return message;
  },
};

function createBaseEventAddMission(): EventAddMission {
  return {
    id: 0,
    owner: "",
    campaignId: 0,
    name: "",
    description: "",
    missionType: 0,
    weight: "",
    claimStartDate: undefined,
  };
}

export const EventAddMission = {
  encode(message: EventAddMission, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(24).uint64(message.campaignId);
    }
    if (message.name !== "") {
      writer.uint32(34).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(42).string(message.description);
    }
    if (message.missionType !== 0) {
      writer.uint32(48).int32(message.missionType);
    }
    if (message.weight !== "") {
      writer.uint32(58).string(message.weight);
    }
    if (message.claimStartDate !== undefined) {
      Timestamp.encode(toTimestamp(message.claimStartDate), writer.uint32(66).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EventAddMission {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventAddMission();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.owner = reader.string();
          break;
        case 3:
          message.campaignId = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.name = reader.string();
          break;
        case 5:
          message.description = reader.string();
          break;
        case 6:
          message.missionType = reader.int32() as any;
          break;
        case 7:
          message.weight = reader.string();
          break;
        case 8:
          message.claimStartDate = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventAddMission {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      missionType: isSet(object.missionType) ? missionTypeFromJSON(object.missionType) : 0,
      weight: isSet(object.weight) ? String(object.weight) : "",
      claimStartDate: isSet(object.claimStartDate) ? fromJsonTimestamp(object.claimStartDate) : undefined,
    };
  },

  toJSON(message: EventAddMission): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.missionType !== undefined && (obj.missionType = missionTypeToJSON(message.missionType));
    message.weight !== undefined && (obj.weight = message.weight);
    message.claimStartDate !== undefined && (obj.claimStartDate = message.claimStartDate.toISOString());
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventAddMission>, I>>(object: I): EventAddMission {
    const message = createBaseEventAddMission();
    message.id = object.id ?? 0;
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

function createBaseEventClaim(): EventClaim {
  return { claimer: "", campaignId: 0, missionId: 0, amount: "" };
}

export const EventClaim = {
  encode(message: EventClaim, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.claimer !== "") {
      writer.uint32(10).string(message.claimer);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    if (message.missionId !== 0) {
      writer.uint32(24).uint64(message.missionId);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EventClaim {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventClaim();
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
        case 4:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventClaim {
    return {
      claimer: isSet(object.claimer) ? String(object.claimer) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      missionId: isSet(object.missionId) ? Number(object.missionId) : 0,
      amount: isSet(object.amount) ? String(object.amount) : "",
    };
  },

  toJSON(message: EventClaim): unknown {
    const obj: any = {};
    message.claimer !== undefined && (obj.claimer = message.claimer);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.missionId !== undefined && (obj.missionId = Math.round(message.missionId));
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventClaim>, I>>(object: I): EventClaim {
    const message = createBaseEventClaim();
    message.claimer = object.claimer ?? "";
    message.campaignId = object.campaignId ?? 0;
    message.missionId = object.missionId ?? 0;
    message.amount = object.amount ?? "";
    return message;
  },
};

function createBaseEventInitialClaim(): EventInitialClaim {
  return { claimer: "", campaignId: 0, addressToClaim: "", amount: "" };
}

export const EventInitialClaim = {
  encode(message: EventInitialClaim, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.claimer !== "") {
      writer.uint32(10).string(message.claimer);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    if (message.addressToClaim !== "") {
      writer.uint32(26).string(message.addressToClaim);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EventInitialClaim {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventInitialClaim();
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
          message.addressToClaim = reader.string();
          break;
        case 4:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventInitialClaim {
    return {
      claimer: isSet(object.claimer) ? String(object.claimer) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      addressToClaim: isSet(object.addressToClaim) ? String(object.addressToClaim) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
    };
  },

  toJSON(message: EventInitialClaim): unknown {
    const obj: any = {};
    message.claimer !== undefined && (obj.claimer = message.claimer);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.addressToClaim !== undefined && (obj.addressToClaim = message.addressToClaim);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventInitialClaim>, I>>(object: I): EventInitialClaim {
    const message = createBaseEventInitialClaim();
    message.claimer = object.claimer ?? "";
    message.campaignId = object.campaignId ?? 0;
    message.addressToClaim = object.addressToClaim ?? "";
    message.amount = object.amount ?? "";
    return message;
  },
};

function createBaseEventAddClaimRecords(): EventAddClaimRecords {
  return { owner: "", campaignId: 0, claimRecordsTotalAmount: "", claimRecordsNumber: 0 };
}

export const EventAddClaimRecords = {
  encode(message: EventAddClaimRecords, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    if (message.claimRecordsTotalAmount !== "") {
      writer.uint32(26).string(message.claimRecordsTotalAmount);
    }
    if (message.claimRecordsNumber !== 0) {
      writer.uint32(32).int64(message.claimRecordsNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EventAddClaimRecords {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventAddClaimRecords();
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
          message.claimRecordsTotalAmount = reader.string();
          break;
        case 4:
          message.claimRecordsNumber = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventAddClaimRecords {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      claimRecordsTotalAmount: isSet(object.claimRecordsTotalAmount) ? String(object.claimRecordsTotalAmount) : "",
      claimRecordsNumber: isSet(object.claimRecordsNumber) ? Number(object.claimRecordsNumber) : 0,
    };
  },

  toJSON(message: EventAddClaimRecords): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.claimRecordsTotalAmount !== undefined && (obj.claimRecordsTotalAmount = message.claimRecordsTotalAmount);
    message.claimRecordsNumber !== undefined && (obj.claimRecordsNumber = Math.round(message.claimRecordsNumber));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventAddClaimRecords>, I>>(object: I): EventAddClaimRecords {
    const message = createBaseEventAddClaimRecords();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? 0;
    message.claimRecordsTotalAmount = object.claimRecordsTotalAmount ?? "";
    message.claimRecordsNumber = object.claimRecordsNumber ?? 0;
    return message;
  },
};

function createBaseEventDeleteClaimRecord(): EventDeleteClaimRecord {
  return { owner: "", campaignId: 0, userAddress: "", claimRecordAmount: "" };
}

export const EventDeleteClaimRecord = {
  encode(message: EventDeleteClaimRecord, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== 0) {
      writer.uint32(16).uint64(message.campaignId);
    }
    if (message.userAddress !== "") {
      writer.uint32(26).string(message.userAddress);
    }
    if (message.claimRecordAmount !== "") {
      writer.uint32(34).string(message.claimRecordAmount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EventDeleteClaimRecord {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventDeleteClaimRecord();
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
        case 4:
          message.claimRecordAmount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EventDeleteClaimRecord {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      userAddress: isSet(object.userAddress) ? String(object.userAddress) : "",
      claimRecordAmount: isSet(object.claimRecordAmount) ? String(object.claimRecordAmount) : "",
    };
  },

  toJSON(message: EventDeleteClaimRecord): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.userAddress !== undefined && (obj.userAddress = message.userAddress);
    message.claimRecordAmount !== undefined && (obj.claimRecordAmount = message.claimRecordAmount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventDeleteClaimRecord>, I>>(object: I): EventDeleteClaimRecord {
    const message = createBaseEventDeleteClaimRecord();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? 0;
    message.userAddress = object.userAddress ?? "";
    message.claimRecordAmount = object.claimRecordAmount ?? "";
    return message;
  },
};

function createBaseEventCompleteMission(): EventCompleteMission {
  return { campaignId: 0, missionId: 0, userAddress: "" };
}

export const EventCompleteMission = {
  encode(message: EventCompleteMission, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.campaignId !== 0) {
      writer.uint32(8).uint64(message.campaignId);
    }
    if (message.missionId !== 0) {
      writer.uint32(16).uint64(message.missionId);
    }
    if (message.userAddress !== "") {
      writer.uint32(26).string(message.userAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EventCompleteMission {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEventCompleteMission();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaignId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.missionId = longToNumber(reader.uint64() as Long);
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

  fromJSON(object: any): EventCompleteMission {
    return {
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      missionId: isSet(object.missionId) ? Number(object.missionId) : 0,
      userAddress: isSet(object.userAddress) ? String(object.userAddress) : "",
    };
  },

  toJSON(message: EventCompleteMission): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.missionId !== undefined && (obj.missionId = Math.round(message.missionId));
    message.userAddress !== undefined && (obj.userAddress = message.userAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EventCompleteMission>, I>>(object: I): EventCompleteMission {
    const message = createBaseEventCompleteMission();
    message.campaignId = object.campaignId ?? 0;
    message.missionId = object.missionId ?? 0;
    message.userAddress = object.userAddress ?? "";
    return message;
  },
};

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
