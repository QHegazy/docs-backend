import express from "express";
import { createServer } from "node:http";
import { Server } from "socket.io";
import cors from "cors"; 
import { config } from "dotenv"; 
import dbConnect from "./model/dbConnection";
config();
var dd: any;
const app = express();
app.use(cors());

// Initialize HTTP server
const server = createServer(app);
dbConnect();
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

  socket.on("joinRoom", (room: string) => {
    if (!rooms[room]) {
      rooms[room] = [];
    }
    socket.join(room);
    console.log(`User joined room: ${room}`);
    socket.emit("get-doc", dd); // Send existing document data
  });

  socket.on("leaveRoom", (room: string) => {
    socket.leave(room);
    console.log(`User left room: ${room}`);
  });

  socket.on("sendRoom", async (room: string, message: any) => {
    try {
      rooms[room] = message; // Update room content
      console.log(dd); // Emit the message to everyone in the room except the sender
      socket.to(room).emit("RoomMessage", message);
    } catch (error) {
      console.error(`Error sending message to room ${room}:`, error);
    }
  });

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
  socket.on("save-doc", (doc) => {
    dd = doc;
  });
  app.get("/doc", (req, res) => {
    res.json(dd);
  });
  socket.on("disconnect", () => {
    console.log(`Client disconnected with id: ${socket.id}`);
  });
});

// Start the server
const PORT = process.env.PORT;
server.listen(PORT, () => {
  console.log(`Server running at http://localhost:${PORT}`);
});
