/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "pendulumlabs.market.market";

/** Params defines the parameters for the module. */
export interface Params {
  /** leader earnings rate */
  earn_rate: string[];
  /** pool burning rate */
  burn_rate: string[];
  /** burn coin */
  burn_coin: string;
}

const baseParams: object = { earn_rate: "", burn_rate: "", burn_coin: "" };

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    for (const v of message.earn_rate) {
      writer.uint32(10).string(v!);
    }
    for (const v of message.burn_rate) {
      writer.uint32(18).string(v!);
    }
    if (message.burn_coin !== "") {
      writer.uint32(26).string(message.burn_coin);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    message.earn_rate = [];
    message.burn_rate = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.earn_rate.push(reader.string());
          break;
        case 2:
          message.burn_rate.push(reader.string());
          break;
        case 3:
          message.burn_coin = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    const message = { ...baseParams } as Params;
    message.earn_rate = [];
    message.burn_rate = [];
    if (object.earn_rate !== undefined && object.earn_rate !== null) {
      for (const e of object.earn_rate) {
        message.earn_rate.push(String(e));
      }
    }
    if (object.burn_rate !== undefined && object.burn_rate !== null) {
      for (const e of object.burn_rate) {
        message.burn_rate.push(String(e));
      }
    }
    if (object.burn_coin !== undefined && object.burn_coin !== null) {
      message.burn_coin = String(object.burn_coin);
    } else {
      message.burn_coin = "";
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    if (message.earn_rate) {
      obj.earn_rate = message.earn_rate.map((e) => e);
    } else {
      obj.earn_rate = [];
    }
    if (message.burn_rate) {
      obj.burn_rate = message.burn_rate.map((e) => e);
    } else {
      obj.burn_rate = [];
    }
    message.burn_coin !== undefined && (obj.burn_coin = message.burn_coin);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    message.earn_rate = [];
    message.burn_rate = [];
    if (object.earn_rate !== undefined && object.earn_rate !== null) {
      for (const e of object.earn_rate) {
        message.earn_rate.push(e);
      }
    }
    if (object.burn_rate !== undefined && object.burn_rate !== null) {
      for (const e of object.burn_rate) {
        message.burn_rate.push(e);
      }
    }
    if (object.burn_coin !== undefined && object.burn_coin !== null) {
      message.burn_coin = object.burn_coin;
    } else {
      message.burn_coin = "";
    }
    return message;
  },
};

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
