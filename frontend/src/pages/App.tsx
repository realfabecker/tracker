import "./App.css";

import Transactions from "@pages/Transactions";
import Protected from "@pages/Protected";

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
