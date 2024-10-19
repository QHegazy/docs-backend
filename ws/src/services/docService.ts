import { Types } from "mongoose";
import { PartialDocDto, DocDto } from "../Dto/doc";
import { QuillData } from "../models/delta";
import { SuccessResponse, ErrorResponse } from "../customResponse/customResponse";

async function createDoc(docDto: DocDto): Promise<SuccessResponse<any> | ErrorResponse> {
    try {
        const newDoc = new QuillData(docDto);
        const savedDoc = await newDoc.save();

        return {
            status: 201,
            message: "Document created successfully",
            data: savedDoc,
        } as SuccessResponse<any>;

    } catch (error) {
        console.error("Error creating document:", error);
        
        let errorMessage = "Could not create the document";

        if (error instanceof Error) {
            errorMessage = error.message;
        }

        return {
            status: 500,
            message: errorMessage,
            error: errorMessage
        } as ErrorResponse;
    }
}

async function updateDoc(docDto: PartialDocDto): Promise<SuccessResponse<any> | ErrorResponse> {
    try {
        let updateQuery: any;

        if (docDto.id) {
            updateQuery = { _id: docDto.id };
        } else if (docDto.title) {
            updateQuery = { title: docDto.title };
        } else {
            return {
                status: 400,
                message: "Invalid document ID or title",
                error: "Document ID or title is required."
            } as ErrorResponse;
        }

        const updatedDoc = await QuillData.findOneAndUpdate(updateQuery, docDto.content, { new: true, runValidators: true });

        if (!updatedDoc) {
            return {
                status: 404,
                message: "Document not found",
                error: "No document matches the provided criteria."
            } as ErrorResponse;
        }

        return {
            status: 200,
            message: "Document updated successfully",
            data: updatedDoc,
        } as SuccessResponse<any>;

    } catch (error) {
        console.error("Error updating document:", error);
        
        let errorMessage = "Could not update the document";

        if (error instanceof Error) {
            errorMessage = error.message;
        }

        return {
            status: 500,
            message: errorMessage,
            error: errorMessage
        } as ErrorResponse;
    }
}

async function getDoc(docDto: PartialDocDto): Promise<SuccessResponse<any> | ErrorResponse> {
    try {
        let query: any;

        if (docDto.id) {
            query = { _id: docDto.id };
        } else if (docDto.title) {
            query = { title: docDto.title };
        } else {
            return {
                status: 400,
                message: "Invalid document ID or title",
                error: "Document ID or title is required."
            } as ErrorResponse;
        }

        const foundDoc = await QuillData.findOne(query);

        if (!foundDoc) {
            return {
                status: 404,
                message: "Document not found",
                error: "No document matches the provided criteria."
            } as ErrorResponse;
        }

        return {
            status: 200,
            message: "Document retrieved successfully",
            data: foundDoc,
        } as SuccessResponse<any>;

    } catch (error) {
        console.error("Error retrieving document:", error);
        
        let errorMessage = "Could not retrieve the document";

        if (error instanceof Error) {
            errorMessage = error.message;
        }

        return {
            status: 500,
            message: errorMessage,
            error: errorMessage
        } as ErrorResponse;
    }
}


async function deleteDoc(docDto: PartialDocDto): Promise<SuccessResponse<any> | ErrorResponse> {
    try {
        let deleteQuery: any;

        if (docDto.id) {
            deleteQuery = { _id: docDto.id };
        } else if (docDto.title) {
            deleteQuery = { title: docDto.title };
        } else {
            return {
                status: 400,
                message: "Invalid document ID or title",
                error: "Document ID or title is required."
            } as ErrorResponse;
        }

        const deletedDoc = await QuillData.findOneAndDelete(deleteQuery);

        if (!deletedDoc) {
            return {
                status: 404,
                message: "Document not found",
                error: "No document matches the provided criteria."
            } as ErrorResponse;
        }

        return {
            status: 200,
            message: "Document deleted successfully",
            data: deletedDoc,
        } as SuccessResponse<any>;

    } catch (error) {
        console.error("Error deleting document:", error);
        
        let errorMessage = "Could not delete the document";

        if (error instanceof Error) {
            errorMessage = error.message;
        }

        return {
            status: 500,
            message: errorMessage,
            error: errorMessage
        } as ErrorResponse;
    }
}

export { createDoc, updateDoc, getDoc, deleteDoc };
