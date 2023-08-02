/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { CertificateType } from "./certificate_type";
import { Params } from "./params";
import { CertificateOffer, UserCertificates } from "./user_certificates";
import { Device, UserDevices } from "./user_devices";

export const protobufPackage = "chain4energy.c4echain.cfetokenization";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetCertificateTypeRequest {
  id: number;
}

export interface QueryDeviceRequest {
  deviceAddress: string;
}

export interface QueryDeviceResponse {
  device: Device | undefined;
}

export interface QueryDeviceAllRequest {
  pagination: PageRequest | undefined;
}

export interface QueryDeviceAllResponse {
  devices: Device[];
  pagination: PageResponse | undefined;
}

export interface QueryMarketplaceCertificatesAllRequest {
  pagination: PageRequest | undefined;
}

export interface QueryMarketplaceCertificatesAllResponse {
  marketplaceCertificates: CertificateOffer[];
  pagination: PageResponse | undefined;
}

export interface QueryMarketplaceCertificateRequest {
  id: number;
}

export interface QueryMarketplaceCertificateResponse {
  marketplaceCertificate: CertificateOffer | undefined;
}

export interface QueryGetCertificateTypeResponse {
  CertificateType: CertificateType | undefined;
}

export interface QueryAllCertificateTypeRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllCertificateTypeResponse {
  CertificateType: CertificateType[];
  pagination: PageResponse | undefined;
}

export interface QueryGetUserDevicesRequest {
  owner: string;
}

export interface QueryGetUserDevicesResponse {
  UserDevices: UserDevices | undefined;
}

export interface QueryAllUserDevicesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllUserDevicesResponse {
  UserDevices: UserDevices[];
  pagination: PageResponse | undefined;
}

export interface QueryGetUserCertificatesRequest {
  owner: string;
}

export interface QueryGetUserCertificatesResponse {
  UserCertificates: UserCertificates | undefined;
}

export interface QueryAllUserCertificatesRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllUserCertificatesResponse {
  UserCertificates: UserCertificates[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
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

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetCertificateTypeRequest(): QueryGetCertificateTypeRequest {
  return { id: 0 };
}

export const QueryGetCertificateTypeRequest = {
  encode(message: QueryGetCertificateTypeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetCertificateTypeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetCertificateTypeRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCertificateTypeRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryGetCertificateTypeRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetCertificateTypeRequest>, I>>(
    object: I,
  ): QueryGetCertificateTypeRequest {
    const message = createBaseQueryGetCertificateTypeRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseQueryDeviceRequest(): QueryDeviceRequest {
  return { deviceAddress: "" };
}

export const QueryDeviceRequest = {
  encode(message: QueryDeviceRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.deviceAddress !== "") {
      writer.uint32(10).string(message.deviceAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryDeviceRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryDeviceRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.deviceAddress = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDeviceRequest {
    return { deviceAddress: isSet(object.deviceAddress) ? String(object.deviceAddress) : "" };
  },

  toJSON(message: QueryDeviceRequest): unknown {
    const obj: any = {};
    message.deviceAddress !== undefined && (obj.deviceAddress = message.deviceAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryDeviceRequest>, I>>(object: I): QueryDeviceRequest {
    const message = createBaseQueryDeviceRequest();
    message.deviceAddress = object.deviceAddress ?? "";
    return message;
  },
};

function createBaseQueryDeviceResponse(): QueryDeviceResponse {
  return { device: undefined };
}

export const QueryDeviceResponse = {
  encode(message: QueryDeviceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.device !== undefined) {
      Device.encode(message.device, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryDeviceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryDeviceResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.device = Device.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDeviceResponse {
    return { device: isSet(object.device) ? Device.fromJSON(object.device) : undefined };
  },

  toJSON(message: QueryDeviceResponse): unknown {
    const obj: any = {};
    message.device !== undefined && (obj.device = message.device ? Device.toJSON(message.device) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryDeviceResponse>, I>>(object: I): QueryDeviceResponse {
    const message = createBaseQueryDeviceResponse();
    message.device = (object.device !== undefined && object.device !== null)
      ? Device.fromPartial(object.device)
      : undefined;
    return message;
  },
};

function createBaseQueryDeviceAllRequest(): QueryDeviceAllRequest {
  return { pagination: undefined };
}

export const QueryDeviceAllRequest = {
  encode(message: QueryDeviceAllRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryDeviceAllRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryDeviceAllRequest();
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

  fromJSON(object: any): QueryDeviceAllRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryDeviceAllRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryDeviceAllRequest>, I>>(object: I): QueryDeviceAllRequest {
    const message = createBaseQueryDeviceAllRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryDeviceAllResponse(): QueryDeviceAllResponse {
  return { devices: [], pagination: undefined };
}

export const QueryDeviceAllResponse = {
  encode(message: QueryDeviceAllResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.devices) {
      Device.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryDeviceAllResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryDeviceAllResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.devices.push(Device.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryDeviceAllResponse {
    return {
      devices: Array.isArray(object?.devices) ? object.devices.map((e: any) => Device.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryDeviceAllResponse): unknown {
    const obj: any = {};
    if (message.devices) {
      obj.devices = message.devices.map((e) => e ? Device.toJSON(e) : undefined);
    } else {
      obj.devices = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryDeviceAllResponse>, I>>(object: I): QueryDeviceAllResponse {
    const message = createBaseQueryDeviceAllResponse();
    message.devices = object.devices?.map((e) => Device.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryMarketplaceCertificatesAllRequest(): QueryMarketplaceCertificatesAllRequest {
  return { pagination: undefined };
}

export const QueryMarketplaceCertificatesAllRequest = {
  encode(message: QueryMarketplaceCertificatesAllRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryMarketplaceCertificatesAllRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryMarketplaceCertificatesAllRequest();
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

  fromJSON(object: any): QueryMarketplaceCertificatesAllRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryMarketplaceCertificatesAllRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryMarketplaceCertificatesAllRequest>, I>>(
    object: I,
  ): QueryMarketplaceCertificatesAllRequest {
    const message = createBaseQueryMarketplaceCertificatesAllRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryMarketplaceCertificatesAllResponse(): QueryMarketplaceCertificatesAllResponse {
  return { marketplaceCertificates: [], pagination: undefined };
}

export const QueryMarketplaceCertificatesAllResponse = {
  encode(message: QueryMarketplaceCertificatesAllResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.marketplaceCertificates) {
      CertificateOffer.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryMarketplaceCertificatesAllResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryMarketplaceCertificatesAllResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.marketplaceCertificates.push(CertificateOffer.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryMarketplaceCertificatesAllResponse {
    return {
      marketplaceCertificates: Array.isArray(object?.marketplaceCertificates)
        ? object.marketplaceCertificates.map((e: any) => CertificateOffer.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryMarketplaceCertificatesAllResponse): unknown {
    const obj: any = {};
    if (message.marketplaceCertificates) {
      obj.marketplaceCertificates = message.marketplaceCertificates.map((e) =>
        e ? CertificateOffer.toJSON(e) : undefined
      );
    } else {
      obj.marketplaceCertificates = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryMarketplaceCertificatesAllResponse>, I>>(
    object: I,
  ): QueryMarketplaceCertificatesAllResponse {
    const message = createBaseQueryMarketplaceCertificatesAllResponse();
    message.marketplaceCertificates = object.marketplaceCertificates?.map((e) => CertificateOffer.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryMarketplaceCertificateRequest(): QueryMarketplaceCertificateRequest {
  return { id: 0 };
}

export const QueryMarketplaceCertificateRequest = {
  encode(message: QueryMarketplaceCertificateRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryMarketplaceCertificateRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryMarketplaceCertificateRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMarketplaceCertificateRequest {
    return { id: isSet(object.id) ? Number(object.id) : 0 };
  },

  toJSON(message: QueryMarketplaceCertificateRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryMarketplaceCertificateRequest>, I>>(
    object: I,
  ): QueryMarketplaceCertificateRequest {
    const message = createBaseQueryMarketplaceCertificateRequest();
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseQueryMarketplaceCertificateResponse(): QueryMarketplaceCertificateResponse {
  return { marketplaceCertificate: undefined };
}

export const QueryMarketplaceCertificateResponse = {
  encode(message: QueryMarketplaceCertificateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.marketplaceCertificate !== undefined) {
      CertificateOffer.encode(message.marketplaceCertificate, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryMarketplaceCertificateResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryMarketplaceCertificateResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.marketplaceCertificate = CertificateOffer.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryMarketplaceCertificateResponse {
    return {
      marketplaceCertificate: isSet(object.marketplaceCertificate)
        ? CertificateOffer.fromJSON(object.marketplaceCertificate)
        : undefined,
    };
  },

  toJSON(message: QueryMarketplaceCertificateResponse): unknown {
    const obj: any = {};
    message.marketplaceCertificate !== undefined && (obj.marketplaceCertificate = message.marketplaceCertificate
      ? CertificateOffer.toJSON(message.marketplaceCertificate)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryMarketplaceCertificateResponse>, I>>(
    object: I,
  ): QueryMarketplaceCertificateResponse {
    const message = createBaseQueryMarketplaceCertificateResponse();
    message.marketplaceCertificate =
      (object.marketplaceCertificate !== undefined && object.marketplaceCertificate !== null)
        ? CertificateOffer.fromPartial(object.marketplaceCertificate)
        : undefined;
    return message;
  },
};

function createBaseQueryGetCertificateTypeResponse(): QueryGetCertificateTypeResponse {
  return { CertificateType: undefined };
}

export const QueryGetCertificateTypeResponse = {
  encode(message: QueryGetCertificateTypeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.CertificateType !== undefined) {
      CertificateType.encode(message.CertificateType, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetCertificateTypeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetCertificateTypeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.CertificateType = CertificateType.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetCertificateTypeResponse {
    return {
      CertificateType: isSet(object.CertificateType) ? CertificateType.fromJSON(object.CertificateType) : undefined,
    };
  },

  toJSON(message: QueryGetCertificateTypeResponse): unknown {
    const obj: any = {};
    message.CertificateType !== undefined
      && (obj.CertificateType = message.CertificateType ? CertificateType.toJSON(message.CertificateType) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetCertificateTypeResponse>, I>>(
    object: I,
  ): QueryGetCertificateTypeResponse {
    const message = createBaseQueryGetCertificateTypeResponse();
    message.CertificateType = (object.CertificateType !== undefined && object.CertificateType !== null)
      ? CertificateType.fromPartial(object.CertificateType)
      : undefined;
    return message;
  },
};

function createBaseQueryAllCertificateTypeRequest(): QueryAllCertificateTypeRequest {
  return { pagination: undefined };
}

export const QueryAllCertificateTypeRequest = {
  encode(message: QueryAllCertificateTypeRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllCertificateTypeRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllCertificateTypeRequest();
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

  fromJSON(object: any): QueryAllCertificateTypeRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllCertificateTypeRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllCertificateTypeRequest>, I>>(
    object: I,
  ): QueryAllCertificateTypeRequest {
    const message = createBaseQueryAllCertificateTypeRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllCertificateTypeResponse(): QueryAllCertificateTypeResponse {
  return { CertificateType: [], pagination: undefined };
}

export const QueryAllCertificateTypeResponse = {
  encode(message: QueryAllCertificateTypeResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.CertificateType) {
      CertificateType.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllCertificateTypeResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllCertificateTypeResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.CertificateType.push(CertificateType.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllCertificateTypeResponse {
    return {
      CertificateType: Array.isArray(object?.CertificateType)
        ? object.CertificateType.map((e: any) => CertificateType.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllCertificateTypeResponse): unknown {
    const obj: any = {};
    if (message.CertificateType) {
      obj.CertificateType = message.CertificateType.map((e) => e ? CertificateType.toJSON(e) : undefined);
    } else {
      obj.CertificateType = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllCertificateTypeResponse>, I>>(
    object: I,
  ): QueryAllCertificateTypeResponse {
    const message = createBaseQueryAllCertificateTypeResponse();
    message.CertificateType = object.CertificateType?.map((e) => CertificateType.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetUserDevicesRequest(): QueryGetUserDevicesRequest {
  return { owner: "" };
}

export const QueryGetUserDevicesRequest = {
  encode(message: QueryGetUserDevicesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetUserDevicesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetUserDevicesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserDevicesRequest {
    return { owner: isSet(object.owner) ? String(object.owner) : "" };
  },

  toJSON(message: QueryGetUserDevicesRequest): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetUserDevicesRequest>, I>>(object: I): QueryGetUserDevicesRequest {
    const message = createBaseQueryGetUserDevicesRequest();
    message.owner = object.owner ?? "";
    return message;
  },
};

function createBaseQueryGetUserDevicesResponse(): QueryGetUserDevicesResponse {
  return { UserDevices: undefined };
}

export const QueryGetUserDevicesResponse = {
  encode(message: QueryGetUserDevicesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.UserDevices !== undefined) {
      UserDevices.encode(message.UserDevices, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetUserDevicesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetUserDevicesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.UserDevices = UserDevices.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserDevicesResponse {
    return { UserDevices: isSet(object.UserDevices) ? UserDevices.fromJSON(object.UserDevices) : undefined };
  },

  toJSON(message: QueryGetUserDevicesResponse): unknown {
    const obj: any = {};
    message.UserDevices !== undefined
      && (obj.UserDevices = message.UserDevices ? UserDevices.toJSON(message.UserDevices) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetUserDevicesResponse>, I>>(object: I): QueryGetUserDevicesResponse {
    const message = createBaseQueryGetUserDevicesResponse();
    message.UserDevices = (object.UserDevices !== undefined && object.UserDevices !== null)
      ? UserDevices.fromPartial(object.UserDevices)
      : undefined;
    return message;
  },
};

function createBaseQueryAllUserDevicesRequest(): QueryAllUserDevicesRequest {
  return { pagination: undefined };
}

export const QueryAllUserDevicesRequest = {
  encode(message: QueryAllUserDevicesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllUserDevicesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllUserDevicesRequest();
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

  fromJSON(object: any): QueryAllUserDevicesRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllUserDevicesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllUserDevicesRequest>, I>>(object: I): QueryAllUserDevicesRequest {
    const message = createBaseQueryAllUserDevicesRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllUserDevicesResponse(): QueryAllUserDevicesResponse {
  return { UserDevices: [], pagination: undefined };
}

export const QueryAllUserDevicesResponse = {
  encode(message: QueryAllUserDevicesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.UserDevices) {
      UserDevices.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllUserDevicesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllUserDevicesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.UserDevices.push(UserDevices.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllUserDevicesResponse {
    return {
      UserDevices: Array.isArray(object?.UserDevices)
        ? object.UserDevices.map((e: any) => UserDevices.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllUserDevicesResponse): unknown {
    const obj: any = {};
    if (message.UserDevices) {
      obj.UserDevices = message.UserDevices.map((e) => e ? UserDevices.toJSON(e) : undefined);
    } else {
      obj.UserDevices = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllUserDevicesResponse>, I>>(object: I): QueryAllUserDevicesResponse {
    const message = createBaseQueryAllUserDevicesResponse();
    message.UserDevices = object.UserDevices?.map((e) => UserDevices.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetUserCertificatesRequest(): QueryGetUserCertificatesRequest {
  return { owner: "" };
}

export const QueryGetUserCertificatesRequest = {
  encode(message: QueryGetUserCertificatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetUserCertificatesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetUserCertificatesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserCertificatesRequest {
    return { owner: isSet(object.owner) ? String(object.owner) : "" };
  },

  toJSON(message: QueryGetUserCertificatesRequest): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetUserCertificatesRequest>, I>>(
    object: I,
  ): QueryGetUserCertificatesRequest {
    const message = createBaseQueryGetUserCertificatesRequest();
    message.owner = object.owner ?? "";
    return message;
  },
};

function createBaseQueryGetUserCertificatesResponse(): QueryGetUserCertificatesResponse {
  return { UserCertificates: undefined };
}

export const QueryGetUserCertificatesResponse = {
  encode(message: QueryGetUserCertificatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.UserCertificates !== undefined) {
      UserCertificates.encode(message.UserCertificates, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetUserCertificatesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetUserCertificatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.UserCertificates = UserCertificates.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetUserCertificatesResponse {
    return {
      UserCertificates: isSet(object.UserCertificates) ? UserCertificates.fromJSON(object.UserCertificates) : undefined,
    };
  },

  toJSON(message: QueryGetUserCertificatesResponse): unknown {
    const obj: any = {};
    message.UserCertificates !== undefined && (obj.UserCertificates = message.UserCertificates
      ? UserCertificates.toJSON(message.UserCertificates)
      : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetUserCertificatesResponse>, I>>(
    object: I,
  ): QueryGetUserCertificatesResponse {
    const message = createBaseQueryGetUserCertificatesResponse();
    message.UserCertificates = (object.UserCertificates !== undefined && object.UserCertificates !== null)
      ? UserCertificates.fromPartial(object.UserCertificates)
      : undefined;
    return message;
  },
};

function createBaseQueryAllUserCertificatesRequest(): QueryAllUserCertificatesRequest {
  return { pagination: undefined };
}

export const QueryAllUserCertificatesRequest = {
  encode(message: QueryAllUserCertificatesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllUserCertificatesRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllUserCertificatesRequest();
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

  fromJSON(object: any): QueryAllUserCertificatesRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllUserCertificatesRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllUserCertificatesRequest>, I>>(
    object: I,
  ): QueryAllUserCertificatesRequest {
    const message = createBaseQueryAllUserCertificatesRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllUserCertificatesResponse(): QueryAllUserCertificatesResponse {
  return { UserCertificates: [], pagination: undefined };
}

export const QueryAllUserCertificatesResponse = {
  encode(message: QueryAllUserCertificatesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.UserCertificates) {
      UserCertificates.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllUserCertificatesResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllUserCertificatesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.UserCertificates.push(UserCertificates.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllUserCertificatesResponse {
    return {
      UserCertificates: Array.isArray(object?.UserCertificates)
        ? object.UserCertificates.map((e: any) => UserCertificates.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllUserCertificatesResponse): unknown {
    const obj: any = {};
    if (message.UserCertificates) {
      obj.UserCertificates = message.UserCertificates.map((e) => e ? UserCertificates.toJSON(e) : undefined);
    } else {
      obj.UserCertificates = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllUserCertificatesResponse>, I>>(
    object: I,
  ): QueryAllUserCertificatesResponse {
    const message = createBaseQueryAllUserCertificatesResponse();
    message.UserCertificates = object.UserCertificates?.map((e) => UserCertificates.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of CertificateType items. */
  CertificateType(request: QueryGetCertificateTypeRequest): Promise<QueryGetCertificateTypeResponse>;
  CertificateTypeAll(request: QueryAllCertificateTypeRequest): Promise<QueryAllCertificateTypeResponse>;
  /** Queries a list of UserDevices items. */
  UserDevices(request: QueryGetUserDevicesRequest): Promise<QueryGetUserDevicesResponse>;
  UserDevicesAll(request: QueryAllUserDevicesRequest): Promise<QueryAllUserDevicesResponse>;
  /** Queries a list of UserCertificates items. */
  UserCertificates(request: QueryGetUserCertificatesRequest): Promise<QueryGetUserCertificatesResponse>;
  UserCertificatesAll(request: QueryAllUserCertificatesRequest): Promise<QueryAllUserCertificatesResponse>;
  Device(request: QueryDeviceRequest): Promise<QueryDeviceResponse>;
  DeviceAll(request: QueryDeviceAllRequest): Promise<QueryDeviceAllResponse>;
  MarketplaceCertificatesAll(
    request: QueryMarketplaceCertificatesAllRequest,
  ): Promise<QueryMarketplaceCertificatesAllResponse>;
  MarketplaceCertificate(request: QueryMarketplaceCertificateRequest): Promise<QueryMarketplaceCertificateResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.CertificateType = this.CertificateType.bind(this);
    this.CertificateTypeAll = this.CertificateTypeAll.bind(this);
    this.UserDevices = this.UserDevices.bind(this);
    this.UserDevicesAll = this.UserDevicesAll.bind(this);
    this.UserCertificates = this.UserCertificates.bind(this);
    this.UserCertificatesAll = this.UserCertificatesAll.bind(this);
    this.Device = this.Device.bind(this);
    this.DeviceAll = this.DeviceAll.bind(this);
    this.MarketplaceCertificatesAll = this.MarketplaceCertificatesAll.bind(this);
    this.MarketplaceCertificate = this.MarketplaceCertificate.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  CertificateType(request: QueryGetCertificateTypeRequest): Promise<QueryGetCertificateTypeResponse> {
    const data = QueryGetCertificateTypeRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Query", "CertificateType", data);
    return promise.then((data) => QueryGetCertificateTypeResponse.decode(new _m0.Reader(data)));
  }

  CertificateTypeAll(request: QueryAllCertificateTypeRequest): Promise<QueryAllCertificateTypeResponse> {
    const data = QueryAllCertificateTypeRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Query", "CertificateTypeAll", data);
    return promise.then((data) => QueryAllCertificateTypeResponse.decode(new _m0.Reader(data)));
  }

  UserDevices(request: QueryGetUserDevicesRequest): Promise<QueryGetUserDevicesResponse> {
    const data = QueryGetUserDevicesRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Query", "UserDevices", data);
    return promise.then((data) => QueryGetUserDevicesResponse.decode(new _m0.Reader(data)));
  }

  UserDevicesAll(request: QueryAllUserDevicesRequest): Promise<QueryAllUserDevicesResponse> {
    const data = QueryAllUserDevicesRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Query", "UserDevicesAll", data);
    return promise.then((data) => QueryAllUserDevicesResponse.decode(new _m0.Reader(data)));
  }

  UserCertificates(request: QueryGetUserCertificatesRequest): Promise<QueryGetUserCertificatesResponse> {
    const data = QueryGetUserCertificatesRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Query", "UserCertificates", data);
    return promise.then((data) => QueryGetUserCertificatesResponse.decode(new _m0.Reader(data)));
  }

  UserCertificatesAll(request: QueryAllUserCertificatesRequest): Promise<QueryAllUserCertificatesResponse> {
    const data = QueryAllUserCertificatesRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Query", "UserCertificatesAll", data);
    return promise.then((data) => QueryAllUserCertificatesResponse.decode(new _m0.Reader(data)));
  }

  Device(request: QueryDeviceRequest): Promise<QueryDeviceResponse> {
    const data = QueryDeviceRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Query", "Device", data);
    return promise.then((data) => QueryDeviceResponse.decode(new _m0.Reader(data)));
  }

  DeviceAll(request: QueryDeviceAllRequest): Promise<QueryDeviceAllResponse> {
    const data = QueryDeviceAllRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Query", "DeviceAll", data);
    return promise.then((data) => QueryDeviceAllResponse.decode(new _m0.Reader(data)));
  }

  MarketplaceCertificatesAll(
    request: QueryMarketplaceCertificatesAllRequest,
  ): Promise<QueryMarketplaceCertificatesAllResponse> {
    const data = QueryMarketplaceCertificatesAllRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Query", "MarketplaceCertificatesAll", data);
    return promise.then((data) => QueryMarketplaceCertificatesAllResponse.decode(new _m0.Reader(data)));
  }

  MarketplaceCertificate(request: QueryMarketplaceCertificateRequest): Promise<QueryMarketplaceCertificateResponse> {
    const data = QueryMarketplaceCertificateRequest.encode(request).finish();
    const promise = this.rpc.request("chain4energy.c4echain.cfetokenization.Query", "MarketplaceCertificate", data);
    return promise.then((data) => QueryMarketplaceCertificateResponse.decode(new _m0.Reader(data)));
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
