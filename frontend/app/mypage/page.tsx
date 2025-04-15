'use client';

import { useSession, signOut } from "next-auth/react";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

export default function MyPage() {
  const { data: session, status } = useSession();
  const router = useRouter();

  useEffect(() => {
    if (status === "unauthenticated") {
      router.push("/");
    }
  }, [status, router]);

  if (status === "loading") {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="animate-spin rounded-full h-32 w-32 border-t-2 border-b-2 border-purple-500"></div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-purple-500 via-pink-500 to-red-500 p-4">
      <div className="max-w-4xl mx-auto bg-white rounded-lg shadow-xl p-8 mt-8">
        <div className="flex justify-between items-center mb-8">
          <h1 className="text-3xl font-bold text-gray-800">マイページ</h1>
          <button
            onClick={() => signOut()}
            className="bg-red-500 hover:bg-red-600 text-white px-4 py-2 rounded-md transition-colors duration-200"
          >
            ログアウト
          </button>
        </div>

        <div className="space-y-6">
          <div className="flex items-center space-x-4">
            {session?.user?.image && (
              <img
                src={session.user.image}
                alt="Profile"
                className="w-20 h-20 rounded-full"
              />
            )}
            <div>
              <h2 className="text-xl font-semibold text-gray-800">
                {session?.user?.name}
              </h2>
              <p className="text-gray-600">{session?.user?.email}</p>
            </div>
          </div>

          <div className="bg-gray-50 p-6 rounded-lg">
            <h3 className="text-lg font-semibold text-gray-800 mb-4">
              アカウント情報
            </h3>
            <div className="space-y-2">
              <p>
                <span className="font-medium">プロバイダー:</span>{" "}
                {session?.user?.provider}
              </p>
              <p>
                <span className="font-medium">最終ログイン:</span>{" "}
                {new Date().toLocaleString()}
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
} 