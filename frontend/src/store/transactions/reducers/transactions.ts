import { ActionStatus } from "@core/domain/domain";
import { TransactionsActions } from "@store/transactions/actions/transactions";

export const transactions = (state: any, action: any) => {
  switch (action.type) {
    case TransactionsActions.TRANSACTION_ADDED:
      return {
        ...state,
        transactions: {
          ...state.transactions,
          list: [...state.transactions.list, action.data.transaction],
        },
      };
    case TransactionsActions.TRANSACTION_LIST_LOADING:
      return {
        ...state,
        transactions: {
          ...state.transactions,
          status: ActionStatus.LOADING,
        },
      };
    case TransactionsActions.TRANSACTION_LIST_ERROR:
      return {
        ...state,
        transactions: {
          ...state.transactions,
          status: ActionStatus.ERROR,
        },
      };
    case TransactionsActions.TRANSACTION_LIST_LOADED:
      return {
        ...state,
        transactions: {
          ...state.transactions,
          list: action.data.transactions,
          status: ActionStatus.DONE,
        },
      };
    default:
      return state;
  }
};
