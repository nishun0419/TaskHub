'use client';
import RegisterForm from '../components/form/RegisterForm';
import { useSession } from 'next-auth/react';
import { useRouter } from 'next/navigation';
import { useEffect, useState } from 'react';

export default function Register() {
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

  return <RegisterForm />;
} 