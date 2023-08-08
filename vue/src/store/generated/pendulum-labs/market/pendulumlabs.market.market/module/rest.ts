/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface MarketBurnings {
  denom?: string;
  amount?: string;
}

export interface MarketDrop {
  /** @format uint64 */
  uid?: string;
  owner?: string;
  pair?: string;
  drops?: string;
  product?: string;
  active?: boolean;
}

export interface MarketLeader {
  address?: string;
  drops?: string;
}

export interface MarketMember {
  pair?: string;
  denomA?: string;
  denomB?: string;
  balance?: string;
  previous?: string;

  /** @format uint64 */
  limit?: string;

  /** @format uint64 */
  stop?: string;
}

export type MarketMsgCancelOrderResponse = object;

export type MarketMsgCreateDropResponse = object;

export type MarketMsgCreateOrderResponse = object;

export type MarketMsgCreatePoolResponse = object;

export type MarketMsgMarketOrderResponse = object;

export type MarketMsgRedeemDropResponse = object;

export interface MarketOrder {
  /** @format uint64 */
  uid?: string;
  owner?: string;
  status?: string;
  orderType?: string;
  denomAsk?: string;
  denomBid?: string;
  amount?: string;
  rate?: string[];

  /** @format uint64 */
  prev?: string;

  /** @format uint64 */
  next?: string;

  /** @format int64 */
  beg_time?: string;

  /** @format int64 */
  end_time?: string;
}

export interface MarketOrderResponse {
  /** @format uint64 */
  uid?: string;
  owner?: string;
  status?: string;
  orderType?: string;
  denomAsk?: string;
  denomBid?: string;
  amount?: string;
  rate?: string[];

  /** @format uint64 */
  prev?: string;

  /** @format uint64 */
  next?: string;

  /** @format int64 */
  beg_time?: string;

  /** @format int64 */
  end_time?: string;
}

export interface MarketOrders {
  uids?: string[];
}

/**
 * Params defines the parameters for the module.
 */
export interface MarketParams {
  earn_rates?: string;
  burn_rate?: string;
  burn_coin?: string;
  market_fee?: string;
}

export interface MarketPool {
  pair?: string;
  denom1?: string;
  denom2?: string;
  leaders?: MarketLeader[];
  drops?: string;

  /** @format uint64 */
  history?: string;
}

