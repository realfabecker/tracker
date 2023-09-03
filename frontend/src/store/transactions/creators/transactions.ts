import {
  Transaction,
  TransactionPeriod,
  TransactionStatus,
} from "@core/domain/domain";

import { AuthService } from "@core/adapters/AuthService";
import { TransactionService } from "@core/adapters/TransactionService";
import { createAsyncThunk } from "@reduxjs/toolkit";

export const getActionDeleteTransaction = createAsyncThunk(
  "transactions/del",
  async (transaction: string, { dispatch }) => {
    const baseUrl = import.meta.env.VITE_API_BASE_URL;

    const authService = new AuthService(baseUrl);
    const accessToken = authService.getAccessToken() as string;

    const transactionService = new TransactionService(baseUrl, accessToken);
    await transactionService.deleteTransaction(transaction);

    dispatch(getActionLoadTransactionsList());
  }
);

export const getActionCreateTransaction = createAsyncThunk(
  "transactions/add",
  async (transaction: Partial<Transaction>, { dispatch }) => {
    const baseUrl = import.meta.env.VITE_API_BASE_URL;

    const authService = new AuthService(baseUrl);
    const accessToken = authService.getAccessToken() as string;

    const transactionService = new TransactionService(baseUrl, accessToken);
    await transactionService.addTransaction(transaction);

    dispatch(getActionLoadTransactionsList());
  }
);

export const getActionLoadTransactionsList = createAsyncThunk(
  "transactions/list",
  async () => {
    const baseUrl = import.meta.env.VITE_API_BASE_URL;

    const authService = new AuthService(baseUrl);
    const accessToken = authService.getAccessToken() as string;

    const transactionService = new TransactionService(baseUrl, accessToken);
    return transactionService.fetchTransactions({
      limit: 50,
      period: TransactionPeriod.THIS_MONTH,
      status: TransactionStatus.ALL,
    });
  }
);
