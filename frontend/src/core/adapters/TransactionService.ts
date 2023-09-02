import {
  PagedDTO,
  Transaction,
  ResponseDTO,
  TransactionPeriod,
  TransactionStatus,
} from "@core/domain/domain";
import { ITransactionService } from "@core/ports/ports";

export class TransactionService implements ITransactionService {
  constructor(
    private readonly baseUrl: string,
    private readonly token: string
  ) {
    this.baseUrl = "http://localhost:3001/api";
  }
  async fetchTransactions({
    period,
    status,
    limit,
    page,
  }: {
    limit: number;
    period: TransactionPeriod;
    status: TransactionStatus;
    page?: string;
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
      headers: { Authorization: `Bearer ${this.token}` },
    });
    if (resp.status !== 200) {
      throw new Error("Não foi possível listar os pagamentos");
    }
    return (await resp.json()) as ResponseDTO<PagedDTO<Transaction>>;
  }

  async addTransaction(body: Transaction): Promise<ResponseDTO<Transaction>> {
    const url = new URL(`${this.baseUrl}/wallet/payments`);
    const resp = await fetch(url.toString(), {
      method: "POST",
      headers: {
        Authorization: `Bearer ${this.token}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(body),
    });
    if (resp.status !== 200) {
      throw new Error("Não foi possível realizar o cadastro do pagamento");
    }
    return (await resp.json()) as ResponseDTO<Transaction>;
  }
}