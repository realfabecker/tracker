import { useInjection } from "inversify-react";
import { ActionStatus } from "@core/domain/domain";
import Login from "@pages/Login";
import Transactions from "@pages/Transactions";
import { useAppSelector } from "@store/store";
import { IAuthService } from "@core/ports/ports";
import { Types } from "@core/container/types";

import "./App.css";

function App() {
  const { status } = useAppSelector((state) => state["auth"]["auth/login"]);
  const service = useInjection<IAuthService>(Types.AuthService);

  if (service.isLoggedIn() && status === ActionStatus.DONE) {
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
