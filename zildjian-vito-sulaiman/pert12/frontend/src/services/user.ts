import api from "./api";
import type { User } from "../types/user";

export async function getMe(): Promise<User> {
  const res = await api.get<User>("/me");
  return res.data;
}

export const updateProfile = async (id: number, formData: FormData) => {
  const { data } = await api.patch(`/profile/${id}`, formData, {
    headers: {
      "Content-Type": "multipart/form-data",
    },
  });
  return data;
};
