import { FormEvent, useCallback, useEffect, useState } from "react";
import { asBrl, asDate } from "@core/lib/formatter";
import {
  ActionStatus,
  Transaction,
  TransactionStatus,
  TransactionType,
} from "@core/domain/domain";
import {
  getActionCreateTransaction,
  getActionDeleteTransaction,
  getActionLoadTransactionsList,
} from "@store/transactions/creators/transactions";
import { useAppDispatch, useAppSelector } from "@store/store";

import "./Transactions.css";

function ItemAdd() {
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
        /(?<type>\+|-)(?<value>\d+(\.\d{2})?)\s{1}(?<title>\w+)/
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

      //@ts-ignore
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

function ItemList() {
  const dispatch = useAppDispatch();

  const transactions = useAppSelector(
    (state) => state.transactions["transactions/list"]
  );

  if (transactions.status === ActionStatus.LOADING) {
    return (
      <div className="transactions loading">
        <span>Loading...</span>
      </div>
    );
  }

  if (transactions.status === ActionStatus.ERROR) {
    return (
      <div className="transactions error">
        <span>Erro ao consultar listagem de transações</span>
      </div>
    );
  }

  return (
    <div className="transactions">
      {transactions.data?.map((t) => (
        <div className="transaction" key={t.id}>
          <div className="left">
            <div>
              <button
                style={{ all: "unset" }}
                onClick={() => dispatch(getActionDeleteTransaction(t.id))}
              >
                <span className="trash">{`\u267B`}</span>
              </button>
            </div>
            <div>
              <div className="name">{t.title}</div>
              <div className="description">{t.description}</div>
            </div>
          </div>
          <div className="right">
            <div className={`price ${t.type}`}>
              {t.type === "expense" ? "-" : "+"}
              {asBrl(t.value)}
            </div>
            <div className="datetime">{asDate(t.dueDate)}</div>
          </div>
        </div>
      ))}
    </div>
  );
}

function ItemHeader() {
  const transactions = useAppSelector(
    (state) => state.transactions["transactions/list"]
  );

  const total = (transactions.data || []).reduce((acc, v) => {
    return acc + (v.type == TransactionType.EXPENSE ? -1 * v.value : v.value);
  }, 0);

  const [n, f] = total.toFixed(2).split(".");
  return (
    <h1 className={`summary`}>
      R$ {n}
      <span>,{f}</span>
    </h1>
  );
}

export default function Transactions() {
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(getActionLoadTransactionsList());
  }, [dispatch]);

  return (
    <main>
      <ItemHeader />
      <ItemAdd />
      <ItemList />
    </main>
  );
}
