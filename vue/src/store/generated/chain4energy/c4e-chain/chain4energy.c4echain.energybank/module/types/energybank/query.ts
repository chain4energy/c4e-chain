/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../energybank/params";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { EnergyToken } from "../energybank/energy_token";
import { TokenParams } from "../energybank/token_params";

export const protobufPackage = "chain4energy.c4echain.cfeenergybank";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryEnergyTokenUserAddressRequest {
  userAddress: string;
  pagination: PageRequest | undefined;
}

export interface QueryEnergyTokenUserAddressResponse {
  EnergyToken: EnergyToken[];
  pagination: PageResponse | undefined;
}

export interface QueryCurrentBalanceRequest {
  userAddress: string;
  tokenName: string;
}

export interface QueryCurrentBalanceResponse {
  balance: number;
}

export interface QueryGetEnergyTokenRequest {
  id: number;
}

export interface QueryGetEnergyTokenResponse {
  EnergyToken: EnergyToken | undefined;
}

export interface QueryAllEnergyTokenRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllEnergyTokenResponse {
  EnergyToken: EnergyToken[];
  pagination: PageResponse | undefined;
}

export interface QueryGetTokenParamsRequest {
  index: string;
}

export interface QueryGetTokenParamsResponse {
  tokenParams: TokenParams | undefined;
}

