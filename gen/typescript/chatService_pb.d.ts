import * as jspb from "google-protobuf"

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';

export class User extends jspb.Message {
  getId(): string;
  setId(value: string): User;

  getName(): string;
  setName(value: string): User;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): User;
  hasCreatedAt(): boolean;
  clearCreatedAt(): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): User.AsObject;
  static toObject(includeInstance: boolean, msg: User): User.AsObject;
  static serializeBinaryToWriter(message: User, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): User;
  static deserializeBinaryFromReader(message: User, reader: jspb.BinaryReader): User;
}

export namespace User {
  export type AsObject = {
    id: string,
    name: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class Room extends jspb.Message {
  getId(): string;
  setId(value: string): Room;

  getTitle(): string;
  setTitle(value: string): Room;

  getDescription(): string;
  setDescription(value: string): Room;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): Room;
  hasCreatedAt(): boolean;
  clearCreatedAt(): Room;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Room.AsObject;
  static toObject(includeInstance: boolean, msg: Room): Room.AsObject;
  static serializeBinaryToWriter(message: Room, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Room;
  static deserializeBinaryFromReader(message: Room, reader: jspb.BinaryReader): Room;
}

export namespace Room {
  export type AsObject = {
    id: string,
    title: string,
    description: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

export class EntryRequest extends jspb.Message {
  getRoomId(): string;
  setRoomId(value: string): EntryRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EntryRequest.AsObject;
  static toObject(includeInstance: boolean, msg: EntryRequest): EntryRequest.AsObject;
  static serializeBinaryToWriter(message: EntryRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EntryRequest;
  static deserializeBinaryFromReader(message: EntryRequest, reader: jspb.BinaryReader): EntryRequest;
}

export namespace EntryRequest {
  export type AsObject = {
    roomId: string,
  }
}

export class ChatMessageRequest extends jspb.Message {
  getRoomId(): string;
  setRoomId(value: string): ChatMessageRequest;

  getMessage(): string;
  setMessage(value: string): ChatMessageRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChatMessageRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ChatMessageRequest): ChatMessageRequest.AsObject;
  static serializeBinaryToWriter(message: ChatMessageRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChatMessageRequest;
  static deserializeBinaryFromReader(message: ChatMessageRequest, reader: jspb.BinaryReader): ChatMessageRequest;
}

export namespace ChatMessageRequest {
  export type AsObject = {
    roomId: string,
    message: string,
  }
}

export class BroadcastResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): BroadcastResponse;

  getRoom(): Room | undefined;
  setRoom(value?: Room): BroadcastResponse;
  hasRoom(): boolean;
  clearRoom(): BroadcastResponse;

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
    room?: Room.AsObject,
  }
}

export class ChatMessageResponse extends jspb.Message {
  getMessageType(): ChatMessageResponse.MessageType;
  setMessageType(value: ChatMessageResponse.MessageType): ChatMessageResponse;

  getSystemJoinNewUser(): System_JoinNewUser | undefined;
  setSystemJoinNewUser(value?: System_JoinNewUser): ChatMessageResponse;
  hasSystemJoinNewUser(): boolean;
  clearSystemJoinNewUser(): ChatMessageResponse;

  getSystemLeaveUser(): System_LeaveUser | undefined;
  setSystemLeaveUser(value?: System_LeaveUser): ChatMessageResponse;
  hasSystemLeaveUser(): boolean;
  clearSystemLeaveUser(): ChatMessageResponse;

  getCommonMessage(): CommonMessage | undefined;
  setCommonMessage(value?: CommonMessage): ChatMessageResponse;
  hasCommonMessage(): boolean;
  clearCommonMessage(): ChatMessageResponse;

  getPayloadCase(): ChatMessageResponse.PayloadCase;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ChatMessageResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ChatMessageResponse): ChatMessageResponse.AsObject;
  static serializeBinaryToWriter(message: ChatMessageResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ChatMessageResponse;
  static deserializeBinaryFromReader(message: ChatMessageResponse, reader: jspb.BinaryReader): ChatMessageResponse;
}

export namespace ChatMessageResponse {
  export type AsObject = {
    messageType: ChatMessageResponse.MessageType,
    systemJoinNewUser?: System_JoinNewUser.AsObject,
    systemLeaveUser?: System_LeaveUser.AsObject,
    commonMessage?: CommonMessage.AsObject,
  }

  export enum MessageType { 
    SYSTEM_JOIN_NEW_USER = 0,
    SYSTEM_LEAVE_USER = 1,
    COMMON_MESSAGE = 2,
  }

  export enum PayloadCase { 
    PAYLOAD_NOT_SET = 0,
    SYSTEM_JOIN_NEW_USER = 2,
    SYSTEM_LEAVE_USER = 3,
    COMMON_MESSAGE = 4,
  }
}

export class System_JoinNewUser extends jspb.Message {
  getUser(): User | undefined;
  setUser(value?: User): System_JoinNewUser;
  hasUser(): boolean;
  clearUser(): System_JoinNewUser;

  getRoom(): Room | undefined;
  setRoom(value?: Room): System_JoinNewUser;
  hasRoom(): boolean;
  clearRoom(): System_JoinNewUser;

  getMessage(): string;
  setMessage(value: string): System_JoinNewUser;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): System_JoinNewUser;
  hasCreatedAt(): boolean;
  clearCreatedAt(): System_JoinNewUser;

  getUserListList(): Array<User>;
  setUserListList(value: Array<User>): System_JoinNewUser;
  clearUserListList(): System_JoinNewUser;
  addUserList(value?: User, index?: number): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): System_JoinNewUser.AsObject;
  static toObject(includeInstance: boolean, msg: System_JoinNewUser): System_JoinNewUser.AsObject;
  static serializeBinaryToWriter(message: System_JoinNewUser, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): System_JoinNewUser;
  static deserializeBinaryFromReader(message: System_JoinNewUser, reader: jspb.BinaryReader): System_JoinNewUser;
}

