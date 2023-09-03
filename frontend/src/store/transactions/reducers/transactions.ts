import { ActionStatus } from "@core/domain/domain";
import { TransactionsActions } from "@store/transactions/actions/transactions";

export const transactions = (state: any, action: any) => {
  switch (action.type) {
    case TransactionsActions.TRANSACTION_DEL_FULFILLED:
      return {
        ...state,
        "transactions/del": {
          status: ActionStatus.DONE,
        },
      };
    case TransactionsActions.TRANSACTION_DEL_REJECTED:
      return {
        ...state,
        "transactions/del": {
          status: ActionStatus.ERROR,
          error: action.error,
        },
      };
    case TransactionsActions.TRANSACTION_DEL_PENDING:
      return {
        ...state,
        "transactions/del": {
          status: ActionStatus.LOADING,
        },
      };
    case TransactionsActions.TRANSACTION_ADD_FULFILLED:
      return {
        ...state,
        "transactions/add": {
          status: ActionStatus.DONE,
        },
      };
    case TransactionsActions.TRANSACTION_ADD_REJECTED:
      return {
        ...state,
        "transactions/add": {
          status: ActionStatus.ERROR,
          error: action.error,
        },
      };
    case TransactionsActions.TRANSACTION_ADD_PENDING:
      return {
        ...state,
        "transactions/add": {
          status: ActionStatus.LOADING,
        },
      };
    case TransactionsActions.TRANSACTION_LIST_PENDING:
      return {
        ...state,
        "transactions/list": {
          ...state["transactions/list"],
          status: ActionStatus.LOADING,
        },
      };
    case TransactionsActions.TRANSACTION_LIST_FULFILLED:
      return {
        ...state,
        "transactions/list": {
          ...state["transactions/list"],
          data: action.payload.data.items,
          status: ActionStatus.DONE,
        },
      };
    case TransactionsActions.TRANSACTION_LIST_REJECTED:
      return {
        ...state,
        "transactions/list": {
          ...state["transactions/list"],
          status: ActionStatus.ERROR,
        },
      };
    default:
      return state;
  }
};
