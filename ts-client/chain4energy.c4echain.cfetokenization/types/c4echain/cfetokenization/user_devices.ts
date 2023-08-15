/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Timestamp } from "../../google/protobuf/timestamp";

export const protobufPackage = "chain4energy.c4echain.cfetokenization";

export interface UserDevices {
  owner: string;
  devices: UserDevice[];
}

export interface UserDevice {
  deviceAddress: string;
  name: string;
  location: string;
}

export interface PendingDevice {
  deviceAddress: string;
  userAddress: string;
}

export interface Device {
  deviceAddress: string;
  measurements: Measurement[];
  activePowerSum: number;
  reversePowerSum: number;
  usedActivePower: number;
  fulfilledReversePower: number;
}

export interface Measurement {
  id: number;
  timestamp: Date | undefined;
  activePower: number;
  usedForCertificate: boolean;
  reversePower: number;
  metadata: string;
}

function createBaseUserDevices(): UserDevices {
  return { owner: "", devices: [] };
}

export const UserDevices = {
  encode(message: UserDevices, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    for (const v of message.devices) {
      UserDevice.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserDevices {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserDevices();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.devices.push(UserDevice.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserDevices {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      devices: Array.isArray(object?.devices) ? object.devices.map((e: any) => UserDevice.fromJSON(e)) : [],
    };
  },

  toJSON(message: UserDevices): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    if (message.devices) {
      obj.devices = message.devices.map((e) => e ? UserDevice.toJSON(e) : undefined);
    } else {
      obj.devices = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UserDevices>, I>>(object: I): UserDevices {
    const message = createBaseUserDevices();
    message.owner = object.owner ?? "";
    message.devices = object.devices?.map((e) => UserDevice.fromPartial(e)) || [];
    return message;
  },
};

function createBaseUserDevice(): UserDevice {
  return { deviceAddress: "", name: "", location: "" };
}

export const UserDevice = {
  encode(message: UserDevice, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.deviceAddress !== "") {
      writer.uint32(10).string(message.deviceAddress);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.location !== "") {
      writer.uint32(26).string(message.location);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UserDevice {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUserDevice();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.deviceAddress = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.location = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): UserDevice {
    return {
      deviceAddress: isSet(object.deviceAddress) ? String(object.deviceAddress) : "",
      name: isSet(object.name) ? String(object.name) : "",
      location: isSet(object.location) ? String(object.location) : "",
    };
  },

  toJSON(message: UserDevice): unknown {
    const obj: any = {};
    message.deviceAddress !== undefined && (obj.deviceAddress = message.deviceAddress);
    message.name !== undefined && (obj.name = message.name);
    message.location !== undefined && (obj.location = message.location);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UserDevice>, I>>(object: I): UserDevice {
    const message = createBaseUserDevice();
    message.deviceAddress = object.deviceAddress ?? "";
    message.name = object.name ?? "";
    message.location = object.location ?? "";
    return message;
  },
};

function createBasePendingDevice(): PendingDevice {
  return { deviceAddress: "", userAddress: "" };
}

export const PendingDevice = {
  encode(message: PendingDevice, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.deviceAddress !== "") {
      writer.uint32(10).string(message.deviceAddress);
    }
    if (message.userAddress !== "") {
      writer.uint32(18).string(message.userAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): PendingDevice {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBasePendingDevice();
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

  fromJSON(object: any): PendingDevice {
    return {
      deviceAddress: isSet(object.deviceAddress) ? String(object.deviceAddress) : "",
      userAddress: isSet(object.userAddress) ? String(object.userAddress) : "",
    };
  },

  toJSON(message: PendingDevice): unknown {
    const obj: any = {};
    message.deviceAddress !== undefined && (obj.deviceAddress = message.deviceAddress);
    message.userAddress !== undefined && (obj.userAddress = message.userAddress);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<PendingDevice>, I>>(object: I): PendingDevice {
    const message = createBasePendingDevice();
    message.deviceAddress = object.deviceAddress ?? "";
    message.userAddress = object.userAddress ?? "";
    return message;
  },
};

function createBaseDevice(): Device {
  return {
    deviceAddress: "",
    measurements: [],
    activePowerSum: 0,
    reversePowerSum: 0,
    usedActivePower: 0,
    fulfilledReversePower: 0,
  };
}

export const Device = {
  encode(message: Device, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.deviceAddress !== "") {
      writer.uint32(10).string(message.deviceAddress);
    }
    for (const v of message.measurements) {
      Measurement.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.activePowerSum !== 0) {
      writer.uint32(24).uint64(message.activePowerSum);
    }
    if (message.reversePowerSum !== 0) {
      writer.uint32(32).uint64(message.reversePowerSum);
    }
    if (message.usedActivePower !== 0) {
      writer.uint32(40).uint64(message.usedActivePower);
    }
    if (message.fulfilledReversePower !== 0) {
      writer.uint32(48).uint64(message.fulfilledReversePower);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Device {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDevice();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.deviceAddress = reader.string();
          break;
        case 2:
          message.measurements.push(Measurement.decode(reader, reader.uint32()));
          break;
        case 3:
          message.activePowerSum = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.reversePowerSum = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.usedActivePower = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.fulfilledReversePower = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Device {
    return {
      deviceAddress: isSet(object.deviceAddress) ? String(object.deviceAddress) : "",
      measurements: Array.isArray(object?.measurements)
        ? object.measurements.map((e: any) => Measurement.fromJSON(e))
        : [],
      activePowerSum: isSet(object.activePowerSum) ? Number(object.activePowerSum) : 0,
      reversePowerSum: isSet(object.reversePowerSum) ? Number(object.reversePowerSum) : 0,
      usedActivePower: isSet(object.usedActivePower) ? Number(object.usedActivePower) : 0,
      fulfilledReversePower: isSet(object.fulfilledReversePower) ? Number(object.fulfilledReversePower) : 0,
    };
  },

  toJSON(message: Device): unknown {
    const obj: any = {};
    message.deviceAddress !== undefined && (obj.deviceAddress = message.deviceAddress);
    if (message.measurements) {
      obj.measurements = message.measurements.map((e) => e ? Measurement.toJSON(e) : undefined);
    } else {
      obj.measurements = [];
    }
    message.activePowerSum !== undefined && (obj.activePowerSum = Math.round(message.activePowerSum));
    message.reversePowerSum !== undefined && (obj.reversePowerSum = Math.round(message.reversePowerSum));
    message.usedActivePower !== undefined && (obj.usedActivePower = Math.round(message.usedActivePower));
    message.fulfilledReversePower !== undefined
      && (obj.fulfilledReversePower = Math.round(message.fulfilledReversePower));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Device>, I>>(object: I): Device {
    const message = createBaseDevice();
    message.deviceAddress = object.deviceAddress ?? "";
    message.measurements = object.measurements?.map((e) => Measurement.fromPartial(e)) || [];
    message.activePowerSum = object.activePowerSum ?? 0;
    message.reversePowerSum = object.reversePowerSum ?? 0;
    message.usedActivePower = object.usedActivePower ?? 0;
    message.fulfilledReversePower = object.fulfilledReversePower ?? 0;
    return message;
  },
};

function createBaseMeasurement(): Measurement {
  return { id: 0, timestamp: undefined, activePower: 0, usedForCertificate: false, reversePower: 0, metadata: "" };
}

export const Measurement = {
  encode(message: Measurement, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.timestamp !== undefined) {
      Timestamp.encode(toTimestamp(message.timestamp), writer.uint32(18).fork()).ldelim();
    }
    if (message.activePower !== 0) {
      writer.uint32(24).uint64(message.activePower);
    }
    if (message.usedForCertificate === true) {
      writer.uint32(32).bool(message.usedForCertificate);
    }
    if (message.reversePower !== 0) {
      writer.uint32(40).uint64(message.reversePower);
    }
    if (message.metadata !== "") {
      writer.uint32(50).string(message.metadata);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Measurement {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMeasurement();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.timestamp = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 3:
          message.activePower = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.usedForCertificate = reader.bool();
          break;
        case 5:
          message.reversePower = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.metadata = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Measurement {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      timestamp: isSet(object.timestamp) ? fromJsonTimestamp(object.timestamp) : undefined,
      activePower: isSet(object.activePower) ? Number(object.activePower) : 0,
      usedForCertificate: isSet(object.usedForCertificate) ? Boolean(object.usedForCertificate) : false,
      reversePower: isSet(object.reversePower) ? Number(object.reversePower) : 0,
      metadata: isSet(object.metadata) ? String(object.metadata) : "",
    };
  },

  toJSON(message: Measurement): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.timestamp !== undefined && (obj.timestamp = message.timestamp.toISOString());
    message.activePower !== undefined && (obj.activePower = Math.round(message.activePower));
    message.usedForCertificate !== undefined && (obj.usedForCertificate = message.usedForCertificate);
    message.reversePower !== undefined && (obj.reversePower = Math.round(message.reversePower));
    message.metadata !== undefined && (obj.metadata = message.metadata);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Measurement>, I>>(object: I): Measurement {
    const message = createBaseMeasurement();
    message.id = object.id ?? 0;
    message.timestamp = object.timestamp ?? undefined;
    message.activePower = object.activePower ?? 0;
    message.usedForCertificate = object.usedForCertificate ?? false;
    message.reversePower = object.reversePower ?? 0;
    message.metadata = object.metadata ?? "";
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