export namespace System_JoinNewUser {
  export type AsObject = {
    user?: User.AsObject,
    room?: Room.AsObject,
    message: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    userListList: Array<User.AsObject>,
  }
}

export class System_LeaveUser extends jspb.Message {
  getUser(): User | undefined;
  setUser(value?: User): System_LeaveUser;
  hasUser(): boolean;
  clearUser(): System_LeaveUser;

  getRoom(): Room | undefined;
  setRoom(value?: Room): System_LeaveUser;
  hasRoom(): boolean;
  clearRoom(): System_LeaveUser;

  getMessage(): string;
  setMessage(value: string): System_LeaveUser;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): System_LeaveUser;
  hasCreatedAt(): boolean;
  clearCreatedAt(): System_LeaveUser;

  getUserListList(): Array<User>;
  setUserListList(value: Array<User>): System_LeaveUser;
  clearUserListList(): System_LeaveUser;
  addUserList(value?: User, index?: number): User;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): System_LeaveUser.AsObject;
  static toObject(includeInstance: boolean, msg: System_LeaveUser): System_LeaveUser.AsObject;
  static serializeBinaryToWriter(message: System_LeaveUser, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): System_LeaveUser;
  static deserializeBinaryFromReader(message: System_LeaveUser, reader: jspb.BinaryReader): System_LeaveUser;
}

export namespace System_LeaveUser {
  export type AsObject = {
    user?: User.AsObject,
    room?: Room.AsObject,
    message: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    userListList: Array<User.AsObject>,
  }
}

export class CommonMessage extends jspb.Message {
  getId(): string;
  setId(value: string): CommonMessage;

  getUser(): User | undefined;
  setUser(value?: User): CommonMessage;
  hasUser(): boolean;
  clearUser(): CommonMessage;

  getMessage(): string;
  setMessage(value: string): CommonMessage;

  getCreatedAt(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setCreatedAt(value?: google_protobuf_timestamp_pb.Timestamp): CommonMessage;
  hasCreatedAt(): boolean;
  clearCreatedAt(): CommonMessage;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CommonMessage.AsObject;
  static toObject(includeInstance: boolean, msg: CommonMessage): CommonMessage.AsObject;
  static serializeBinaryToWriter(message: CommonMessage, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CommonMessage;
  static deserializeBinaryFromReader(message: CommonMessage, reader: jspb.BinaryReader): CommonMessage;
}

export namespace CommonMessage {
  export type AsObject = {
    id: string,
    user?: User.AsObject,
    message: string,
    createdAt?: google_protobuf_timestamp_pb.Timestamp.AsObject,
  }
}

