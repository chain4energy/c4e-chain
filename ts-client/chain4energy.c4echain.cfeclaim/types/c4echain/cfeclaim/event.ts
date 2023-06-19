/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeclaim";

export interface NewCampaign {
  id: string;
  owner: string;
  name: string;
  description: string;
  campaignType: string;
  feegrantAmount: string;
  initialClaimFreeAmount: string;
  enabled: string;
  startTime: string;
  endTime: string;
  lockupPeriod: string;
  vestingPeriod: string;
  vestingPoolName: string;
}

export interface CloseCampaign {
  owner: string;
  campaignId: string;
  campaignCloseAction: string;
}

export interface RemoveCampaign {
  owner: string;
  campaignId: string;
}

export interface EnableCampaign {
  owner: string;
  campaignId: string;
}

export interface AddMission {
  id: string;
  owner: string;
  campaignId: string;
  name: string;
  description: string;
  missionType: string;
  weight: string;
  claimStartDate: string;
}

export interface Claim {
  claimer: string;
  campaignId: string;
  missionId: string;
  amount: string;
}

export interface InitialClaim {
  claimer: string;
  campaignId: string;
  addressToClaim: string;
  amount: string;
}

export interface AddClaimRecords {
  owner: string;
  campaignId: string;
  claimRecordsTotalAmount: string;
  claimRecordsNumber: string;
}

export interface DeleteClaimRecord {
  owner: string;
  campaignId: string;
  userAddress: string;
  claimRecordAmount: string;
}

export interface CompleteMission {
  campaignId: string;
  missionId: string;
  userAddress: string;
}

function createBaseNewCampaign(): NewCampaign {
  return {
    id: "",
    owner: "",
    name: "",
    description: "",
    campaignType: "",
    feegrantAmount: "",
    initialClaimFreeAmount: "",
    enabled: "",
    startTime: "",
    endTime: "",
    lockupPeriod: "",
    vestingPeriod: "",
    vestingPoolName: "",
  };
}

export const NewCampaign = {
  encode(message: NewCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
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
    if (message.campaignType !== "") {
      writer.uint32(42).string(message.campaignType);
    }
    if (message.feegrantAmount !== "") {
      writer.uint32(50).string(message.feegrantAmount);
    }
    if (message.initialClaimFreeAmount !== "") {
      writer.uint32(58).string(message.initialClaimFreeAmount);
    }
    if (message.enabled !== "") {
      writer.uint32(66).string(message.enabled);
    }
    if (message.startTime !== "") {
      writer.uint32(74).string(message.startTime);
    }
    if (message.endTime !== "") {
      writer.uint32(82).string(message.endTime);
    }
    if (message.lockupPeriod !== "") {
      writer.uint32(90).string(message.lockupPeriod);
    }
    if (message.vestingPeriod !== "") {
      writer.uint32(98).string(message.vestingPeriod);
    }
    if (message.vestingPoolName !== "") {
      writer.uint32(106).string(message.vestingPoolName);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): NewCampaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseNewCampaign();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
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
          message.campaignType = reader.string();
          break;
        case 6:
          message.feegrantAmount = reader.string();
          break;
        case 7:
          message.initialClaimFreeAmount = reader.string();
          break;
        case 8:
          message.enabled = reader.string();
          break;
        case 9:
          message.startTime = reader.string();
          break;
        case 10:
          message.endTime = reader.string();
          break;
        case 11:
          message.lockupPeriod = reader.string();
          break;
        case 12:
          message.vestingPeriod = reader.string();
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

  fromJSON(object: any): NewCampaign {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      owner: isSet(object.owner) ? String(object.owner) : "",
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      campaignType: isSet(object.campaignType) ? String(object.campaignType) : "",
      feegrantAmount: isSet(object.feegrantAmount) ? String(object.feegrantAmount) : "",
      initialClaimFreeAmount: isSet(object.initialClaimFreeAmount) ? String(object.initialClaimFreeAmount) : "",
      enabled: isSet(object.enabled) ? String(object.enabled) : "",
      startTime: isSet(object.startTime) ? String(object.startTime) : "",
      endTime: isSet(object.endTime) ? String(object.endTime) : "",
      lockupPeriod: isSet(object.lockupPeriod) ? String(object.lockupPeriod) : "",
      vestingPeriod: isSet(object.vestingPeriod) ? String(object.vestingPeriod) : "",
      vestingPoolName: isSet(object.vestingPoolName) ? String(object.vestingPoolName) : "",
    };
  },

  toJSON(message: NewCampaign): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.owner !== undefined && (obj.owner = message.owner);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.campaignType !== undefined && (obj.campaignType = message.campaignType);
    message.feegrantAmount !== undefined && (obj.feegrantAmount = message.feegrantAmount);
    message.initialClaimFreeAmount !== undefined && (obj.initialClaimFreeAmount = message.initialClaimFreeAmount);
    message.enabled !== undefined && (obj.enabled = message.enabled);
    message.startTime !== undefined && (obj.startTime = message.startTime);
    message.endTime !== undefined && (obj.endTime = message.endTime);
    message.lockupPeriod !== undefined && (obj.lockupPeriod = message.lockupPeriod);
    message.vestingPeriod !== undefined && (obj.vestingPeriod = message.vestingPeriod);
    message.vestingPoolName !== undefined && (obj.vestingPoolName = message.vestingPoolName);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<NewCampaign>, I>>(object: I): NewCampaign {
    const message = createBaseNewCampaign();
    message.id = object.id ?? "";
    message.owner = object.owner ?? "";
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.campaignType = object.campaignType ?? "";
    message.feegrantAmount = object.feegrantAmount ?? "";
    message.initialClaimFreeAmount = object.initialClaimFreeAmount ?? "";
    message.enabled = object.enabled ?? "";
    message.startTime = object.startTime ?? "";
    message.endTime = object.endTime ?? "";
    message.lockupPeriod = object.lockupPeriod ?? "";
    message.vestingPeriod = object.vestingPeriod ?? "";
    message.vestingPoolName = object.vestingPoolName ?? "";
    return message;
  },
};

