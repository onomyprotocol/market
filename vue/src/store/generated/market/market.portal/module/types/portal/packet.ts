/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "market.portal";

export interface PortalPacketData {
  noData: NoData | undefined;
  /** this line is used by starport scaffolding # ibc/packet/proto/field */
  subscribeRatePacket: SubscribeRatePacketData | undefined;
}

export interface NoData {}

/** SubscribeRatePacketData defines a struct for the packet payload */
export interface SubscribeRatePacketData {
  denomA: string;
  denomB: string;
}

/** SubscribeRatePacketAck defines a struct for the packet acknowledgment */
export interface SubscribeRatePacketAck {}

const basePortalPacketData: object = {};

export const PortalPacketData = {
  encode(message: PortalPacketData, writer: Writer = Writer.create()): Writer {
    if (message.noData !== undefined) {
      NoData.encode(message.noData, writer.uint32(10).fork()).ldelim();
    }
    if (message.subscribeRatePacket !== undefined) {
      SubscribeRatePacketData.encode(
        message.subscribeRatePacket,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): PortalPacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...basePortalPacketData } as PortalPacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.noData = NoData.decode(reader, reader.uint32());
          break;
        case 2:
          message.subscribeRatePacket = SubscribeRatePacketData.decode(
            reader,
            reader.uint32()
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): PortalPacketData {
    const message = { ...basePortalPacketData } as PortalPacketData;
    if (object.noData !== undefined && object.noData !== null) {
      message.noData = NoData.fromJSON(object.noData);
    } else {
      message.noData = undefined;
    }
    if (
      object.subscribeRatePacket !== undefined &&
      object.subscribeRatePacket !== null
    ) {
      message.subscribeRatePacket = SubscribeRatePacketData.fromJSON(
        object.subscribeRatePacket
      );
    } else {
      message.subscribeRatePacket = undefined;
    }
    return message;
  },

  toJSON(message: PortalPacketData): unknown {
    const obj: any = {};
    message.noData !== undefined &&
      (obj.noData = message.noData ? NoData.toJSON(message.noData) : undefined);
    message.subscribeRatePacket !== undefined &&
      (obj.subscribeRatePacket = message.subscribeRatePacket
        ? SubscribeRatePacketData.toJSON(message.subscribeRatePacket)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<PortalPacketData>): PortalPacketData {
    const message = { ...basePortalPacketData } as PortalPacketData;
    if (object.noData !== undefined && object.noData !== null) {
      message.noData = NoData.fromPartial(object.noData);
    } else {
      message.noData = undefined;
    }
    if (
      object.subscribeRatePacket !== undefined &&
      object.subscribeRatePacket !== null
    ) {
      message.subscribeRatePacket = SubscribeRatePacketData.fromPartial(
        object.subscribeRatePacket
      );
    } else {
      message.subscribeRatePacket = undefined;
    }
    return message;
  },
};

const baseNoData: object = {};

export const NoData = {
  encode(_: NoData, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): NoData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseNoData } as NoData;
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

  fromJSON(_: any): NoData {
    const message = { ...baseNoData } as NoData;
    return message;
  },

  toJSON(_: NoData): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<NoData>): NoData {
    const message = { ...baseNoData } as NoData;
    return message;
  },
};

const baseSubscribeRatePacketData: object = { denomA: "", denomB: "" };

export const SubscribeRatePacketData = {
  encode(
    message: SubscribeRatePacketData,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.denomA !== "") {
      writer.uint32(10).string(message.denomA);
    }
    if (message.denomB !== "") {
      writer.uint32(18).string(message.denomB);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SubscribeRatePacketData {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseSubscribeRatePacketData,
    } as SubscribeRatePacketData;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.denomA = reader.string();
          break;
        case 2:
          message.denomB = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): SubscribeRatePacketData {
    const message = {
      ...baseSubscribeRatePacketData,
    } as SubscribeRatePacketData;
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
    return message;
  },

  toJSON(message: SubscribeRatePacketData): unknown {
    const obj: any = {};
    message.denomA !== undefined && (obj.denomA = message.denomA);
    message.denomB !== undefined && (obj.denomB = message.denomB);
    return obj;
  },

  fromPartial(
    object: DeepPartial<SubscribeRatePacketData>
  ): SubscribeRatePacketData {
    const message = {
      ...baseSubscribeRatePacketData,
    } as SubscribeRatePacketData;
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
    return message;
  },
};

const baseSubscribeRatePacketAck: object = {};

export const SubscribeRatePacketAck = {
  encode(_: SubscribeRatePacketAck, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SubscribeRatePacketAck {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSubscribeRatePacketAck } as SubscribeRatePacketAck;
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

  fromJSON(_: any): SubscribeRatePacketAck {
    const message = { ...baseSubscribeRatePacketAck } as SubscribeRatePacketAck;
    return message;
  },

  toJSON(_: SubscribeRatePacketAck): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<SubscribeRatePacketAck>): SubscribeRatePacketAck {
    const message = { ...baseSubscribeRatePacketAck } as SubscribeRatePacketAck;
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
