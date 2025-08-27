'use client';

import { useEffect, useState, use, useRef } from 'react';
import { useRouter } from 'next/navigation';
import { API_ENDPOINTS } from '@/constants/api';

interface Todo {
  todo_id: number;
  title: string;
  description?: string;
  completed: boolean;
  due_date?: string;
  assigned_to?: string;
  customer_id: number;
}

interface Team {
  team_id: number;
  name: string;
  description?: string;
  role: string;
}

interface TodoFormData {
  title: string;
  description: string;
  completed?: boolean;
}

interface InviteResponse {
  data: string;
  message: string;
}

export default function TeamDetailPage({ params }: { params: Promise<{ teamId: number }> }) {
  const resolvedParams = use(params);
  const router = useRouter();
  const [team, setTeam] = useState<Team | null>(null);
  const [todos, setTodos] = useState<Todo[]>([]);
  const [error, setError] = useState('');
  const [isAddModalOpen, setIsAddModalOpen] = useState(false);
  const [isEditModalOpen, setIsEditModalOpen] = useState(false);
  const [isInviteModalOpen, setIsInviteModalOpen] = useState(false);
  const [inviteEmail, setInviteEmail] = useState('');
  const [inviteToken, setInviteToken] = useState<string | null>(null);
  const [selectedTodo, setSelectedTodo] = useState<Todo | null>(null);
  const [formData, setFormData] = useState<TodoFormData>({
    title: '',
    description: '',
  });
  const hasFetched = useRef(false);

  useEffect(() => {
    if (hasFetched.current) return;
    hasFetched.current = true;
    const fetchTeamAndTodos = async () => {
      try {
        const token = localStorage.getItem('token');
        if (!token) {
          router.push('/login');
          return;
        }

        // チーム情報の取得
        const teamResponse = await fetch(`${API_ENDPOINTS.TEAM(resolvedParams.teamId)}`, {
          headers: {
            'Authorization': `Bearer ${token}`,
          },
        });

        if (!teamResponse.ok) {
          throw new Error('チーム情報の取得に失敗しました');
        }

        const teamData = await teamResponse.json();
        if (!teamData.data.name) {
          router.push('/teams');
          return;
        }
        setTeam(teamData.data);

        // TODOの取得
        const todosResponse = await fetch(`${API_ENDPOINTS.TEAM_TODO(resolvedParams.teamId)}`, {
          headers: {
            'Authorization': `Bearer ${token}`,
          },
        });

        if (!todosResponse.ok) {
          throw new Error('TODOの取得に失敗しました');
        }

        const todosData = await todosResponse.json();
        setTodos(todosData.data);
      } catch (err) {
        setError(err instanceof Error ? err.message : 'データの取得に失敗しました');
      }
    };

    fetchTeamAndTodos();
  }, [resolvedParams.teamId]);

  const handleAddTodo = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const token = localStorage.getItem('token');
      if (!token) {
        router.push('/login');
        return;
      }

      const response = await fetch(`${API_ENDPOINTS.TEAM_TODO(resolvedParams.teamId)}`, {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          ...formData,
          completed: false,
        }),
      });

      if (!response.ok) {
        throw new Error('TODOの追加に失敗しました');
      }

      const newTodo = await response.json();
      setTodos([...todos, newTodo.data]);
      setIsAddModalOpen(false);
      setFormData({ title: '', description: '' });
    } catch (err) {
      setError(err instanceof Error ? err.message : 'TODOの追加に失敗しました');
    }
  };

  const handleEditTodo = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!selectedTodo) return;

    try {
      const token = localStorage.getItem('token');
      if (!token) {
        router.push('/login');
        return;
      }
      const response = await fetch(`${API_ENDPOINTS.TODO(selectedTodo.todo_id)}`, {
        method: 'PUT',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          title: formData.title,
          description: formData.description,
          completed: Boolean(formData.completed)
        }),
      });
      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message);
      }

      setTodos(todos.map(todo => 
        todo.todo_id === selectedTodo.todo_id ? {
          ...todo,
          title: formData.title,
          description: formData.description,
          completed: formData.completed ?? todo.completed
        } : todo
      ));
      setIsEditModalOpen(false);
      setSelectedTodo(null);
      setFormData({ title: '', description: '' });
    } catch (err) {
      setError(err instanceof Error ? err.message : 'TODOの更新に失敗しました');
    }
  };

  const handleStatusChange = async (todoId: number, newStatus: boolean) => {
    try {
      const token = localStorage.getItem('token');
      if (!token) {
        router.push('/login');
        return;
      }


      const response = await fetch(`${API_ENDPOINTS.TODO_STATUS_CHANGE(todoId)}`, {
        method: 'PUT',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ completed: newStatus}),
      });

      if (!response.ok) {
        throw new Error('ステータスの更新に失敗しました');
      }

      setTodos(todos.map(todo => 
        todo.todo_id === todoId ? {
          ...todo,
          completed: newStatus
        } : todo
      ));
    } catch (err) {
      setError(err instanceof Error ? err.message : 'ステータスの更新に失敗しました');
    }
  };

  const handleDeleteTodo = async (todoId: number) => {
    if (!confirm('本当に削除してもよろしいですか？')) {
      return;
    }

    try {
      const token = localStorage.getItem('token');
      if (!token) {
        router.push('/login');
        return;
      }

      const response = await fetch(`${API_ENDPOINTS.TODO(todoId)}`, {
        method: 'DELETE',
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      });

      if (!response.ok) {
        throw new Error('TODOの削除に失敗しました');
      }

      setTodos(todos.filter(todo => todo.todo_id !== todoId));
    } catch (err) {
      setError(err instanceof Error ? err.message : 'TODOの削除に失敗しました');
    }
  };

  const openEditModal = (todo: Todo) => {
    setSelectedTodo(todo);
    setFormData({
      title: todo.title,
      description: todo.description || '',
      completed: todo.completed,
    });
    setIsEditModalOpen(true);
  };

  if (error) {
    return (
      <div className="min-h-screen bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-3xl mx-auto">
          <div className="bg-white shadow overflow-hidden sm:rounded-lg p-6">
            <div className="text-red-500">{error}</div>
          </div>
        </div>
      </div>
    );
  }

  if (!team) {
    return (
      <div className="min-h-screen bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
        <div className="max-w-3xl mx-auto">
          <div className="bg-white shadow overflow-hidden sm:rounded-lg p-6">
            <div className="text-gray-500">読み込み中...</div>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-3xl mx-auto">
        <div className="bg-white shadow overflow-hidden sm:rounded-lg">
          <div className="px-4 py-5 sm:px-6">
            <h3 className="text-lg leading-6 font-medium text-gray-900">
              {team.name}
            </h3>
            {team.description && (
              <p className="mt-1 max-w-2xl text-sm text-gray-500">
                {team.description}
              </p>
            )}
            <div className="mt-2 flex justify-between items-center">
              <span className="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-green-100 text-green-800">
                {team.role}
              </span>
              {team.role === 'owner' && (
                <button
                  onClick={() => setIsInviteModalOpen(true)}
                  className="bg-green-500 hover:bg-green-600 text-white px-4 py-2 rounded-md text-sm font-medium"
                >
                  チームに招待
                </button>
              )}
            </div>
          </div>

          <div className="border-t border-gray-200">
            <div className="px-4 py-5 sm:px-6">
              <div className="flex justify-between items-center mb-4">
                <h4 className="text-lg font-medium text-gray-900">TODO一覧</h4>
                <button
                  onClick={() => setIsAddModalOpen(true)}
                  className="bg-blue-500 hover:bg-blue-600 text-white px-4 py-2 rounded-md text-sm font-medium"
                >
                  新しいTODOを追加
                </button>
              </div>
              {todos.length === 0 ? (
                <p className="text-gray-500">TODOはありません</p>
              ) : (
                <ul className="divide-y divide-gray-200">
                  {todos.map((todo) => (
                    <li key={todo.todo_id} className="py-4">
                      <div className="flex items-center justify-between">
                        <div className="flex-1 cursor-pointer" onClick={() => openEditModal(todo)}>
                          <p className="text-sm font-medium text-gray-900">{todo.title}</p>
                          {todo.description && (
                            <p className="text-sm text-gray-500">{todo.description}</p>
                          )}
                        </div>
                        <div className="flex items-center space-x-4">
                          <select
                            value={todo.completed ? "1" : "0"}
                            onChange={(e) => handleStatusChange(todo.todo_id, e.target.value === "1")}
                            className={`px-2 py-1 text-xs font-semibold rounded-full ${
                              todo.completed ? 'bg-green-100 text-green-800' : 'bg-gray-100 text-gray-800'
                            }`}
                          >
                            <option value="0">未完了</option>
                            <option value="1">完了</option>
                          </select>
                          {todo.customer_id === JSON.parse(localStorage.getItem('user') || '{}').customer_id && (
                            <button
                              onClick={() => handleDeleteTodo(todo.todo_id)}
                              className="px-2 py-1 text-xs font-semibold rounded-full bg-red-100 text-red-800 hover:bg-red-200"
                            >
                              削除
                            </button>
                          )}
                          {todo.due_date && (
                            <span className="text-sm text-gray-500">
                              期限: {new Date(todo.due_date).toLocaleDateString()}
                            </span>
                          )}
                        </div>
                      </div>
                    </li>
                  ))}
                </ul>
              )}
            </div>
          </div>
        </div>
      </div>

      {/* Add Todo Modal */}
      {isAddModalOpen && (
        <div className="fixed inset-0 bg-gray-500 bg-opacity-75 flex items-center justify-center">
          <div className="bg-white rounded-lg p-6 max-w-md w-full">
            <h3 className="text-lg font-medium mb-4">新しいTODOを追加</h3>
            <form onSubmit={handleAddTodo}>
              <div className="mb-4">
                <label className="block text-sm font-medium text-gray-700">タイトル</label>
                <input
                  type="text"
                  value={formData.title}
                  onChange={(e) => setFormData({ ...formData, title: e.target.value })}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                  required
                />
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium text-gray-700">説明</label>
                <textarea
                  value={formData.description}
                  onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                  rows={3}
                />
              </div>
              <div className="flex justify-end space-x-3">
                <button
                  type="button"
                  onClick={() => {
                    setIsAddModalOpen(false);
                    setFormData({ title: '', description: '' });
                  }}
                  className="px-4 py-2 text-sm font-medium text-gray-700 bg-gray-100 hover:bg-gray-200 rounded-md"
                >
                  キャンセル
                </button>
                <button
                  type="submit"
                  className="px-4 py-2 text-sm font-medium text-white bg-blue-500 hover:bg-blue-600 rounded-md"
                >
                  追加
                </button>
              </div>
            </form>
          </div>
        </div>
      )}

      {/* Edit Todo Modal */}
      {isEditModalOpen && selectedTodo && (
        <div className="fixed inset-0 bg-gray-500 bg-opacity-75 flex items-center justify-center">
          <div className="bg-white rounded-lg p-6 max-w-md w-full">
            <h3 className="text-lg font-medium mb-4">TODOを編集</h3>
            <form onSubmit={handleEditTodo}>
              <div className="mb-4">
                <label className="block text-sm font-medium text-gray-700">タイトル</label>
                <input
                  type="text"
                  value={formData.title}
                  onChange={(e) => setFormData({ ...formData, title: e.target.value })}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                  required
                />
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium text-gray-700">説明</label>
                <textarea
                  value={formData.description}
                  onChange={(e) => setFormData({ ...formData, description: e.target.value })}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                  rows={3}
                />
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium text-gray-700">ステータス</label>
                <select
                  value={formData.completed ? "1" : "0"}
                  onChange={(e) => setFormData({ ...formData, completed: e.target.value === "1" })}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                >
                  <option value="0">未完了</option>
                  <option value="1">完了</option>
                </select>
              </div>
              <div className="flex justify-end space-x-3">
                <button
                  type="button"
                  onClick={() => {
                    setIsEditModalOpen(false);
                    setSelectedTodo(null);
                    setFormData({ title: '', description: '' });
                  }}
                  className="px-4 py-2 text-sm font-medium text-gray-700 bg-gray-100 hover:bg-gray-200 rounded-md"
                >
                  キャンセル
                </button>
                <button
                  type="submit"
                  className="px-4 py-2 text-sm font-medium text-white bg-blue-500 hover:bg-blue-600 rounded-md"
                >
                  更新
                </button>
              </div>
            </form>
          </div>
        </div>
      )}

      {/* Invite Modal */}
      {isInviteModalOpen && (
        <div className="fixed inset-0 bg-gray-500 bg-opacity-75 flex items-center justify-center">
          <div className="bg-white rounded-lg p-6 max-w-md w-full">
            <h3 className="text-lg font-medium mb-4">チームに招待</h3>
            {!inviteToken ? (
              <form onSubmit={async (e) => {
                e.preventDefault();
                try {
                  const token = localStorage.getItem('token');
                  if (!token) {
                    router.push('/login');
                    return;
                  }

                  const response = await fetch(API_ENDPOINTS.TEAM_INVITE(resolvedParams.teamId), {
                    method: 'POST',
                    headers: {
                      'Authorization': `Bearer ${token}`,
                      'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ mail: inviteEmail }),
                  });

                  if (!response.ok) {
                    throw new Error('招待の送信に失敗しました');
                  }
                  const data: InviteResponse = await response.json();
                  setInviteToken(data.data);
                } catch (err) {
                  setError(err instanceof Error ? err.message : '招待の送信に失敗しました');
                }
              }}>
                <div className="mb-4">
                  <label className="block text-sm font-medium text-gray-700">メールアドレス</label>
                  <input
                    type="email"
                    value={inviteEmail}
                    onChange={(e) => setInviteEmail(e.target.value)}
                    className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                    required
                  />
                </div>
                <div className="flex justify-end space-x-3">
                  <button
                    type="button"
                    onClick={() => {
                      setIsInviteModalOpen(false);
                      setInviteEmail('');
                      setInviteToken(null);
                    }}
                    className="px-4 py-2 text-sm font-medium text-gray-700 bg-gray-100 hover:bg-gray-200 rounded-md"
                  >
                    キャンセル
                  </button>
                  <button
                    type="submit"
                    className="px-4 py-2 text-sm font-medium text-white bg-green-500 hover:bg-green-600 rounded-md"
                  >
                    招待を送信
                  </button>
                </div>
              </form>
            ) : (
              <div>
                <p className="mb-4 text-sm text-gray-600">以下の招待URLを共有してください：</p>
                <div className="bg-gray-50 p-3 rounded-md break-all">
                  <code className="text-sm">
                    {`${window.location.origin}/join?token=${inviteToken}`}
                  </code>
                </div>
                <div className="mt-4 flex justify-end">
                  <button
                    onClick={() => {
                      setIsInviteModalOpen(false);
                      setInviteEmail('');
                      setInviteToken(null);
                    }}
                    className="px-4 py-2 text-sm font-medium text-gray-700 bg-gray-100 hover:bg-gray-200 rounded-md"
                  >
                    閉じる
                  </button>
                </div>
              </div>
            )}
          </div>
        </div>
      )}
    </div>
  );
} 