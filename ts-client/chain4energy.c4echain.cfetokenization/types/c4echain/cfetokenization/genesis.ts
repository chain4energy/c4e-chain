/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { CertificateType } from "./certificate_type";
import { Params } from "./params";
import { CertificateOffer, UserCertificates } from "./user_certificates";
import { Device, UserDevices } from "./user_devices";

export const protobufPackage = "chain4energy.c4echain.cfetokenization";

/** GenesisState defines the cfetokenization module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  certificateTypeList: CertificateType[];
  certificateTypeCount: number;
  userDevicesList: UserDevices[];
  userCertificatesList: UserCertificates[];
  devicesList: Device[];
  certificates: CertificateOffer[];
}

function createBaseGenesisState(): GenesisState {
  return {
    params: undefined,
    certificateTypeList: [],
    certificateTypeCount: 0,
    userDevicesList: [],
    userCertificatesList: [],
    devicesList: [],
    certificates: [],
  };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.certificateTypeList) {
      CertificateType.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.certificateTypeCount !== 0) {
      writer.uint32(24).uint64(message.certificateTypeCount);
    }
    for (const v of message.userDevicesList) {
      UserDevices.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.userCertificatesList) {
      UserCertificates.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.devicesList) {
      Device.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.certificates) {
      CertificateOffer.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.certificateTypeList.push(CertificateType.decode(reader, reader.uint32()));
          break;
        case 3:
          message.certificateTypeCount = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.userDevicesList.push(UserDevices.decode(reader, reader.uint32()));
          break;
        case 5:
          message.userCertificatesList.push(UserCertificates.decode(reader, reader.uint32()));
          break;
        case 6:
          message.devicesList.push(Device.decode(reader, reader.uint32()));
          break;
        case 7:
          message.certificates.push(CertificateOffer.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      certificateTypeList: Array.isArray(object?.certificateTypeList)
        ? object.certificateTypeList.map((e: any) => CertificateType.fromJSON(e))
        : [],
      certificateTypeCount: isSet(object.certificateTypeCount) ? Number(object.certificateTypeCount) : 0,
      userDevicesList: Array.isArray(object?.userDevicesList)
        ? object.userDevicesList.map((e: any) => UserDevices.fromJSON(e))
        : [],
      userCertificatesList: Array.isArray(object?.userCertificatesList)
        ? object.userCertificatesList.map((e: any) => UserCertificates.fromJSON(e))
        : [],
      devicesList: Array.isArray(object?.devicesList) ? object.devicesList.map((e: any) => Device.fromJSON(e)) : [],
      certificates: Array.isArray(object?.certificates)
        ? object.certificates.map((e: any) => CertificateOffer.fromJSON(e))
        : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.certificateTypeList) {
      obj.certificateTypeList = message.certificateTypeList.map((e) => e ? CertificateType.toJSON(e) : undefined);
    } else {
      obj.certificateTypeList = [];
    }
    message.certificateTypeCount !== undefined && (obj.certificateTypeCount = Math.round(message.certificateTypeCount));
    if (message.userDevicesList) {
      obj.userDevicesList = message.userDevicesList.map((e) => e ? UserDevices.toJSON(e) : undefined);
    } else {
      obj.userDevicesList = [];
    }
    if (message.userCertificatesList) {
      obj.userCertificatesList = message.userCertificatesList.map((e) => e ? UserCertificates.toJSON(e) : undefined);
    } else {
      obj.userCertificatesList = [];
    }
    if (message.devicesList) {
      obj.devicesList = message.devicesList.map((e) => e ? Device.toJSON(e) : undefined);
    } else {
      obj.devicesList = [];
    }
    if (message.certificates) {
      obj.certificates = message.certificates.map((e) => e ? CertificateOffer.toJSON(e) : undefined);
    } else {
      obj.certificates = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.certificateTypeList = object.certificateTypeList?.map((e) => CertificateType.fromPartial(e)) || [];
    message.certificateTypeCount = object.certificateTypeCount ?? 0;
    message.userDevicesList = object.userDevicesList?.map((e) => UserDevices.fromPartial(e)) || [];
    message.userCertificatesList = object.userCertificatesList?.map((e) => UserCertificates.fromPartial(e)) || [];
    message.devicesList = object.devicesList?.map((e) => Device.fromPartial(e)) || [];
    message.certificates = object.certificates?.map((e) => CertificateOffer.fromPartial(e)) || [];
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
