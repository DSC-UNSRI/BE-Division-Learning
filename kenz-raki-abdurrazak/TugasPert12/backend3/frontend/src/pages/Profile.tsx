import { useState, useEffect } from "react";
import { useUpdateProfile, useUser } from "../hooks/user";

export default function ProfilePage() {
  const { user, userLoading, setUser } = useUser();
  const { handleUpdateProfile, loading, success } = useUpdateProfile();

  const [name, setName] = useState("");
  const [password, setPassword] = useState("");
  const [profilePicture, setProfilePicture] = useState<File | null>(null);

  useEffect(() => {
    if (user) {
      setName(user.name || "");
    }
  }, [user]);

  if (userLoading) return <p className="text-center">Loading...</p>;

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      if (!user?.id) throw new Error("User ID is missing");
      const updated = await handleUpdateProfile(user.id, {
        name,
        password: password || undefined,
        profile_picture: profilePicture,
      });
      setUser(updated);
    } catch {
      //
    }
  };

  return (
    <div className="max-w-md mx-auto p-6 bg-white shadow-lg rounded-xl">
      <h2 className="text-2xl font-semibold mb-4 text-center">
        Update Profile
      </h2>
      <form onSubmit={handleSubmit} className="space-y-4">
        <div>
          <label className="block text-sm font-medium">Nama</label>
          <input
            type="text"
            className="w-full mt-1 border rounded-md p-2"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </div>

        <div>
          <label className="block text-sm font-medium">Password Baru</label>
          <input
            type="password"
            className="w-full mt-1 border rounded-md p-2"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Kosongkan jika tidak diganti"
          />
        </div>

        <div>
          <label className="block text-sm font-medium">Foto Profil</label>
          <input
            type="file"
            className="w-full mt-1"
            accept="image/*"
            onChange={(e) =>
              setProfilePicture(e.target.files ? e.target.files[0] : null)
            }
          />
        </div>

        <button
          type="submit"
          disabled={loading}
          className="w-full bg-blue-600 text-white py-2 rounded-md hover:bg-blue-700 disabled:opacity-50"
        >
          {loading ? "Menyimpan..." : "Simpan"}
        </button>
      </form>

      {success && (
        <p className="mt-4 text-green-600 text-center">{success}</p>
      )}
    </div>
  );
}
