interface BaseResponse {
    status: number;   
    message: string;  
}

interface SuccessResponse<T> extends BaseResponse {
    data: T;  
}

interface ErrorResponse extends BaseResponse {
    error: any; 
}

export type { SuccessResponse, ErrorResponse };