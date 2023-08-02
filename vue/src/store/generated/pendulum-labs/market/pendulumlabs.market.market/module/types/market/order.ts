/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "pendulumlabs.market.market";

export interface Order {
  uid: number;
  owner: string;
  status: string;
  orderType: string;
  denomAsk: string;
  denomBid: string;
  amount: string;
  rate: string[];
  prev: number;
  next: number;
  beg_time: number;
  end_time: number;
}

export interface Orders {
  uids: number[];
}

export interface OrderResponse {
  uid: number;
  owner: string;
  status: string;
  orderType: string;
  denomAsk: string;
  denomBid: string;
  amount: string;
  rate: string[];
  prev: number;
  next: number;
  beg_time: number;
  end_time: number;
}

const baseOrder: object = {
  uid: 0,
  owner: "",
  status: "",
  orderType: "",
  denomAsk: "",
  denomBid: "",
  amount: "",
  rate: "",
  prev: 0,
  next: 0,
  beg_time: 0,
  end_time: 0,
};

export const Order = {
  encode(message: Order, writer: Writer = Writer.create()): Writer {
    if (message.uid !== 0) {
      writer.uint32(8).uint64(message.uid);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    if (message.status !== "") {
      writer.uint32(26).string(message.status);
    }
    if (message.orderType !== "") {
      writer.uint32(34).string(message.orderType);
    }
    if (message.denomAsk !== "") {
      writer.uint32(42).string(message.denomAsk);
    }
    if (message.denomBid !== "") {
      writer.uint32(50).string(message.denomBid);
    }
    if (message.amount !== "") {
      writer.uint32(58).string(message.amount);
    }
    for (const v of message.rate) {
      writer.uint32(66).string(v!);
    }
    if (message.prev !== 0) {
      writer.uint32(72).uint64(message.prev);
    }
    if (message.next !== 0) {
      writer.uint32(80).uint64(message.next);
    }
    if (message.beg_time !== 0) {
      writer.uint32(88).uint64(message.beg_time);
    }
    if (message.end_time !== 0) {
      writer.uint32(96).uint64(message.end_time);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Order {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOrder } as Order;
    message.rate = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uid = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.owner = reader.string();
          break;
        case 3:
          message.status = reader.string();
          break;
        case 4:
          message.orderType = reader.string();
          break;
        case 5:
          message.denomAsk = reader.string();
          break;
        case 6:
          message.denomBid = reader.string();
          break;
        case 7:
          message.amount = reader.string();
          break;
        case 8:
          message.rate.push(reader.string());
          break;
        case 9:
          message.prev = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          message.next = longToNumber(reader.uint64() as Long);
          break;
        case 11:
          message.beg_time = longToNumber(reader.uint64() as Long);
          break;
        case 12:
          message.end_time = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Order {
    const message = { ...baseOrder } as Order;
    message.rate = [];
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = Number(object.uid);
    } else {
      message.uid = 0;
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = String(object.orderType);
    } else {
      message.orderType = "";
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
      message.prev = Number(object.prev);
    } else {
      message.prev = 0;
    }
    if (object.next !== undefined && object.next !== null) {
      message.next = Number(object.next);
    } else {
      message.next = 0;
    }
    if (object.beg_time !== undefined && object.beg_time !== null) {
      message.beg_time = Number(object.beg_time);
    } else {
      message.beg_time = 0;
    }
    if (object.end_time !== undefined && object.end_time !== null) {
      message.end_time = Number(object.end_time);
    } else {
      message.end_time = 0;
    }
    return message;
  },

  toJSON(message: Order): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    message.owner !== undefined && (obj.owner = message.owner);
    message.status !== undefined && (obj.status = message.status);
    message.orderType !== undefined && (obj.orderType = message.orderType);
    message.denomAsk !== undefined && (obj.denomAsk = message.denomAsk);
    message.denomBid !== undefined && (obj.denomBid = message.denomBid);
    message.amount !== undefined && (obj.amount = message.amount);
    if (message.rate) {
      obj.rate = message.rate.map((e) => e);
    } else {
      obj.rate = [];
    }
    message.prev !== undefined && (obj.prev = message.prev);
    message.next !== undefined && (obj.next = message.next);
    message.beg_time !== undefined && (obj.beg_time = message.beg_time);
    message.end_time !== undefined && (obj.end_time = message.end_time);
    return obj;
  },

  fromPartial(object: DeepPartial<Order>): Order {
    const message = { ...baseOrder } as Order;
    message.rate = [];
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = 0;
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = object.orderType;
    } else {
      message.orderType = "";
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
      message.prev = 0;
    }
    if (object.next !== undefined && object.next !== null) {
      message.next = object.next;
    } else {
      message.next = 0;
    }
    if (object.beg_time !== undefined && object.beg_time !== null) {
      message.beg_time = object.beg_time;
    } else {
      message.beg_time = 0;
    }
    if (object.end_time !== undefined && object.end_time !== null) {
      message.end_time = object.end_time;
    } else {
      message.end_time = 0;
    }
    return message;
  },
};

const baseOrders: object = { uids: 0 };

