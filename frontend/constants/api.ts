export const API_BASE_URL = 'http://localhost:8080';

export const API_ENDPOINTS = {
  REGISTER: `${API_BASE_URL}/api/register`,
  LOGIN: `${API_BASE_URL}/api/login`,
  // 他のエンドポイントもここに追加できます
} as const; 