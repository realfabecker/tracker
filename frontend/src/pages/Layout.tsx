import { Types } from "@core/container/types";
import { IAuthService } from "@core/ports/ports";
import { useInjection } from "inversify-react";
import { Navigate, Outlet } from "react-router";

export const PubLayout = () => {
  const service = useInjection<IAuthService>(Types.AuthService);

  if (service.isLoggedIn()) {
    return <Navigate to="/" />;
  }

  return <Outlet />;
};

export const PrivLayout = () => {
  const service = useInjection<IAuthService>(Types.AuthService);

  if (!service.isLoggedIn()) {
    return <Navigate to="/login" />;
  }

  return <Outlet />;
};
