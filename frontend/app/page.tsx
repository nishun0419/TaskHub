'use client';

import { useSession } from 'next-auth/react';
import { useRouter } from 'next/navigation';
import { useEffect, useState } from 'react';
import Link from 'next/link';

export default function Home() {
  const { status } = useSession();
  const router = useRouter();
  const [storedUser, setStoredUser] = useState<string | null>(null);

  useEffect(() => {
    const user = localStorage.getItem('user');
    setStoredUser(user);
  }, []);

  useEffect(() => {
    if (status === 'authenticated' || storedUser) {
      router.push('/mypage');
    }
  }, [status, storedUser, router]);

  if (status === 'loading') {
    return (
      <div className="min-h-screen flex items-center justify-center">
        <div className="animate-spin rounded-full h-32 w-32 border-t-2 border-b-2 border-purple-500"></div>
      </div>
    );
  }

  return (
    <div className="bg-white">
      {/* Hero section */}
      <div className="relative isolate px-6 pt-14 lg:px-8">
        <div className="mx-auto max-w-2xl py-32 sm:py-48 lg:py-56">
          <div className="text-center">
            <h1 className="text-4xl font-bold tracking-tight text-gray-900 sm:text-6xl">
              GoFlow
              <span className="text-indigo-600"> - シンプルな会員管理機能を一瞬で作成</span>
            </h1>
            <p className="mt-6 text-lg leading-8 text-gray-600">
              Go（Gin）とNext.js 14を使用した、JWT認証・Google認証対応の
              シンプルな会員管理機能を一瞬で作成できます。
            </p>
            <div className="mt-10 flex items-center justify-center gap-x-6">
              <Link
                href="/register"
                className="rounded-md bg-indigo-600 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
              >
                会員登録
              </Link>
              <Link href="/login" className="text-sm font-semibold leading-6 text-gray-900">
                ログイン <span aria-hidden="true">→</span>
              </Link>
            </div>
          </div>
        </div>
      </div>

      {/* Feature section */}
      <div className="bg-gray-50 py-24 sm:py-32">
        <div className="mx-auto max-w-7xl px-6 lg:px-8">
          <div className="mx-auto max-w-2xl lg:text-center">
            <h2 className="text-base font-semibold leading-7 text-indigo-600">主要機能</h2>
            <p className="mt-2 text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">
              シンプルで使いやすい会員管理
            </p>
          </div>
          <div className="mx-auto mt-16 max-w-2xl sm:mt-20 lg:mt-24 lg:max-w-none">
            <dl className="grid max-w-xl grid-cols-1 gap-x-8 gap-y-16 lg:max-w-none lg:grid-cols-3">
              {features.map((feature) => (
                <div key={feature.name} className="flex flex-col">
                  <dt className="flex items-center gap-x-3 text-base font-semibold leading-7 text-gray-900">
                    {feature.name}
                  </dt>
                  <dd className="mt-4 flex flex-auto flex-col text-base leading-7 text-gray-600">
                    <p className="flex-auto">{feature.description}</p>
                  </dd>
                </div>
              ))}
            </dl>
          </div>
        </div>
      </div>

      {/* Tech stack section */}
      <div className="py-24 sm:py-32">
        <div className="mx-auto max-w-7xl px-6 lg:px-8">
          <div className="mx-auto max-w-2xl lg:text-center">
            <h2 className="text-base font-semibold leading-7 text-indigo-600">使用技術</h2>
            <p className="mt-2 text-3xl font-bold tracking-tight text-gray-900 sm:text-4xl">
              最新の技術スタック
            </p>
          </div>
          <div className="mx-auto mt-16 max-w-2xl sm:mt-20 lg:mt-24 lg:max-w-none">
            <dl className="grid max-w-xl grid-cols-1 gap-x-8 gap-y-16 lg:max-w-none lg:grid-cols-2">
              {techStacks.map((stack) => (
                <div key={stack.category} className="flex flex-col">
                  <dt className="text-base font-semibold leading-7 text-gray-900">
                    {stack.category}
                  </dt>
                  <dd className="mt-4 flex flex-auto flex-col text-base leading-7 text-gray-600">
                    <ul className="list-disc pl-5 space-y-2">
                      {stack.items.map((item) => (
                        <li key={item}>{item}</li>
                      ))}
                    </ul>
                  </dd>
                </div>
              ))}
            </dl>
          </div>
        </div>
      </div>
    </div>
  );
}

const features = [
  {
    name: 'ユーザー登録・認証',
    description: 'メールアドレス・パスワードによる登録と、GoogleアカウントによるOAuthログインをサポート。',
  },
  {
    name: 'JWT認証',
    description: 'セキュアなJWTトークンによる認証システムを実装。安全なユーザー管理を実現。',
  },
  {
    name: 'マイページ機能',
    description: '認証後のユーザー専用マイページが表示されます。',
  },
];

const techStacks = [
  {
    category: 'バックエンド',
    items: [
      'Go 1.23',
      'Gin',
      'GORM',
      'Goose（マイグレーションツール）',
      'MySQL',
    ],
  },
  {
    category: 'フロントエンド',
    items: [
      'Next.js 14 (App Router)',
      'Tailwind CSS',
      'NextAuth.js（Google OAuth認証）',
      'JWT認証（アクセストークン）',
    ],
  },
];
