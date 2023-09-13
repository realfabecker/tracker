import "reflect-metadata";

import { Provider } from "react-redux";
// import React from "react";
import ReactDOM from "react-dom/client";
import { Provider as Container } from "inversify-react";
import App from "@pages/App.tsx";
import { store } from "@store/store.ts";
import { container } from "@core/container";

ReactDOM.createRoot(document.getElementById("root")!).render(
  // <React.StrictMode>
  <Container container={container}>
    <Provider store={store}>
      <App />
    </Provider>
  </Container>
  // </React.StrictMode>
);
