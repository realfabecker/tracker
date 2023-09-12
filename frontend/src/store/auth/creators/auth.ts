import { createAsyncThunk } from "@reduxjs/toolkit";
import { Container } from "inversify";
import { Types } from "@core/container/types";
import { IAuthService } from "@core/ports/ports";
import { NavigateFunction } from "react-router";
import { RoutesEnum } from "@core/domain/domain";

export const getActionAuthLogin = createAsyncThunk(
  "auth/login",
  async (
    {
      email,
      password,
      navigate,
    }: { email: string; password: string; navigate: NavigateFunction },
    { extra }
  ) => {
    const container = (<any>extra).container as Container;
    const authService = container.get<IAuthService>(Types.AuthService);
    await authService.login({ email, password });
    navigate(RoutesEnum.Transactions);
  }
);

export const getActionAuthLogout = createAsyncThunk(
  "auth/logout",
  async ({ navigate }: { navigate: NavigateFunction }, { extra }) => {
    const container = (<any>extra).container as Container;
    const authService = container.get<IAuthService>(Types.AuthService);
    authService.logout();
    navigate(RoutesEnum.Login);
  }
);
