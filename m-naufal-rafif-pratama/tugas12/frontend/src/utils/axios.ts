import { AxiosError } from "axios";

export function isAxiosError(err: unknown): err is AxiosError<{ message: string }> {
  return typeof err === "object" && err !== null && "isAxiosError" in err;
}
