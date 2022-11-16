/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfeairdrop";

export interface MsgClaim {
  cclaimer: string;
  campaignId: string;
  missionId: string;
}

export interface MsgClaimResponse {}

const baseMsgClaim: object = { cclaimer: "", campaignId: "", missionId: "" };

export const MsgClaim = {
  encode(message: MsgClaim, writer: Writer = Writer.create()): Writer {
    if (message.cclaimer !== "") {
      writer.uint32(10).string(message.cclaimer);
    }
    if (message.campaignId !== "") {
      writer.uint32(18).string(message.campaignId);
    }
    if (message.missionId !== "") {
      writer.uint32(26).string(message.missionId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgClaim {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgClaim } as MsgClaim;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.cclaimer = reader.string();
          break;
        case 2:
          message.campaignId = reader.string();
          break;
        case 3:
          message.missionId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgClaim {
    const message = { ...baseMsgClaim } as MsgClaim;
    if (object.cclaimer !== undefined && object.cclaimer !== null) {
      message.cclaimer = String(object.cclaimer);
    } else {
      message.cclaimer = "";
    }
    if (object.campaignId !== undefined && object.campaignId !== null) {
      message.campaignId = String(object.campaignId);
    } else {
      message.campaignId = "";
    }
    if (object.missionId !== undefined && object.missionId !== null) {
      message.missionId = String(object.missionId);
    } else {
      message.missionId = "";
    }
    return message;
  },

  toJSON(message: MsgClaim): unknown {
    const obj: any = {};
    message.cclaimer !== undefined && (obj.cclaimer = message.cclaimer);
    message.campaignId !== undefined && (obj.campaignId = message.campaignId);
    message.missionId !== undefined && (obj.missionId = message.missionId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgClaim>): MsgClaim {
    const message = { ...baseMsgClaim } as MsgClaim;
    if (object.cclaimer !== undefined && object.cclaimer !== null) {
      message.cclaimer = object.cclaimer;
    } else {
      message.cclaimer = "";
    }
    if (object.campaignId !== undefined && object.campaignId !== null) {
      message.campaignId = object.campaignId;
    } else {
      message.campaignId = "";
    }
    if (object.missionId !== undefined && object.missionId !== null) {
      message.missionId = object.missionId;
    } else {
      message.missionId = "";
    }
    return message;
  },
};

const baseMsgClaimResponse: object = {};

export const MsgClaimResponse = {
  encode(_: MsgClaimResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgClaimResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgClaimResponse } as MsgClaimResponse;
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

  fromJSON(_: any): MsgClaimResponse {
    const message = { ...baseMsgClaimResponse } as MsgClaimResponse;
    return message;
  },

  toJSON(_: MsgClaimResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgClaimResponse>): MsgClaimResponse {
    const message = { ...baseMsgClaimResponse } as MsgClaimResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Claim(request: MsgClaim): Promise<MsgClaimResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Claim(request: MsgClaim): Promise<MsgClaimResponse> {
    const data = MsgClaim.encode(request).finish();
    const promise = this.rpc.request(
      "chain4energy.c4echain.cfeairdrop.Msg",
      "Claim",
      data
    );
    return promise.then((data) => MsgClaimResponse.decode(new Reader(data)));
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
