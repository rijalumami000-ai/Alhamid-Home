export interface User {
  id: number;
  username: string;
  email: string;
  full_name: string;
  role: string;
}

export interface Banner {
  id: number;
  title: string;
  subtitle: string;
  image_url: string;
  button_text: string;
  button_link: string;
  badge: string;
  order: number;
  is_active: boolean;
}

export interface Category {
  id: number;
  name: string;
  slug: string;
  type: 'berita' | 'galeri_foto' | 'galeri_video' | 'prestasi';
  description?: string;
}

export interface Article {
  id: number;
  title: string;
  slug: string;
  category_id: number;
  category: Category;
  excerpt: string;
  content: string;
  thumbnail_url: string;
  author: string;
  status: 'published' | 'draft';
  is_featured: boolean;
  view_count: number;
  published_at: string;
  created_at: string;
}

export interface EducationUnit {
  id: number;
  name: string;
  slug: string;
  type: 'formal' | 'non_formal';
  tagline: string;
  hero_image: string;
  badge: string;
  short_desc: string;
  full_desc: string;
  curriculum: string;
  requirements: string;
  facilities_list: string;
  headmaster_name: string;
  headmaster_photo: string;
  order: number;
  is_active: boolean;
}

export interface Caregiver {
  id: number;
  name: string;
  role_title: string;
  photo_url: string;
  bio: string;
  sanad_keilmuan: string;
  order: number;
  is_active: boolean;
}

export interface Facility {
  id: number;
  name: string;
  category: string;
  description: string;
  image_url: string;
  order: number;
}

export interface Achievement {
  id: number;
  title: string;
  year: number;
  level: string;
  category: string;
  recipient_name: string;
  description: string;
  image_url: string;
}

export interface Testimonial {
  id: number;
  name: string;
  role: string;
  content: string;
  avatar_url: string;
  rating: number;
  order: number;
  is_active: boolean;
}

export interface Alumni {
  id: number;
  name: string;
  graduation_year: number;
  education_unit: string;
  profession: string;
  current_location: string;
  story: string;
  photo_url: string;
  social_link: string;
  order: number;
}

export interface MediaPhoto {
  id: number;
  title: string;
  category_id: number;
  category?: Category;
  image_url: string;
  caption: string;
  created_at: string;
}

export interface MediaVideo {
  id: number;
  title: string;
  category_id: number;
  category?: Category;
  youtube_url: string;
  youtube_id: string;
  description: string;
  created_at: string;
}

export interface ContactMessage {
  id: number;
  name: string;
  email: string;
  phone: string;
  subject: string;
  message: string;
  is_read: boolean;
  created_at: string;
}

export interface HomeData {
  banners: Banner[];
  education_units: EducationUnit[];
  featured_articles: Article[];
  caregivers: Caregiver[];
  facilities: Facility[];
  achievements: Achievement[];
  testimonials: Testimonial[];
  recent_photos: MediaPhoto[];
  stats: {
    santri_count: string;
    asatidz_count: string;
    alumni_count: string;
    prestasi_count: string;
  };
}

export interface SearchResult {
  articles: Article[];
  education_units: EducationUnit[];
  caregivers: Caregiver[];
  facilities: Facility[];
}

