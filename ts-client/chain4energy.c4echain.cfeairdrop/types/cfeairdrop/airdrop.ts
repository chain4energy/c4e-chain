/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../cosmos/base/v1beta1/coin";
import { Duration } from "../google/protobuf/duration";
import { Timestamp } from "../google/protobuf/timestamp";

export const protobufPackage = "chain4energy.c4echain.cfeairdrop";

export enum MissionType {
  MISSION_TYPE_UNSPECIFIED = 0,
  INITIAL_CLAIM = 1,
  DELEGATION = 2,
  VOTE = 3,
  CLAIM = 4,
  UNRECOGNIZED = -1,
}

export function missionTypeFromJSON(object: any): MissionType {
  switch (object) {
    case 0:
    case "MISSION_TYPE_UNSPECIFIED":
      return MissionType.MISSION_TYPE_UNSPECIFIED;
    case 1:
    case "INITIAL_CLAIM":
      return MissionType.INITIAL_CLAIM;
    case 2:
    case "DELEGATION":
      return MissionType.DELEGATION;
    case 3:
    case "VOTE":
      return MissionType.VOTE;
    case 4:
    case "CLAIM":
      return MissionType.CLAIM;
    case -1:
    case "UNRECOGNIZED":
    default:
      return MissionType.UNRECOGNIZED;
  }
}

