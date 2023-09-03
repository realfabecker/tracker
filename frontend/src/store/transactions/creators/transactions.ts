import {
  Transaction,
  TransactionPeriod,
  TransactionStatus,
} from "@core/domain/domain";

import { createAsyncThunk } from "@reduxjs/toolkit";
import { Container } from "inversify";
import { IAuthService, ITransactionService } from "@core/ports/ports";
import { Types } from "@core/container/types";

export const getActionDeleteTransaction = createAsyncThunk(
  "transactions/del",
  async (transaction: string, { dispatch, extra }) => {
    const container = (<any>extra).container as Container;

    const authService = container.get<IAuthService>(Types.AuthService);
    const accessToken = authService.getAccessToken() as string;

    const tranService = container.get<ITransactionService>(
      Types.TransactionsService
    );
    await tranService.deleteTransaction(transaction, accessToken);
    dispatch(getActionLoadTransactionsList());
  }
);

export const getActionCreateTransaction = createAsyncThunk(
  "transactions/add",
  async (transaction: Partial<Transaction>, { dispatch, extra }) => {
    const container = (<any>extra).container as Container;

    const authService = container.get<IAuthService>(Types.AuthService);
    const accessToken = authService.getAccessToken() as string;

    const tranService = container.get<ITransactionService>(
      Types.TransactionsService
    );
    await tranService.addTransaction(transaction, accessToken);
    dispatch(getActionLoadTransactionsList());
  }
);

export const getActionLoadTransactionsList = createAsyncThunk(
  "transactions/list",
  async (_, { extra }) => {
    const container = (<any>extra).container as Container;

    const authService = container.get<IAuthService>(Types.AuthService);
    const accessToken = authService.getAccessToken() as string;

    const tranService = container.get<ITransactionService>(
      Types.TransactionsService
    );
    return tranService.fetchTransactions({
      limit: 50,
      period: TransactionPeriod.THIS_MONTH,
      status: TransactionStatus.ALL,
      token: accessToken,
    });
  }
);
