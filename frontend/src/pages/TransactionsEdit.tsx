import { useInjection } from "inversify-react";
import { FormEvent, useCallback, useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router";
import { Types } from "@core/container/types";
import { ActionStatus, Transaction } from "@core/domain/domain";
import { ITransactionService } from "@core/ports/ports";
import { useAppDispatch, useAppSelector } from "@store/store";
import { getActionUpdateTransaction } from "@store/transactions/creators/transactions";

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
      const transaction = await service.getTransaction(id + "", "");
      setName(`${transaction.data.value} ${transaction.data.title}`);
      setDesc(transaction.data.description);
      setDate(transaction.data.dueDate.slice(0, 10));
    })();
  }, [service, id]);

  const handleFormSubmit = useCallback(
    (e: FormEvent) => {
      e.preventDefault();
      const {
        groups: { value, title },
      } = name.match(/(?<value>\d+(\.\d{2})?)\s{1}(?<title>.+)/) as {
        groups: any;
      };
      const transaction: Partial<Transaction> = {
        title: title,
        description: desc,
        value: Number(value),
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
            placeholder="+200 Bethesda Starfield"
            pattern="^\d+(\.\d{2})?\s{1}.+"
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
        <button type="submit" disabled={store.status === ActionStatus.LOADING}>
          {store.status === ActionStatus.LOADING ? "Loading..." : "Atualizar"}
        </button>
        {store.error?.message && (
          <div className="error">{store.error.message}</div>
        )}
      </form>
    </>
  );
}
export default TransactionsEdit;
