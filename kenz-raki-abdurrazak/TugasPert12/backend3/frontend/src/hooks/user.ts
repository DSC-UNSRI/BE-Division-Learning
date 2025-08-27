import { useState, useEffect } from "react";
import type { User } from "../types/user";
import { getMe, updateProfile } from "../services/user";

export function useUser() {
  const [user, setUser] = useState<User | null>(null);
  const [userLoading, setLoading] = useState(true);

  useEffect(() => {
    getMe()
      .then((data) => setUser(data))
      .finally(() => setLoading(false));
  }, []);

  return { user, userLoading, setUser };
}

export function useUpdateProfile() {
  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState<string | null>(null);

  const handleUpdateProfile = async (id: number, payload: {
    name?: string;
    password?: string;
    profile_picture?: File | null;
  }) => {
    setLoading(true);
    setSuccess(null);

    const formData = new FormData();
    if (payload.name) formData.append("name", payload.name);
    if (payload.password) formData.append("password", payload.password);
    if (payload.profile_picture) {
      formData.append("profile_picture", payload.profile_picture);
    }

    try {
      const res = await updateProfile(id, formData);
      setSuccess("Profile berhasil diperbarui.");
      return res;
    } finally {
      setLoading(false);
    }
  };

  return { handleUpdateProfile, loading, success };
}