import { ActionStatus, Transaction } from "@core/domain/domain";

export enum TransactionsActions {
  TRANSACTION_ADDED = "TRANSACTION_ADDED",

  TRANSACTION_LIST_LOADING = "TRANSACTION_LIST_LOADING",
  TRANSACTION_LIST_ERROR = "TRANSACTION_LIST_ERROR",
  TRANSACTION_LIST_LOADED = "TRANSACTION_LIST_LOADED",
}

export const transactionAdded = (transaction: Transaction) => ({
  type: TransactionsActions.TRANSACTION_ADDED,
  data: { transaction },
});

export const transactionListError = () => ({
  type: TransactionsActions.TRANSACTION_LIST_ERROR,
  data: { status: ActionStatus.ERROR },
});

export const transactionListLoading = () => ({
  type: TransactionsActions.TRANSACTION_LIST_LOADING,
  data: { status: ActionStatus.LOADING },
});

export const transactionListLoaded = (transactions: Transaction[]) => ({
  type: TransactionsActions.TRANSACTION_LIST_LOADED,
  data: { status: ActionStatus.DONE, transactions },
});
