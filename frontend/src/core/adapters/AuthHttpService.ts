import { LoginDTO, ResponseDTO } from "@core/domain/domain";
import { IAuthService } from "@core/ports/ports";
import { injectable } from "inversify";

@injectable()
export class AuthHttpService implements IAuthService {
  constructor(
    private readonly baseUrl: string = import.meta.env.VITE_API_BASE_URL,
    private readonly storage = localStorage,
    private readonly authKey = "tracker"
  ) {}

  async login({
    email,
    password,
  }: {
    email: string;
    password: string;
  }): Promise<void> {
    const resp = await fetch(`${this.baseUrl}/auth/login`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password }),
    });
    if (resp.status !== 200) {
      throw new Error("Credenciais inválidas");
    }
    const auth = (await resp.json()) as ResponseDTO<LoginDTO>;
    if (!auth.data.AccessToken || !auth.data.RefreshToken) {
      throw new Error("Não foi possível capturar as credenciais");
    }
    this.storage.setItem(this.authKey, JSON.stringify(auth.data));
  }

  getAccessToken(): string | undefined {
    const data = this.storage.getItem(this.authKey);
    if (!data) return;
    const auth = JSON.parse(data) as { AccessToken: string };
    return auth.AccessToken;
  }

  isLoggedIn(): boolean {
    const data = this.storage.getItem(this.authKey);
    if (!data) return false;

    const auth = JSON.parse(data) as { AccessToken: string };
    if (!auth.AccessToken) return false;

    const [, body] = auth.AccessToken.split(".");
    if (!body) return false;

    try {
      const token = JSON.parse(atob(body)) as {
        exp: number;
        [key: string]: unknown;
      };
      return new Date(token.exp * 1000).getTime() > new Date().getTime();
    } catch (e) {
      return false;
    }
  }

  logout(): void {
    this.storage.removeItem(this.authKey);
  }
}
