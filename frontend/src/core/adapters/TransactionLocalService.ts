import { injectable } from "inversify";
import {
  PagedDTO,
  Transaction,
  ResponseDTO,
  TransactionPeriod,
  TransactionStatus,
} from "@core/domain/domain";
import { ITransactionService } from "@core/ports/ports";

@injectable()
export class TransactionsLocalService implements ITransactionService {
  async fetchTransactions({
    period,
    status,
    limit,
    page,
    token,
  }: {
    limit: number;
    period: TransactionPeriod;
    status: TransactionStatus;
    page?: string;
    token: string;
  }): Promise<ResponseDTO<PagedDTO<Transaction>>> {
    console.log({ period, status, page, limit, token });
    const items = JSON.parse(localStorage.getItem("transactions") || "[]");
    return {
      status: "success",
      data: {
        page_count: items.length,
        items: items,
        has_more: false,
      },
    };
  }

  async addTransaction(
    body: Partial<Transaction>,
    token: string
  ): Promise<ResponseDTO<Transaction>> {
    console.log({ body, token });
    const items: Partial<Transaction>[] = JSON.parse(
      localStorage.getItem("transactions") || "[]"
    );
    items.push({
      id: Math.random().toString(32).slice(2),
      ...body,
    });
    localStorage.setItem("transactions", JSON.stringify(items));
    return {
      status: "success",
      data: items[items.length - 1] as Transaction,
    };
  }

  async deleteTransaction(id: string, token: string): Promise<void> {
    console.log({ id, token });
    const items: Partial<Transaction>[] = JSON.parse(
      localStorage.getItem("transactions") || "[]"
    );
    const data = items.filter((i) => i.id !== id);
    localStorage.setItem("transactions", JSON.stringify(data));
  }

  async getTransaction(id: string): Promise<ResponseDTO<Transaction>> {
    const items: Partial<Transaction>[] = JSON.parse(
      localStorage.getItem("transactions") || "[]"
    );
    const data = items.filter((i) => i.id == id);
    return {
      status: "success",
      data: data[0] as Transaction,
    };
  }

  async editTransaction(
    id: string,
    body: Partial<Transaction>,
    token: string
  ): Promise<ResponseDTO<Transaction>> {
    console.log({ id, body, token });
    const items: Partial<Transaction>[] = JSON.parse(
      localStorage.getItem("transactions") || "[]"
    );
    const index = items.findIndex((i) => i.id == id);
    items[index] = {
      ...items[index],
      ...body,
    };
    localStorage.setItem("transactions", JSON.stringify(items));
    return {
      status: "success",
      data: items[index] as Transaction,
    };
  }
}
