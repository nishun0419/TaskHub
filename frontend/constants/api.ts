export const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

export const API_ENDPOINTS = {
  REGISTER: `${API_BASE_URL}/api/register`,
  LOGIN: `${API_BASE_URL}/api/login`,
  TEAMS: `${API_BASE_URL}/api/team`,
  TEAM: (id: number) => `${API_BASE_URL}/api/team/${id}`,
  TEAM_TODO: (teamID: number) => `${API_BASE_URL}/api/todo/team/${teamID}`,
  TODO: (id: number) => `${API_BASE_URL}/api/todo/${id}`,
  TODO_STATUS_CHANGE: (id: number) => `${API_BASE_URL}/api/todo/${id}/status`,
  TEAM_INVITE: (id: number) => `${API_BASE_URL}/api/team/${id}/invite`,
  // 他のエンドポイントもここに追加できます
} as const; 