import { Types } from "mongoose";
import { PartialDocDto, DocDto } from "../Dto/doc";
import { QuillData } from "../models/delta";
import { SuccessResponse, ErrorResponse } from "../customResponse/customResponse";

async function createDoc(title: string): Promise<SuccessResponse<any> | ErrorResponse> {
    try {
        const newDoc = new QuillData({
            title: title,
            content: ""
        });
        
        const savedDoc = await newDoc.save();
        var res: SuccessResponse<any> = {
            status: 201,
            message: "Document created successfully",
            data: savedDoc,
        };
        
        
        return res;


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

        // Set the query to find the document based on either id or title
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

        // Prepare the update object based on provided fields in docDto
        const updateData: Partial<DocDto> = {};
        if (docDto.content) {
            updateData.content = docDto.content;
        }
        if (docDto.title) {
            updateData.title = docDto.title;
        }

        // Perform the update operation with only the provided fields
        const updatedDoc = await QuillData.findOneAndUpdate(updateQuery, updateData, { new: true, runValidators: true });

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
