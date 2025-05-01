'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { API_ENDPOINTS } from '@/constants/api';

interface Team {
  TeamID: number;
  Name: string;
  Description?: string;
  CreatedAt: string;
  UpdatedAt: string;
}

export default function TeamsPage() {
  const router = useRouter();
  const [teams, setTeams] = useState<Team[]>([]);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchTeams = async () => {
      try {
        const token = localStorage.getItem('token');
        if (!token) {
          router.push('/login');
          return;
        }

        const response = await fetch(API_ENDPOINTS.TEAMS, {
          headers: {
            'Authorization': `Bearer ${token}`,
          },
        });

        if (!response.ok) {
          throw new Error('チームの取得に失敗しました');
        }
        const data = await response.json();
        setTeams(data.data);
      } catch (err) {
        setError(err instanceof Error ? err.message : 'チームの取得に失敗しました');
      }
    };

    fetchTeams();
  }, [router]);

  return (
    <div className="min-h-screen bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-3xl mx-auto">
        <div className="bg-white shadow overflow-hidden sm:rounded-lg">
          <div className="px-4 py-5 sm:px-6 flex justify-between items-center">
            <h3 className="text-lg leading-6 font-medium text-gray-900">
              チーム一覧
            </h3>
            <a
              href="/teams/create"
              className="inline-flex items-center px-4 py-2 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
            >
              チーム作成
            </a>
          </div>
          <div className="border-t border-gray-200">
            {error ? (
              <div className="px-4 py-5 sm:px-6 text-red-500">
                {error}
              </div>
            ) : teams.length === 0 ? (
              <div className="px-4 py-5 sm:px-6 text-gray-500">
                所属しているチームはありません
              </div>
            ) : (
              <ul className="divide-y divide-gray-200">
                {teams.map((team) => (
                  <li key={team.TeamID}>
                    <a
                      href={`/teams/${team.TeamID}`}
                      className="block hover:bg-gray-50"
                    >
                      <div className="px-4 py-4 sm:px-6">
                        <div className="flex items-center justify-between">
                          <p className="text-sm font-medium text-indigo-600 truncate">
                            {team.Name}
                          </p>
                          <div className="ml-2 flex-shrink-0 flex">
                            <p className="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">
                              メンバー
                            </p>
                          </div>
                        </div>
                        {team.Description && (
                          <div className="mt-2 sm:flex sm:justify-between">
                            <div className="sm:flex">
                              <p className="flex items-center text-sm text-gray-500">
                                {team.Description}
                              </p>
                            </div>
                          </div>
                        )}
                      </div>
                    </a>
                  </li>
                ))}
              </ul>
            )}
          </div>
        </div>
      </div>
    </div>
  );
} 