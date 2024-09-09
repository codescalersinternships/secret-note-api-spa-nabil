import axios from "axios";
import { Note } from "./types/Note";
import { User } from "./types/User";
const api = axios.create({
  baseURL: "http://localhost:8090",
  headers: {
    "Content-Type": "application/json",
  },
});

interface LoginData {
  email: string;
  password: string;
}

interface SignUpData {
  name: string;
  email: string;
  password: string;
}

export const login = async (data: LoginData): Promise<void> => {
  try {
    const response = await api.post("/signin", data);
    const token = response.data.acces_token;
    sessionStorage.setItem("token", token);
    const userId = response.data.id;
    sessionStorage.setItem("id", userId);
  } catch (error) {
    console.error("Failed to login", error);
    throw error;
  }
};

export const signUp = async (data: SignUpData): Promise<User | null> => {
  try {
    const response = await api.post("/signup", data);
    return response.data;
  } catch (error) {
    console.error("Failed to fetch user notes", error);
    return null;
  }
};

const getToken = () => sessionStorage.getItem("token");
const getUserID = () => sessionStorage.getItem("id");

interface CreateNoteData {
  text: string;
  noteremvisits: number;
  expiredat: string;
  userid: string | null;
}

export const createNote = async (
  data: CreateNoteData
): Promise<Note | null> => {
  try {
    data.userid = getUserID();
    console.log(data);
    const response = await api.post("/create", data);
    return response.data;
  } catch (error) {
    console.error("Failed to fetch user notes", error);
    console.log(data);
    return null;
  }
};

// Add a request interceptor to include the token in every request
api.interceptors.request.use((config) => {
  const token = getToken();
  if (token) {
    config.headers["Authorization"] = `Bearer ${token}`;
  }
  return config;
});

export const getUserNotes = async (): Promise<Note[]> => {
  if (getUserID === null) {
    return [];
  }
  try {
    const response = await api.get<Note[]>(`/${getUserID}`);
    return response.data;
  } catch (error) {
    console.error("Failed to fetch user notes", error);
    return [];
  }
};

export const getNote = async (noteId: string): Promise<Note | null> => {
  try {
    const response = await api.get<Note>(`/note/${noteId}`);
    return response.data;
  } catch (error) {
    console.error("Failed to fetch note", error);
    return null;
  }
};
