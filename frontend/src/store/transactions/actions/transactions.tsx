export enum TransactionsActions {
  TRANSACTION_DEL_PENDING = "transactions/del/pending",
  TRANSACTION_DEL_REJECTED = "transactions/del/rejected",
  TRANSACTION_DEL_FULFILLED = "transactions/del/fulfilled",

  TRANSACTION_ADD_PENDING = "transactions/add/pending",
  TRANSACTION_ADD_REJECTED = "transactions/add/rejected",
  TRANSACTION_ADD_FULFILLED = "transactions/add/fulfilled",

  TRANSACTION_LIST_PENDING = "transactions/list/pending",
  TRANSACTION_LIST_REJECTED = "transactions/list/rejected",
  TRANSACTION_LIST_FULFILLED = "transactions/list/fulfilled",
}
