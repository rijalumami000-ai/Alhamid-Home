import type {
  HomeData,
  Article,
  Category,
  EducationUnit,
  Caregiver,
  Facility,
  Achievement,
  Testimonial,
  Alumni,
  MediaPhoto,
  MediaVideo,
  ContactMessage,
  SearchResult,
  User,
  Banner
} from './types';

const API_BASE = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8088/api';

// Helper for JWT Auth Token in browser
export function getToken(): string | null {
  if (typeof window !== 'undefined') {
    return localStorage.getItem('alhamid_admin_token');
  }
  return null;
}

export function setToken(token: string) {
  if (typeof window !== 'undefined') {
    localStorage.setItem('alhamid_admin_token', token);
  }
}

export function removeToken() {
  if (typeof window !== 'undefined') {
    localStorage.removeItem('alhamid_admin_token');
  }
}

async function request<T>(endpoint: string, options: RequestInit = {}): Promise<T> {
  const token = getToken();
  const headers = new Headers(options.headers || {});

  if (!headers.has('Content-Type') && !(options.body instanceof FormData)) {
    headers.set('Content-Type', 'application/json');
  }

  if (token && !headers.has('Authorization')) {
    headers.set('Authorization', `Bearer ${token}`);
  }

  const response = await fetch(`${API_BASE}${endpoint}`, {
    ...options,
    headers,
  });

  if (!response.ok) {
    const errorBody = await response.json().catch(() => ({}));
    throw new Error(errorBody.error || `Request failed with status ${response.status}`);
  }

  return response.json();
}

