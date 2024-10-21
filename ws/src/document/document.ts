// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.2.4
//   protoc               v3.19.6
// source: document.proto

/* eslint-disable */
import { BinaryReader, BinaryWriter } from "@bufbuild/protobuf/wire";

export const protobufPackage = "document";

export interface DocumentRequest {
  title: string;
}

export interface DocumentResponse {
  document_id: string;
}

export interface DeleteDocumentRequest {
  document_id: string;
  title: string;
}

export interface DeleteDocumentResponse {
  message: string;
}

function createBaseDocumentRequest(): DocumentRequest {
  return { title: "" };
}

export const DocumentRequest: MessageFns<DocumentRequest> = {
  encode(message: DocumentRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.title !== "") {
      writer.uint32(10).string(message.title);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): DocumentRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDocumentRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.title = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DocumentRequest {
    return { title: isSet(object.title) ? globalThis.String(object.title) : "" };
  },

  toJSON(message: DocumentRequest): unknown {
    const obj: any = {};
    if (message.title !== "") {
      obj.title = message.title;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DocumentRequest>, I>>(base?: I): DocumentRequest {
    return DocumentRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DocumentRequest>, I>>(object: I): DocumentRequest {
    const message = createBaseDocumentRequest();
    message.title = object.title ?? "";
    return message;
  },
};

function createBaseDocumentResponse(): DocumentResponse {
  return { document_id: "" };
}

export const DocumentResponse: MessageFns<DocumentResponse> = {
  encode(message: DocumentResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.document_id !== "") {
      writer.uint32(10).string(message.document_id);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): DocumentResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDocumentResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.document_id = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DocumentResponse {
    return { document_id: isSet(object.document_id) ? globalThis.String(object.document_id) : "" };
  },

  toJSON(message: DocumentResponse): unknown {
    const obj: any = {};
    if (message.document_id !== "") {
      obj.document_id = message.document_id;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DocumentResponse>, I>>(base?: I): DocumentResponse {
    return DocumentResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DocumentResponse>, I>>(object: I): DocumentResponse {
    const message = createBaseDocumentResponse();
    message.document_id = object.document_id ?? "";
    return message;
  },
};

function createBaseDeleteDocumentRequest(): DeleteDocumentRequest {
  return { document_id: "", title: "" };
}

export const DeleteDocumentRequest: MessageFns<DeleteDocumentRequest> = {
  encode(message: DeleteDocumentRequest, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.document_id !== "") {
      writer.uint32(10).string(message.document_id);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): DeleteDocumentRequest {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteDocumentRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.document_id = reader.string();
          continue;
        }
        case 2: {
          if (tag !== 18) {
            break;
          }

          message.title = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteDocumentRequest {
    return {
      document_id: isSet(object.document_id) ? globalThis.String(object.document_id) : "",
      title: isSet(object.title) ? globalThis.String(object.title) : "",
    };
  },

  toJSON(message: DeleteDocumentRequest): unknown {
    const obj: any = {};
    if (message.document_id !== "") {
      obj.document_id = message.document_id;
    }
    if (message.title !== "") {
      obj.title = message.title;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteDocumentRequest>, I>>(base?: I): DeleteDocumentRequest {
    return DeleteDocumentRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DeleteDocumentRequest>, I>>(object: I): DeleteDocumentRequest {
    const message = createBaseDeleteDocumentRequest();
    message.document_id = object.document_id ?? "";
    message.title = object.title ?? "";
    return message;
  },
};

function createBaseDeleteDocumentResponse(): DeleteDocumentResponse {
  return { message: "" };
}

export const DeleteDocumentResponse: MessageFns<DeleteDocumentResponse> = {
  encode(message: DeleteDocumentResponse, writer: BinaryWriter = new BinaryWriter()): BinaryWriter {
    if (message.message !== "") {
      writer.uint32(10).string(message.message);
    }
    return writer;
  },

  decode(input: BinaryReader | Uint8Array, length?: number): DeleteDocumentResponse {
    const reader = input instanceof BinaryReader ? input : new BinaryReader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDeleteDocumentResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1: {
          if (tag !== 10) {
            break;
          }

          message.message = reader.string();
          continue;
        }
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skip(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DeleteDocumentResponse {
    return { message: isSet(object.message) ? globalThis.String(object.message) : "" };
  },

  toJSON(message: DeleteDocumentResponse): unknown {
    const obj: any = {};
    if (message.message !== "") {
      obj.message = message.message;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DeleteDocumentResponse>, I>>(base?: I): DeleteDocumentResponse {
    return DeleteDocumentResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DeleteDocumentResponse>, I>>(object: I): DeleteDocumentResponse {
    const message = createBaseDeleteDocumentResponse();
    message.message = object.message ?? "";
    return message;
  },
};

export interface NewDocument {
  InsertDocument(request: DocumentRequest): Promise<DocumentResponse>;
}

export const NewDocumentServiceName = "document.NewDocument";
export class NewDocumentClientImpl implements NewDocument {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || NewDocumentServiceName;
    this.rpc = rpc;
    this.InsertDocument = this.InsertDocument.bind(this);
  }
  InsertDocument(request: DocumentRequest): Promise<DocumentResponse> {
    const data = DocumentRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "InsertDocument", data);
    return promise.then((data) => DocumentResponse.decode(new BinaryReader(data)));
  }
}

export interface RemoveDocument {
  DeleteDocument(request: DeleteDocumentRequest): Promise<DeleteDocumentResponse>;
}

export const RemoveDocumentServiceName = "document.RemoveDocument";
export class RemoveDocumentClientImpl implements RemoveDocument {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || RemoveDocumentServiceName;
    this.rpc = rpc;
    this.DeleteDocument = this.DeleteDocument.bind(this);
  }
  DeleteDocument(request: DeleteDocumentRequest): Promise<DeleteDocumentResponse> {
    const data = DeleteDocumentRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DeleteDocument", data);
    return promise.then((data) => DeleteDocumentResponse.decode(new BinaryReader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

export interface MessageFns<T> {
  encode(message: T, writer?: BinaryWriter): BinaryWriter;
  decode(input: BinaryReader | Uint8Array, length?: number): T;
  fromJSON(object: any): T;
  toJSON(message: T): unknown;
  create<I extends Exact<DeepPartial<T>, I>>(base?: I): T;
  fromPartial<I extends Exact<DeepPartial<T>, I>>(object: I): T;
}
