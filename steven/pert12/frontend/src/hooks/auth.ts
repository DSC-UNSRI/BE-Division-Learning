import { useState } from "react";
import { login, logout, signup } from "../services/auth";
import type { LoginResponse, SignupResponse } from "../types/auth";
import { isAxiosError } from "../utils/axios";

export function useLogin() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [message, setMessage] = useState<string | null>(null);
  const [user, setUser] = useState<LoginResponse["user"] | null>(null);

  const doLogin = async (email: string, password: string) => {
    setLoading(true);
    setError(null);

    try {
      const data = await login(email, password);
      setUser(data.user);
      setMessage(data.message)
      return data;
    } catch (err: unknown) {
      if (isAxiosError(err)) {
        setError(err.response?.data?.message || err.message);
      } else {
        setError("Unknown error");
      }
      throw err
    } finally {
      setLoading(false);
    }
  };

  return { loading, message , error, user, doLogin };
}

export function useSignup() {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [message, setMessage] = useState<string | null>(null);
  const [user, setUser] = useState<SignupResponse["user"] | null>(null);

  const doSignup = async (name: string, email: string, password: string, role: "user" | "admin") => {
    setLoading(true);
    setError(null);

    try {
      const data = await signup(name, email, password, role);
      setUser(data.user);
      setMessage(data.message)
      return data;
    } catch (err: unknown) {
      if (isAxiosError(err)) {
        setError(err.response?.data?.message || err.message);
      } else {
        setError("Unknown error");
      }
      throw err
    } finally {
      setLoading(false);
    }
  };

  return { loading, message , error, user, doSignup };
}

export function useLogout() {
  const [logoutLoading, setLoading] = useState(false);
  const [logoutError, setError] = useState<Error | null>(null);
    const [logoutSuccess, setSuccess] = useState<string | null>(null);

  async function handleLogout() {
    setLoading(true);
    setError(null);
    setSuccess(null);

    try {
      const res = await logout();
      setSuccess(res.message); 
    } catch (err) {
      setError(err as Error);
    } finally {
      setLoading(false);
    }
  }

  return { handleLogout, logoutLoading, logoutError, logoutSuccess };
}