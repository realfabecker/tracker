import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import { PrivLayout, PubLayout } from "@pages/Layout";
import TransactionsAdd from "@pages/Transactions/TransactionsAdd";
import TransactionsEdit from "@pages/Transactions/TransactionsEdit";

import "./App.css";
import { RoutesEnum } from "@core/domain/domain";
import Dashboard from "@pages/Dashboard/Dashboard.tsx";
import Login from "@pages/Login/Login.tsx";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route element={<PubLayout />}>
          <Route path={RoutesEnum.Login} element={<Login />} />
        </Route>
        <Route element={<PrivLayout />}>
          <Route path={RoutesEnum.Transactions} element={<Dashboard />}>
            <Route path="" element={<TransactionsAdd />} />
            <Route path=":id" element={<TransactionsEdit />} />
          </Route>
        </Route>
        <Route path="*" element={<Navigate to={RoutesEnum.Transactions} />} />
      </Routes>
    </BrowserRouter>
  );
}
export default App;
