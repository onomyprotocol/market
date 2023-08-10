/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";
import { Params } from "../market/params";
import { Pool } from "../market/pool";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Drop } from "../market/drop";
import { Member } from "../market/member";
import { Burnings } from "../market/burnings";
import { Order, Orders, OrderResponse } from "../market/order";

export const protobufPackage = "pendulumlabs.market.market";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetPoolRequest {
  pair: string;
}

export interface QueryGetPoolResponse {
  pool: Pool | undefined;
}

export interface QueryAllPoolRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllPoolResponse {
  pool: Pool[];
  pagination: PageResponse | undefined;
}

export interface QueryDropRequest {
  uid: number;
}

export interface QueryDropResponse {
  drop: Drop | undefined;
}

export interface QueryAllDropRequest {
  pagination: PageRequest | undefined;
}

export interface QueryDropsResponse {
  drops: Drop[];
  pagination: PageResponse | undefined;
}

export interface QueryGetMemberRequest {
  denomA: string;
  denomB: string;
}

export interface QueryGetMemberResponse {
  member: Member | undefined;
}

export interface QueryAllMemberRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllMemberResponse {
  member: Member[];
  pagination: PageResponse | undefined;
}

export interface QueryGetBurningsRequest {
  denom: string;
}

export interface QueryGetBurningsResponse {
  burnings: Burnings | undefined;
}

export interface QueryAllBurningsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllBurningsResponse {
  burnings: Burnings[];
  pagination: PageResponse | undefined;
}

export interface QueryGetOrderRequest {
  uid: number;
}

export interface QueryGetOrderResponse {
  order: Order | undefined;
}

export interface QueryOrdersResponse {
  orders: Order[];
  pagination: PageResponse | undefined;
}

export interface QueryAllOrderRequest {
  pagination: PageRequest | undefined;
}

export interface QueryOrderOwnerRequest {
  address: string;
  pagination: PageRequest | undefined;
}

export interface QueryOrderOwnerUidsResponse {
  orders: Orders | undefined;
  pagination: PageResponse | undefined;
}

export interface QueryOrderOwnerPairRequest {
  address: string;
  pair: string;
  pagination: PageRequest | undefined;
}

export interface QueryOrderOwnerPairResponse {
  order: Order[];
  pagination: PageResponse | undefined;
}

export interface QueryBookRequest {
  denomA: string;
  denomB: string;
  orderType: string;
  pagination: PageRequest | undefined;
}

export interface QueryBookResponse {
  book: OrderResponse[];
  pagination: PageResponse | undefined;
}

export interface QueryBookendsRequest {
  coinA: string;
  coinB: string;
  orderType: string;
  rate: string[];
}

export interface QueryBookendsResponse {
  coinA: string;
  coinB: string;
  orderType: string;
  rate: string[];
  prev: number;
  next: number;
}

export interface QueryHistoryRequest {
  pair: string;
  length: string;
  pagination: PageRequest | undefined;
}

