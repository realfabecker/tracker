import { Outlet, useNavigate } from "react-router";
import { useEffect, useState } from "react";
import { asBrl, asDate } from "@core/lib/formatter";
import {
  ActionStatus,
  TransactionPeriod,
  TransactionType,
} from "@core/domain/domain";
import {
  getActionDeleteTransaction,
  getActionLoadTransactionsList,
} from "@store/transactions/creators/transactions";
import { useAppDispatch, useAppSelector } from "@store/store";
import { getActionAuthLogout } from "@store/auth/creators/auth";

import "./Transactions.css";

function ItemList() {
  const dispatch = useAppDispatch();
  const navigate = useNavigate();
  const [period, setPeriod] = useState("this_month");

  const transactions = useAppSelector(
    (state) => state.transactions["transactions/list"]
  );

  useEffect(() => {
    dispatch(getActionLoadTransactionsList(period));
  }, [dispatch, period]);

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
      <div className="filters">
        <select value={period} onChange={(e) => setPeriod(e.target.value)}>
          <option value={TransactionPeriod.THIS_MONTH}>This Month</option>
          <option value={TransactionPeriod.LAST_MONTH}>Last Month</option>
          <option value={TransactionPeriod.NEXT_MONTH}>Next Month</option>
        </select>
      </div>

      {transactions.data?.map((t) => (
        <div className="transaction" key={t.id}>
          <div className="left">
            <div style={{ display: "flex", flexDirection: "column" }}>
              <button
                id="edit"
                style={{ all: "unset" }}
                title="Edit"
                onClick={() => navigate(`/transactions/${t.id}`)}
              >
                <span className="edit">{`\u270E`}</span>
              </button>
              <button
                id="remove"
                title="Remove"
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
  const dispatch = useAppDispatch();
  const navigate = useNavigate();

  const transactions = useAppSelector(
    (state) => state.transactions["transactions/list"]
  );

  const total = (transactions.data || []).reduce((acc, v) => {
    return acc + (v.type == TransactionType.EXPENSE ? -1 * v.value : v.value);
  }, 0);

  const [n, f] = total.toFixed(2).split(".");
  return (
    <header>
      <h1 className={`summary`}>
        R$ {n}
        <span>,{f}</span>
      </h1>
      <button
        title="Logout"
        onClick={() => dispatch(getActionAuthLogout({ navigate }))}
      >{`\u27F6`}</button>
    </header>
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
      <Outlet />
      <ItemList />
    </main>
  );
}
