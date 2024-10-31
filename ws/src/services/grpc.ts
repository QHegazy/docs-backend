import * as grpc from '@grpc/grpc-js';
import { loadSync } from '@grpc/proto-loader';
import * as path from 'path';
import { DocumentRequest, DocumentResponse ,DeleteDocumentRequest,DeleteDocumentResponse} from '../document/document';
import { authenticate } from '../middlewares/grpc_auth';
import { createDoc } from './docService';


const PROTO_PATH = path.join(__dirname, '../../../document.proto');
const packageDefinition = loadSync(PROTO_PATH, { keepCase: true, longs: String, defaults: true, oneofs: true });
const documentProto = grpc.loadPackageDefinition(packageDefinition) as any;

const insertDocument = async (call: grpc.ServerUnaryCall<DocumentRequest, DocumentResponse>, callback: grpc.sendUnaryData<DocumentResponse>) => {
  authenticate(call, callback, async () => {
    const { title } = call.request;
    if (!title) {
      return callback({ code: grpc.status.INVALID_ARGUMENT, message: 'Title is required' }, null);
    }

    try {
      const result: any = await createDoc(title);

      if (result.status !== 201) {
        return callback({ code: grpc.status.INTERNAL, message: "Failed to create document" }, null);
      }

      const documentId = result.data._id;

      callback(null, { document_id: documentId });
    } catch (error) {
      console.error("Error creating document:", error);
      callback({ code: grpc.status.INTERNAL, message: 'Failed to create document' }, null);
    }
  });
};

const deleteDocument = (call: grpc.ServerUnaryCall<DeleteDocumentRequest, DeleteDocumentResponse>, callback: grpc.sendUnaryData<DeleteDocumentResponse>) => {
  authenticate(call, callback, () => {
      const documentId = call.request.document_id;
      if (!documentId) {
          return callback({ code: grpc.status.INVALID_ARGUMENT, message: 'Document ID is required' }, null);
      }
      callback(null, { message: `Document ${documentId} deleted successfully` });
  });
};





export { insertDocument, deleteDocument, documentProto };