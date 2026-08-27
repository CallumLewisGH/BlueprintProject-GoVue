import type { ErrorModel } from "@/api";

export type ApiResult<T> = { data: T | null; error: ErrorModel | null; status: number };