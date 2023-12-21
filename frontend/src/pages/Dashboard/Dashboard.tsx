import { TransactionsList } from "@pages/Transactions/TransactionsList.tsx";
import { Outlet } from "react-router";
import { useAppDispatch } from "@store/store.ts";
import { useEffect } from "react";
import { getActionLoadTransactionsList } from "@store/transactions/creators/transactions.ts";

function Dashboard() {
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(getActionLoadTransactionsList());
  }, [dispatch]);
  return (
    <>
      <div id="app">
        <aside>
          <Outlet />
        </aside>
        <main>
          <TransactionsList />
        </main>
      </div>
    </>
  );
}

export default Dashboard;
