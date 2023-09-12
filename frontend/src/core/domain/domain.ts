export interface ResponseDTO<T> {
  status: "success" | "error";
  data: T;
}

export interface PagedDTO<T> {
  page_count: number;
  items: T[];
  page_token?: string;
  has_more: boolean;
}

export enum TransactionStatus {
  ALL = "all",
  PENDING = "pending",
  PAID = "paid",
  CANCELLED = "cancelled",
}

export enum TransactionPeriod {
  THIS_WEEK = "this_week",
  THIS_MONTH = "this_month",
  LAST_MONTH = "last_month",
  NEXT_MONTH = "next_month",
}

export interface Transaction {
  id: string;
  userId: string;
  title: string;
  type: string;
  description: string;
  value: number;
  dueDate: string;
  status: TransactionStatus;
  createdAt: string;
}

export enum TransactionType {
  EXPENSE = "expense",
  INCOME = "income",
}

export enum ActionStatus {
  IDLE = "idle",
  DONE = "done",
  ERROR = "error",
  LOADING = "loading",
}

export interface LoginDTO {
  RefreshToken: string;
  AccessToken: string;
}

export enum RoutesEnum {
  Login = "/login",
  Transactions = "/transactions",
}
