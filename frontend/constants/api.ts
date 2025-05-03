export const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL;

export const API_ENDPOINTS = {
  REGISTER: 'http://localhost:8080/api/register',
  LOGIN: 'http://localhost:8080/api/login',
  TEAMS: 'http://localhost:8080/api/team',
  TEAM: (id: number) => `http://localhost:8080/api/team/${id}`,
  TEAM_TODO: (teamID: number) => `http://localhost:8080/api/todo/team/${teamID}`,
  TEAM_INVITE: (id: number) => `http://localhost:8080/api/team/${id}/invite`,
  // 他のエンドポイントもここに追加できます
} as const; 