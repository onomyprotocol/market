/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "pendulumlabs.market.market";

/** Params defines the parameters for the module. */
export interface Params {
  /**
   * leader earnings rates
   * 1,2,3 Comma separated, no space
   */
  earn_rates: string;
  /** pool burning rate */
  burn_rate: string;
  /** burn coin */
  burn_coin: string;
  /** market_fee (parameter / 1000), 9999 representing as 99.99% */
  market_fee: string;
}

const baseParams: object = {
  earn_rates: "",
  burn_rate: "",
  burn_coin: "",
  market_fee: "",
};

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.earn_rates !== "") {
      writer.uint32(10).string(message.earn_rates);
    }
    if (message.burn_rate !== "") {
      writer.uint32(18).string(message.burn_rate);
    }
    if (message.burn_coin !== "") {
      writer.uint32(26).string(message.burn_coin);
    }
    if (message.market_fee !== "") {
      writer.uint32(34).string(message.market_fee);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseParams } as Params;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.earn_rates = reader.string();
          break;
        case 2:
          message.burn_rate = reader.string();
          break;
        case 3:
          message.burn_coin = reader.string();
          break;
        case 4:
          message.market_fee = reader.string();
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
    if (object.earn_rates !== undefined && object.earn_rates !== null) {
      message.earn_rates = String(object.earn_rates);
    } else {
      message.earn_rates = "";
    }
    if (object.burn_rate !== undefined && object.burn_rate !== null) {
      message.burn_rate = String(object.burn_rate);
    } else {
      message.burn_rate = "";
    }
    if (object.burn_coin !== undefined && object.burn_coin !== null) {
      message.burn_coin = String(object.burn_coin);
    } else {
      message.burn_coin = "";
    }
    if (object.market_fee !== undefined && object.market_fee !== null) {
      message.market_fee = String(object.market_fee);
    } else {
      message.market_fee = "";
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.earn_rates !== undefined && (obj.earn_rates = message.earn_rates);
    message.burn_rate !== undefined && (obj.burn_rate = message.burn_rate);
    message.burn_coin !== undefined && (obj.burn_coin = message.burn_coin);
    message.market_fee !== undefined && (obj.market_fee = message.market_fee);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (object.earn_rates !== undefined && object.earn_rates !== null) {
      message.earn_rates = object.earn_rates;
    } else {
      message.earn_rates = "";
    }
    if (object.burn_rate !== undefined && object.burn_rate !== null) {
      message.burn_rate = object.burn_rate;
    } else {
      message.burn_rate = "";
    }
    if (object.burn_coin !== undefined && object.burn_coin !== null) {
      message.burn_coin = object.burn_coin;
    } else {
      message.burn_coin = "";
    }
    if (object.market_fee !== undefined && object.market_fee !== null) {
      message.market_fee = object.market_fee;
    } else {
      message.market_fee = "";
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
