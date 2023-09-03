import {
  Transaction,
  TransactionPeriod,
  TransactionStatus,
} from "@core/domain/domain";
import {
  transactionListError,
  transactionListLoaded,
  transactionListLoading,
} from "../actions/transactions";
import { AuthService } from "@core/adapters/AuthService";
import { TransactionService } from "@core/adapters/TransactionService";

export const getActionCreateTransaction = (transaction: Transaction) => {
  return async () => {
    const baseUrl = import.meta.env.VITE_API_BASE_URL;

    const authService = new AuthService(baseUrl);
    const accessToken = authService.getAccessToken() as string;

    const transactionService = new TransactionService(baseUrl, accessToken);
    await transactionService.addTransaction(transaction);
  };
};

export const getActionLoadTransactionsList = () => {
  return async (dispatch: any) => {
    dispatch(transactionListLoading());

    try {
      const baseUrl = import.meta.env.VITE_API_BASE_URL;

      const authService = new AuthService(baseUrl);
      const accessToken = authService.getAccessToken() as string;

      const transactionService = new TransactionService(baseUrl, accessToken);
      const transactions = await transactionService.fetchTransactions({
        limit: 10,
        period: TransactionPeriod.THIS_MONTH,
        status: TransactionStatus.ALL,
      });

      dispatch(transactionListLoaded(transactions.data.items));
    } catch (e) {
      dispatch(transactionListError());
    }
  };
};
