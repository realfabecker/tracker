import { Container as InversifyContainer } from "inversify";
import { AuthService } from "@core/adapters/AuthService";
import { TransactionService } from "@core/adapters/TransactionService";
import { Types } from "@core/container/types";

export const container = new InversifyContainer();
container.bind(Types.AuthService).to(AuthService);
container.bind(Types.TransactionsService).to(TransactionService);
