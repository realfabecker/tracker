import { TypedUseSelectorHook, useDispatch, useSelector } from "react-redux";
import { configureStore } from "@reduxjs/toolkit";
import { ActionStatus } from "@core/domain/domain";
import transactionsSlice from "@store/transactions/reducers/transactions";
import authSlice from "@store/auth/reducers/auth";
import { container } from "@core/container";

export interface State<T = any> {
  data?: T;
  status: ActionStatus;
  error?: { message: string };
}

export const store = configureStore({
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      thunk: { extraArgument: { container } },
    }),
  reducer: {
    transactions: transactionsSlice,
    auth: authSlice,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;
