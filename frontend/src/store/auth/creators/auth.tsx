import { createAsyncThunk } from "@reduxjs/toolkit";
import { AuthService } from "@core/adapters/AuthService";

export const getActionAuthLogin = createAsyncThunk(
  "auth/login",
  async ({ email, password }: { email: string; password: string }) => {
    const baseUrl = import.meta.env.VITE_API_BASE_URL;

    const authService = new AuthService(baseUrl);
    return authService.login({ email, password });
  }
);
