/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "pendulum-labs.market.market";

export interface MsgCreatePool {
  creator: string;
  coinA: string;
  coinB: string;
  rateA: string[];
  rateB: string[];
}

export interface MsgCreatePoolResponse {}

export interface MsgCreateDrop {
  creator: string;
  pair: string;
  drops: string;
  rate1: string[];
  prev1: string;
  next1: string;
  rate2: string[];
  prev2: string;
  next2: string;
}

export interface MsgCreateDropResponse {}

export interface MsgRedeemDrop {
  creator: string;
  uid: string;
}

export interface MsgRedeemDropResponse {}

export interface MsgCreateOrder {
  creator: string;
  denomAsk: string;
  denomBid: string;
  orderType: string;
  amount: string;
  rate: string[];
  prev: string;
  next: string;
}

export interface MsgCreateOrderResponse {}

export interface MsgCancelOrder {
  creator: string;
  uid: string;
}

export interface MsgCancelOrderResponse {}

export interface MsgMarketOrder {
  creator: string;
  denomAsk: string;
  denomBid: string;
  amountBid: string;
  quoteAsk: string;
  slippage: string;
}

export interface MsgMarketOrderResponse {}

const baseMsgCreatePool: object = {
  creator: "",
  coinA: "",
  coinB: "",
  rateA: "",
  rateB: "",
};

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
    for (const v of message.rateA) {
      writer.uint32(34).string(v!);
    }
    for (const v of message.rateB) {
      writer.uint32(42).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreatePool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreatePool } as MsgCreatePool;
    message.rateA = [];
    message.rateB = [];
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
        case 4:
          message.rateA.push(reader.string());
          break;
        case 5:
          message.rateB.push(reader.string());
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
    message.rateA = [];
    message.rateB = [];
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
    if (object.rateA !== undefined && object.rateA !== null) {
      for (const e of object.rateA) {
        message.rateA.push(String(e));
      }
    }
    if (object.rateB !== undefined && object.rateB !== null) {
      for (const e of object.rateB) {
        message.rateB.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: MsgCreatePool): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.coinA !== undefined && (obj.coinA = message.coinA);
    message.coinB !== undefined && (obj.coinB = message.coinB);
    if (message.rateA) {
      obj.rateA = message.rateA.map((e) => e);
    } else {
      obj.rateA = [];
    }
    if (message.rateB) {
      obj.rateB = message.rateB.map((e) => e);
    } else {
      obj.rateB = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreatePool>): MsgCreatePool {
    const message = { ...baseMsgCreatePool } as MsgCreatePool;
    message.rateA = [];
    message.rateB = [];
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
    if (object.rateA !== undefined && object.rateA !== null) {
      for (const e of object.rateA) {
        message.rateA.push(e);
      }
    }
    if (object.rateB !== undefined && object.rateB !== null) {
      for (const e of object.rateB) {
        message.rateB.push(e);
      }
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

const baseMsgCreateDrop: object = {
  creator: "",
  pair: "",
  drops: "",
  rate1: "",
  prev1: "",
  next1: "",
  rate2: "",
  prev2: "",
  next2: "",
};

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
    for (const v of message.rate1) {
      writer.uint32(34).string(v!);
    }
    if (message.prev1 !== "") {
      writer.uint32(42).string(message.prev1);
    }
    if (message.next1 !== "") {
      writer.uint32(50).string(message.next1);
    }
    for (const v of message.rate2) {
      writer.uint32(58).string(v!);
    }
    if (message.prev2 !== "") {
      writer.uint32(66).string(message.prev2);
    }
    if (message.next2 !== "") {
      writer.uint32(74).string(message.next2);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateDrop {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateDrop } as MsgCreateDrop;
    message.rate1 = [];
    message.rate2 = [];
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
        case 4:
          message.rate1.push(reader.string());
          break;
        case 5:
          message.prev1 = reader.string();
          break;
        case 6:
          message.next1 = reader.string();
          break;
        case 7:
          message.rate2.push(reader.string());
          break;
        case 8:
          message.prev2 = reader.string();
          break;
        case 9:
          message.next2 = reader.string();
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
    message.rate1 = [];
    message.rate2 = [];
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
    if (object.rate1 !== undefined && object.rate1 !== null) {
      for (const e of object.rate1) {
        message.rate1.push(String(e));
      }
    }
    if (object.prev1 !== undefined && object.prev1 !== null) {
      message.prev1 = String(object.prev1);
    } else {
      message.prev1 = "";
    }
    if (object.next1 !== undefined && object.next1 !== null) {
      message.next1 = String(object.next1);
    } else {
      message.next1 = "";
    }
    if (object.rate2 !== undefined && object.rate2 !== null) {
      for (const e of object.rate2) {
        message.rate2.push(String(e));
      }
    }
    if (object.prev2 !== undefined && object.prev2 !== null) {
      message.prev2 = String(object.prev2);
    } else {
      message.prev2 = "";
    }
    if (object.next2 !== undefined && object.next2 !== null) {
      message.next2 = String(object.next2);
    } else {
      message.next2 = "";
    }
    return message;
  },

  toJSON(message: MsgCreateDrop): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.pair !== undefined && (obj.pair = message.pair);
    message.drops !== undefined && (obj.drops = message.drops);
    if (message.rate1) {
      obj.rate1 = message.rate1.map((e) => e);
    } else {
      obj.rate1 = [];
    }
    message.prev1 !== undefined && (obj.prev1 = message.prev1);
    message.next1 !== undefined && (obj.next1 = message.next1);
    if (message.rate2) {
      obj.rate2 = message.rate2.map((e) => e);
    } else {
      obj.rate2 = [];
    }
    message.prev2 !== undefined && (obj.prev2 = message.prev2);
    message.next2 !== undefined && (obj.next2 = message.next2);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateDrop>): MsgCreateDrop {
    const message = { ...baseMsgCreateDrop } as MsgCreateDrop;
    message.rate1 = [];
    message.rate2 = [];
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
    if (object.rate1 !== undefined && object.rate1 !== null) {
      for (const e of object.rate1) {
        message.rate1.push(e);
      }
    }
    if (object.prev1 !== undefined && object.prev1 !== null) {
      message.prev1 = object.prev1;
    } else {
      message.prev1 = "";
    }
    if (object.next1 !== undefined && object.next1 !== null) {
      message.next1 = object.next1;
    } else {
      message.next1 = "";
    }
    if (object.rate2 !== undefined && object.rate2 !== null) {
      for (const e of object.rate2) {
        message.rate2.push(e);
      }
    }
    if (object.prev2 !== undefined && object.prev2 !== null) {
      message.prev2 = object.prev2;
    } else {
      message.prev2 = "";
    }
    if (object.next2 !== undefined && object.next2 !== null) {
      message.next2 = object.next2;
    } else {
      message.next2 = "";
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

const baseMsgCreateOrder: object = {
  creator: "",
  denomAsk: "",
  denomBid: "",
  orderType: "",
  amount: "",
  rate: "",
  prev: "",
  next: "",
};

export const MsgCreateOrder = {
  encode(message: MsgCreateOrder, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.denomAsk !== "") {
      writer.uint32(18).string(message.denomAsk);
    }
    if (message.denomBid !== "") {
      writer.uint32(26).string(message.denomBid);
    }
    if (message.orderType !== "") {
      writer.uint32(34).string(message.orderType);
    }
    if (message.amount !== "") {
      writer.uint32(42).string(message.amount);
    }
    for (const v of message.rate) {
      writer.uint32(50).string(v!);
    }
    if (message.prev !== "") {
      writer.uint32(58).string(message.prev);
    }
    if (message.next !== "") {
      writer.uint32(66).string(message.next);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateOrder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateOrder } as MsgCreateOrder;
    message.rate = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.denomAsk = reader.string();
          break;
        case 3:
          message.denomBid = reader.string();
          break;
        case 4:
          message.orderType = reader.string();
          break;
        case 5:
          message.amount = reader.string();
          break;
        case 6:
          message.rate.push(reader.string());
          break;
        case 7:
          message.prev = reader.string();
          break;
        case 8:
          message.next = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateOrder {
    const message = { ...baseMsgCreateOrder } as MsgCreateOrder;
    message.rate = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.denomAsk !== undefined && object.denomAsk !== null) {
      message.denomAsk = String(object.denomAsk);
    } else {
      message.denomAsk = "";
    }
    if (object.denomBid !== undefined && object.denomBid !== null) {
      message.denomBid = String(object.denomBid);
    } else {
      message.denomBid = "";
    }
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = String(object.orderType);
    } else {
      message.orderType = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.rate !== undefined && object.rate !== null) {
      for (const e of object.rate) {
        message.rate.push(String(e));
      }
    }
    if (object.prev !== undefined && object.prev !== null) {
      message.prev = String(object.prev);
    } else {
      message.prev = "";
    }
    if (object.next !== undefined && object.next !== null) {
      message.next = String(object.next);
    } else {
      message.next = "";
    }
    return message;
  },

  toJSON(message: MsgCreateOrder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.denomAsk !== undefined && (obj.denomAsk = message.denomAsk);
    message.denomBid !== undefined && (obj.denomBid = message.denomBid);
    message.orderType !== undefined && (obj.orderType = message.orderType);
    message.amount !== undefined && (obj.amount = message.amount);
    if (message.rate) {
      obj.rate = message.rate.map((e) => e);
    } else {
      obj.rate = [];
    }
    message.prev !== undefined && (obj.prev = message.prev);
    message.next !== undefined && (obj.next = message.next);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateOrder>): MsgCreateOrder {
    const message = { ...baseMsgCreateOrder } as MsgCreateOrder;
    message.rate = [];
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.denomAsk !== undefined && object.denomAsk !== null) {
      message.denomAsk = object.denomAsk;
    } else {
      message.denomAsk = "";
    }
    if (object.denomBid !== undefined && object.denomBid !== null) {
      message.denomBid = object.denomBid;
    } else {
      message.denomBid = "";
    }
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = object.orderType;
    } else {
      message.orderType = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.rate !== undefined && object.rate !== null) {
      for (const e of object.rate) {
        message.rate.push(e);
      }
    }
    if (object.prev !== undefined && object.prev !== null) {
      message.prev = object.prev;
    } else {
      message.prev = "";
    }
    if (object.next !== undefined && object.next !== null) {
      message.next = object.next;
    } else {
      message.next = "";
    }
    return message;
  },
};

