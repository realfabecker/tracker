import { AuthService } from "@core/adapters/AuthService";

export default function Protected({
  children,
}: {
  children?: React.ReactElement;
}) {
  const service = new AuthService();
  if (!service.isLoggedIn()) {
    return <div>Forbidden</div>;
  }
  return children;
}
