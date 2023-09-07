// @generated by protobuf-ts 2.9.1 with parameter generate_dependencies
// @generated from protobuf file "vince/api/v1/api.proto" (package "v1", syntax proto3)
// tslint:disable
import { Build } from "../../config/v1/config";
import { Empty } from "../../../google/protobuf/empty";
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { Cluster_Config } from "../../config/v1/config";
import { Client_Auth } from "../../config/v1/config";
import { Duration } from "../../../google/protobuf/duration";
/**
 * @generated from protobuf message v1.Token
 */
export interface Token {
    /**
     * @generated from protobuf field: bytes pub_key = 1;
     */
    pubKey: Uint8Array;
}
/**
 * @generated from protobuf enum v1.Token.Issuer
 */
export enum Token_Issuer {
    /**
     * @generated from protobuf enum value: SERVER = 0;
     */
    SERVER = 0,
    /**
     * @generated from protobuf enum value: CLIENT = 1;
     */
    CLIENT = 1
}
/**
 * Defines the resources that this token has access to
 *
 * @generated from protobuf enum v1.Token.Scope
 */
export enum Token_Scope {
    /**
     * Token has access to all vince resources
     *
     * @generated from protobuf enum value: ALL = 0;
     */
    ALL = 0,
    /**
     * Token  has access to only cluster resources.
     *
     * @generated from protobuf enum value: CLUSTER = 2;
     */
    CLUSTER = 2
}
/**
 * @generated from protobuf message v1.LoginRequest
 */
export interface LoginRequest {
    /**
     * @generated from protobuf field: string token = 3;
     */
    token: string;
    /**
     * @generated from protobuf field: bytes public_key = 4;
     */
    publicKey: Uint8Array;
    /**
     * When true, the token will be generated by the server.
     *
     * @generated from protobuf field: bool generate = 5;
     */
    generate: boolean;
    /**
     * @generated from protobuf field: google.protobuf.Duration ttl = 6;
     */
    ttl?: Duration;
}
/**
 * @generated from protobuf message v1.LoginResponse
 */
export interface LoginResponse {
    /**
     * @generated from protobuf field: v1.Client.Auth auth = 1;
     */
    auth?: Client_Auth;
}
/**
 * @generated from protobuf message v1.Account
 */
export interface Account {
    /**
     * @generated from protobuf field: string name = 1;
     */
    name: string;
    /**
     * @generated from protobuf field: bytes hashed_password = 2;
     */
    hashedPassword: Uint8Array;
}
/**
 * @generated from protobuf message v1.Error
 */
export interface Error {
    /**
     * @generated from protobuf field: string error = 1;
     */
    error: string;
    /**
     * @generated from protobuf field: int32 code = 2;
     */
    code: number;
}
/**
 * @generated from protobuf message v1.ApplyClusterRequest
 */
export interface ApplyClusterRequest {
    /**
     * @generated from protobuf field: v1.Cluster.Config config = 1;
     */
    config?: Cluster_Config;
}
/**
 * @generated from protobuf message v1.ApplyClusterResponse
 */
export interface ApplyClusterResponse {
    /**
     * @generated from protobuf field: string ok = 1;
     */
    ok: string;
}
/**
 * @generated from protobuf message v1.GetClusterRequest
 */
export interface GetClusterRequest {
}
/**
 * @generated from protobuf message v1.GetClusterResponse
 */
export interface GetClusterResponse {
    /**
     * @generated from protobuf field: v1.Cluster.Config config = 1;
     */
    config?: Cluster_Config;
}
/**
 * @generated from protobuf message v1.Status
 */
export interface Status {
}
/**
 * @generated from protobuf message v1.Notice
 */
export interface Notice {
}
/**
 * @generated from protobuf message v1.Event
 */