export function missionTypeToJSON(object: MissionType): string {
  switch (object) {
    case MissionType.MISSION_TYPE_UNSPECIFIED:
      return "MISSION_TYPE_UNSPECIFIED";
    case MissionType.INITIAL_CLAIM:
      return "INITIAL_CLAIM";
    case MissionType.DELEGATION:
      return "DELEGATION";
    case MissionType.VOTE:
      return "VOTE";
    case MissionType.CLAIM:
      return "CLAIM";
    case MissionType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum AirdropCloseAction {
  AIRDROP_CLOSE_ACTION_UNSPECIFIED = 0,
  SEND_TO_COMMUNITY_POOL = 1,
  BURN = 2,
  SEND_TO_OWNER = 3,
  UNRECOGNIZED = -1,
}

export function airdropCloseActionFromJSON(object: any): AirdropCloseAction {
  switch (object) {
    case 0:
    case "AIRDROP_CLOSE_ACTION_UNSPECIFIED":
      return AirdropCloseAction.AIRDROP_CLOSE_ACTION_UNSPECIFIED;
    case 1:
    case "SEND_TO_COMMUNITY_POOL":
      return AirdropCloseAction.SEND_TO_COMMUNITY_POOL;
    case 2:
    case "BURN":
      return AirdropCloseAction.BURN;
    case 3:
    case "SEND_TO_OWNER":
      return AirdropCloseAction.SEND_TO_OWNER;
    case -1:
    case "UNRECOGNIZED":
    default:
      return AirdropCloseAction.UNRECOGNIZED;
  }
}

export function airdropCloseActionToJSON(object: AirdropCloseAction): string {
  switch (object) {
    case AirdropCloseAction.AIRDROP_CLOSE_ACTION_UNSPECIFIED:
      return "AIRDROP_CLOSE_ACTION_UNSPECIFIED";
    case AirdropCloseAction.SEND_TO_COMMUNITY_POOL:
      return "SEND_TO_COMMUNITY_POOL";
    case AirdropCloseAction.BURN:
      return "BURN";
    case AirdropCloseAction.SEND_TO_OWNER:
      return "SEND_TO_OWNER";
    case AirdropCloseAction.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface UserEntry {
  address: string;
  claimAddress: string;
  airdropEntries: ClaimRecord[];
}

export interface ClaimRecord {
  campaignId: number;
  address: string;
  airdropCoins: Coin[];
  completedMissions: number[];
  claimedMissions: number[];
}

export interface ClaimRecords {
  airdropEntries: ClaimRecord[];
}

export interface AirdropDistrubitions {
  campaignId: number;
  airdropCoins: Coin[];
}

export interface AirdropClaimsLeft {
  campaignId: number;
  airdropCoins: Coin[];
}

export interface Campaign {
  id: number;
  owner: string;
  name: string;
  description: string;
  allowFeegrant: boolean;
  initialClaimFreeAmount: string;
  enabled: boolean;
  startTime: Date | undefined;
  endTime:
    | Date
    | undefined;
  /** period of locked coins from claim */
  lockupPeriod:
    | Duration
    | undefined;
  /** period of vesting coins after lockup period */
  vestingPeriod: Duration | undefined;
}

export interface Mission {
  id: number;
  campaignId: number;
  name: string;
  description: string;
  missionType: MissionType;
  weight: string;
  claimStartDate: Date | undefined;
}

function createBaseUsersEntries(): UserEntry {
  return { address: "", claimAddress: "", airdropEntries: [] };
}

export const UserEntry = {
  encode(message: UserEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.claimAddress !== "") {
      writer.uint32(18).string(message.claimAddress);
    }
    for (const v of message.airdropEntries) {
      ClaimRecord.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUsersEntries();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.claimAddress = reader.string();
          break;
        case 3:
          message.airdropEntries.push(ClaimRecord.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserEntry {
    return {
      address: isSet(object.address) ? String(object.address) : "",
      claimAddress: isSet(object.claimAddress) ? String(object.claimAddress) : "",
      airdropEntries: Array.isArray(object?.airdropEntries)
        ? object.airdropEntries.map((e: any) => ClaimRecord.fromJSON(e))
        : [],
    };
  },

  toJSON(message: UserEntry): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.claimAddress !== undefined && (obj.claimAddress = message.claimAddress);
    if (message.airdropEntries) {
      obj.airdropEntries = message.airdropEntries.map((e) => e ? ClaimRecord.toJSON(e) : undefined);
    } else {
      obj.airdropEntries = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UserEntry>, I>>(object: I): UserEntry {
    const message = createBaseUsersEntries();
    message.address = object.address ?? "";
    message.claimAddress = object.claimAddress ?? "";
    message.airdropEntries = object.airdropEntries?.map((e) => ClaimRecord.fromPartial(e)) || [];
    return message;
  },
};

function createBaseAirdropEntry(): ClaimRecord {
  return { campaignId: 0, address: "", airdropCoins: [], completedMissions: [], claimedMissions: [] };
}

export const ClaimRecord = {
  encode(message: ClaimRecord, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.campaignId !== 0) {
      writer.uint32(8).uint64(message.campaignId);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    for (const v of message.airdropCoins) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    writer.uint32(34).fork();
    for (const v of message.completedMissions) {
      writer.uint64(v);
    }
    writer.ldelim();
    writer.uint32(42).fork();
    for (const v of message.claimedMissions) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ClaimRecord {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAirdropEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaignId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.airdropCoins.push(Coin.decode(reader, reader.uint32()));
          break;
        case 4:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.completedMissions.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.completedMissions.push(longToNumber(reader.uint64() as Long));
          }
          break;
        case 5:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.claimedMissions.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.claimedMissions.push(longToNumber(reader.uint64() as Long));
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ClaimRecord {
    return {
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      address: isSet(object.address) ? String(object.address) : "",
      airdropCoins: Array.isArray(object?.airdropCoins) ? object.airdropCoins.map((e: any) => Coin.fromJSON(e)) : [],
      completedMissions: Array.isArray(object?.completedMissions)
        ? object.completedMissions.map((e: any) => Number(e))
        : [],
      claimedMissions: Array.isArray(object?.claimedMissions) ? object.claimedMissions.map((e: any) => Number(e)) : [],
    };
  },

  toJSON(message: ClaimRecord): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.address !== undefined && (obj.address = message.address);
    if (message.airdropCoins) {
      obj.airdropCoins = message.airdropCoins.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.airdropCoins = [];
    }
    if (message.completedMissions) {
      obj.completedMissions = message.completedMissions.map((e) => Math.round(e));
    } else {
      obj.completedMissions = [];
    }
    if (message.claimedMissions) {
      obj.claimedMissions = message.claimedMissions.map((e) => Math.round(e));
    } else {
      obj.claimedMissions = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ClaimRecord>, I>>(object: I): ClaimRecord {
    const message = createBaseAirdropEntry();
    message.campaignId = object.campaignId ?? 0;
    message.address = object.address ?? "";
    message.airdropCoins = object.airdropCoins?.map((e) => Coin.fromPartial(e)) || [];
    message.completedMissions = object.completedMissions?.map((e) => e) || [];
    message.claimedMissions = object.claimedMissions?.map((e) => e) || [];
    return message;
  },
};

function createBaseAirdropEntries(): ClaimRecords {
  return { airdropEntries: [] };
}

export const ClaimRecords = {
  encode(message: ClaimRecords, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.airdropEntries) {
      ClaimRecord.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): ClaimRecords {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAirdropEntries();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.airdropEntries.push(ClaimRecord.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): ClaimRecords {
    return {
      airdropEntries: Array.isArray(object?.airdropEntries)
        ? object.airdropEntries.map((e: any) => ClaimRecord.fromJSON(e))
        : [],
    };
  },

  toJSON(message: ClaimRecords): unknown {
    const obj: any = {};
    if (message.airdropEntries) {
      obj.airdropEntries = message.airdropEntries.map((e) => e ? ClaimRecord.toJSON(e) : undefined);
    } else {
      obj.airdropEntries = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<ClaimRecords>, I>>(object: I): ClaimRecords {
    const message = createBaseAirdropEntries();
    message.airdropEntries = object.airdropEntries?.map((e) => ClaimRecord.fromPartial(e)) || [];
    return message;
  },
};

function createBaseAirdropDistrubitions(): AirdropDistrubitions {
  return { campaignId: 0, airdropCoins: [] };
}

export const AirdropDistrubitions = {
  encode(message: AirdropDistrubitions, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.campaignId !== 0) {
      writer.uint32(8).uint64(message.campaignId);
    }
    for (const v of message.airdropCoins) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AirdropDistrubitions {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAirdropDistrubitions();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaignId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.airdropCoins.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AirdropDistrubitions {
    return {
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      airdropCoins: Array.isArray(object?.airdropCoins) ? object.airdropCoins.map((e: any) => Coin.fromJSON(e)) : [],
    };
  },

  toJSON(message: AirdropDistrubitions): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    if (message.airdropCoins) {
      obj.airdropCoins = message.airdropCoins.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.airdropCoins = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AirdropDistrubitions>, I>>(object: I): AirdropDistrubitions {
    const message = createBaseAirdropDistrubitions();
    message.campaignId = object.campaignId ?? 0;
    message.airdropCoins = object.airdropCoins?.map((e) => Coin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseAirdropClaimsLeft(): AirdropClaimsLeft {
  return { campaignId: 0, airdropCoins: [] };
}

export const AirdropClaimsLeft = {
  encode(message: AirdropClaimsLeft, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.campaignId !== 0) {
      writer.uint32(8).uint64(message.campaignId);
    }
    for (const v of message.airdropCoins) {
      Coin.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AirdropClaimsLeft {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAirdropClaimsLeft();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaignId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.airdropCoins.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AirdropClaimsLeft {
    return {
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      airdropCoins: Array.isArray(object?.airdropCoins) ? object.airdropCoins.map((e: any) => Coin.fromJSON(e)) : [],
    };
  },

  toJSON(message: AirdropClaimsLeft): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    if (message.airdropCoins) {
      obj.airdropCoins = message.airdropCoins.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.airdropCoins = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AirdropClaimsLeft>, I>>(object: I): AirdropClaimsLeft {
    const message = createBaseAirdropClaimsLeft();
    message.campaignId = object.campaignId ?? 0;
    message.airdropCoins = object.airdropCoins?.map((e) => Coin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseCampaign(): Campaign {
  return {
    id: 0,
    owner: "",
    name: "",
    description: "",
    allowFeegrant: false,
    initialClaimFreeAmount: "",
    enabled: false,
    startTime: undefined,
    endTime: undefined,
    lockupPeriod: undefined,
    vestingPeriod: undefined,
  };
}

export const Campaign = {
  encode(message: Campaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
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
    if (message.allowFeegrant === true) {
      writer.uint32(40).bool(message.allowFeegrant);
    }
    if (message.initialClaimFreeAmount !== "") {
      writer.uint32(50).string(message.initialClaimFreeAmount);
    }
    if (message.enabled === true) {
      writer.uint32(56).bool(message.enabled);
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

  decode(input: _m0.Reader | Uint8Array, length?: number): Campaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCampaign();
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
          message.allowFeegrant = reader.bool();
          break;
        case 6:
          message.initialClaimFreeAmount = reader.string();
          break;
        case 7:
          message.enabled = reader.bool();
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

  fromJSON(object: any): Campaign {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      owner: isSet(object.owner) ? String(object.owner) : "",
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      allowFeegrant: isSet(object.allowFeegrant) ? Boolean(object.allowFeegrant) : false,
      initialClaimFreeAmount: isSet(object.initialClaimFreeAmount) ? String(object.initialClaimFreeAmount) : "",
      enabled: isSet(object.enabled) ? Boolean(object.enabled) : false,
      startTime: isSet(object.startTime) ? fromJsonTimestamp(object.startTime) : undefined,
      endTime: isSet(object.endTime) ? fromJsonTimestamp(object.endTime) : undefined,
      lockupPeriod: isSet(object.lockupPeriod) ? Duration.fromJSON(object.lockupPeriod) : undefined,
      vestingPeriod: isSet(object.vestingPeriod) ? Duration.fromJSON(object.vestingPeriod) : undefined,
    };
  },

  toJSON(message: Campaign): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.owner !== undefined && (obj.owner = message.owner);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.allowFeegrant !== undefined && (obj.allowFeegrant = message.allowFeegrant);
    message.initialClaimFreeAmount !== undefined && (obj.initialClaimFreeAmount = message.initialClaimFreeAmount);
    message.enabled !== undefined && (obj.enabled = message.enabled);
    message.startTime !== undefined && (obj.startTime = message.startTime.toISOString());
    message.endTime !== undefined && (obj.endTime = message.endTime.toISOString());
    message.lockupPeriod !== undefined
      && (obj.lockupPeriod = message.lockupPeriod ? Duration.toJSON(message.lockupPeriod) : undefined);
    message.vestingPeriod !== undefined
      && (obj.vestingPeriod = message.vestingPeriod ? Duration.toJSON(message.vestingPeriod) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Campaign>, I>>(object: I): Campaign {
    const message = createBaseCampaign();
    message.id = object.id ?? 0;
    message.owner = object.owner ?? "";
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.allowFeegrant = object.allowFeegrant ?? false;
    message.initialClaimFreeAmount = object.initialClaimFreeAmount ?? "";
    message.enabled = object.enabled ?? false;
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

function createBaseMission(): Mission {
  return { id: 0, campaignId: 0, name: "", description: "", missionType: 0, weight: "", claimStartDate: undefined };
}

export const Mission = {
  encode(message: Mission, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
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

  decode(input: _m0.Reader | Uint8Array, length?: number): Mission {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMission();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
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

  fromJSON(object: any): Mission {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      missionType: isSet(object.missionType) ? missionTypeFromJSON(object.missionType) : 0,
      weight: isSet(object.weight) ? String(object.weight) : "",
      claimStartDate: isSet(object.claimStartDate) ? fromJsonTimestamp(object.claimStartDate) : undefined,
    };
  },

  toJSON(message: Mission): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.missionType !== undefined && (obj.missionType = missionTypeToJSON(message.missionType));
    message.weight !== undefined && (obj.weight = message.weight);
    message.claimStartDate !== undefined && (obj.claimStartDate = message.claimStartDate.toISOString());
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Mission>, I>>(object: I): Mission {
    const message = createBaseMission();
    message.id = object.id ?? 0;
    message.campaignId = object.campaignId ?? 0;
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.missionType = object.missionType ?? 0;
    message.weight = object.weight ?? "";
    message.claimStartDate = object.claimStartDate ?? undefined;
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
