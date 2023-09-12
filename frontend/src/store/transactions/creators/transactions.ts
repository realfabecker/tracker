import {
  RoutesEnum,
  Transaction,
  TransactionPeriod,
  TransactionStatus,
} from "@core/domain/domain";

import { createAsyncThunk } from "@reduxjs/toolkit";
import { Container } from "inversify";
import { ITransactionService } from "@core/ports/ports";
import { Types } from "@core/container/types";
import { NavigateFunction } from "react-router";

export const getActionDeleteTransaction = createAsyncThunk(
  "transactions/del",
  async (transaction: string, { dispatch, extra }) => {
    const container = (<any>extra).container as Container;
    const service = container.get<ITransactionService>(
      Types.TransactionsService
    );
    await service.deleteTransaction(transaction);
    dispatch(getActionLoadTransactionsList());
  }
);

export const getActionCreateTransaction = createAsyncThunk(
  "transactions/add",
  async (transaction: Partial<Transaction>, { dispatch, extra }) => {
    const container = (<any>extra).container as Container;
    const service = container.get<ITransactionService>(
      Types.TransactionsService
    );
    await service.addTransaction(transaction);
    dispatch(getActionLoadTransactionsList());
  }
);

export const getActionLoadTransactionsList = createAsyncThunk(
  "transactions/list",
  async (_, { extra }) => {
    const container = (<any>extra).container as Container;
    const service = container.get<ITransactionService>(
      Types.TransactionsService
    );
    return service.fetchTransactions({
      limit: 50,
      period: TransactionPeriod.THIS_MONTH,
      status: TransactionStatus.ALL,
    });
  }
);

export const getActionUpdateTransaction = createAsyncThunk(
  "transactions/edit",
  async (
    props: {
      id: string;
      transaction: Partial<Transaction>;
      navigate: NavigateFunction;
    },
    { dispatch, extra }
  ) => {
    const container = (<any>extra).container as Container;
    const service = container.get<ITransactionService>(
      Types.TransactionsService
    );
    await service.editTransaction(props.id, props.transaction);
    dispatch(getActionLoadTransactionsList());
    props.navigate(RoutesEnum.Transactions);
  }
);
