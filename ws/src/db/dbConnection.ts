import mongoose from "mongoose";
import dotenv from "dotenv";

// Load environment variables from .env file
dotenv.config();
// Database connection
async function dbConnect() {
  const { MONGO_USER, MONGO_PASSWORD, MONGO_PORT, MONGO_DB_NAME } = process.env;

  const MONGO_URI = `mongodb://${MONGO_USER}:${MONGO_PASSWORD}@localhost:${MONGO_PORT}/${MONGO_DB_NAME}`;

  try {
    await mongoose.connect(MONGO_URI);
    console.log("Connected to MongoDB");
  } catch (error) {
    console.error("Database connection error:", error);
  }
}

export default dbConnect;
