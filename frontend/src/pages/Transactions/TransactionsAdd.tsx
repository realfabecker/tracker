import {
  ActionStatus,
  Transaction,
  TransactionStatus,
  TransactionType,
} from "@core/domain/domain";
import { useAppDispatch, useAppSelector } from "@store/store";
import { getActionCreateTransaction } from "@store/transactions/creators/transactions";
import { FormEvent, useCallback, useState } from "react";

const initialState = {
  name: "",
  desc: "",
  date: new Date().toISOString().slice(0, 10),
};

function TransactionsAdd() {
  const dispatch = useAppDispatch();
  const store = useAppSelector((d) => d.transactions["transactions/add"]);

  const [d, setD] = useState(initialState);
  const handleFormSubmit = useCallback(
    (e: FormEvent) => {
      e.preventDefault();

      const {
        groups: { type, value },
      } = d.name.match(/(?<type>[+-])?(?<value>\d+(,\d{2})?)$/) as {
        groups: any;
      };

      const transaction: Partial<Transaction> = {
        type: type === "+" ? TransactionType.INCOME : TransactionType.EXPENSE,
        title: d.desc,
        description: d.desc,
        value: Number(value.replace(",", ".")),
        dueDate: new Date(d.date).toISOString(),
        status: TransactionStatus.PAID,
      };
      dispatch(getActionCreateTransaction(transaction)).then(() => {
        setD(initialState);
      });
    },
    [d, dispatch]
  );

  return (
    <>
      <form id="transaction" onSubmit={handleFormSubmit}>
        <div className="input-wrapper">
          <label htmlFor="title">
            Valor <span>(valor do lançamento)</span>
          </label>
          <input
            type="text"
            id="title"
            value={d.name}
            onChange={(e) =>
              setD((prevState) => ({
                ...prevState,
                name: e.target.value,
              }))
            }
            placeholder="R$ 2,50"
            title="Ex.: 2,50"
            pattern="^(\+|-)?\d+(,\d{2})?$"
            required
          ></input>
        </div>
        <div className="input-wrapper">
          <label htmlFor="description">Descrição</label>
          <input
            type="text"
            id="description"
            value={d.desc}
            onChange={(e) =>
              setD((prevState) => ({
                ...prevState,
                desc: e.target.value,
              }))
            }
            placeholder="Título do lançamento"
            required
          ></input>
        </div>
        <div className="input-wrapper">
          <label htmlFor="dueDate">Vencimento</label>
          <input
            type="date"
            id="dueDate"
            value={d.date}
            onChange={(e) =>
              setD((prevState) => ({
                ...prevState,
                date: e.target.value,
              }))
            }
            required
          />
        </div>

        <div className="actions">
          <button
            type="submit"
            disabled={store.status === ActionStatus.LOADING}
          >
            {store.status === ActionStatus.LOADING ? "Loading..." : "Novo"}
          </button>
        </div>

        {store.error?.message && (
          <div className="error">{store.error.message}</div>
        )}
      </form>
    </>
  );
}

export default TransactionsAdd;
