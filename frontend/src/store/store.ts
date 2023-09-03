import { configureStore } from "@reduxjs/toolkit";
import thunk from "redux-thunk";
import { TypedUseSelectorHook, useDispatch, useSelector } from "react-redux";
import { ActionStatus } from "@core/domain/domain";
import transactionsSlice from "@store/transactions/reducers/transactions";
import authSlice from "@store/auth/reducers/auth";

export interface State<T = any> {
  data?: T;
  status: ActionStatus;
  error?: { message: string };
}

export const store = configureStore({
  middleware: [thunk],
  reducer: {
    transactions: transactionsSlice,
    auth: authSlice,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;
