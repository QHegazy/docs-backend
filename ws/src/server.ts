import express from "express";
import { createServer } from "node:http";
import { Server } from "socket.io";
import cors from "cors";
import { config } from "dotenv";
import dbConnect from "./db/dbConnection";
import { QuillData } from "./models/delta";
import { getDoc, updateDoc } from "./services/docService";
import { Types } from "mongoose";
import { DocDto, PartialDocDto } from "./Dto/doc";
import { grpcServer } from "./grpc/grpc_server";
import redisAdapter from "socket.io-redis";
import Redis from "ioredis";

config();

const app = express();
app.use(cors({
  origin: process.env.ORIGIN,
  credentials: true,
}));

const server = createServer(app);
dbConnect();
grpcServer();

const redisClient = new Redis({
  host: process.env.REDIS_HOST, 
  port: parseInt(process.env.REDIS_PORT || '6379', 10),
});

const io = new Server(server, {
  cors: {
    origin: process.env.ORIGIN,
    methods: ["GET", "POST"],
    credentials: true,
  },
});

io.adapter(redisAdapter({ host: process.env.REDIS_HOST, port: parseInt(process.env.REDIS_PORT || '6379', 10) }));

io.on("connection", (socket) => {
  console.log(`New connection: ${socket.id}`);

  socket.on("joinRoom", async (room: string) => {
    socket.join(room);
    console.log(`User joined room: ${room}`);
  });

  socket.on("leaveRoom", (room: string) => {
    socket.leave(room);
    console.log(`User left room: ${room}`);
  });

  socket.on("save-doc", async (id, message) => {
    try {
      const updateData: PartialDocDto = { id: new Types.ObjectId(id), content: message };

      const doc = await updateDoc(updateData);
      socket.to(id).emit("document-updated", doc);
    } catch (error) {
      console.error("Error saving document:", error);
      socket.emit("document-error", error);
    }
  });

  socket.on("sendRoom", async (room: string, message: any) => {
    try {
      socket.to(room).emit("RoomMessage", message); 
    } catch (error) {
      console.error(`Error sending message to room ${room}:`, error);
    }
  });

  socket.on("disconnect", () => {
    console.log(`Client disconnected with id: ${socket.id}`);
  });
});

app.get("/doc/:id", async (req, res) => {
  try {
    const docId = req.params.id;

    const latestDoc = await QuillData.findOne({ _id: docId }).sort({ createdAt: -1 });

    if (latestDoc) {
      const responseDoc = {
        ...latestDoc.toObject(),
        _id: latestDoc._id.toString(),
      };
      responseDoc.content = latestDoc.content;
      res.json(responseDoc.content);
    } else {
      res.status(404).json({ message: "Document not found" });
    }
  } catch (error) {
    console.error("Error fetching document:", error);
    res.status(500).json({ error: "Error fetching document" });
  }
});

const PORT = process.env.PORT || 3000;
server.listen(PORT, () => {
  console.log(`Server running at http://localhost:${PORT}`);
});

