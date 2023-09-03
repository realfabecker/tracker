import { configureStore } from "@reduxjs/toolkit";
import thunk from "redux-thunk";

import { transactions } from "@store/transactions/reducers/transactions";
import { ActionStatus, IRootStore } from "@core/domain/domain";
import { TypedUseSelectorHook, useDispatch, useSelector } from "react-redux";

const initialState: IRootStore = {
  "transactions/list": {
    status: ActionStatus.IDLE,
    data: [],
  },
  "transactions/add": {
    status: ActionStatus.IDLE,
  },
  "transactions/del": {
    status: ActionStatus.IDLE,
  },
};

export const store = configureStore({
  middleware: [thunk],
  preloadedState: initialState,
  reducer: transactions,
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export const useAppDispatch = () => useDispatch<AppDispatch>();
export const useAppSelector: TypedUseSelectorHook<RootState> = useSelector;
