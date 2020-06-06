/**
 * @fileoverview gRPC-Web generated client stub for plaza
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import {
  BroadcastResponse,
  ChatMessageRequest,
  ChatMessageResponse,
  EntryRequest} from './plaza_pb';

export class PlazaClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: string; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoEntry = new grpcWeb.AbstractClientBase.MethodInfo(
    ChatMessageResponse,
    (request: EntryRequest) => {
      return request.serializeBinary();
    },
    ChatMessageResponse.deserializeBinary
  );

  entry(
    request: EntryRequest,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/plaza.Plaza/Entry',
      request,
      metadata || {},
      this.methodInfoEntry);
  }

  methodInfoBroadcast = new grpcWeb.AbstractClientBase.MethodInfo(
    BroadcastResponse,
    (request: ChatMessageRequest) => {
      return request.serializeBinary();
    },
    BroadcastResponse.deserializeBinary
  );

  broadcast(
    request: ChatMessageRequest,
    metadata: grpcWeb.Metadata | null): Promise<BroadcastResponse>;

  broadcast(
    request: ChatMessageRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.Error,
               response: BroadcastResponse) => void): grpcWeb.ClientReadableStream<BroadcastResponse>;

  broadcast(
    request: ChatMessageRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.Error,
               response: BroadcastResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/plaza.Plaza/Broadcast',
        request,
        metadata || {},
        this.methodInfoBroadcast,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/plaza.Plaza/Broadcast',
    request,
    metadata || {},
    this.methodInfoBroadcast);
  }

}

