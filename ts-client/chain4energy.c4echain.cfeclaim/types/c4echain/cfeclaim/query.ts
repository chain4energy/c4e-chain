/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Campaign } from "./campaign";
import { UserEntry } from "./claim_record";
import { Mission } from "./mission";

export const protobufPackage = "chain4energy.c4echain.cfeclaim";

export interface QueryUserEntryRequest {
  address: string;
}

export interface QueryUserEntryResponse {
  userEntry: UserEntry | undefined;
}

export interface QueryUsersEntriesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryUsersEntriesResponse {
  usersEntries: UserEntry[];
  pagination: PageResponse | undefined;
}

export interface QueryMissionRequest {
  campaignId: number;
  missionId: number;
}

export interface QueryMissionResponse {
  mission: Mission | undefined;
}

export interface QueryMissionsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryMissionsResponse {
  missions: Mission[];
  pagination: PageResponse | undefined;
}

export interface QueryCampaignMissionsRequest {
  campaignId: number;
}

export interface QueryCampaignMissionsResponse {
  missions: Mission[];
}

export interface QueryCampaignsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryCampaignsResponse {
  campaigns: Campaign[];
  pagination: PageResponse | undefined;
}

export interface QueryCampaignRequest {
  campaignId: number;
}

export interface QueryCampaignResponse {
  campaign: Campaign | undefined;
}

function createBaseQueryUserEntryRequest(): QueryUserEntryRequest {
  return { address: "" };
}

