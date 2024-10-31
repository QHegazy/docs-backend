import * as jwt from 'jsonwebtoken';
import * as grpc from '@grpc/grpc-js';
const SECRET_KEY =process.env.SECRET_KEY as string;

function authenticate(call: grpc.ServerUnaryCall<any, any>, callback: grpc.sendUnaryData<any>, next: Function) {
    console.log(call.metadata.get('authorization'))
    const token = call.metadata.get('authorization')[0];
    if (!token) {
        return callback({
            code: grpc.status.UNAUTHENTICATED,
            message: 'No token provided',
        }, null);
    }
  
    jwt.verify(token as string, SECRET_KEY, (err: any) => {
        if (err) {
            return callback({
                code: grpc.status.UNAUTHENTICATED,
                message: 'Invalid token',
            }, null);
        }
        next();
    });
  }

export {authenticate}