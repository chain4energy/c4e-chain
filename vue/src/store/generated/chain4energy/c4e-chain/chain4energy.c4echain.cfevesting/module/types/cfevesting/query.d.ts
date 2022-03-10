import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../cfevesting/params";
import { VestingTypes } from "../cfevesting/vesting_types";
import { Coin } from "../cosmos/base/v1beta1/coin";
export declare const protobufPackage = "chain4energy.c4echain.cfevesting";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QueryVestingTypeRequest {
}
export interface QueryVestingTypeResponse {
    vestingTypes: VestingTypes | undefined;
}
export interface QueryVestingRequest {
    address: string;
}
export interface QueryVestingResponse {
    delegableAddress: string;
    vestings: VestingInfo[];
}
export interface VestingInfo {
    vestingType: string;
    vestingStartHeight: number;
    lockEndHeight: number;
    vestingEndHeight: number;
    withdrawable: string;
    delegationAllowed: boolean;
    vested: Coin | undefined;
    currentVestedAmount: string;
}
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
};
export declare const QueryVestingTypeRequest: {
    encode(_: QueryVestingTypeRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryVestingTypeRequest;
    fromJSON(_: any): QueryVestingTypeRequest;
    toJSON(_: QueryVestingTypeRequest): unknown;
    fromPartial(_: DeepPartial<QueryVestingTypeRequest>): QueryVestingTypeRequest;
};
export declare const QueryVestingTypeResponse: {
    encode(message: QueryVestingTypeResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryVestingTypeResponse;
    fromJSON(object: any): QueryVestingTypeResponse;
    toJSON(message: QueryVestingTypeResponse): unknown;
    fromPartial(object: DeepPartial<QueryVestingTypeResponse>): QueryVestingTypeResponse;
};
export declare const QueryVestingRequest: {
    encode(message: QueryVestingRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryVestingRequest;
    fromJSON(object: any): QueryVestingRequest;
    toJSON(message: QueryVestingRequest): unknown;
    fromPartial(object: DeepPartial<QueryVestingRequest>): QueryVestingRequest;
};
export declare const QueryVestingResponse: {
    encode(message: QueryVestingResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryVestingResponse;
    fromJSON(object: any): QueryVestingResponse;
    toJSON(message: QueryVestingResponse): unknown;
    fromPartial(object: DeepPartial<QueryVestingResponse>): QueryVestingResponse;
};
export declare const VestingInfo: {
    encode(message: VestingInfo, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): VestingInfo;
    fromJSON(object: any): VestingInfo;
    toJSON(message: VestingInfo): unknown;
    fromPartial(object: DeepPartial<VestingInfo>): VestingInfo;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a list of VestingType items. */
    VestingType(request: QueryVestingTypeRequest): Promise<QueryVestingTypeResponse>;
    /** Queries a list of Vesting items. */
    Vesting(request: QueryVestingRequest): Promise<QueryVestingResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    VestingType(request: QueryVestingTypeRequest): Promise<QueryVestingTypeResponse>;
    Vesting(request: QueryVestingRequest): Promise<QueryVestingResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
