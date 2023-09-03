import { PagedDTO, Transaction, ResponseDTO } from "@core/domain/domain";

export interface ITransactionService {
  fetchTransactions({
    limit,
    period,
    status,
    page,
    token,
  }: {
    limit: number;
    period: string;
    status: string;
    page?: string;
    token: string;
  }): Promise<ResponseDTO<PagedDTO<Transaction>>>;
  addTransaction(
    body: Partial<Transaction>,
    token: string
  ): Promise<ResponseDTO<Transaction>>;
  deleteTransaction(id: string, token: string): Promise<void>;
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
