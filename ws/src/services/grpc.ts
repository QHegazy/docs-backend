import * as grpc from '@grpc/grpc-js';
import { loadSync } from '@grpc/proto-loader';
import * as path from 'path';
import { DocumentRequest, DocumentResponse } from '../document/document';
import { authenticate } from '../middlewares/grpc_auth';


const PROTO_PATH = path.join(__dirname, '../../../document.proto');
const packageDefinition = loadSync(PROTO_PATH, { keepCase: true, longs: String, defaults: true, oneofs: true });
const documentProto = grpc.loadPackageDefinition(packageDefinition) as any;

const insertDocument = (call: grpc.ServerUnaryCall<DocumentRequest, DocumentResponse>, callback: grpc.sendUnaryData<DocumentResponse>) => {
  authenticate(call, callback, () => {
      const title = call.request.title;
      if (!title) {
          return callback({ code: grpc.status.INVALID_ARGUMENT, message: 'Title is required' }, null);
      }
      const documentId = `doc_${Math.random().toString(36).substring(2, 15)}`;
      callback(null, { document_id: documentId });
  });
};




export { insertDocument,documentProto };