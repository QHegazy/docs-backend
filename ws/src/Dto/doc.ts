import { Types } from "mongoose";

type DocDto = {
    id: Types.ObjectId;  
    title: string;
    content: object;  
};

type PartialDocDto = Partial<DocDto>;


export type { DocDto , PartialDocDto };