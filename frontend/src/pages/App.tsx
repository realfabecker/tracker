import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "@pages/Login";
import Transactions from "@pages/Transactions/Transactions";
import { PrivLayout, PubLayout } from "@pages/Layout";
import TransactionsAdd from "@pages/Transactions/TransactionsAdd";
import TransactionsEdit from "@pages/Transactions/TransactionsEdit";

import "./App.css";
import { RoutesEnum } from "@core/domain/domain";

function App() {
  return (
    <main>
      <BrowserRouter>
        <Routes>
          <Route element={<PubLayout />}>
            <Route path={RoutesEnum.Login} element={<Login />} />
          </Route>
          <Route element={<PrivLayout />}>
            <Route path={RoutesEnum.Transactions} element={<Transactions />}>
              <Route path="" element={<TransactionsAdd />} />
              <Route path=":id" element={<TransactionsEdit />} />
            </Route>
          </Route>
        </Routes>
      </BrowserRouter>
    </main>
  );
}
export default App;
