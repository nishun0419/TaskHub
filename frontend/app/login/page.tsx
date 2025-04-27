'use client';

import { useState } from 'react';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import { API_ENDPOINTS } from '@/constants/api';
import LoginForm from '../components/form/LoginForm';
export default function LoginPage() {
  

  return (
    <LoginForm />
  );
} 