import { Schema, model } from "mongoose";

// Define the schema for Quill data as a Delta JSON
const quillDataSchema = new Schema(
  {
    title: {
      type: String,
      required: true,
      trim: true,
    },
    content: {
      type: Object, 
      required: true,
    },
    createdAt: {
      type: Date,
      default: Date.now, 
    },
    updatedAt: {
      type: Date,
      default: Date.now,
    },
    deletedAt: {
      type: Date,
      default: Date.now,
    },
    
  },
  { timestamps: true } 
);

// Create the model from the schema
const QuillData = model("QuillData", quillDataSchema);

export { QuillData };
