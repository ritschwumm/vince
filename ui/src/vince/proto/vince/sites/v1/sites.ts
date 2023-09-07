// @generated by protobuf-ts 2.9.1 with parameter generate_dependencies
// @generated from protobuf file "vince/sites/v1/sites.proto" (package "v1", syntax proto3)
// tslint:disable
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
import { Goal } from "../../goals/v1/goals";
/**
 * @generated from protobuf message v1.Site
 */
export interface Site {
    /**
     * @generated from protobuf field: string domain = 1;
     */
    domain: string;
    /**
     * @generated from protobuf field: repeated v1.Goal goals = 2;
     */
    goals: Goal[];
}
/**
 * @generated from protobuf message v1.CreateSiteRequest
 */
export interface CreateSiteRequest {
    /**
     * @generated from protobuf field: string domain = 1;
     */
    domain: string;
}
/**
 * @generated from protobuf message v1.CreateSiteResponse
 */
export interface CreateSiteResponse {
    /**
     * @generated from protobuf field: v1.Site site = 1;
     */
    site?: Site;
}
/**
 * @generated from protobuf message v1.GetSiteRequest
 */
export interface GetSiteRequest {
    /**
     * @generated from protobuf field: string domain = 1;
     */
    domain: string;
}
/**
 * @generated from protobuf message v1.GetSiteResponse
 */
export interface GetSiteResponse {
    /**
     * @generated from protobuf field: v1.Site site = 1;
     */
    site?: Site;
}
/**
 * @generated from protobuf message v1.ListSitesRequest
 */
export interface ListSitesRequest {
}
/**
 * @generated from protobuf message v1.ListSitesResponse
 */
export interface ListSitesResponse {
    /**
     * @generated from protobuf field: repeated v1.Site list = 1;
     */
    list: Site[];
}
/**
 * @generated from protobuf message v1.DeleteSiteRequest
 */
export interface DeleteSiteRequest {
    /**
     * @generated from protobuf field: string domain = 1;
     */
    domain: string;
}
/**
 * @generated from protobuf message v1.DeleteSiteResponse
 */
