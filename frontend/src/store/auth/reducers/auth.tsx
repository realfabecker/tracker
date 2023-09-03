import { ActionStatus, Transaction } from "@core/domain/domain";
import { createSlice } from "@reduxjs/toolkit";
import { State } from "@store/store";
import { getActionAuthLogin } from "@store/auth/creators/auth";

const initialState = {
  "auth/login": {
    status: ActionStatus.IDLE,
  } as State<Transaction[]>,
};

export type AuthState = typeof initialState;

export const authSlice = createSlice({
  name: "auth",
  initialState: initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(getActionAuthLogin.pending, (state: AuthState) => {
      state["auth/login"]["status"] = ActionStatus.LOADING;
    });
    builder.addCase(getActionAuthLogin.fulfilled, (state: AuthState) => {
      state["auth/login"] = {
        status: ActionStatus.DONE,
      };
    });
    builder.addCase(getActionAuthLogin.rejected, (state: AuthState, action) => {
      state["auth/login"] = {
        status: ActionStatus.ERROR,
        error: {
          message: action?.error?.message || "Erro",
        },
      };
    });
  },
});

export default authSlice.reducer;