export interface QueryAllTokenParamsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllTokenParamsResponse {
  tokenParams: TokenParams[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
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
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryEnergyTokenUserAddressRequest: object = { userAddress: "" };

export const QueryEnergyTokenUserAddressRequest = {
  encode(
    message: QueryEnergyTokenUserAddressRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.userAddress !== "") {
      writer.uint32(10).string(message.userAddress);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryEnergyTokenUserAddressRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryEnergyTokenUserAddressRequest,
    } as QueryEnergyTokenUserAddressRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userAddress = reader.string();
          break;
        case 2:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryEnergyTokenUserAddressRequest {
    const message = {
      ...baseQueryEnergyTokenUserAddressRequest,
    } as QueryEnergyTokenUserAddressRequest;
    if (object.userAddress !== undefined && object.userAddress !== null) {
      message.userAddress = String(object.userAddress);
    } else {
      message.userAddress = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryEnergyTokenUserAddressRequest): unknown {
    const obj: any = {};
    message.userAddress !== undefined &&
      (obj.userAddress = message.userAddress);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryEnergyTokenUserAddressRequest>
  ): QueryEnergyTokenUserAddressRequest {
    const message = {
      ...baseQueryEnergyTokenUserAddressRequest,
    } as QueryEnergyTokenUserAddressRequest;
    if (object.userAddress !== undefined && object.userAddress !== null) {
      message.userAddress = object.userAddress;
    } else {
      message.userAddress = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryEnergyTokenUserAddressResponse: object = {};

export const QueryEnergyTokenUserAddressResponse = {
  encode(
    message: QueryEnergyTokenUserAddressResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.EnergyToken) {
      EnergyToken.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryEnergyTokenUserAddressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryEnergyTokenUserAddressResponse,
    } as QueryEnergyTokenUserAddressResponse;
    message.EnergyToken = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.EnergyToken.push(EnergyToken.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryEnergyTokenUserAddressResponse {
    const message = {
      ...baseQueryEnergyTokenUserAddressResponse,
    } as QueryEnergyTokenUserAddressResponse;
    message.EnergyToken = [];
    if (object.EnergyToken !== undefined && object.EnergyToken !== null) {
      for (const e of object.EnergyToken) {
        message.EnergyToken.push(EnergyToken.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryEnergyTokenUserAddressResponse): unknown {
    const obj: any = {};
    if (message.EnergyToken) {
      obj.EnergyToken = message.EnergyToken.map((e) =>
        e ? EnergyToken.toJSON(e) : undefined
      );
    } else {
      obj.EnergyToken = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryEnergyTokenUserAddressResponse>
  ): QueryEnergyTokenUserAddressResponse {
    const message = {
      ...baseQueryEnergyTokenUserAddressResponse,
    } as QueryEnergyTokenUserAddressResponse;
    message.EnergyToken = [];
    if (object.EnergyToken !== undefined && object.EnergyToken !== null) {
      for (const e of object.EnergyToken) {
        message.EnergyToken.push(EnergyToken.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryCurrentBalanceRequest: object = {
  userAddress: "",
  tokenName: "",
};

export const QueryCurrentBalanceRequest = {
  encode(
    message: QueryCurrentBalanceRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.userAddress !== "") {
      writer.uint32(10).string(message.userAddress);
    }
    if (message.tokenName !== "") {
      writer.uint32(18).string(message.tokenName);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryCurrentBalanceRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCurrentBalanceRequest,
    } as QueryCurrentBalanceRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.userAddress = reader.string();
          break;
        case 2:
          message.tokenName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCurrentBalanceRequest {
    const message = {
      ...baseQueryCurrentBalanceRequest,
    } as QueryCurrentBalanceRequest;
    if (object.userAddress !== undefined && object.userAddress !== null) {
      message.userAddress = String(object.userAddress);
    } else {
      message.userAddress = "";
    }
    if (object.tokenName !== undefined && object.tokenName !== null) {
      message.tokenName = String(object.tokenName);
    } else {
      message.tokenName = "";
    }
    return message;
  },

  toJSON(message: QueryCurrentBalanceRequest): unknown {
    const obj: any = {};
    message.userAddress !== undefined &&
      (obj.userAddress = message.userAddress);
    message.tokenName !== undefined && (obj.tokenName = message.tokenName);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCurrentBalanceRequest>
  ): QueryCurrentBalanceRequest {
    const message = {
      ...baseQueryCurrentBalanceRequest,
    } as QueryCurrentBalanceRequest;
    if (object.userAddress !== undefined && object.userAddress !== null) {
      message.userAddress = object.userAddress;
    } else {
      message.userAddress = "";
    }
    if (object.tokenName !== undefined && object.tokenName !== null) {
      message.tokenName = object.tokenName;
    } else {
      message.tokenName = "";
    }
    return message;
  },
};

const baseQueryCurrentBalanceResponse: object = { balance: 0 };

export const QueryCurrentBalanceResponse = {
  encode(
    message: QueryCurrentBalanceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.balance !== 0) {
      writer.uint32(8).uint64(message.balance);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryCurrentBalanceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryCurrentBalanceResponse,
    } as QueryCurrentBalanceResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.balance = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryCurrentBalanceResponse {
    const message = {
      ...baseQueryCurrentBalanceResponse,
    } as QueryCurrentBalanceResponse;
    if (object.balance !== undefined && object.balance !== null) {
      message.balance = Number(object.balance);
    } else {
      message.balance = 0;
    }
    return message;
  },

  toJSON(message: QueryCurrentBalanceResponse): unknown {
    const obj: any = {};
    message.balance !== undefined && (obj.balance = message.balance);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryCurrentBalanceResponse>
  ): QueryCurrentBalanceResponse {
    const message = {
      ...baseQueryCurrentBalanceResponse,
    } as QueryCurrentBalanceResponse;
    if (object.balance !== undefined && object.balance !== null) {
      message.balance = object.balance;
    } else {
      message.balance = 0;
    }
    return message;
  },
};

const baseQueryGetEnergyTokenRequest: object = { id: 0 };

export const QueryGetEnergyTokenRequest = {
  encode(
    message: QueryGetEnergyTokenRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetEnergyTokenRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetEnergyTokenRequest,
    } as QueryGetEnergyTokenRequest;
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

  fromJSON(object: any): QueryGetEnergyTokenRequest {
    const message = {
      ...baseQueryGetEnergyTokenRequest,
    } as QueryGetEnergyTokenRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: QueryGetEnergyTokenRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetEnergyTokenRequest>
  ): QueryGetEnergyTokenRequest {
    const message = {
      ...baseQueryGetEnergyTokenRequest,
    } as QueryGetEnergyTokenRequest;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseQueryGetEnergyTokenResponse: object = {};

export const QueryGetEnergyTokenResponse = {
  encode(
    message: QueryGetEnergyTokenResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.EnergyToken !== undefined) {
      EnergyToken.encode(
        message.EnergyToken,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetEnergyTokenResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetEnergyTokenResponse,
    } as QueryGetEnergyTokenResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.EnergyToken = EnergyToken.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetEnergyTokenResponse {
    const message = {
      ...baseQueryGetEnergyTokenResponse,
    } as QueryGetEnergyTokenResponse;
    if (object.EnergyToken !== undefined && object.EnergyToken !== null) {
      message.EnergyToken = EnergyToken.fromJSON(object.EnergyToken);
    } else {
      message.EnergyToken = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetEnergyTokenResponse): unknown {
    const obj: any = {};
    message.EnergyToken !== undefined &&
      (obj.EnergyToken = message.EnergyToken
        ? EnergyToken.toJSON(message.EnergyToken)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetEnergyTokenResponse>
  ): QueryGetEnergyTokenResponse {
    const message = {
      ...baseQueryGetEnergyTokenResponse,
    } as QueryGetEnergyTokenResponse;
    if (object.EnergyToken !== undefined && object.EnergyToken !== null) {
      message.EnergyToken = EnergyToken.fromPartial(object.EnergyToken);
    } else {
      message.EnergyToken = undefined;
    }
    return message;
  },
};

const baseQueryAllEnergyTokenRequest: object = {};

export const QueryAllEnergyTokenRequest = {
  encode(
    message: QueryAllEnergyTokenRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllEnergyTokenRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllEnergyTokenRequest,
    } as QueryAllEnergyTokenRequest;
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

  fromJSON(object: any): QueryAllEnergyTokenRequest {
    const message = {
      ...baseQueryAllEnergyTokenRequest,
    } as QueryAllEnergyTokenRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllEnergyTokenRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllEnergyTokenRequest>
  ): QueryAllEnergyTokenRequest {
    const message = {
      ...baseQueryAllEnergyTokenRequest,
    } as QueryAllEnergyTokenRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllEnergyTokenResponse: object = {};

export const QueryAllEnergyTokenResponse = {
  encode(
    message: QueryAllEnergyTokenResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.EnergyToken) {
      EnergyToken.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllEnergyTokenResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllEnergyTokenResponse,
    } as QueryAllEnergyTokenResponse;
    message.EnergyToken = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.EnergyToken.push(EnergyToken.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllEnergyTokenResponse {
    const message = {
      ...baseQueryAllEnergyTokenResponse,
    } as QueryAllEnergyTokenResponse;
    message.EnergyToken = [];
    if (object.EnergyToken !== undefined && object.EnergyToken !== null) {
      for (const e of object.EnergyToken) {
        message.EnergyToken.push(EnergyToken.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllEnergyTokenResponse): unknown {
    const obj: any = {};
    if (message.EnergyToken) {
      obj.EnergyToken = message.EnergyToken.map((e) =>
        e ? EnergyToken.toJSON(e) : undefined
      );
    } else {
      obj.EnergyToken = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllEnergyTokenResponse>
  ): QueryAllEnergyTokenResponse {
    const message = {
      ...baseQueryAllEnergyTokenResponse,
    } as QueryAllEnergyTokenResponse;
    message.EnergyToken = [];
    if (object.EnergyToken !== undefined && object.EnergyToken !== null) {
      for (const e of object.EnergyToken) {
        message.EnergyToken.push(EnergyToken.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetTokenParamsRequest: object = { index: "" };

export const QueryGetTokenParamsRequest = {
  encode(
    message: QueryGetTokenParamsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTokenParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTokenParamsRequest,
    } as QueryGetTokenParamsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTokenParamsRequest {
    const message = {
      ...baseQueryGetTokenParamsRequest,
    } as QueryGetTokenParamsRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetTokenParamsRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTokenParamsRequest>
  ): QueryGetTokenParamsRequest {
    const message = {
      ...baseQueryGetTokenParamsRequest,
    } as QueryGetTokenParamsRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetTokenParamsResponse: object = {};

export const QueryGetTokenParamsResponse = {
  encode(
    message: QueryGetTokenParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.tokenParams !== undefined) {
      TokenParams.encode(
        message.tokenParams,
        writer.uint32(10).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetTokenParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetTokenParamsResponse,
    } as QueryGetTokenParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tokenParams = TokenParams.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetTokenParamsResponse {
    const message = {
      ...baseQueryGetTokenParamsResponse,
    } as QueryGetTokenParamsResponse;
    if (object.tokenParams !== undefined && object.tokenParams !== null) {
      message.tokenParams = TokenParams.fromJSON(object.tokenParams);
    } else {
      message.tokenParams = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetTokenParamsResponse): unknown {
    const obj: any = {};
    message.tokenParams !== undefined &&
      (obj.tokenParams = message.tokenParams
        ? TokenParams.toJSON(message.tokenParams)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetTokenParamsResponse>
  ): QueryGetTokenParamsResponse {
    const message = {
      ...baseQueryGetTokenParamsResponse,
    } as QueryGetTokenParamsResponse;
    if (object.tokenParams !== undefined && object.tokenParams !== null) {
      message.tokenParams = TokenParams.fromPartial(object.tokenParams);
    } else {
      message.tokenParams = undefined;
    }
    return message;
  },
};

const baseQueryAllTokenParamsRequest: object = {};

export const QueryAllTokenParamsRequest = {
  encode(
    message: QueryAllTokenParamsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllTokenParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllTokenParamsRequest,
    } as QueryAllTokenParamsRequest;
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

  fromJSON(object: any): QueryAllTokenParamsRequest {
    const message = {
      ...baseQueryAllTokenParamsRequest,
    } as QueryAllTokenParamsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTokenParamsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTokenParamsRequest>
  ): QueryAllTokenParamsRequest {
    const message = {
      ...baseQueryAllTokenParamsRequest,
    } as QueryAllTokenParamsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllTokenParamsResponse: object = {};

export const QueryAllTokenParamsResponse = {
  encode(
    message: QueryAllTokenParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.tokenParams) {
      TokenParams.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllTokenParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllTokenParamsResponse,
    } as QueryAllTokenParamsResponse;
    message.tokenParams = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.tokenParams.push(TokenParams.decode(reader, reader.uint32()));
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

  fromJSON(object: any): QueryAllTokenParamsResponse {
    const message = {
      ...baseQueryAllTokenParamsResponse,
    } as QueryAllTokenParamsResponse;
    message.tokenParams = [];
    if (object.tokenParams !== undefined && object.tokenParams !== null) {
      for (const e of object.tokenParams) {
        message.tokenParams.push(TokenParams.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllTokenParamsResponse): unknown {
    const obj: any = {};
    if (message.tokenParams) {
      obj.tokenParams = message.tokenParams.map((e) =>
        e ? TokenParams.toJSON(e) : undefined
      );
    } else {
      obj.tokenParams = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllTokenParamsResponse>
  ): QueryAllTokenParamsResponse {
    const message = {
      ...baseQueryAllTokenParamsResponse,
    } as QueryAllTokenParamsResponse;
    message.tokenParams = [];
    if (object.tokenParams !== undefined && object.tokenParams !== null) {
      for (const e of object.tokenParams) {
        message.tokenParams.push(TokenParams.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of EnergyTokenUserAddress items. */
  EnergyTokenUserAddress(
    request: QueryEnergyTokenUserAddressRequest
  ): Promise<QueryEnergyTokenUserAddressResponse>;
  /** Queries a list of CurrentBalance items. */
  CurrentBalance(
    request: QueryCurrentBalanceRequest
  ): Promise<QueryCurrentBalanceResponse>;
  /** Queries a EnergyToken by id. */
  EnergyToken(
    request: QueryGetEnergyTokenRequest
  ): Promise<QueryGetEnergyTokenResponse>;
  /** Queries a list of EnergyToken items. */
  EnergyTokenAll(
    request: QueryAllEnergyTokenRequest
  ): Promise<QueryAllEnergyTokenResponse>;
  /** Queries a TokenParams by index. */
  TokenParams(
    request: QueryGetTokenParamsRequest
  ): Promise<QueryGetTokenParamsResponse>;
  /** Queries a list of TokenParams items. */
  TokenParamsAll(
    request: QueryAllTokenParamsRequest
  ): Promise<QueryAllTokenParamsResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeenergybank.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  EnergyTokenUserAddress(
    request: QueryEnergyTokenUserAddressRequest
  ): Promise<QueryEnergyTokenUserAddressResponse> {
    const data = QueryEnergyTokenUserAddressRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeenergybank.Query",
      "EnergyTokenUserAddress",
      data
    );
    return promise.then((data) =>
      QueryEnergyTokenUserAddressResponse.decode(new Reader(data))
    );
  }

  CurrentBalance(
    request: QueryCurrentBalanceRequest
  ): Promise<QueryCurrentBalanceResponse> {
    const data = QueryCurrentBalanceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeenergybank.Query",
      "CurrentBalance",
      data
    );
    return promise.then((data) =>
      QueryCurrentBalanceResponse.decode(new Reader(data))
    );
  }

  EnergyToken(
    request: QueryGetEnergyTokenRequest
  ): Promise<QueryGetEnergyTokenResponse> {
    const data = QueryGetEnergyTokenRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeenergybank.Query",
      "EnergyToken",
      data
    );
    return promise.then((data) =>
      QueryGetEnergyTokenResponse.decode(new Reader(data))
    );
  }

  EnergyTokenAll(
    request: QueryAllEnergyTokenRequest
  ): Promise<QueryAllEnergyTokenResponse> {
    const data = QueryAllEnergyTokenRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeenergybank.Query",
      "EnergyTokenAll",
      data
    );
    return promise.then((data) =>
      QueryAllEnergyTokenResponse.decode(new Reader(data))
    );
  }

  TokenParams(
    request: QueryGetTokenParamsRequest
  ): Promise<QueryGetTokenParamsResponse> {
    const data = QueryGetTokenParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeenergybank.Query",
      "TokenParams",
      data
    );
    return promise.then((data) =>
      QueryGetTokenParamsResponse.decode(new Reader(data))
    );
  }

  TokenParamsAll(
    request: QueryAllTokenParamsRequest
  ): Promise<QueryAllTokenParamsResponse> {
    const data = QueryAllTokenParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeenergybank.Query",
      "TokenParamsAll",
      data
    );
    return promise.then((data) =>
      QueryAllTokenParamsResponse.decode(new Reader(data))
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
