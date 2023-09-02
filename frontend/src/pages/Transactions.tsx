import { FormEvent, useEffect, useState } from "react";
import { useDispatch, useSelector } from "react-redux";

import { asBrl, asDate } from "@core/lib/formatter";
import {
  ActionStatus,
  IRootStore,
  Transaction,
  TransactionStatus,
  TransactionType,
} from "@core/domain/domain";
import {
  getActionCreateTransaction,
  getActionLoadTransactionsList,
} from "@store/transactions/creators/transactions";

import "./Transactions.css";

function ItemAdd() {
  const dispatch = useDispatch();

  const [name, setName] = useState("");
  const [date, setDate] = useState(new Date().toISOString().slice(0, 10));
  const [desc, setDesc] = useState("");

  function handleFormSubmit(e: FormEvent) {
    e.preventDefault();

    const {
      groups: { type, value, title },
    } = name.match(/(?<type>\+|-)(?<value>\d+)\s{1}(?<title>\w+)/) as {
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
  }

  return (
    <form onSubmit={handleFormSubmit}>
      <div className="basic">
        <input
          type="text"
          id="title"
          value={name}
          onChange={(e) => setName(e.target.value)}
          placeholder="+200 Bethesda Starfield"
          pattern="^(\+|-)\d+\s{1}\w+"
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
      <button type="submit">Novo</button>
    </form>
  );
}

function ItemList() {
  const transactions = useSelector((state: IRootStore) => state.transactions);

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
        <span>Erro...</span>
      </div>
    );
  }

  return (
    <div className="transactions">
      {transactions.list.map((t) => (
        <div className="transaction" key={t.id}>
          <div className="left">
            <div className="name">{t.title}</div>
            <div className="description">{t.description}</div>
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
  const transactions = useSelector((state: IRootStore) => state.transactions);

  const total = transactions.list.reduce((acc, v) => {
    return acc + (v.type == TransactionType.EXPENSE ? -1 * v.value : v.value);
  }, 0);

  const type = total < 0 ? "expense" : "income";

  const [n, f] = total.toFixed(2).split(".");
  return (
    <h1 className={`summary ${type}`}>
      R$ {n}
      <span>,{f}</span>
    </h1>
  );
}

export default function Transactions() {
  const dispatch = useDispatch();

  useEffect(() => {
    //@ts-ignore
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
