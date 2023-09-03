import { createAsyncThunk } from "@reduxjs/toolkit";
import { Container } from "inversify";
import { Types } from "@core/container/types";
import { IAuthService } from "@core/ports/ports";

export const getActionAuthLogin = createAsyncThunk(
  "auth/login",
  async (
    { email, password }: { email: string; password: string },
    { extra }
  ) => {
    const container = (<any>extra).container as Container;
    const authService = container.get<IAuthService>(Types.AuthService);
    return authService.login({ email, password });
  }
);
