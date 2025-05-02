'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import { useSession } from 'next-auth/react';
import LoginForm from '../components/form/LoginForm';
export default function LoginPage() {
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

  return (
    <LoginForm />
  );
} 