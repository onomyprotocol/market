/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "pendulumlabs.market.market";

export interface Pool {
  pair: string;
  denom1: string;
  denom2: string;
  volume1: Volume | undefined;
  volume2: Volume | undefined;
  leaders: Leader[];
  drops: string;
  history: number;
}

export interface Leader {
  address: string;
  drops: string;
}

export interface Volume {
  denom: string;
  amount: string;
}

const basePool: object = {
  pair: "",
  denom1: "",
  denom2: "",
  drops: "",
  history: 0,
};

export const Pool = {
  encode(message: Pool, writer: Writer = Writer.create()): Writer {
    if (message.pair !== "") {
      writer.uint32(10).string(message.pair);
    }
    if (message.denom1 !== "") {
      writer.uint32(18).string(message.denom1);
    }
    if (message.denom2 !== "") {
      writer.uint32(26).string(message.denom2);
    }
    if (message.volume1 !== undefined) {
      Volume.encode(message.volume1, writer.uint32(34).fork()).ldelim();
    }
    if (message.volume2 !== undefined) {
      Volume.encode(message.volume2, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.leaders) {
      Leader.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    if (message.drops !== "") {
      writer.uint32(58).string(message.drops);
    }
    if (message.history !== 0) {
      writer.uint32(64).uint64(message.history);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Pool {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePool } as Pool;
    message.leaders = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pair = reader.string();
          break;
        case 2:
          message.denom1 = reader.string();
          break;
        case 3:
          message.denom2 = reader.string();
          break;
        case 4:
          message.volume1 = Volume.decode(reader, reader.uint32());
          break;
        case 5:
          message.volume2 = Volume.decode(reader, reader.uint32());
          break;
        case 6:
          message.leaders.push(Leader.decode(reader, reader.uint32()));
          break;
        case 7:
          message.drops = reader.string();
          break;
        case 8:
          message.history = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Pool {
    const message = { ...basePool } as Pool;
    message.leaders = [];
    if (object.pair !== undefined && object.pair !== null) {
      message.pair = String(object.pair);
    } else {
      message.pair = "";
    }
    if (object.denom1 !== undefined && object.denom1 !== null) {
      message.denom1 = String(object.denom1);
    } else {
      message.denom1 = "";
    }
    if (object.denom2 !== undefined && object.denom2 !== null) {
      message.denom2 = String(object.denom2);
    } else {
      message.denom2 = "";
    }
    if (object.volume1 !== undefined && object.volume1 !== null) {
      message.volume1 = Volume.fromJSON(object.volume1);
    } else {
      message.volume1 = undefined;
    }
    if (object.volume2 !== undefined && object.volume2 !== null) {
      message.volume2 = Volume.fromJSON(object.volume2);
    } else {
      message.volume2 = undefined;
    }
    if (object.leaders !== undefined && object.leaders !== null) {
      for (const e of object.leaders) {
        message.leaders.push(Leader.fromJSON(e));
      }
    }
    if (object.drops !== undefined && object.drops !== null) {
      message.drops = String(object.drops);
    } else {
      message.drops = "";
    }
    if (object.history !== undefined && object.history !== null) {
      message.history = Number(object.history);
    } else {
      message.history = 0;
    }
    return message;
  },

  toJSON(message: Pool): unknown {
    const obj: any = {};
    message.pair !== undefined && (obj.pair = message.pair);
    message.denom1 !== undefined && (obj.denom1 = message.denom1);
    message.denom2 !== undefined && (obj.denom2 = message.denom2);
    message.volume1 !== undefined &&
      (obj.volume1 = message.volume1
        ? Volume.toJSON(message.volume1)
        : undefined);
    message.volume2 !== undefined &&
      (obj.volume2 = message.volume2
        ? Volume.toJSON(message.volume2)
        : undefined);
    if (message.leaders) {
      obj.leaders = message.leaders.map((e) =>
        e ? Leader.toJSON(e) : undefined
      );
    } else {
      obj.leaders = [];
    }
    message.drops !== undefined && (obj.drops = message.drops);
    message.history !== undefined && (obj.history = message.history);
    return obj;
  },

  fromPartial(object: DeepPartial<Pool>): Pool {
    const message = { ...basePool } as Pool;
    message.leaders = [];
    if (object.pair !== undefined && object.pair !== null) {
      message.pair = object.pair;
    } else {
      message.pair = "";
    }
    if (object.denom1 !== undefined && object.denom1 !== null) {
      message.denom1 = object.denom1;
    } else {
      message.denom1 = "";
    }
    if (object.denom2 !== undefined && object.denom2 !== null) {
      message.denom2 = object.denom2;
    } else {
      message.denom2 = "";
    }
    if (object.volume1 !== undefined && object.volume1 !== null) {
      message.volume1 = Volume.fromPartial(object.volume1);
    } else {
      message.volume1 = undefined;
    }
    if (object.volume2 !== undefined && object.volume2 !== null) {
      message.volume2 = Volume.fromPartial(object.volume2);
    } else {
      message.volume2 = undefined;
    }
    if (object.leaders !== undefined && object.leaders !== null) {
      for (const e of object.leaders) {
        message.leaders.push(Leader.fromPartial(e));
      }
    }
    if (object.drops !== undefined && object.drops !== null) {
      message.drops = object.drops;
    } else {
      message.drops = "";
    }
    if (object.history !== undefined && object.history !== null) {
      message.history = object.history;
    } else {
      message.history = 0;
    }
    return message;
  },
};

const baseLeader: object = { address: "", drops: "" };

export const Leader = {
  encode(message: Leader, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.drops !== "") {
      writer.uint32(18).string(message.drops);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Leader {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseLeader } as Leader;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.drops = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Leader {
    const message = { ...baseLeader } as Leader;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.drops !== undefined && object.drops !== null) {
      message.drops = String(object.drops);
    } else {
      message.drops = "";
    }
    return message;
  },

  toJSON(message: Leader): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.drops !== undefined && (obj.drops = message.drops);
    return obj;
  },

  fromPartial(object: DeepPartial<Leader>): Leader {
    const message = { ...baseLeader } as Leader;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.drops !== undefined && object.drops !== null) {
      message.drops = object.drops;
    } else {
      message.drops = "";
    }
    return message;
  },
};

const baseVolume: object = { denom: "", amount: "" };

export const Volume = {
  encode(message: Volume, writer: Writer = Writer.create()): Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Volume {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVolume } as Volume;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;
        case 2:
          message.amount = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Volume {
    const message = { ...baseVolume } as Volume;
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom);
    } else {
      message.denom = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    return message;
  },

  toJSON(message: Volume): unknown {
    const obj: any = {};
    message.denom !== undefined && (obj.denom = message.denom);
    message.amount !== undefined && (obj.amount = message.amount);
    return obj;
  },

  fromPartial(object: DeepPartial<Volume>): Volume {
    const message = { ...baseVolume } as Volume;
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom;
    } else {
      message.denom = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
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
