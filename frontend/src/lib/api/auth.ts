import { api } from "./client";

export interface RegisterRequest {
  username: string;
  email: string;
  password: string;
}

export interface User {
  id: number;
  username: string;
  email: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export function register(data: RegisterRequest) {
  return api<User>("/api/auth/register", {
    method: "POST",
    body: JSON.stringify(data),
  });
}

export function login(data: LoginRequest) {
  return api<User>("/api/auth/login", {
    method: "POST",
    body: JSON.stringify(data),
  });
}
