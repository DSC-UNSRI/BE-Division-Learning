import type { User } from "./user";

export type LoginResponse = {
  message: string;
  token: string;
  user: {
    id: number;
    role: "user" | "admin";
    name: string;
    email: string;
    profilePic: string;
  };
};

export type SignupResponse = {
  message: string;
  user: User;
};

export type LogoutResponse = {
  message: string;
};