export interface MarketQueryAllBurningsResponse {
  burnings?: MarketBurnings[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface MarketQueryAllDropResponse {
  drop?: MarketDrop[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface MarketQueryAllMemberResponse {
  member?: MarketMember[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface MarketQueryAllOrderResponse {
  order?: MarketOrder[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface MarketQueryAllPoolResponse {
  pool?: MarketPool[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface MarketQueryBookResponse {
  book?: MarketOrderResponse[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface MarketQueryBookendsResponse {
  coinA?: string;
  coinB?: string;
  orderType?: string;
  rate?: string[];

  /** @format uint64 */
  prev?: string;

  /** @format uint64 */
  next?: string;
}

export interface MarketQueryGetBurningsResponse {
  burnings?: MarketBurnings;
}

export interface MarketQueryGetDropResponse {
  drop?: MarketDrop;
}

export interface MarketQueryGetMemberResponse {
  member?: MarketMember;
}

export interface MarketQueryGetOrderResponse {
  order?: MarketOrder;
}

export interface MarketQueryGetPoolResponse {
  pool?: MarketPool;
}

export interface MarketQueryHistoryResponse {
  history?: MarketOrderResponse[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface MarketQueryOrderOwnerPairResponse {
  order?: MarketOrder[];

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

export interface MarketQueryOrderOwnerResponse {
  orders?: MarketOrders;

  /**
   * PageResponse is to be embedded in gRPC response messages where the
   * corresponding request message has used PageRequest.
   *
   *  message SomeResponse {
   *          repeated Bar results = 1;
   *          PageResponse page = 2;
   *  }
   */
  pagination?: V1Beta1PageResponse;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface MarketQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: MarketParams;
}

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

/**
* message SomeRequest {
         Foo some_parameter = 1;
         PageRequest pagination = 2;
 }
*/
export interface V1Beta1PageRequest {
  /**
   * key is a value returned in PageResponse.next_key to begin
   * querying the next page most efficiently. Only one of offset or key
   * should be set.
   * @format byte
   */
  key?: string;

  /**
   * offset is a numeric offset that can be used when key is unavailable.
   * It is less efficient than using key. Only one of offset or key should
   * be set.
   * @format uint64
   */
  offset?: string;

  /**
   * limit is the total number of results to be returned in the result page.
   * If left empty it will default to a value to be set by each app.
   * @format uint64
   */
  limit?: string;

  /**
   * count_total is set to true  to indicate that the result set should include
   * a count of the total number of items available for pagination in UIs.
   * count_total is only respected when offset is used. It is ignored when key
   * is set.
   */
  count_total?: boolean;

  /**
   * reverse is set to true if results are to be returned in the descending order.
   *
   * Since: cosmos-sdk 0.43
   */
  reverse?: boolean;
}

/**
* PageResponse is to be embedded in gRPC response messages where the
corresponding request message has used PageRequest.

 message SomeResponse {
         repeated Bar results = 1;
         PageResponse page = 2;
 }
*/
export interface V1Beta1PageResponse {
  /** @format byte */
  next_key?: string;

  /** @format uint64 */
  total?: string;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: keyof Omit<Body, "body" | "bodyUsed">;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (securityData: SecurityDataType) => RequestParams | void;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown> extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "";
  private securityData: SecurityDataType = null as any;
  private securityWorker: null | ApiConfig<SecurityDataType>["securityWorker"] = null;
  private abortControllers = new Map<CancelToken, AbortController>();

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType) => {
    this.securityData = data;
  };

  private addQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];

    return (
      encodeURIComponent(key) +
      "=" +
      encodeURIComponent(Array.isArray(value) ? value.join(",") : typeof value === "number" ? value : `${value}`)
    );
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter((key) => "undefined" !== typeof query[key]);
    return keys
      .map((key) =>
        typeof query[key] === "object" && !Array.isArray(query[key])
          ? this.toQueryString(query[key] as QueryParamsType)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string") ? JSON.stringify(input) : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((data, key) => {
        data.append(key, input[key]);
        return data;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  private mergeRequestParams(params1: RequestParams, params2?: RequestParams): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createAbortSignal = (cancelToken: CancelToken): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format = "json",
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams = (secure && this.securityWorker && this.securityWorker(this.securityData)) || {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];

    return fetch(`${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`, {
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      signal: cancelToken ? this.createAbortSignal(cancelToken) : void 0,
      body: typeof body === "undefined" || body === null ? null : payloadFormatter(body),
    }).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = (null as unknown) as T;
      r.error = (null as unknown) as E;

      const data = await response[format]()
        .then((data) => {
          if (r.ok) {
            r.data = data;
          } else {
            r.error = data;
          }
          return r;
        })
        .catch((e) => {
          r.error = e;
          return r;
        });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title market/burnings.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryBook
   * @summary Queries a list of Book items.
   * @request GET:/pendulum-labs/market/market/book/{denomA}/{denomB}/{orderType}
   */
  queryBook = (
    denomA: string,
    denomB: string,
    orderType: string,
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<MarketQueryBookResponse, RpcStatus>({
      path: `/pendulum-labs/market/market/book/${denomA}/${denomB}/${orderType}`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryBookends
   * @summary Queries a list of Bookends items.
   * @request GET:/pendulum-labs/market/market/bookends/{coinA}/{coinB}/{orderType}/{rate}
   */
  queryBookends = (coinA: string, coinB: string, orderType: string, rate: string[], params: RequestParams = {}) =>
    this.request<MarketQueryBookendsResponse, RpcStatus>({
      path: `/pendulum-labs/market/market/bookends/${coinA}/${coinB}/${orderType}/${rate}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryBurningsAll
   * @summary Queries a list of Burnings items.
   * @request GET:/pendulum-labs/market/market/burnings
   */
  queryBurningsAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<MarketQueryAllBurningsResponse, RpcStatus>({
      path: `/pendulum-labs/market/market/burnings`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryBurnings
   * @summary Queries a Burnings by index.
   * @request GET:/pendulum-labs/market/market/burnings/{denom}
   */
  queryBurnings = (denom: string, params: RequestParams = {}) =>
    this.request<MarketQueryGetBurningsResponse, RpcStatus>({
      path: `/pendulum-labs/market/market/burnings/${denom}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryDropAll
   * @summary Queries a list of Drop items.
   * @request GET:/pendulum-labs/market/market/drop
   */
  queryDropAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<MarketQueryAllDropResponse, RpcStatus>({
      path: `/pendulum-labs/market/market/drop`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryDrop
   * @summary Queries a Drop by index.
   * @request GET:/pendulum-labs/market/market/drop/{uid}
   */
  queryDrop = (uid: string, params: RequestParams = {}) =>
    this.request<MarketQueryGetDropResponse, RpcStatus>({
      path: `/pendulum-labs/market/market/drop/${uid}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryHistory
   * @summary Queries pool trade history.
   * @request GET:/pendulum-labs/market/market/history/{pair}
   */
  queryHistory = (
    pair: string,
    query?: {
      length?: string;
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<MarketQueryHistoryResponse, RpcStatus>({
      path: `/pendulum-labs/market/market/history/${pair}`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryMemberAll
   * @summary Queries a list of Member items.
   * @request GET:/pendulum-labs/market/market/member
   */
  queryMemberAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<MarketQueryAllMemberResponse, RpcStatus>({
      path: `/pendulum-labs/market/market/member`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryMember
   * @summary Queries a Member by index.
   * @request GET:/pendulum-labs/market/market/member/{denomA}/{denomB}
   */
  queryMember = (denomA: string, denomB: string, params: RequestParams = {}) =>
    this.request<MarketQueryGetMemberResponse, RpcStatus>({
      path: `/pendulum-labs/market/market/member/${denomA}/${denomB}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryOrderAll
   * @summary Queries a list of Order items.
   * @request GET:/pendulum-labs/market/market/order
   */
  queryOrderAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<MarketQueryAllOrderResponse, RpcStatus>({
      path: `/pendulum-labs/market/market/order`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryOrderOwner
   * @summary Queries a list of Order items.
   * @request GET:/pendulum-labs/market/market/order/{address}
   */
  queryOrderOwner = (
    address: string,
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<MarketQueryOrderOwnerResponse, RpcStatus>({
      path: `/pendulum-labs/market/market/order/${address}`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryOrderOwnerPair
   * @summary Queries a list of Order items.
   * @request GET:/pendulum-labs/market/market/order/{address}/{pair}
   */
  queryOrderOwnerPair = (
    address: string,
    pair: string,
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<MarketQueryOrderOwnerPairResponse, RpcStatus>({
      path: `/pendulum-labs/market/market/order/${address}/${pair}`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryOrder
   * @summary Queries a Order by index.
   * @request GET:/pendulum-labs/market/market/order/{uid}
   */
  queryOrder = (uid: string, params: RequestParams = {}) =>
    this.request<MarketQueryGetOrderResponse, RpcStatus>({
      path: `/pendulum-labs/market/market/order/${uid}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/pendulum-labs/market/market/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<MarketQueryParamsResponse, RpcStatus>({
      path: `/pendulum-labs/market/market/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPoolAll
   * @summary Queries a list of Pool items.
   * @request GET:/pendulum-labs/market/market/pool
   */
  queryPoolAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<MarketQueryAllPoolResponse, RpcStatus>({
      path: `/pendulum-labs/market/market/pool`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPool
   * @summary Queries a Pool by index.
   * @request GET:/pendulum-labs/market/market/pool/{pair}
   */
  queryPool = (pair: string, params: RequestParams = {}) =>
    this.request<MarketQueryGetPoolResponse, RpcStatus>({
      path: `/pendulum-labs/market/market/pool/${pair}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
