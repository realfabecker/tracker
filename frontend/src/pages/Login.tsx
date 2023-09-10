import { useState } from "react";
import { useAppDispatch, useAppSelector } from "@store/store";
import { getActionAuthLogin } from "@store/auth/creators/auth";
import { ActionStatus } from "@core/domain/domain";

import "./Login.css";
import { useNavigate } from "react-router";

export default function Login() {
  const dispatch = useAppDispatch();
  const navigate = useNavigate();

  const store = useAppSelector((state) => state["auth"]["auth/login"]);

  const [email, setEmail] = useState("");
  const [passw, setPassw] = useState("");

  return (
    <main>
      <form
        id="login"
        onSubmit={(e) => {
          e.preventDefault();
          dispatch(getActionAuthLogin({ email, password: passw, navigate }));
        }}
      >
        <div className="basic">
          <input
            type="email"
            value={email}
            placeholder="E-mail"
            onChange={(e) => setEmail(e.target.value)}
          />
          <input
            type="password"
            value={passw}
            placeholder="Password"
            onChange={(e) => setPassw(e.target.value)}
          />
        </div>
        <button type="submit" disabled={store.status === ActionStatus.LOADING}>
          {store.status === ActionStatus.LOADING ? "Loading..." : "Login"}
        </button>
        {store.error?.message && (
          <div className="error">{store.error.message}</div>
        )}
      </form>
    </main>
  );
}