function createBaseCloseCampaign(): CloseCampaign {
  return { owner: "", campaignId: "", campaignCloseAction: "" };
}

export const CloseCampaign = {
  encode(message: CloseCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== "") {
      writer.uint32(18).string(message.campaignId);
    }
    if (message.campaignCloseAction !== "") {
      writer.uint32(26).string(message.campaignCloseAction);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CloseCampaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCloseCampaign();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.campaignId = reader.string();
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

  fromJSON(object: any): CloseCampaign {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? String(object.campaignId) : "",
      campaignCloseAction: isSet(object.campaignCloseAction) ? String(object.campaignCloseAction) : "",
    };
  },

  toJSON(message: CloseCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    message.campaignCloseAction !== undefined && (obj.campaignCloseAction = message.campaignCloseAction);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CloseCampaign>, I>>(object: I): CloseCampaign {
    const message = createBaseCloseCampaign();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? "";
    message.campaignCloseAction = object.campaignCloseAction ?? "";
    return message;
  },
};

function createBaseRemoveCampaign(): RemoveCampaign {
  return { owner: "", campaignId: "" };
}

export const RemoveCampaign = {
  encode(message: RemoveCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== "") {
      writer.uint32(18).string(message.campaignId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RemoveCampaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRemoveCampaign();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.campaignId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RemoveCampaign {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? String(object.campaignId) : "",
    };
  },

  toJSON(message: RemoveCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RemoveCampaign>, I>>(object: I): RemoveCampaign {
    const message = createBaseRemoveCampaign();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? "";
    return message;
  },
};

function createBaseEnableCampaign(): EnableCampaign {
  return { owner: "", campaignId: "" };
}

export const EnableCampaign = {
  encode(message: EnableCampaign, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== "") {
      writer.uint32(18).string(message.campaignId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): EnableCampaign {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEnableCampaign();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.campaignId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): EnableCampaign {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? String(object.campaignId) : "",
    };
  },

  toJSON(message: EnableCampaign): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<EnableCampaign>, I>>(object: I): EnableCampaign {
    const message = createBaseEnableCampaign();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? "";
    return message;
  },
};

function createBaseAddMission(): AddMission {
  return {
    id: "",
    owner: "",
    campaignId: "",
    name: "",
    description: "",
    missionType: "",
    weight: "",
    claimStartDate: "",
  };
}

export const AddMission = {
  encode(message: AddMission, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    if (message.campaignId !== "") {
      writer.uint32(26).string(message.campaignId);
    }
    if (message.name !== "") {
      writer.uint32(34).string(message.name);
    }
    if (message.description !== "") {
      writer.uint32(42).string(message.description);
    }
    if (message.missionType !== "") {
      writer.uint32(50).string(message.missionType);
    }
    if (message.weight !== "") {
      writer.uint32(58).string(message.weight);
    }
    if (message.claimStartDate !== "") {
      writer.uint32(66).string(message.claimStartDate);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AddMission {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAddMission();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.string();
          break;
        case 2:
          message.owner = reader.string();
          break;
        case 3:
          message.campaignId = reader.string();
          break;
        case 4:
          message.name = reader.string();
          break;
        case 5:
          message.description = reader.string();
          break;
        case 6:
          message.missionType = reader.string();
          break;
        case 7:
          message.weight = reader.string();
          break;
        case 8:
          message.claimStartDate = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AddMission {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? String(object.campaignId) : "",
      name: isSet(object.name) ? String(object.name) : "",
      description: isSet(object.description) ? String(object.description) : "",
      missionType: isSet(object.missionType) ? String(object.missionType) : "",
      weight: isSet(object.weight) ? String(object.weight) : "",
      claimStartDate: isSet(object.claimStartDate) ? String(object.claimStartDate) : "",
    };
  },

  toJSON(message: AddMission): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    message.name !== undefined && (obj.name = message.name);
    message.description !== undefined && (obj.description = message.description);
    message.missionType !== undefined && (obj.missionType = message.missionType);
    message.weight !== undefined && (obj.weight = message.weight);
    message.claimStartDate !== undefined && (obj.claimStartDate = message.claimStartDate);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AddMission>, I>>(object: I): AddMission {
    const message = createBaseAddMission();
    message.id = object.id ?? "";
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? "";
    message.name = object.name ?? "";
    message.description = object.description ?? "";
    message.missionType = object.missionType ?? "";
    message.weight = object.weight ?? "";
    message.claimStartDate = object.claimStartDate ?? "";
    return message;
  },
};

function createBaseClaim(): Claim {
  return { claimer: "", campaignId: "", missionId: "", amount: "" };
}

export const Claim = {
  encode(message: Claim, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.claimer !== "") {
      writer.uint32(10).string(message.claimer);
    }
    if (message.campaignId !== "") {
      writer.uint32(18).string(message.campaignId);
    }
    if (message.missionId !== "") {
      writer.uint32(26).string(message.missionId);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Claim {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseClaim();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.claimer = reader.string();
          break;
        case 2:
          message.campaignId = reader.string();
          break;
        case 3:
          message.missionId = reader.string();
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

  fromJSON(object: any): Claim {
    return {
      claimer: isSet(object.claimer) ? String(object.claimer) : "",
      campaignId: isSet(object.campaignId) ? String(object.campaignId) : "",
      missionId: isSet(object.missionId) ? String(object.missionId) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
    };
  },

  toJSON(message: Claim): unknown {
    const obj: any = {};
    message.claimer !== undefined && (obj.claimer = message.claimer);
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    message.missionId !== undefined && (obj.missionId = message.missionId);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Claim>, I>>(object: I): Claim {
    const message = createBaseClaim();
    message.claimer = object.claimer ?? "";
    message.campaignId = object.campaignId ?? "";
    message.missionId = object.missionId ?? "";
    message.amount = object.amount ?? "";
    return message;
  },
};

function createBaseInitialClaim(): InitialClaim {
  return { claimer: "", campaignId: "", addressToClaim: "", amount: "" };
}

export const InitialClaim = {
  encode(message: InitialClaim, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.claimer !== "") {
      writer.uint32(10).string(message.claimer);
    }
    if (message.campaignId !== "") {
      writer.uint32(18).string(message.campaignId);
    }
    if (message.addressToClaim !== "") {
      writer.uint32(26).string(message.addressToClaim);
    }
    if (message.amount !== "") {
      writer.uint32(34).string(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): InitialClaim {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseInitialClaim();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.claimer = reader.string();
          break;
        case 2:
          message.campaignId = reader.string();
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

  fromJSON(object: any): InitialClaim {
    return {
      claimer: isSet(object.claimer) ? String(object.claimer) : "",
      campaignId: isSet(object.campaignId) ? String(object.campaignId) : "",
      addressToClaim: isSet(object.addressToClaim) ? String(object.addressToClaim) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
    };
  },

  toJSON(message: InitialClaim): unknown {
    const obj: any = {};
    message.claimer !== undefined && (obj.claimer = message.claimer);
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    message.addressToClaim !== undefined && (obj.addressToClaim = message.addressToClaim);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<InitialClaim>, I>>(object: I): InitialClaim {
    const message = createBaseInitialClaim();
    message.claimer = object.claimer ?? "";
    message.campaignId = object.campaignId ?? "";
    message.addressToClaim = object.addressToClaim ?? "";
    message.amount = object.amount ?? "";
    return message;
  },
};

function createBaseAddClaimRecords(): AddClaimRecords {
  return { owner: "", campaignId: "", claimRecordsTotalAmount: "", claimRecordsNumber: "" };
}

export const AddClaimRecords = {
  encode(message: AddClaimRecords, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== "") {
      writer.uint32(18).string(message.campaignId);
    }
    if (message.claimRecordsTotalAmount !== "") {
      writer.uint32(26).string(message.claimRecordsTotalAmount);
    }
    if (message.claimRecordsNumber !== "") {
      writer.uint32(34).string(message.claimRecordsNumber);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): AddClaimRecords {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseAddClaimRecords();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.campaignId = reader.string();
          break;
        case 3:
          message.claimRecordsTotalAmount = reader.string();
          break;
        case 4:
          message.claimRecordsNumber = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AddClaimRecords {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? String(object.campaignId) : "",
      claimRecordsTotalAmount: isSet(object.claimRecordsTotalAmount) ? String(object.claimRecordsTotalAmount) : "",
      claimRecordsNumber: isSet(object.claimRecordsNumber) ? String(object.claimRecordsNumber) : "",
    };
  },

  toJSON(message: AddClaimRecords): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    message.claimRecordsTotalAmount !== undefined && (obj.claimRecordsTotalAmount = message.claimRecordsTotalAmount);
    message.claimRecordsNumber !== undefined && (obj.claimRecordsNumber = message.claimRecordsNumber);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<AddClaimRecords>, I>>(object: I): AddClaimRecords {
    const message = createBaseAddClaimRecords();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? "";
    message.claimRecordsTotalAmount = object.claimRecordsTotalAmount ?? "";
    message.claimRecordsNumber = object.claimRecordsNumber ?? "";
    return message;
  },
};

function createBaseDeleteClaimRecord(): DeleteClaimRecord {
  return { owner: "", campaignId: "", userAddress: "", claimRecordAmount: "" };
}

export const DeleteClaimRecord = {
  encode(message: DeleteClaimRecord, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.campaignId !== "") {
      writer.uint32(18).string(message.campaignId);
    }
    if (message.userAddress !== "") {
      writer.uint32(26).string(message.userAddress);
    }
    if (message.claimRecordAmount !== "") {
      writer.uint32(34).string(message.claimRecordAmount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DeleteClaimRecord {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteClaimRecord();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.campaignId = reader.string();
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

  fromJSON(object: any): DeleteClaimRecord {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      campaignId: isSet(object.campaignId) ? String(object.campaignId) : "",
      userAddress: isSet(object.userAddress) ? String(object.userAddress) : "",
      claimRecordAmount: isSet(object.claimRecordAmount) ? String(object.claimRecordAmount) : "",
    };
  },

  toJSON(message: DeleteClaimRecord): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    message.userAddress !== undefined && (obj.userAddress = message.userAddress);
    message.claimRecordAmount !== undefined && (obj.claimRecordAmount = message.claimRecordAmount);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DeleteClaimRecord>, I>>(object: I): DeleteClaimRecord {
    const message = createBaseDeleteClaimRecord();
    message.owner = object.owner ?? "";
    message.campaignId = object.campaignId ?? "";
    message.userAddress = object.userAddress ?? "";
    message.claimRecordAmount = object.claimRecordAmount ?? "";
    return message;
  },
};

function createBaseCompleteMission(): CompleteMission {
  return { campaignId: "", missionId: "", userAddress: "" };
}

export const CompleteMission = {
  encode(message: CompleteMission, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.campaignId !== "") {
      writer.uint32(10).string(message.campaignId);
    }
    if (message.missionId !== "") {
      writer.uint32(18).string(message.missionId);
    }
    if (message.userAddress !== "") {
      writer.uint32(26).string(message.userAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CompleteMission {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCompleteMission();
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
          message.userAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CompleteMission {
    return {
      campaignId: isSet(object.campaignId) ? String(object.campaignId) : "",
      missionId: isSet(object.missionId) ? String(object.missionId) : "",
      userAddress: isSet(object.userAddress) ? String(object.userAddress) : "",
    };
  },

  toJSON(message: CompleteMission): unknown {
    const obj: any = {};
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    message.missionId !== undefined && (obj.missionId = message.missionId);
    message.userAddress !== undefined && (obj.userAddress = message.userAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CompleteMission>, I>>(object: I): CompleteMission {
    const message = createBaseCompleteMission();
    message.campaignId = object.campaignId ?? "";
    message.missionId = object.missionId ?? "";
    message.userAddress = object.userAddress ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
