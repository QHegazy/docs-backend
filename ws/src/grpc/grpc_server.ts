import * as grpc from '@grpc/grpc-js';
import { insertDocument, documentProto } from '../services/grpc';

function grpcServer() {
    const server = new grpc.Server();
    server.addService(documentProto.document.NewDocument.service, { InsertDocument: insertDocument });

    const PORT = 50051;
    server.bindAsync(`localhost:${PORT}`, grpc.ServerCredentials.createInsecure(), (error, port) => {
        if (error) {
            console.error(`Error binding server: ${error}`);
            return;
        }
        console.log(`Server running at http://localhost:${port}`);
    });
}

export { grpcServer };