export interface DeleteSiteResponse {
}
// @generated message type with reflection information, may provide speed optimized methods
class Site$Type extends MessageType<Site> {
    constructor() {
        super("v1.Site", [
            { no: 1, name: "domain", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "goals", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Goal }
        ]);
    }
    create(value?: PartialMessage<Site>): Site {
        const message = { domain: "", goals: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Site>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Site): Site {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string domain */ 1:
                    message.domain = reader.string();
                    break;
                case /* repeated v1.Goal goals */ 2:
                    message.goals.push(Goal.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: Site, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string domain = 1; */
        if (message.domain !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.domain);
        /* repeated v1.Goal goals = 2; */
        for (let i = 0; i < message.goals.length; i++)
            Goal.internalBinaryWrite(message.goals[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.Site
 */
export const Site = new Site$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateSiteRequest$Type extends MessageType<CreateSiteRequest> {
    constructor() {
        super("v1.CreateSiteRequest", [
            { no: 1, name: "domain", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "buf.validate.field": { required: true, string: { hostname: true } } } }
        ]);
    }
    create(value?: PartialMessage<CreateSiteRequest>): CreateSiteRequest {
        const message = { domain: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<CreateSiteRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CreateSiteRequest): CreateSiteRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string domain */ 1:
                    message.domain = reader.string();
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
    internalBinaryWrite(message: CreateSiteRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string domain = 1; */
        if (message.domain !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.domain);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.CreateSiteRequest
 */
export const CreateSiteRequest = new CreateSiteRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CreateSiteResponse$Type extends MessageType<CreateSiteResponse> {
    constructor() {
        super("v1.CreateSiteResponse", [
            { no: 1, name: "site", kind: "message", T: () => Site }
        ]);
    }
    create(value?: PartialMessage<CreateSiteResponse>): CreateSiteResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<CreateSiteResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: CreateSiteResponse): CreateSiteResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* v1.Site site */ 1:
                    message.site = Site.internalBinaryRead(reader, reader.uint32(), options, message.site);
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
    internalBinaryWrite(message: CreateSiteResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* v1.Site site = 1; */
        if (message.site)
            Site.internalBinaryWrite(message.site, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.CreateSiteResponse
 */
export const CreateSiteResponse = new CreateSiteResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetSiteRequest$Type extends MessageType<GetSiteRequest> {
    constructor() {
        super("v1.GetSiteRequest", [
            { no: 1, name: "domain", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "buf.validate.field": { required: true, string: { hostname: true } } } }
        ]);
    }
    create(value?: PartialMessage<GetSiteRequest>): GetSiteRequest {
        const message = { domain: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<GetSiteRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetSiteRequest): GetSiteRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string domain */ 1:
                    message.domain = reader.string();
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
    internalBinaryWrite(message: GetSiteRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string domain = 1; */
        if (message.domain !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.domain);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.GetSiteRequest
 */
export const GetSiteRequest = new GetSiteRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetSiteResponse$Type extends MessageType<GetSiteResponse> {
    constructor() {
        super("v1.GetSiteResponse", [
            { no: 1, name: "site", kind: "message", T: () => Site }
        ]);
    }
    create(value?: PartialMessage<GetSiteResponse>): GetSiteResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<GetSiteResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetSiteResponse): GetSiteResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* v1.Site site */ 1:
                    message.site = Site.internalBinaryRead(reader, reader.uint32(), options, message.site);
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
    internalBinaryWrite(message: GetSiteResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* v1.Site site = 1; */
        if (message.site)
            Site.internalBinaryWrite(message.site, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.GetSiteResponse
 */
export const GetSiteResponse = new GetSiteResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListSitesRequest$Type extends MessageType<ListSitesRequest> {
    constructor() {
        super("v1.ListSitesRequest", []);
    }
    create(value?: PartialMessage<ListSitesRequest>): ListSitesRequest {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ListSitesRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListSitesRequest): ListSitesRequest {
        return target ?? this.create();
    }
    internalBinaryWrite(message: ListSitesRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.ListSitesRequest
 */
export const ListSitesRequest = new ListSitesRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListSitesResponse$Type extends MessageType<ListSitesResponse> {
    constructor() {
        super("v1.ListSitesResponse", [
            { no: 1, name: "list", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Site }
        ]);
    }
    create(value?: PartialMessage<ListSitesResponse>): ListSitesResponse {
        const message = { list: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ListSitesResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListSitesResponse): ListSitesResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated v1.Site list */ 1:
                    message.list.push(Site.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: ListSitesResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated v1.Site list = 1; */
        for (let i = 0; i < message.list.length; i++)
            Site.internalBinaryWrite(message.list[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.ListSitesResponse
 */
export const ListSitesResponse = new ListSitesResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteSiteRequest$Type extends MessageType<DeleteSiteRequest> {
    constructor() {
        super("v1.DeleteSiteRequest", [
            { no: 1, name: "domain", kind: "scalar", T: 9 /*ScalarType.STRING*/, options: { "buf.validate.field": { required: true, string: { hostname: true } } } }
        ]);
    }
    create(value?: PartialMessage<DeleteSiteRequest>): DeleteSiteRequest {
        const message = { domain: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<DeleteSiteRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DeleteSiteRequest): DeleteSiteRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string domain */ 1:
                    message.domain = reader.string();
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
    internalBinaryWrite(message: DeleteSiteRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string domain = 1; */
        if (message.domain !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.domain);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.DeleteSiteRequest
 */
export const DeleteSiteRequest = new DeleteSiteRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeleteSiteResponse$Type extends MessageType<DeleteSiteResponse> {
    constructor() {
        super("v1.DeleteSiteResponse", []);
    }
    create(value?: PartialMessage<DeleteSiteResponse>): DeleteSiteResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<DeleteSiteResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: DeleteSiteResponse): DeleteSiteResponse {
        return target ?? this.create();
    }
    internalBinaryWrite(message: DeleteSiteResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.DeleteSiteResponse
 */
export const DeleteSiteResponse = new DeleteSiteResponse$Type();
/**
 * @generated ServiceType for protobuf service v1.Sites
 */
export const Sites = new ServiceType("v1.Sites", [
    { name: "CreateSite", options: { "google.api.http": { post: "/v1/sites" } }, I: CreateSiteRequest, O: CreateSiteResponse },
    { name: "GetSite", options: { "google.api.http": { get: "/v1/site" } }, I: GetSiteRequest, O: GetSiteResponse },
    { name: "ListSites", options: { "google.api.http": { get: "/v1/sites" } }, I: ListSitesRequest, O: ListSitesResponse },
    { name: "DeleteSite", options: { "google.api.http": { delete: "/v1/site" } }, I: DeleteSiteRequest, O: DeleteSiteResponse }
]);