export const api = {
  // Public
  getHome: () => request<HomeData>('/public/home'),
  getSettings: () => request<Record<string, string>>('/public/settings'),
  getAbout: () => request<{
    profile: string;
    vision: string;
    mission: string;
    site_name: string;
    site_tagline: string;
    caregivers: Caregiver[];
    facilities: Facility[];
    achievements: Achievement[];
    testimonials: Testimonial[];
    alumni: Alumni[];
  }>('/public/about'),
  getCaregivers: () => request<Caregiver[]>('/public/caregivers'),
  getFacilities: () => request<Facility[]>('/public/facilities'),
  getAchievements: () => request<Achievement[]>('/public/achievements'),
  getTestimonials: () => request<Testimonial[]>('/public/testimonials'),
  getAlumni: () => request<Alumni[]>('/public/alumni'),
  getEducationUnits: (type?: string) =>
    request<EducationUnit[]>(`/public/education-units${type ? `?type=${type}` : ''}`),
  getEducationUnitBySlug: (slug: string) =>
    request<EducationUnit>(`/public/education-units/${slug}`),
  getCategories: (type?: string) =>
    request<Category[]>(`/public/categories${type ? `?type=${type}` : ''}`),
  getArticles: (params?: { category?: string; q?: string; page?: number; limit?: number }) => {
    const q = new URLSearchParams();
    if (params?.category) q.set('category', params.category);
    if (params?.q) q.set('q', params.q);
    if (params?.page) q.set('page', params.page.toString());
    if (params?.limit) q.set('limit', params.limit.toString());
    return request<{ data: Article[]; total: number; page: number; limit: number; pages: number }>(
      `/public/articles?${q.toString()}`
    );
  },
  getArticleBySlug: (slug: string) =>
    request<{ article: Article; related: Article[] }>(`/public/articles/${slug}`),
  getPhotos: (category?: string) =>
    request<MediaPhoto[]>(`/public/gallery/photos${category ? `?category=${category}` : ''}`),
  getVideos: (category?: string) =>
    request<MediaVideo[]>(`/public/gallery/videos${category ? `?category=${category}` : ''}`),
  submitContact: (data: { name: string; email: string; phone?: string; subject: string; message: string }) =>
    request<{ message: string }>('/public/contact', {
      method: 'POST',
      body: JSON.stringify(data),
    }),
  search: (query: string) =>
    request<SearchResult>(`/public/search?q=${encodeURIComponent(query)}`),

  // Auth
  login: (credentials: { username: string; password: string }) =>
    request<{ message: string; token: string; user: User }>('/auth/login', {
      method: 'POST',
      body: JSON.stringify(credentials),
    }),
  me: () => request<User>('/auth/me'),
  logout: () => request<{ message: string }>('/auth/logout', { method: 'POST' }),

  // Admin CMS
  admin: {
    getDashboard: () =>
      request<{
        stats: {
          total_articles: number;
          total_units: number;
          total_photos: number;
          total_videos: number;
          unread_messages: number;
          total_caregivers: number;
        };
        recent_articles: Article[];
        recent_messages: ContactMessage[];
      }>('/admin/dashboard'),

    // Articles
    getArticles: () => request<Article[]>('/admin/articles'),
    createArticle: (data: Partial<Article>) =>
      request<Article>('/admin/articles', { method: 'POST', body: JSON.stringify(data) }),
    updateArticle: (id: number, data: Partial<Article>) =>
      request<Article>(`/admin/articles/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
    deleteArticle: (id: number) =>
      request<{ message: string }>(`/admin/articles/${id}`, { method: 'DELETE' }),

    // Categories
    getCategories: () => request<Category[]>('/admin/categories'),
    createCategory: (data: Partial<Category>) =>
      request<Category>('/admin/categories', { method: 'POST', body: JSON.stringify(data) }),
    updateCategory: (id: number, data: Partial<Category>) =>
      request<Category>(`/admin/categories/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
    deleteCategory: (id: number) =>
      request<{ message: string }>(`/admin/categories/${id}`, { method: 'DELETE' }),

    // Education Units
    getEducationUnits: () => request<EducationUnit[]>('/admin/education-units'),
    createEducationUnit: (data: Partial<EducationUnit>) =>
      request<EducationUnit>('/admin/education-units', { method: 'POST', body: JSON.stringify(data) }),
    updateEducationUnit: (id: number, data: Partial<EducationUnit>) =>
      request<EducationUnit>(`/admin/education-units/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
    deleteEducationUnit: (id: number) =>
      request<{ message: string }>(`/admin/education-units/${id}`, { method: 'DELETE' }),

    // Caregivers
    getCaregivers: () => request<Caregiver[]>('/admin/caregivers'),
    createCaregiver: (data: Partial<Caregiver>) =>
      request<Caregiver>('/admin/caregivers', { method: 'POST', body: JSON.stringify(data) }),
    updateCaregiver: (id: number, data: Partial<Caregiver>) =>
      request<Caregiver>(`/admin/caregivers/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
    deleteCaregiver: (id: number) =>
      request<{ message: string }>(`/admin/caregivers/${id}`, { method: 'DELETE' }),

    // Facilities
    getFacilities: () => request<Facility[]>('/admin/facilities'),
    createFacility: (data: Partial<Facility>) =>
      request<Facility>('/admin/facilities', { method: 'POST', body: JSON.stringify(data) }),
    updateFacility: (id: number, data: Partial<Facility>) =>
      request<Facility>(`/admin/facilities/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
    deleteFacility: (id: number) =>
      request<{ message: string }>(`/admin/facilities/${id}`, { method: 'DELETE' }),

    // Achievements
    getAchievements: () => request<Achievement[]>('/admin/achievements'),
    createAchievement: (data: Partial<Achievement>) =>
      request<Achievement>('/admin/achievements', { method: 'POST', body: JSON.stringify(data) }),
    updateAchievement: (id: number, data: Partial<Achievement>) =>
      request<Achievement>(`/admin/achievements/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
    deleteAchievement: (id: number) =>
      request<{ message: string }>(`/admin/achievements/${id}`, { method: 'DELETE' }),

    // Testimonials
    getTestimonials: () => request<Testimonial[]>('/admin/testimonials'),
    createTestimonial: (data: Partial<Testimonial>) =>
      request<Testimonial>('/admin/testimonials', { method: 'POST', body: JSON.stringify(data) }),
    updateTestimonial: (id: number, data: Partial<Testimonial>) =>
      request<Testimonial>(`/admin/testimonials/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
    deleteTestimonial: (id: number) =>
      request<{ message: string }>(`/admin/testimonials/${id}`, { method: 'DELETE' }),

    // Alumni
    getAlumni: () => request<Alumni[]>('/admin/alumni'),
    createAlumni: (data: Partial<Alumni>) =>
      request<Alumni>('/admin/alumni', { method: 'POST', body: JSON.stringify(data) }),
    updateAlumni: (id: number, data: Partial<Alumni>) =>
      request<Alumni>(`/admin/alumni/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
    deleteAlumni: (id: number) =>
      request<{ message: string }>(`/admin/alumni/${id}`, { method: 'DELETE' }),

    // Media Photos
    getPhotos: () => request<MediaPhoto[]>('/admin/photos'),
    createPhoto: (data: Partial<MediaPhoto>) =>
      request<MediaPhoto>('/admin/photos', { method: 'POST', body: JSON.stringify(data) }),
    deletePhoto: (id: number) =>
      request<{ message: string }>(`/admin/photos/${id}`, { method: 'DELETE' }),

    // Media Videos
    getVideos: () => request<MediaVideo[]>('/admin/videos'),
    createVideo: (data: Partial<MediaVideo>) =>
      request<MediaVideo>('/admin/videos', { method: 'POST', body: JSON.stringify(data) }),
    deleteVideo: (id: number) =>
      request<{ message: string }>(`/admin/videos/${id}`, { method: 'DELETE' }),

    // Banners
    getBanners: () => request<Banner[]>('/admin/banners'),
    createBanner: (data: Partial<Banner>) =>
      request<Banner>('/admin/banners', { method: 'POST', body: JSON.stringify(data) }),
    updateBanner: (id: number, data: Partial<Banner>) =>
      request<Banner>(`/admin/banners/${id}`, { method: 'PUT', body: JSON.stringify(data) }),
    deleteBanner: (id: number) =>
      request<{ message: string }>(`/admin/banners/${id}`, { method: 'DELETE' }),

    // Settings
    getSettings: () => request<{ key: string; value: string; category: string }[]>('/admin/settings'),
    updateSettings: (settingsMap: Record<string, string>) =>
      request<{ message: string }>('/admin/settings', {
        method: 'PUT',
        body: JSON.stringify(settingsMap),
      }),

    // Messages
    getMessages: () => request<ContactMessage[]>('/admin/messages'),
    markMessageRead: (id: number) =>
      request<{ message: string }>(`/admin/messages/${id}/read`, { method: 'PUT' }),
    deleteMessage: (id: number) =>
      request<{ message: string }>(`/admin/messages/${id}`, { method: 'DELETE' }),

    // Upload
    uploadFile: async (file: File) => {
      const formData = new FormData();
      formData.append('file', file);
      const token = getToken();
      const res = await fetch(`${API_BASE}/admin/upload`, {
        method: 'POST',
        headers: token ? { Authorization: `Bearer ${token}` } : {},
        body: formData,
      });
      if (!res.ok) throw new Error('Upload file gagal');
      return res.json() as Promise<{ message: string; file_url: string; filename: string }>;
    },
  },
};

