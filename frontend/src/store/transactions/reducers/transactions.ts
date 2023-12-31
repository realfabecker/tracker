import { ActionStatus, Transaction } from "@core/domain/domain";
import { PayloadAction, createSlice } from "@reduxjs/toolkit";
import { State } from "@store/store";
import {
  getActionCreateTransaction,
  getActionDeleteTransaction,
  getActionLoadTransactionsList,
  getActionUpdateTransaction,
} from "@store/transactions/creators/transactions";

const initialState = {
  "transactions/list": {
    status: ActionStatus.IDLE,
    filters: {
      period: "this_month",
    },
    data: [],
  } as State<Transaction[]>,
  "transactions/add": {
    status: ActionStatus.IDLE,
  } as State,
  "transactions/edit": {
    status: ActionStatus.IDLE,
  } as State,
  "transactions/del": {
    status: ActionStatus.IDLE,
  } as State,
};

export type TransactionsState = typeof initialState;

export const transactionSlice = createSlice({
  name: "transactions",
  initialState: initialState,
  reducers: {
    filters_set: (state, action: PayloadAction<{ period: string }>) => {
      state["transactions/list"].filters = {
        period: action.payload!.period as string,
      };
    },
  },
  extraReducers: (builder) => {
    builder.addCase(
      getActionLoadTransactionsList.pending,
      (state: TransactionsState) => {
        state["transactions/list"]["status"] = ActionStatus.LOADING;
      }
    );
    builder.addCase(
      getActionLoadTransactionsList.fulfilled,
      (state: TransactionsState, action: any) => {
        state["transactions/list"] = {
          ...state["transactions/list"],
          status: ActionStatus.DONE,
          data: action.payload.data.items,
        };
      }
    );
    builder.addCase(
      getActionLoadTransactionsList.rejected,
      (state: TransactionsState) => {
        state["transactions/list"]["status"] = ActionStatus.ERROR;
      }
    );
    builder.addCase(
      getActionDeleteTransaction.pending,
      (state: TransactionsState) => {
        state["transactions/del"]["status"] = ActionStatus.LOADING;
      }
    );
    builder.addCase(
      getActionDeleteTransaction.fulfilled,
      (state: TransactionsState) => {
        state["transactions/del"] = {
          status: ActionStatus.DONE,
        };
      }
    );
    builder.addCase(
      getActionDeleteTransaction.rejected,
      (state: TransactionsState) => {
        state["transactions/del"]["status"] = ActionStatus.ERROR;
      }
    );
    builder.addCase(
      getActionCreateTransaction.pending,
      (state: TransactionsState) => {
        state["transactions/add"]["status"] = ActionStatus.LOADING;
      }
    );
    builder.addCase(
      getActionCreateTransaction.fulfilled,
      (state: TransactionsState) => {
        state["transactions/add"] = {
          status: ActionStatus.DONE,
        };
      }
    );
    builder.addCase(
      getActionCreateTransaction.rejected,
      (state: TransactionsState) => {
        state["transactions/add"]["status"] = ActionStatus.ERROR;
      }
    );
    builder.addCase(
      getActionUpdateTransaction.pending,
      (state: TransactionsState) => {
        state["transactions/edit"]["status"] = ActionStatus.LOADING;
      }
    );
    builder.addCase(
      getActionUpdateTransaction.fulfilled,
      (state: TransactionsState) => {
        state["transactions/edit"] = {
          status: ActionStatus.DONE,
        };
      }
    );
    builder.addCase(
      getActionUpdateTransaction.rejected,
      (state: TransactionsState) => {
        state["transactions/edit"]["status"] = ActionStatus.ERROR;
      }
    );
  },
});

export default transactionSlice.reducer;
