/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "market.portal";

/** Params defines the parameters for the module. */
export interface Params {
  /** Onomy channel */
  onomy_channel: string;
  /** Reserve channel */
  reserve_channel: string;
}

const baseParams: object = { onomy_channel: "", reserve_channel: "" };

export const Params = {
  encode(message: Params, writer: Writer = Writer.create()): Writer {
    if (message.onomy_channel !== "") {
      writer.uint32(10).string(message.onomy_channel);
    }
    if (message.reserve_channel !== "") {
      writer.uint32(18).string(message.reserve_channel);
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
          message.onomy_channel = reader.string();
          break;
        case 2:
          message.reserve_channel = reader.string();
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
    if (object.onomy_channel !== undefined && object.onomy_channel !== null) {
      message.onomy_channel = String(object.onomy_channel);
    } else {
      message.onomy_channel = "";
    }
    if (
      object.reserve_channel !== undefined &&
      object.reserve_channel !== null
    ) {
      message.reserve_channel = String(object.reserve_channel);
    } else {
      message.reserve_channel = "";
    }
    return message;
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.onomy_channel !== undefined &&
      (obj.onomy_channel = message.onomy_channel);
    message.reserve_channel !== undefined &&
      (obj.reserve_channel = message.reserve_channel);
    return obj;
  },

  fromPartial(object: DeepPartial<Params>): Params {
    const message = { ...baseParams } as Params;
    if (object.onomy_channel !== undefined && object.onomy_channel !== null) {
      message.onomy_channel = object.onomy_channel;
    } else {
      message.onomy_channel = "";
    }
    if (
      object.reserve_channel !== undefined &&
      object.reserve_channel !== null
    ) {
      message.reserve_channel = object.reserve_channel;
    } else {
      message.reserve_channel = "";
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