export interface QueryHistoryResponse {
  history: OrderResponse[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetPoolRequest: object = { pair: "" };

export const QueryGetPoolRequest = {
  encode(
    message: QueryGetPoolRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pair !== "") {
      writer.uint32(10).string(message.pair);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPoolRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetPoolRequest } as QueryGetPoolRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pair = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPoolRequest {
    const message = { ...baseQueryGetPoolRequest } as QueryGetPoolRequest;
    if (object.pair !== undefined && object.pair !== null) {
      message.pair = String(object.pair);
    } else {
      message.pair = "";
    }
    return message;
  },

  toJSON(message: QueryGetPoolRequest): unknown {
    const obj: any = {};
    message.pair !== undefined && (obj.pair = message.pair);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetPoolRequest>): QueryGetPoolRequest {
    const message = { ...baseQueryGetPoolRequest } as QueryGetPoolRequest;
    if (object.pair !== undefined && object.pair !== null) {
      message.pair = object.pair;
    } else {
      message.pair = "";
    }
    return message;
  },
};

const baseQueryGetPoolResponse: object = {};

export const QueryGetPoolResponse = {
  encode(
    message: QueryGetPoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pool !== undefined) {
      Pool.encode(message.pool, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetPoolResponse } as QueryGetPoolResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pool = Pool.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetPoolResponse {
    const message = { ...baseQueryGetPoolResponse } as QueryGetPoolResponse;
    if (object.pool !== undefined && object.pool !== null) {
      message.pool = Pool.fromJSON(object.pool);
    } else {
      message.pool = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetPoolResponse): unknown {
    const obj: any = {};
    message.pool !== undefined &&
      (obj.pool = message.pool ? Pool.toJSON(message.pool) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetPoolResponse>): QueryGetPoolResponse {
    const message = { ...baseQueryGetPoolResponse } as QueryGetPoolResponse;
    if (object.pool !== undefined && object.pool !== null) {
      message.pool = Pool.fromPartial(object.pool);
    } else {
      message.pool = undefined;
    }
    return message;
  },
};

const baseQueryAllPoolRequest: object = {};

export const QueryAllPoolRequest = {
  encode(
    message: QueryAllPoolRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPoolRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllPoolRequest } as QueryAllPoolRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPoolRequest {
    const message = { ...baseQueryAllPoolRequest } as QueryAllPoolRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPoolRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllPoolRequest>): QueryAllPoolRequest {
    const message = { ...baseQueryAllPoolRequest } as QueryAllPoolRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllPoolResponse: object = {};

export const QueryAllPoolResponse = {
  encode(
    message: QueryAllPoolResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.pool) {
      Pool.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllPoolResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllPoolResponse } as QueryAllPoolResponse;
    message.pool = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pool.push(Pool.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllPoolResponse {
    const message = { ...baseQueryAllPoolResponse } as QueryAllPoolResponse;
    message.pool = [];
    if (object.pool !== undefined && object.pool !== null) {
      for (const e of object.pool) {
        message.pool.push(Pool.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllPoolResponse): unknown {
    const obj: any = {};
    if (message.pool) {
      obj.pool = message.pool.map((e) => (e ? Pool.toJSON(e) : undefined));
    } else {
      obj.pool = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllPoolResponse>): QueryAllPoolResponse {
    const message = { ...baseQueryAllPoolResponse } as QueryAllPoolResponse;
    message.pool = [];
    if (object.pool !== undefined && object.pool !== null) {
      for (const e of object.pool) {
        message.pool.push(Pool.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryDropRequest: object = { uid: 0 };

export const QueryDropRequest = {
  encode(message: QueryDropRequest, writer: Writer = Writer.create()): Writer {
    if (message.uid !== 0) {
      writer.uint32(8).uint64(message.uid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryDropRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryDropRequest } as QueryDropRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uid = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDropRequest {
    const message = { ...baseQueryDropRequest } as QueryDropRequest;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = Number(object.uid);
    } else {
      message.uid = 0;
    }
    return message;
  },

  toJSON(message: QueryDropRequest): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryDropRequest>): QueryDropRequest {
    const message = { ...baseQueryDropRequest } as QueryDropRequest;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = 0;
    }
    return message;
  },
};

const baseQueryDropResponse: object = {};

export const QueryDropResponse = {
  encode(message: QueryDropResponse, writer: Writer = Writer.create()): Writer {
    if (message.drop !== undefined) {
      Drop.encode(message.drop, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryDropResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryDropResponse } as QueryDropResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.drop = Drop.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDropResponse {
    const message = { ...baseQueryDropResponse } as QueryDropResponse;
    if (object.drop !== undefined && object.drop !== null) {
      message.drop = Drop.fromJSON(object.drop);
    } else {
      message.drop = undefined;
    }
    return message;
  },

  toJSON(message: QueryDropResponse): unknown {
    const obj: any = {};
    message.drop !== undefined &&
      (obj.drop = message.drop ? Drop.toJSON(message.drop) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryDropResponse>): QueryDropResponse {
    const message = { ...baseQueryDropResponse } as QueryDropResponse;
    if (object.drop !== undefined && object.drop !== null) {
      message.drop = Drop.fromPartial(object.drop);
    } else {
      message.drop = undefined;
    }
    return message;
  },
};

const baseQueryAllDropRequest: object = {};

export const QueryAllDropRequest = {
  encode(
    message: QueryAllDropRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllDropRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllDropRequest } as QueryAllDropRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllDropRequest {
    const message = { ...baseQueryAllDropRequest } as QueryAllDropRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllDropRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllDropRequest>): QueryAllDropRequest {
    const message = { ...baseQueryAllDropRequest } as QueryAllDropRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryDropsResponse: object = {};

export const QueryDropsResponse = {
  encode(
    message: QueryDropsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.drops) {
      Drop.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryDropsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryDropsResponse } as QueryDropsResponse;
    message.drops = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.drops.push(Drop.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryDropsResponse {
    const message = { ...baseQueryDropsResponse } as QueryDropsResponse;
    message.drops = [];
    if (object.drops !== undefined && object.drops !== null) {
      for (const e of object.drops) {
        message.drops.push(Drop.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryDropsResponse): unknown {
    const obj: any = {};
    if (message.drops) {
      obj.drops = message.drops.map((e) => (e ? Drop.toJSON(e) : undefined));
    } else {
      obj.drops = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryDropsResponse>): QueryDropsResponse {
    const message = { ...baseQueryDropsResponse } as QueryDropsResponse;
    message.drops = [];
    if (object.drops !== undefined && object.drops !== null) {
      for (const e of object.drops) {
        message.drops.push(Drop.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetMemberRequest: object = { denomA: "", denomB: "" };

export const QueryGetMemberRequest = {
  encode(
    message: QueryGetMemberRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.denomA !== "") {
      writer.uint32(18).string(message.denomA);
    }
    if (message.denomB !== "") {
      writer.uint32(26).string(message.denomB);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetMemberRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetMemberRequest } as QueryGetMemberRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 2:
          message.denomA = reader.string();
          break;
        case 3:
          message.denomB = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMemberRequest {
    const message = { ...baseQueryGetMemberRequest } as QueryGetMemberRequest;
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

  toJSON(message: QueryGetMemberRequest): unknown {
    const obj: any = {};
    message.denomA !== undefined && (obj.denomA = message.denomA);
    message.denomB !== undefined && (obj.denomB = message.denomB);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetMemberRequest>
  ): QueryGetMemberRequest {
    const message = { ...baseQueryGetMemberRequest } as QueryGetMemberRequest;
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

const baseQueryGetMemberResponse: object = {};

export const QueryGetMemberResponse = {
  encode(
    message: QueryGetMemberResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.member !== undefined) {
      Member.encode(message.member, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetMemberResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetMemberResponse } as QueryGetMemberResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.member = Member.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetMemberResponse {
    const message = { ...baseQueryGetMemberResponse } as QueryGetMemberResponse;
    if (object.member !== undefined && object.member !== null) {
      message.member = Member.fromJSON(object.member);
    } else {
      message.member = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetMemberResponse): unknown {
    const obj: any = {};
    message.member !== undefined &&
      (obj.member = message.member ? Member.toJSON(message.member) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetMemberResponse>
  ): QueryGetMemberResponse {
    const message = { ...baseQueryGetMemberResponse } as QueryGetMemberResponse;
    if (object.member !== undefined && object.member !== null) {
      message.member = Member.fromPartial(object.member);
    } else {
      message.member = undefined;
    }
    return message;
  },
};

const baseQueryAllMemberRequest: object = {};

export const QueryAllMemberRequest = {
  encode(
    message: QueryAllMemberRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllMemberRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllMemberRequest } as QueryAllMemberRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllMemberRequest {
    const message = { ...baseQueryAllMemberRequest } as QueryAllMemberRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllMemberRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllMemberRequest>
  ): QueryAllMemberRequest {
    const message = { ...baseQueryAllMemberRequest } as QueryAllMemberRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllMemberResponse: object = {};

export const QueryAllMemberResponse = {
  encode(
    message: QueryAllMemberResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.member) {
      Member.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllMemberResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllMemberResponse } as QueryAllMemberResponse;
    message.member = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.member.push(Member.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllMemberResponse {
    const message = { ...baseQueryAllMemberResponse } as QueryAllMemberResponse;
    message.member = [];
    if (object.member !== undefined && object.member !== null) {
      for (const e of object.member) {
        message.member.push(Member.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllMemberResponse): unknown {
    const obj: any = {};
    if (message.member) {
      obj.member = message.member.map((e) =>
        e ? Member.toJSON(e) : undefined
      );
    } else {
      obj.member = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllMemberResponse>
  ): QueryAllMemberResponse {
    const message = { ...baseQueryAllMemberResponse } as QueryAllMemberResponse;
    message.member = [];
    if (object.member !== undefined && object.member !== null) {
      for (const e of object.member) {
        message.member.push(Member.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetBurningsRequest: object = { denom: "" };

export const QueryGetBurningsRequest = {
  encode(
    message: QueryGetBurningsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.denom !== "") {
      writer.uint32(10).string(message.denom);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBurningsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetBurningsRequest,
    } as QueryGetBurningsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.denom = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetBurningsRequest {
    const message = {
      ...baseQueryGetBurningsRequest,
    } as QueryGetBurningsRequest;
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = String(object.denom);
    } else {
      message.denom = "";
    }
    return message;
  },

  toJSON(message: QueryGetBurningsRequest): unknown {
    const obj: any = {};
    message.denom !== undefined && (obj.denom = message.denom);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetBurningsRequest>
  ): QueryGetBurningsRequest {
    const message = {
      ...baseQueryGetBurningsRequest,
    } as QueryGetBurningsRequest;
    if (object.denom !== undefined && object.denom !== null) {
      message.denom = object.denom;
    } else {
      message.denom = "";
    }
    return message;
  },
};

const baseQueryGetBurningsResponse: object = {};

export const QueryGetBurningsResponse = {
  encode(
    message: QueryGetBurningsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.burnings !== undefined) {
      Burnings.encode(message.burnings, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryGetBurningsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetBurningsResponse,
    } as QueryGetBurningsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.burnings = Burnings.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetBurningsResponse {
    const message = {
      ...baseQueryGetBurningsResponse,
    } as QueryGetBurningsResponse;
    if (object.burnings !== undefined && object.burnings !== null) {
      message.burnings = Burnings.fromJSON(object.burnings);
    } else {
      message.burnings = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetBurningsResponse): unknown {
    const obj: any = {};
    message.burnings !== undefined &&
      (obj.burnings = message.burnings
        ? Burnings.toJSON(message.burnings)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetBurningsResponse>
  ): QueryGetBurningsResponse {
    const message = {
      ...baseQueryGetBurningsResponse,
    } as QueryGetBurningsResponse;
    if (object.burnings !== undefined && object.burnings !== null) {
      message.burnings = Burnings.fromPartial(object.burnings);
    } else {
      message.burnings = undefined;
    }
    return message;
  },
};

const baseQueryAllBurningsRequest: object = {};

export const QueryAllBurningsRequest = {
  encode(
    message: QueryAllBurningsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllBurningsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllBurningsRequest,
    } as QueryAllBurningsRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllBurningsRequest {
    const message = {
      ...baseQueryAllBurningsRequest,
    } as QueryAllBurningsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllBurningsRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllBurningsRequest>
  ): QueryAllBurningsRequest {
    const message = {
      ...baseQueryAllBurningsRequest,
    } as QueryAllBurningsRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllBurningsResponse: object = {};

export const QueryAllBurningsResponse = {
  encode(
    message: QueryAllBurningsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.burnings) {
      Burnings.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryAllBurningsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllBurningsResponse,
    } as QueryAllBurningsResponse;
    message.burnings = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.burnings.push(Burnings.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllBurningsResponse {
    const message = {
      ...baseQueryAllBurningsResponse,
    } as QueryAllBurningsResponse;
    message.burnings = [];
    if (object.burnings !== undefined && object.burnings !== null) {
      for (const e of object.burnings) {
        message.burnings.push(Burnings.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllBurningsResponse): unknown {
    const obj: any = {};
    if (message.burnings) {
      obj.burnings = message.burnings.map((e) =>
        e ? Burnings.toJSON(e) : undefined
      );
    } else {
      obj.burnings = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllBurningsResponse>
  ): QueryAllBurningsResponse {
    const message = {
      ...baseQueryAllBurningsResponse,
    } as QueryAllBurningsResponse;
    message.burnings = [];
    if (object.burnings !== undefined && object.burnings !== null) {
      for (const e of object.burnings) {
        message.burnings.push(Burnings.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetOrderRequest: object = { uid: 0 };

export const QueryGetOrderRequest = {
  encode(
    message: QueryGetOrderRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.uid !== 0) {
      writer.uint32(8).uint64(message.uid);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetOrderRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetOrderRequest } as QueryGetOrderRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.uid = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOrderRequest {
    const message = { ...baseQueryGetOrderRequest } as QueryGetOrderRequest;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = Number(object.uid);
    } else {
      message.uid = 0;
    }
    return message;
  },

  toJSON(message: QueryGetOrderRequest): unknown {
    const obj: any = {};
    message.uid !== undefined && (obj.uid = message.uid);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetOrderRequest>): QueryGetOrderRequest {
    const message = { ...baseQueryGetOrderRequest } as QueryGetOrderRequest;
    if (object.uid !== undefined && object.uid !== null) {
      message.uid = object.uid;
    } else {
      message.uid = 0;
    }
    return message;
  },
};

const baseQueryGetOrderResponse: object = {};

export const QueryGetOrderResponse = {
  encode(
    message: QueryGetOrderResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.order !== undefined) {
      Order.encode(message.order, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetOrderResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetOrderResponse } as QueryGetOrderResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.order = Order.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOrderResponse {
    const message = { ...baseQueryGetOrderResponse } as QueryGetOrderResponse;
    if (object.order !== undefined && object.order !== null) {
      message.order = Order.fromJSON(object.order);
    } else {
      message.order = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetOrderResponse): unknown {
    const obj: any = {};
    message.order !== undefined &&
      (obj.order = message.order ? Order.toJSON(message.order) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetOrderResponse>
  ): QueryGetOrderResponse {
    const message = { ...baseQueryGetOrderResponse } as QueryGetOrderResponse;
    if (object.order !== undefined && object.order !== null) {
      message.order = Order.fromPartial(object.order);
    } else {
      message.order = undefined;
    }
    return message;
  },
};

const baseQueryOrdersResponse: object = {};

export const QueryOrdersResponse = {
  encode(
    message: QueryOrdersResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.orders) {
      Order.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryOrdersResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryOrdersResponse } as QueryOrdersResponse;
    message.orders = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.orders.push(Order.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryOrdersResponse {
    const message = { ...baseQueryOrdersResponse } as QueryOrdersResponse;
    message.orders = [];
    if (object.orders !== undefined && object.orders !== null) {
      for (const e of object.orders) {
        message.orders.push(Order.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryOrdersResponse): unknown {
    const obj: any = {};
    if (message.orders) {
      obj.orders = message.orders.map((e) => (e ? Order.toJSON(e) : undefined));
    } else {
      obj.orders = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryOrdersResponse>): QueryOrdersResponse {
    const message = { ...baseQueryOrdersResponse } as QueryOrdersResponse;
    message.orders = [];
    if (object.orders !== undefined && object.orders !== null) {
      for (const e of object.orders) {
        message.orders.push(Order.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllOrderRequest: object = {};

export const QueryAllOrderRequest = {
  encode(
    message: QueryAllOrderRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllOrderRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllOrderRequest } as QueryAllOrderRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOrderRequest {
    const message = { ...baseQueryAllOrderRequest } as QueryAllOrderRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllOrderRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllOrderRequest>): QueryAllOrderRequest {
    const message = { ...baseQueryAllOrderRequest } as QueryAllOrderRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryOrderOwnerRequest: object = { address: "" };

export const QueryOrderOwnerRequest = {
  encode(
    message: QueryOrderOwnerRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryOrderOwnerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryOrderOwnerRequest } as QueryOrderOwnerRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryOrderOwnerRequest {
    const message = { ...baseQueryOrderOwnerRequest } as QueryOrderOwnerRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryOrderOwnerRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryOrderOwnerRequest>
  ): QueryOrderOwnerRequest {
    const message = { ...baseQueryOrderOwnerRequest } as QueryOrderOwnerRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryOrderOwnerUidsResponse: object = {};

export const QueryOrderOwnerUidsResponse = {
  encode(
    message: QueryOrderOwnerUidsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.orders !== undefined) {
      Orders.encode(message.orders, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryOrderOwnerUidsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryOrderOwnerUidsResponse,
    } as QueryOrderOwnerUidsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.orders = Orders.decode(reader, reader.uint32());
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryOrderOwnerUidsResponse {
    const message = {
      ...baseQueryOrderOwnerUidsResponse,
    } as QueryOrderOwnerUidsResponse;
    if (object.orders !== undefined && object.orders !== null) {
      message.orders = Orders.fromJSON(object.orders);
    } else {
      message.orders = undefined;
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryOrderOwnerUidsResponse): unknown {
    const obj: any = {};
    message.orders !== undefined &&
      (obj.orders = message.orders ? Orders.toJSON(message.orders) : undefined);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryOrderOwnerUidsResponse>
  ): QueryOrderOwnerUidsResponse {
    const message = {
      ...baseQueryOrderOwnerUidsResponse,
    } as QueryOrderOwnerUidsResponse;
    if (object.orders !== undefined && object.orders !== null) {
      message.orders = Orders.fromPartial(object.orders);
    } else {
      message.orders = undefined;
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryOrderOwnerPairRequest: object = { address: "", pair: "" };

export const QueryOrderOwnerPairRequest = {
  encode(
    message: QueryOrderOwnerPairRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.pair !== "") {
      writer.uint32(18).string(message.pair);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryOrderOwnerPairRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryOrderOwnerPairRequest,
    } as QueryOrderOwnerPairRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.pair = reader.string();
          break;
        case 3:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryOrderOwnerPairRequest {
    const message = {
      ...baseQueryOrderOwnerPairRequest,
    } as QueryOrderOwnerPairRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.pair !== undefined && object.pair !== null) {
      message.pair = String(object.pair);
    } else {
      message.pair = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryOrderOwnerPairRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.pair !== undefined && (obj.pair = message.pair);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryOrderOwnerPairRequest>
  ): QueryOrderOwnerPairRequest {
    const message = {
      ...baseQueryOrderOwnerPairRequest,
    } as QueryOrderOwnerPairRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.pair !== undefined && object.pair !== null) {
      message.pair = object.pair;
    } else {
      message.pair = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryOrderOwnerPairResponse: object = {};

export const QueryOrderOwnerPairResponse = {
  encode(
    message: QueryOrderOwnerPairResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.order) {
      Order.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): QueryOrderOwnerPairResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryOrderOwnerPairResponse,
    } as QueryOrderOwnerPairResponse;
    message.order = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.order.push(Order.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryOrderOwnerPairResponse {
    const message = {
      ...baseQueryOrderOwnerPairResponse,
    } as QueryOrderOwnerPairResponse;
    message.order = [];
    if (object.order !== undefined && object.order !== null) {
      for (const e of object.order) {
        message.order.push(Order.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryOrderOwnerPairResponse): unknown {
    const obj: any = {};
    if (message.order) {
      obj.order = message.order.map((e) => (e ? Order.toJSON(e) : undefined));
    } else {
      obj.order = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryOrderOwnerPairResponse>
  ): QueryOrderOwnerPairResponse {
    const message = {
      ...baseQueryOrderOwnerPairResponse,
    } as QueryOrderOwnerPairResponse;
    message.order = [];
    if (object.order !== undefined && object.order !== null) {
      for (const e of object.order) {
        message.order.push(Order.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryBookRequest: object = { denomA: "", denomB: "", orderType: "" };

export const QueryBookRequest = {
  encode(message: QueryBookRequest, writer: Writer = Writer.create()): Writer {
    if (message.denomA !== "") {
      writer.uint32(10).string(message.denomA);
    }
    if (message.denomB !== "") {
      writer.uint32(18).string(message.denomB);
    }
    if (message.orderType !== "") {
      writer.uint32(26).string(message.orderType);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryBookRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBookRequest } as QueryBookRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.denomA = reader.string();
          break;
        case 2:
          message.denomB = reader.string();
          break;
        case 3:
          message.orderType = reader.string();
          break;
        case 4:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryBookRequest {
    const message = { ...baseQueryBookRequest } as QueryBookRequest;
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
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = String(object.orderType);
    } else {
      message.orderType = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryBookRequest): unknown {
    const obj: any = {};
    message.denomA !== undefined && (obj.denomA = message.denomA);
    message.denomB !== undefined && (obj.denomB = message.denomB);
    message.orderType !== undefined && (obj.orderType = message.orderType);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryBookRequest>): QueryBookRequest {
    const message = { ...baseQueryBookRequest } as QueryBookRequest;
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
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = object.orderType;
    } else {
      message.orderType = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryBookResponse: object = {};

export const QueryBookResponse = {
  encode(message: QueryBookResponse, writer: Writer = Writer.create()): Writer {
    for (const v of message.book) {
      OrderResponse.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryBookResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBookResponse } as QueryBookResponse;
    message.book = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.book.push(OrderResponse.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryBookResponse {
    const message = { ...baseQueryBookResponse } as QueryBookResponse;
    message.book = [];
    if (object.book !== undefined && object.book !== null) {
      for (const e of object.book) {
        message.book.push(OrderResponse.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryBookResponse): unknown {
    const obj: any = {};
    if (message.book) {
      obj.book = message.book.map((e) =>
        e ? OrderResponse.toJSON(e) : undefined
      );
    } else {
      obj.book = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryBookResponse>): QueryBookResponse {
    const message = { ...baseQueryBookResponse } as QueryBookResponse;
    message.book = [];
    if (object.book !== undefined && object.book !== null) {
      for (const e of object.book) {
        message.book.push(OrderResponse.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryBookendsRequest: object = {
  coinA: "",
  coinB: "",
  orderType: "",
  rate: "",
};

export const QueryBookendsRequest = {
  encode(
    message: QueryBookendsRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.coinA !== "") {
      writer.uint32(10).string(message.coinA);
    }
    if (message.coinB !== "") {
      writer.uint32(18).string(message.coinB);
    }
    if (message.orderType !== "") {
      writer.uint32(26).string(message.orderType);
    }
    for (const v of message.rate) {
      writer.uint32(34).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryBookendsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBookendsRequest } as QueryBookendsRequest;
    message.rate = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.coinA = reader.string();
          break;
        case 2:
          message.coinB = reader.string();
          break;
        case 3:
          message.orderType = reader.string();
          break;
        case 4:
          message.rate.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryBookendsRequest {
    const message = { ...baseQueryBookendsRequest } as QueryBookendsRequest;
    message.rate = [];
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
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = String(object.orderType);
    } else {
      message.orderType = "";
    }
    if (object.rate !== undefined && object.rate !== null) {
      for (const e of object.rate) {
        message.rate.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: QueryBookendsRequest): unknown {
    const obj: any = {};
    message.coinA !== undefined && (obj.coinA = message.coinA);
    message.coinB !== undefined && (obj.coinB = message.coinB);
    message.orderType !== undefined && (obj.orderType = message.orderType);
    if (message.rate) {
      obj.rate = message.rate.map((e) => e);
    } else {
      obj.rate = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<QueryBookendsRequest>): QueryBookendsRequest {
    const message = { ...baseQueryBookendsRequest } as QueryBookendsRequest;
    message.rate = [];
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
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = object.orderType;
    } else {
      message.orderType = "";
    }
    if (object.rate !== undefined && object.rate !== null) {
      for (const e of object.rate) {
        message.rate.push(e);
      }
    }
    return message;
  },
};

const baseQueryBookendsResponse: object = {
  coinA: "",
  coinB: "",
  orderType: "",
  rate: "",
  prev: 0,
  next: 0,
};

export const QueryBookendsResponse = {
  encode(
    message: QueryBookendsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.coinA !== "") {
      writer.uint32(10).string(message.coinA);
    }
    if (message.coinB !== "") {
      writer.uint32(18).string(message.coinB);
    }
    if (message.orderType !== "") {
      writer.uint32(26).string(message.orderType);
    }
    for (const v of message.rate) {
      writer.uint32(34).string(v!);
    }
    if (message.prev !== 0) {
      writer.uint32(40).uint64(message.prev);
    }
    if (message.next !== 0) {
      writer.uint32(48).uint64(message.next);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryBookendsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryBookendsResponse } as QueryBookendsResponse;
    message.rate = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.coinA = reader.string();
          break;
        case 2:
          message.coinB = reader.string();
          break;
        case 3:
          message.orderType = reader.string();
          break;
        case 4:
          message.rate.push(reader.string());
          break;
        case 5:
          message.prev = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.next = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryBookendsResponse {
    const message = { ...baseQueryBookendsResponse } as QueryBookendsResponse;
    message.rate = [];
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
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = String(object.orderType);
    } else {
      message.orderType = "";
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
    return message;
  },

  toJSON(message: QueryBookendsResponse): unknown {
    const obj: any = {};
    message.coinA !== undefined && (obj.coinA = message.coinA);
    message.coinB !== undefined && (obj.coinB = message.coinB);
    message.orderType !== undefined && (obj.orderType = message.orderType);
    if (message.rate) {
      obj.rate = message.rate.map((e) => e);
    } else {
      obj.rate = [];
    }
    message.prev !== undefined && (obj.prev = message.prev);
    message.next !== undefined && (obj.next = message.next);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryBookendsResponse>
  ): QueryBookendsResponse {
    const message = { ...baseQueryBookendsResponse } as QueryBookendsResponse;
    message.rate = [];
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
    if (object.orderType !== undefined && object.orderType !== null) {
      message.orderType = object.orderType;
    } else {
      message.orderType = "";
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
    return message;
  },
};

const baseQueryHistoryRequest: object = { pair: "", length: "" };

export const QueryHistoryRequest = {
  encode(
    message: QueryHistoryRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pair !== "") {
      writer.uint32(10).string(message.pair);
    }
    if (message.length !== "") {
      writer.uint32(18).string(message.length);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryHistoryRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryHistoryRequest } as QueryHistoryRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pair = reader.string();
          break;
        case 2:
          message.length = reader.string();
          break;
        case 3:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryHistoryRequest {
    const message = { ...baseQueryHistoryRequest } as QueryHistoryRequest;
    if (object.pair !== undefined && object.pair !== null) {
      message.pair = String(object.pair);
    } else {
      message.pair = "";
    }
    if (object.length !== undefined && object.length !== null) {
      message.length = String(object.length);
    } else {
      message.length = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryHistoryRequest): unknown {
    const obj: any = {};
    message.pair !== undefined && (obj.pair = message.pair);
    message.length !== undefined && (obj.length = message.length);
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryHistoryRequest>): QueryHistoryRequest {
    const message = { ...baseQueryHistoryRequest } as QueryHistoryRequest;
    if (object.pair !== undefined && object.pair !== null) {
      message.pair = object.pair;
    } else {
      message.pair = "";
    }
    if (object.length !== undefined && object.length !== null) {
      message.length = object.length;
    } else {
      message.length = "";
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryHistoryResponse: object = {};

export const QueryHistoryResponse = {
  encode(
    message: QueryHistoryResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.history) {
      OrderResponse.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryHistoryResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryHistoryResponse } as QueryHistoryResponse;
    message.history = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.history.push(OrderResponse.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryHistoryResponse {
    const message = { ...baseQueryHistoryResponse } as QueryHistoryResponse;
    message.history = [];
    if (object.history !== undefined && object.history !== null) {
      for (const e of object.history) {
        message.history.push(OrderResponse.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryHistoryResponse): unknown {
    const obj: any = {};
    if (message.history) {
      obj.history = message.history.map((e) =>
        e ? OrderResponse.toJSON(e) : undefined
      );
    } else {
      obj.history = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryHistoryResponse>): QueryHistoryResponse {
    const message = { ...baseQueryHistoryResponse } as QueryHistoryResponse;
    message.history = [];
    if (object.history !== undefined && object.history !== null) {
      for (const e of object.history) {
        message.history.push(OrderResponse.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Pool by index. */
  Pool(request: QueryGetPoolRequest): Promise<QueryGetPoolResponse>;
  /** Queries a list of Pool items. */
  PoolAll(request: QueryAllPoolRequest): Promise<QueryAllPoolResponse>;
  /** Queries a Drop by index. */
  Drop(request: QueryDropRequest): Promise<QueryDropResponse>;
  /** Queries a list of Drop items. */
  DropAll(request: QueryAllDropRequest): Promise<QueryDropsResponse>;
  /** Queries a Member by index. */
  Member(request: QueryGetMemberRequest): Promise<QueryGetMemberResponse>;
  /** Queries a list of Member items. */
  MemberAll(request: QueryAllMemberRequest): Promise<QueryAllMemberResponse>;
  /** Queries a Burnings by index. */
  Burnings(request: QueryGetBurningsRequest): Promise<QueryGetBurningsResponse>;
  /** Queries a list of Burnings items. */
  BurningsAll(
    request: QueryAllBurningsRequest
  ): Promise<QueryAllBurningsResponse>;
  /** Queries a Order by index. */
  Order(request: QueryGetOrderRequest): Promise<QueryGetOrderResponse>;
  /** Queries a list of Order items. */
  OrderAll(request: QueryAllOrderRequest): Promise<QueryOrdersResponse>;
  /** Queries a list of Order items. */
  OrderOwner(request: QueryOrderOwnerRequest): Promise<QueryOrdersResponse>;
  /** Queries a list of Order items. */
  OrderOwnerUids(
    request: QueryOrderOwnerRequest
  ): Promise<QueryOrderOwnerUidsResponse>;
  /** Queries a list of Book items. */
  Book(request: QueryBookRequest): Promise<QueryBookResponse>;
  /** Queries a list of Bookends items. */
  Bookends(request: QueryBookendsRequest): Promise<QueryBookendsResponse>;
  /** Queries pool trade history. */
  History(request: QueryHistoryRequest): Promise<QueryHistoryResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pendulumlabs.market.market.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Pool(request: QueryGetPoolRequest): Promise<QueryGetPoolResponse> {
    const data = QueryGetPoolRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pendulumlabs.market.market.Query",
      "Pool",
      data
    );
    return promise.then((data) =>
      QueryGetPoolResponse.decode(new Reader(data))
    );
  }

  PoolAll(request: QueryAllPoolRequest): Promise<QueryAllPoolResponse> {
    const data = QueryAllPoolRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pendulumlabs.market.market.Query",
      "PoolAll",
      data
    );
    return promise.then((data) =>
      QueryAllPoolResponse.decode(new Reader(data))
    );
  }

  Drop(request: QueryDropRequest): Promise<QueryDropResponse> {
    const data = QueryDropRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pendulumlabs.market.market.Query",
      "Drop",
      data
    );
    return promise.then((data) => QueryDropResponse.decode(new Reader(data)));
  }

  DropAll(request: QueryAllDropRequest): Promise<QueryDropsResponse> {
    const data = QueryAllDropRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pendulumlabs.market.market.Query",
      "DropAll",
      data
    );
    return promise.then((data) => QueryDropsResponse.decode(new Reader(data)));
  }

  Member(request: QueryGetMemberRequest): Promise<QueryGetMemberResponse> {
    const data = QueryGetMemberRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pendulumlabs.market.market.Query",
      "Member",
      data
    );
    return promise.then((data) =>
      QueryGetMemberResponse.decode(new Reader(data))
    );
  }

  MemberAll(request: QueryAllMemberRequest): Promise<QueryAllMemberResponse> {
    const data = QueryAllMemberRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pendulumlabs.market.market.Query",
      "MemberAll",
      data
    );
    return promise.then((data) =>
      QueryAllMemberResponse.decode(new Reader(data))
    );
  }

  Burnings(
    request: QueryGetBurningsRequest
  ): Promise<QueryGetBurningsResponse> {
    const data = QueryGetBurningsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pendulumlabs.market.market.Query",
      "Burnings",
      data
    );
    return promise.then((data) =>
      QueryGetBurningsResponse.decode(new Reader(data))
    );
  }

  BurningsAll(
    request: QueryAllBurningsRequest
  ): Promise<QueryAllBurningsResponse> {
    const data = QueryAllBurningsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pendulumlabs.market.market.Query",
      "BurningsAll",
      data
    );
    return promise.then((data) =>
      QueryAllBurningsResponse.decode(new Reader(data))
    );
  }

  Order(request: QueryGetOrderRequest): Promise<QueryGetOrderResponse> {
    const data = QueryGetOrderRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pendulumlabs.market.market.Query",
      "Order",
      data
    );
    return promise.then((data) =>
      QueryGetOrderResponse.decode(new Reader(data))
    );
  }

  OrderAll(request: QueryAllOrderRequest): Promise<QueryOrdersResponse> {
    const data = QueryAllOrderRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pendulumlabs.market.market.Query",
      "OrderAll",
      data
    );
    return promise.then((data) => QueryOrdersResponse.decode(new Reader(data)));
  }

  OrderOwner(request: QueryOrderOwnerRequest): Promise<QueryOrdersResponse> {
    const data = QueryOrderOwnerRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pendulumlabs.market.market.Query",
      "OrderOwner",
      data
    );
    return promise.then((data) => QueryOrdersResponse.decode(new Reader(data)));
  }

  OrderOwnerUids(
    request: QueryOrderOwnerRequest
  ): Promise<QueryOrderOwnerUidsResponse> {
    const data = QueryOrderOwnerRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pendulumlabs.market.market.Query",
      "OrderOwnerUids",
      data
    );
    return promise.then((data) =>
      QueryOrderOwnerUidsResponse.decode(new Reader(data))
    );
  }

  Book(request: QueryBookRequest): Promise<QueryBookResponse> {
    const data = QueryBookRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pendulumlabs.market.market.Query",
      "Book",
      data
    );
    return promise.then((data) => QueryBookResponse.decode(new Reader(data)));
  }

  Bookends(request: QueryBookendsRequest): Promise<QueryBookendsResponse> {
    const data = QueryBookendsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pendulumlabs.market.market.Query",
      "Bookends",
      data
    );
    return promise.then((data) =>
      QueryBookendsResponse.decode(new Reader(data))
    );
  }

  History(request: QueryHistoryRequest): Promise<QueryHistoryResponse> {
    const data = QueryHistoryRequest.encode(request).finish();
    const promise = this.rpc.request(
      "pendulumlabs.market.market.Query",
      "History",
      data
    );
    return promise.then((data) =>
      QueryHistoryResponse.decode(new Reader(data))
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