const baseMsgCreateOrderResponse: object = {};

export const MsgCreateOrderResponse = {
  encode(_: MsgCreateOrderResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateOrderResponse } as MsgCreateOrderResponse;
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

  fromJSON(_: any): MsgCreateOrderResponse {
    const message = { ...baseMsgCreateOrderResponse } as MsgCreateOrderResponse;
    return message;
  },

  toJSON(_: MsgCreateOrderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgCreateOrderResponse>): MsgCreateOrderResponse {
    const message = { ...baseMsgCreateOrderResponse } as MsgCreateOrderResponse;
    return message;
  },
};

const baseMsgCancelOrder: object = { creator: "", uid: "" };

export const MsgCancelOrder = {
  encode(message: MsgCancelOrder, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.uid !== "") {
      writer.uint32(18).string(message.uid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCancelOrder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCancelOrder } as MsgCancelOrder;
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

  fromJSON(object: any): MsgCancelOrder {
    const message = { ...baseMsgCancelOrder } as MsgCancelOrder;
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

  toJSON(message: MsgCancelOrder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.uid !== undefined && (obj.uid = message.uid);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCancelOrder>): MsgCancelOrder {
    const message = { ...baseMsgCancelOrder } as MsgCancelOrder;
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

const baseMsgCancelOrderResponse: object = {};

export const MsgCancelOrderResponse = {
  encode(_: MsgCancelOrderResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCancelOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCancelOrderResponse } as MsgCancelOrderResponse;
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

  fromJSON(_: any): MsgCancelOrderResponse {
    const message = { ...baseMsgCancelOrderResponse } as MsgCancelOrderResponse;
    return message;
  },

  toJSON(_: MsgCancelOrderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgCancelOrderResponse>): MsgCancelOrderResponse {
    const message = { ...baseMsgCancelOrderResponse } as MsgCancelOrderResponse;
    return message;
  },
};

const baseMsgMarketOrder: object = {
  creator: "",
  denomAsk: "",
  denomBid: "",
  amountBid: "",
  quoteAsk: "",
  slippage: "",
};

export const MsgMarketOrder = {
  encode(message: MsgMarketOrder, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.denomAsk !== "") {
      writer.uint32(18).string(message.denomAsk);
    }
    if (message.denomBid !== "") {
      writer.uint32(26).string(message.denomBid);
    }
    if (message.amountBid !== "") {
      writer.uint32(34).string(message.amountBid);
    }
    if (message.quoteAsk !== "") {
      writer.uint32(42).string(message.quoteAsk);
    }
    if (message.slippage !== "") {
      writer.uint32(50).string(message.slippage);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMarketOrder {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgMarketOrder } as MsgMarketOrder;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.denomAsk = reader.string();
          break;
        case 3:
          message.denomBid = reader.string();
          break;
        case 4:
          message.amountBid = reader.string();
          break;
        case 5:
          message.quoteAsk = reader.string();
          break;
        case 6:
          message.slippage = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgMarketOrder {
    const message = { ...baseMsgMarketOrder } as MsgMarketOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.denomAsk !== undefined && object.denomAsk !== null) {
      message.denomAsk = String(object.denomAsk);
    } else {
      message.denomAsk = "";
    }
    if (object.denomBid !== undefined && object.denomBid !== null) {
      message.denomBid = String(object.denomBid);
    } else {
      message.denomBid = "";
    }
    if (object.amountBid !== undefined && object.amountBid !== null) {
      message.amountBid = String(object.amountBid);
    } else {
      message.amountBid = "";
    }
    if (object.quoteAsk !== undefined && object.quoteAsk !== null) {
      message.quoteAsk = String(object.quoteAsk);
    } else {
      message.quoteAsk = "";
    }
    if (object.slippage !== undefined && object.slippage !== null) {
      message.slippage = String(object.slippage);
    } else {
      message.slippage = "";
    }
    return message;
  },

  toJSON(message: MsgMarketOrder): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.denomAsk !== undefined && (obj.denomAsk = message.denomAsk);
    message.denomBid !== undefined && (obj.denomBid = message.denomBid);
    message.amountBid !== undefined && (obj.amountBid = message.amountBid);
    message.quoteAsk !== undefined && (obj.quoteAsk = message.quoteAsk);
    message.slippage !== undefined && (obj.slippage = message.slippage);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgMarketOrder>): MsgMarketOrder {
    const message = { ...baseMsgMarketOrder } as MsgMarketOrder;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.denomAsk !== undefined && object.denomAsk !== null) {
      message.denomAsk = object.denomAsk;
    } else {
      message.denomAsk = "";
    }
    if (object.denomBid !== undefined && object.denomBid !== null) {
      message.denomBid = object.denomBid;
    } else {
      message.denomBid = "";
    }
    if (object.amountBid !== undefined && object.amountBid !== null) {
      message.amountBid = object.amountBid;
    } else {
      message.amountBid = "";
    }
    if (object.quoteAsk !== undefined && object.quoteAsk !== null) {
      message.quoteAsk = object.quoteAsk;
    } else {
      message.quoteAsk = "";
    }
    if (object.slippage !== undefined && object.slippage !== null) {
      message.slippage = object.slippage;
    } else {
      message.slippage = "";
    }
    return message;
  },
};

const baseMsgMarketOrderResponse: object = {};

export const MsgMarketOrderResponse = {
  encode(_: MsgMarketOrderResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMarketOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgMarketOrderResponse } as MsgMarketOrderResponse;
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

  fromJSON(_: any): MsgMarketOrderResponse {
    const message = { ...baseMsgMarketOrderResponse } as MsgMarketOrderResponse;
    return message;
  },

  toJSON(_: MsgMarketOrderResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgMarketOrderResponse>): MsgMarketOrderResponse {
    const message = { ...baseMsgMarketOrderResponse } as MsgMarketOrderResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreatePool(request: MsgCreatePool): Promise<MsgCreatePoolResponse>;
  CreateDrop(request: MsgCreateDrop): Promise<MsgCreateDropResponse>;
  RedeemDrop(request: MsgRedeemDrop): Promise<MsgRedeemDropResponse>;
  CreateOrder(request: MsgCreateOrder): Promise<MsgCreateOrderResponse>;
  CancelOrder(request: MsgCancelOrder): Promise<MsgCancelOrderResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  MarketOrder(request: MsgMarketOrder): Promise<MsgMarketOrderResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreatePool(request: MsgCreatePool): Promise<MsgCreatePoolResponse> {
    const data = MsgCreatePool.encode(request).finish();
    const promise = this.rpc.request(
      "pendulum-labs.market.market.Msg",
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
      "pendulum-labs.market.market.Msg",
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
      "pendulum-labs.market.market.Msg",
      "RedeemDrop",
      data
    );
    return promise.then((data) =>
      MsgRedeemDropResponse.decode(new Reader(data))
    );
  }

  CreateOrder(request: MsgCreateOrder): Promise<MsgCreateOrderResponse> {
    const data = MsgCreateOrder.encode(request).finish();
    const promise = this.rpc.request(
      "pendulum-labs.market.market.Msg",
      "CreateOrder",
      data
    );
    return promise.then((data) =>
      MsgCreateOrderResponse.decode(new Reader(data))
    );
  }

  CancelOrder(request: MsgCancelOrder): Promise<MsgCancelOrderResponse> {
    const data = MsgCancelOrder.encode(request).finish();
    const promise = this.rpc.request(
      "pendulum-labs.market.market.Msg",
      "CancelOrder",
      data
    );
    return promise.then((data) =>
      MsgCancelOrderResponse.decode(new Reader(data))
    );
  }

  MarketOrder(request: MsgMarketOrder): Promise<MsgMarketOrderResponse> {
    const data = MsgMarketOrder.encode(request).finish();
    const promise = this.rpc.request(
      "pendulum-labs.market.market.Msg",
      "MarketOrder",
      data
    );
    return promise.then((data) =>
      MsgMarketOrderResponse.decode(new Reader(data))
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
