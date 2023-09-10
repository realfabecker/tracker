import { LoginDTO, ResponseDTO } from "@core/domain/domain";
import { IAuthService } from "@core/ports/ports";
import { injectable } from "inversify";

@injectable()
export class AuthLocalService implements IAuthService {
  constructor(
    private readonly baseUrl: string = import.meta.env.VITE_API_BASE_URL,
    private readonly storage = localStorage,
    private readonly authKey = "tracker_local"
  ) {}

  async login({
    email,
    password,
  }: {
    email: string;
    password: string;
  }): Promise<void> {
    console.log({ baseUrl: this.baseUrl });
    const token = btoa(`${email}:${password}`);
    const data: ResponseDTO<LoginDTO> = {
      status: "success",
      data: { AccessToken: token, RefreshToken: token },
    };
    this.storage.setItem(this.authKey, JSON.stringify(data.data));
  }

  getAccessToken(): string | undefined {
    const data = this.storage.getItem(this.authKey);
    if (!data) return;
    const auth = JSON.parse(data) as { AccessToken: string };
    return auth.AccessToken;
  }

  isLoggedIn(): boolean {
    return !!this.storage.getItem(this.authKey);
  }
}
