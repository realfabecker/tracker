import { useNavigate } from "react-router";
import { useEffect, useState } from "react";
import { asBrl, asDate } from "@core/lib/formatter";
import {
  ActionStatus,
  TransactionPeriod,
  TransactionType,
} from "@core/domain/domain";
import { getActionLoadTransactionsList } from "@store/transactions/creators/transactions";
import { useAppDispatch, useAppSelector } from "@store/store";

export function TransactionsList() {
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
        <div className="input-wrapper">
          <label htmlFor="period">Selecione o período:</label>
          <select
            id="period"
            value={period}
            onChange={(e) => setPeriod(e.target.value)}
          >
            <option value={TransactionPeriod.THIS_MONTH}>Esse Mês</option>
            <option value={TransactionPeriod.LAST_MONTH}>Mês Passado</option>
            <option value={TransactionPeriod.NEXT_MONTH}>Próximo Mês</option>
          </select>
        </div>
      </div>
      <div className="summary">
        <TransactionSummary />
      </div>
      <div className="items">
        {!transactions?.data?.length && (
          <div className={"transaction"}>
            <ul>
              <li className="name">
                <span>Descrição:</span>
                <span>--</span>
              </li>
              <li className="status">
                <span>Situação:</span>
                <span>--</span>
              </li>
            </ul>
            <ul>
              <li>
                <span>Valor:</span>
                <span>--</span>
              </li>
              <li className="datetime">
                <span>Vencimento:</span>
                <span>--</span>
              </li>
            </ul>
          </div>
        )}
        {(transactions?.data?.length || 0) > 0 &&
          transactions.data?.map((t) => (
            <div
              className={"transaction"}
              key={t.paymentId}
              onClick={() => {
                navigate(`/transactions/${t.paymentId}`);
              }}
            >
              <ul>
                <li className="name">
                  <span>Descrição:</span>
                  <span>{t.title}</span>
                </li>
                <li className="status">
                  <span>Situação:</span>
                  <span>{t.status.toUpperCase()}</span>
                </li>
              </ul>
              <ul>
                <li className={`price ${t.type}`}>
                  <span>Valor:</span>
                  <span>
                    {t.type === "expense" ? "-" : "+"} {asBrl(t.value)}
                  </span>
                </li>
                <li className="datetime">
                  <span>Vencimento:</span>
                  <span>{asDate(t.dueDate)}</span>
                </li>
              </ul>
            </div>
          ))}
      </div>
    </div>
  );
}

export function TransactionSummary() {
  const transactions = useAppSelector(
    (state) => state.transactions["transactions/list"]
  );

  const total = (transactions.data || []).reduce((acc, v) => {
    return acc + (v.type == TransactionType.EXPENSE ? -1 * v.value : v.value);
  }, 0);

  const [n, f] = total.toFixed(2).split(".");
  return (
    <div className={"summary"}>
      <h1>
        R$ {n}
        <span>,{f}</span>
      </h1>
    </div>
  );
}
