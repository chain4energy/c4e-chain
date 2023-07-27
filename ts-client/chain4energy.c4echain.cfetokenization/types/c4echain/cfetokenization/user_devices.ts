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
}

export interface PendingDevice {
  deviceAddress: string;
  userAddress: string;
}

export interface Device {
  deviceAddress: string;
  measurements: Measurement[];
  powerSum: number;
  usedPower: number;
}

export interface Measurement {
  timestamp: Date | undefined;
  power: number;
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
  return { deviceAddress: "", name: "" };
}

export const UserDevice = {
  encode(message: UserDevice, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.deviceAddress !== "") {
      writer.uint32(10).string(message.deviceAddress);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
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
    };
  },

  toJSON(message: UserDevice): unknown {
    const obj: any = {};
    message.deviceAddress !== undefined && (obj.deviceAddress = message.deviceAddress);
    message.name !== undefined && (obj.name = message.name);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<UserDevice>, I>>(object: I): UserDevice {
    const message = createBaseUserDevice();
    message.deviceAddress = object.deviceAddress ?? "";
    message.name = object.name ?? "";
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
  return { deviceAddress: "", measurements: [], powerSum: 0, usedPower: 0 };
}

export const Device = {
  encode(message: Device, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.deviceAddress !== "") {
      writer.uint32(10).string(message.deviceAddress);
    }
    for (const v of message.measurements) {
      Measurement.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.powerSum !== 0) {
      writer.uint32(24).uint64(message.powerSum);
    }
    if (message.usedPower !== 0) {
      writer.uint32(32).uint64(message.usedPower);
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
          message.powerSum = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.usedPower = longToNumber(reader.uint64() as Long);
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
      powerSum: isSet(object.powerSum) ? Number(object.powerSum) : 0,
      usedPower: isSet(object.usedPower) ? Number(object.usedPower) : 0,
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
    message.powerSum !== undefined && (obj.powerSum = Math.round(message.powerSum));
    message.usedPower !== undefined && (obj.usedPower = Math.round(message.usedPower));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Device>, I>>(object: I): Device {
    const message = createBaseDevice();
    message.deviceAddress = object.deviceAddress ?? "";
    message.measurements = object.measurements?.map((e) => Measurement.fromPartial(e)) || [];
    message.powerSum = object.powerSum ?? 0;
    message.usedPower = object.usedPower ?? 0;
    return message;
  },
};

function createBaseMeasurement(): Measurement {
  return { timestamp: undefined, power: 0 };
}

export const Measurement = {
  encode(message: Measurement, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.timestamp !== undefined) {
      Timestamp.encode(toTimestamp(message.timestamp), writer.uint32(10).fork()).ldelim();
    }
    if (message.power !== 0) {
      writer.uint32(16).uint64(message.power);
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
          message.timestamp = fromTimestamp(Timestamp.decode(reader, reader.uint32()));
          break;
        case 2:
          message.power = longToNumber(reader.uint64() as Long);
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
      timestamp: isSet(object.timestamp) ? fromJsonTimestamp(object.timestamp) : undefined,
      power: isSet(object.power) ? Number(object.power) : 0,
    };
  },

  toJSON(message: Measurement): unknown {
    const obj: any = {};
    message.timestamp !== undefined && (obj.timestamp = message.timestamp.toISOString());
    message.power !== undefined && (obj.power = Math.round(message.power));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Measurement>, I>>(object: I): Measurement {
    const message = createBaseMeasurement();
    message.timestamp = object.timestamp ?? undefined;
    message.power = object.power ?? 0;
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
