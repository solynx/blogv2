interface ErrorResponse {
  status: boolean;
  code: number;
  message?: string;
}

interface UserTokenResponse {
  token: string;
  metadata?: any;
  status: boolean;
  code: number;
}
export type { ErrorResponse, UserTokenResponse };
