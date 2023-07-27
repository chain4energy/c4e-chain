/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "chain4energy.c4echain.cfetokenization";

export enum CertificateStatus {
  UNKNOWN_CERTIFICATE_STATUS = 0,
  VALID = 1,
  INVALID = 2,
  BURNED = 3,
  UNRECOGNIZED = -1,
}

export function certificateStatusFromJSON(object: any): CertificateStatus {
  switch (object) {
    case 0:
    case "UNKNOWN_CERTIFICATE_STATUS":
      return CertificateStatus.UNKNOWN_CERTIFICATE_STATUS;
    case 1:
    case "VALID":
      return CertificateStatus.VALID;
    case 2:
    case "INVALID":
      return CertificateStatus.INVALID;
    case 3:
    case "BURNED":
      return CertificateStatus.BURNED;
    case -1:
    case "UNRECOGNIZED":
    default:
      return CertificateStatus.UNRECOGNIZED;
  }
}

export function certificateStatusToJSON(object: CertificateStatus): string {
  switch (object) {
    case CertificateStatus.UNKNOWN_CERTIFICATE_STATUS:
      return "UNKNOWN_CERTIFICATE_STATUS";
    case CertificateStatus.VALID:
      return "VALID";
    case CertificateStatus.INVALID:
      return "INVALID";
    case CertificateStatus.BURNED:
      return "BURNED";
    case CertificateStatus.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface UserCertificates {
  owner: string;
  certificates: Certificate[];
}

export interface Certificate {
  id: number;
  certyficateTypeId: number;
  power: number;
  deviceAddress: string;
  allowedAuthorities: string[];
  authority: string;
  certificateStatus: CertificateStatus;
}

export interface CertificateOffer {
  id: number;
  certificateId: number;
  owner: string;
  buyer: string;
  price: Coin[];
}

function createBaseUserCertificates(): UserCertificates {
  return { owner: "", certificates: [] };
}

export const UserCertificates = {
  encode(message: UserCertificates, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    for (const v of message.certificates) {
      Certificate.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserCertificates {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserCertificates();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.certificates.push(Certificate.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserCertificates {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      certificates: Array.isArray(object?.certificates)
        ? object.certificates.map((e: any) => Certificate.fromJSON(e))
        : [],
    };
  },

  toJSON(message: UserCertificates): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    if (message.certificates) {
      obj.certificates = message.certificates.map((e) => e ? Certificate.toJSON(e) : undefined);
    } else {
      obj.certificates = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UserCertificates>, I>>(object: I): UserCertificates {
    const message = createBaseUserCertificates();
    message.owner = object.owner ?? "";
    message.certificates = object.certificates?.map((e) => Certificate.fromPartial(e)) || [];
    return message;
  },
};

function createBaseCertificate(): Certificate {
  return {
    id: 0,
    certyficateTypeId: 0,
    power: 0,
    deviceAddress: "",
    allowedAuthorities: [],
    authority: "",
    certificateStatus: 0,
  };
}

export const Certificate = {
  encode(message: Certificate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.certyficateTypeId !== 0) {
      writer.uint32(16).uint64(message.certyficateTypeId);
    }
    if (message.power !== 0) {
      writer.uint32(24).uint64(message.power);
    }
    if (message.deviceAddress !== "") {
      writer.uint32(34).string(message.deviceAddress);
    }
    for (const v of message.allowedAuthorities) {
      writer.uint32(42).string(v!);
    }
    if (message.authority !== "") {
      writer.uint32(50).string(message.authority);
    }
    if (message.certificateStatus !== 0) {
      writer.uint32(56).int32(message.certificateStatus);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Certificate {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCertificate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.certyficateTypeId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.power = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.deviceAddress = reader.string();
          break;
        case 5:
          message.allowedAuthorities.push(reader.string());
          break;
        case 6:
          message.authority = reader.string();
          break;
        case 7:
          message.certificateStatus = reader.int32() as any;
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Certificate {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      certyficateTypeId: isSet(object.certyficateTypeId) ? Number(object.certyficateTypeId) : 0,
      power: isSet(object.power) ? Number(object.power) : 0,
      deviceAddress: isSet(object.deviceAddress) ? String(object.deviceAddress) : "",
      allowedAuthorities: Array.isArray(object?.allowedAuthorities)
        ? object.allowedAuthorities.map((e: any) => String(e))
        : [],
      authority: isSet(object.authority) ? String(object.authority) : "",
      certificateStatus: isSet(object.certificateStatus) ? certificateStatusFromJSON(object.certificateStatus) : 0,
    };
  },

  toJSON(message: Certificate): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.certyficateTypeId !== undefined && (obj.certyficateTypeId = Math.round(message.certyficateTypeId));
    message.power !== undefined && (obj.power = Math.round(message.power));
    message.deviceAddress !== undefined && (obj.deviceAddress = message.deviceAddress);
    if (message.allowedAuthorities) {
      obj.allowedAuthorities = message.allowedAuthorities.map((e) => e);
    } else {
      obj.allowedAuthorities = [];
    }
    message.authority !== undefined && (obj.authority = message.authority);
    message.certificateStatus !== undefined
      && (obj.certificateStatus = certificateStatusToJSON(message.certificateStatus));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Certificate>, I>>(object: I): Certificate {
    const message = createBaseCertificate();
    message.id = object.id ?? 0;
    message.certyficateTypeId = object.certyficateTypeId ?? 0;
    message.power = object.power ?? 0;
    message.deviceAddress = object.deviceAddress ?? "";
    message.allowedAuthorities = object.allowedAuthorities?.map((e) => e) || [];
    message.authority = object.authority ?? "";
    message.certificateStatus = object.certificateStatus ?? 0;
    return message;
  },
};

function createBaseCertificateOffer(): CertificateOffer {
  return { id: 0, certificateId: 0, owner: "", buyer: "", price: [] };
}

export const CertificateOffer = {
  encode(message: CertificateOffer, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.certificateId !== 0) {
      writer.uint32(16).uint64(message.certificateId);
    }
    if (message.owner !== "") {
      writer.uint32(26).string(message.owner);
    }
    if (message.buyer !== "") {
      writer.uint32(34).string(message.buyer);
    }
    for (const v of message.price) {
      Coin.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CertificateOffer {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCertificateOffer();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.certificateId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.owner = reader.string();
          break;
        case 4:
          message.buyer = reader.string();
          break;
        case 5:
          message.price.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): CertificateOffer {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      certificateId: isSet(object.certificateId) ? Number(object.certificateId) : 0,
      owner: isSet(object.owner) ? String(object.owner) : "",
      buyer: isSet(object.buyer) ? String(object.buyer) : "",
      price: Array.isArray(object?.price) ? object.price.map((e: any) => Coin.fromJSON(e)) : [],
    };
  },

  toJSON(message: CertificateOffer): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.certificateId !== undefined && (obj.certificateId = Math.round(message.certificateId));
    message.owner !== undefined && (obj.owner = message.owner);
    message.buyer !== undefined && (obj.buyer = message.buyer);
    if (message.price) {
      obj.price = message.price.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.price = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<CertificateOffer>, I>>(object: I): CertificateOffer {
    const message = createBaseCertificateOffer();
    message.id = object.id ?? 0;
    message.certificateId = object.certificateId ?? 0;
    message.owner = object.owner ?? "";
    message.buyer = object.buyer ?? "";
    message.price = object.price?.map((e) => Coin.fromPartial(e)) || [];
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
