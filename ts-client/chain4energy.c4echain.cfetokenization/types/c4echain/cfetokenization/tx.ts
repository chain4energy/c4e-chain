/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";
import { Timestamp } from "../../google/protobuf/timestamp";

export const protobufPackage = "chain4energy.c4echain.cfetokenization";

export interface MsgAssignDeviceToUser {
  deviceAddress: string;
  userAddress: string;
}

export interface MsgAssignDeviceToUserResponse {
}

export interface MsgAcceptDevice {
  userAddress: string;
  deviceAddress: string;
  deviceName: string;
  deviceLocation: string;
}

export interface MsgAcceptDeviceResponse {
}

export interface MsgCreateUserCertificates {
  owner: string;
  deviceAddress: string;
  allowedAuthorities: string[];
  certyficateTypeId: number;
  measurements: number[];
}

export interface MsgCreateUserCertificatesResponse {
}

export interface MsgAddMeasurement {
  deviceAddress: string;
  timestamp: Date | undefined;
  activePower: number;
  reversePower: number;
  metadata: string;
}

export interface MsgAddMeasurementResponse {
}

export interface MsgAddCertificateToMarketplace {
  owner: string;
  certificateId: number;
  price: Coin[];
}

export interface MsgAddCertificateToMarketplaceResponse {
}

export interface MsgBuyCertificate {
  buyer: string;
  marketplaceCertificateId: number;
}

export interface MsgBuyCertificateResponse {
}

export interface MsgBurnCertificate {
  owner: string;
  certificateId: number;
  deviceAddress: string;
}

export interface MsgBurnCertificateResponse {
}

export interface MsgAuthorizeCertificate {
  authorizer: string;
  userAddress: string;
  certificateId: number;
  validUntil: Date | undefined;
}

export interface MsgAuthorizeCertificateResponse {
}

function createBaseMsgAssignDeviceToUser(): MsgAssignDeviceToUser {
  return { deviceAddress: "", userAddress: "" };
}

