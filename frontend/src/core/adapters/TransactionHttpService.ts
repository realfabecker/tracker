import {
  PagedDTO,
  Transaction,
  ResponseDTO,
  TransactionPeriod,
  TransactionStatus,
} from "@core/domain/domain";
import { ITransactionService } from "@core/ports/ports";
import { injectable } from "inversify";

@injectable()
export class TransactionsHttpService implements ITransactionService {
  constructor(
    private readonly baseUrl: string = import.meta.env.VITE_API_BASE_URL
  ) {}
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
    const url = new URL(`${this.baseUrl}/wallet/payments`);
    url.searchParams.append("limit", "" + limit);
    url.searchParams.append("period", period.toLowerCase());
    if (status.toLowerCase() !== "all") {
      url.searchParams.append("status", status.toLowerCase());
    }
    if (page) {
      url.searchParams.append("page_token", page);
    }
    const resp = await fetch(url.toString(), {
      method: "GET",
      headers: { Authorization: `Bearer ${token}` },
    });
    if (resp.status !== 200) {
      throw new Error("Não foi possível listar as transações");
    }
    return (await resp.json()) as ResponseDTO<PagedDTO<Transaction>>;
  }

  async addTransaction(
    body: Partial<Transaction>,
    token: string
  ): Promise<ResponseDTO<Transaction>> {
    const url = new URL(`${this.baseUrl}/wallet/payments`);
    const resp = await fetch(url.toString(), {
      method: "POST",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(body),
    });
    if (resp.status !== 200) {
      throw new Error("Não foi possível realizar o cadastro da transação");
    }
    return (await resp.json()) as ResponseDTO<Transaction>;
  }

  async deleteTransaction(id: string, token: string): Promise<void> {
    const url = new URL(`${this.baseUrl}/wallet/payments/${id}`);
    const resp = await fetch(url.toString(), {
      method: "DELETE",
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
    if (resp.status !== 204) {
      throw new Error("Não foi possível excluir a transação");
    }
  }
}