export interface Event {
    /**
     * / EventName
     *
     * @generated from protobuf field: string n = 1;
     */
    n: string;
    /**
     * @generated from protobuf field: string url = 2;
     */
    url: string;
    /**
     * Domain
     *
     * @generated from protobuf field: string d = 3;
     */
    d: string;
    /**
     * Screen width
     *
     * @generated from protobuf field: int32 w = 4;
     */
    w: number;
    /**
     * Hash mode
     *
     * @generated from protobuf field: bool h = 5;
     */
    h: boolean;
    /**
     * @generated from protobuf field: string ip = 6;
     */
    ip: string;
    /**
     * user agent
     *
     * @generated from protobuf field: string ua = 7;
     */
    ua: string;
    /**
     * Referrer
     *
     * @generated from protobuf field: string r = 8;
     */
    r: string;
}
// @generated message type with reflection information, may provide speed optimized methods
class Token$Type extends MessageType<Token> {
    constructor() {
        super("v1.Token", [
            { no: 1, name: "pub_key", kind: "scalar", T: 12 /*ScalarType.BYTES*/ }
        ]);
    }
    create(value?: PartialMessage<Token>): Token {
        const message = { pubKey: new Uint8Array(0) };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Token>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Token): Token {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bytes pub_key */ 1:
                    message.pubKey = reader.bytes();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Token, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bytes pub_key = 1; */
        if (message.pubKey.length)
            writer.tag(1, WireType.LengthDelimited).bytes(message.pubKey);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.Token
 */
export const Token = new Token$Type();
// @generated message type with reflection information, may provide speed optimized methods
class LoginRequest$Type extends MessageType<LoginRequest> {
    constructor() {
        super("v1.LoginRequest", [
            { no: 3, name: "token", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "public_key", kind: "scalar", T: 12 /*ScalarType.BYTES*/, options: { "buf.validate.field": { bytes: { len: "32" } } } },
            { no: 5, name: "generate", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 6, name: "ttl", kind: "message", T: () => Duration }
        ], { "buf.validate.message": { cel: [{ id: "client.token", message: "token  is required", expression: "this.generate ? true : size(this.token)>0" }, { id: "client.pub_key", message: "public_key  is required", expression: "this.generate ? true : size(this.public_key)>0" }] } });
    }
    create(value?: PartialMessage<LoginRequest>): LoginRequest {
        const message = { token: "", publicKey: new Uint8Array(0), generate: false };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<LoginRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: LoginRequest): LoginRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string token */ 3:
                    message.token = reader.string();
                    break;
                case /* bytes public_key */ 4:
                    message.publicKey = reader.bytes();
                    break;
                case /* bool generate */ 5:
                    message.generate = reader.bool();
                    break;
                case /* google.protobuf.Duration ttl */ 6:
                    message.ttl = Duration.internalBinaryRead(reader, reader.uint32(), options, message.ttl);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: LoginRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string token = 3; */
        if (message.token !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.token);
        /* bytes public_key = 4; */
        if (message.publicKey.length)
            writer.tag(4, WireType.LengthDelimited).bytes(message.publicKey);
        /* bool generate = 5; */
        if (message.generate !== false)
            writer.tag(5, WireType.Varint).bool(message.generate);
        /* google.protobuf.Duration ttl = 6; */
        if (message.ttl)
            Duration.internalBinaryWrite(message.ttl, writer.tag(6, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.LoginRequest
 */
export const LoginRequest = new LoginRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class LoginResponse$Type extends MessageType<LoginResponse> {
    constructor() {
        super("v1.LoginResponse", [
            { no: 1, name: "auth", kind: "message", T: () => Client_Auth }
        ]);
    }
    create(value?: PartialMessage<LoginResponse>): LoginResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<LoginResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: LoginResponse): LoginResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* v1.Client.Auth auth */ 1:
                    message.auth = Client_Auth.internalBinaryRead(reader, reader.uint32(), options, message.auth);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: LoginResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* v1.Client.Auth auth = 1; */
        if (message.auth)
            Client_Auth.internalBinaryWrite(message.auth, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.LoginResponse
 */
export const LoginResponse = new LoginResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Account$Type extends MessageType<Account> {
    constructor() {
        super("v1.Account", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "hashed_password", kind: "scalar", T: 12 /*ScalarType.BYTES*/ }
        ]);
    }
    create(value?: PartialMessage<Account>): Account {
        const message = { name: "", hashedPassword: new Uint8Array(0) };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Account>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Account): Account {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string name */ 1:
                    message.name = reader.string();
                    break;
                case /* bytes hashed_password */ 2:
                    message.hashedPassword = reader.bytes();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Account, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string name = 1; */
        if (message.name !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.name);
        /* bytes hashed_password = 2; */
        if (message.hashedPassword.length)
            writer.tag(2, WireType.LengthDelimited).bytes(message.hashedPassword);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.Account
 */
export const Account = new Account$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Error$Type extends MessageType<Error> {
    constructor() {
        super("v1.Error", [
            { no: 1, name: "error", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "code", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value?: PartialMessage<Error>): Error {
        const message = { error: "", code: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Error>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Error): Error {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string error */ 1:
                    message.error = reader.string();
                    break;
                case /* int32 code */ 2:
                    message.code = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Error, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string error = 1; */
        if (message.error !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.error);
        /* int32 code = 2; */
        if (message.code !== 0)
            writer.tag(2, WireType.Varint).int32(message.code);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.Error
 */
export const Error = new Error$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ApplyClusterRequest$Type extends MessageType<ApplyClusterRequest> {
    constructor() {
        super("v1.ApplyClusterRequest", [
            { no: 1, name: "config", kind: "message", T: () => Cluster_Config }
        ]);
    }
    create(value?: PartialMessage<ApplyClusterRequest>): ApplyClusterRequest {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ApplyClusterRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ApplyClusterRequest): ApplyClusterRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* v1.Cluster.Config config */ 1:
                    message.config = Cluster_Config.internalBinaryRead(reader, reader.uint32(), options, message.config);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: ApplyClusterRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* v1.Cluster.Config config = 1; */
        if (message.config)
            Cluster_Config.internalBinaryWrite(message.config, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.ApplyClusterRequest
 */
export const ApplyClusterRequest = new ApplyClusterRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ApplyClusterResponse$Type extends MessageType<ApplyClusterResponse> {
    constructor() {
        super("v1.ApplyClusterResponse", [
            { no: 1, name: "ok", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<ApplyClusterResponse>): ApplyClusterResponse {
        const message = { ok: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ApplyClusterResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ApplyClusterResponse): ApplyClusterResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string ok */ 1:
                    message.ok = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: ApplyClusterResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string ok = 1; */
        if (message.ok !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.ok);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.ApplyClusterResponse
 */
export const ApplyClusterResponse = new ApplyClusterResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetClusterRequest$Type extends MessageType<GetClusterRequest> {
    constructor() {
        super("v1.GetClusterRequest", []);
    }
    create(value?: PartialMessage<GetClusterRequest>): GetClusterRequest {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<GetClusterRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetClusterRequest): GetClusterRequest {
        return target ?? this.create();
    }
    internalBinaryWrite(message: GetClusterRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.GetClusterRequest
 */
export const GetClusterRequest = new GetClusterRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetClusterResponse$Type extends MessageType<GetClusterResponse> {
    constructor() {
        super("v1.GetClusterResponse", [
            { no: 1, name: "config", kind: "message", T: () => Cluster_Config }
        ]);
    }
    create(value?: PartialMessage<GetClusterResponse>): GetClusterResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<GetClusterResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetClusterResponse): GetClusterResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* v1.Cluster.Config config */ 1:
                    message.config = Cluster_Config.internalBinaryRead(reader, reader.uint32(), options, message.config);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: GetClusterResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* v1.Cluster.Config config = 1; */
        if (message.config)
            Cluster_Config.internalBinaryWrite(message.config, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.GetClusterResponse
 */
export const GetClusterResponse = new GetClusterResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Status$Type extends MessageType<Status> {
    constructor() {
        super("v1.Status", []);
    }
    create(value?: PartialMessage<Status>): Status {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Status>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Status): Status {
        return target ?? this.create();
    }
    internalBinaryWrite(message: Status, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.Status
 */
export const Status = new Status$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Notice$Type extends MessageType<Notice> {
    constructor() {
        super("v1.Notice", []);
    }
    create(value?: PartialMessage<Notice>): Notice {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Notice>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Notice): Notice {
        return target ?? this.create();
    }
    internalBinaryWrite(message: Notice, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.Notice
 */
export const Notice = new Notice$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Event$Type extends MessageType<Event> {
    constructor() {
        super("v1.Event", [
            { no: 1, name: "n", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "buf.validate.field": { required: true } } },
            { no: 2, name: "url", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "buf.validate.field": { required: true, string: { uri: true } } } },
            { no: 3, name: "d", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "buf.validate.field": { required: true, string: { hostname: true } } } },
            { no: 4, name: "w", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "h", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 6, name: "ip", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "buf.validate.field": { string: { ip: true } } } },
            { no: 7, name: "ua", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 8, name: "r", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<Event>): Event {
        const message = { n: "", url: "", d: "", w: 0, h: false, ip: "", ua: "", r: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Event>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Event): Event {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string n */ 1:
                    message.n = reader.string();
                    break;
                case /* string url */ 2:
                    message.url = reader.string();
                    break;
                case /* string d */ 3:
                    message.d = reader.string();
                    break;
                case /* int32 w */ 4:
                    message.w = reader.int32();
                    break;
                case /* bool h */ 5:
                    message.h = reader.bool();
                    break;
                case /* string ip */ 6:
                    message.ip = reader.string();
                    break;
                case /* string ua */ 7:
                    message.ua = reader.string();
                    break;
                case /* string r */ 8:
                    message.r = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Event, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string n = 1; */
        if (message.n !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.n);
        /* string url = 2; */
        if (message.url !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.url);
        /* string d = 3; */
        if (message.d !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.d);
        /* int32 w = 4; */
        if (message.w !== 0)
            writer.tag(4, WireType.Varint).int32(message.w);
        /* bool h = 5; */
        if (message.h !== false)
            writer.tag(5, WireType.Varint).bool(message.h);
        /* string ip = 6; */
        if (message.ip !== "")
            writer.tag(6, WireType.LengthDelimited).string(message.ip);
        /* string ua = 7; */
        if (message.ua !== "")
            writer.tag(7, WireType.LengthDelimited).string(message.ua);
        /* string r = 8; */
        if (message.r !== "")
            writer.tag(8, WireType.LengthDelimited).string(message.r);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.Event
 */
export const Event = new Event$Type();
/**
 * @generated ServiceType for protobuf service v1.Vince
 */
export const Vince = new ServiceType("v1.Vince", [
    { name: "Login", options: { "google.api.http": { post: "/v1/login" } }, I: LoginRequest, O: LoginResponse },
    { name: "ApplyCluster", options: { "google.api.http": { post: "/v1/cluster/apply" } }, I: ApplyClusterRequest, O: ApplyClusterResponse },
    { name: "GetCluster", options: { "google.api.http": { get: "/v1/cluster" } }, I: GetClusterRequest, O: GetClusterResponse },
    { name: "Version", options: { "google.api.http": { get: "/v1/version" } }, I: Empty, O: Build },
    { name: "SendEvent", options: { "google.api.http": { post: "/api/event" } }, I: Event, O: Empty }
]);
