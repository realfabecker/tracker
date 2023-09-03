import { ActionStatus } from "@core/domain/domain";
import "./App.css";

import { AuthService } from "@core/adapters/AuthService";
import Login from "@pages/Login";
import Transactions from "@pages/Transactions";
import { useAppSelector } from "@store/store";

function App() {
  const { status } = useAppSelector((state) => state["auth"]["auth/login"]);

  if (new AuthService().isLoggedIn() || status === ActionStatus.DONE) {
    return (
      <main>
        <Transactions />
      </main>
    );
  }

  return (
    <main>
      <Login />
    </main>
  );
}
export default App;
