import {
  ActionStatus,
  Transaction,
  TransactionStatus,
  TransactionType,
} from "@core/domain/domain";
import { useAppDispatch, useAppSelector } from "@store/store";
import { getActionCreateTransaction } from "@store/transactions/creators/transactions";
import { FormEvent, useCallback, useState } from "react";

function TransactionsAdd() {
  const dispatch = useAppDispatch();
  const store = useAppSelector((d) => d.transactions["transactions/add"]);

  const [name, setName] = useState("");
  const [desc, setDesc] = useState("");
  const [date, setDate] = useState(new Date().toISOString().slice(0, 10));

  const handleFormSubmit = useCallback(
    (e: FormEvent) => {
      e.preventDefault();

      const {
        groups: { type, value, title },
      } = name.match(
        /(?<type>\+|-)(?<value>\d+(\.\d{2})?)\s{1}(?<title>.+)/
      ) as {
        groups: any;
      };

      const transaction: Partial<Transaction> = {
        type: type === "+" ? TransactionType.INCOME : TransactionType.EXPENSE,
        title: title,
        description: desc,
        value: Number(value),
        dueDate: new Date(date).toISOString(),
        status: TransactionStatus.PENDING,
      };
      dispatch(getActionCreateTransaction(transaction));
    },
    [date, desc, name, dispatch]
  );

  return (
    <>
      <form id="transaction" onSubmit={handleFormSubmit}>
        <div className="basic">
          <input
            type="text"
            id="title"
            value={name}
            onChange={(e) => setName(e.target.value)}
            placeholder="+200 Bethesda Starfield"
            pattern="^(\+|-)\d+(\.\d{2})?\s{1}.+"
            required
          ></input>
          <input
            type="date"
            id="dueDate"
            value={date}
            onChange={(e) => setDate(e.target.value)}
            required
          />
        </div>
        <div className="description">
          <input
            type="text"
            id="description"
            value={desc}
            onChange={(e) => setDesc(e.target.value)}
            placeholder="descrição"
            required
          ></input>
        </div>
        <button type="submit" disabled={store.status === ActionStatus.LOADING}>
          {store.status === ActionStatus.LOADING ? "Loading..." : "Novo"}
        </button>
        {store.error?.message && (
          <div className="error">{store.error.message}</div>
        )}
      </form>
    </>
  );
}
export default TransactionsAdd;
