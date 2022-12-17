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

export interface MarketAsset {
  active?: boolean;
  owner?: string;
  assetType?: string;

  /** @format uint64 */
  uid?: string;
}

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
  sum?: string;
  active?: boolean;
}

export interface MarketMember {
  pair?: string;
  denomA?: string;
  denomB?: string;
  balance?: string;

  /** @format uint64 */
  limit?: string;

  /** @format uint64 */
  stop?: string;
}

export type MarketMsgCancelOrderResponse = object;

export type MarketMsgCreateDropResponse = object;

export type MarketMsgCreateOrderResponse = object;

export type MarketMsgCreatePoolResponse = object;

export type MarketMsgRedeemDropResponse = object;

export interface MarketOrder {
  /** @format uint64 */
  uid?: string;
  owner?: string;
  active?: boolean;
  orderType?: string;
  denomAsk?: string;
  denomBid?: string;
  amount?: string;
  rate?: string[];

  /** @format uint64 */
  prev?: string;

  /** @format uint64 */
  next?: string;
}

/**
 * Params defines the parameters for the module.
 */
export interface MarketParams {
  earn_rate?: string[];
  burn_rate?: string[];
}

export interface MarketPool {
  pair?: string;
  denom1?: string;
  denom2?: string;
  leader?: string;
  drops?: string;
}

export interface MarketQueryAllAssetResponse {
  asset?: MarketAsset[];

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

export interface MarketQueryGetAssetResponse {
  asset?: MarketAsset;
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
 * @title market/asset.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryAssetAll
   * @summary Queries a list of Asset items.
   * @request GET:/onomyprotocol/market/market/asset
   */
  queryAssetAll = (
    query?: {
      "pagination.key"?: string;
      "pagination.offset"?: string;
      "pagination.limit"?: string;
      "pagination.count_total"?: boolean;
      "pagination.reverse"?: boolean;
    },
    params: RequestParams = {},
  ) =>
    this.request<MarketQueryAllAssetResponse, RpcStatus>({
      path: `/onomyprotocol/market/market/asset`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryAsset
   * @summary Queries a Asset by index.
   * @request GET:/onomyprotocol/market/market/asset/{active}/{owner}/{assetType}
   */
  queryAsset = (active: boolean, owner: string, assetType: string, params: RequestParams = {}) =>
    this.request<MarketQueryGetAssetResponse, RpcStatus>({
      path: `/onomyprotocol/market/market/asset/${active}/${owner}/${assetType}`,
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
   * @request GET:/onomyprotocol/market/market/burnings
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
      path: `/onomyprotocol/market/market/burnings`,
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
   * @request GET:/onomyprotocol/market/market/burnings/{denom}
   */
  queryBurnings = (denom: string, params: RequestParams = {}) =>
    this.request<MarketQueryGetBurningsResponse, RpcStatus>({
      path: `/onomyprotocol/market/market/burnings/${denom}`,
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
   * @request GET:/onomyprotocol/market/market/drop
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
      path: `/onomyprotocol/market/market/drop`,
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
   * @request GET:/onomyprotocol/market/market/drop/{uid}/{owner}/{pair}
   */
  queryDrop = (uid: string, owner: string, pair: string, params: RequestParams = {}) =>
    this.request<MarketQueryGetDropResponse, RpcStatus>({
      path: `/onomyprotocol/market/market/drop/${uid}/${owner}/${pair}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryMemberAll
   * @summary Queries a list of Member items.
   * @request GET:/onomyprotocol/market/market/member
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
      path: `/onomyprotocol/market/market/member`,
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
   * @request GET:/onomyprotocol/market/market/member/{pair}/{denomA}/{denomB}
   */
  queryMember = (pair: string, denomA: string, denomB: string, params: RequestParams = {}) =>
    this.request<MarketQueryGetMemberResponse, RpcStatus>({
      path: `/onomyprotocol/market/market/member/${pair}/${denomA}/${denomB}`,
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
   * @request GET:/onomyprotocol/market/market/order
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
      path: `/onomyprotocol/market/market/order`,
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
   * @request GET:/onomyprotocol/market/market/order/{uid}
   */
  queryOrder = (
    uid: string,
    query?: { owner?: string; active?: boolean; orderType?: string; denomAsk?: string; denomBid?: string },
    params: RequestParams = {},
  ) =>
    this.request<MarketQueryGetOrderResponse, RpcStatus>({
      path: `/onomyprotocol/market/market/order/${uid}`,
      method: "GET",
      query: query,
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/onomyprotocol/market/market/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<MarketQueryParamsResponse, RpcStatus>({
      path: `/onomyprotocol/market/market/params`,
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
   * @request GET:/onomyprotocol/market/market/pool
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
      path: `/onomyprotocol/market/market/pool`,
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
   * @request GET:/onomyprotocol/market/market/pool/{pair}/{denom1}/{denom2}/{leader}
   */
  queryPool = (pair: string, denom1: string, denom2: string, leader: string, params: RequestParams = {}) =>
    this.request<MarketQueryGetPoolResponse, RpcStatus>({
      path: `/onomyprotocol/market/market/pool/${pair}/${denom1}/${denom2}/${leader}`,
      method: "GET",
      format: "json",
      ...params,
    });
}
