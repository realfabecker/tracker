import { configureStore } from "@reduxjs/toolkit";
import thunk from "redux-thunk";

import { transactions } from "@store/transactions/reducers/transactions";
import { ActionStatus, IRootStore } from "@core/domain/domain";

const initialState: IRootStore = {
  transactions: {
    status: ActionStatus.IDLE,
    list: [],
  },
};

export const store = configureStore({
  middleware: [thunk],
  preloadedState: initialState,
  reducer: transactions,
});
