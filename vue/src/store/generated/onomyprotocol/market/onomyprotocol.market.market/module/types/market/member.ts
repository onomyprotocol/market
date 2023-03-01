/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "pendulum-labs.market.market";

export interface Member {
  pair: string;
  denomA: string;
  denomB: string;
  balance: string;
  previous: string;
  limit: number;
  stop: number;
  protect: number;
}

const baseMember: object = {
  pair: "",
  denomA: "",
  denomB: "",
  balance: "",
  previous: "",
  limit: 0,
  stop: 0,
  protect: 0,
};

export const Member = {
  encode(message: Member, writer: Writer = Writer.create()): Writer {
    if (message.pair !== "") {
      writer.uint32(10).string(message.pair);
    }
    if (message.denomA !== "") {
      writer.uint32(18).string(message.denomA);
    }
    if (message.denomB !== "") {
      writer.uint32(26).string(message.denomB);
    }
    if (message.balance !== "") {
      writer.uint32(34).string(message.balance);
    }
    if (message.previous !== "") {
      writer.uint32(42).string(message.previous);
    }
    if (message.limit !== 0) {
      writer.uint32(48).uint64(message.limit);
    }
    if (message.stop !== 0) {
      writer.uint32(56).uint64(message.stop);
    }
    if (message.protect !== 0) {
      writer.uint32(64).uint64(message.protect);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Member {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMember } as Member;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pair = reader.string();
          break;
        case 2:
          message.denomA = reader.string();
          break;
        case 3:
          message.denomB = reader.string();
          break;
        case 4:
          message.balance = reader.string();
          break;
        case 5:
          message.previous = reader.string();
          break;
        case 6:
          message.limit = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.stop = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.protect = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Member {
    const message = { ...baseMember } as Member;
    if (object.pair !== undefined && object.pair !== null) {
      message.pair = String(object.pair);
    } else {
      message.pair = "";
    }
    if (object.denomA !== undefined && object.denomA !== null) {
      message.denomA = String(object.denomA);
    } else {
      message.denomA = "";
    }
    if (object.denomB !== undefined && object.denomB !== null) {
      message.denomB = String(object.denomB);
    } else {
      message.denomB = "";
    }
    if (object.balance !== undefined && object.balance !== null) {
      message.balance = String(object.balance);
    } else {
      message.balance = "";
    }
    if (object.previous !== undefined && object.previous !== null) {
      message.previous = String(object.previous);
    } else {
      message.previous = "";
    }
    if (object.limit !== undefined && object.limit !== null) {
      message.limit = Number(object.limit);
    } else {
      message.limit = 0;
    }
    if (object.stop !== undefined && object.stop !== null) {
      message.stop = Number(object.stop);
    } else {
      message.stop = 0;
    }
    if (object.protect !== undefined && object.protect !== null) {
      message.protect = Number(object.protect);
    } else {
      message.protect = 0;
    }
    return message;
  },

  toJSON(message: Member): unknown {
    const obj: any = {};
    message.pair !== undefined && (obj.pair = message.pair);
    message.denomA !== undefined && (obj.denomA = message.denomA);
    message.denomB !== undefined && (obj.denomB = message.denomB);
    message.balance !== undefined && (obj.balance = message.balance);
    message.previous !== undefined && (obj.previous = message.previous);
    message.limit !== undefined && (obj.limit = message.limit);
    message.stop !== undefined && (obj.stop = message.stop);
    message.protect !== undefined && (obj.protect = message.protect);
    return obj;
  },

  fromPartial(object: DeepPartial<Member>): Member {
    const message = { ...baseMember } as Member;
    if (object.pair !== undefined && object.pair !== null) {
      message.pair = object.pair;
    } else {
      message.pair = "";
    }
    if (object.denomA !== undefined && object.denomA !== null) {
      message.denomA = object.denomA;
    } else {
      message.denomA = "";
    }
    if (object.denomB !== undefined && object.denomB !== null) {
      message.denomB = object.denomB;
    } else {
      message.denomB = "";
    }
    if (object.balance !== undefined && object.balance !== null) {
      message.balance = object.balance;
    } else {
      message.balance = "";
    }
    if (object.previous !== undefined && object.previous !== null) {
      message.previous = object.previous;
    } else {
      message.previous = "";
    }
    if (object.limit !== undefined && object.limit !== null) {
      message.limit = object.limit;
    } else {
      message.limit = 0;
    }
    if (object.stop !== undefined && object.stop !== null) {
      message.stop = object.stop;
    } else {
      message.stop = 0;
    }
    if (object.protect !== undefined && object.protect !== null) {
      message.protect = object.protect;
    } else {
      message.protect = 0;
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
