import "./App.css";

import Transactions from "@pages/Transactions";
import Protected from "./Protected";

function App() {
  return (
    <main>
      <Protected>
        <Transactions />
      </Protected>
    </main>
  );
}

export default App;
