import { Container as InversifyContainer } from "inversify";

import { Types } from "@core/container/types";
import { AuthHttpService } from "@core/adapters/AuthHttpService";
import { TransactionsHttpService } from "@core/adapters/TransactionHttpService";
import { TransactionsLocalService } from "@core/adapters/TransactionLocalService";
import { AuthLocalService } from "@core/adapters/AuthLocalService";

export const container = new InversifyContainer();
if (import.meta.env.MODE === "development") {
  container.bind(Types.AuthService).to(AuthLocalService);
  container.bind(Types.TransactionsService).to(TransactionsLocalService);
} else {
  container.bind(Types.AuthService).to(AuthHttpService);
  container.bind(Types.TransactionsService).to(TransactionsHttpService);
}
