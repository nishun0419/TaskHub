export const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL;

export const API_ENDPOINTS = {
  REGISTER: `${API_BASE_URL}/register`,
  LOGIN: `${API_BASE_URL}/login`,
  TEAMS: `${API_BASE_URL}/team`,
  TEAM: (id: number) => `${API_BASE_URL}/team/${id}`,
  TEAM_TODO: (teamID: number) => `${API_BASE_URL}/todo/team/${teamID}`,
  TEAM_INVITE: (id: number) => `${API_BASE_URL}/team/${id}/invite`,
  // 他のエンドポイントもここに追加できます
} as const; 