import * as jspb from "google-protobuf"

export class EntryRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EntryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: EntryRequest): EntryRequest.AsObject;
  static serializeBinaryToWriter(message: EntryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EntryRequest;
  static deserializeBinaryFromReader(message: EntryRequest, reader: jspb.BinaryReader): EntryRequest;
}

export namespace EntryRequest {
  export type AsObject = {
  }
}

export class ChatMessageRequest extends jspb.Message {
  getContent(): string;
  setContent(value: string): ChatMessageRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChatMessageRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ChatMessageRequest): ChatMessageRequest.AsObject;
  static serializeBinaryToWriter(message: ChatMessageRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChatMessageRequest;
  static deserializeBinaryFromReader(message: ChatMessageRequest, reader: jspb.BinaryReader): ChatMessageRequest;
}

export namespace ChatMessageRequest {
  export type AsObject = {
    content: string,
  }
}

export class BroadcastResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): BroadcastResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BroadcastResponse.AsObject;
  static toObject(includeInstance: boolean, msg: BroadcastResponse): BroadcastResponse.AsObject;
  static serializeBinaryToWriter(message: BroadcastResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BroadcastResponse;
  static deserializeBinaryFromReader(message: BroadcastResponse, reader: jspb.BinaryReader): BroadcastResponse;
}

export namespace BroadcastResponse {
  export type AsObject = {
    message: string,
  }
}

export class ChatMessageResponse extends jspb.Message {
  getName(): string;
  setName(value: string): ChatMessageResponse;

  getContent(): string;
  setContent(value: string): ChatMessageResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChatMessageResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ChatMessageResponse): ChatMessageResponse.AsObject;
  static serializeBinaryToWriter(message: ChatMessageResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChatMessageResponse;
  static deserializeBinaryFromReader(message: ChatMessageResponse, reader: jspb.BinaryReader): ChatMessageResponse;
}

export namespace ChatMessageResponse {
  export type AsObject = {
    name: string,
    content: string,
  }
}

