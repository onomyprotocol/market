/* eslint-disable */
import { Params } from "../market/params";
import { Pool } from "../market/pool";
import { Drop } from "../market/drop";
import { Member } from "../market/member";
import { Burnings } from "../market/burnings";
import { Order } from "../market/order";
import { Asset } from "../market/asset";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "pendulumlabs.market.market";

/** GenesisState defines the market module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  poolList: Pool[];
  dropList: Drop[];
  memberList: Member[];
  burningsList: Burnings[];
  orderList: Order[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  assetList: Asset[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.poolList) {
      Pool.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.dropList) {
      Drop.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.memberList) {
      Member.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    for (const v of message.burningsList) {
      Burnings.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    for (const v of message.orderList) {
      Order.encode(v!, writer.uint32(50).fork()).ldelim();
    }
    for (const v of message.assetList) {
      Asset.encode(v!, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.poolList = [];
    message.dropList = [];
    message.memberList = [];
    message.burningsList = [];
    message.orderList = [];
    message.assetList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.poolList.push(Pool.decode(reader, reader.uint32()));
          break;
        case 3:
          message.dropList.push(Drop.decode(reader, reader.uint32()));
          break;
        case 4:
          message.memberList.push(Member.decode(reader, reader.uint32()));
          break;
        case 5:
          message.burningsList.push(Burnings.decode(reader, reader.uint32()));
          break;
        case 6:
          message.orderList.push(Order.decode(reader, reader.uint32()));
          break;
        case 7:
          message.assetList.push(Asset.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.poolList = [];
    message.dropList = [];
    message.memberList = [];
    message.burningsList = [];
    message.orderList = [];
    message.assetList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.poolList !== undefined && object.poolList !== null) {
      for (const e of object.poolList) {
        message.poolList.push(Pool.fromJSON(e));
      }
    }
    if (object.dropList !== undefined && object.dropList !== null) {
      for (const e of object.dropList) {
        message.dropList.push(Drop.fromJSON(e));
      }
    }
    if (object.memberList !== undefined && object.memberList !== null) {
      for (const e of object.memberList) {
        message.memberList.push(Member.fromJSON(e));
      }
    }
    if (object.burningsList !== undefined && object.burningsList !== null) {
      for (const e of object.burningsList) {
        message.burningsList.push(Burnings.fromJSON(e));
      }
    }
    if (object.orderList !== undefined && object.orderList !== null) {
      for (const e of object.orderList) {
        message.orderList.push(Order.fromJSON(e));
      }
    }
    if (object.assetList !== undefined && object.assetList !== null) {
      for (const e of object.assetList) {
        message.assetList.push(Asset.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.poolList) {
      obj.poolList = message.poolList.map((e) =>
        e ? Pool.toJSON(e) : undefined
      );
    } else {
      obj.poolList = [];
    }
    if (message.dropList) {
      obj.dropList = message.dropList.map((e) =>
        e ? Drop.toJSON(e) : undefined
      );
    } else {
      obj.dropList = [];
    }
    if (message.memberList) {
      obj.memberList = message.memberList.map((e) =>
        e ? Member.toJSON(e) : undefined
      );
    } else {
      obj.memberList = [];
    }
    if (message.burningsList) {
      obj.burningsList = message.burningsList.map((e) =>
        e ? Burnings.toJSON(e) : undefined
      );
    } else {
      obj.burningsList = [];
    }
    if (message.orderList) {
      obj.orderList = message.orderList.map((e) =>
        e ? Order.toJSON(e) : undefined
      );
    } else {
      obj.orderList = [];
    }
    if (message.assetList) {
      obj.assetList = message.assetList.map((e) =>
        e ? Asset.toJSON(e) : undefined
      );
    } else {
      obj.assetList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.poolList = [];
    message.dropList = [];
    message.memberList = [];
    message.burningsList = [];
    message.orderList = [];
    message.assetList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.poolList !== undefined && object.poolList !== null) {
      for (const e of object.poolList) {
        message.poolList.push(Pool.fromPartial(e));
      }
    }
    if (object.dropList !== undefined && object.dropList !== null) {
      for (const e of object.dropList) {
        message.dropList.push(Drop.fromPartial(e));
      }
    }
    if (object.memberList !== undefined && object.memberList !== null) {
      for (const e of object.memberList) {
        message.memberList.push(Member.fromPartial(e));
      }
    }
    if (object.burningsList !== undefined && object.burningsList !== null) {
      for (const e of object.burningsList) {
        message.burningsList.push(Burnings.fromPartial(e));
      }
    }
    if (object.orderList !== undefined && object.orderList !== null) {
      for (const e of object.orderList) {
        message.orderList.push(Order.fromPartial(e));
      }
    }
    if (object.assetList !== undefined && object.assetList !== null) {
      for (const e of object.assetList) {
        message.assetList.push(Asset.fromPartial(e));
      }
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