export const QueryUserEntryRequest = {
  encode(message: QueryUserEntryRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUserEntryRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUserEntryRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryUserEntryRequest {
    return { address: isSet(object.address) ? String(object.address) : "" };
  },

  toJSON(message: QueryUserEntryRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryUserEntryRequest>, I>>(object: I): QueryUserEntryRequest {
    const message = createBaseQueryUserEntryRequest();
    message.address = object.address ?? "";
    return message;
  },
};

function createBaseQueryUserEntryResponse(): QueryUserEntryResponse {
  return { userEntry: undefined };
}

export const QueryUserEntryResponse = {
  encode(message: QueryUserEntryResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userEntry !== undefined) {
      UserEntry.encode(message.userEntry, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUserEntryResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUserEntryResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userEntry = UserEntry.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryUserEntryResponse {
    return { userEntry: isSet(object.userEntry) ? UserEntry.fromJSON(object.userEntry) : undefined };
  },

  toJSON(message: QueryUserEntryResponse): unknown {
    const obj: any = {};
    message.userEntry !== undefined
      && (obj.userEntry = message.userEntry ? UserEntry.toJSON(message.userEntry) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryUserEntryResponse>, I>>(object: I): QueryUserEntryResponse {
    const message = createBaseQueryUserEntryResponse();
    message.userEntry = (object.userEntry !== undefined && object.userEntry !== null)
      ? UserEntry.fromPartial(object.userEntry)
      : undefined;
    return message;
  },
};

function createBaseQueryUsersEntriesRequest(): QueryUsersEntriesRequest {
  return { pagination: undefined };
}

export const QueryUsersEntriesRequest = {
  encode(message: QueryUsersEntriesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUsersEntriesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUsersEntriesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryUsersEntriesRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryUsersEntriesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryUsersEntriesRequest>, I>>(object: I): QueryUsersEntriesRequest {
    const message = createBaseQueryUsersEntriesRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryUsersEntriesResponse(): QueryUsersEntriesResponse {
  return { usersEntries: [], pagination: undefined };
}

export const QueryUsersEntriesResponse = {
  encode(message: QueryUsersEntriesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.usersEntries) {
      UserEntry.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUsersEntriesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUsersEntriesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.usersEntries.push(UserEntry.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryUsersEntriesResponse {
    return {
      usersEntries: Array.isArray(object?.usersEntries)
        ? object.usersEntries.map((e: any) => UserEntry.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryUsersEntriesResponse): unknown {
    const obj: any = {};
    if (message.usersEntries) {
      obj.usersEntries = message.usersEntries.map((e) => e ? UserEntry.toJSON(e) : undefined);
    } else {
      obj.usersEntries = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryUsersEntriesResponse>, I>>(object: I): QueryUsersEntriesResponse {
    const message = createBaseQueryUsersEntriesResponse();
    message.usersEntries = object.usersEntries?.map((e) => UserEntry.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryMissionRequest(): QueryMissionRequest {
  return { campaignId: 0, missionId: 0 };
}

export const QueryMissionRequest = {
  encode(message: QueryMissionRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.campaignId !== 0) {
      writer.uint32(8).uint64(message.campaignId);
    }
    if (message.missionId !== 0) {
      writer.uint32(16).uint64(message.missionId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryMissionRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryMissionRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaignId = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.missionId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMissionRequest {
    return {
      campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0,
      missionId: isSet(object.missionId) ? Number(object.missionId) : 0,
    };
  },

  toJSON(message: QueryMissionRequest): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    message.missionId !== undefined && (obj.missionId = Math.round(message.missionId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryMissionRequest>, I>>(object: I): QueryMissionRequest {
    const message = createBaseQueryMissionRequest();
    message.campaignId = object.campaignId ?? 0;
    message.missionId = object.missionId ?? 0;
    return message;
  },
};

function createBaseQueryMissionResponse(): QueryMissionResponse {
  return { mission: undefined };
}

export const QueryMissionResponse = {
  encode(message: QueryMissionResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.mission !== undefined) {
      Mission.encode(message.mission, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryMissionResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryMissionResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.mission = Mission.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMissionResponse {
    return { mission: isSet(object.mission) ? Mission.fromJSON(object.mission) : undefined };
  },

  toJSON(message: QueryMissionResponse): unknown {
    const obj: any = {};
    message.mission !== undefined && (obj.mission = message.mission ? Mission.toJSON(message.mission) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryMissionResponse>, I>>(object: I): QueryMissionResponse {
    const message = createBaseQueryMissionResponse();
    message.mission = (object.mission !== undefined && object.mission !== null)
      ? Mission.fromPartial(object.mission)
      : undefined;
    return message;
  },
};

function createBaseQueryMissionsRequest(): QueryMissionsRequest {
  return { pagination: undefined };
}

export const QueryMissionsRequest = {
  encode(message: QueryMissionsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryMissionsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryMissionsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMissionsRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryMissionsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryMissionsRequest>, I>>(object: I): QueryMissionsRequest {
    const message = createBaseQueryMissionsRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryMissionsResponse(): QueryMissionsResponse {
  return { missions: [], pagination: undefined };
}

export const QueryMissionsResponse = {
  encode(message: QueryMissionsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.missions) {
      Mission.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryMissionsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryMissionsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.missions.push(Mission.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMissionsResponse {
    return {
      missions: Array.isArray(object?.missions) ? object.missions.map((e: any) => Mission.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryMissionsResponse): unknown {
    const obj: any = {};
    if (message.missions) {
      obj.missions = message.missions.map((e) => e ? Mission.toJSON(e) : undefined);
    } else {
      obj.missions = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryMissionsResponse>, I>>(object: I): QueryMissionsResponse {
    const message = createBaseQueryMissionsResponse();
    message.missions = object.missions?.map((e) => Mission.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryCampaignMissionsRequest(): QueryCampaignMissionsRequest {
  return { campaignId: 0 };
}

export const QueryCampaignMissionsRequest = {
  encode(message: QueryCampaignMissionsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.campaignId !== 0) {
      writer.uint32(8).uint64(message.campaignId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCampaignMissionsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCampaignMissionsRequest();
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

  fromJSON(object: any): QueryCampaignMissionsRequest {
    return { campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0 };
  },

  toJSON(message: QueryCampaignMissionsRequest): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCampaignMissionsRequest>, I>>(object: I): QueryCampaignMissionsRequest {
    const message = createBaseQueryCampaignMissionsRequest();
    message.campaignId = object.campaignId ?? 0;
    return message;
  },
};

function createBaseQueryCampaignMissionsResponse(): QueryCampaignMissionsResponse {
  return { missions: [] };
}

export const QueryCampaignMissionsResponse = {
  encode(message: QueryCampaignMissionsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.missions) {
      Mission.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCampaignMissionsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCampaignMissionsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.missions.push(Mission.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCampaignMissionsResponse {
    return { missions: Array.isArray(object?.missions) ? object.missions.map((e: any) => Mission.fromJSON(e)) : [] };
  },

  toJSON(message: QueryCampaignMissionsResponse): unknown {
    const obj: any = {};
    if (message.missions) {
      obj.missions = message.missions.map((e) => e ? Mission.toJSON(e) : undefined);
    } else {
      obj.missions = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCampaignMissionsResponse>, I>>(
    object: I,
  ): QueryCampaignMissionsResponse {
    const message = createBaseQueryCampaignMissionsResponse();
    message.missions = object.missions?.map((e) => Mission.fromPartial(e)) || [];
    return message;
  },
};

function createBaseQueryCampaignsRequest(): QueryCampaignsRequest {
  return { pagination: undefined };
}

export const QueryCampaignsRequest = {
  encode(message: QueryCampaignsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCampaignsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCampaignsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCampaignsRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryCampaignsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCampaignsRequest>, I>>(object: I): QueryCampaignsRequest {
    const message = createBaseQueryCampaignsRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryCampaignsResponse(): QueryCampaignsResponse {
  return { campaigns: [], pagination: undefined };
}

export const QueryCampaignsResponse = {
  encode(message: QueryCampaignsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.campaigns) {
      Campaign.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCampaignsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCampaignsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaigns.push(Campaign.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCampaignsResponse {
    return {
      campaigns: Array.isArray(object?.campaigns) ? object.campaigns.map((e: any) => Campaign.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryCampaignsResponse): unknown {
    const obj: any = {};
    if (message.campaigns) {
      obj.campaigns = message.campaigns.map((e) => e ? Campaign.toJSON(e) : undefined);
    } else {
      obj.campaigns = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCampaignsResponse>, I>>(object: I): QueryCampaignsResponse {
    const message = createBaseQueryCampaignsResponse();
    message.campaigns = object.campaigns?.map((e) => Campaign.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryCampaignRequest(): QueryCampaignRequest {
  return { campaignId: 0 };
}

export const QueryCampaignRequest = {
  encode(message: QueryCampaignRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.campaignId !== 0) {
      writer.uint32(8).uint64(message.campaignId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCampaignRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCampaignRequest();
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

  fromJSON(object: any): QueryCampaignRequest {
    return { campaignId: isSet(object.campaignId) ? Number(object.campaignId) : 0 };
  },

  toJSON(message: QueryCampaignRequest): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = Math.round(message.campaignId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCampaignRequest>, I>>(object: I): QueryCampaignRequest {
    const message = createBaseQueryCampaignRequest();
    message.campaignId = object.campaignId ?? 0;
    return message;
  },
};

function createBaseQueryCampaignResponse(): QueryCampaignResponse {
  return { campaign: undefined };
}

export const QueryCampaignResponse = {
  encode(message: QueryCampaignResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.campaign !== undefined) {
      Campaign.encode(message.campaign, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryCampaignResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryCampaignResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.campaign = Campaign.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCampaignResponse {
    return { campaign: isSet(object.campaign) ? Campaign.fromJSON(object.campaign) : undefined };
  },

  toJSON(message: QueryCampaignResponse): unknown {
    const obj: any = {};
    message.campaign !== undefined && (obj.campaign = message.campaign ? Campaign.toJSON(message.campaign) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryCampaignResponse>, I>>(object: I): QueryCampaignResponse {
    const message = createBaseQueryCampaignResponse();
    message.campaign = (object.campaign !== undefined && object.campaign !== null)
      ? Campaign.fromPartial(object.campaign)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Queries a UserEntry by index. */
  UserEntry(request: QueryUserEntryRequest): Promise<QueryUserEntryResponse>;
  /** Queries a list of all UserEntry items. */
  UsersEntries(request: QueryUsersEntriesRequest): Promise<QueryUsersEntriesResponse>;
  /** Queries a Mission by campaign id and mission id. */
  Mission(request: QueryMissionRequest): Promise<QueryMissionResponse>;
  /** Queries a list of Mission items for a given campaign. */
  CampaignMissions(request: QueryCampaignMissionsRequest): Promise<QueryCampaignMissionsResponse>;
  /** Queries a list of all Missions items. */
  Missions(request: QueryMissionsRequest): Promise<QueryMissionsResponse>;
  /** Queries a Campaign by id. */
  Campaign(request: QueryCampaignRequest): Promise<QueryCampaignResponse>;
  /** Queries a list of all Campaigns items. */
  Campaigns(request: QueryCampaignsRequest): Promise<QueryCampaignsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.UserEntry = this.UserEntry.bind(this);
    this.UsersEntries = this.UsersEntries.bind(this);
    this.Mission = this.Mission.bind(this);
    this.CampaignMissions = this.CampaignMissions.bind(this);
    this.Missions = this.Missions.bind(this);
    this.Campaign = this.Campaign.bind(this);
    this.Campaigns = this.Campaigns.bind(this);
  }
  UserEntry(request: QueryUserEntryRequest): Promise<QueryUserEntryResponse> {
    const data = QueryUserEntryRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeclaim.Query", "UserEntry", data);
    return promise.then((data) => QueryUserEntryResponse.decode(new _m0.Reader(data)));
  }

  UsersEntries(request: QueryUsersEntriesRequest): Promise<QueryUsersEntriesResponse> {
    const data = QueryUsersEntriesRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeclaim.Query", "UsersEntries", data);
    return promise.then((data) => QueryUsersEntriesResponse.decode(new _m0.Reader(data)));
  }

  Mission(request: QueryMissionRequest): Promise<QueryMissionResponse> {
    const data = QueryMissionRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeclaim.Query", "Mission", data);
    return promise.then((data) => QueryMissionResponse.decode(new _m0.Reader(data)));
  }

  CampaignMissions(request: QueryCampaignMissionsRequest): Promise<QueryCampaignMissionsResponse> {
    const data = QueryCampaignMissionsRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeclaim.Query", "CampaignMissions", data);
    return promise.then((data) => QueryCampaignMissionsResponse.decode(new _m0.Reader(data)));
  }

  Missions(request: QueryMissionsRequest): Promise<QueryMissionsResponse> {
    const data = QueryMissionsRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeclaim.Query", "Missions", data);
    return promise.then((data) => QueryMissionsResponse.decode(new _m0.Reader(data)));
  }

  Campaign(request: QueryCampaignRequest): Promise<QueryCampaignResponse> {
    const data = QueryCampaignRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeclaim.Query", "Campaign", data);
    return promise.then((data) => QueryCampaignResponse.decode(new _m0.Reader(data)));
  }

  Campaigns(request: QueryCampaignsRequest): Promise<QueryCampaignsResponse> {
    const data = QueryCampaignsRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfeclaim.Query", "Campaigns", data);
    return promise.then((data) => QueryCampaignsResponse.decode(new _m0.Reader(data)));
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
