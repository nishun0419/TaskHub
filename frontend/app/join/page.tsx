'use client';

import { useEffect, useState, Suspense, useRef } from 'react';
import { useRouter, useSearchParams } from 'next/navigation';
import { API_ENDPOINTS } from '@/constants/api';

interface JoinResponse {
  data: number;
  message: string;
}

function JoinContent() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [error, setError] = useState<string>('');
  const [isProcessing, setIsProcessing] = useState(true);
  const hasProcessed = useRef(false);
  const token = searchParams.get('token');

  useEffect(() => {
    if (!token || hasProcessed.current) return;
  
    const processInvitation = async () => {
      if (hasProcessed.current) return;
      hasProcessed.current = true;
  
      const storedToken = localStorage.getItem('token');
      if (!storedToken) {
        localStorage.setItem('inviteRedirectUrl', window.location.href);
        router.replace('/login');
        return;
      }
  
      try {
        const response = await fetch(`${API_ENDPOINTS.TEAM_JOIN}`, {
          method: 'POST',
          headers: {
            Authorization: `Bearer ${storedToken}`,
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ token }),
        });
  
        if (!response.ok) {
          const errorData = await response.json();
          throw new Error(errorData.message || '招待の処理に失敗しました');
        }
  
        const data: JoinResponse = await response.json();
        router.replace(`/teams/${data.data}`);
      } catch (err) {
        setError(err instanceof Error ? err.message : '招待の処理中にエラーが発生しました');
        setIsProcessing(false);
      }
    };
  
    processInvitation();
  }, [token]);
  

  if (error) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-md w-full space-y-8">
          <div className="text-center">
            <h2 className="mt-6 text-3xl font-extrabold text-gray-900">
              エラー
            </h2>
            <p className="mt-2 text-sm text-red-600">
              {error}
            </p>
            <button
              onClick={() => router.push('/teams')}
              className="mt-4 inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md shadow-sm text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500"
            >
              チーム一覧に戻る
            </button>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md w-full space-y-8">
        <div className="text-center">
          <h2 className="mt-6 text-3xl font-extrabold text-gray-900">
            チームに参加中
          </h2>
          <p className="mt-2 text-sm text-gray-600">
            しばらくお待ちください...
          </p>
        </div>
      </div>
    </div>
  );
}

export default function JoinPage() {
  return (
    <Suspense fallback={
      <div className="min-h-screen bg-gray-50 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-md w-full space-y-8">
          <div className="text-center">
            <h2 className="mt-6 text-3xl font-extrabold text-gray-900">
              読み込み中
            </h2>
            <p className="mt-2 text-sm text-gray-600">
              しばらくお待ちください...
            </p>
          </div>
        </div>
      </div>
    }>
      <JoinContent />
    </Suspense>
  );
} 