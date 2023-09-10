import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "@pages/Login";
import Transactions from "@pages/Transactions";
import { PrivLayout, PubLayout } from "@pages/Layout";

import "./App.css";
import TransactionsAdd from "./TransactionsAdd";
import TransactionsEdit from "./TransactionsEdit";

function App() {
  return (
    <main>
      <BrowserRouter>
        <Routes>
          <Route element={<PubLayout />}>
            <Route path="/login" element={<Login />} />
          </Route>
          <Route element={<PrivLayout />}>
            <Route path="/transactions" element={<Transactions />}>
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
