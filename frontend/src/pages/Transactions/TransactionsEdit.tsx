import { useInjection } from "inversify-react";
import { FormEvent, useCallback, useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router";
import { Types } from "@core/container/types";
import { ActionStatus, Transaction } from "@core/domain/domain";
import { ITransactionService } from "@core/ports/ports";
import { useAppDispatch, useAppSelector } from "@store/store";
import {
  getActionDeleteTransaction,
  getActionUpdateTransaction,
} from "@store/transactions/creators/transactions";
import { asBrl } from "@core/lib/formatter";

function TransactionsEdit() {
  const dispatch = useAppDispatch();
  const navigate = useNavigate();
  const store = useAppSelector((d) => d.transactions["transactions/edit"]);
  const service = useInjection<ITransactionService>(Types.TransactionsService);

  const [name, setName] = useState("");
  const [desc, setDesc] = useState("");
  const [date, setDate] = useState(new Date().toISOString().slice(0, 10));
  const { id } = useParams();

  useEffect(() => {
    (async () => {
      const transaction = await service.getTransaction(id as string);
      setName(asBrl(transaction.data.value));
      setDesc(transaction.data.title);
      setDate(transaction.data.dueDate.slice(0, 10));
    })();
  }, [service, id]);

  const handleFormSubmit = useCallback(
    (e: FormEvent) => {
      e.preventDefault();
      const {
        groups: { value },
      } = name.match(/R\$\s(?<value>\d+(,\d{2})?)$/) as {
        groups: any;
      };
      const transaction: Partial<Transaction> = {
        title: desc,
        value: Number(value.replace(",", ".")),
      };
      const payload = { id: id as string, transaction, navigate };
      dispatch(getActionUpdateTransaction(payload));
    },
    [id, desc, name, dispatch, navigate]
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
            placeholder="200 Bethesda Skyrim"
            pattern="^R\$\s{1}\d+(,\d{2})?"
            required
          ></input>
          <input
            type="date"
            id="dueDate"
            value={date}
            onChange={(e) => setDate(e.target.value)}
            required
            readOnly={true}
            disabled={true}
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
        <div className="actions">
          <button
            id="update"
            type="submit"
            title="Salvar"
            disabled={store.status === ActionStatus.LOADING}
          >
            <span>&#128440;</span>
          </button>
          <button
            id="remove"
            title="Remover"
            onClick={() =>
              id &&
              dispatch(getActionDeleteTransaction(id)).then(() =>
                navigate("/transactions")
              )
            }
          >
            <span>{`\u267B`}</span>
          </button>
        </div>
        {store.error?.message && (
          <div className="error">{store.error.message}</div>
        )}
      </form>
    </>
  );
}
export default TransactionsEdit;