export const Orders = {
  encode(message: Orders, writer: Writer = Writer.create()): Writer {
    writer.uint32(10).fork();
    for (const v of message.uids) {
      writer.uint64(v);
    }
    writer.ldelim();
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Orders {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOrders } as Orders;
    message.uids = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if ((tag & 7) === 2) {
            const end2 = reader.uint32() + reader.pos;
            while (reader.pos < end2) {
              message.uids.push(longToNumber(reader.uint64() as Long));
            }
          } else {
            message.uids.push(longToNumber(reader.uint64() as Long));
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Orders {
    const message = { ...baseOrders } as Orders;
    message.uids = [];
    if (object.uids !== undefined && object.uids !== null) {
      for (const e of object.uids) {
        message.uids.push(Number(e));
      }
    }
    return message;
  },

  toJSON(message: Orders): unknown {
    const obj: any = {};
    if (message.uids) {
      obj.uids = message.uids.map((e) => e);
    } else {
      obj.uids = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Orders>): Orders {
    const message = { ...baseOrders } as Orders;
    message.uids = [];
    if (object.uids !== undefined && object.uids !== null) {
      for (const e of object.uids) {
        message.uids.push(e);
      }
    }
    return message;
  },
};

const baseOrderResponse: object = {
  uid: 0,
  owner: "",
  status: "",
  orderType: "",
  denomAsk: "",
  denomBid: "",
  amount: "",
  rate: "",
  prev: 0,
  next: 0,
  beg_time: 0,
  end_time: 0,
};

export const OrderResponse = {
  encode(message: OrderResponse, writer: Writer = Writer.create()): Writer {
    if (message.uid !== 0) {
      writer.uint32(8).uint64(message.uid);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    if (message.status !== "") {
      writer.uint32(26).string(message.status);
    }
    if (message.orderType !== "") {
      writer.uint32(34).string(message.orderType);
    }
    if (message.denomAsk !== "") {
      writer.uint32(42).string(message.denomAsk);
    }
    if (message.denomBid !== "") {
      writer.uint32(50).string(message.denomBid);
    }
    if (message.amount !== "") {
      writer.uint32(58).string(message.amount);
    }
    for (const v of message.rate) {
      writer.uint32(66).string(v!);
    }
    if (message.prev !== 0) {
      writer.uint32(72).uint64(message.prev);
    }
    if (message.next !== 0) {
      writer.uint32(80).uint64(message.next);
    }
    if (message.beg_time !== 0) {
      writer.uint32(88).uint64(message.beg_time);
    }
    if (message.end_time !== 0) {
      writer.uint32(96).uint64(message.end_time);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOrderResponse } as OrderResponse;
    message.rate = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uid = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.owner = reader.string();
          break;
        case 3:
          message.status = reader.string();
          break;
        case 4:
          message.orderType = reader.string();
          break;
        case 5:
          message.denomAsk = reader.string();
          break;
        case 6:
          message.denomBid = reader.string();
          break;
        case 7:
          message.amount = reader.string();
          break;
        case 8:
          message.rate.push(reader.string());
          break;
        case 9:
          message.prev = longToNumber(reader.uint64() as Long);
          break;
        case 10:
          message.next = longToNumber(reader.uint64() as Long);
          break;
        case 11:
          message.beg_time = longToNumber(reader.uint64() as Long);
          break;
        case 12:
          message.end_time = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OrderResponse {
    const message = { ...baseOrderResponse } as OrderResponse;
    message.rate = [];
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = Number(object.uid);
    } else {
      message.uid = 0;
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = String(object.status);
    } else {
      message.status = "";
    }
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = String(object.orderType);
    } else {
      message.orderType = "";
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
      message.prev = Number(object.prev);
    } else {
      message.prev = 0;
    }
    if (object.next !== undefined && object.next !== null) {
      message.next = Number(object.next);
    } else {
      message.next = 0;
    }
    if (object.beg_time !== undefined && object.beg_time !== null) {
      message.beg_time = Number(object.beg_time);
    } else {
      message.beg_time = 0;
    }
    if (object.end_time !== undefined && object.end_time !== null) {
      message.end_time = Number(object.end_time);
    } else {
      message.end_time = 0;
    }
    return message;
  },

  toJSON(message: OrderResponse): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    message.owner !== undefined && (obj.owner = message.owner);
    message.status !== undefined && (obj.status = message.status);
    message.orderType !== undefined && (obj.orderType = message.orderType);
    message.denomAsk !== undefined && (obj.denomAsk = message.denomAsk);
    message.denomBid !== undefined && (obj.denomBid = message.denomBid);
    message.amount !== undefined && (obj.amount = message.amount);
    if (message.rate) {
      obj.rate = message.rate.map((e) => e);
    } else {
      obj.rate = [];
    }
    message.prev !== undefined && (obj.prev = message.prev);
    message.next !== undefined && (obj.next = message.next);
    message.beg_time !== undefined && (obj.beg_time = message.beg_time);
    message.end_time !== undefined && (obj.end_time = message.end_time);
    return obj;
  },

  fromPartial(object: DeepPartial<OrderResponse>): OrderResponse {
    const message = { ...baseOrderResponse } as OrderResponse;
    message.rate = [];
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = 0;
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.status !== undefined && object.status !== null) {
      message.status = object.status;
    } else {
      message.status = "";
    }
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = object.orderType;
    } else {
      message.orderType = "";
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
      message.prev = 0;
    }
    if (object.next !== undefined && object.next !== null) {
      message.next = object.next;
    } else {
      message.next = 0;
    }
    if (object.beg_time !== undefined && object.beg_time !== null) {
      message.beg_time = object.beg_time;
    } else {
      message.beg_time = 0;
    }
    if (object.end_time !== undefined && object.end_time !== null) {
      message.end_time = object.end_time;
    } else {
      message.end_time = 0;
    }
    return message;
  },
};

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
