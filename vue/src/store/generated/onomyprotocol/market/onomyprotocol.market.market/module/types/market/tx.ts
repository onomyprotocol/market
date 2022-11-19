/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "onomyprotocol.market.market";

export interface MsgCreatePool {
  creator: string;
  coinA: string;
  coinB: string;
}

export interface MsgCreatePoolResponse {}

export interface MsgCreateDrop {
  creator: string;
  pair: string;
  drops: string;
}

export interface MsgCreateDropResponse {}

export interface MsgRedeemDrop {
  creator: string;
  uid: string;
}

export interface MsgRedeemDropResponse {}

const baseMsgCreatePool: object = { creator: "", coinA: "", coinB: "" };

export const MsgCreatePool = {
  encode(message: MsgCreatePool, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.coinA !== "") {
      writer.uint32(18).string(message.coinA);
    }
    if (message.coinB !== "") {
      writer.uint32(26).string(message.coinB);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreatePool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreatePool } as MsgCreatePool;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.coinA = reader.string();
          break;
        case 3:
          message.coinB = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreatePool {
    const message = { ...baseMsgCreatePool } as MsgCreatePool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.coinA !== undefined && object.coinA !== null) {
      message.coinA = String(object.coinA);
    } else {
      message.coinA = "";
    }
    if (object.coinB !== undefined && object.coinB !== null) {
      message.coinB = String(object.coinB);
    } else {
      message.coinB = "";
    }
    return message;
  },

  toJSON(message: MsgCreatePool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.coinA !== undefined && (obj.coinA = message.coinA);
    message.coinB !== undefined && (obj.coinB = message.coinB);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreatePool>): MsgCreatePool {
    const message = { ...baseMsgCreatePool } as MsgCreatePool;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.coinA !== undefined && object.coinA !== null) {
      message.coinA = object.coinA;
    } else {
      message.coinA = "";
    }
    if (object.coinB !== undefined && object.coinB !== null) {
      message.coinB = object.coinB;
    } else {
      message.coinB = "";
    }
    return message;
  },
};

const baseMsgCreatePoolResponse: object = {};

export const MsgCreatePoolResponse = {
  encode(_: MsgCreatePoolResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreatePoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreatePoolResponse } as MsgCreatePoolResponse;
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

  fromJSON(_: any): MsgCreatePoolResponse {
    const message = { ...baseMsgCreatePoolResponse } as MsgCreatePoolResponse;
    return message;
  },

  toJSON(_: MsgCreatePoolResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgCreatePoolResponse>): MsgCreatePoolResponse {
    const message = { ...baseMsgCreatePoolResponse } as MsgCreatePoolResponse;
    return message;
  },
};

const baseMsgCreateDrop: object = { creator: "", pair: "", drops: "" };

export const MsgCreateDrop = {
  encode(message: MsgCreateDrop, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.pair !== "") {
      writer.uint32(18).string(message.pair);
    }
    if (message.drops !== "") {
      writer.uint32(26).string(message.drops);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateDrop {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateDrop } as MsgCreateDrop;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.pair = reader.string();
          break;
        case 3:
          message.drops = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateDrop {
    const message = { ...baseMsgCreateDrop } as MsgCreateDrop;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.pair !== undefined && object.pair !== null) {
      message.pair = String(object.pair);
    } else {
      message.pair = "";
    }
    if (object.drops !== undefined && object.drops !== null) {
      message.drops = String(object.drops);
    } else {
      message.drops = "";
    }
    return message;
  },

  toJSON(message: MsgCreateDrop): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.pair !== undefined && (obj.pair = message.pair);
    message.drops !== undefined && (obj.drops = message.drops);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateDrop>): MsgCreateDrop {
    const message = { ...baseMsgCreateDrop } as MsgCreateDrop;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.pair !== undefined && object.pair !== null) {
      message.pair = object.pair;
    } else {
      message.pair = "";
    }
    if (object.drops !== undefined && object.drops !== null) {
      message.drops = object.drops;
    } else {
      message.drops = "";
    }
    return message;
  },
};

const baseMsgCreateDropResponse: object = {};

export const MsgCreateDropResponse = {
  encode(_: MsgCreateDropResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateDropResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateDropResponse } as MsgCreateDropResponse;
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

  fromJSON(_: any): MsgCreateDropResponse {
    const message = { ...baseMsgCreateDropResponse } as MsgCreateDropResponse;
    return message;
  },

  toJSON(_: MsgCreateDropResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgCreateDropResponse>): MsgCreateDropResponse {
    const message = { ...baseMsgCreateDropResponse } as MsgCreateDropResponse;
    return message;
  },
};

const baseMsgRedeemDrop: object = { creator: "", uid: "" };

export const MsgRedeemDrop = {
  encode(message: MsgRedeemDrop, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.uid !== "") {
      writer.uint32(18).string(message.uid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRedeemDrop {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRedeemDrop } as MsgRedeemDrop;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.uid = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRedeemDrop {
    const message = { ...baseMsgRedeemDrop } as MsgRedeemDrop;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = String(object.uid);
    } else {
      message.uid = "";
    }
    return message;
  },

  toJSON(message: MsgRedeemDrop): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.uid !== undefined && (obj.uid = message.uid);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRedeemDrop>): MsgRedeemDrop {
    const message = { ...baseMsgRedeemDrop } as MsgRedeemDrop;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = "";
    }
    return message;
  },
};

const baseMsgRedeemDropResponse: object = {};

export const MsgRedeemDropResponse = {
  encode(_: MsgRedeemDropResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRedeemDropResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRedeemDropResponse } as MsgRedeemDropResponse;
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

  fromJSON(_: any): MsgRedeemDropResponse {
    const message = { ...baseMsgRedeemDropResponse } as MsgRedeemDropResponse;
    return message;
  },

  toJSON(_: MsgRedeemDropResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgRedeemDropResponse>): MsgRedeemDropResponse {
    const message = { ...baseMsgRedeemDropResponse } as MsgRedeemDropResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreatePool(request: MsgCreatePool): Promise<MsgCreatePoolResponse>;
  CreateDrop(request: MsgCreateDrop): Promise<MsgCreateDropResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RedeemDrop(request: MsgRedeemDrop): Promise<MsgRedeemDropResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreatePool(request: MsgCreatePool): Promise<MsgCreatePoolResponse> {
    const data = MsgCreatePool.encode(request).finish();
    const promise = this.rpc.request(
      "onomyprotocol.market.market.Msg",
      "CreatePool",
      data
    );
    return promise.then((data) =>
      MsgCreatePoolResponse.decode(new Reader(data))
    );
  }

  CreateDrop(request: MsgCreateDrop): Promise<MsgCreateDropResponse> {
    const data = MsgCreateDrop.encode(request).finish();
    const promise = this.rpc.request(
      "onomyprotocol.market.market.Msg",
      "CreateDrop",
      data
    );
    return promise.then((data) =>
      MsgCreateDropResponse.decode(new Reader(data))
    );
  }

  RedeemDrop(request: MsgRedeemDrop): Promise<MsgRedeemDropResponse> {
    const data = MsgRedeemDrop.encode(request).finish();
    const promise = this.rpc.request(
      "onomyprotocol.market.market.Msg",
      "RedeemDrop",
      data
    );
    return promise.then((data) =>
      MsgRedeemDropResponse.decode(new Reader(data))
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
