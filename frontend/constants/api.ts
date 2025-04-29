export const API_BASE_URL = 'http://localhost:8080';

export const API_ENDPOINTS = {
  REGISTER: 'http://localhost:8080/api/register',
  LOGIN: 'http://localhost:8080/api/login',
  TEAMS: 'http://localhost:8080/api/teams',
  TEAM: (id: number) => `http://localhost:8080/api/teams/${id}`,
  TEAM_INVITE: (id: number) => `http://localhost:8080/api/teams/${id}/invite`,
  // 他のエンドポイントもここに追加できます
} as const; 