export const MsgAssignDeviceToUser = {
  encode(message: MsgAssignDeviceToUser, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.deviceAddress !== "") {
      writer.uint32(10).string(message.deviceAddress);
    }
    if (message.userAddress !== "") {
      writer.uint32(18).string(message.userAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAssignDeviceToUser {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAssignDeviceToUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.deviceAddress = reader.string();
          break;
        case 2:
          message.userAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAssignDeviceToUser {
    return {
      deviceAddress: isSet(object.deviceAddress) ? String(object.deviceAddress) : "",
      userAddress: isSet(object.userAddress) ? String(object.userAddress) : "",
    };
  },

  toJSON(message: MsgAssignDeviceToUser): unknown {
    const obj: any = {};
    message.deviceAddress !== undefined && (obj.deviceAddress = message.deviceAddress);
    message.userAddress !== undefined && (obj.userAddress = message.userAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAssignDeviceToUser>, I>>(object: I): MsgAssignDeviceToUser {
    const message = createBaseMsgAssignDeviceToUser();
    message.deviceAddress = object.deviceAddress ?? "";
    message.userAddress = object.userAddress ?? "";
    return message;
  },
};

function createBaseMsgAssignDeviceToUserResponse(): MsgAssignDeviceToUserResponse {
  return {};
}

export const MsgAssignDeviceToUserResponse = {
  encode(_: MsgAssignDeviceToUserResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAssignDeviceToUserResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAssignDeviceToUserResponse();
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

  fromJSON(_: any): MsgAssignDeviceToUserResponse {
    return {};
  },

  toJSON(_: MsgAssignDeviceToUserResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAssignDeviceToUserResponse>, I>>(_: I): MsgAssignDeviceToUserResponse {
    const message = createBaseMsgAssignDeviceToUserResponse();
    return message;
  },
};

function createBaseMsgAcceptDevice(): MsgAcceptDevice {
  return { userAddress: "", deviceAddress: "", deviceName: "", deviceLocation: "" };
}

export const MsgAcceptDevice = {
  encode(message: MsgAcceptDevice, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.userAddress !== "") {
      writer.uint32(10).string(message.userAddress);
    }
    if (message.deviceAddress !== "") {
      writer.uint32(18).string(message.deviceAddress);
    }
    if (message.deviceName !== "") {
      writer.uint32(26).string(message.deviceName);
    }
    if (message.deviceLocation !== "") {
      writer.uint32(34).string(message.deviceLocation);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAcceptDevice {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAcceptDevice();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userAddress = reader.string();
          break;
        case 2:
          message.deviceAddress = reader.string();
          break;
        case 3:
          message.deviceName = reader.string();
          break;
        case 4:
          message.deviceLocation = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAcceptDevice {
    return {
      userAddress: isSet(object.userAddress) ? String(object.userAddress) : "",
      deviceAddress: isSet(object.deviceAddress) ? String(object.deviceAddress) : "",
      deviceName: isSet(object.deviceName) ? String(object.deviceName) : "",
      deviceLocation: isSet(object.deviceLocation) ? String(object.deviceLocation) : "",
    };
  },

  toJSON(message: MsgAcceptDevice): unknown {
    const obj: any = {};
    message.userAddress !== undefined && (obj.userAddress = message.userAddress);
    message.deviceAddress !== undefined && (obj.deviceAddress = message.deviceAddress);
    message.deviceName !== undefined && (obj.deviceName = message.deviceName);
    message.deviceLocation !== undefined && (obj.deviceLocation = message.deviceLocation);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAcceptDevice>, I>>(object: I): MsgAcceptDevice {
    const message = createBaseMsgAcceptDevice();
    message.userAddress = object.userAddress ?? "";
    message.deviceAddress = object.deviceAddress ?? "";
    message.deviceName = object.deviceName ?? "";
    message.deviceLocation = object.deviceLocation ?? "";
    return message;
  },
};

function createBaseMsgAcceptDeviceResponse(): MsgAcceptDeviceResponse {
  return {};
}

export const MsgAcceptDeviceResponse = {
  encode(_: MsgAcceptDeviceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAcceptDeviceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAcceptDeviceResponse();
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

  fromJSON(_: any): MsgAcceptDeviceResponse {
    return {};
  },

  toJSON(_: MsgAcceptDeviceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAcceptDeviceResponse>, I>>(_: I): MsgAcceptDeviceResponse {
    const message = createBaseMsgAcceptDeviceResponse();
    return message;
  },
};

function createBaseMsgCreateUserCertificates(): MsgCreateUserCertificates {
  return { owner: "", deviceAddress: "", allowedAuthorities: [], certyficateTypeId: 0, measurements: [] };
}

export const MsgCreateUserCertificates = {
  encode(message: MsgCreateUserCertificates, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.deviceAddress !== "") {
      writer.uint32(18).string(message.deviceAddress);
    }
    for (const v of message.allowedAuthorities) {
      writer.uint32(34).string(v!);
    }
    if (message.certyficateTypeId !== 0) {
      writer.uint32(40).uint64(message.certyficateTypeId);
    }
    writer.uint32(50).fork();
    for (const v of message.measurements) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateUserCertificates {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateUserCertificates();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.deviceAddress = reader.string();
          break;
        case 4:
          message.allowedAuthorities.push(reader.string());
          break;
        case 5:
          message.certyficateTypeId = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.measurements.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.measurements.push(longToNumber(reader.uint64() as Long));
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateUserCertificates {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      deviceAddress: isSet(object.deviceAddress) ? String(object.deviceAddress) : "",
      allowedAuthorities: Array.isArray(object?.allowedAuthorities)
        ? object.allowedAuthorities.map((e: any) => String(e))
        : [],
      certyficateTypeId: isSet(object.certyficateTypeId) ? Number(object.certyficateTypeId) : 0,
      measurements: Array.isArray(object?.measurements) ? object.measurements.map((e: any) => Number(e)) : [],
    };
  },

  toJSON(message: MsgCreateUserCertificates): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.deviceAddress !== undefined && (obj.deviceAddress = message.deviceAddress);
    if (message.allowedAuthorities) {
      obj.allowedAuthorities = message.allowedAuthorities.map((e) => e);
    } else {
      obj.allowedAuthorities = [];
    }
    message.certyficateTypeId !== undefined && (obj.certyficateTypeId = Math.round(message.certyficateTypeId));
    if (message.measurements) {
      obj.measurements = message.measurements.map((e) => Math.round(e));
    } else {
      obj.measurements = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateUserCertificates>, I>>(object: I): MsgCreateUserCertificates {
    const message = createBaseMsgCreateUserCertificates();
    message.owner = object.owner ?? "";
    message.deviceAddress = object.deviceAddress ?? "";
    message.allowedAuthorities = object.allowedAuthorities?.map((e) => e) || [];
    message.certyficateTypeId = object.certyficateTypeId ?? 0;
    message.measurements = object.measurements?.map((e) => e) || [];
    return message;
  },
};

function createBaseMsgCreateUserCertificatesResponse(): MsgCreateUserCertificatesResponse {
  return {};
}

export const MsgCreateUserCertificatesResponse = {
  encode(_: MsgCreateUserCertificatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateUserCertificatesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateUserCertificatesResponse();
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

  fromJSON(_: any): MsgCreateUserCertificatesResponse {
    return {};
  },

  toJSON(_: MsgCreateUserCertificatesResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateUserCertificatesResponse>, I>>(
    _: I,
  ): MsgCreateUserCertificatesResponse {
    const message = createBaseMsgCreateUserCertificatesResponse();
    return message;
  },
};

function createBaseMsgAddMeasurement(): MsgAddMeasurement {
  return { deviceAddress: "", timestamp: undefined, activePower: 0, reversePower: 0, metadata: "" };
}

export const MsgAddMeasurement = {
  encode(message: MsgAddMeasurement, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.deviceAddress !== "") {
      writer.uint32(10).string(message.deviceAddress);
    }
    if (message.timestamp !== undefined) {
      Timestamp.encode(toTimestamp(message.timestamp), writer.uint32(18).fork()).ldelim();
    }
    if (message.activePower !== 0) {
      writer.uint32(24).uint64(message.activePower);
    }
    if (message.reversePower !== 0) {
      writer.uint32(32).uint64(message.reversePower);
    }
    if (message.metadata !== "") {
      writer.uint32(42).string(message.metadata);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddMeasurement {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddMeasurement();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.deviceAddress = reader.string();
          break;
        case 2:
          message.timestamp = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 3:
          message.activePower = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.reversePower = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.metadata = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddMeasurement {
    return {
      deviceAddress: isSet(object.deviceAddress) ? String(object.deviceAddress) : "",
      timestamp: isSet(object.timestamp) ? fromJsonTimestamp(object.timestamp) : undefined,
      activePower: isSet(object.activePower) ? Number(object.activePower) : 0,
      reversePower: isSet(object.reversePower) ? Number(object.reversePower) : 0,
      metadata: isSet(object.metadata) ? String(object.metadata) : "",
    };
  },

  toJSON(message: MsgAddMeasurement): unknown {
    const obj: any = {};
    message.deviceAddress !== undefined && (obj.deviceAddress = message.deviceAddress);
    message.timestamp !== undefined && (obj.timestamp = message.timestamp.toISOString());
    message.activePower !== undefined && (obj.activePower = Math.round(message.activePower));
    message.reversePower !== undefined && (obj.reversePower = Math.round(message.reversePower));
    message.metadata !== undefined && (obj.metadata = message.metadata);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddMeasurement>, I>>(object: I): MsgAddMeasurement {
    const message = createBaseMsgAddMeasurement();
    message.deviceAddress = object.deviceAddress ?? "";
    message.timestamp = object.timestamp ?? undefined;
    message.activePower = object.activePower ?? 0;
    message.reversePower = object.reversePower ?? 0;
    message.metadata = object.metadata ?? "";
    return message;
  },
};

function createBaseMsgAddMeasurementResponse(): MsgAddMeasurementResponse {
  return {};
}

export const MsgAddMeasurementResponse = {
  encode(_: MsgAddMeasurementResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddMeasurementResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddMeasurementResponse();
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

  fromJSON(_: any): MsgAddMeasurementResponse {
    return {};
  },

  toJSON(_: MsgAddMeasurementResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddMeasurementResponse>, I>>(_: I): MsgAddMeasurementResponse {
    const message = createBaseMsgAddMeasurementResponse();
    return message;
  },
};

function createBaseMsgAddCertificateToMarketplace(): MsgAddCertificateToMarketplace {
  return { owner: "", certificateId: 0, price: [] };
}

export const MsgAddCertificateToMarketplace = {
  encode(message: MsgAddCertificateToMarketplace, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.certificateId !== 0) {
      writer.uint32(16).uint64(message.certificateId);
    }
    for (const v of message.price) {
      Coin.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddCertificateToMarketplace {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddCertificateToMarketplace();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.certificateId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.price.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAddCertificateToMarketplace {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      certificateId: isSet(object.certificateId) ? Number(object.certificateId) : 0,
      price: Array.isArray(object?.price) ? object.price.map((e: any) => Coin.fromJSON(e)) : [],
    };
  },

  toJSON(message: MsgAddCertificateToMarketplace): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.certificateId !== undefined && (obj.certificateId = Math.round(message.certificateId));
    if (message.price) {
      obj.price = message.price.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.price = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddCertificateToMarketplace>, I>>(
    object: I,
  ): MsgAddCertificateToMarketplace {
    const message = createBaseMsgAddCertificateToMarketplace();
    message.owner = object.owner ?? "";
    message.certificateId = object.certificateId ?? 0;
    message.price = object.price?.map((e) => Coin.fromPartial(e)) || [];
    return message;
  },
};

function createBaseMsgAddCertificateToMarketplaceResponse(): MsgAddCertificateToMarketplaceResponse {
  return {};
}

export const MsgAddCertificateToMarketplaceResponse = {
  encode(_: MsgAddCertificateToMarketplaceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAddCertificateToMarketplaceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAddCertificateToMarketplaceResponse();
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

  fromJSON(_: any): MsgAddCertificateToMarketplaceResponse {
    return {};
  },

  toJSON(_: MsgAddCertificateToMarketplaceResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAddCertificateToMarketplaceResponse>, I>>(
    _: I,
  ): MsgAddCertificateToMarketplaceResponse {
    const message = createBaseMsgAddCertificateToMarketplaceResponse();
    return message;
  },
};

function createBaseMsgBuyCertificate(): MsgBuyCertificate {
  return { buyer: "", marketplaceCertificateId: 0 };
}

export const MsgBuyCertificate = {
  encode(message: MsgBuyCertificate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.buyer !== "") {
      writer.uint32(10).string(message.buyer);
    }
    if (message.marketplaceCertificateId !== 0) {
      writer.uint32(16).uint64(message.marketplaceCertificateId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBuyCertificate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBuyCertificate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.buyer = reader.string();
          break;
        case 2:
          message.marketplaceCertificateId = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBuyCertificate {
    return {
      buyer: isSet(object.buyer) ? String(object.buyer) : "",
      marketplaceCertificateId: isSet(object.marketplaceCertificateId) ? Number(object.marketplaceCertificateId) : 0,
    };
  },

  toJSON(message: MsgBuyCertificate): unknown {
    const obj: any = {};
    message.buyer !== undefined && (obj.buyer = message.buyer);
    message.marketplaceCertificateId !== undefined
      && (obj.marketplaceCertificateId = Math.round(message.marketplaceCertificateId));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBuyCertificate>, I>>(object: I): MsgBuyCertificate {
    const message = createBaseMsgBuyCertificate();
    message.buyer = object.buyer ?? "";
    message.marketplaceCertificateId = object.marketplaceCertificateId ?? 0;
    return message;
  },
};

function createBaseMsgBuyCertificateResponse(): MsgBuyCertificateResponse {
  return {};
}

export const MsgBuyCertificateResponse = {
  encode(_: MsgBuyCertificateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBuyCertificateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBuyCertificateResponse();
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

  fromJSON(_: any): MsgBuyCertificateResponse {
    return {};
  },

  toJSON(_: MsgBuyCertificateResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBuyCertificateResponse>, I>>(_: I): MsgBuyCertificateResponse {
    const message = createBaseMsgBuyCertificateResponse();
    return message;
  },
};

function createBaseMsgBurnCertificate(): MsgBurnCertificate {
  return { owner: "", certificateId: 0, deviceAddress: "" };
}

export const MsgBurnCertificate = {
  encode(message: MsgBurnCertificate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    if (message.certificateId !== 0) {
      writer.uint32(16).uint64(message.certificateId);
    }
    if (message.deviceAddress !== "") {
      writer.uint32(26).string(message.deviceAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBurnCertificate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBurnCertificate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.certificateId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.deviceAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgBurnCertificate {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      certificateId: isSet(object.certificateId) ? Number(object.certificateId) : 0,
      deviceAddress: isSet(object.deviceAddress) ? String(object.deviceAddress) : "",
    };
  },

  toJSON(message: MsgBurnCertificate): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    message.certificateId !== undefined && (obj.certificateId = Math.round(message.certificateId));
    message.deviceAddress !== undefined && (obj.deviceAddress = message.deviceAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBurnCertificate>, I>>(object: I): MsgBurnCertificate {
    const message = createBaseMsgBurnCertificate();
    message.owner = object.owner ?? "";
    message.certificateId = object.certificateId ?? 0;
    message.deviceAddress = object.deviceAddress ?? "";
    return message;
  },
};

function createBaseMsgBurnCertificateResponse(): MsgBurnCertificateResponse {
  return {};
}

export const MsgBurnCertificateResponse = {
  encode(_: MsgBurnCertificateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgBurnCertificateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgBurnCertificateResponse();
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

  fromJSON(_: any): MsgBurnCertificateResponse {
    return {};
  },

  toJSON(_: MsgBurnCertificateResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgBurnCertificateResponse>, I>>(_: I): MsgBurnCertificateResponse {
    const message = createBaseMsgBurnCertificateResponse();
    return message;
  },
};

function createBaseMsgAuthorizeCertificate(): MsgAuthorizeCertificate {
  return { authorizer: "", userAddress: "", certificateId: 0, validUntil: undefined };
}

export const MsgAuthorizeCertificate = {
  encode(message: MsgAuthorizeCertificate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.authorizer !== "") {
      writer.uint32(10).string(message.authorizer);
    }
    if (message.userAddress !== "") {
      writer.uint32(18).string(message.userAddress);
    }
    if (message.certificateId !== 0) {
      writer.uint32(24).uint64(message.certificateId);
    }
    if (message.validUntil !== undefined) {
      Timestamp.encode(toTimestamp(message.validUntil), writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAuthorizeCertificate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAuthorizeCertificate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.authorizer = reader.string();
          break;
        case 2:
          message.userAddress = reader.string();
          break;
        case 3:
          message.certificateId = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.validUntil = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgAuthorizeCertificate {
    return {
      authorizer: isSet(object.authorizer) ? String(object.authorizer) : "",
      userAddress: isSet(object.userAddress) ? String(object.userAddress) : "",
      certificateId: isSet(object.certificateId) ? Number(object.certificateId) : 0,
      validUntil: isSet(object.validUntil) ? fromJsonTimestamp(object.validUntil) : undefined,
    };
  },

  toJSON(message: MsgAuthorizeCertificate): unknown {
    const obj: any = {};
    message.authorizer !== undefined && (obj.authorizer = message.authorizer);
    message.userAddress !== undefined && (obj.userAddress = message.userAddress);
    message.certificateId !== undefined && (obj.certificateId = Math.round(message.certificateId));
    message.validUntil !== undefined && (obj.validUntil = message.validUntil.toISOString());
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAuthorizeCertificate>, I>>(object: I): MsgAuthorizeCertificate {
    const message = createBaseMsgAuthorizeCertificate();
    message.authorizer = object.authorizer ?? "";
    message.userAddress = object.userAddress ?? "";
    message.certificateId = object.certificateId ?? 0;
    message.validUntil = object.validUntil ?? undefined;
    return message;
  },
};

function createBaseMsgAuthorizeCertificateResponse(): MsgAuthorizeCertificateResponse {
  return {};
}

export const MsgAuthorizeCertificateResponse = {
  encode(_: MsgAuthorizeCertificateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgAuthorizeCertificateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgAuthorizeCertificateResponse();
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

  fromJSON(_: any): MsgAuthorizeCertificateResponse {
    return {};
  },

  toJSON(_: MsgAuthorizeCertificateResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgAuthorizeCertificateResponse>, I>>(_: I): MsgAuthorizeCertificateResponse {
    const message = createBaseMsgAuthorizeCertificateResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  AssignDeviceToUser(request: MsgAssignDeviceToUser): Promise<MsgAssignDeviceToUserResponse>;
  AcceptDevice(request: MsgAcceptDevice): Promise<MsgAcceptDeviceResponse>;
  CreateUserCertificate(request: MsgCreateUserCertificates): Promise<MsgCreateUserCertificatesResponse>;
  AuthorizeCertificate(request: MsgAuthorizeCertificate): Promise<MsgAuthorizeCertificateResponse>;
  AddCertificateToMarketplace(request: MsgAddCertificateToMarketplace): Promise<MsgAddCertificateToMarketplaceResponse>;
  BurnCertificate(request: MsgBurnCertificate): Promise<MsgBurnCertificateResponse>;
  BuyCertificate(request: MsgBuyCertificate): Promise<MsgBuyCertificateResponse>;
  AddMeasurement(request: MsgAddMeasurement): Promise<MsgAddMeasurementResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.AssignDeviceToUser = this.AssignDeviceToUser.bind(this);
    this.AcceptDevice = this.AcceptDevice.bind(this);
    this.CreateUserCertificate = this.CreateUserCertificate.bind(this);
    this.AuthorizeCertificate = this.AuthorizeCertificate.bind(this);
    this.AddCertificateToMarketplace = this.AddCertificateToMarketplace.bind(this);
    this.BurnCertificate = this.BurnCertificate.bind(this);
    this.BuyCertificate = this.BuyCertificate.bind(this);
    this.AddMeasurement = this.AddMeasurement.bind(this);
  }
  AssignDeviceToUser(request: MsgAssignDeviceToUser): Promise<MsgAssignDeviceToUserResponse> {
    const data = MsgAssignDeviceToUser.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Msg", "AssignDeviceToUser", data);
    return promise.then((data) => MsgAssignDeviceToUserResponse.decode(new _m0.Reader(data)));
  }

  AcceptDevice(request: MsgAcceptDevice): Promise<MsgAcceptDeviceResponse> {
    const data = MsgAcceptDevice.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Msg", "AcceptDevice", data);
    return promise.then((data) => MsgAcceptDeviceResponse.decode(new _m0.Reader(data)));
  }

  CreateUserCertificate(request: MsgCreateUserCertificates): Promise<MsgCreateUserCertificatesResponse> {
    const data = MsgCreateUserCertificates.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Msg", "CreateUserCertificate", data);
    return promise.then((data) => MsgCreateUserCertificatesResponse.decode(new _m0.Reader(data)));
  }

  AuthorizeCertificate(request: MsgAuthorizeCertificate): Promise<MsgAuthorizeCertificateResponse> {
    const data = MsgAuthorizeCertificate.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Msg", "AuthorizeCertificate", data);
    return promise.then((data) => MsgAuthorizeCertificateResponse.decode(new _m0.Reader(data)));
  }

  AddCertificateToMarketplace(
    request: MsgAddCertificateToMarketplace,
  ): Promise<MsgAddCertificateToMarketplaceResponse> {
    const data = MsgAddCertificateToMarketplace.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Msg", "AddCertificateToMarketplace", data);
    return promise.then((data) => MsgAddCertificateToMarketplaceResponse.decode(new _m0.Reader(data)));
  }

  BurnCertificate(request: MsgBurnCertificate): Promise<MsgBurnCertificateResponse> {
    const data = MsgBurnCertificate.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Msg", "BurnCertificate", data);
    return promise.then((data) => MsgBurnCertificateResponse.decode(new _m0.Reader(data)));
  }

  BuyCertificate(request: MsgBuyCertificate): Promise<MsgBuyCertificateResponse> {
    const data = MsgBuyCertificate.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Msg", "BuyCertificate", data);
    return promise.then((data) => MsgBuyCertificateResponse.decode(new _m0.Reader(data)));
  }

  AddMeasurement(request: MsgAddMeasurement): Promise<MsgAddMeasurementResponse> {
    const data = MsgAddMeasurement.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Msg", "AddMeasurement", data);
    return promise.then((data) => MsgAddMeasurementResponse.decode(new _m0.Reader(data)));
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
