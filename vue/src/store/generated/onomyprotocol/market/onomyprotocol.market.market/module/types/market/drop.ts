/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "onomyprotocol.market.market";

export interface Drop {
  uid: number;
  owner: string;
  pair: string;
  drops: string;
  sum: string;
  active: boolean;
  rate1: string[];
  rate2: string[];
}

const baseDrop: object = {
  uid: 0,
  owner: "",
  pair: "",
  drops: "",
  sum: "",
  active: false,
  rate1: "",
  rate2: "",
};

export const Drop = {
  encode(message: Drop, writer: Writer = Writer.create()): Writer {
    if (message.uid !== 0) {
      writer.uint32(8).uint64(message.uid);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    if (message.pair !== "") {
      writer.uint32(26).string(message.pair);
    }
    if (message.drops !== "") {
      writer.uint32(34).string(message.drops);
    }
    if (message.sum !== "") {
      writer.uint32(42).string(message.sum);
    }
    if (message.active === true) {
      writer.uint32(48).bool(message.active);
    }
    for (const v of message.rate1) {
      writer.uint32(58).string(v!);
    }
    for (const v of message.rate2) {
      writer.uint32(66).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Drop {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseDrop } as Drop;
    message.rate1 = [];
    message.rate2 = [];
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
          message.pair = reader.string();
          break;
        case 4:
          message.drops = reader.string();
          break;
        case 5:
          message.sum = reader.string();
          break;
        case 6:
          message.active = reader.bool();
          break;
        case 7:
          message.rate1.push(reader.string());
          break;
        case 8:
          message.rate2.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Drop {
    const message = { ...baseDrop } as Drop;
    message.rate1 = [];
    message.rate2 = [];
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
    if (object.sum !== undefined && object.sum !== null) {
      message.sum = String(object.sum);
    } else {
      message.sum = "";
    }
    if (object.active !== undefined && object.active !== null) {
      message.active = Boolean(object.active);
    } else {
      message.active = false;
    }
    if (object.rate1 !== undefined && object.rate1 !== null) {
      for (const e of object.rate1) {
        message.rate1.push(String(e));
      }
    }
    if (object.rate2 !== undefined && object.rate2 !== null) {
      for (const e of object.rate2) {
        message.rate2.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: Drop): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    message.owner !== undefined && (obj.owner = message.owner);
    message.pair !== undefined && (obj.pair = message.pair);
    message.drops !== undefined && (obj.drops = message.drops);
    message.sum !== undefined && (obj.sum = message.sum);
    message.active !== undefined && (obj.active = message.active);
    if (message.rate1) {
      obj.rate1 = message.rate1.map((e) => e);
    } else {
      obj.rate1 = [];
    }
    if (message.rate2) {
      obj.rate2 = message.rate2.map((e) => e);
    } else {
      obj.rate2 = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Drop>): Drop {
    const message = { ...baseDrop } as Drop;
    message.rate1 = [];
    message.rate2 = [];
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
    if (object.sum !== undefined && object.sum !== null) {
      message.sum = object.sum;
    } else {
      message.sum = "";
    }
    if (object.active !== undefined && object.active !== null) {
      message.active = object.active;
    } else {
      message.active = false;
    }
    if (object.rate1 !== undefined && object.rate1 !== null) {
      for (const e of object.rate1) {
        message.rate1.push(e);
      }
    }
    if (object.rate2 !== undefined && object.rate2 !== null) {
      for (const e of object.rate2) {
        message.rate2.push(e);
      }
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
