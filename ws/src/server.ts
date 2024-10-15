import express from "express";
import { createServer } from "node:http";
import { Server } from "socket.io";
import cors from "cors";
import { config } from "dotenv";
import dbConnect from "./db/dbConnection";
import { QuillData } from "./models/delta";
import { getDoc, updateDoc } from "./services/docService";
import { Types } from "mongoose";

config(); // Load environment variables

const app = express();
app.use(cors({
  origin: process.env.ORIGIN,
  credentials: true,
}));

// Initialize HTTP server
const server = createServer(app);
dbConnect(); // Connect to MongoDB

// Initialize Socket.IO server
const io = new Server(server, {
  cors: {
    origin: process.env.ORIGIN,
    methods: ["GET", "POST"],
    credentials: true,
  },
});

// Initialize rooms data structure
const rooms: Record<string, any[]> = {};

// Handle socket connections
io.on("connection", (socket) => {
  console.log(`New connection: ${socket.id}`);

  socket.on("joinRoom", async (room: string) => {
    if (!rooms[room]) {
      rooms[room] = [];
    }
    socket.join(room);
    console.log(`User joined room: ${room}`);
  });

  // Handle user leaving a room
  socket.on("leaveRoom", (room: string) => {
    socket.leave(room);
    console.log(`User left room: ${room}`);
  });

  socket.on("save-doc", async (id,message) => {
   const doc = await getDoc({id:id as Types.ObjectId});
   if (doc){
    console.log(doc)
   }

   
  });

  // Handle broadcasting messages to a room
  socket.on("sendRoom", async (room: string, message: any) => {
    try {
      rooms[room] = message; // Update room content
      console.log(`Room ${room} updated with message`);
      socket.to(room).emit("RoomMessage", message); // Emit the message to everyone in the room
    } catch (error) {
      console.error(`Error sending message to room ${room}:`, error);
    }
  });

  // Handle general messages
  socket.on("message", (data: any, callback?: (response: string) => void) => {
    try {
      console.log("New message received:", data);
      socket.broadcast.emit("message", data); // Broadcast to other clients
      if (callback) {
        callback("Message received and broadcasted");
      }
    } catch (error) {
      console.error("Error handling message:", error);
      if (callback) {
        callback("Error processing message");
      }
    }
  });

  // Handle disconnections
  socket.on("disconnect", () => {
    console.log(`Client disconnected with id: ${socket.id}`);
  });
});

// API endpoint to get the latest document
app.get("/doc/:id", async (req, res) => {
  try {
    const docId = req.params.id;
    const latestDoc = await QuillData.findOne({_id:docId}).sort({ createdAt: -1 });
    res.json(latestDoc || {});
  } catch (error) {
    res.status(500).json({ error: "Error fetching document" });
  }
});

// Start the server
const PORT = process.env.PORT || 3000;
server.listen(PORT, () => {
  console.log(`Server running at http://localhost:${PORT}`);
});
