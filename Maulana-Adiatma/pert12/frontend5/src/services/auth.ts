import type { LoginResponse, LogoutResponse, SignupResponse } from "../types/auth";
import api from "./api";

export async function login(email: string, password: string): Promise<LoginResponse> {
  const formData = new FormData();
  formData.append("email", email);
  formData.append("password", password);

  const res = await api.post<LoginResponse>("/auth/login", formData, {
    headers: { "Content-Type": "multipart/form-data" },
  });

  return res.data;
}

export async function signup(name: string, email: string, password: string, role: "user" | "admin"): Promise<SignupResponse> {
  const formData = new FormData();
  formData.append("name", name);
  formData.append("email", email);
  formData.append("password", password);
  formData.append("role", role);

  const res = await api.post<SignupResponse>("/auth/register", formData, {
    headers: { "Content-Type": "multipart/form-data" },
  });

  return res.data;
}

export async function logout(): Promise<{ message: string }> {
  const res = await api.post<LogoutResponse>("/auth/logout");
  return res.data
}