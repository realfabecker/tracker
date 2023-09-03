import { PagedDTO, Transaction, ResponseDTO } from "@core/domain/domain";

export interface ITransactionService {
  fetchTransactions({
    limit,
    period,
    status,
    page,
  }: {
    limit: number;
    period: string;
    status: string;
    page?: string;
  }): Promise<ResponseDTO<PagedDTO<Transaction>>>;
  addTransaction(body: Transaction): Promise<ResponseDTO<Transaction>>;
  deleteTransaction(id: string): Promise<void>;
}

export interface IAuthService {
  login: ({
    email,
    password,
  }: {
    email: string;
    password: string;
  }) => Promise<void>;
  isLoggedIn(): boolean;
  getAccessToken(): string | undefined;
}
