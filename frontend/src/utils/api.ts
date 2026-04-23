const API_BASE_URL = import.meta.env.VITE_API_URL ?? "http://localhost:8080";

// ── Types ─────────────────────────────────────────────────────────────────────

export interface Review {
  id: number;
  name: string;
  rating: number;
  body: string;
  pros: string[];
  cons: string[];
  doItAgain: boolean;
  approved: boolean;
  date: string; // "February 11, 2026" — formatted by the server
}

export interface ReviewStats {
  totalReviews: number;
  averageRating: number;
  ratingCounts: Record<number, number>; // e.g. {1: 10, 2: 5, 3: 20, 4: 50, 5: 100}
}

export type CreateReviewPayload = Pick<Review, "name" | "rating"> &
  Partial<Pick<Review, "body" | "pros" | "cons" | "doItAgain">>;

interface ApiError {
  error: string;
}

// ── Helpers ───────────────────────────────────────────────────────────────────

async function parseResponse<T>(res: Response): Promise<T> {
  const data = await res.json();
  if (!res.ok) {
    const err = data as ApiError;
    throw new Error(err.error ?? `Request failed (${res.status})`);
  }
  return data as T;
}

function adminHeaders(adminKey: string) {
  return {
    "Content-Type": "application/json",
    "X-Admin-Key": adminKey,
  };
}

// ── Public endpoints ──────────────────────────────────────────────────────────

export async function fetchReviews(
  limit: number = 10,
  offset: number = 0,
): Promise<Review[]> {
  const res = await fetch(
    `${API_BASE_URL}/reviews?limit=${limit}&offset=${offset}`,
  );
  return parseResponse<Review[]>(res);
}

export async function createReview(
  payload: CreateReviewPayload,
): Promise<Review> {
  const res = await fetch(`${API_BASE_URL}/reviews`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
  return parseResponse<Review>(res);
}

export async function fetchStats(): Promise<ReviewStats> {
  const res = await fetch(`${API_BASE_URL}/stats`);
  return parseResponse<ReviewStats>(res);
}

// ── Admin endpoints ───────────────────────────────────────────────────────────

export async function fetchAllReviews(
  adminKey: string,
  limit: number = 20,
  offset: number = 0,
): Promise<Review[]> {
  const res = await fetch(
    `${API_BASE_URL}/admin/reviews?limit=${limit}&offset=${offset}`,
    { headers: adminHeaders(adminKey) },
  );
  return parseResponse<Review[]>(res);
}

export async function approveReview(
  adminKey: string,
  id: number,
): Promise<Review> {
  const res = await fetch(`${API_BASE_URL}/approve/${id}`, {
    method: "PUT",
    headers: adminHeaders(adminKey),
  });
  return parseResponse<Review>(res);
}

export async function deleteReview(
  adminKey: string,
  id: number,
): Promise<void> {
  const res = await fetch(`${API_BASE_URL}/reviews/${id}`, {
    method: "DELETE",
    headers: adminHeaders(adminKey),
  });
  if (!res.ok) {
    const data = await res.json().catch(() => ({}));
    throw new Error(
      (data as ApiError).error ?? `Delete failed (${res.status})`,
    );
  }
}
